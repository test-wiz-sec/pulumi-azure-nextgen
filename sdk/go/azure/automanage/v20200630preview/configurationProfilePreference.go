// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v20200630preview

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Definition of the configuration profile preference.
type ConfigurationProfilePreference struct {
	pulumi.CustomResourceState

	// The geo-location where the resource lives
	Location pulumi.StringOutput `pulumi:"location"`
	// The name of the resource
	Name pulumi.StringOutput `pulumi:"name"`
	// Properties of the configuration profile preference.
	Properties ConfigurationProfilePreferencePropertiesResponseOutput `pulumi:"properties"`
	// Resource tags.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// The type of the resource. Ex- Microsoft.Compute/virtualMachines or Microsoft.Storage/storageAccounts.
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewConfigurationProfilePreference registers a new resource with the given unique name, arguments, and options.
func NewConfigurationProfilePreference(ctx *pulumi.Context,
	name string, args *ConfigurationProfilePreferenceArgs, opts ...pulumi.ResourceOption) (*ConfigurationProfilePreference, error) {
	if args == nil || args.ConfigurationProfilePreferenceName == nil {
		return nil, errors.New("missing required argument 'ConfigurationProfilePreferenceName'")
	}
	if args == nil || args.Location == nil {
		return nil, errors.New("missing required argument 'Location'")
	}
	if args == nil || args.ResourceGroupName == nil {
		return nil, errors.New("missing required argument 'ResourceGroupName'")
	}
	if args == nil {
		args = &ConfigurationProfilePreferenceArgs{}
	}
	var resource ConfigurationProfilePreference
	err := ctx.RegisterResource("azure-nextgen:automanage/v20200630preview:ConfigurationProfilePreference", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetConfigurationProfilePreference gets an existing ConfigurationProfilePreference resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetConfigurationProfilePreference(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ConfigurationProfilePreferenceState, opts ...pulumi.ResourceOption) (*ConfigurationProfilePreference, error) {
	var resource ConfigurationProfilePreference
	err := ctx.ReadResource("azure-nextgen:automanage/v20200630preview:ConfigurationProfilePreference", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ConfigurationProfilePreference resources.
type configurationProfilePreferenceState struct {
	// The geo-location where the resource lives
	Location *string `pulumi:"location"`
	// The name of the resource
	Name *string `pulumi:"name"`
	// Properties of the configuration profile preference.
	Properties *ConfigurationProfilePreferencePropertiesResponse `pulumi:"properties"`
	// Resource tags.
	Tags map[string]string `pulumi:"tags"`
	// The type of the resource. Ex- Microsoft.Compute/virtualMachines or Microsoft.Storage/storageAccounts.
	Type *string `pulumi:"type"`
}

type ConfigurationProfilePreferenceState struct {
	// The geo-location where the resource lives
	Location pulumi.StringPtrInput
	// The name of the resource
	Name pulumi.StringPtrInput
	// Properties of the configuration profile preference.
	Properties ConfigurationProfilePreferencePropertiesResponsePtrInput
	// Resource tags.
	Tags pulumi.StringMapInput
	// The type of the resource. Ex- Microsoft.Compute/virtualMachines or Microsoft.Storage/storageAccounts.
	Type pulumi.StringPtrInput
}

func (ConfigurationProfilePreferenceState) ElementType() reflect.Type {
	return reflect.TypeOf((*configurationProfilePreferenceState)(nil)).Elem()
}

type configurationProfilePreferenceArgs struct {
	// Name of the configuration profile preference.
	ConfigurationProfilePreferenceName string `pulumi:"configurationProfilePreferenceName"`
	// The geo-location where the resource lives
	Location string `pulumi:"location"`
	// Properties of the configuration profile preference.
	Properties *ConfigurationProfilePreferenceProperties `pulumi:"properties"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// Resource tags.
	Tags map[string]string `pulumi:"tags"`
}

// The set of arguments for constructing a ConfigurationProfilePreference resource.
type ConfigurationProfilePreferenceArgs struct {
	// Name of the configuration profile preference.
	ConfigurationProfilePreferenceName pulumi.StringInput
	// The geo-location where the resource lives
	Location pulumi.StringInput
	// Properties of the configuration profile preference.
	Properties ConfigurationProfilePreferencePropertiesPtrInput
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput
	// Resource tags.
	Tags pulumi.StringMapInput
}

func (ConfigurationProfilePreferenceArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*configurationProfilePreferenceArgs)(nil)).Elem()
}