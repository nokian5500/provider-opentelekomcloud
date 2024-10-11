/*
Copyright 2022 Upbound Inc.
*/
// Code generated by angryjet. DO NOT EDIT.

package v1alpha1

import (
	"context"
	reference "github.com/crossplane/crossplane-runtime/pkg/reference"
	v1alpha1 "github.com/opentelekomcloud/provider-opentelekomcloud/apis/vpc/v1alpha1"
	errors "github.com/pkg/errors"
	client "sigs.k8s.io/controller-runtime/pkg/client"
)

// ResolveReferences of this PtrrecordV2.
func (mg *PtrrecordV2) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.FloatingipID),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.FloatingipIDRef,
		Selector:     mg.Spec.ForProvider.FloatingipIDSelector,
		To: reference.To{
			List:    &v1alpha1.EIPV1List{},
			Managed: &v1alpha1.EIPV1{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.FloatingipID")
	}
	mg.Spec.ForProvider.FloatingipID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.FloatingipIDRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.FloatingipID),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.InitProvider.FloatingipIDRef,
		Selector:     mg.Spec.InitProvider.FloatingipIDSelector,
		To: reference.To{
			List:    &v1alpha1.EIPV1List{},
			Managed: &v1alpha1.EIPV1{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.FloatingipID")
	}
	mg.Spec.InitProvider.FloatingipID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.InitProvider.FloatingipIDRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this RecordsetV2.
func (mg *RecordsetV2) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.ZoneID),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.ZoneIDRef,
		Selector:     mg.Spec.ForProvider.ZoneIDSelector,
		To: reference.To{
			List:    &ZoneV2List{},
			Managed: &ZoneV2{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.ZoneID")
	}
	mg.Spec.ForProvider.ZoneID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.ZoneIDRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.ZoneID),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.InitProvider.ZoneIDRef,
		Selector:     mg.Spec.InitProvider.ZoneIDSelector,
		To: reference.To{
			List:    &ZoneV2List{},
			Managed: &ZoneV2{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.ZoneID")
	}
	mg.Spec.InitProvider.ZoneID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.InitProvider.ZoneIDRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this ZoneV2.
func (mg *ZoneV2) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	for i3 := 0; i3 < len(mg.Spec.ForProvider.Router); i3++ {
		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.Router[i3].RouterID),
			Extract:      reference.ExternalName(),
			Reference:    mg.Spec.ForProvider.Router[i3].RouterIDRef,
			Selector:     mg.Spec.ForProvider.Router[i3].RouterIDSelector,
			To: reference.To{
				List:    &v1alpha1.VpcV1List{},
				Managed: &v1alpha1.VpcV1{},
			},
		})
		if err != nil {
			return errors.Wrap(err, "mg.Spec.ForProvider.Router[i3].RouterID")
		}
		mg.Spec.ForProvider.Router[i3].RouterID = reference.ToPtrValue(rsp.ResolvedValue)
		mg.Spec.ForProvider.Router[i3].RouterIDRef = rsp.ResolvedReference

	}
	for i3 := 0; i3 < len(mg.Spec.InitProvider.Router); i3++ {
		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.Router[i3].RouterID),
			Extract:      reference.ExternalName(),
			Reference:    mg.Spec.InitProvider.Router[i3].RouterIDRef,
			Selector:     mg.Spec.InitProvider.Router[i3].RouterIDSelector,
			To: reference.To{
				List:    &v1alpha1.VpcV1List{},
				Managed: &v1alpha1.VpcV1{},
			},
		})
		if err != nil {
			return errors.Wrap(err, "mg.Spec.InitProvider.Router[i3].RouterID")
		}
		mg.Spec.InitProvider.Router[i3].RouterID = reference.ToPtrValue(rsp.ResolvedValue)
		mg.Spec.InitProvider.Router[i3].RouterIDRef = rsp.ResolvedReference

	}

	return nil
}
