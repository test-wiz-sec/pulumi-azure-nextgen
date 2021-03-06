// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v20181031preview

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Resource information with extended details.
type DedicatedHsm struct {
	pulumi.CustomResourceState

	// The supported Azure location where the dedicated HSM should be created.
	Location pulumi.StringOutput `pulumi:"location"`
	// The name of the dedicated HSM.
	Name pulumi.StringOutput `pulumi:"name"`
	// Specifies the network interfaces of the dedicated hsm.
	NetworkProfile NetworkProfileResponsePtrOutput `pulumi:"networkProfile"`
	// Provisioning state.
	ProvisioningState pulumi.StringOutput `pulumi:"provisioningState"`
	// SKU details
	Sku SkuResponseOutput `pulumi:"sku"`
	// This field will be used when RP does not support Availability zones.
	StampId pulumi.StringPtrOutput `pulumi:"stampId"`
	// Resource Status Message.
	StatusMessage pulumi.StringOutput `pulumi:"statusMessage"`
	// Resource tags
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// The resource type of the dedicated HSM.
	Type pulumi.StringOutput `pulumi:"type"`
	// The Dedicated Hsm zones.
	Zones pulumi.StringArrayOutput `pulumi:"zones"`
}

// NewDedicatedHsm registers a new resource with the given unique name, arguments, and options.
func NewDedicatedHsm(ctx *pulumi.Context,
	name string, args *DedicatedHsmArgs, opts ...pulumi.ResourceOption) (*DedicatedHsm, error) {
	if args == nil || args.Location == nil {
		return nil, errors.New("missing required argument 'Location'")
	}
	if args == nil || args.Name == nil {
		return nil, errors.New("missing required argument 'Name'")
	}
	if args == nil || args.ResourceGroupName == nil {
		return nil, errors.New("missing required argument 'ResourceGroupName'")
	}
	if args == nil || args.Sku == nil {
		return nil, errors.New("missing required argument 'Sku'")
	}
	if args == nil {
		args = &DedicatedHsmArgs{}
	}
	var resource DedicatedHsm
	err := ctx.RegisterResource("azure-nextgen:hardwaresecuritymodules/v20181031preview:DedicatedHsm", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetDedicatedHsm gets an existing DedicatedHsm resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetDedicatedHsm(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *DedicatedHsmState, opts ...pulumi.ResourceOption) (*DedicatedHsm, error) {
	var resource DedicatedHsm
	err := ctx.ReadResource("azure-nextgen:hardwaresecuritymodules/v20181031preview:DedicatedHsm", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering DedicatedHsm resources.
type dedicatedHsmState struct {
	// The supported Azure location where the dedicated HSM should be created.
	Location *string `pulumi:"location"`
	// The name of the dedicated HSM.
	Name *string `pulumi:"name"`
	// Specifies the network interfaces of the dedicated hsm.
	NetworkProfile *NetworkProfileResponse `pulumi:"networkProfile"`
	// Provisioning state.
	ProvisioningState *string `pulumi:"provisioningState"`
	// SKU details
	Sku *SkuResponse `pulumi:"sku"`
	// This field will be used when RP does not support Availability zones.
	StampId *string `pulumi:"stampId"`
	// Resource Status Message.
	StatusMessage *string `pulumi:"statusMessage"`
	// Resource tags
	Tags map[string]string `pulumi:"tags"`
	// The resource type of the dedicated HSM.
	Type *string `pulumi:"type"`
	// The Dedicated Hsm zones.
	Zones []string `pulumi:"zones"`
}

type DedicatedHsmState struct {
	// The supported Azure location where the dedicated HSM should be created.
	Location pulumi.StringPtrInput
	// The name of the dedicated HSM.
	Name pulumi.StringPtrInput
	// Specifies the network interfaces of the dedicated hsm.
	NetworkProfile NetworkProfileResponsePtrInput
	// Provisioning state.
	ProvisioningState pulumi.StringPtrInput
	// SKU details
	Sku SkuResponsePtrInput
	// This field will be used when RP does not support Availability zones.
	StampId pulumi.StringPtrInput
	// Resource Status Message.
	StatusMessage pulumi.StringPtrInput
	// Resource tags
	Tags pulumi.StringMapInput
	// The resource type of the dedicated HSM.
	Type pulumi.StringPtrInput
	// The Dedicated Hsm zones.
	Zones pulumi.StringArrayInput
}

func (DedicatedHsmState) ElementType() reflect.Type {
	return reflect.TypeOf((*dedicatedHsmState)(nil)).Elem()
}

type dedicatedHsmArgs struct {
	// The supported Azure location where the dedicated HSM should be created.
	Location string `pulumi:"location"`
	// Name of the dedicated Hsm
	Name string `pulumi:"name"`
	// Specifies the network interfaces of the dedicated hsm.
	NetworkProfile *NetworkProfile `pulumi:"networkProfile"`
	// The name of the Resource Group to which the resource belongs.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// SKU details
	Sku Sku `pulumi:"sku"`
	// This field will be used when RP does not support Availability zones.
	StampId *string `pulumi:"stampId"`
	// Resource tags
	Tags map[string]string `pulumi:"tags"`
	// The Dedicated Hsm zones.
	Zones []string `pulumi:"zones"`
}

// The set of arguments for constructing a DedicatedHsm resource.
type DedicatedHsmArgs struct {
	// The supported Azure location where the dedicated HSM should be created.
	Location pulumi.StringInput
	// Name of the dedicated Hsm
	Name pulumi.StringInput
	// Specifies the network interfaces of the dedicated hsm.
	NetworkProfile NetworkProfilePtrInput
	// The name of the Resource Group to which the resource belongs.
	ResourceGroupName pulumi.StringInput
	// SKU details
	Sku SkuInput
	// This field will be used when RP does not support Availability zones.
	StampId pulumi.StringPtrInput
	// Resource tags
	Tags pulumi.StringMapInput
	// The Dedicated Hsm zones.
	Zones pulumi.StringArrayInput
}

func (DedicatedHsmArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*dedicatedHsmArgs)(nil)).Elem()
}
