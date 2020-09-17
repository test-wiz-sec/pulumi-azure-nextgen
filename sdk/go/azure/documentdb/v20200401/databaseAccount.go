// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v20200401

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// An Azure Cosmos DB database account.
type DatabaseAccount struct {
	pulumi.CustomResourceState

	// API specific properties.
	ApiProperties ApiPropertiesResponsePtrOutput `pulumi:"apiProperties"`
	// List of Cosmos DB capabilities for the account
	Capabilities CapabilityResponseArrayOutput `pulumi:"capabilities"`
	// The cassandra connector offer type for the Cosmos DB database C* account.
	ConnectorOffer pulumi.StringPtrOutput `pulumi:"connectorOffer"`
	// The consistency policy for the Cosmos DB database account.
	ConsistencyPolicy ConsistencyPolicyResponsePtrOutput `pulumi:"consistencyPolicy"`
	// The CORS policy for the Cosmos DB database account.
	Cors CorsPolicyResponseArrayOutput `pulumi:"cors"`
	// The offer type for the Cosmos DB database account. Default value: Standard.
	DatabaseAccountOfferType pulumi.StringOutput `pulumi:"databaseAccountOfferType"`
	// Disable write operations on metadata resources (databases, containers, throughput) via account keys
	DisableKeyBasedMetadataWriteAccess pulumi.BoolPtrOutput `pulumi:"disableKeyBasedMetadataWriteAccess"`
	// The connection endpoint for the Cosmos DB database account.
	DocumentEndpoint pulumi.StringOutput `pulumi:"documentEndpoint"`
	// Flag to indicate whether to enable storage analytics.
	EnableAnalyticalStorage pulumi.BoolPtrOutput `pulumi:"enableAnalyticalStorage"`
	// Enables automatic failover of the write region in the rare event that the region is unavailable due to an outage. Automatic failover will result in a new write region for the account and is chosen based on the failover priorities configured for the account.
	EnableAutomaticFailover pulumi.BoolPtrOutput `pulumi:"enableAutomaticFailover"`
	// Enables the cassandra connector on the Cosmos DB C* account
	EnableCassandraConnector pulumi.BoolPtrOutput `pulumi:"enableCassandraConnector"`
	// Flag to indicate whether Free Tier is enabled.
	EnableFreeTier pulumi.BoolPtrOutput `pulumi:"enableFreeTier"`
	// Enables the account to write in multiple locations
	EnableMultipleWriteLocations pulumi.BoolPtrOutput `pulumi:"enableMultipleWriteLocations"`
	// An array that contains the regions ordered by their failover priorities.
	FailoverPolicies FailoverPolicyResponseArrayOutput `pulumi:"failoverPolicies"`
	// List of IpRules.
	IpRules IpAddressOrRangeResponseArrayOutput `pulumi:"ipRules"`
	// Flag to indicate whether to enable/disable Virtual Network ACL rules.
	IsVirtualNetworkFilterEnabled pulumi.BoolPtrOutput `pulumi:"isVirtualNetworkFilterEnabled"`
	// The URI of the key vault
	KeyVaultKeyUri pulumi.StringPtrOutput `pulumi:"keyVaultKeyUri"`
	// Indicates the type of database account. This can only be set at database account creation.
	Kind pulumi.StringPtrOutput `pulumi:"kind"`
	// The location of the resource group to which the resource belongs.
	Location pulumi.StringPtrOutput `pulumi:"location"`
	// An array that contains all of the locations enabled for the Cosmos DB account.
	Locations LocationResponseArrayOutput `pulumi:"locations"`
	// The name of the ARM resource.
	Name pulumi.StringOutput `pulumi:"name"`
	// List of Private Endpoint Connections configured for the Cosmos DB account.
	PrivateEndpointConnections PrivateEndpointConnectionResponseArrayOutput `pulumi:"privateEndpointConnections"`
	// The status of the Cosmos DB account at the time the operation was called. The status can be one of following. 'Creating' – the Cosmos DB account is being created. When an account is in Creating state, only properties that are specified as input for the Create Cosmos DB account operation are returned. 'Succeeded' – the Cosmos DB account is active for use. 'Updating' – the Cosmos DB account is being updated. 'Deleting' – the Cosmos DB account is being deleted. 'Failed' – the Cosmos DB account failed creation. 'DeletionFailed' – the Cosmos DB account deletion failed.
	ProvisioningState pulumi.StringPtrOutput `pulumi:"provisioningState"`
	// Whether requests from Public Network are allowed
	PublicNetworkAccess pulumi.StringPtrOutput `pulumi:"publicNetworkAccess"`
	// An array that contains of the read locations enabled for the Cosmos DB account.
	ReadLocations LocationResponseArrayOutput `pulumi:"readLocations"`
	// Tags are a list of key-value pairs that describe the resource. These tags can be used in viewing and grouping this resource (across resource groups). A maximum of 15 tags can be provided for a resource. Each tag must have a key no greater than 128 characters and value no greater than 256 characters. For example, the default experience for a template type is set with "defaultExperience": "Cassandra". Current "defaultExperience" values also include "Table", "Graph", "DocumentDB", and "MongoDB".
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// The type of Azure resource.
	Type pulumi.StringOutput `pulumi:"type"`
	// List of Virtual Network ACL rules configured for the Cosmos DB account.
	VirtualNetworkRules VirtualNetworkRuleResponseArrayOutput `pulumi:"virtualNetworkRules"`
	// An array that contains the write location for the Cosmos DB account.
	WriteLocations LocationResponseArrayOutput `pulumi:"writeLocations"`
}

// NewDatabaseAccount registers a new resource with the given unique name, arguments, and options.
func NewDatabaseAccount(ctx *pulumi.Context,
	name string, args *DatabaseAccountArgs, opts ...pulumi.ResourceOption) (*DatabaseAccount, error) {
	if args == nil || args.AccountName == nil {
		return nil, errors.New("missing required argument 'AccountName'")
	}
	if args == nil || args.DatabaseAccountOfferType == nil {
		return nil, errors.New("missing required argument 'DatabaseAccountOfferType'")
	}
	if args == nil || args.Locations == nil {
		return nil, errors.New("missing required argument 'Locations'")
	}
	if args == nil || args.ResourceGroupName == nil {
		return nil, errors.New("missing required argument 'ResourceGroupName'")
	}
	if args == nil {
		args = &DatabaseAccountArgs{}
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-nextgen:documentdb/latest:DatabaseAccount"),
		},
		{
			Type: pulumi.String("azure-nextgen:documentdb/v20150401:DatabaseAccount"),
		},
		{
			Type: pulumi.String("azure-nextgen:documentdb/v20150408:DatabaseAccount"),
		},
		{
			Type: pulumi.String("azure-nextgen:documentdb/v20151106:DatabaseAccount"),
		},
		{
			Type: pulumi.String("azure-nextgen:documentdb/v20160319:DatabaseAccount"),
		},
		{
			Type: pulumi.String("azure-nextgen:documentdb/v20160331:DatabaseAccount"),
		},
		{
			Type: pulumi.String("azure-nextgen:documentdb/v20190801:DatabaseAccount"),
		},
		{
			Type: pulumi.String("azure-nextgen:documentdb/v20191212:DatabaseAccount"),
		},
		{
			Type: pulumi.String("azure-nextgen:documentdb/v20200301:DatabaseAccount"),
		},
		{
			Type: pulumi.String("azure-nextgen:documentdb/v20200601preview:DatabaseAccount"),
		},
	})
	opts = append(opts, aliases)
	var resource DatabaseAccount
	err := ctx.RegisterResource("azure-nextgen:documentdb/v20200401:DatabaseAccount", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetDatabaseAccount gets an existing DatabaseAccount resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetDatabaseAccount(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *DatabaseAccountState, opts ...pulumi.ResourceOption) (*DatabaseAccount, error) {
	var resource DatabaseAccount
	err := ctx.ReadResource("azure-nextgen:documentdb/v20200401:DatabaseAccount", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering DatabaseAccount resources.
type databaseAccountState struct {
	// API specific properties.
	ApiProperties *ApiPropertiesResponse `pulumi:"apiProperties"`
	// List of Cosmos DB capabilities for the account
	Capabilities []CapabilityResponse `pulumi:"capabilities"`
	// The cassandra connector offer type for the Cosmos DB database C* account.
	ConnectorOffer *string `pulumi:"connectorOffer"`
	// The consistency policy for the Cosmos DB database account.
	ConsistencyPolicy *ConsistencyPolicyResponse `pulumi:"consistencyPolicy"`
	// The CORS policy for the Cosmos DB database account.
	Cors []CorsPolicyResponse `pulumi:"cors"`
	// The offer type for the Cosmos DB database account. Default value: Standard.
	DatabaseAccountOfferType *string `pulumi:"databaseAccountOfferType"`
	// Disable write operations on metadata resources (databases, containers, throughput) via account keys
	DisableKeyBasedMetadataWriteAccess *bool `pulumi:"disableKeyBasedMetadataWriteAccess"`
	// The connection endpoint for the Cosmos DB database account.
	DocumentEndpoint *string `pulumi:"documentEndpoint"`
	// Flag to indicate whether to enable storage analytics.
	EnableAnalyticalStorage *bool `pulumi:"enableAnalyticalStorage"`
	// Enables automatic failover of the write region in the rare event that the region is unavailable due to an outage. Automatic failover will result in a new write region for the account and is chosen based on the failover priorities configured for the account.
	EnableAutomaticFailover *bool `pulumi:"enableAutomaticFailover"`
	// Enables the cassandra connector on the Cosmos DB C* account
	EnableCassandraConnector *bool `pulumi:"enableCassandraConnector"`
	// Flag to indicate whether Free Tier is enabled.
	EnableFreeTier *bool `pulumi:"enableFreeTier"`
	// Enables the account to write in multiple locations
	EnableMultipleWriteLocations *bool `pulumi:"enableMultipleWriteLocations"`
	// An array that contains the regions ordered by their failover priorities.
	FailoverPolicies []FailoverPolicyResponse `pulumi:"failoverPolicies"`
	// List of IpRules.
	IpRules []IpAddressOrRangeResponse `pulumi:"ipRules"`
	// Flag to indicate whether to enable/disable Virtual Network ACL rules.
	IsVirtualNetworkFilterEnabled *bool `pulumi:"isVirtualNetworkFilterEnabled"`
	// The URI of the key vault
	KeyVaultKeyUri *string `pulumi:"keyVaultKeyUri"`
	// Indicates the type of database account. This can only be set at database account creation.
	Kind *string `pulumi:"kind"`
	// The location of the resource group to which the resource belongs.
	Location *string `pulumi:"location"`
	// An array that contains all of the locations enabled for the Cosmos DB account.
	Locations []LocationResponse `pulumi:"locations"`
	// The name of the ARM resource.
	Name *string `pulumi:"name"`
	// List of Private Endpoint Connections configured for the Cosmos DB account.
	PrivateEndpointConnections []PrivateEndpointConnectionResponse `pulumi:"privateEndpointConnections"`
	// The status of the Cosmos DB account at the time the operation was called. The status can be one of following. 'Creating' – the Cosmos DB account is being created. When an account is in Creating state, only properties that are specified as input for the Create Cosmos DB account operation are returned. 'Succeeded' – the Cosmos DB account is active for use. 'Updating' – the Cosmos DB account is being updated. 'Deleting' – the Cosmos DB account is being deleted. 'Failed' – the Cosmos DB account failed creation. 'DeletionFailed' – the Cosmos DB account deletion failed.
	ProvisioningState *string `pulumi:"provisioningState"`
	// Whether requests from Public Network are allowed
	PublicNetworkAccess *string `pulumi:"publicNetworkAccess"`
	// An array that contains of the read locations enabled for the Cosmos DB account.
	ReadLocations []LocationResponse `pulumi:"readLocations"`
	// Tags are a list of key-value pairs that describe the resource. These tags can be used in viewing and grouping this resource (across resource groups). A maximum of 15 tags can be provided for a resource. Each tag must have a key no greater than 128 characters and value no greater than 256 characters. For example, the default experience for a template type is set with "defaultExperience": "Cassandra". Current "defaultExperience" values also include "Table", "Graph", "DocumentDB", and "MongoDB".
	Tags map[string]string `pulumi:"tags"`
	// The type of Azure resource.
	Type *string `pulumi:"type"`
	// List of Virtual Network ACL rules configured for the Cosmos DB account.
	VirtualNetworkRules []VirtualNetworkRuleResponse `pulumi:"virtualNetworkRules"`
	// An array that contains the write location for the Cosmos DB account.
	WriteLocations []LocationResponse `pulumi:"writeLocations"`
}

type DatabaseAccountState struct {
	// API specific properties.
	ApiProperties ApiPropertiesResponsePtrInput
	// List of Cosmos DB capabilities for the account
	Capabilities CapabilityResponseArrayInput
	// The cassandra connector offer type for the Cosmos DB database C* account.
	ConnectorOffer pulumi.StringPtrInput
	// The consistency policy for the Cosmos DB database account.
	ConsistencyPolicy ConsistencyPolicyResponsePtrInput
	// The CORS policy for the Cosmos DB database account.
	Cors CorsPolicyResponseArrayInput
	// The offer type for the Cosmos DB database account. Default value: Standard.
	DatabaseAccountOfferType pulumi.StringPtrInput
	// Disable write operations on metadata resources (databases, containers, throughput) via account keys
	DisableKeyBasedMetadataWriteAccess pulumi.BoolPtrInput
	// The connection endpoint for the Cosmos DB database account.
	DocumentEndpoint pulumi.StringPtrInput
	// Flag to indicate whether to enable storage analytics.
	EnableAnalyticalStorage pulumi.BoolPtrInput
	// Enables automatic failover of the write region in the rare event that the region is unavailable due to an outage. Automatic failover will result in a new write region for the account and is chosen based on the failover priorities configured for the account.
	EnableAutomaticFailover pulumi.BoolPtrInput
	// Enables the cassandra connector on the Cosmos DB C* account
	EnableCassandraConnector pulumi.BoolPtrInput
	// Flag to indicate whether Free Tier is enabled.
	EnableFreeTier pulumi.BoolPtrInput
	// Enables the account to write in multiple locations
	EnableMultipleWriteLocations pulumi.BoolPtrInput
	// An array that contains the regions ordered by their failover priorities.
	FailoverPolicies FailoverPolicyResponseArrayInput
	// List of IpRules.
	IpRules IpAddressOrRangeResponseArrayInput
	// Flag to indicate whether to enable/disable Virtual Network ACL rules.
	IsVirtualNetworkFilterEnabled pulumi.BoolPtrInput
	// The URI of the key vault
	KeyVaultKeyUri pulumi.StringPtrInput
	// Indicates the type of database account. This can only be set at database account creation.
	Kind pulumi.StringPtrInput
	// The location of the resource group to which the resource belongs.
	Location pulumi.StringPtrInput
	// An array that contains all of the locations enabled for the Cosmos DB account.
	Locations LocationResponseArrayInput
	// The name of the ARM resource.
	Name pulumi.StringPtrInput
	// List of Private Endpoint Connections configured for the Cosmos DB account.
	PrivateEndpointConnections PrivateEndpointConnectionResponseArrayInput
	// The status of the Cosmos DB account at the time the operation was called. The status can be one of following. 'Creating' – the Cosmos DB account is being created. When an account is in Creating state, only properties that are specified as input for the Create Cosmos DB account operation are returned. 'Succeeded' – the Cosmos DB account is active for use. 'Updating' – the Cosmos DB account is being updated. 'Deleting' – the Cosmos DB account is being deleted. 'Failed' – the Cosmos DB account failed creation. 'DeletionFailed' – the Cosmos DB account deletion failed.
	ProvisioningState pulumi.StringPtrInput
	// Whether requests from Public Network are allowed
	PublicNetworkAccess pulumi.StringPtrInput
	// An array that contains of the read locations enabled for the Cosmos DB account.
	ReadLocations LocationResponseArrayInput
	// Tags are a list of key-value pairs that describe the resource. These tags can be used in viewing and grouping this resource (across resource groups). A maximum of 15 tags can be provided for a resource. Each tag must have a key no greater than 128 characters and value no greater than 256 characters. For example, the default experience for a template type is set with "defaultExperience": "Cassandra". Current "defaultExperience" values also include "Table", "Graph", "DocumentDB", and "MongoDB".
	Tags pulumi.StringMapInput
	// The type of Azure resource.
	Type pulumi.StringPtrInput
	// List of Virtual Network ACL rules configured for the Cosmos DB account.
	VirtualNetworkRules VirtualNetworkRuleResponseArrayInput
	// An array that contains the write location for the Cosmos DB account.
	WriteLocations LocationResponseArrayInput
}

func (DatabaseAccountState) ElementType() reflect.Type {
	return reflect.TypeOf((*databaseAccountState)(nil)).Elem()
}

type databaseAccountArgs struct {
	// Cosmos DB database account name.
	AccountName string `pulumi:"accountName"`
	// API specific properties. Currently, supported only for MongoDB API.
	ApiProperties *ApiProperties `pulumi:"apiProperties"`
	// List of Cosmos DB capabilities for the account
	Capabilities []Capability `pulumi:"capabilities"`
	// The cassandra connector offer type for the Cosmos DB database C* account.
	ConnectorOffer *string `pulumi:"connectorOffer"`
	// The consistency policy for the Cosmos DB account.
	ConsistencyPolicy *ConsistencyPolicy `pulumi:"consistencyPolicy"`
	// The CORS policy for the Cosmos DB database account.
	Cors []CorsPolicy `pulumi:"cors"`
	// The offer type for the database
	DatabaseAccountOfferType string `pulumi:"databaseAccountOfferType"`
	// Disable write operations on metadata resources (databases, containers, throughput) via account keys
	DisableKeyBasedMetadataWriteAccess *bool `pulumi:"disableKeyBasedMetadataWriteAccess"`
	// Flag to indicate whether to enable storage analytics.
	EnableAnalyticalStorage *bool `pulumi:"enableAnalyticalStorage"`
	// Enables automatic failover of the write region in the rare event that the region is unavailable due to an outage. Automatic failover will result in a new write region for the account and is chosen based on the failover priorities configured for the account.
	EnableAutomaticFailover *bool `pulumi:"enableAutomaticFailover"`
	// Enables the cassandra connector on the Cosmos DB C* account
	EnableCassandraConnector *bool `pulumi:"enableCassandraConnector"`
	// Flag to indicate whether Free Tier is enabled.
	EnableFreeTier *bool `pulumi:"enableFreeTier"`
	// Enables the account to write in multiple locations
	EnableMultipleWriteLocations *bool `pulumi:"enableMultipleWriteLocations"`
	// List of IpRules.
	IpRules []IpAddressOrRange `pulumi:"ipRules"`
	// Flag to indicate whether to enable/disable Virtual Network ACL rules.
	IsVirtualNetworkFilterEnabled *bool `pulumi:"isVirtualNetworkFilterEnabled"`
	// The URI of the key vault
	KeyVaultKeyUri *string `pulumi:"keyVaultKeyUri"`
	// Indicates the type of database account. This can only be set at database account creation.
	Kind *string `pulumi:"kind"`
	// The location of the resource group to which the resource belongs.
	Location *string `pulumi:"location"`
	// An array that contains the georeplication locations enabled for the Cosmos DB account.
	Locations []Location `pulumi:"locations"`
	// Whether requests from Public Network are allowed
	PublicNetworkAccess *string `pulumi:"publicNetworkAccess"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// Tags are a list of key-value pairs that describe the resource. These tags can be used in viewing and grouping this resource (across resource groups). A maximum of 15 tags can be provided for a resource. Each tag must have a key no greater than 128 characters and value no greater than 256 characters. For example, the default experience for a template type is set with "defaultExperience": "Cassandra". Current "defaultExperience" values also include "Table", "Graph", "DocumentDB", and "MongoDB".
	Tags map[string]string `pulumi:"tags"`
	// List of Virtual Network ACL rules configured for the Cosmos DB account.
	VirtualNetworkRules []VirtualNetworkRule `pulumi:"virtualNetworkRules"`
}

// The set of arguments for constructing a DatabaseAccount resource.
type DatabaseAccountArgs struct {
	// Cosmos DB database account name.
	AccountName pulumi.StringInput
	// API specific properties. Currently, supported only for MongoDB API.
	ApiProperties ApiPropertiesPtrInput
	// List of Cosmos DB capabilities for the account
	Capabilities CapabilityArrayInput
	// The cassandra connector offer type for the Cosmos DB database C* account.
	ConnectorOffer pulumi.StringPtrInput
	// The consistency policy for the Cosmos DB account.
	ConsistencyPolicy ConsistencyPolicyPtrInput
	// The CORS policy for the Cosmos DB database account.
	Cors CorsPolicyArrayInput
	// The offer type for the database
	DatabaseAccountOfferType pulumi.StringInput
	// Disable write operations on metadata resources (databases, containers, throughput) via account keys
	DisableKeyBasedMetadataWriteAccess pulumi.BoolPtrInput
	// Flag to indicate whether to enable storage analytics.
	EnableAnalyticalStorage pulumi.BoolPtrInput
	// Enables automatic failover of the write region in the rare event that the region is unavailable due to an outage. Automatic failover will result in a new write region for the account and is chosen based on the failover priorities configured for the account.
	EnableAutomaticFailover pulumi.BoolPtrInput
	// Enables the cassandra connector on the Cosmos DB C* account
	EnableCassandraConnector pulumi.BoolPtrInput
	// Flag to indicate whether Free Tier is enabled.
	EnableFreeTier pulumi.BoolPtrInput
	// Enables the account to write in multiple locations
	EnableMultipleWriteLocations pulumi.BoolPtrInput
	// List of IpRules.
	IpRules IpAddressOrRangeArrayInput
	// Flag to indicate whether to enable/disable Virtual Network ACL rules.
	IsVirtualNetworkFilterEnabled pulumi.BoolPtrInput
	// The URI of the key vault
	KeyVaultKeyUri pulumi.StringPtrInput
	// Indicates the type of database account. This can only be set at database account creation.
	Kind pulumi.StringPtrInput
	// The location of the resource group to which the resource belongs.
	Location pulumi.StringPtrInput
	// An array that contains the georeplication locations enabled for the Cosmos DB account.
	Locations LocationArrayInput
	// Whether requests from Public Network are allowed
	PublicNetworkAccess pulumi.StringPtrInput
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput
	// Tags are a list of key-value pairs that describe the resource. These tags can be used in viewing and grouping this resource (across resource groups). A maximum of 15 tags can be provided for a resource. Each tag must have a key no greater than 128 characters and value no greater than 256 characters. For example, the default experience for a template type is set with "defaultExperience": "Cassandra". Current "defaultExperience" values also include "Table", "Graph", "DocumentDB", and "MongoDB".
	Tags pulumi.StringMapInput
	// List of Virtual Network ACL rules configured for the Cosmos DB account.
	VirtualNetworkRules VirtualNetworkRuleArrayInput
}

func (DatabaseAccountArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*databaseAccountArgs)(nil)).Elem()
}