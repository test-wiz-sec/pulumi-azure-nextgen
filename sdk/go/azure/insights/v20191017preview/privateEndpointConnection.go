// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v20191017preview

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// A private endpoint connection
type PrivateEndpointConnection struct {
	pulumi.CustomResourceState

	// Azure resource name
	Name pulumi.StringOutput `pulumi:"name"`
	// Private endpoint which the connection belongs to.
	PrivateEndpoint PrivateEndpointPropertyResponsePtrOutput `pulumi:"privateEndpoint"`
	// Connection state of the private endpoint connection.
	PrivateLinkServiceConnectionState PrivateLinkServiceConnectionStatePropertyResponsePtrOutput `pulumi:"privateLinkServiceConnectionState"`
	// State of the private endpoint connection.
	ProvisioningState pulumi.StringOutput `pulumi:"provisioningState"`
	// Azure resource type
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewPrivateEndpointConnection registers a new resource with the given unique name, arguments, and options.
func NewPrivateEndpointConnection(ctx *pulumi.Context,
	name string, args *PrivateEndpointConnectionArgs, opts ...pulumi.ResourceOption) (*PrivateEndpointConnection, error) {
	if args == nil || args.PrivateEndpointConnectionName == nil {
		return nil, errors.New("missing required argument 'PrivateEndpointConnectionName'")
	}
	if args == nil || args.ResourceGroupName == nil {
		return nil, errors.New("missing required argument 'ResourceGroupName'")
	}
	if args == nil || args.ScopeName == nil {
		return nil, errors.New("missing required argument 'ScopeName'")
	}
	if args == nil {
		args = &PrivateEndpointConnectionArgs{}
	}
	var resource PrivateEndpointConnection
	err := ctx.RegisterResource("azure-nextgen:insights/v20191017preview:PrivateEndpointConnection", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetPrivateEndpointConnection gets an existing PrivateEndpointConnection resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetPrivateEndpointConnection(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *PrivateEndpointConnectionState, opts ...pulumi.ResourceOption) (*PrivateEndpointConnection, error) {
	var resource PrivateEndpointConnection
	err := ctx.ReadResource("azure-nextgen:insights/v20191017preview:PrivateEndpointConnection", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering PrivateEndpointConnection resources.
type privateEndpointConnectionState struct {
	// Azure resource name
	Name *string `pulumi:"name"`
	// Private endpoint which the connection belongs to.
	PrivateEndpoint *PrivateEndpointPropertyResponse `pulumi:"privateEndpoint"`
	// Connection state of the private endpoint connection.
	PrivateLinkServiceConnectionState *PrivateLinkServiceConnectionStatePropertyResponse `pulumi:"privateLinkServiceConnectionState"`
	// State of the private endpoint connection.
	ProvisioningState *string `pulumi:"provisioningState"`
	// Azure resource type
	Type *string `pulumi:"type"`
}

type PrivateEndpointConnectionState struct {
	// Azure resource name
	Name pulumi.StringPtrInput
	// Private endpoint which the connection belongs to.
	PrivateEndpoint PrivateEndpointPropertyResponsePtrInput
	// Connection state of the private endpoint connection.
	PrivateLinkServiceConnectionState PrivateLinkServiceConnectionStatePropertyResponsePtrInput
	// State of the private endpoint connection.
	ProvisioningState pulumi.StringPtrInput
	// Azure resource type
	Type pulumi.StringPtrInput
}

func (PrivateEndpointConnectionState) ElementType() reflect.Type {
	return reflect.TypeOf((*privateEndpointConnectionState)(nil)).Elem()
}

type privateEndpointConnectionArgs struct {
	// Private endpoint which the connection belongs to.
	PrivateEndpoint *PrivateEndpointProperty `pulumi:"privateEndpoint"`
	// The name of the private endpoint connection.
	PrivateEndpointConnectionName string `pulumi:"privateEndpointConnectionName"`
	// Connection state of the private endpoint connection.
	PrivateLinkServiceConnectionState *PrivateLinkServiceConnectionStateProperty `pulumi:"privateLinkServiceConnectionState"`
	// The name of the resource group.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The name of the Azure Monitor PrivateLinkScope resource.
	ScopeName string `pulumi:"scopeName"`
}

// The set of arguments for constructing a PrivateEndpointConnection resource.
type PrivateEndpointConnectionArgs struct {
	// Private endpoint which the connection belongs to.
	PrivateEndpoint PrivateEndpointPropertyPtrInput
	// The name of the private endpoint connection.
	PrivateEndpointConnectionName pulumi.StringInput
	// Connection state of the private endpoint connection.
	PrivateLinkServiceConnectionState PrivateLinkServiceConnectionStatePropertyPtrInput
	// The name of the resource group.
	ResourceGroupName pulumi.StringInput
	// The name of the Azure Monitor PrivateLinkScope resource.
	ScopeName pulumi.StringInput
}

func (PrivateEndpointConnectionArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*privateEndpointConnectionArgs)(nil)).Elem()
}
