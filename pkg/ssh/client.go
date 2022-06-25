package ssh

import (
	"fmt"
	"io"
	"log"
	"net"
	"os"
	"strings"

	s "golang.org/x/crypto/ssh"
	"golang.org/x/crypto/ssh/terminal"
)

type SSH struct {
	c *s.Client
}

func NewSSH() *SSH { return new(SSH) }

func (ssh *SSH) Connect(user, password, host, port string) *SSH {

	conf := &s.ClientConfig{
		User:            user,
		HostKeyCallback: s.InsecureIgnoreHostKey(),
		Auth: []s.AuthMethod{
			s.Password(password),
		},
	}

	conn, err := s.Dial("tcp", "localhost:"+port, conf)
	if err != nil {
		panic(err)
	}

	ssh.c = conn

	return ssh

}

func (ssh *SSH) Interact() {
	session, err := ssh.c.NewSession()
	if err != nil {
		panic(err)
	}
	defer session.Close()
	defer ssh.c.Close()

	fd := int(os.Stdin.Fd())
	state, err := terminal.MakeRaw(fd)
	if err != nil {
		panic(err)
	}
	defer terminal.Restore(fd, state)

	w, h, err := terminal.GetSize(fd)
	if err != nil {
		panic(err)
	}

	modes := s.TerminalModes{
		s.ECHO:          1,
		s.TTY_OP_ISPEED: 14400,
		s.TTY_OP_OSPEED: 14400,
	}

	term := os.Getenv("TERM")
	if term == "" {
		term = "xterm-256color"
	}
	if err := session.RequestPty(term, h, w, modes); err != nil {
		panic(err)
	}

	session.Stdout = os.Stdout
	session.Stderr = os.Stderr
	session.Stdin = os.Stdin

	if err := session.Shell(); err != nil {
		panic(err)
	}

	if err := session.Wait(); err != nil {
		if e, ok := err.(*s.ExitError); ok {
			switch e.ExitStatus() {
			case 130:
				panic(err)
			}
		}
		panic(err)
	}
	panic(err)

}

func (ssh *SSH) PortForward(addr string) {
	splited := strings.Split(addr, ":")
	// Establish connection with remote server
	remote, err := ssh.c.Dial("tcp", fmt.Sprintf("%s:%s", splited[1], splited[2]))
	if err != nil {
		panic(err)
	}

	// Start local server to forward traffic to remote connection
	local, err := net.Listen("tcp", "localhost:"+splited[0])
	if err != nil {
		panic(err)
	}
	defer local.Close()

	// Handle incoming connections
	for {
		client, err := local.Accept()
		if err != nil {
			panic(err)
		}

		ssh.handleClient(client, remote)
	}
}

func (ssh *SSH) handleClient(client net.Conn, remote net.Conn) {
	chDone := make(chan bool)
	// Start remote -> local data transfer
	go func() {
		_, err := io.Copy(client, remote)
		if err != nil {
			log.Println("error while copy remote->local:", err)
		}
		chDone <- true
	}()

	// Start local -> remote data transfer
	go func() {
		_, err := io.Copy(remote, client)
		if err != nil {
			log.Println(err)
		}
		chDone <- true
	}()

	<-chDone
}
