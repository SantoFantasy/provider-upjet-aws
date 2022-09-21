/*
Copyright 2022 Upbound Inc.
*/

// Code generated by upjet. DO NOT EDIT.

package v1beta1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	v1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
)

type VPCDHCPOptionsObservation struct {

	// The ARN of the DHCP Options Set.
	Arn *string `json:"arn,omitempty" tf:"arn,omitempty"`

	// The ID of the DHCP Options Set.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The ID of the AWS account that owns the DHCP options set.
	OwnerID *string `json:"ownerId,omitempty" tf:"owner_id,omitempty"`

	// A map of tags assigned to the resource, including those inherited from the provider default_tags configuration block.
	TagsAll map[string]*string `json:"tagsAll,omitempty" tf:"tags_all,omitempty"`
}

type VPCDHCPOptionsParameters struct {

	// the suffix domain name to use by default when resolving non Fully Qualified Domain Names. In other words, this is what ends up being the search value in the /etc/resolv.conf file.
	// +kubebuilder:validation:Optional
	DomainName *string `json:"domainName,omitempty" tf:"domain_name,omitempty"`

	// +kubebuilder:validation:Optional
	DomainNameServers []*string `json:"domainNameServers,omitempty" tf:"domain_name_servers,omitempty"`

	// List of NETBIOS name servers.
	// +kubebuilder:validation:Optional
	NetbiosNameServers []*string `json:"netbiosNameServers,omitempty" tf:"netbios_name_servers,omitempty"`

	// The NetBIOS node type (1, 2, 4, or 8). AWS recommends to specify 2 since broadcast and multicast are not supported in their network. For more information about these node types, see RFC 2132.
	// +kubebuilder:validation:Optional
	NetbiosNodeType *string `json:"netbiosNodeType,omitempty" tf:"netbios_node_type,omitempty"`

	// List of NTP servers to configure.
	// +kubebuilder:validation:Optional
	NtpServers []*string `json:"ntpServers,omitempty" tf:"ntp_servers,omitempty"`

	// Region is the region you'd like your resource to be created in.
	// +terrajet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`

	// A map of tags to assign to the resource. If configured with a provider default_tags configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	// +kubebuilder:validation:Optional
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`
}

// VPCDHCPOptionsSpec defines the desired state of VPCDHCPOptions
type VPCDHCPOptionsSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     VPCDHCPOptionsParameters `json:"forProvider"`
}

// VPCDHCPOptionsStatus defines the observed state of VPCDHCPOptions.
type VPCDHCPOptionsStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        VPCDHCPOptionsObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// VPCDHCPOptions is the Schema for the VPCDHCPOptionss API. Provides a VPC DHCP Options resource.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type VPCDHCPOptions struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              VPCDHCPOptionsSpec   `json:"spec"`
	Status            VPCDHCPOptionsStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// VPCDHCPOptionsList contains a list of VPCDHCPOptionss
type VPCDHCPOptionsList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []VPCDHCPOptions `json:"items"`
}

// Repository type metadata.
var (
	VPCDHCPOptions_Kind             = "VPCDHCPOptions"
	VPCDHCPOptions_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: VPCDHCPOptions_Kind}.String()
	VPCDHCPOptions_KindAPIVersion   = VPCDHCPOptions_Kind + "." + CRDGroupVersion.String()
	VPCDHCPOptions_GroupVersionKind = CRDGroupVersion.WithKind(VPCDHCPOptions_Kind)
)

func init() {
	SchemeBuilder.Register(&VPCDHCPOptions{}, &VPCDHCPOptionsList{})
}