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

type ConfigObservation struct {
}

type ConfigParameters struct {

	// +kubebuilder:validation:Required
	AdminPasswordSecretRef v1.SecretKeySelector `json:"adminPasswordSecretRef" tf:"-"`

	// +kubebuilder:validation:Required
	// (Required) Configuration for Elasticsearch data nodes subcluster. The structure is documented below.
	DataNode []DataNodeParameters `json:"dataNode" tf:"data_node,omitempty"`

	// +kubebuilder:validation:Optional
	// (Optional) Edition of Elasticsearch. For more information, see [the official documentation](https://cloud.yandex.com/en-ru/docs/managed-elasticsearch/concepts/es-editions).
	Edition *string `json:"edition,omitempty" tf:"edition,omitempty"`

	// +kubebuilder:validation:Optional
	// (Optional) Configuration for Elasticsearch master nodes subcluster. The structure is documented below.
	MasterNode []MasterNodeParameters `json:"masterNode,omitempty" tf:"master_node,omitempty"`

	// +kubebuilder:validation:Optional
	// (Optional) A set of Elasticsearch plugins to install.
	Plugins []*string `json:"plugins,omitempty" tf:"plugins,omitempty"`

	// +kubebuilder:validation:Optional
	// (Optional) Version of Elasticsearch.
	Version *string `json:"version,omitempty" tf:"version,omitempty"`
}

type DataNodeObservation struct {
}

type DataNodeParameters struct {

	// +kubebuilder:validation:Required
	// (Required) Resources allocated to hosts of the Elasticsearch data nodes subcluster. The structure is documented below.
	Resources []ResourcesParameters `json:"resources" tf:"resources,omitempty"`
}

type ElasticsearchClusterObservation struct {
	// Creation timestamp of the key.
	CreatedAt *string `json:"createdAt,omitempty" tf:"created_at,omitempty"`

	// Aggregated health of the cluster. Can be either `ALIVE`, `DEGRADED`, `DEAD` or `HEALTH_UNKNOWN`.
	Health *string `json:"health,omitempty" tf:"health,omitempty"`

	// +kubebuilder:validation:Optional
	// (Required) A host of the Elasticsearch cluster. The structure is documented below.
	Host []HostObservation `json:"host,omitempty" tf:"host,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// Status of the cluster. Can be either `CREATING`, `STARTING`, `RUNNING`, `UPDATING`, `STOPPING`, `STOPPED`, `ERROR` or `STATUS_UNKNOWN`.
	Status *string `json:"status,omitempty" tf:"status,omitempty"`
}

type ElasticsearchClusterParameters struct {

	// +kubebuilder:validation:Required
	// (Required) Configuration of the Elasticsearch cluster. The structure is documented below.
	Config []ConfigParameters `json:"config" tf:"config,omitempty"`

	// +kubebuilder:validation:Optional
	// (Optional) Inhibits deletion of the cluster.  Can be either `true` or `false`.
	DeletionProtection *bool `json:"deletionProtection,omitempty" tf:"deletion_protection,omitempty"`

	// +kubebuilder:validation:Optional
	// (Optional) Description of the Elasticsearch cluster.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// +kubebuilder:validation:Required
	// (Required) Deployment environment of the Elasticsearch cluster. Can be either `PRESTABLE` or `PRODUCTION`.
	Environment *string `json:"environment" tf:"environment,omitempty"`

	// +crossplane:generate:reference:type=github.com/yandex-cloud/provider-jet-yc/apis/resourcemanager/v1alpha1.Folder
	// +kubebuilder:validation:Optional
	// (Optional) The ID of the folder that the resource belongs to. If it is not provided, the default provider folder is used.
	FolderID *string `json:"folderId,omitempty" tf:"folder_id,omitempty"`

	// Reference to a Folder in resourcemanager to populate folderId.
	// +kubebuilder:validation:Optional
	FolderIDRef *v1.Reference `json:"folderIdRef,omitempty" tf:"-"`

	// Selector for a Folder in resourcemanager to populate folderId.
	// +kubebuilder:validation:Optional
	FolderIDSelector *v1.Selector `json:"folderIdSelector,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	// (Required) A host of the Elasticsearch cluster. The structure is documented below.
	Host []HostParameters `json:"host,omitempty" tf:"host,omitempty"`

	// +kubebuilder:validation:Optional
	// (Optional) A set of key/value label pairs to assign to the Elasticsearch cluster.
	Labels map[string]*string `json:"labels,omitempty" tf:"labels,omitempty"`

	// +kubebuilder:validation:Optional
	MaintenanceWindow []MaintenanceWindowParameters `json:"maintenanceWindow,omitempty" tf:"maintenance_window,omitempty"`

	// +kubebuilder:validation:Required
	// (Required) Name of the Elasticsearch cluster. Provided by the client when the cluster is created.
	Name *string `json:"name" tf:"name,omitempty"`

	// +crossplane:generate:reference:type=github.com/yandex-cloud/provider-jet-yc/apis/vpc/v1alpha1.Network
	// +kubebuilder:validation:Optional
	// (Required) ID of the network, to which the Elasticsearch cluster belongs.
	NetworkID *string `json:"networkId,omitempty" tf:"network_id,omitempty"`

	// Reference to a Network in vpc to populate networkId.
	// +kubebuilder:validation:Optional
	NetworkIDRef *v1.Reference `json:"networkIdRef,omitempty" tf:"-"`

	// Selector for a Network in vpc to populate networkId.
	// +kubebuilder:validation:Optional
	NetworkIDSelector *v1.Selector `json:"networkIdSelector,omitempty" tf:"-"`

	// +crossplane:generate:reference:type=github.com/yandex-cloud/provider-jet-yc/apis/vpc/v1alpha1.SecurityGroup
	// +kubebuilder:validation:Optional
	// (Optional) A set of ids of security groups assigned to hosts of the cluster.
	SecurityGroupIds []*string `json:"securityGroupIds,omitempty" tf:"security_group_ids,omitempty"`

	// References to SecurityGroup in vpc to populate securityGroupIds.
	// +kubebuilder:validation:Optional
	SecurityGroupIdsRefs []v1.Reference `json:"securityGroupIdsRefs,omitempty" tf:"-"`

	// Selector for a list of SecurityGroup in vpc to populate securityGroupIds.
	// +kubebuilder:validation:Optional
	SecurityGroupIdsSelector *v1.Selector `json:"securityGroupIdsSelector,omitempty" tf:"-"`

	// +crossplane:generate:reference:type=github.com/yandex-cloud/provider-jet-yc/apis/iam/v1alpha1.ServiceAccount
	// +kubebuilder:validation:Optional
	// (Optional) ID of the service account authorized for this cluster.
	ServiceAccountID *string `json:"serviceAccountId,omitempty" tf:"service_account_id,omitempty"`

	// Reference to a ServiceAccount in iam to populate serviceAccountId.
	// +kubebuilder:validation:Optional
	ServiceAccountIDRef *v1.Reference `json:"serviceAccountIdRef,omitempty" tf:"-"`

	// Selector for a ServiceAccount in iam to populate serviceAccountId.
	// +kubebuilder:validation:Optional
	ServiceAccountIDSelector *v1.Selector `json:"serviceAccountIdSelector,omitempty" tf:"-"`
}

type HostObservation struct {
	Fqdn *string `json:"fqdn,omitempty" tf:"fqdn,omitempty"`
}

type HostParameters struct {

	// +kubebuilder:validation:Optional
	AssignPublicIP *bool `json:"assignPublicIp,omitempty" tf:"assign_public_ip,omitempty"`

	// +kubebuilder:validation:Required
	// (Required) Name of the Elasticsearch cluster. Provided by the client when the cluster is created.
	Name *string `json:"name" tf:"name,omitempty"`

	// +crossplane:generate:reference:type=github.com/yandex-cloud/provider-jet-yc/apis/vpc/v1alpha1.Subnet
	// +kubebuilder:validation:Optional
	SubnetID *string `json:"subnetId,omitempty" tf:"subnet_id,omitempty"`

	// Reference to a Subnet in vpc to populate subnetId.
	// +kubebuilder:validation:Optional
	SubnetIDRef *v1.Reference `json:"subnetIdRef,omitempty" tf:"-"`

	// Selector for a Subnet in vpc to populate subnetId.
	// +kubebuilder:validation:Optional
	SubnetIDSelector *v1.Selector `json:"subnetIdSelector,omitempty" tf:"-"`

	// +kubebuilder:validation:Required
	// (Required) The type of the host to be deployed. Can be either `DATA_NODE` or `MASTER_NODE`.
	Type *string `json:"type" tf:"type,omitempty"`

	// +kubebuilder:validation:Required
	// (Required) The availability zone where the Elasticsearch host will be created.
	Zone *string `json:"zone" tf:"zone,omitempty"`
}

type MaintenanceWindowObservation struct {
}

type MaintenanceWindowParameters struct {

	// +kubebuilder:validation:Optional
	// (Optional) Day of week for maintenance window if window type is weekly. Possible values: `MON`, `TUE`, `WED`, `THU`, `FRI`, `SAT`, `SUN`.
	Day *string `json:"day,omitempty" tf:"day,omitempty"`

	// +kubebuilder:validation:Optional
	// (Optional) Hour of day in UTC time zone (1-24) for maintenance window if window type is weekly.
	Hour *float64 `json:"hour,omitempty" tf:"hour,omitempty"`

	// +kubebuilder:validation:Required
	// (Required) The type of the host to be deployed. Can be either `DATA_NODE` or `MASTER_NODE`.
	Type *string `json:"type" tf:"type,omitempty"`
}

type MasterNodeObservation struct {
}

type MasterNodeParameters struct {

	// +kubebuilder:validation:Required
	// (Required) Resources allocated to hosts of the Elasticsearch data nodes subcluster. The structure is documented below.
	Resources []MasterNodeResourcesParameters `json:"resources" tf:"resources,omitempty"`
}

type MasterNodeResourcesObservation struct {
}

type MasterNodeResourcesParameters struct {

	// +kubebuilder:validation:Required
	// (Required) Volume of the storage available to a host, in gigabytes.
	DiskSize *float64 `json:"diskSize" tf:"disk_size,omitempty"`

	// +kubebuilder:validation:Required
	// (Required) Type of the storage of Elasticsearch hosts.
	DiskTypeID *string `json:"diskTypeId" tf:"disk_type_id,omitempty"`

	// +kubebuilder:validation:Required
	ResourcePresetID *string `json:"resourcePresetId" tf:"resource_preset_id,omitempty"`
}

type ResourcesObservation struct {
}

type ResourcesParameters struct {

	// +kubebuilder:validation:Required
	// (Required) Volume of the storage available to a host, in gigabytes.
	DiskSize *float64 `json:"diskSize" tf:"disk_size,omitempty"`

	// +kubebuilder:validation:Required
	// (Required) Type of the storage of Elasticsearch hosts.
	DiskTypeID *string `json:"diskTypeId" tf:"disk_type_id,omitempty"`

	// +kubebuilder:validation:Required
	ResourcePresetID *string `json:"resourcePresetId" tf:"resource_preset_id,omitempty"`
}

// ElasticsearchClusterSpec defines the desired state of ElasticsearchCluster
type ElasticsearchClusterSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     ElasticsearchClusterParameters `json:"forProvider"`
}

// ElasticsearchClusterStatus defines the observed state of ElasticsearchCluster.
type ElasticsearchClusterStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        ElasticsearchClusterObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// ElasticsearchCluster is the Schema for the ElasticsearchClusters API. <no value>
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,yandex-cloud}
type ElasticsearchCluster struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              ElasticsearchClusterSpec   `json:"spec"`
	Status            ElasticsearchClusterStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ElasticsearchClusterList contains a list of ElasticsearchClusters
type ElasticsearchClusterList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ElasticsearchCluster `json:"items"`
}

// Repository type metadata.
var (
	ElasticsearchCluster_Kind             = "ElasticsearchCluster"
	ElasticsearchCluster_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: ElasticsearchCluster_Kind}.String()
	ElasticsearchCluster_KindAPIVersion   = ElasticsearchCluster_Kind + "." + CRDGroupVersion.String()
	ElasticsearchCluster_GroupVersionKind = CRDGroupVersion.WithKind(ElasticsearchCluster_Kind)
)

func init() {
	SchemeBuilder.Register(&ElasticsearchCluster{}, &ElasticsearchClusterList{})
}
