/*
Copyright 2022 Upbound Inc.
*/
// Code generated by angryjet. DO NOT EDIT.

package v1beta1

import (
	"context"
	reference "github.com/crossplane/crossplane-runtime/pkg/reference"
	errors "github.com/pkg/errors"
	v1beta1 "github.com/upbound/provider-aws/apis/ec2/v1beta1"
	resource "github.com/upbound/upjet/pkg/resource"
	client "sigs.k8s.io/controller-runtime/pkg/client"
)

// ResolveReferences of this Service.
func (mg *Service) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	for i3 := 0; i3 < len(mg.Spec.ForProvider.NetworkConfiguration); i3++ {
		for i4 := 0; i4 < len(mg.Spec.ForProvider.NetworkConfiguration[i3].EgressConfiguration); i4++ {
			rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
				CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.NetworkConfiguration[i3].EgressConfiguration[i4].VPCConnectorArn),
				Extract:      resource.ExtractParamPath("arn", true),
				Reference:    mg.Spec.ForProvider.NetworkConfiguration[i3].EgressConfiguration[i4].VPCConnectorArnRef,
				Selector:     mg.Spec.ForProvider.NetworkConfiguration[i3].EgressConfiguration[i4].VPCConnectorArnSelector,
				To: reference.To{
					List:    &VPCConnectorList{},
					Managed: &VPCConnector{},
				},
			})
			if err != nil {
				return errors.Wrap(err, "mg.Spec.ForProvider.NetworkConfiguration[i3].EgressConfiguration[i4].VPCConnectorArn")
			}
			mg.Spec.ForProvider.NetworkConfiguration[i3].EgressConfiguration[i4].VPCConnectorArn = reference.ToPtrValue(rsp.ResolvedValue)
			mg.Spec.ForProvider.NetworkConfiguration[i3].EgressConfiguration[i4].VPCConnectorArnRef = rsp.ResolvedReference

		}
	}
	for i3 := 0; i3 < len(mg.Spec.ForProvider.ObservabilityConfiguration); i3++ {
		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.ObservabilityConfiguration[i3].ObservabilityConfigurationArn),
			Extract:      resource.ExtractParamPath("arn", true),
			Reference:    mg.Spec.ForProvider.ObservabilityConfiguration[i3].ObservabilityConfigurationArnRef,
			Selector:     mg.Spec.ForProvider.ObservabilityConfiguration[i3].ObservabilityConfigurationArnSelector,
			To: reference.To{
				List:    &ObservabilityConfigurationList{},
				Managed: &ObservabilityConfiguration{},
			},
		})
		if err != nil {
			return errors.Wrap(err, "mg.Spec.ForProvider.ObservabilityConfiguration[i3].ObservabilityConfigurationArn")
		}
		mg.Spec.ForProvider.ObservabilityConfiguration[i3].ObservabilityConfigurationArn = reference.ToPtrValue(rsp.ResolvedValue)
		mg.Spec.ForProvider.ObservabilityConfiguration[i3].ObservabilityConfigurationArnRef = rsp.ResolvedReference

	}
	for i3 := 0; i3 < len(mg.Spec.ForProvider.SourceConfiguration); i3++ {
		for i4 := 0; i4 < len(mg.Spec.ForProvider.SourceConfiguration[i3].AuthenticationConfiguration); i4++ {
			rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
				CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.SourceConfiguration[i3].AuthenticationConfiguration[i4].ConnectionArn),
				Extract:      resource.ExtractParamPath("arn", true),
				Reference:    mg.Spec.ForProvider.SourceConfiguration[i3].AuthenticationConfiguration[i4].ConnectionArnRef,
				Selector:     mg.Spec.ForProvider.SourceConfiguration[i3].AuthenticationConfiguration[i4].ConnectionArnSelector,
				To: reference.To{
					List:    &ConnectionList{},
					Managed: &Connection{},
				},
			})
			if err != nil {
				return errors.Wrap(err, "mg.Spec.ForProvider.SourceConfiguration[i3].AuthenticationConfiguration[i4].ConnectionArn")
			}
			mg.Spec.ForProvider.SourceConfiguration[i3].AuthenticationConfiguration[i4].ConnectionArn = reference.ToPtrValue(rsp.ResolvedValue)
			mg.Spec.ForProvider.SourceConfiguration[i3].AuthenticationConfiguration[i4].ConnectionArnRef = rsp.ResolvedReference

		}
	}

	return nil
}

// ResolveReferences of this VPCConnector.
func (mg *VPCConnector) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var mrsp reference.MultiResolutionResponse
	var err error

	mrsp, err = r.ResolveMultiple(ctx, reference.MultiResolutionRequest{
		CurrentValues: reference.FromPtrValues(mg.Spec.ForProvider.SecurityGroups),
		Extract:       reference.ExternalName(),
		References:    mg.Spec.ForProvider.SecurityGroupRefs,
		Selector:      mg.Spec.ForProvider.SecurityGroupSelector,
		To: reference.To{
			List:    &v1beta1.SecurityGroupList{},
			Managed: &v1beta1.SecurityGroup{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.SecurityGroups")
	}
	mg.Spec.ForProvider.SecurityGroups = reference.ToPtrValues(mrsp.ResolvedValues)
	mg.Spec.ForProvider.SecurityGroupRefs = mrsp.ResolvedReferences

	mrsp, err = r.ResolveMultiple(ctx, reference.MultiResolutionRequest{
		CurrentValues: reference.FromPtrValues(mg.Spec.ForProvider.Subnets),
		Extract:       reference.ExternalName(),
		References:    mg.Spec.ForProvider.SubnetRefs,
		Selector:      mg.Spec.ForProvider.SubnetSelector,
		To: reference.To{
			List:    &v1beta1.SubnetList{},
			Managed: &v1beta1.Subnet{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.Subnets")
	}
	mg.Spec.ForProvider.Subnets = reference.ToPtrValues(mrsp.ResolvedValues)
	mg.Spec.ForProvider.SubnetRefs = mrsp.ResolvedReferences

	return nil
}
