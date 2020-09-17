// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v20151101preview

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// A user-defined logical grouping of machines.
type MachineGroup struct {
	pulumi.CustomResourceState

	// Count of machines in this group. The value of count may be bigger than the number of machines in case of the group has been truncated due to exceeding the max number of machines a group can handle.
	Count pulumi.IntPtrOutput `pulumi:"count"`
	// User defined name for the group
	DisplayName pulumi.StringOutput `pulumi:"displayName"`
	// Resource ETAG.
	Etag pulumi.StringPtrOutput `pulumi:"etag"`
	// Type of the machine group
	GroupType pulumi.StringPtrOutput `pulumi:"groupType"`
	// Additional resource type qualifier.
	Kind pulumi.StringOutput `pulumi:"kind"`
	// References of the machines in this group. The hints within each reference do not represent the current value of the corresponding fields. They are a snapshot created during the last time the machine group was updated.
	Machines MachineReferenceWithHintsResponseArrayOutput `pulumi:"machines"`
	// Resource name.
	Name pulumi.StringOutput `pulumi:"name"`
	// Resource type.
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewMachineGroup registers a new resource with the given unique name, arguments, and options.
func NewMachineGroup(ctx *pulumi.Context,
	name string, args *MachineGroupArgs, opts ...pulumi.ResourceOption) (*MachineGroup, error) {
	if args == nil || args.DisplayName == nil {
		return nil, errors.New("missing required argument 'DisplayName'")
	}
	if args == nil || args.Kind == nil {
		return nil, errors.New("missing required argument 'Kind'")
	}
	if args == nil || args.MachineGroupName == nil {
		return nil, errors.New("missing required argument 'MachineGroupName'")
	}
	if args == nil || args.ResourceGroupName == nil {
		return nil, errors.New("missing required argument 'ResourceGroupName'")
	}
	if args == nil || args.WorkspaceName == nil {
		return nil, errors.New("missing required argument 'WorkspaceName'")
	}
	if args == nil {
		args = &MachineGroupArgs{}
	}
	args.Kind = pulumi.String("machineGroup")
	var resource MachineGroup
	err := ctx.RegisterResource("azure-nextgen:operationalinsights/v20151101preview:MachineGroup", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetMachineGroup gets an existing MachineGroup resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetMachineGroup(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *MachineGroupState, opts ...pulumi.ResourceOption) (*MachineGroup, error) {
	var resource MachineGroup
	err := ctx.ReadResource("azure-nextgen:operationalinsights/v20151101preview:MachineGroup", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering MachineGroup resources.
type machineGroupState struct {
	// Count of machines in this group. The value of count may be bigger than the number of machines in case of the group has been truncated due to exceeding the max number of machines a group can handle.
	Count *int `pulumi:"count"`
	// User defined name for the group
	DisplayName *string `pulumi:"displayName"`
	// Resource ETAG.
	Etag *string `pulumi:"etag"`
	// Type of the machine group
	GroupType *string `pulumi:"groupType"`
	// Additional resource type qualifier.
	Kind *string `pulumi:"kind"`
	// References of the machines in this group. The hints within each reference do not represent the current value of the corresponding fields. They are a snapshot created during the last time the machine group was updated.
	Machines []MachineReferenceWithHintsResponse `pulumi:"machines"`
	// Resource name.
	Name *string `pulumi:"name"`
	// Resource type.
	Type *string `pulumi:"type"`
}

type MachineGroupState struct {
	// Count of machines in this group. The value of count may be bigger than the number of machines in case of the group has been truncated due to exceeding the max number of machines a group can handle.
	Count pulumi.IntPtrInput
	// User defined name for the group
	DisplayName pulumi.StringPtrInput
	// Resource ETAG.
	Etag pulumi.StringPtrInput
	// Type of the machine group
	GroupType pulumi.StringPtrInput
	// Additional resource type qualifier.
	Kind pulumi.StringPtrInput
	// References of the machines in this group. The hints within each reference do not represent the current value of the corresponding fields. They are a snapshot created during the last time the machine group was updated.
	Machines MachineReferenceWithHintsResponseArrayInput
	// Resource name.
	Name pulumi.StringPtrInput
	// Resource type.
	Type pulumi.StringPtrInput
}

func (MachineGroupState) ElementType() reflect.Type {
	return reflect.TypeOf((*machineGroupState)(nil)).Elem()
}

type machineGroupArgs struct {
	// Count of machines in this group. The value of count may be bigger than the number of machines in case of the group has been truncated due to exceeding the max number of machines a group can handle.
	Count *int `pulumi:"count"`
	// User defined name for the group
	DisplayName string `pulumi:"displayName"`
	// Resource ETAG.
	Etag *string `pulumi:"etag"`
	// Type of the machine group
	GroupType *string `pulumi:"groupType"`
	// Additional resource type qualifier.
	Kind string `pulumi:"kind"`
	// Machine Group resource name.
	MachineGroupName string `pulumi:"machineGroupName"`
	// References of the machines in this group. The hints within each reference do not represent the current value of the corresponding fields. They are a snapshot created during the last time the machine group was updated.
	Machines []MachineReferenceWithHints `pulumi:"machines"`
	// Resource group name within the specified subscriptionId.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// OMS workspace containing the resources of interest.
	WorkspaceName string `pulumi:"workspaceName"`
}

// The set of arguments for constructing a MachineGroup resource.
type MachineGroupArgs struct {
	// Count of machines in this group. The value of count may be bigger than the number of machines in case of the group has been truncated due to exceeding the max number of machines a group can handle.
	Count pulumi.IntPtrInput
	// User defined name for the group
	DisplayName pulumi.StringInput
	// Resource ETAG.
	Etag pulumi.StringPtrInput
	// Type of the machine group
	GroupType pulumi.StringPtrInput
	// Additional resource type qualifier.
	Kind pulumi.StringInput
	// Machine Group resource name.
	MachineGroupName pulumi.StringInput
	// References of the machines in this group. The hints within each reference do not represent the current value of the corresponding fields. They are a snapshot created during the last time the machine group was updated.
	Machines MachineReferenceWithHintsArrayInput
	// Resource group name within the specified subscriptionId.
	ResourceGroupName pulumi.StringInput
	// OMS workspace containing the resources of interest.
	WorkspaceName pulumi.StringInput
}

func (MachineGroupArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*machineGroupArgs)(nil)).Elem()
}