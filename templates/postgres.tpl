
name="{{.Values.Spec.Configuration.Name}}"
ip="{{.Values.Spec.Configuration.Network.PrivateIP}}"


Vagrant.configure("2") do |config|

    config.vm.define name do |node|
        node.vm.provider "docker" do |d|
            d.image = "rofrano/vagrant-provider:ubuntu"
            d.remains_running = true
            d.has_ssh = true
            d.privileged = true
            d.volumes = ["/sys/fs/cgroup:/sys/fs/cgroup:rw", "overlay2_"+name+":/var/lib/docker:rw", "/sys/fs/:/sys/fs/:rw"]
            d.create_args = [   
                "--cgroupns=host"
            ]
        end
        node.vm.network "private_network", ip: ip
        node.vm.hostname = name
        node.vm.provision "ansible" do |ansible|
            ansible.playbook = "{{.AnsiblePath}}"
        end
    end

  
end
