/*
Copyright 2022 Upbound Inc.
*/
// Code generated by angryjet. DO NOT EDIT.

package v1beta1

import (
	"context"
	reference "github.com/crossplane/crossplane-runtime/pkg/reference"
	errors "github.com/pkg/errors"
	v1beta1 "github.com/upbound/provider-azure/apis/authorization/v1beta1"
	resource "github.com/upbound/upjet/pkg/resource"
	client "sigs.k8s.io/controller-runtime/pkg/client"
)

// ResolveReferences of this SubscriptionPolicyRemediation.
func (mg *SubscriptionPolicyRemediation) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.PolicyAssignmentID),
		Extract:      resource.ExtractResourceID(),
		Reference:    mg.Spec.ForProvider.PolicyAssignmentIDRef,
		Selector:     mg.Spec.ForProvider.PolicyAssignmentIDSelector,
		To: reference.To{
			List:    &v1beta1.SubscriptionPolicyAssignmentList{},
			Managed: &v1beta1.SubscriptionPolicyAssignment{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.PolicyAssignmentID")
	}
	mg.Spec.ForProvider.PolicyAssignmentID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.PolicyAssignmentIDRef = rsp.ResolvedReference

	return nil
}
