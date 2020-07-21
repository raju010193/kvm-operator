package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOUSpec TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.
type Image struct{
	Path string `json:"path"`
	Format string `json:"format"`
	Device string `json:"device"`
    Type string `json:"type"`
}
type Status struct{
	Status string `json:"status"`
}
type Memory  struct{
	Current uint64 `json:"current"`
	Max uint64 `json:"max"`
}
type vcpu struct{
	Current string `json:"current"`
	Max string `json:"max"`
}

type Drive struct{
    Drive string `json:"drive"`
	Type string `json:"type"`
}
 type NetworkInterface struct{
	InterfaceType string `json:"interfaceType"`
    Bridge string `json:"bridge"` 
    Model string `json:"model"` 
    MacAddress string `json:"macAddress"` 
 }
   
// KvmSpec defines the desired state of Kvm
type KvmSpec struct {
	// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
	// Important: Run "operator-sdk generate k8s" to regenerate code after modifying this file
	// Add custom validation using kubebuilder tags: https://book-v1.book.kubebuilder.io/beyond_basics/generating_crd.html
	Host  string `json:"host"`
	Image Image `json:"image"`
	Memory Memory `json:"memory"`
	VCPU vcpu 	`json:"VCPU"`
	OStype string `json:"OStype"`
	Connection string `json:"connection",omitempty`
	NetworkInterface NetworkInterface `json:"interface"`
	Drive Drive `json:"drive"`
	StatusSpec Status `json:"statusSpec"`
	


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
