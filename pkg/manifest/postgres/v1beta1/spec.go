package nodesv1beta1

type Spec struct {
	ClusterName   string         `json:"cluster_name,omitempty" yaml:"cluster"`
	Configuration *Configuration `json:"configuration,omitempty" yaml:"configuration"`
}
