// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v20150615

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Application gateway resource
type ApplicationGateway struct {
	pulumi.CustomResourceState

	// Backend address pool of the application gateway resource.
	BackendAddressPools ApplicationGatewayBackendAddressPoolResponseArrayOutput `pulumi:"backendAddressPools"`
	// Backend http settings of the application gateway resource.
	BackendHttpSettingsCollection ApplicationGatewayBackendHttpSettingsResponseArrayOutput `pulumi:"backendHttpSettingsCollection"`
	// A unique read-only string that changes whenever the resource is updated.
	Etag pulumi.StringPtrOutput `pulumi:"etag"`
	// Frontend IP addresses of the application gateway resource.
	FrontendIPConfigurations ApplicationGatewayFrontendIPConfigurationResponseArrayOutput `pulumi:"frontendIPConfigurations"`
	// Frontend ports of the application gateway resource.
	FrontendPorts ApplicationGatewayFrontendPortResponseArrayOutput `pulumi:"frontendPorts"`
	// Gets or sets subnets of application gateway resource
	GatewayIPConfigurations ApplicationGatewayIPConfigurationResponseArrayOutput `pulumi:"gatewayIPConfigurations"`
	// Http listeners of the application gateway resource.
	HttpListeners ApplicationGatewayHttpListenerResponseArrayOutput `pulumi:"httpListeners"`
	// Resource location.
	Location pulumi.StringPtrOutput `pulumi:"location"`
	// Resource name.
	Name pulumi.StringOutput `pulumi:"name"`
	// Operational state of the application gateway resource. Possible values are: 'Stopped', 'Started', 'Running', and 'Stopping'.
	OperationalState pulumi.StringOutput `pulumi:"operationalState"`
	// Probes of the application gateway resource.
	Probes ApplicationGatewayProbeResponseArrayOutput `pulumi:"probes"`
	// Provisioning state of the application gateway resource. Possible values are: 'Updating', 'Deleting', and 'Failed'.
	ProvisioningState pulumi.StringPtrOutput `pulumi:"provisioningState"`
	// Request routing rules of the application gateway resource.
	RequestRoutingRules ApplicationGatewayRequestRoutingRuleResponseArrayOutput `pulumi:"requestRoutingRules"`
	// Resource GUID property of the application gateway resource.
	ResourceGuid pulumi.StringPtrOutput `pulumi:"resourceGuid"`
	// SKU of the application gateway resource.
	Sku ApplicationGatewaySkuResponsePtrOutput `pulumi:"sku"`
	// SSL certificates of the application gateway resource.
	SslCertificates ApplicationGatewaySslCertificateResponseArrayOutput `pulumi:"sslCertificates"`
	// Resource tags.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// Resource type.
	Type pulumi.StringOutput `pulumi:"type"`
	// URL path map of the application gateway resource.
	UrlPathMaps ApplicationGatewayUrlPathMapResponseArrayOutput `pulumi:"urlPathMaps"`
}

// NewApplicationGateway registers a new resource with the given unique name, arguments, and options.
func NewApplicationGateway(ctx *pulumi.Context,
	name string, args *ApplicationGatewayArgs, opts ...pulumi.ResourceOption) (*ApplicationGateway, error) {
	if args == nil || args.ApplicationGatewayName == nil {
		return nil, errors.New("missing required argument 'ApplicationGatewayName'")
	}
	if args == nil || args.ResourceGroupName == nil {
		return nil, errors.New("missing required argument 'ResourceGroupName'")
	}
	if args == nil {
		args = &ApplicationGatewayArgs{}
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-nextgen:network/latest:ApplicationGateway"),
		},
		{
			Type: pulumi.String("azure-nextgen:network/v20150501preview:ApplicationGateway"),
		},
		{
			Type: pulumi.String("azure-nextgen:network/v20160330:ApplicationGateway"),
		},
		{
			Type: pulumi.String("azure-nextgen:network/v20160601:ApplicationGateway"),
		},
		{
			Type: pulumi.String("azure-nextgen:network/v20160901:ApplicationGateway"),
		},
		{
			Type: pulumi.String("azure-nextgen:network/v20161201:ApplicationGateway"),
		},
		{
			Type: pulumi.String("azure-nextgen:network/v20170301:ApplicationGateway"),
		},
		{
			Type: pulumi.String("azure-nextgen:network/v20170601:ApplicationGateway"),
		},
		{
			Type: pulumi.String("azure-nextgen:network/v20170801:ApplicationGateway"),
		},
		{
			Type: pulumi.String("azure-nextgen:network/v20170901:ApplicationGateway"),
		},
		{
			Type: pulumi.String("azure-nextgen:network/v20171001:ApplicationGateway"),
		},
		{
			Type: pulumi.String("azure-nextgen:network/v20171101:ApplicationGateway"),
		},
		{
			Type: pulumi.String("azure-nextgen:network/v20180101:ApplicationGateway"),
		},
		{
			Type: pulumi.String("azure-nextgen:network/v20180201:ApplicationGateway"),
		},
		{
			Type: pulumi.String("azure-nextgen:network/v20180401:ApplicationGateway"),
		},
		{
			Type: pulumi.String("azure-nextgen:network/v20180601:ApplicationGateway"),
		},
		{
			Type: pulumi.String("azure-nextgen:network/v20180701:ApplicationGateway"),
		},
		{
			Type: pulumi.String("azure-nextgen:network/v20180801:ApplicationGateway"),
		},
		{
			Type: pulumi.String("azure-nextgen:network/v20181001:ApplicationGateway"),
		},
		{
			Type: pulumi.String("azure-nextgen:network/v20181101:ApplicationGateway"),
		},
		{
			Type: pulumi.String("azure-nextgen:network/v20181201:ApplicationGateway"),
		},
		{
			Type: pulumi.String("azure-nextgen:network/v20190201:ApplicationGateway"),
		},
		{
			Type: pulumi.String("azure-nextgen:network/v20190401:ApplicationGateway"),
		},
		{
			Type: pulumi.String("azure-nextgen:network/v20190601:ApplicationGateway"),
		},
		{
			Type: pulumi.String("azure-nextgen:network/v20190701:ApplicationGateway"),
		},
		{
			Type: pulumi.String("azure-nextgen:network/v20190801:ApplicationGateway"),
		},
		{
			Type: pulumi.String("azure-nextgen:network/v20190901:ApplicationGateway"),
		},
		{
			Type: pulumi.String("azure-nextgen:network/v20191101:ApplicationGateway"),
		},
		{
			Type: pulumi.String("azure-nextgen:network/v20191201:ApplicationGateway"),
		},
		{
			Type: pulumi.String("azure-nextgen:network/v20200301:ApplicationGateway"),
		},
		{
			Type: pulumi.String("azure-nextgen:network/v20200401:ApplicationGateway"),
		},
		{
			Type: pulumi.String("azure-nextgen:network/v20200501:ApplicationGateway"),
		},
		{
			Type: pulumi.String("azure-nextgen:network/v20200601:ApplicationGateway"),
		},
		{
			Type: pulumi.String("azure-nextgen:network/v20200701:ApplicationGateway"),
		},
	})
	opts = append(opts, aliases)
	var resource ApplicationGateway
	err := ctx.RegisterResource("azure-nextgen:network/v20150615:ApplicationGateway", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetApplicationGateway gets an existing ApplicationGateway resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetApplicationGateway(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ApplicationGatewayState, opts ...pulumi.ResourceOption) (*ApplicationGateway, error) {
	var resource ApplicationGateway
	err := ctx.ReadResource("azure-nextgen:network/v20150615:ApplicationGateway", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ApplicationGateway resources.
type applicationGatewayState struct {
	// Backend address pool of the application gateway resource.
	BackendAddressPools []ApplicationGatewayBackendAddressPoolResponse `pulumi:"backendAddressPools"`
	// Backend http settings of the application gateway resource.
	BackendHttpSettingsCollection []ApplicationGatewayBackendHttpSettingsResponse `pulumi:"backendHttpSettingsCollection"`
	// A unique read-only string that changes whenever the resource is updated.
	Etag *string `pulumi:"etag"`
	// Frontend IP addresses of the application gateway resource.
	FrontendIPConfigurations []ApplicationGatewayFrontendIPConfigurationResponse `pulumi:"frontendIPConfigurations"`
	// Frontend ports of the application gateway resource.
	FrontendPorts []ApplicationGatewayFrontendPortResponse `pulumi:"frontendPorts"`
	// Gets or sets subnets of application gateway resource
	GatewayIPConfigurations []ApplicationGatewayIPConfigurationResponse `pulumi:"gatewayIPConfigurations"`
	// Http listeners of the application gateway resource.
	HttpListeners []ApplicationGatewayHttpListenerResponse `pulumi:"httpListeners"`
	// Resource location.
	Location *string `pulumi:"location"`
	// Resource name.
	Name *string `pulumi:"name"`
	// Operational state of the application gateway resource. Possible values are: 'Stopped', 'Started', 'Running', and 'Stopping'.
	OperationalState *string `pulumi:"operationalState"`
	// Probes of the application gateway resource.
	Probes []ApplicationGatewayProbeResponse `pulumi:"probes"`
	// Provisioning state of the application gateway resource. Possible values are: 'Updating', 'Deleting', and 'Failed'.
	ProvisioningState *string `pulumi:"provisioningState"`
	// Request routing rules of the application gateway resource.
	RequestRoutingRules []ApplicationGatewayRequestRoutingRuleResponse `pulumi:"requestRoutingRules"`
	// Resource GUID property of the application gateway resource.
	ResourceGuid *string `pulumi:"resourceGuid"`
	// SKU of the application gateway resource.
	Sku *ApplicationGatewaySkuResponse `pulumi:"sku"`
	// SSL certificates of the application gateway resource.
	SslCertificates []ApplicationGatewaySslCertificateResponse `pulumi:"sslCertificates"`
	// Resource tags.
	Tags map[string]string `pulumi:"tags"`
	// Resource type.
	Type *string `pulumi:"type"`
	// URL path map of the application gateway resource.
	UrlPathMaps []ApplicationGatewayUrlPathMapResponse `pulumi:"urlPathMaps"`
}

type ApplicationGatewayState struct {
	// Backend address pool of the application gateway resource.
	BackendAddressPools ApplicationGatewayBackendAddressPoolResponseArrayInput
	// Backend http settings of the application gateway resource.
	BackendHttpSettingsCollection ApplicationGatewayBackendHttpSettingsResponseArrayInput
	// A unique read-only string that changes whenever the resource is updated.
	Etag pulumi.StringPtrInput
	// Frontend IP addresses of the application gateway resource.
	FrontendIPConfigurations ApplicationGatewayFrontendIPConfigurationResponseArrayInput
	// Frontend ports of the application gateway resource.
	FrontendPorts ApplicationGatewayFrontendPortResponseArrayInput
	// Gets or sets subnets of application gateway resource
	GatewayIPConfigurations ApplicationGatewayIPConfigurationResponseArrayInput
	// Http listeners of the application gateway resource.
	HttpListeners ApplicationGatewayHttpListenerResponseArrayInput
	// Resource location.
	Location pulumi.StringPtrInput
	// Resource name.
	Name pulumi.StringPtrInput
	// Operational state of the application gateway resource. Possible values are: 'Stopped', 'Started', 'Running', and 'Stopping'.
	OperationalState pulumi.StringPtrInput
	// Probes of the application gateway resource.
	Probes ApplicationGatewayProbeResponseArrayInput
	// Provisioning state of the application gateway resource. Possible values are: 'Updating', 'Deleting', and 'Failed'.
	ProvisioningState pulumi.StringPtrInput
	// Request routing rules of the application gateway resource.
	RequestRoutingRules ApplicationGatewayRequestRoutingRuleResponseArrayInput
	// Resource GUID property of the application gateway resource.
	ResourceGuid pulumi.StringPtrInput
	// SKU of the application gateway resource.
	Sku ApplicationGatewaySkuResponsePtrInput
	// SSL certificates of the application gateway resource.
	SslCertificates ApplicationGatewaySslCertificateResponseArrayInput
	// Resource tags.
	Tags pulumi.StringMapInput
	// Resource type.
	Type pulumi.StringPtrInput
	// URL path map of the application gateway resource.
	UrlPathMaps ApplicationGatewayUrlPathMapResponseArrayInput
}

func (ApplicationGatewayState) ElementType() reflect.Type {
	return reflect.TypeOf((*applicationGatewayState)(nil)).Elem()
}

type applicationGatewayArgs struct {
	// The name of the application gateway.
	ApplicationGatewayName string `pulumi:"applicationGatewayName"`
	// Backend address pool of the application gateway resource.
	BackendAddressPools []ApplicationGatewayBackendAddressPool `pulumi:"backendAddressPools"`
	// Backend http settings of the application gateway resource.
	BackendHttpSettingsCollection []ApplicationGatewayBackendHttpSettings `pulumi:"backendHttpSettingsCollection"`
	// A unique read-only string that changes whenever the resource is updated.
	Etag *string `pulumi:"etag"`
	// Frontend IP addresses of the application gateway resource.
	FrontendIPConfigurations []ApplicationGatewayFrontendIPConfiguration `pulumi:"frontendIPConfigurations"`
	// Frontend ports of the application gateway resource.
	FrontendPorts []ApplicationGatewayFrontendPort `pulumi:"frontendPorts"`
	// Gets or sets subnets of application gateway resource
	GatewayIPConfigurations []ApplicationGatewayIPConfiguration `pulumi:"gatewayIPConfigurations"`
	// Http listeners of the application gateway resource.
	HttpListeners []ApplicationGatewayHttpListener `pulumi:"httpListeners"`
	// Resource Identifier.
	Id *string `pulumi:"id"`
	// Resource location.
	Location *string `pulumi:"location"`
	// Probes of the application gateway resource.
	Probes []ApplicationGatewayProbe `pulumi:"probes"`
	// Provisioning state of the application gateway resource. Possible values are: 'Updating', 'Deleting', and 'Failed'.
	ProvisioningState *string `pulumi:"provisioningState"`
	// Request routing rules of the application gateway resource.
	RequestRoutingRules []ApplicationGatewayRequestRoutingRule `pulumi:"requestRoutingRules"`
	// The name of the resource group.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// Resource GUID property of the application gateway resource.
	ResourceGuid *string `pulumi:"resourceGuid"`
	// SKU of the application gateway resource.
	Sku *ApplicationGatewaySku `pulumi:"sku"`
	// SSL certificates of the application gateway resource.
	SslCertificates []ApplicationGatewaySslCertificate `pulumi:"sslCertificates"`
	// Resource tags.
	Tags map[string]string `pulumi:"tags"`
	// URL path map of the application gateway resource.
	UrlPathMaps []ApplicationGatewayUrlPathMap `pulumi:"urlPathMaps"`
}

// The set of arguments for constructing a ApplicationGateway resource.
type ApplicationGatewayArgs struct {
	// The name of the application gateway.
	ApplicationGatewayName pulumi.StringInput
	// Backend address pool of the application gateway resource.
	BackendAddressPools ApplicationGatewayBackendAddressPoolArrayInput
	// Backend http settings of the application gateway resource.
	BackendHttpSettingsCollection ApplicationGatewayBackendHttpSettingsArrayInput
	// A unique read-only string that changes whenever the resource is updated.
	Etag pulumi.StringPtrInput
	// Frontend IP addresses of the application gateway resource.
	FrontendIPConfigurations ApplicationGatewayFrontendIPConfigurationArrayInput
	// Frontend ports of the application gateway resource.
	FrontendPorts ApplicationGatewayFrontendPortArrayInput
	// Gets or sets subnets of application gateway resource
	GatewayIPConfigurations ApplicationGatewayIPConfigurationArrayInput
	// Http listeners of the application gateway resource.
	HttpListeners ApplicationGatewayHttpListenerArrayInput
	// Resource Identifier.
	Id pulumi.StringPtrInput
	// Resource location.
	Location pulumi.StringPtrInput
	// Probes of the application gateway resource.
	Probes ApplicationGatewayProbeArrayInput
	// Provisioning state of the application gateway resource. Possible values are: 'Updating', 'Deleting', and 'Failed'.
	ProvisioningState pulumi.StringPtrInput
	// Request routing rules of the application gateway resource.
	RequestRoutingRules ApplicationGatewayRequestRoutingRuleArrayInput
	// The name of the resource group.
	ResourceGroupName pulumi.StringInput
	// Resource GUID property of the application gateway resource.
	ResourceGuid pulumi.StringPtrInput
	// SKU of the application gateway resource.
	Sku ApplicationGatewaySkuPtrInput
	// SSL certificates of the application gateway resource.
	SslCertificates ApplicationGatewaySslCertificateArrayInput
	// Resource tags.
	Tags pulumi.StringMapInput
	// URL path map of the application gateway resource.
	UrlPathMaps ApplicationGatewayUrlPathMapArrayInput
}

func (ApplicationGatewayArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*applicationGatewayArgs)(nil)).Elem()
}
