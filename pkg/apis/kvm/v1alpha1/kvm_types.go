package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// KvmSpec defines the desired state of Kvm
type KvmSpec struct {
	// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
	// Important: Run "operator-sdk generate k8s" to regenerate code after modifying this file
	// Add custom validation using kubebuilder tags: https://book-v1.book.kubebuilder.io/beyond_basics/generating_crd.html
	Host  string `json:"host"`
	Imagepath string `json:"imagepath"`
	Memory string `json:"memory"`
	VCPU string 	`json:"VCPU"`
	OStype string `json:"OStype"`


}

// KvmStatus defines the observed state of Kvm
type KvmStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "operator-sdk generate k8s" to regenerate code after modifying this file
	// Add custom validation using kubebuilder tags: https://book-v1.book.kubebuilder.io/beyond_basics/generating_crd.html
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// Kvm is the Schema for the kvms API
// +kubebuilder:subresource:status
// +kubebuilder:resource:path=kvms,scope=Namespaced
type Kvm struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   KvmSpec   `json:"spec,omitempty"`
	Status KvmStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// KvmList contains a list of Kvm
type KvmList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Kvm `json:"items"`
}

func init() {
	SchemeBuilder.Register(&Kvm{}, &KvmList{})
}
