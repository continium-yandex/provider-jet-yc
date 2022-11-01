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

type DatabaseDedicatedObservation struct {
	// The Yandex Database cluster creation timestamp.
	CreatedAt *string `json:"createdAt,omitempty" tf:"created_at,omitempty"`

	// Full database path of the Yandex Database cluster.
	DatabasePath *string `json:"databasePath,omitempty" tf:"database_path,omitempty"`

	// (Required) Region ID for the Yandex Database cluster.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// Status of the Yandex Database cluster.
	Status *string `json:"status,omitempty" tf:"status,omitempty"`

	// Whether TLS is enabled for the Yandex Database cluster.
	TLSEnabled *bool `json:"tlsEnabled,omitempty" tf:"tls_enabled,omitempty"`

	// API endpoint of the Yandex Database cluster.
	YdbAPIEndpoint *string `json:"ydbApiEndpoint,omitempty" tf:"ydb_api_endpoint,omitempty"`

	// Full endpoint of the Yandex Database cluster.
	YdbFullEndpoint *string `json:"ydbFullEndpoint,omitempty" tf:"ydb_full_endpoint,omitempty"`
}

type DatabaseDedicatedParameters struct {

	// +kubebuilder:validation:Optional
	// (Optional) Whether public IP addresses should be assigned to the Yandex Database cluster.
	AssignPublicIps *bool `json:"assignPublicIps,omitempty" tf:"assign_public_ips,omitempty"`

	// +kubebuilder:validation:Optional
	// (Optional) A description for the Yandex Database cluster.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// +crossplane:generate:reference:type=github.com/yandex-cloud/provider-jet-yc/apis/resourcemanager/v1alpha1.Folder
	// +kubebuilder:validation:Optional
	// (Optional) ID of the folder that the Yandex Database cluster belongs to.
	FolderID *string `json:"folderId,omitempty" tf:"folder_id,omitempty"`

	// Reference to a Folder in resourcemanager to populate folderId.
	// +kubebuilder:validation:Optional
	FolderIDRef *v1.Reference `json:"folderIdRef,omitempty" tf:"-"`

	// Selector for a Folder in resourcemanager to populate folderId.
	// +kubebuilder:validation:Optional
	FolderIDSelector *v1.Selector `json:"folderIdSelector,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	// (Optional) A set of key/value label pairs to assign to the Yandex Database cluster.
	Labels map[string]*string `json:"labels,omitempty" tf:"labels,omitempty"`

	// +kubebuilder:validation:Optional
	// (Optional) Location for the Yandex Database cluster.
	Location []LocationParameters `json:"location,omitempty" tf:"location,omitempty"`

	// +kubebuilder:validation:Optional
	// (Optional) Location ID for the Yandex Database cluster.
	LocationID *string `json:"locationId,omitempty" tf:"location_id,omitempty"`

	// +kubebuilder:validation:Required
	// (Required) Name of the Yandex Database cluster.
	Name *string `json:"name" tf:"name,omitempty"`

	// +crossplane:generate:reference:type=github.com/yandex-cloud/provider-jet-yc/apis/vpc/v1alpha1.Network
	// +kubebuilder:validation:Optional
	// (Required) ID of the network to attach the Yandex Database cluster to.
	NetworkID *string `json:"networkId,omitempty" tf:"network_id,omitempty"`

	// Reference to a Network in vpc to populate networkId.
	// +kubebuilder:validation:Optional
	NetworkIDRef *v1.Reference `json:"networkIdRef,omitempty" tf:"-"`

	// Selector for a Network in vpc to populate networkId.
	// +kubebuilder:validation:Optional
	NetworkIDSelector *v1.Selector `json:"networkIdSelector,omitempty" tf:"-"`

	// +kubebuilder:validation:Required
	// (Required) The Yandex Database cluster preset.
	ResourcePresetID *string `json:"resourcePresetId" tf:"resource_preset_id,omitempty"`

	// +kubebuilder:validation:Required
	// (Required) Scaling policy for the Yandex Database cluster.
	ScalePolicy []ScalePolicyParameters `json:"scalePolicy" tf:"scale_policy,omitempty"`

	// +kubebuilder:validation:Required
	// (Required) A list of storage configuration options for the Yandex Database cluster.
	StorageConfig []StorageConfigParameters `json:"storageConfig" tf:"storage_config,omitempty"`

	// +crossplane:generate:reference:type=github.com/yandex-cloud/provider-jet-yc/apis/vpc/v1alpha1.Subnet
	// +kubebuilder:validation:Optional
	// (Required) List of subnet IDs to attach the Yandex Database cluster to.
	SubnetIds []*string `json:"subnetIds,omitempty" tf:"subnet_ids,omitempty"`

	// References to Subnet in vpc to populate subnetIds.
	// +kubebuilder:validation:Optional
	SubnetIdsRefs []v1.Reference `json:"subnetIdsRefs,omitempty" tf:"-"`

	// Selector for a list of Subnet in vpc to populate subnetIds.
	// +kubebuilder:validation:Optional
	SubnetIdsSelector *v1.Selector `json:"subnetIdsSelector,omitempty" tf:"-"`
}

type FixedScaleObservation struct {
}

type FixedScaleParameters struct {

	// +kubebuilder:validation:Required
	// (Required) Number of instances for the Yandex Database cluster.
	Size *float64 `json:"size" tf:"size,omitempty"`
}

type LocationObservation struct {
}

type LocationParameters struct {

	// +kubebuilder:validation:Optional
	// (Optional) Region for the Yandex Database cluster.
	Region []RegionParameters `json:"region,omitempty" tf:"region,omitempty"`
}

type RegionObservation struct {
}

type RegionParameters struct {

	// +kubebuilder:validation:Required
	// (Required) Region ID for the Yandex Database cluster.
	ID *string `json:"id" tf:"id,omitempty"`
}

type ScalePolicyObservation struct {
}

type ScalePolicyParameters struct {

	// +kubebuilder:validation:Required
	// (Required) Fixed scaling policy for the Yandex Database cluster.
	FixedScale []FixedScaleParameters `json:"fixedScale" tf:"fixed_scale,omitempty"`
}

type StorageConfigObservation struct {
}

type StorageConfigParameters struct {

	// +kubebuilder:validation:Required
	// (Required) Amount of storage groups of selected type for the Yandex Database cluster.
	GroupCount *float64 `json:"groupCount" tf:"group_count,omitempty"`

	// +kubebuilder:validation:Required
	// (Required) Storage type ID for the Yandex Database cluster.
	StorageTypeID *string `json:"storageTypeId" tf:"storage_type_id,omitempty"`
}

// DatabaseDedicatedSpec defines the desired state of DatabaseDedicated
type DatabaseDedicatedSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     DatabaseDedicatedParameters `json:"forProvider"`
}

// DatabaseDedicatedStatus defines the observed state of DatabaseDedicated.
type DatabaseDedicatedStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        DatabaseDedicatedObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// DatabaseDedicated is the Schema for the DatabaseDedicateds API. <no value>
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,yandex-cloud}
type DatabaseDedicated struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              DatabaseDedicatedSpec   `json:"spec"`
	Status            DatabaseDedicatedStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// DatabaseDedicatedList contains a list of DatabaseDedicateds
type DatabaseDedicatedList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []DatabaseDedicated `json:"items"`
}

// Repository type metadata.
var (
	DatabaseDedicated_Kind             = "DatabaseDedicated"
	DatabaseDedicated_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: DatabaseDedicated_Kind}.String()
	DatabaseDedicated_KindAPIVersion   = DatabaseDedicated_Kind + "." + CRDGroupVersion.String()
	DatabaseDedicated_GroupVersionKind = CRDGroupVersion.WithKind(DatabaseDedicated_Kind)
)

func init() {
	SchemeBuilder.Register(&DatabaseDedicated{}, &DatabaseDedicatedList{})
}
