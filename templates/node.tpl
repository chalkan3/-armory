
workerName="{{.Values.Spec.Node.Name}}"
workerIp="{{.Values.Spec.Node.Network.PrivateIP}}"
kubernetesHaIp="{{.Values.Spec.Node.Network.LoadBalancerIP}}"

HA_IP = "193.168.50.5"

Vagrant.configure("2") do |config|

    config.vm.define workerName do |node|
        node.vm.provider "docker" do |d|
            d.image = "rofrano/vagrant-provider:ubuntu"
            d.remains_running = true
            d.has_ssh = true
            d.privileged = true
            d.volumes = ["/sys/fs/cgroup:/sys/fs/cgroup:rw", "overlay2_"+workerName+":/var/lib/docker:rw", "/sys/fs/:/sys/fs/:rw"]
            d.create_args = [   
                "--cgroupns=host"
            ]
        end
        node.vm.network "private_network", ip: workerIp
        node.vm.hostname = workerName
        node.vm.provision "ansible" do |ansible|
            ansible.playbook = "{{.AnsiblePath}}"
            ansible.extra_vars = {
                node_ip: workerIp,
                ha_ip: kubernetesHaIp
            }
        end
    end

  
end
