// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v20171111preview

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Represents a Blueprint assignment.
type Assignment struct {
	pulumi.CustomResourceState

	// ID of the Blueprint definition resource.
	BlueprintId pulumi.StringPtrOutput `pulumi:"blueprintId"`
	// Multi-line explain this resource.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// One-liner string explain this resource.
	DisplayName pulumi.StringPtrOutput `pulumi:"displayName"`
	// Managed Service Identity for this Blueprint assignment
	Identity ManagedServiceIdentityResponseOutput `pulumi:"identity"`
	// The location of this Blueprint assignment.
	Location pulumi.StringOutput `pulumi:"location"`
	// Defines how Blueprint-managed resources will be locked.
	Locks AssignmentLockSettingsResponsePtrOutput `pulumi:"locks"`
	// Name of this resource.
	Name pulumi.StringOutput `pulumi:"name"`
	// Blueprint parameter values.
	Parameters ParameterValueBaseResponseMapOutput `pulumi:"parameters"`
	// State of the assignment.
	ProvisioningState pulumi.StringOutput `pulumi:"provisioningState"`
	// Names and locations of resource group placeholders.
	ResourceGroups ResourceGroupValueResponseMapOutput `pulumi:"resourceGroups"`
	// Status of Blueprint assignment. This field is readonly.
	Status AssignmentStatusResponseOutput `pulumi:"status"`
	// Type of this resource.
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewAssignment registers a new resource with the given unique name, arguments, and options.
func NewAssignment(ctx *pulumi.Context,
	name string, args *AssignmentArgs, opts ...pulumi.ResourceOption) (*Assignment, error) {
	if args == nil || args.AssignmentName == nil {
		return nil, errors.New("missing required argument 'AssignmentName'")
	}
	if args == nil || args.Identity == nil {
		return nil, errors.New("missing required argument 'Identity'")
	}
	if args == nil || args.Location == nil {
		return nil, errors.New("missing required argument 'Location'")
	}
	if args == nil || args.Parameters == nil {
		return nil, errors.New("missing required argument 'Parameters'")
	}
	if args == nil || args.ResourceGroups == nil {
		return nil, errors.New("missing required argument 'ResourceGroups'")
	}
	if args == nil {
		args = &AssignmentArgs{}
	}
	var resource Assignment
	err := ctx.RegisterResource("azure-nextgen:blueprint/v20171111preview:Assignment", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetAssignment gets an existing Assignment resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetAssignment(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *AssignmentState, opts ...pulumi.ResourceOption) (*Assignment, error) {
	var resource Assignment
	err := ctx.ReadResource("azure-nextgen:blueprint/v20171111preview:Assignment", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Assignment resources.
type assignmentState struct {
	// ID of the Blueprint definition resource.
	BlueprintId *string `pulumi:"blueprintId"`
	// Multi-line explain this resource.
	Description *string `pulumi:"description"`
	// One-liner string explain this resource.
	DisplayName *string `pulumi:"displayName"`
	// Managed Service Identity for this Blueprint assignment
	Identity *ManagedServiceIdentityResponse `pulumi:"identity"`
	// The location of this Blueprint assignment.
	Location *string `pulumi:"location"`
	// Defines how Blueprint-managed resources will be locked.
	Locks *AssignmentLockSettingsResponse `pulumi:"locks"`
	// Name of this resource.
	Name *string `pulumi:"name"`
	// Blueprint parameter values.
	Parameters map[string]ParameterValueBaseResponse `pulumi:"parameters"`
	// State of the assignment.
	ProvisioningState *string `pulumi:"provisioningState"`
	// Names and locations of resource group placeholders.
	ResourceGroups map[string]ResourceGroupValueResponse `pulumi:"resourceGroups"`
	// Status of Blueprint assignment. This field is readonly.
	Status *AssignmentStatusResponse `pulumi:"status"`
	// Type of this resource.
	Type *string `pulumi:"type"`
}

type AssignmentState struct {
	// ID of the Blueprint definition resource.
	BlueprintId pulumi.StringPtrInput
	// Multi-line explain this resource.
	Description pulumi.StringPtrInput
	// One-liner string explain this resource.
	DisplayName pulumi.StringPtrInput
	// Managed Service Identity for this Blueprint assignment
	Identity ManagedServiceIdentityResponsePtrInput
	// The location of this Blueprint assignment.
	Location pulumi.StringPtrInput
	// Defines how Blueprint-managed resources will be locked.
	Locks AssignmentLockSettingsResponsePtrInput
	// Name of this resource.
	Name pulumi.StringPtrInput
	// Blueprint parameter values.
	Parameters ParameterValueBaseResponseMapInput
	// State of the assignment.
	ProvisioningState pulumi.StringPtrInput
	// Names and locations of resource group placeholders.
	ResourceGroups ResourceGroupValueResponseMapInput
	// Status of Blueprint assignment. This field is readonly.
	Status AssignmentStatusResponsePtrInput
	// Type of this resource.
	Type pulumi.StringPtrInput
}

func (AssignmentState) ElementType() reflect.Type {
	return reflect.TypeOf((*assignmentState)(nil)).Elem()
}

type assignmentArgs struct {
	// name of the assignment.
	AssignmentName string `pulumi:"assignmentName"`
	// ID of the Blueprint definition resource.
	BlueprintId *string `pulumi:"blueprintId"`
	// Multi-line explain this resource.
	Description *string `pulumi:"description"`
	// One-liner string explain this resource.
	DisplayName *string `pulumi:"displayName"`
	// Managed Service Identity for this Blueprint assignment
	Identity ManagedServiceIdentity `pulumi:"identity"`
	// The location of this Blueprint assignment.
	Location string `pulumi:"location"`
	// Defines how Blueprint-managed resources will be locked.
	Locks *AssignmentLockSettings `pulumi:"locks"`
	// Blueprint parameter values.
	Parameters map[string]ParameterValueBase `pulumi:"parameters"`
	// Names and locations of resource group placeholders.
	ResourceGroups map[string]ResourceGroupValue `pulumi:"resourceGroups"`
}

// The set of arguments for constructing a Assignment resource.
type AssignmentArgs struct {
	// name of the assignment.
	AssignmentName pulumi.StringInput
	// ID of the Blueprint definition resource.
	BlueprintId pulumi.StringPtrInput
	// Multi-line explain this resource.
	Description pulumi.StringPtrInput
	// One-liner string explain this resource.
	DisplayName pulumi.StringPtrInput
	// Managed Service Identity for this Blueprint assignment
	Identity ManagedServiceIdentityInput
	// The location of this Blueprint assignment.
	Location pulumi.StringInput
	// Defines how Blueprint-managed resources will be locked.
	Locks AssignmentLockSettingsPtrInput
	// Blueprint parameter values.
	Parameters ParameterValueBaseMapInput
	// Names and locations of resource group placeholders.
	ResourceGroups ResourceGroupValueMapInput
}

func (AssignmentArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*assignmentArgs)(nil)).Elem()
}
