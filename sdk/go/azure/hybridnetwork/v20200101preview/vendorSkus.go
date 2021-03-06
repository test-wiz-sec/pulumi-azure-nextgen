// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v20200101preview

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Sku sub resource.
type VendorSkus struct {
	pulumi.CustomResourceState

	// The sku deployment mode.
	DeploymentMode pulumi.StringPtrOutput `pulumi:"deploymentMode"`
	// The parameters for the managed application to be supplied by the vendor.
	ManagedApplicationParameters pulumi.AnyOutput `pulumi:"managedApplicationParameters"`
	// The template for the managed application deployment.
	ManagedApplicationTemplate pulumi.AnyOutput `pulumi:"managedApplicationTemplate"`
	// The name of the resource
	Name pulumi.StringOutput `pulumi:"name"`
	// The template definition of the network function.
	NetworkFunctionTemplate NetworkFunctionTemplateResponsePtrOutput `pulumi:"networkFunctionTemplate"`
	// Indicates if the vendor sku is in preview mode.
	Preview pulumi.BoolPtrOutput `pulumi:"preview"`
	// The provisioning state of the vendor sku sub resource.
	ProvisioningState pulumi.StringOutput `pulumi:"provisioningState"`
	// The sku type.
	SkuType pulumi.StringPtrOutput `pulumi:"skuType"`
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewVendorSkus registers a new resource with the given unique name, arguments, and options.
func NewVendorSkus(ctx *pulumi.Context,
	name string, args *VendorSkusArgs, opts ...pulumi.ResourceOption) (*VendorSkus, error) {
	if args == nil || args.SkuName == nil {
		return nil, errors.New("missing required argument 'SkuName'")
	}
	if args == nil || args.VendorName == nil {
		return nil, errors.New("missing required argument 'VendorName'")
	}
	if args == nil {
		args = &VendorSkusArgs{}
	}
	var resource VendorSkus
	err := ctx.RegisterResource("azure-nextgen:hybridnetwork/v20200101preview:VendorSkus", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetVendorSkus gets an existing VendorSkus resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetVendorSkus(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *VendorSkusState, opts ...pulumi.ResourceOption) (*VendorSkus, error) {
	var resource VendorSkus
	err := ctx.ReadResource("azure-nextgen:hybridnetwork/v20200101preview:VendorSkus", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering VendorSkus resources.
type vendorSkusState struct {
	// The sku deployment mode.
	DeploymentMode *string `pulumi:"deploymentMode"`
	// The parameters for the managed application to be supplied by the vendor.
	ManagedApplicationParameters interface{} `pulumi:"managedApplicationParameters"`
	// The template for the managed application deployment.
	ManagedApplicationTemplate interface{} `pulumi:"managedApplicationTemplate"`
	// The name of the resource
	Name *string `pulumi:"name"`
	// The template definition of the network function.
	NetworkFunctionTemplate *NetworkFunctionTemplateResponse `pulumi:"networkFunctionTemplate"`
	// Indicates if the vendor sku is in preview mode.
	Preview *bool `pulumi:"preview"`
	// The provisioning state of the vendor sku sub resource.
	ProvisioningState *string `pulumi:"provisioningState"`
	// The sku type.
	SkuType *string `pulumi:"skuType"`
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type *string `pulumi:"type"`
}

type VendorSkusState struct {
	// The sku deployment mode.
	DeploymentMode pulumi.StringPtrInput
	// The parameters for the managed application to be supplied by the vendor.
	ManagedApplicationParameters pulumi.Input
	// The template for the managed application deployment.
	ManagedApplicationTemplate pulumi.Input
	// The name of the resource
	Name pulumi.StringPtrInput
	// The template definition of the network function.
	NetworkFunctionTemplate NetworkFunctionTemplateResponsePtrInput
	// Indicates if the vendor sku is in preview mode.
	Preview pulumi.BoolPtrInput
	// The provisioning state of the vendor sku sub resource.
	ProvisioningState pulumi.StringPtrInput
	// The sku type.
	SkuType pulumi.StringPtrInput
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type pulumi.StringPtrInput
}

func (VendorSkusState) ElementType() reflect.Type {
	return reflect.TypeOf((*vendorSkusState)(nil)).Elem()
}

type vendorSkusArgs struct {
	// The sku deployment mode.
	DeploymentMode *string `pulumi:"deploymentMode"`
	// The parameters for the managed application to be supplied by the vendor.
	ManagedApplicationParameters interface{} `pulumi:"managedApplicationParameters"`
	// The template for the managed application deployment.
	ManagedApplicationTemplate interface{} `pulumi:"managedApplicationTemplate"`
	// The template definition of the network function.
	NetworkFunctionTemplate *NetworkFunctionTemplate `pulumi:"networkFunctionTemplate"`
	// Indicates if the vendor sku is in preview mode.
	Preview *bool `pulumi:"preview"`
	// The name of the sku.
	SkuName string `pulumi:"skuName"`
	// The sku type.
	SkuType *string `pulumi:"skuType"`
	// The name of the vendor.
	VendorName string `pulumi:"vendorName"`
}

// The set of arguments for constructing a VendorSkus resource.
type VendorSkusArgs struct {
	// The sku deployment mode.
	DeploymentMode pulumi.StringPtrInput
	// The parameters for the managed application to be supplied by the vendor.
	ManagedApplicationParameters pulumi.Input
	// The template for the managed application deployment.
	ManagedApplicationTemplate pulumi.Input
	// The template definition of the network function.
	NetworkFunctionTemplate NetworkFunctionTemplatePtrInput
	// Indicates if the vendor sku is in preview mode.
	Preview pulumi.BoolPtrInput
	// The name of the sku.
	SkuName pulumi.StringInput
	// The sku type.
	SkuType pulumi.StringPtrInput
	// The name of the vendor.
	VendorName pulumi.StringInput
}

func (VendorSkusArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*vendorSkusArgs)(nil)).Elem()
}
