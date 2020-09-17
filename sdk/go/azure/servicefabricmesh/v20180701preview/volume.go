// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v20180701preview

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// This type describes a volume resource.
type Volume struct {
	pulumi.CustomResourceState

	// This type describes a volume provided by an Azure Files file share.
	AzureFileParameters VolumeProviderParametersAzureFileResponsePtrOutput `pulumi:"azureFileParameters"`
	// User readable description of the volume.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// The geo-location where the resource lives
	Location pulumi.StringOutput `pulumi:"location"`
	// The name of the resource
	Name pulumi.StringOutput `pulumi:"name"`
	// Provider of the volume.
	Provider pulumi.StringOutput `pulumi:"provider"`
	// State of the resource.
	ProvisioningState pulumi.StringOutput `pulumi:"provisioningState"`
	// Resource tags.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// The type of the resource. Ex- Microsoft.Compute/virtualMachines or Microsoft.Storage/storageAccounts.
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewVolume registers a new resource with the given unique name, arguments, and options.
func NewVolume(ctx *pulumi.Context,
	name string, args *VolumeArgs, opts ...pulumi.ResourceOption) (*Volume, error) {
	if args == nil || args.Location == nil {
		return nil, errors.New("missing required argument 'Location'")
	}
	if args == nil || args.Provider == nil {
		return nil, errors.New("missing required argument 'Provider'")
	}
	if args == nil || args.ResourceGroupName == nil {
		return nil, errors.New("missing required argument 'ResourceGroupName'")
	}
	if args == nil || args.VolumeName == nil {
		return nil, errors.New("missing required argument 'VolumeName'")
	}
	if args == nil {
		args = &VolumeArgs{}
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-nextgen:servicefabricmesh/v20180901preview:Volume"),
		},
	})
	opts = append(opts, aliases)
	var resource Volume
	err := ctx.RegisterResource("azure-nextgen:servicefabricmesh/v20180701preview:Volume", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetVolume gets an existing Volume resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetVolume(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *VolumeState, opts ...pulumi.ResourceOption) (*Volume, error) {
	var resource Volume
	err := ctx.ReadResource("azure-nextgen:servicefabricmesh/v20180701preview:Volume", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Volume resources.
type volumeState struct {
	// This type describes a volume provided by an Azure Files file share.
	AzureFileParameters *VolumeProviderParametersAzureFileResponse `pulumi:"azureFileParameters"`
	// User readable description of the volume.
	Description *string `pulumi:"description"`
	// The geo-location where the resource lives
	Location *string `pulumi:"location"`
	// The name of the resource
	Name *string `pulumi:"name"`
	// Provider of the volume.
	Provider *string `pulumi:"provider"`
	// State of the resource.
	ProvisioningState *string `pulumi:"provisioningState"`
	// Resource tags.
	Tags map[string]string `pulumi:"tags"`
	// The type of the resource. Ex- Microsoft.Compute/virtualMachines or Microsoft.Storage/storageAccounts.
	Type *string `pulumi:"type"`
}

type VolumeState struct {
	// This type describes a volume provided by an Azure Files file share.
	AzureFileParameters VolumeProviderParametersAzureFileResponsePtrInput
	// User readable description of the volume.
	Description pulumi.StringPtrInput
	// The geo-location where the resource lives
	Location pulumi.StringPtrInput
	// The name of the resource
	Name pulumi.StringPtrInput
	// Provider of the volume.
	Provider pulumi.StringPtrInput
	// State of the resource.
	ProvisioningState pulumi.StringPtrInput
	// Resource tags.
	Tags pulumi.StringMapInput
	// The type of the resource. Ex- Microsoft.Compute/virtualMachines or Microsoft.Storage/storageAccounts.
	Type pulumi.StringPtrInput
}

func (VolumeState) ElementType() reflect.Type {
	return reflect.TypeOf((*volumeState)(nil)).Elem()
}

type volumeArgs struct {
	// This type describes a volume provided by an Azure Files file share.
	AzureFileParameters *VolumeProviderParametersAzureFile `pulumi:"azureFileParameters"`
	// User readable description of the volume.
	Description *string `pulumi:"description"`
	// The geo-location where the resource lives
	Location string `pulumi:"location"`
	// Provider of the volume.
	Provider string `pulumi:"provider"`
	// Azure resource group name
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// Resource tags.
	Tags map[string]string `pulumi:"tags"`
	// The identity of the volume.
	VolumeName string `pulumi:"volumeName"`
}

// The set of arguments for constructing a Volume resource.
type VolumeArgs struct {
	// This type describes a volume provided by an Azure Files file share.
	AzureFileParameters VolumeProviderParametersAzureFilePtrInput
	// User readable description of the volume.
	Description pulumi.StringPtrInput
	// The geo-location where the resource lives
	Location pulumi.StringInput
	// Provider of the volume.
	Provider pulumi.StringInput
	// Azure resource group name
	ResourceGroupName pulumi.StringInput
	// Resource tags.
	Tags pulumi.StringMapInput
	// The identity of the volume.
	VolumeName pulumi.StringInput
}

func (VolumeArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*volumeArgs)(nil)).Elem()
}