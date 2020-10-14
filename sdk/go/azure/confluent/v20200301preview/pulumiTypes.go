// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v20200301preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Confluent offer detail
type OrganizationResourcePropertiesOfferDetail struct {
	// Offer Id
	Id *string `pulumi:"id"`
	// Offer Plan Id
	PlanId *string `pulumi:"planId"`
	// Offer Plan Name
	PlanName *string `pulumi:"planName"`
	// Publisher Id
	PublisherId *string `pulumi:"publisherId"`
	// SaaS Offer Status
	Status *string `pulumi:"status"`
	// Offer Plan Term unit
	TermUnit *string `pulumi:"termUnit"`
}

// OrganizationResourcePropertiesOfferDetailInput is an input type that accepts OrganizationResourcePropertiesOfferDetailArgs and OrganizationResourcePropertiesOfferDetailOutput values.
// You can construct a concrete instance of `OrganizationResourcePropertiesOfferDetailInput` via:
//
//          OrganizationResourcePropertiesOfferDetailArgs{...}
type OrganizationResourcePropertiesOfferDetailInput interface {
	pulumi.Input

	ToOrganizationResourcePropertiesOfferDetailOutput() OrganizationResourcePropertiesOfferDetailOutput
	ToOrganizationResourcePropertiesOfferDetailOutputWithContext(context.Context) OrganizationResourcePropertiesOfferDetailOutput
}

// Confluent offer detail
type OrganizationResourcePropertiesOfferDetailArgs struct {
	// Offer Id
	Id pulumi.StringPtrInput `pulumi:"id"`
	// Offer Plan Id
	PlanId pulumi.StringPtrInput `pulumi:"planId"`
	// Offer Plan Name
	PlanName pulumi.StringPtrInput `pulumi:"planName"`
	// Publisher Id
	PublisherId pulumi.StringPtrInput `pulumi:"publisherId"`
	// SaaS Offer Status
	Status pulumi.StringPtrInput `pulumi:"status"`
	// Offer Plan Term unit
	TermUnit pulumi.StringPtrInput `pulumi:"termUnit"`
}

func (OrganizationResourcePropertiesOfferDetailArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*OrganizationResourcePropertiesOfferDetail)(nil)).Elem()
}

func (i OrganizationResourcePropertiesOfferDetailArgs) ToOrganizationResourcePropertiesOfferDetailOutput() OrganizationResourcePropertiesOfferDetailOutput {
	return i.ToOrganizationResourcePropertiesOfferDetailOutputWithContext(context.Background())
}

func (i OrganizationResourcePropertiesOfferDetailArgs) ToOrganizationResourcePropertiesOfferDetailOutputWithContext(ctx context.Context) OrganizationResourcePropertiesOfferDetailOutput {
	return pulumi.ToOutputWithContext(ctx, i).(OrganizationResourcePropertiesOfferDetailOutput)
}

func (i OrganizationResourcePropertiesOfferDetailArgs) ToOrganizationResourcePropertiesOfferDetailPtrOutput() OrganizationResourcePropertiesOfferDetailPtrOutput {
	return i.ToOrganizationResourcePropertiesOfferDetailPtrOutputWithContext(context.Background())
}

func (i OrganizationResourcePropertiesOfferDetailArgs) ToOrganizationResourcePropertiesOfferDetailPtrOutputWithContext(ctx context.Context) OrganizationResourcePropertiesOfferDetailPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(OrganizationResourcePropertiesOfferDetailOutput).ToOrganizationResourcePropertiesOfferDetailPtrOutputWithContext(ctx)
}

// OrganizationResourcePropertiesOfferDetailPtrInput is an input type that accepts OrganizationResourcePropertiesOfferDetailArgs, OrganizationResourcePropertiesOfferDetailPtr and OrganizationResourcePropertiesOfferDetailPtrOutput values.
// You can construct a concrete instance of `OrganizationResourcePropertiesOfferDetailPtrInput` via:
//
//          OrganizationResourcePropertiesOfferDetailArgs{...}
//
//  or:
//
//          nil
type OrganizationResourcePropertiesOfferDetailPtrInput interface {
	pulumi.Input

	ToOrganizationResourcePropertiesOfferDetailPtrOutput() OrganizationResourcePropertiesOfferDetailPtrOutput
	ToOrganizationResourcePropertiesOfferDetailPtrOutputWithContext(context.Context) OrganizationResourcePropertiesOfferDetailPtrOutput
}

type organizationResourcePropertiesOfferDetailPtrType OrganizationResourcePropertiesOfferDetailArgs

func OrganizationResourcePropertiesOfferDetailPtr(v *OrganizationResourcePropertiesOfferDetailArgs) OrganizationResourcePropertiesOfferDetailPtrInput {
	return (*organizationResourcePropertiesOfferDetailPtrType)(v)
}

func (*organizationResourcePropertiesOfferDetailPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**OrganizationResourcePropertiesOfferDetail)(nil)).Elem()
}

func (i *organizationResourcePropertiesOfferDetailPtrType) ToOrganizationResourcePropertiesOfferDetailPtrOutput() OrganizationResourcePropertiesOfferDetailPtrOutput {
	return i.ToOrganizationResourcePropertiesOfferDetailPtrOutputWithContext(context.Background())
}

func (i *organizationResourcePropertiesOfferDetailPtrType) ToOrganizationResourcePropertiesOfferDetailPtrOutputWithContext(ctx context.Context) OrganizationResourcePropertiesOfferDetailPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(OrganizationResourcePropertiesOfferDetailPtrOutput)
}

// Confluent offer detail
type OrganizationResourcePropertiesOfferDetailOutput struct{ *pulumi.OutputState }

func (OrganizationResourcePropertiesOfferDetailOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*OrganizationResourcePropertiesOfferDetail)(nil)).Elem()
}

func (o OrganizationResourcePropertiesOfferDetailOutput) ToOrganizationResourcePropertiesOfferDetailOutput() OrganizationResourcePropertiesOfferDetailOutput {
	return o
}

func (o OrganizationResourcePropertiesOfferDetailOutput) ToOrganizationResourcePropertiesOfferDetailOutputWithContext(ctx context.Context) OrganizationResourcePropertiesOfferDetailOutput {
	return o
}

func (o OrganizationResourcePropertiesOfferDetailOutput) ToOrganizationResourcePropertiesOfferDetailPtrOutput() OrganizationResourcePropertiesOfferDetailPtrOutput {
	return o.ToOrganizationResourcePropertiesOfferDetailPtrOutputWithContext(context.Background())
}

func (o OrganizationResourcePropertiesOfferDetailOutput) ToOrganizationResourcePropertiesOfferDetailPtrOutputWithContext(ctx context.Context) OrganizationResourcePropertiesOfferDetailPtrOutput {
	return o.ApplyT(func(v OrganizationResourcePropertiesOfferDetail) *OrganizationResourcePropertiesOfferDetail {
		return &v
	}).(OrganizationResourcePropertiesOfferDetailPtrOutput)
}

// Offer Id
func (o OrganizationResourcePropertiesOfferDetailOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v OrganizationResourcePropertiesOfferDetail) *string { return v.Id }).(pulumi.StringPtrOutput)
}

// Offer Plan Id
func (o OrganizationResourcePropertiesOfferDetailOutput) PlanId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v OrganizationResourcePropertiesOfferDetail) *string { return v.PlanId }).(pulumi.StringPtrOutput)
}

// Offer Plan Name
func (o OrganizationResourcePropertiesOfferDetailOutput) PlanName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v OrganizationResourcePropertiesOfferDetail) *string { return v.PlanName }).(pulumi.StringPtrOutput)
}

// Publisher Id
func (o OrganizationResourcePropertiesOfferDetailOutput) PublisherId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v OrganizationResourcePropertiesOfferDetail) *string { return v.PublisherId }).(pulumi.StringPtrOutput)
}

// SaaS Offer Status
func (o OrganizationResourcePropertiesOfferDetailOutput) Status() pulumi.StringPtrOutput {
	return o.ApplyT(func(v OrganizationResourcePropertiesOfferDetail) *string { return v.Status }).(pulumi.StringPtrOutput)
}

// Offer Plan Term unit
func (o OrganizationResourcePropertiesOfferDetailOutput) TermUnit() pulumi.StringPtrOutput {
	return o.ApplyT(func(v OrganizationResourcePropertiesOfferDetail) *string { return v.TermUnit }).(pulumi.StringPtrOutput)
}

type OrganizationResourcePropertiesOfferDetailPtrOutput struct{ *pulumi.OutputState }

func (OrganizationResourcePropertiesOfferDetailPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**OrganizationResourcePropertiesOfferDetail)(nil)).Elem()
}

func (o OrganizationResourcePropertiesOfferDetailPtrOutput) ToOrganizationResourcePropertiesOfferDetailPtrOutput() OrganizationResourcePropertiesOfferDetailPtrOutput {
	return o
}

func (o OrganizationResourcePropertiesOfferDetailPtrOutput) ToOrganizationResourcePropertiesOfferDetailPtrOutputWithContext(ctx context.Context) OrganizationResourcePropertiesOfferDetailPtrOutput {
	return o
}

func (o OrganizationResourcePropertiesOfferDetailPtrOutput) Elem() OrganizationResourcePropertiesOfferDetailOutput {
	return o.ApplyT(func(v *OrganizationResourcePropertiesOfferDetail) OrganizationResourcePropertiesOfferDetail {
		return *v
	}).(OrganizationResourcePropertiesOfferDetailOutput)
}

// Offer Id
func (o OrganizationResourcePropertiesOfferDetailPtrOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *OrganizationResourcePropertiesOfferDetail) *string {
		if v == nil {
			return nil
		}
		return v.Id
	}).(pulumi.StringPtrOutput)
}

// Offer Plan Id
func (o OrganizationResourcePropertiesOfferDetailPtrOutput) PlanId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *OrganizationResourcePropertiesOfferDetail) *string {
		if v == nil {
			return nil
		}
		return v.PlanId
	}).(pulumi.StringPtrOutput)
}

// Offer Plan Name
func (o OrganizationResourcePropertiesOfferDetailPtrOutput) PlanName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *OrganizationResourcePropertiesOfferDetail) *string {
		if v == nil {
			return nil
		}
		return v.PlanName
	}).(pulumi.StringPtrOutput)
}

// Publisher Id
func (o OrganizationResourcePropertiesOfferDetailPtrOutput) PublisherId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *OrganizationResourcePropertiesOfferDetail) *string {
		if v == nil {
			return nil
		}
		return v.PublisherId
	}).(pulumi.StringPtrOutput)
}

// SaaS Offer Status
func (o OrganizationResourcePropertiesOfferDetailPtrOutput) Status() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *OrganizationResourcePropertiesOfferDetail) *string {
		if v == nil {
			return nil
		}
		return v.Status
	}).(pulumi.StringPtrOutput)
}

// Offer Plan Term unit
func (o OrganizationResourcePropertiesOfferDetailPtrOutput) TermUnit() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *OrganizationResourcePropertiesOfferDetail) *string {
		if v == nil {
			return nil
		}
		return v.TermUnit
	}).(pulumi.StringPtrOutput)
}

// Confluent offer detail
type OrganizationResourcePropertiesResponseOfferDetail struct {
	// Offer Id
	Id *string `pulumi:"id"`
	// Offer Plan Id
	PlanId *string `pulumi:"planId"`
	// Offer Plan Name
	PlanName *string `pulumi:"planName"`
	// Publisher Id
	PublisherId *string `pulumi:"publisherId"`
	// SaaS Offer Status
	Status *string `pulumi:"status"`
	// Offer Plan Term unit
	TermUnit *string `pulumi:"termUnit"`
}

// OrganizationResourcePropertiesResponseOfferDetailInput is an input type that accepts OrganizationResourcePropertiesResponseOfferDetailArgs and OrganizationResourcePropertiesResponseOfferDetailOutput values.
// You can construct a concrete instance of `OrganizationResourcePropertiesResponseOfferDetailInput` via:
//
//          OrganizationResourcePropertiesResponseOfferDetailArgs{...}
type OrganizationResourcePropertiesResponseOfferDetailInput interface {
	pulumi.Input

	ToOrganizationResourcePropertiesResponseOfferDetailOutput() OrganizationResourcePropertiesResponseOfferDetailOutput
	ToOrganizationResourcePropertiesResponseOfferDetailOutputWithContext(context.Context) OrganizationResourcePropertiesResponseOfferDetailOutput
}

// Confluent offer detail
type OrganizationResourcePropertiesResponseOfferDetailArgs struct {
	// Offer Id
	Id pulumi.StringPtrInput `pulumi:"id"`
	// Offer Plan Id
	PlanId pulumi.StringPtrInput `pulumi:"planId"`
	// Offer Plan Name
	PlanName pulumi.StringPtrInput `pulumi:"planName"`
	// Publisher Id
	PublisherId pulumi.StringPtrInput `pulumi:"publisherId"`
	// SaaS Offer Status
	Status pulumi.StringPtrInput `pulumi:"status"`
	// Offer Plan Term unit
	TermUnit pulumi.StringPtrInput `pulumi:"termUnit"`
}

func (OrganizationResourcePropertiesResponseOfferDetailArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*OrganizationResourcePropertiesResponseOfferDetail)(nil)).Elem()
}

func (i OrganizationResourcePropertiesResponseOfferDetailArgs) ToOrganizationResourcePropertiesResponseOfferDetailOutput() OrganizationResourcePropertiesResponseOfferDetailOutput {
	return i.ToOrganizationResourcePropertiesResponseOfferDetailOutputWithContext(context.Background())
}

func (i OrganizationResourcePropertiesResponseOfferDetailArgs) ToOrganizationResourcePropertiesResponseOfferDetailOutputWithContext(ctx context.Context) OrganizationResourcePropertiesResponseOfferDetailOutput {
	return pulumi.ToOutputWithContext(ctx, i).(OrganizationResourcePropertiesResponseOfferDetailOutput)
}

func (i OrganizationResourcePropertiesResponseOfferDetailArgs) ToOrganizationResourcePropertiesResponseOfferDetailPtrOutput() OrganizationResourcePropertiesResponseOfferDetailPtrOutput {
	return i.ToOrganizationResourcePropertiesResponseOfferDetailPtrOutputWithContext(context.Background())
}

func (i OrganizationResourcePropertiesResponseOfferDetailArgs) ToOrganizationResourcePropertiesResponseOfferDetailPtrOutputWithContext(ctx context.Context) OrganizationResourcePropertiesResponseOfferDetailPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(OrganizationResourcePropertiesResponseOfferDetailOutput).ToOrganizationResourcePropertiesResponseOfferDetailPtrOutputWithContext(ctx)
}

// OrganizationResourcePropertiesResponseOfferDetailPtrInput is an input type that accepts OrganizationResourcePropertiesResponseOfferDetailArgs, OrganizationResourcePropertiesResponseOfferDetailPtr and OrganizationResourcePropertiesResponseOfferDetailPtrOutput values.
// You can construct a concrete instance of `OrganizationResourcePropertiesResponseOfferDetailPtrInput` via:
//
//          OrganizationResourcePropertiesResponseOfferDetailArgs{...}
//
//  or:
//
//          nil
type OrganizationResourcePropertiesResponseOfferDetailPtrInput interface {
	pulumi.Input

	ToOrganizationResourcePropertiesResponseOfferDetailPtrOutput() OrganizationResourcePropertiesResponseOfferDetailPtrOutput
	ToOrganizationResourcePropertiesResponseOfferDetailPtrOutputWithContext(context.Context) OrganizationResourcePropertiesResponseOfferDetailPtrOutput
}

type organizationResourcePropertiesResponseOfferDetailPtrType OrganizationResourcePropertiesResponseOfferDetailArgs

func OrganizationResourcePropertiesResponseOfferDetailPtr(v *OrganizationResourcePropertiesResponseOfferDetailArgs) OrganizationResourcePropertiesResponseOfferDetailPtrInput {
	return (*organizationResourcePropertiesResponseOfferDetailPtrType)(v)
}

func (*organizationResourcePropertiesResponseOfferDetailPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**OrganizationResourcePropertiesResponseOfferDetail)(nil)).Elem()
}

func (i *organizationResourcePropertiesResponseOfferDetailPtrType) ToOrganizationResourcePropertiesResponseOfferDetailPtrOutput() OrganizationResourcePropertiesResponseOfferDetailPtrOutput {
	return i.ToOrganizationResourcePropertiesResponseOfferDetailPtrOutputWithContext(context.Background())
}

func (i *organizationResourcePropertiesResponseOfferDetailPtrType) ToOrganizationResourcePropertiesResponseOfferDetailPtrOutputWithContext(ctx context.Context) OrganizationResourcePropertiesResponseOfferDetailPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(OrganizationResourcePropertiesResponseOfferDetailPtrOutput)
}

// Confluent offer detail
type OrganizationResourcePropertiesResponseOfferDetailOutput struct{ *pulumi.OutputState }

func (OrganizationResourcePropertiesResponseOfferDetailOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*OrganizationResourcePropertiesResponseOfferDetail)(nil)).Elem()
}

func (o OrganizationResourcePropertiesResponseOfferDetailOutput) ToOrganizationResourcePropertiesResponseOfferDetailOutput() OrganizationResourcePropertiesResponseOfferDetailOutput {
	return o
}

func (o OrganizationResourcePropertiesResponseOfferDetailOutput) ToOrganizationResourcePropertiesResponseOfferDetailOutputWithContext(ctx context.Context) OrganizationResourcePropertiesResponseOfferDetailOutput {
	return o
}

func (o OrganizationResourcePropertiesResponseOfferDetailOutput) ToOrganizationResourcePropertiesResponseOfferDetailPtrOutput() OrganizationResourcePropertiesResponseOfferDetailPtrOutput {
	return o.ToOrganizationResourcePropertiesResponseOfferDetailPtrOutputWithContext(context.Background())
}

func (o OrganizationResourcePropertiesResponseOfferDetailOutput) ToOrganizationResourcePropertiesResponseOfferDetailPtrOutputWithContext(ctx context.Context) OrganizationResourcePropertiesResponseOfferDetailPtrOutput {
	return o.ApplyT(func(v OrganizationResourcePropertiesResponseOfferDetail) *OrganizationResourcePropertiesResponseOfferDetail {
		return &v
	}).(OrganizationResourcePropertiesResponseOfferDetailPtrOutput)
}

// Offer Id
func (o OrganizationResourcePropertiesResponseOfferDetailOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v OrganizationResourcePropertiesResponseOfferDetail) *string { return v.Id }).(pulumi.StringPtrOutput)
}

// Offer Plan Id
func (o OrganizationResourcePropertiesResponseOfferDetailOutput) PlanId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v OrganizationResourcePropertiesResponseOfferDetail) *string { return v.PlanId }).(pulumi.StringPtrOutput)
}

// Offer Plan Name
func (o OrganizationResourcePropertiesResponseOfferDetailOutput) PlanName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v OrganizationResourcePropertiesResponseOfferDetail) *string { return v.PlanName }).(pulumi.StringPtrOutput)
}

// Publisher Id
func (o OrganizationResourcePropertiesResponseOfferDetailOutput) PublisherId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v OrganizationResourcePropertiesResponseOfferDetail) *string { return v.PublisherId }).(pulumi.StringPtrOutput)
}

// SaaS Offer Status
func (o OrganizationResourcePropertiesResponseOfferDetailOutput) Status() pulumi.StringPtrOutput {
	return o.ApplyT(func(v OrganizationResourcePropertiesResponseOfferDetail) *string { return v.Status }).(pulumi.StringPtrOutput)
}

// Offer Plan Term unit
func (o OrganizationResourcePropertiesResponseOfferDetailOutput) TermUnit() pulumi.StringPtrOutput {
	return o.ApplyT(func(v OrganizationResourcePropertiesResponseOfferDetail) *string { return v.TermUnit }).(pulumi.StringPtrOutput)
}

type OrganizationResourcePropertiesResponseOfferDetailPtrOutput struct{ *pulumi.OutputState }

func (OrganizationResourcePropertiesResponseOfferDetailPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**OrganizationResourcePropertiesResponseOfferDetail)(nil)).Elem()
}

func (o OrganizationResourcePropertiesResponseOfferDetailPtrOutput) ToOrganizationResourcePropertiesResponseOfferDetailPtrOutput() OrganizationResourcePropertiesResponseOfferDetailPtrOutput {
	return o
}

func (o OrganizationResourcePropertiesResponseOfferDetailPtrOutput) ToOrganizationResourcePropertiesResponseOfferDetailPtrOutputWithContext(ctx context.Context) OrganizationResourcePropertiesResponseOfferDetailPtrOutput {
	return o
}

func (o OrganizationResourcePropertiesResponseOfferDetailPtrOutput) Elem() OrganizationResourcePropertiesResponseOfferDetailOutput {
	return o.ApplyT(func(v *OrganizationResourcePropertiesResponseOfferDetail) OrganizationResourcePropertiesResponseOfferDetail {
		return *v
	}).(OrganizationResourcePropertiesResponseOfferDetailOutput)
}

// Offer Id
func (o OrganizationResourcePropertiesResponseOfferDetailPtrOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *OrganizationResourcePropertiesResponseOfferDetail) *string {
		if v == nil {
			return nil
		}
		return v.Id
	}).(pulumi.StringPtrOutput)
}

// Offer Plan Id
func (o OrganizationResourcePropertiesResponseOfferDetailPtrOutput) PlanId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *OrganizationResourcePropertiesResponseOfferDetail) *string {
		if v == nil {
			return nil
		}
		return v.PlanId
	}).(pulumi.StringPtrOutput)
}

// Offer Plan Name
func (o OrganizationResourcePropertiesResponseOfferDetailPtrOutput) PlanName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *OrganizationResourcePropertiesResponseOfferDetail) *string {
		if v == nil {
			return nil
		}
		return v.PlanName
	}).(pulumi.StringPtrOutput)
}

// Publisher Id
func (o OrganizationResourcePropertiesResponseOfferDetailPtrOutput) PublisherId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *OrganizationResourcePropertiesResponseOfferDetail) *string {
		if v == nil {
			return nil
		}
		return v.PublisherId
	}).(pulumi.StringPtrOutput)
}

// SaaS Offer Status
func (o OrganizationResourcePropertiesResponseOfferDetailPtrOutput) Status() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *OrganizationResourcePropertiesResponseOfferDetail) *string {
		if v == nil {
			return nil
		}
		return v.Status
	}).(pulumi.StringPtrOutput)
}

// Offer Plan Term unit
func (o OrganizationResourcePropertiesResponseOfferDetailPtrOutput) TermUnit() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *OrganizationResourcePropertiesResponseOfferDetail) *string {
		if v == nil {
			return nil
		}
		return v.TermUnit
	}).(pulumi.StringPtrOutput)
}

// Subscriber detail
type OrganizationResourcePropertiesResponseUserDetail struct {
	// Email address
	EmailAddress *string `pulumi:"emailAddress"`
	// First name
	FirstName *string `pulumi:"firstName"`
	// Last name
	LastName *string `pulumi:"lastName"`
}

// OrganizationResourcePropertiesResponseUserDetailInput is an input type that accepts OrganizationResourcePropertiesResponseUserDetailArgs and OrganizationResourcePropertiesResponseUserDetailOutput values.
// You can construct a concrete instance of `OrganizationResourcePropertiesResponseUserDetailInput` via:
//
//          OrganizationResourcePropertiesResponseUserDetailArgs{...}
type OrganizationResourcePropertiesResponseUserDetailInput interface {
	pulumi.Input

	ToOrganizationResourcePropertiesResponseUserDetailOutput() OrganizationResourcePropertiesResponseUserDetailOutput
	ToOrganizationResourcePropertiesResponseUserDetailOutputWithContext(context.Context) OrganizationResourcePropertiesResponseUserDetailOutput
}

// Subscriber detail
type OrganizationResourcePropertiesResponseUserDetailArgs struct {
	// Email address
	EmailAddress pulumi.StringPtrInput `pulumi:"emailAddress"`
	// First name
	FirstName pulumi.StringPtrInput `pulumi:"firstName"`
	// Last name
	LastName pulumi.StringPtrInput `pulumi:"lastName"`
}

func (OrganizationResourcePropertiesResponseUserDetailArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*OrganizationResourcePropertiesResponseUserDetail)(nil)).Elem()
}

func (i OrganizationResourcePropertiesResponseUserDetailArgs) ToOrganizationResourcePropertiesResponseUserDetailOutput() OrganizationResourcePropertiesResponseUserDetailOutput {
	return i.ToOrganizationResourcePropertiesResponseUserDetailOutputWithContext(context.Background())
}

func (i OrganizationResourcePropertiesResponseUserDetailArgs) ToOrganizationResourcePropertiesResponseUserDetailOutputWithContext(ctx context.Context) OrganizationResourcePropertiesResponseUserDetailOutput {
	return pulumi.ToOutputWithContext(ctx, i).(OrganizationResourcePropertiesResponseUserDetailOutput)
}

func (i OrganizationResourcePropertiesResponseUserDetailArgs) ToOrganizationResourcePropertiesResponseUserDetailPtrOutput() OrganizationResourcePropertiesResponseUserDetailPtrOutput {
	return i.ToOrganizationResourcePropertiesResponseUserDetailPtrOutputWithContext(context.Background())
}

func (i OrganizationResourcePropertiesResponseUserDetailArgs) ToOrganizationResourcePropertiesResponseUserDetailPtrOutputWithContext(ctx context.Context) OrganizationResourcePropertiesResponseUserDetailPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(OrganizationResourcePropertiesResponseUserDetailOutput).ToOrganizationResourcePropertiesResponseUserDetailPtrOutputWithContext(ctx)
}

// OrganizationResourcePropertiesResponseUserDetailPtrInput is an input type that accepts OrganizationResourcePropertiesResponseUserDetailArgs, OrganizationResourcePropertiesResponseUserDetailPtr and OrganizationResourcePropertiesResponseUserDetailPtrOutput values.
// You can construct a concrete instance of `OrganizationResourcePropertiesResponseUserDetailPtrInput` via:
//
//          OrganizationResourcePropertiesResponseUserDetailArgs{...}
//
//  or:
//
//          nil
type OrganizationResourcePropertiesResponseUserDetailPtrInput interface {
	pulumi.Input

	ToOrganizationResourcePropertiesResponseUserDetailPtrOutput() OrganizationResourcePropertiesResponseUserDetailPtrOutput
	ToOrganizationResourcePropertiesResponseUserDetailPtrOutputWithContext(context.Context) OrganizationResourcePropertiesResponseUserDetailPtrOutput
}

type organizationResourcePropertiesResponseUserDetailPtrType OrganizationResourcePropertiesResponseUserDetailArgs

func OrganizationResourcePropertiesResponseUserDetailPtr(v *OrganizationResourcePropertiesResponseUserDetailArgs) OrganizationResourcePropertiesResponseUserDetailPtrInput {
	return (*organizationResourcePropertiesResponseUserDetailPtrType)(v)
}

func (*organizationResourcePropertiesResponseUserDetailPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**OrganizationResourcePropertiesResponseUserDetail)(nil)).Elem()
}

func (i *organizationResourcePropertiesResponseUserDetailPtrType) ToOrganizationResourcePropertiesResponseUserDetailPtrOutput() OrganizationResourcePropertiesResponseUserDetailPtrOutput {
	return i.ToOrganizationResourcePropertiesResponseUserDetailPtrOutputWithContext(context.Background())
}

func (i *organizationResourcePropertiesResponseUserDetailPtrType) ToOrganizationResourcePropertiesResponseUserDetailPtrOutputWithContext(ctx context.Context) OrganizationResourcePropertiesResponseUserDetailPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(OrganizationResourcePropertiesResponseUserDetailPtrOutput)
}

// Subscriber detail
type OrganizationResourcePropertiesResponseUserDetailOutput struct{ *pulumi.OutputState }

func (OrganizationResourcePropertiesResponseUserDetailOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*OrganizationResourcePropertiesResponseUserDetail)(nil)).Elem()
}

func (o OrganizationResourcePropertiesResponseUserDetailOutput) ToOrganizationResourcePropertiesResponseUserDetailOutput() OrganizationResourcePropertiesResponseUserDetailOutput {
	return o
}

func (o OrganizationResourcePropertiesResponseUserDetailOutput) ToOrganizationResourcePropertiesResponseUserDetailOutputWithContext(ctx context.Context) OrganizationResourcePropertiesResponseUserDetailOutput {
	return o
}

func (o OrganizationResourcePropertiesResponseUserDetailOutput) ToOrganizationResourcePropertiesResponseUserDetailPtrOutput() OrganizationResourcePropertiesResponseUserDetailPtrOutput {
	return o.ToOrganizationResourcePropertiesResponseUserDetailPtrOutputWithContext(context.Background())
}

func (o OrganizationResourcePropertiesResponseUserDetailOutput) ToOrganizationResourcePropertiesResponseUserDetailPtrOutputWithContext(ctx context.Context) OrganizationResourcePropertiesResponseUserDetailPtrOutput {
	return o.ApplyT(func(v OrganizationResourcePropertiesResponseUserDetail) *OrganizationResourcePropertiesResponseUserDetail {
		return &v
	}).(OrganizationResourcePropertiesResponseUserDetailPtrOutput)
}

// Email address
func (o OrganizationResourcePropertiesResponseUserDetailOutput) EmailAddress() pulumi.StringPtrOutput {
	return o.ApplyT(func(v OrganizationResourcePropertiesResponseUserDetail) *string { return v.EmailAddress }).(pulumi.StringPtrOutput)
}

// First name
func (o OrganizationResourcePropertiesResponseUserDetailOutput) FirstName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v OrganizationResourcePropertiesResponseUserDetail) *string { return v.FirstName }).(pulumi.StringPtrOutput)
}

// Last name
func (o OrganizationResourcePropertiesResponseUserDetailOutput) LastName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v OrganizationResourcePropertiesResponseUserDetail) *string { return v.LastName }).(pulumi.StringPtrOutput)
}

type OrganizationResourcePropertiesResponseUserDetailPtrOutput struct{ *pulumi.OutputState }

func (OrganizationResourcePropertiesResponseUserDetailPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**OrganizationResourcePropertiesResponseUserDetail)(nil)).Elem()
}

func (o OrganizationResourcePropertiesResponseUserDetailPtrOutput) ToOrganizationResourcePropertiesResponseUserDetailPtrOutput() OrganizationResourcePropertiesResponseUserDetailPtrOutput {
	return o
}

func (o OrganizationResourcePropertiesResponseUserDetailPtrOutput) ToOrganizationResourcePropertiesResponseUserDetailPtrOutputWithContext(ctx context.Context) OrganizationResourcePropertiesResponseUserDetailPtrOutput {
	return o
}

func (o OrganizationResourcePropertiesResponseUserDetailPtrOutput) Elem() OrganizationResourcePropertiesResponseUserDetailOutput {
	return o.ApplyT(func(v *OrganizationResourcePropertiesResponseUserDetail) OrganizationResourcePropertiesResponseUserDetail {
		return *v
	}).(OrganizationResourcePropertiesResponseUserDetailOutput)
}

// Email address
func (o OrganizationResourcePropertiesResponseUserDetailPtrOutput) EmailAddress() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *OrganizationResourcePropertiesResponseUserDetail) *string {
		if v == nil {
			return nil
		}
		return v.EmailAddress
	}).(pulumi.StringPtrOutput)
}

// First name
func (o OrganizationResourcePropertiesResponseUserDetailPtrOutput) FirstName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *OrganizationResourcePropertiesResponseUserDetail) *string {
		if v == nil {
			return nil
		}
		return v.FirstName
	}).(pulumi.StringPtrOutput)
}

// Last name
func (o OrganizationResourcePropertiesResponseUserDetailPtrOutput) LastName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *OrganizationResourcePropertiesResponseUserDetail) *string {
		if v == nil {
			return nil
		}
		return v.LastName
	}).(pulumi.StringPtrOutput)
}

// Subscriber detail
type OrganizationResourcePropertiesUserDetail struct {
	// Email address
	EmailAddress *string `pulumi:"emailAddress"`
	// First name
	FirstName *string `pulumi:"firstName"`
	// Last name
	LastName *string `pulumi:"lastName"`
}

// OrganizationResourcePropertiesUserDetailInput is an input type that accepts OrganizationResourcePropertiesUserDetailArgs and OrganizationResourcePropertiesUserDetailOutput values.
// You can construct a concrete instance of `OrganizationResourcePropertiesUserDetailInput` via:
//
//          OrganizationResourcePropertiesUserDetailArgs{...}
type OrganizationResourcePropertiesUserDetailInput interface {
	pulumi.Input

	ToOrganizationResourcePropertiesUserDetailOutput() OrganizationResourcePropertiesUserDetailOutput
	ToOrganizationResourcePropertiesUserDetailOutputWithContext(context.Context) OrganizationResourcePropertiesUserDetailOutput
}

// Subscriber detail
type OrganizationResourcePropertiesUserDetailArgs struct {
	// Email address
	EmailAddress pulumi.StringPtrInput `pulumi:"emailAddress"`
	// First name
	FirstName pulumi.StringPtrInput `pulumi:"firstName"`
	// Last name
	LastName pulumi.StringPtrInput `pulumi:"lastName"`
}

func (OrganizationResourcePropertiesUserDetailArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*OrganizationResourcePropertiesUserDetail)(nil)).Elem()
}

func (i OrganizationResourcePropertiesUserDetailArgs) ToOrganizationResourcePropertiesUserDetailOutput() OrganizationResourcePropertiesUserDetailOutput {
	return i.ToOrganizationResourcePropertiesUserDetailOutputWithContext(context.Background())
}

func (i OrganizationResourcePropertiesUserDetailArgs) ToOrganizationResourcePropertiesUserDetailOutputWithContext(ctx context.Context) OrganizationResourcePropertiesUserDetailOutput {
	return pulumi.ToOutputWithContext(ctx, i).(OrganizationResourcePropertiesUserDetailOutput)
}

func (i OrganizationResourcePropertiesUserDetailArgs) ToOrganizationResourcePropertiesUserDetailPtrOutput() OrganizationResourcePropertiesUserDetailPtrOutput {
	return i.ToOrganizationResourcePropertiesUserDetailPtrOutputWithContext(context.Background())
}

func (i OrganizationResourcePropertiesUserDetailArgs) ToOrganizationResourcePropertiesUserDetailPtrOutputWithContext(ctx context.Context) OrganizationResourcePropertiesUserDetailPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(OrganizationResourcePropertiesUserDetailOutput).ToOrganizationResourcePropertiesUserDetailPtrOutputWithContext(ctx)
}

// OrganizationResourcePropertiesUserDetailPtrInput is an input type that accepts OrganizationResourcePropertiesUserDetailArgs, OrganizationResourcePropertiesUserDetailPtr and OrganizationResourcePropertiesUserDetailPtrOutput values.
// You can construct a concrete instance of `OrganizationResourcePropertiesUserDetailPtrInput` via:
//
//          OrganizationResourcePropertiesUserDetailArgs{...}
//
//  or:
//
//          nil
type OrganizationResourcePropertiesUserDetailPtrInput interface {
	pulumi.Input

	ToOrganizationResourcePropertiesUserDetailPtrOutput() OrganizationResourcePropertiesUserDetailPtrOutput
	ToOrganizationResourcePropertiesUserDetailPtrOutputWithContext(context.Context) OrganizationResourcePropertiesUserDetailPtrOutput
}

type organizationResourcePropertiesUserDetailPtrType OrganizationResourcePropertiesUserDetailArgs

func OrganizationResourcePropertiesUserDetailPtr(v *OrganizationResourcePropertiesUserDetailArgs) OrganizationResourcePropertiesUserDetailPtrInput {
	return (*organizationResourcePropertiesUserDetailPtrType)(v)
}

func (*organizationResourcePropertiesUserDetailPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**OrganizationResourcePropertiesUserDetail)(nil)).Elem()
}

func (i *organizationResourcePropertiesUserDetailPtrType) ToOrganizationResourcePropertiesUserDetailPtrOutput() OrganizationResourcePropertiesUserDetailPtrOutput {
	return i.ToOrganizationResourcePropertiesUserDetailPtrOutputWithContext(context.Background())
}

func (i *organizationResourcePropertiesUserDetailPtrType) ToOrganizationResourcePropertiesUserDetailPtrOutputWithContext(ctx context.Context) OrganizationResourcePropertiesUserDetailPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(OrganizationResourcePropertiesUserDetailPtrOutput)
}

// Subscriber detail
type OrganizationResourcePropertiesUserDetailOutput struct{ *pulumi.OutputState }

func (OrganizationResourcePropertiesUserDetailOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*OrganizationResourcePropertiesUserDetail)(nil)).Elem()
}

func (o OrganizationResourcePropertiesUserDetailOutput) ToOrganizationResourcePropertiesUserDetailOutput() OrganizationResourcePropertiesUserDetailOutput {
	return o
}

func (o OrganizationResourcePropertiesUserDetailOutput) ToOrganizationResourcePropertiesUserDetailOutputWithContext(ctx context.Context) OrganizationResourcePropertiesUserDetailOutput {
	return o
}

func (o OrganizationResourcePropertiesUserDetailOutput) ToOrganizationResourcePropertiesUserDetailPtrOutput() OrganizationResourcePropertiesUserDetailPtrOutput {
	return o.ToOrganizationResourcePropertiesUserDetailPtrOutputWithContext(context.Background())
}

func (o OrganizationResourcePropertiesUserDetailOutput) ToOrganizationResourcePropertiesUserDetailPtrOutputWithContext(ctx context.Context) OrganizationResourcePropertiesUserDetailPtrOutput {
	return o.ApplyT(func(v OrganizationResourcePropertiesUserDetail) *OrganizationResourcePropertiesUserDetail {
		return &v
	}).(OrganizationResourcePropertiesUserDetailPtrOutput)
}

// Email address
func (o OrganizationResourcePropertiesUserDetailOutput) EmailAddress() pulumi.StringPtrOutput {
	return o.ApplyT(func(v OrganizationResourcePropertiesUserDetail) *string { return v.EmailAddress }).(pulumi.StringPtrOutput)
}

// First name
func (o OrganizationResourcePropertiesUserDetailOutput) FirstName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v OrganizationResourcePropertiesUserDetail) *string { return v.FirstName }).(pulumi.StringPtrOutput)
}

// Last name
func (o OrganizationResourcePropertiesUserDetailOutput) LastName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v OrganizationResourcePropertiesUserDetail) *string { return v.LastName }).(pulumi.StringPtrOutput)
}

type OrganizationResourcePropertiesUserDetailPtrOutput struct{ *pulumi.OutputState }

func (OrganizationResourcePropertiesUserDetailPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**OrganizationResourcePropertiesUserDetail)(nil)).Elem()
}

func (o OrganizationResourcePropertiesUserDetailPtrOutput) ToOrganizationResourcePropertiesUserDetailPtrOutput() OrganizationResourcePropertiesUserDetailPtrOutput {
	return o
}

func (o OrganizationResourcePropertiesUserDetailPtrOutput) ToOrganizationResourcePropertiesUserDetailPtrOutputWithContext(ctx context.Context) OrganizationResourcePropertiesUserDetailPtrOutput {
	return o
}

func (o OrganizationResourcePropertiesUserDetailPtrOutput) Elem() OrganizationResourcePropertiesUserDetailOutput {
	return o.ApplyT(func(v *OrganizationResourcePropertiesUserDetail) OrganizationResourcePropertiesUserDetail { return *v }).(OrganizationResourcePropertiesUserDetailOutput)
}

// Email address
func (o OrganizationResourcePropertiesUserDetailPtrOutput) EmailAddress() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *OrganizationResourcePropertiesUserDetail) *string {
		if v == nil {
			return nil
		}
		return v.EmailAddress
	}).(pulumi.StringPtrOutput)
}

// First name
func (o OrganizationResourcePropertiesUserDetailPtrOutput) FirstName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *OrganizationResourcePropertiesUserDetail) *string {
		if v == nil {
			return nil
		}
		return v.FirstName
	}).(pulumi.StringPtrOutput)
}

// Last name
func (o OrganizationResourcePropertiesUserDetailPtrOutput) LastName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *OrganizationResourcePropertiesUserDetail) *string {
		if v == nil {
			return nil
		}
		return v.LastName
	}).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(OrganizationResourcePropertiesOfferDetailOutput{})
	pulumi.RegisterOutputType(OrganizationResourcePropertiesOfferDetailPtrOutput{})
	pulumi.RegisterOutputType(OrganizationResourcePropertiesResponseOfferDetailOutput{})
	pulumi.RegisterOutputType(OrganizationResourcePropertiesResponseOfferDetailPtrOutput{})
	pulumi.RegisterOutputType(OrganizationResourcePropertiesResponseUserDetailOutput{})
	pulumi.RegisterOutputType(OrganizationResourcePropertiesResponseUserDetailPtrOutput{})
	pulumi.RegisterOutputType(OrganizationResourcePropertiesUserDetailOutput{})
	pulumi.RegisterOutputType(OrganizationResourcePropertiesUserDetailPtrOutput{})
}
