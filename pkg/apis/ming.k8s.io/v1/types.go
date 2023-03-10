package v1

import(
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)


// BarSpec defines the desired state of Bar
type BarSpec struct {
	// INSERT ADDITIONAL SPEC FIELDS -- desired state of cluster
	Name string `json:"name,omitempty"`
}


// BarStatus defines the observed state of Bar.
// It should always be reconstructable from the state of the cluster and/or outside world.
type BarStatus struct {
	// INSERT ADDITIONAL STATUS FIELDS -- observed state of cluster
}

// Bar is the Schema for the foobars API
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +k8s:openapi-gen=true

type Bar struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   BarSpec   `json:"spec,omitempty"`
	Status BarStatus `json:"status,omitempty"`
}

// BarList contains a list of Bar
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type BarList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Bar `json:"items"`
}


