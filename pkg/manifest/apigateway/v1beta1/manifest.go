package v1beta1

type Kong struct {
	ApiVersion string    `json:"api_version,omitempty" yaml:"apiVersion"`
	Kind       string    `json:"kind,omitempty" yaml:"kind"`
	Metadata   *Metadata `json:"metadata,omitempty" yaml:"metadata"`
	Spec       *Spec     `json:"spec,omitempty" yaml:"spec"`
}
