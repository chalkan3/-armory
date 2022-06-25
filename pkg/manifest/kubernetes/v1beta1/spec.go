package nodesv1beta1

type Spec struct {
	ClusterName string `json:"cluster_name,omitempty" yaml:"cluster"`
	Node        *Node  `json:"node,omitempty" yaml:"node"`
}
