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
// Code generated by angryjet. DO NOT EDIT.

package v1alpha1

import (
	"context"
	reference "github.com/crossplane/crossplane-runtime/pkg/reference"
	errors "github.com/pkg/errors"
	v1alpha13 "github.com/yandex-cloud/provider-jet-yc/apis/iam/v1alpha1"
	v1alpha11 "github.com/yandex-cloud/provider-jet-yc/apis/kms/v1alpha1"
	v1alpha1 "github.com/yandex-cloud/provider-jet-yc/apis/resourcemanager/v1alpha1"
	v1alpha12 "github.com/yandex-cloud/provider-jet-yc/apis/vpc/v1alpha1"
	client "sigs.k8s.io/controller-runtime/pkg/client"
)

// ResolveReferences of this Cluster.
func (mg *Cluster) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var mrsp reference.MultiResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.FolderID),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.FolderIDRef,
		Selector:     mg.Spec.ForProvider.FolderIDSelector,
		To: reference.To{
			List:    &v1alpha1.FolderList{},
			Managed: &v1alpha1.Folder{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.FolderID")
	}
	mg.Spec.ForProvider.FolderID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.FolderIDRef = rsp.ResolvedReference

	for i3 := 0; i3 < len(mg.Spec.ForProvider.KMSProvider); i3++ {
		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.KMSProvider[i3].KeyID),
			Extract:      reference.ExternalName(),
			Reference:    mg.Spec.ForProvider.KMSProvider[i3].KeyIDRef,
			Selector:     mg.Spec.ForProvider.KMSProvider[i3].KeyIDSelector,
			To: reference.To{
				List:    &v1alpha11.SymmetricKeyList{},
				Managed: &v1alpha11.SymmetricKey{},
			},
		})
		if err != nil {
			return errors.Wrap(err, "mg.Spec.ForProvider.KMSProvider[i3].KeyID")
		}
		mg.Spec.ForProvider.KMSProvider[i3].KeyID = reference.ToPtrValue(rsp.ResolvedValue)
		mg.Spec.ForProvider.KMSProvider[i3].KeyIDRef = rsp.ResolvedReference

	}
	for i3 := 0; i3 < len(mg.Spec.ForProvider.Master); i3++ {
		for i4 := 0; i4 < len(mg.Spec.ForProvider.Master[i3].Regional); i4++ {
			for i5 := 0; i5 < len(mg.Spec.ForProvider.Master[i3].Regional[i4].Location); i5++ {
				rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
					CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.Master[i3].Regional[i4].Location[i5].SubnetID),
					Extract:      reference.ExternalName(),
					Reference:    mg.Spec.ForProvider.Master[i3].Regional[i4].Location[i5].SubnetIDRef,
					Selector:     mg.Spec.ForProvider.Master[i3].Regional[i4].Location[i5].SubnetIDSelector,
					To: reference.To{
						List:    &v1alpha12.SubnetList{},
						Managed: &v1alpha12.Subnet{},
					},
				})
				if err != nil {
					return errors.Wrap(err, "mg.Spec.ForProvider.Master[i3].Regional[i4].Location[i5].SubnetID")
				}
				mg.Spec.ForProvider.Master[i3].Regional[i4].Location[i5].SubnetID = reference.ToPtrValue(rsp.ResolvedValue)
				mg.Spec.ForProvider.Master[i3].Regional[i4].Location[i5].SubnetIDRef = rsp.ResolvedReference

			}
		}
	}
	for i3 := 0; i3 < len(mg.Spec.ForProvider.Master); i3++ {
		mrsp, err = r.ResolveMultiple(ctx, reference.MultiResolutionRequest{
			CurrentValues: reference.FromPtrValues(mg.Spec.ForProvider.Master[i3].SecurityGroupIds),
			Extract:       reference.ExternalName(),
			References:    mg.Spec.ForProvider.Master[i3].SecurityGroupIdsRefs,
			Selector:      mg.Spec.ForProvider.Master[i3].SecurityGroupIdsSelector,
			To: reference.To{
				List:    &v1alpha12.SecurityGroupList{},
				Managed: &v1alpha12.SecurityGroup{},
			},
		})
		if err != nil {
			return errors.Wrap(err, "mg.Spec.ForProvider.Master[i3].SecurityGroupIds")
		}
		mg.Spec.ForProvider.Master[i3].SecurityGroupIds = reference.ToPtrValues(mrsp.ResolvedValues)
		mg.Spec.ForProvider.Master[i3].SecurityGroupIdsRefs = mrsp.ResolvedReferences

	}
	for i3 := 0; i3 < len(mg.Spec.ForProvider.Master); i3++ {
		for i4 := 0; i4 < len(mg.Spec.ForProvider.Master[i3].Zonal); i4++ {
			rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
				CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.Master[i3].Zonal[i4].SubnetID),
				Extract:      reference.ExternalName(),
				Reference:    mg.Spec.ForProvider.Master[i3].Zonal[i4].SubnetIDRef,
				Selector:     mg.Spec.ForProvider.Master[i3].Zonal[i4].SubnetIDSelector,
				To: reference.To{
					List:    &v1alpha12.SubnetList{},
					Managed: &v1alpha12.Subnet{},
				},
			})
			if err != nil {
				return errors.Wrap(err, "mg.Spec.ForProvider.Master[i3].Zonal[i4].SubnetID")
			}
			mg.Spec.ForProvider.Master[i3].Zonal[i4].SubnetID = reference.ToPtrValue(rsp.ResolvedValue)
			mg.Spec.ForProvider.Master[i3].Zonal[i4].SubnetIDRef = rsp.ResolvedReference

		}
	}
	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.NetworkID),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.NetworkIDRef,
		Selector:     mg.Spec.ForProvider.NetworkIDSelector,
		To: reference.To{
			List:    &v1alpha12.NetworkList{},
			Managed: &v1alpha12.Network{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.NetworkID")
	}
	mg.Spec.ForProvider.NetworkID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.NetworkIDRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.NodeServiceAccountID),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.NodeServiceAccountIDRef,
		Selector:     mg.Spec.ForProvider.NodeServiceAccountIDSelector,
		To: reference.To{
			List:    &v1alpha13.ServiceAccountList{},
			Managed: &v1alpha13.ServiceAccount{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.NodeServiceAccountID")
	}
	mg.Spec.ForProvider.NodeServiceAccountID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.NodeServiceAccountIDRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.ServiceAccountID),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.ServiceAccountIDRef,
		Selector:     mg.Spec.ForProvider.ServiceAccountIDSelector,
		To: reference.To{
			List:    &v1alpha13.ServiceAccountList{},
			Managed: &v1alpha13.ServiceAccount{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.ServiceAccountID")
	}
	mg.Spec.ForProvider.ServiceAccountID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.ServiceAccountIDRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this NodeGroup.
func (mg *NodeGroup) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var mrsp reference.MultiResolutionResponse
	var err error

	for i3 := 0; i3 < len(mg.Spec.ForProvider.AllocationPolicy); i3++ {
		for i4 := 0; i4 < len(mg.Spec.ForProvider.AllocationPolicy[i3].Location); i4++ {
			rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
				CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.AllocationPolicy[i3].Location[i4].SubnetID),
				Extract:      reference.ExternalName(),
				Reference:    mg.Spec.ForProvider.AllocationPolicy[i3].Location[i4].SubnetIDRef,
				Selector:     mg.Spec.ForProvider.AllocationPolicy[i3].Location[i4].SubnetIDSelector,
				To: reference.To{
					List:    &v1alpha12.SubnetList{},
					Managed: &v1alpha12.Subnet{},
				},
			})
			if err != nil {
				return errors.Wrap(err, "mg.Spec.ForProvider.AllocationPolicy[i3].Location[i4].SubnetID")
			}
			mg.Spec.ForProvider.AllocationPolicy[i3].Location[i4].SubnetID = reference.ToPtrValue(rsp.ResolvedValue)
			mg.Spec.ForProvider.AllocationPolicy[i3].Location[i4].SubnetIDRef = rsp.ResolvedReference

		}
	}
	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.ClusterID),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.ClusterIDRef,
		Selector:     mg.Spec.ForProvider.ClusterIDSelector,
		To: reference.To{
			List:    &ClusterList{},
			Managed: &Cluster{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.ClusterID")
	}
	mg.Spec.ForProvider.ClusterID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.ClusterIDRef = rsp.ResolvedReference

	for i3 := 0; i3 < len(mg.Spec.ForProvider.InstanceTemplate); i3++ {
		for i4 := 0; i4 < len(mg.Spec.ForProvider.InstanceTemplate[i3].NetworkInterface); i4++ {
			mrsp, err = r.ResolveMultiple(ctx, reference.MultiResolutionRequest{
				CurrentValues: reference.FromPtrValues(mg.Spec.ForProvider.InstanceTemplate[i3].NetworkInterface[i4].SecurityGroupIds),
				Extract:       reference.ExternalName(),
				References:    mg.Spec.ForProvider.InstanceTemplate[i3].NetworkInterface[i4].SecurityGroupIdsRefs,
				Selector:      mg.Spec.ForProvider.InstanceTemplate[i3].NetworkInterface[i4].SecurityGroupIdsSelector,
				To: reference.To{
					List:    &v1alpha12.SecurityGroupList{},
					Managed: &v1alpha12.SecurityGroup{},
				},
			})
			if err != nil {
				return errors.Wrap(err, "mg.Spec.ForProvider.InstanceTemplate[i3].NetworkInterface[i4].SecurityGroupIds")
			}
			mg.Spec.ForProvider.InstanceTemplate[i3].NetworkInterface[i4].SecurityGroupIds = reference.ToPtrValues(mrsp.ResolvedValues)
			mg.Spec.ForProvider.InstanceTemplate[i3].NetworkInterface[i4].SecurityGroupIdsRefs = mrsp.ResolvedReferences

		}
	}
	for i3 := 0; i3 < len(mg.Spec.ForProvider.InstanceTemplate); i3++ {
		for i4 := 0; i4 < len(mg.Spec.ForProvider.InstanceTemplate[i3].NetworkInterface); i4++ {
			mrsp, err = r.ResolveMultiple(ctx, reference.MultiResolutionRequest{
				CurrentValues: reference.FromPtrValues(mg.Spec.ForProvider.InstanceTemplate[i3].NetworkInterface[i4].SubnetIds),
				Extract:       reference.ExternalName(),
				References:    mg.Spec.ForProvider.InstanceTemplate[i3].NetworkInterface[i4].SubnetIdsRefs,
				Selector:      mg.Spec.ForProvider.InstanceTemplate[i3].NetworkInterface[i4].SubnetIdsSelector,
				To: reference.To{
					List:    &v1alpha12.SubnetList{},
					Managed: &v1alpha12.Subnet{},
				},
			})
			if err != nil {
				return errors.Wrap(err, "mg.Spec.ForProvider.InstanceTemplate[i3].NetworkInterface[i4].SubnetIds")
			}
			mg.Spec.ForProvider.InstanceTemplate[i3].NetworkInterface[i4].SubnetIds = reference.ToPtrValues(mrsp.ResolvedValues)
			mg.Spec.ForProvider.InstanceTemplate[i3].NetworkInterface[i4].SubnetIdsRefs = mrsp.ResolvedReferences

		}
	}

	return nil
}
