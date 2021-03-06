// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v20190801

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// An Azure Cosmos DB MongoDB collection.
type MongoDBResourceMongoDBCollection struct {
	pulumi.CustomResourceState

	// The location of the resource group to which the resource belongs.
	Location pulumi.StringPtrOutput `pulumi:"location"`
	// The name of the ARM resource.
	Name     pulumi.StringOutput                                     `pulumi:"name"`
	Resource MongoDBCollectionGetPropertiesResponseResourcePtrOutput `pulumi:"resource"`
	// Tags are a list of key-value pairs that describe the resource. These tags can be used in viewing and grouping this resource (across resource groups). A maximum of 15 tags can be provided for a resource. Each tag must have a key no greater than 128 characters and value no greater than 256 characters. For example, the default experience for a template type is set with "defaultExperience": "Cassandra". Current "defaultExperience" values also include "Table", "Graph", "DocumentDB", and "MongoDB".
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// The type of Azure resource.
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewMongoDBResourceMongoDBCollection registers a new resource with the given unique name, arguments, and options.
func NewMongoDBResourceMongoDBCollection(ctx *pulumi.Context,
	name string, args *MongoDBResourceMongoDBCollectionArgs, opts ...pulumi.ResourceOption) (*MongoDBResourceMongoDBCollection, error) {
	if args == nil || args.AccountName == nil {
		return nil, errors.New("missing required argument 'AccountName'")
	}
	if args == nil || args.CollectionName == nil {
		return nil, errors.New("missing required argument 'CollectionName'")
	}
	if args == nil || args.DatabaseName == nil {
		return nil, errors.New("missing required argument 'DatabaseName'")
	}
	if args == nil || args.Options == nil {
		return nil, errors.New("missing required argument 'Options'")
	}
	if args == nil || args.Resource == nil {
		return nil, errors.New("missing required argument 'Resource'")
	}
	if args == nil || args.ResourceGroupName == nil {
		return nil, errors.New("missing required argument 'ResourceGroupName'")
	}
	if args == nil {
		args = &MongoDBResourceMongoDBCollectionArgs{}
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-nextgen:documentdb/latest:MongoDBResourceMongoDBCollection"),
		},
		{
			Type: pulumi.String("azure-nextgen:documentdb/v20191212:MongoDBResourceMongoDBCollection"),
		},
		{
			Type: pulumi.String("azure-nextgen:documentdb/v20200301:MongoDBResourceMongoDBCollection"),
		},
		{
			Type: pulumi.String("azure-nextgen:documentdb/v20200401:MongoDBResourceMongoDBCollection"),
		},
		{
			Type: pulumi.String("azure-nextgen:documentdb/v20200601preview:MongoDBResourceMongoDBCollection"),
		},
		{
			Type: pulumi.String("azure-nextgen:documentdb/v20200901:MongoDBResourceMongoDBCollection"),
		},
	})
	opts = append(opts, aliases)
	var resource MongoDBResourceMongoDBCollection
	err := ctx.RegisterResource("azure-nextgen:documentdb/v20190801:MongoDBResourceMongoDBCollection", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetMongoDBResourceMongoDBCollection gets an existing MongoDBResourceMongoDBCollection resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetMongoDBResourceMongoDBCollection(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *MongoDBResourceMongoDBCollectionState, opts ...pulumi.ResourceOption) (*MongoDBResourceMongoDBCollection, error) {
	var resource MongoDBResourceMongoDBCollection
	err := ctx.ReadResource("azure-nextgen:documentdb/v20190801:MongoDBResourceMongoDBCollection", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering MongoDBResourceMongoDBCollection resources.
type mongoDBResourceMongoDBCollectionState struct {
	// The location of the resource group to which the resource belongs.
	Location *string `pulumi:"location"`
	// The name of the ARM resource.
	Name     *string                                         `pulumi:"name"`
	Resource *MongoDBCollectionGetPropertiesResponseResource `pulumi:"resource"`
	// Tags are a list of key-value pairs that describe the resource. These tags can be used in viewing and grouping this resource (across resource groups). A maximum of 15 tags can be provided for a resource. Each tag must have a key no greater than 128 characters and value no greater than 256 characters. For example, the default experience for a template type is set with "defaultExperience": "Cassandra". Current "defaultExperience" values also include "Table", "Graph", "DocumentDB", and "MongoDB".
	Tags map[string]string `pulumi:"tags"`
	// The type of Azure resource.
	Type *string `pulumi:"type"`
}

type MongoDBResourceMongoDBCollectionState struct {
	// The location of the resource group to which the resource belongs.
	Location pulumi.StringPtrInput
	// The name of the ARM resource.
	Name     pulumi.StringPtrInput
	Resource MongoDBCollectionGetPropertiesResponseResourcePtrInput
	// Tags are a list of key-value pairs that describe the resource. These tags can be used in viewing and grouping this resource (across resource groups). A maximum of 15 tags can be provided for a resource. Each tag must have a key no greater than 128 characters and value no greater than 256 characters. For example, the default experience for a template type is set with "defaultExperience": "Cassandra". Current "defaultExperience" values also include "Table", "Graph", "DocumentDB", and "MongoDB".
	Tags pulumi.StringMapInput
	// The type of Azure resource.
	Type pulumi.StringPtrInput
}

func (MongoDBResourceMongoDBCollectionState) ElementType() reflect.Type {
	return reflect.TypeOf((*mongoDBResourceMongoDBCollectionState)(nil)).Elem()
}

type mongoDBResourceMongoDBCollectionArgs struct {
	// Cosmos DB database account name.
	AccountName string `pulumi:"accountName"`
	// Cosmos DB collection name.
	CollectionName string `pulumi:"collectionName"`
	// Cosmos DB database name.
	DatabaseName string `pulumi:"databaseName"`
	// The location of the resource group to which the resource belongs.
	Location *string `pulumi:"location"`
	// A key-value pair of options to be applied for the request. This corresponds to the headers sent with the request.
	Options map[string]string `pulumi:"options"`
	// The standard JSON format of a MongoDB collection
	Resource MongoDBCollectionResource `pulumi:"resource"`
	// Name of an Azure resource group.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// Tags are a list of key-value pairs that describe the resource. These tags can be used in viewing and grouping this resource (across resource groups). A maximum of 15 tags can be provided for a resource. Each tag must have a key no greater than 128 characters and value no greater than 256 characters. For example, the default experience for a template type is set with "defaultExperience": "Cassandra". Current "defaultExperience" values also include "Table", "Graph", "DocumentDB", and "MongoDB".
	Tags map[string]string `pulumi:"tags"`
}

// The set of arguments for constructing a MongoDBResourceMongoDBCollection resource.
type MongoDBResourceMongoDBCollectionArgs struct {
	// Cosmos DB database account name.
	AccountName pulumi.StringInput
	// Cosmos DB collection name.
	CollectionName pulumi.StringInput
	// Cosmos DB database name.
	DatabaseName pulumi.StringInput
	// The location of the resource group to which the resource belongs.
	Location pulumi.StringPtrInput
	// A key-value pair of options to be applied for the request. This corresponds to the headers sent with the request.
	Options pulumi.StringMapInput
	// The standard JSON format of a MongoDB collection
	Resource MongoDBCollectionResourceInput
	// Name of an Azure resource group.
	ResourceGroupName pulumi.StringInput
	// Tags are a list of key-value pairs that describe the resource. These tags can be used in viewing and grouping this resource (across resource groups). A maximum of 15 tags can be provided for a resource. Each tag must have a key no greater than 128 characters and value no greater than 256 characters. For example, the default experience for a template type is set with "defaultExperience": "Cassandra". Current "defaultExperience" values also include "Table", "Graph", "DocumentDB", and "MongoDB".
	Tags pulumi.StringMapInput
}

func (MongoDBResourceMongoDBCollectionArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*mongoDBResourceMongoDBCollectionArgs)(nil)).Elem()
}
