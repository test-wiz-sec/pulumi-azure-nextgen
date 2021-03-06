// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v20150320

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// The top level storage insight resource container.
type StorageInsight struct {
	pulumi.CustomResourceState

	// The names of the blob containers that the workspace should read
	Containers pulumi.StringArrayOutput `pulumi:"containers"`
	// The ETag of the storage insight.
	ETag pulumi.StringPtrOutput `pulumi:"eTag"`
	// Resource name.
	Name pulumi.StringOutput `pulumi:"name"`
	// The status of the storage insight
	Status StorageInsightStatusResponseOutput `pulumi:"status"`
	// The storage account connection details
	StorageAccount StorageAccountResponseOutput `pulumi:"storageAccount"`
	// The names of the Azure tables that the workspace should read
	Tables pulumi.StringArrayOutput `pulumi:"tables"`
	// Resource tags
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// Resource type.
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewStorageInsight registers a new resource with the given unique name, arguments, and options.
func NewStorageInsight(ctx *pulumi.Context,
	name string, args *StorageInsightArgs, opts ...pulumi.ResourceOption) (*StorageInsight, error) {
	if args == nil || args.ResourceGroupName == nil {
		return nil, errors.New("missing required argument 'ResourceGroupName'")
	}
	if args == nil || args.StorageAccount == nil {
		return nil, errors.New("missing required argument 'StorageAccount'")
	}
	if args == nil || args.StorageInsightName == nil {
		return nil, errors.New("missing required argument 'StorageInsightName'")
	}
	if args == nil || args.WorkspaceName == nil {
		return nil, errors.New("missing required argument 'WorkspaceName'")
	}
	if args == nil {
		args = &StorageInsightArgs{}
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-nextgen:operationalinsights/latest:StorageInsight"),
		},
		{
			Type: pulumi.String("azure-nextgen:operationalinsights/v20200301preview:StorageInsight"),
		},
		{
			Type: pulumi.String("azure-nextgen:operationalinsights/v20200801:StorageInsight"),
		},
	})
	opts = append(opts, aliases)
	var resource StorageInsight
	err := ctx.RegisterResource("azure-nextgen:operationalinsights/v20150320:StorageInsight", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetStorageInsight gets an existing StorageInsight resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetStorageInsight(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *StorageInsightState, opts ...pulumi.ResourceOption) (*StorageInsight, error) {
	var resource StorageInsight
	err := ctx.ReadResource("azure-nextgen:operationalinsights/v20150320:StorageInsight", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering StorageInsight resources.
type storageInsightState struct {
	// The names of the blob containers that the workspace should read
	Containers []string `pulumi:"containers"`
	// The ETag of the storage insight.
	ETag *string `pulumi:"eTag"`
	// Resource name.
	Name *string `pulumi:"name"`
	// The status of the storage insight
	Status *StorageInsightStatusResponse `pulumi:"status"`
	// The storage account connection details
	StorageAccount *StorageAccountResponse `pulumi:"storageAccount"`
	// The names of the Azure tables that the workspace should read
	Tables []string `pulumi:"tables"`
	// Resource tags
	Tags map[string]string `pulumi:"tags"`
	// Resource type.
	Type *string `pulumi:"type"`
}

type StorageInsightState struct {
	// The names of the blob containers that the workspace should read
	Containers pulumi.StringArrayInput
	// The ETag of the storage insight.
	ETag pulumi.StringPtrInput
	// Resource name.
	Name pulumi.StringPtrInput
	// The status of the storage insight
	Status StorageInsightStatusResponsePtrInput
	// The storage account connection details
	StorageAccount StorageAccountResponsePtrInput
	// The names of the Azure tables that the workspace should read
	Tables pulumi.StringArrayInput
	// Resource tags
	Tags pulumi.StringMapInput
	// Resource type.
	Type pulumi.StringPtrInput
}

func (StorageInsightState) ElementType() reflect.Type {
	return reflect.TypeOf((*storageInsightState)(nil)).Elem()
}

type storageInsightArgs struct {
	// The names of the blob containers that the workspace should read
	Containers []string `pulumi:"containers"`
	// The ETag of the storage insight.
	ETag *string `pulumi:"eTag"`
	// The Resource Group name.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The storage account connection details
	StorageAccount StorageAccount `pulumi:"storageAccount"`
	// Name of the storageInsightsConfigs resource
	StorageInsightName string `pulumi:"storageInsightName"`
	// The names of the Azure tables that the workspace should read
	Tables []string `pulumi:"tables"`
	// Resource tags
	Tags map[string]string `pulumi:"tags"`
	// The Log Analytics Workspace name.
	WorkspaceName string `pulumi:"workspaceName"`
}

// The set of arguments for constructing a StorageInsight resource.
type StorageInsightArgs struct {
	// The names of the blob containers that the workspace should read
	Containers pulumi.StringArrayInput
	// The ETag of the storage insight.
	ETag pulumi.StringPtrInput
	// The Resource Group name.
	ResourceGroupName pulumi.StringInput
	// The storage account connection details
	StorageAccount StorageAccountInput
	// Name of the storageInsightsConfigs resource
	StorageInsightName pulumi.StringInput
	// The names of the Azure tables that the workspace should read
	Tables pulumi.StringArrayInput
	// Resource tags
	Tags pulumi.StringMapInput
	// The Log Analytics Workspace name.
	WorkspaceName pulumi.StringInput
}

func (StorageInsightArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*storageInsightArgs)(nil)).Elem()
}
