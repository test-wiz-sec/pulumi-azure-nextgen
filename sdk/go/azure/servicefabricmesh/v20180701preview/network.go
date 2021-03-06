// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v20180701preview

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// This type describes a network resource.
type Network struct {
	pulumi.CustomResourceState

	// the address prefix for this network.
	AddressPrefix pulumi.StringOutput `pulumi:"addressPrefix"`
	// User readable description of the network.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// Configuration for public connectivity for this network.
	IngressConfig IngressConfigResponsePtrOutput `pulumi:"ingressConfig"`
	// The geo-location where the resource lives
	Location pulumi.StringOutput `pulumi:"location"`
	// The name of the resource
	Name pulumi.StringOutput `pulumi:"name"`
	// State of the resource.
	ProvisioningState pulumi.StringOutput `pulumi:"provisioningState"`
	// Resource tags.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// The type of the resource. Ex- Microsoft.Compute/virtualMachines or Microsoft.Storage/storageAccounts.
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewNetwork registers a new resource with the given unique name, arguments, and options.
func NewNetwork(ctx *pulumi.Context,
	name string, args *NetworkArgs, opts ...pulumi.ResourceOption) (*Network, error) {
	if args == nil || args.AddressPrefix == nil {
		return nil, errors.New("missing required argument 'AddressPrefix'")
	}
	if args == nil || args.Location == nil {
		return nil, errors.New("missing required argument 'Location'")
	}
	if args == nil || args.NetworkName == nil {
		return nil, errors.New("missing required argument 'NetworkName'")
	}
	if args == nil || args.ResourceGroupName == nil {
		return nil, errors.New("missing required argument 'ResourceGroupName'")
	}
	if args == nil {
		args = &NetworkArgs{}
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-nextgen:servicefabricmesh/v20180901preview:Network"),
		},
	})
	opts = append(opts, aliases)
	var resource Network
	err := ctx.RegisterResource("azure-nextgen:servicefabricmesh/v20180701preview:Network", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetNetwork gets an existing Network resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetNetwork(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *NetworkState, opts ...pulumi.ResourceOption) (*Network, error) {
	var resource Network
	err := ctx.ReadResource("azure-nextgen:servicefabricmesh/v20180701preview:Network", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Network resources.
type networkState struct {
	// the address prefix for this network.
	AddressPrefix *string `pulumi:"addressPrefix"`
	// User readable description of the network.
	Description *string `pulumi:"description"`
	// Configuration for public connectivity for this network.
	IngressConfig *IngressConfigResponse `pulumi:"ingressConfig"`
	// The geo-location where the resource lives
	Location *string `pulumi:"location"`
	// The name of the resource
	Name *string `pulumi:"name"`
	// State of the resource.
	ProvisioningState *string `pulumi:"provisioningState"`
	// Resource tags.
	Tags map[string]string `pulumi:"tags"`
	// The type of the resource. Ex- Microsoft.Compute/virtualMachines or Microsoft.Storage/storageAccounts.
	Type *string `pulumi:"type"`
}

type NetworkState struct {
	// the address prefix for this network.
	AddressPrefix pulumi.StringPtrInput
	// User readable description of the network.
	Description pulumi.StringPtrInput
	// Configuration for public connectivity for this network.
	IngressConfig IngressConfigResponsePtrInput
	// The geo-location where the resource lives
	Location pulumi.StringPtrInput
	// The name of the resource
	Name pulumi.StringPtrInput
	// State of the resource.
	ProvisioningState pulumi.StringPtrInput
	// Resource tags.
	Tags pulumi.StringMapInput
	// The type of the resource. Ex- Microsoft.Compute/virtualMachines or Microsoft.Storage/storageAccounts.
	Type pulumi.StringPtrInput
}

func (NetworkState) ElementType() reflect.Type {
	return reflect.TypeOf((*networkState)(nil)).Elem()
}

type networkArgs struct {
	// the address prefix for this network.
	AddressPrefix string `pulumi:"addressPrefix"`
	// User readable description of the network.
	Description *string `pulumi:"description"`
	// Configuration for public connectivity for this network.
	IngressConfig *IngressConfig `pulumi:"ingressConfig"`
	// The geo-location where the resource lives
	Location string `pulumi:"location"`
	// The identity of the network.
	NetworkName string `pulumi:"networkName"`
	// Azure resource group name
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// Resource tags.
	Tags map[string]string `pulumi:"tags"`
}

// The set of arguments for constructing a Network resource.
type NetworkArgs struct {
	// the address prefix for this network.
	AddressPrefix pulumi.StringInput
	// User readable description of the network.
	Description pulumi.StringPtrInput
	// Configuration for public connectivity for this network.
	IngressConfig IngressConfigPtrInput
	// The geo-location where the resource lives
	Location pulumi.StringInput
	// The identity of the network.
	NetworkName pulumi.StringInput
	// Azure resource group name
	ResourceGroupName pulumi.StringInput
	// Resource tags.
	Tags pulumi.StringMapInput
}

func (NetworkArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*networkArgs)(nil)).Elem()
}
