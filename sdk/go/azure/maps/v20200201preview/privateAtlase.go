// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v20200201preview

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// An Azure resource which represents which will provision the ability to create private location data.
type PrivateAtlase struct {
	pulumi.CustomResourceState

	// The geo-location where the resource lives
	Location pulumi.StringOutput `pulumi:"location"`
	// The name of the resource
	Name pulumi.StringOutput `pulumi:"name"`
	// The Private Atlas resource properties.
	Properties PrivateAtlasPropertiesResponseOutput `pulumi:"properties"`
	// Resource tags.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewPrivateAtlase registers a new resource with the given unique name, arguments, and options.
func NewPrivateAtlase(ctx *pulumi.Context,
	name string, args *PrivateAtlaseArgs, opts ...pulumi.ResourceOption) (*PrivateAtlase, error) {
	if args == nil || args.AccountName == nil {
		return nil, errors.New("missing required argument 'AccountName'")
	}
	if args == nil || args.Location == nil {
		return nil, errors.New("missing required argument 'Location'")
	}
	if args == nil || args.PrivateAtlasName == nil {
		return nil, errors.New("missing required argument 'PrivateAtlasName'")
	}
	if args == nil || args.ResourceGroupName == nil {
		return nil, errors.New("missing required argument 'ResourceGroupName'")
	}
	if args == nil {
		args = &PrivateAtlaseArgs{}
	}
	var resource PrivateAtlase
	err := ctx.RegisterResource("azure-nextgen:maps/v20200201preview:PrivateAtlase", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetPrivateAtlase gets an existing PrivateAtlase resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetPrivateAtlase(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *PrivateAtlaseState, opts ...pulumi.ResourceOption) (*PrivateAtlase, error) {
	var resource PrivateAtlase
	err := ctx.ReadResource("azure-nextgen:maps/v20200201preview:PrivateAtlase", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering PrivateAtlase resources.
type privateAtlaseState struct {
	// The geo-location where the resource lives
	Location *string `pulumi:"location"`
	// The name of the resource
	Name *string `pulumi:"name"`
	// The Private Atlas resource properties.
	Properties *PrivateAtlasPropertiesResponse `pulumi:"properties"`
	// Resource tags.
	Tags map[string]string `pulumi:"tags"`
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type *string `pulumi:"type"`
}

type PrivateAtlaseState struct {
	// The geo-location where the resource lives
	Location pulumi.StringPtrInput
	// The name of the resource
	Name pulumi.StringPtrInput
	// The Private Atlas resource properties.
	Properties PrivateAtlasPropertiesResponsePtrInput
	// Resource tags.
	Tags pulumi.StringMapInput
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type pulumi.StringPtrInput
}

func (PrivateAtlaseState) ElementType() reflect.Type {
	return reflect.TypeOf((*privateAtlaseState)(nil)).Elem()
}

type privateAtlaseArgs struct {
	// The name of the Maps Account.
	AccountName string `pulumi:"accountName"`
	// The location of the resource.
	Location string `pulumi:"location"`
	// The name of the Private Atlas instance.
	PrivateAtlasName string `pulumi:"privateAtlasName"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// Gets or sets a list of key value pairs that describe the resource. These tags can be used in viewing and grouping this resource (across resource groups). A maximum of 15 tags can be provided for a resource. Each tag must have a key no greater than 128 characters and value no greater than 256 characters.
	Tags map[string]string `pulumi:"tags"`
}

// The set of arguments for constructing a PrivateAtlase resource.
type PrivateAtlaseArgs struct {
	// The name of the Maps Account.
	AccountName pulumi.StringInput
	// The location of the resource.
	Location pulumi.StringInput
	// The name of the Private Atlas instance.
	PrivateAtlasName pulumi.StringInput
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput
	// Gets or sets a list of key value pairs that describe the resource. These tags can be used in viewing and grouping this resource (across resource groups). A maximum of 15 tags can be provided for a resource. Each tag must have a key no greater than 128 characters and value no greater than 256 characters.
	Tags pulumi.StringMapInput
}

func (PrivateAtlaseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*privateAtlaseArgs)(nil)).Elem()
}
