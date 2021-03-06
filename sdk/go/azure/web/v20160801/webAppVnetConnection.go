// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v20160801

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Virtual Network information contract.
type WebAppVnetConnection struct {
	pulumi.CustomResourceState

	// A certificate file (.cer) blob containing the public key of the private key used to authenticate a
	// Point-To-Site VPN connection.
	CertBlob pulumi.StringPtrOutput `pulumi:"certBlob"`
	// The client certificate thumbprint.
	CertThumbprint pulumi.StringOutput `pulumi:"certThumbprint"`
	// DNS servers to be used by this Virtual Network. This should be a comma-separated list of IP addresses.
	DnsServers pulumi.StringPtrOutput `pulumi:"dnsServers"`
	// Kind of resource.
	Kind pulumi.StringPtrOutput `pulumi:"kind"`
	// Resource Name.
	Name pulumi.StringOutput `pulumi:"name"`
	// <code>true</code> if a resync is required; otherwise, <code>false</code>.
	ResyncRequired pulumi.BoolOutput `pulumi:"resyncRequired"`
	// The routes that this Virtual Network connection uses.
	Routes VnetRouteResponseArrayOutput `pulumi:"routes"`
	// Resource type.
	Type pulumi.StringOutput `pulumi:"type"`
	// The Virtual Network's resource ID.
	VnetResourceId pulumi.StringPtrOutput `pulumi:"vnetResourceId"`
}

// NewWebAppVnetConnection registers a new resource with the given unique name, arguments, and options.
func NewWebAppVnetConnection(ctx *pulumi.Context,
	name string, args *WebAppVnetConnectionArgs, opts ...pulumi.ResourceOption) (*WebAppVnetConnection, error) {
	if args == nil || args.Name == nil {
		return nil, errors.New("missing required argument 'Name'")
	}
	if args == nil || args.ResourceGroupName == nil {
		return nil, errors.New("missing required argument 'ResourceGroupName'")
	}
	if args == nil || args.VnetName == nil {
		return nil, errors.New("missing required argument 'VnetName'")
	}
	if args == nil {
		args = &WebAppVnetConnectionArgs{}
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-nextgen:web/latest:WebAppVnetConnection"),
		},
		{
			Type: pulumi.String("azure-nextgen:web/v20150801:WebAppVnetConnection"),
		},
		{
			Type: pulumi.String("azure-nextgen:web/v20180201:WebAppVnetConnection"),
		},
		{
			Type: pulumi.String("azure-nextgen:web/v20181101:WebAppVnetConnection"),
		},
		{
			Type: pulumi.String("azure-nextgen:web/v20190801:WebAppVnetConnection"),
		},
		{
			Type: pulumi.String("azure-nextgen:web/v20200601:WebAppVnetConnection"),
		},
	})
	opts = append(opts, aliases)
	var resource WebAppVnetConnection
	err := ctx.RegisterResource("azure-nextgen:web/v20160801:WebAppVnetConnection", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetWebAppVnetConnection gets an existing WebAppVnetConnection resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetWebAppVnetConnection(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *WebAppVnetConnectionState, opts ...pulumi.ResourceOption) (*WebAppVnetConnection, error) {
	var resource WebAppVnetConnection
	err := ctx.ReadResource("azure-nextgen:web/v20160801:WebAppVnetConnection", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering WebAppVnetConnection resources.
type webAppVnetConnectionState struct {
	// A certificate file (.cer) blob containing the public key of the private key used to authenticate a
	// Point-To-Site VPN connection.
	CertBlob *string `pulumi:"certBlob"`
	// The client certificate thumbprint.
	CertThumbprint *string `pulumi:"certThumbprint"`
	// DNS servers to be used by this Virtual Network. This should be a comma-separated list of IP addresses.
	DnsServers *string `pulumi:"dnsServers"`
	// Kind of resource.
	Kind *string `pulumi:"kind"`
	// Resource Name.
	Name *string `pulumi:"name"`
	// <code>true</code> if a resync is required; otherwise, <code>false</code>.
	ResyncRequired *bool `pulumi:"resyncRequired"`
	// The routes that this Virtual Network connection uses.
	Routes []VnetRouteResponse `pulumi:"routes"`
	// Resource type.
	Type *string `pulumi:"type"`
	// The Virtual Network's resource ID.
	VnetResourceId *string `pulumi:"vnetResourceId"`
}

type WebAppVnetConnectionState struct {
	// A certificate file (.cer) blob containing the public key of the private key used to authenticate a
	// Point-To-Site VPN connection.
	CertBlob pulumi.StringPtrInput
	// The client certificate thumbprint.
	CertThumbprint pulumi.StringPtrInput
	// DNS servers to be used by this Virtual Network. This should be a comma-separated list of IP addresses.
	DnsServers pulumi.StringPtrInput
	// Kind of resource.
	Kind pulumi.StringPtrInput
	// Resource Name.
	Name pulumi.StringPtrInput
	// <code>true</code> if a resync is required; otherwise, <code>false</code>.
	ResyncRequired pulumi.BoolPtrInput
	// The routes that this Virtual Network connection uses.
	Routes VnetRouteResponseArrayInput
	// Resource type.
	Type pulumi.StringPtrInput
	// The Virtual Network's resource ID.
	VnetResourceId pulumi.StringPtrInput
}

func (WebAppVnetConnectionState) ElementType() reflect.Type {
	return reflect.TypeOf((*webAppVnetConnectionState)(nil)).Elem()
}

type webAppVnetConnectionArgs struct {
	// A certificate file (.cer) blob containing the public key of the private key used to authenticate a
	// Point-To-Site VPN connection.
	CertBlob *string `pulumi:"certBlob"`
	// DNS servers to be used by this Virtual Network. This should be a comma-separated list of IP addresses.
	DnsServers *string `pulumi:"dnsServers"`
	// Kind of resource.
	Kind *string `pulumi:"kind"`
	// Name of the app.
	Name string `pulumi:"name"`
	// Name of the resource group to which the resource belongs.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// Name of an existing Virtual Network.
	VnetName string `pulumi:"vnetName"`
	// The Virtual Network's resource ID.
	VnetResourceId *string `pulumi:"vnetResourceId"`
}

// The set of arguments for constructing a WebAppVnetConnection resource.
type WebAppVnetConnectionArgs struct {
	// A certificate file (.cer) blob containing the public key of the private key used to authenticate a
	// Point-To-Site VPN connection.
	CertBlob pulumi.StringPtrInput
	// DNS servers to be used by this Virtual Network. This should be a comma-separated list of IP addresses.
	DnsServers pulumi.StringPtrInput
	// Kind of resource.
	Kind pulumi.StringPtrInput
	// Name of the app.
	Name pulumi.StringInput
	// Name of the resource group to which the resource belongs.
	ResourceGroupName pulumi.StringInput
	// Name of an existing Virtual Network.
	VnetName pulumi.StringInput
	// The Virtual Network's resource ID.
	VnetResourceId pulumi.StringPtrInput
}

func (WebAppVnetConnectionArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*webAppVnetConnectionArgs)(nil)).Elem()
}
