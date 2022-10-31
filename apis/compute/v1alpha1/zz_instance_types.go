/*
Copyright 2021 The Crossplane Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Code generated by upjet. DO NOT EDIT.

package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	v1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
)

type BootDiskObservation struct {
}

type BootDiskParameters struct {

	// +kubebuilder:validation:Optional
	// (Optional) Defines whether the disk will be auto-deleted when the instance
	AutoDelete *bool `json:"autoDelete,omitempty" tf:"auto_delete,omitempty"`

	// +kubebuilder:validation:Optional
	// (Optional) Name that can be used to access an attached disk.
	DeviceName *string `json:"deviceName,omitempty" tf:"device_name,omitempty"`

	// +kubebuilder:validation:Optional
	// (Optional) The ID of the existing disk (such as those managed by
	DiskID *string `json:"diskId,omitempty" tf:"disk_id,omitempty"`

	// +kubebuilder:validation:Optional
	// (Optional) Parameters for a new disk that will be created
	InitializeParams []InitializeParamsParameters `json:"initializeParams,omitempty" tf:"initialize_params,omitempty"`

	// +kubebuilder:validation:Optional
	// (Optional) Type of access to the disk resource. By default, a disk is attached in `READ_WRITE` mode.
	Mode *string `json:"mode,omitempty" tf:"mode,omitempty"`
}

type DNSRecordObservation struct {
}

type DNSRecordParameters struct {

	// +kubebuilder:validation:Optional
	// (Optional) DNS zone ID (if not set, private zone used).
	DNSZoneID *string `json:"dnsZoneId,omitempty" tf:"dns_zone_id,omitempty"`

	// +kubebuilder:validation:Required
	// (Required) DNS record FQDN (must have a dot at the end).
	Fqdn *string `json:"fqdn" tf:"fqdn,omitempty"`

	// +kubebuilder:validation:Optional
	// (Optional) When set to true, also create a PTR DNS record.
	Ptr *bool `json:"ptr,omitempty" tf:"ptr,omitempty"`

	// +kubebuilder:validation:Optional
	// (Optional) DNS record TTL. in seconds
	TTL *float64 `json:"ttl,omitempty" tf:"ttl,omitempty"`
}

type HostAffinityRulesObservation struct {
}

type HostAffinityRulesParameters struct {

	// +kubebuilder:validation:Optional
	// (Required) Affinity label or one of reserved values - `yc.hostId`, `yc.hostGroupId`.
	Key *string `json:"key,omitempty" tf:"key"`

	// +kubebuilder:validation:Optional
	// (Required) Affinity action. The only value supported is `IN`.
	Op *string `json:"op,omitempty" tf:"op"`

	// +kubebuilder:validation:Optional
	Values []*string `json:"values,omitempty" tf:"values"`
}

type IPv6DNSRecordObservation struct {
}

type IPv6DNSRecordParameters struct {

	// +kubebuilder:validation:Optional
	// (Optional) DNS zone ID (if not set, private zone used).
	DNSZoneID *string `json:"dnsZoneId,omitempty" tf:"dns_zone_id,omitempty"`

	// +kubebuilder:validation:Required
	// (Required) DNS record FQDN (must have a dot at the end).
	Fqdn *string `json:"fqdn" tf:"fqdn,omitempty"`

	// +kubebuilder:validation:Optional
	// (Optional) When set to true, also create a PTR DNS record.
	Ptr *bool `json:"ptr,omitempty" tf:"ptr,omitempty"`

	// +kubebuilder:validation:Optional
	// (Optional) DNS record TTL. in seconds
	TTL *float64 `json:"ttl,omitempty" tf:"ttl,omitempty"`
}

type InitializeParamsObservation struct {
}

type InitializeParamsParameters struct {

	// +kubebuilder:validation:Optional
	BlockSize *float64 `json:"blockSize,omitempty" tf:"block_size,omitempty"`

	// +kubebuilder:validation:Optional
	// (Optional) Description of the instance.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// +kubebuilder:validation:Optional
	// (Optional) A disk image to initialize this disk from.
	ImageID *string `json:"imageId,omitempty" tf:"image_id,omitempty"`

	// +kubebuilder:validation:Optional
	// (Optional) Resource name.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// +kubebuilder:validation:Optional
	// (Optional) Size of the disk in GB.
	Size *float64 `json:"size,omitempty" tf:"size,omitempty"`

	// +kubebuilder:validation:Optional
	// (Optional) A snapshot to initialize this disk from.
	SnapshotID *string `json:"snapshotId,omitempty" tf:"snapshot_id,omitempty"`

	// +kubebuilder:validation:Optional
	// (Optional) Disk type.
	Type *string `json:"type,omitempty" tf:"type,omitempty"`
}

type InstanceObservation struct {
	// Creation timestamp of the instance.
	CreatedAt *string `json:"createdAt,omitempty" tf:"created_at,omitempty"`

	// (Required) DNS record FQDN (must have a dot at the end).
	Fqdn *string `json:"fqdn,omitempty" tf:"fqdn,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// +kubebuilder:validation:Optional
	// (Optional) List of local disks that are attached to the instance. Structure is documented below.
	LocalDisk []LocalDiskObservation `json:"localDisk,omitempty" tf:"local_disk,omitempty"`

	// +kubebuilder:validation:Required
	// (Required) Networks to attach to the instance. This can
	NetworkInterface []NetworkInterfaceObservation `json:"networkInterface,omitempty" tf:"network_interface,omitempty"`

	// The status of this instance.
	Status *string `json:"status,omitempty" tf:"status,omitempty"`
}

type InstanceParameters struct {

	// +kubebuilder:validation:Optional
	AllowRecreate *bool `json:"allowRecreate,omitempty" tf:"allow_recreate,omitempty"`

	// +kubebuilder:validation:Optional
	// (Optional) If true, allows Terraform to stop the instance in order to update its properties.
	AllowStoppingForUpdate *bool `json:"allowStoppingForUpdate,omitempty" tf:"allow_stopping_for_update,omitempty"`

	// +kubebuilder:validation:Required
	// (Required) The boot disk for the instance. The structure is documented below.
	BootDisk []BootDiskParameters `json:"bootDisk" tf:"boot_disk,omitempty"`

	// +kubebuilder:validation:Optional
	// (Optional) Description of the instance.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// +kubebuilder:validation:Optional
	// (Optional) The ID of the folder that the resource belongs to. If it
	FolderID *string `json:"folderId,omitempty" tf:"folder_id,omitempty"`

	// +kubebuilder:validation:Optional
	// (Optional) Host name for the instance. This field is used to generate the instance `fqdn` value. 
	Hostname *string `json:"hostname,omitempty" tf:"hostname,omitempty"`

	// +kubebuilder:validation:Optional
	// (Optional) A set of key/value label pairs to assign to the instance.
	Labels map[string]*string `json:"labels,omitempty" tf:"labels,omitempty"`

	// +kubebuilder:validation:Optional
	// (Optional) List of local disks that are attached to the instance. Structure is documented below.
	LocalDisk []LocalDiskParameters `json:"localDisk,omitempty" tf:"local_disk,omitempty"`

	// +kubebuilder:validation:Optional
	// (Optional) Metadata key/value pairs to make available from
	Metadata map[string]*string `json:"metadata,omitempty" tf:"metadata,omitempty"`

	// +kubebuilder:validation:Optional
	// (Optional) Type of network acceleration. The default is `standard`. Values: `standard`, `software_accelerated`
	NetworkAccelerationType *string `json:"networkAccelerationType,omitempty" tf:"network_acceleration_type,omitempty"`

	// +kubebuilder:validation:Required
	// (Required) Networks to attach to the instance. This can
	NetworkInterface []NetworkInterfaceParameters `json:"networkInterface" tf:"network_interface,omitempty"`

	// +kubebuilder:validation:Optional
	// (Optional) The placement policy configuration. The structure is documented below.
	PlacementPolicy []PlacementPolicyParameters `json:"placementPolicy,omitempty" tf:"placement_policy,omitempty"`

	// +kubebuilder:validation:Optional
	// (Optional) The type of virtual machine to create. The default is 'standard-v1'.
	PlatformID *string `json:"platformId,omitempty" tf:"platform_id,omitempty"`

	// +kubebuilder:validation:Required
	// (Required) Compute resources that are allocated for the instance. The structure is documented below.
	Resources []ResourcesParameters `json:"resources" tf:"resources,omitempty"`

	// +kubebuilder:validation:Optional
	// (Optional) Scheduling policy configuration. The structure is documented below.
	SchedulingPolicy []SchedulingPolicyParameters `json:"schedulingPolicy,omitempty" tf:"scheduling_policy,omitempty"`

	// +kubebuilder:validation:Optional
	// (Optional) A list of disks to attach to the instance. The structure is documented below.
	SecondaryDisk []SecondaryDiskParameters `json:"secondaryDisk,omitempty" tf:"secondary_disk,omitempty"`

	// +crossplane:generate:reference:type=github.com/yandex-cloud/provider-jet-yc/apis/iam/v1alpha1.ServiceAccount
	// +kubebuilder:validation:Optional
	// (Optional) ID of the service account authorized for this instance.
	ServiceAccountID *string `json:"serviceAccountId,omitempty" tf:"service_account_id,omitempty"`

	// Reference to a ServiceAccount in iam to populate serviceAccountId.
	// +kubebuilder:validation:Optional
	ServiceAccountIDRef *v1.Reference `json:"serviceAccountIdRef,omitempty" tf:"-"`

	// Selector for a ServiceAccount in iam to populate serviceAccountId.
	// +kubebuilder:validation:Optional
	ServiceAccountIDSelector *v1.Selector `json:"serviceAccountIdSelector,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	// (Optional) The availability zone where the virtual machine will be created. If it is not provided,
	Zone *string `json:"zone,omitempty" tf:"zone,omitempty"`
}

type LocalDiskObservation struct {
	// (Optional) Name that can be used to access an attached disk.
	DeviceName *string `json:"deviceName,omitempty" tf:"device_name,omitempty"`
}

type LocalDiskParameters struct {

	// +kubebuilder:validation:Required
	// (Required) Size of the disk, specified in bytes.
	SizeBytes *float64 `json:"sizeBytes" tf:"size_bytes,omitempty"`
}

type NATDNSRecordObservation struct {
}

type NATDNSRecordParameters struct {

	// +kubebuilder:validation:Optional
	// (Optional) DNS zone ID (if not set, private zone used).
	DNSZoneID *string `json:"dnsZoneId,omitempty" tf:"dns_zone_id,omitempty"`

	// +kubebuilder:validation:Required
	// (Required) DNS record FQDN (must have a dot at the end).
	Fqdn *string `json:"fqdn" tf:"fqdn,omitempty"`

	// +kubebuilder:validation:Optional
	// (Optional) When set to true, also create a PTR DNS record.
	Ptr *bool `json:"ptr,omitempty" tf:"ptr,omitempty"`

	// +kubebuilder:validation:Optional
	// (Optional) DNS record TTL. in seconds
	TTL *float64 `json:"ttl,omitempty" tf:"ttl,omitempty"`
}

type NetworkInterfaceObservation struct {
	Index *float64 `json:"index,omitempty" tf:"index,omitempty"`

	MacAddress *string `json:"macAddress,omitempty" tf:"mac_address,omitempty"`

	NATIPVersion *string `json:"natIpVersion,omitempty" tf:"nat_ip_version,omitempty"`
}

type NetworkInterfaceParameters struct {

	// +kubebuilder:validation:Optional
	// (Optional) List of configurations for creating ipv4 DNS records. The structure is documented below.
	DNSRecord []DNSRecordParameters `json:"dnsRecord,omitempty" tf:"dns_record,omitempty"`

	// +kubebuilder:validation:Optional
	// (Optional) The private IP address to assign to the instance. If
	IPAddress *string `json:"ipAddress,omitempty" tf:"ip_address,omitempty"`

	// +kubebuilder:validation:Optional
	// (Optional) Allocate an IPv4 address for the interface. The default value is `true`.
	IPv4 *bool `json:"ipv4,omitempty" tf:"ipv4,omitempty"`

	// +kubebuilder:validation:Optional
	// (Optional) If true, allocate an IPv6 address for the interface.
	IPv6 *bool `json:"ipv6,omitempty" tf:"ipv6,omitempty"`

	// +kubebuilder:validation:Optional
	// (Optional) The private IPv6 address to assign to the instance.
	IPv6Address *string `json:"ipv6Address,omitempty" tf:"ipv6_address,omitempty"`

	// +kubebuilder:validation:Optional
	// (Optional) List of configurations for creating ipv6 DNS records. The structure is documented below.
	IPv6DNSRecord []IPv6DNSRecordParameters `json:"ipv6DnsRecord,omitempty" tf:"ipv6_dns_record,omitempty"`

	// +kubebuilder:validation:Optional
	// (Optional) Provide a public address, for instance, to access the internet over NAT.
	NAT *bool `json:"nat,omitempty" tf:"nat,omitempty"`

	// +kubebuilder:validation:Optional
	// (Optional) List of configurations for creating ipv4 NAT DNS records. The structure is documented below.
	NATDNSRecord []NATDNSRecordParameters `json:"natDnsRecord,omitempty" tf:"nat_dns_record,omitempty"`

	// +kubebuilder:validation:Optional
	// (Optional) Provide a public address, for instance, to access the internet over NAT. Address should be already reserved in web UI.
	NATIPAddress *string `json:"natIpAddress,omitempty" tf:"nat_ip_address,omitempty"`

	// +crossplane:generate:reference:type=github.com/yandex-cloud/provider-jet-yc/apis/vpc/v1alpha1.SecurityGroup
	// +kubebuilder:validation:Optional
	// (Optional) Security group ids for network interface.
	SecurityGroupIds []*string `json:"securityGroupIds,omitempty" tf:"security_group_ids,omitempty"`

	// References to SecurityGroup in vpc to populate securityGroupIds.
	// +kubebuilder:validation:Optional
	SecurityGroupIdsRefs []v1.Reference `json:"securityGroupIdsRefs,omitempty" tf:"-"`

	// Selector for a list of SecurityGroup in vpc to populate securityGroupIds.
	// +kubebuilder:validation:Optional
	SecurityGroupIdsSelector *v1.Selector `json:"securityGroupIdsSelector,omitempty" tf:"-"`

	// +crossplane:generate:reference:type=github.com/yandex-cloud/provider-jet-yc/apis/vpc/v1alpha1.Subnet
	// +kubebuilder:validation:Optional
	// (Required) ID of the subnet to attach this
	SubnetID *string `json:"subnetId,omitempty" tf:"subnet_id,omitempty"`

	// Reference to a Subnet in vpc to populate subnetId.
	// +kubebuilder:validation:Optional
	SubnetIDRef *v1.Reference `json:"subnetIdRef,omitempty" tf:"-"`

	// Selector for a Subnet in vpc to populate subnetId.
	// +kubebuilder:validation:Optional
	SubnetIDSelector *v1.Selector `json:"subnetIdSelector,omitempty" tf:"-"`
}

type PlacementPolicyObservation struct {
}

type PlacementPolicyParameters struct {

	// +kubebuilder:validation:Optional
	// (Optional) List of host affinity rules. The structure is documented below.
	HostAffinityRules []HostAffinityRulesParameters `json:"hostAffinityRules,omitempty" tf:"host_affinity_rules,omitempty"`

	// +kubebuilder:validation:Optional
	// (Optional) Specifies the id of the Placement Group to assign to the instance.
	PlacementGroupID *string `json:"placementGroupId,omitempty" tf:"placement_group_id,omitempty"`
}

type ResourcesObservation struct {
}

type ResourcesParameters struct {

	// +kubebuilder:validation:Optional
	// (Optional) If provided, specifies baseline performance for a core as a percent.
	CoreFraction *float64 `json:"coreFraction,omitempty" tf:"core_fraction,omitempty"`

	// +kubebuilder:validation:Required
	// (Required) CPU cores for the instance.
	Cores *float64 `json:"cores" tf:"cores,omitempty"`

	// +kubebuilder:validation:Optional
	Gpus *float64 `json:"gpus,omitempty" tf:"gpus,omitempty"`

	// +kubebuilder:validation:Required
	// (Required) Memory size in GB.
	Memory *float64 `json:"memory" tf:"memory,omitempty"`
}

type SchedulingPolicyObservation struct {
}

type SchedulingPolicyParameters struct {

	// +kubebuilder:validation:Optional
	// (Optional) Specifies if the instance is preemptible. Defaults to false.
	Preemptible *bool `json:"preemptible,omitempty" tf:"preemptible,omitempty"`
}

type SecondaryDiskObservation struct {
}

type SecondaryDiskParameters struct {

	// +kubebuilder:validation:Optional
	// (Optional) Defines whether the disk will be auto-deleted when the instance
	AutoDelete *bool `json:"autoDelete,omitempty" tf:"auto_delete,omitempty"`

	// +kubebuilder:validation:Optional
	// (Optional) Name that can be used to access an attached disk.
	DeviceName *string `json:"deviceName,omitempty" tf:"device_name,omitempty"`

	// +kubebuilder:validation:Required
	// (Optional) The ID of the existing disk (such as those managed by
	DiskID *string `json:"diskId" tf:"disk_id,omitempty"`

	// +kubebuilder:validation:Optional
	// (Optional) Type of access to the disk resource. By default, a disk is attached in `READ_WRITE` mode.
	Mode *string `json:"mode,omitempty" tf:"mode,omitempty"`
}

// InstanceSpec defines the desired state of Instance
type InstanceSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     InstanceParameters `json:"forProvider"`
}

// InstanceStatus defines the observed state of Instance.
type InstanceStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        InstanceObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// Instance is the Schema for the Instances API. <no value>
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,yandex-cloud}
type Instance struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              InstanceSpec   `json:"spec"`
	Status            InstanceStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// InstanceList contains a list of Instances
type InstanceList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Instance `json:"items"`
}

// Repository type metadata.
var (
	Instance_Kind             = "Instance"
	Instance_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Instance_Kind}.String()
	Instance_KindAPIVersion   = Instance_Kind + "." + CRDGroupVersion.String()
	Instance_GroupVersionKind = CRDGroupVersion.WithKind(Instance_Kind)
)

func init() {
	SchemeBuilder.Register(&Instance{}, &InstanceList{})
}
