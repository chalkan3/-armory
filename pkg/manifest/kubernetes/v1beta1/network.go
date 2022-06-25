package nodesv1beta1

type Network struct {
	PrivateIP      string `json:"private_ip,omitempty" yaml:"privateIp"`
	LoadBalancerIP string `json:"load_balancer_ip,omitempty" yaml:"loadBalancerIp"`
}
