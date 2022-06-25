package nodesv1beta1

type Configuration struct {
	Name    string   `json:"name,omitempty" yaml:"name"`
	Network *Network `json:"network,omitempty" yaml:"network"`
}
