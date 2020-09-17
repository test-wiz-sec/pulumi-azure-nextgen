// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v20181015

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Represents settings of an environment, from which environment instances would be created
type EnvironmentSetting struct {
	pulumi.CustomResourceState

	// Describes the user's progress in configuring their environment setting
	ConfigurationState pulumi.StringPtrOutput `pulumi:"configurationState"`
	// Describes the environment and its resource settings
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// Time when the template VM was last changed.
	LastChanged pulumi.StringOutput `pulumi:"lastChanged"`
	// Time when the template VM was last sent for publishing.
	LastPublished pulumi.StringOutput `pulumi:"lastPublished"`
	// The details of the latest operation. ex: status, error
	LatestOperationResult LatestOperationResultResponseOutput `pulumi:"latestOperationResult"`
	// The location of the resource.
	Location pulumi.StringPtrOutput `pulumi:"location"`
	// The name of the resource.
	Name pulumi.StringOutput `pulumi:"name"`
	// The provisioning status of the resource.
	ProvisioningState pulumi.StringPtrOutput `pulumi:"provisioningState"`
	// Describes the readiness of this environment setting
	PublishingState pulumi.StringOutput `pulumi:"publishingState"`
	// The resource specific settings
	ResourceSettings ResourceSettingsResponseOutput `pulumi:"resourceSettings"`
	// The tags of the resource.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// Brief title describing the environment and its resource settings
	Title pulumi.StringPtrOutput `pulumi:"title"`
	// The type of the resource.
	Type pulumi.StringOutput `pulumi:"type"`
	// The unique immutable identifier of a resource (Guid).
	UniqueIdentifier pulumi.StringPtrOutput `pulumi:"uniqueIdentifier"`
}

// NewEnvironmentSetting registers a new resource with the given unique name, arguments, and options.
func NewEnvironmentSetting(ctx *pulumi.Context,
	name string, args *EnvironmentSettingArgs, opts ...pulumi.ResourceOption) (*EnvironmentSetting, error) {
	if args == nil || args.EnvironmentSettingName == nil {
		return nil, errors.New("missing required argument 'EnvironmentSettingName'")
	}
	if args == nil || args.LabAccountName == nil {
		return nil, errors.New("missing required argument 'LabAccountName'")
	}
	if args == nil || args.LabName == nil {
		return nil, errors.New("missing required argument 'LabName'")
	}
	if args == nil || args.ResourceGroupName == nil {
		return nil, errors.New("missing required argument 'ResourceGroupName'")
	}
	if args == nil || args.ResourceSettings == nil {
		return nil, errors.New("missing required argument 'ResourceSettings'")
	}
	if args == nil {
		args = &EnvironmentSettingArgs{}
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-nextgen:labservices/latest:EnvironmentSetting"),
		},
	})
	opts = append(opts, aliases)
	var resource EnvironmentSetting
	err := ctx.RegisterResource("azure-nextgen:labservices/v20181015:EnvironmentSetting", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetEnvironmentSetting gets an existing EnvironmentSetting resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetEnvironmentSetting(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *EnvironmentSettingState, opts ...pulumi.ResourceOption) (*EnvironmentSetting, error) {
	var resource EnvironmentSetting
	err := ctx.ReadResource("azure-nextgen:labservices/v20181015:EnvironmentSetting", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering EnvironmentSetting resources.
type environmentSettingState struct {
	// Describes the user's progress in configuring their environment setting
	ConfigurationState *string `pulumi:"configurationState"`
	// Describes the environment and its resource settings
	Description *string `pulumi:"description"`
	// Time when the template VM was last changed.
	LastChanged *string `pulumi:"lastChanged"`
	// Time when the template VM was last sent for publishing.
	LastPublished *string `pulumi:"lastPublished"`
	// The details of the latest operation. ex: status, error
	LatestOperationResult *LatestOperationResultResponse `pulumi:"latestOperationResult"`
	// The location of the resource.
	Location *string `pulumi:"location"`
	// The name of the resource.
	Name *string `pulumi:"name"`
	// The provisioning status of the resource.
	ProvisioningState *string `pulumi:"provisioningState"`
	// Describes the readiness of this environment setting
	PublishingState *string `pulumi:"publishingState"`
	// The resource specific settings
	ResourceSettings *ResourceSettingsResponse `pulumi:"resourceSettings"`
	// The tags of the resource.
	Tags map[string]string `pulumi:"tags"`
	// Brief title describing the environment and its resource settings
	Title *string `pulumi:"title"`
	// The type of the resource.
	Type *string `pulumi:"type"`
	// The unique immutable identifier of a resource (Guid).
	UniqueIdentifier *string `pulumi:"uniqueIdentifier"`
}

type EnvironmentSettingState struct {
	// Describes the user's progress in configuring their environment setting
	ConfigurationState pulumi.StringPtrInput
	// Describes the environment and its resource settings
	Description pulumi.StringPtrInput
	// Time when the template VM was last changed.
	LastChanged pulumi.StringPtrInput
	// Time when the template VM was last sent for publishing.
	LastPublished pulumi.StringPtrInput
	// The details of the latest operation. ex: status, error
	LatestOperationResult LatestOperationResultResponsePtrInput
	// The location of the resource.
	Location pulumi.StringPtrInput
	// The name of the resource.
	Name pulumi.StringPtrInput
	// The provisioning status of the resource.
	ProvisioningState pulumi.StringPtrInput
	// Describes the readiness of this environment setting
	PublishingState pulumi.StringPtrInput
	// The resource specific settings
	ResourceSettings ResourceSettingsResponsePtrInput
	// The tags of the resource.
	Tags pulumi.StringMapInput
	// Brief title describing the environment and its resource settings
	Title pulumi.StringPtrInput
	// The type of the resource.
	Type pulumi.StringPtrInput
	// The unique immutable identifier of a resource (Guid).
	UniqueIdentifier pulumi.StringPtrInput
}

func (EnvironmentSettingState) ElementType() reflect.Type {
	return reflect.TypeOf((*environmentSettingState)(nil)).Elem()
}

type environmentSettingArgs struct {
	// Describes the user's progress in configuring their environment setting
	ConfigurationState *string `pulumi:"configurationState"`
	// Describes the environment and its resource settings
	Description *string `pulumi:"description"`
	// The name of the environment Setting.
	EnvironmentSettingName string `pulumi:"environmentSettingName"`
	// The name of the lab Account.
	LabAccountName string `pulumi:"labAccountName"`
	// The name of the lab.
	LabName string `pulumi:"labName"`
	// The location of the resource.
	Location *string `pulumi:"location"`
	// The provisioning status of the resource.
	ProvisioningState *string `pulumi:"provisioningState"`
	// The name of the resource group.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The resource specific settings
	ResourceSettings ResourceSettings `pulumi:"resourceSettings"`
	// The tags of the resource.
	Tags map[string]string `pulumi:"tags"`
	// Brief title describing the environment and its resource settings
	Title *string `pulumi:"title"`
	// The unique immutable identifier of a resource (Guid).
	UniqueIdentifier *string `pulumi:"uniqueIdentifier"`
}

// The set of arguments for constructing a EnvironmentSetting resource.
type EnvironmentSettingArgs struct {
	// Describes the user's progress in configuring their environment setting
	ConfigurationState pulumi.StringPtrInput
	// Describes the environment and its resource settings
	Description pulumi.StringPtrInput
	// The name of the environment Setting.
	EnvironmentSettingName pulumi.StringInput
	// The name of the lab Account.
	LabAccountName pulumi.StringInput
	// The name of the lab.
	LabName pulumi.StringInput
	// The location of the resource.
	Location pulumi.StringPtrInput
	// The provisioning status of the resource.
	ProvisioningState pulumi.StringPtrInput
	// The name of the resource group.
	ResourceGroupName pulumi.StringInput
	// The resource specific settings
	ResourceSettings ResourceSettingsInput
	// The tags of the resource.
	Tags pulumi.StringMapInput
	// Brief title describing the environment and its resource settings
	Title pulumi.StringPtrInput
	// The unique immutable identifier of a resource (Guid).
	UniqueIdentifier pulumi.StringPtrInput
}

func (EnvironmentSettingArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*environmentSettingArgs)(nil)).Elem()
}