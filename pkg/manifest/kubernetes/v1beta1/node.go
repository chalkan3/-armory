package nodesv1beta1

type Node struct {
	Network *Network `json:"network,omitempty" yaml:"network"`
	Name    string   `json:"name,omitempty" yaml:"name"`
	Types   string   `json:"types,omitempty" yaml:"types"`
	Primary bool     `json:"primary,omitempty" yaml:"primary"`
}
