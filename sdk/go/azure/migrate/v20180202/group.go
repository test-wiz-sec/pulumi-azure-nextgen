// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v20180202

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// A group created in a Migration project.
type Group struct {
	pulumi.CustomResourceState

	// List of References to Assessments created on this group.
	Assessments pulumi.StringArrayOutput `pulumi:"assessments"`
	// Time when this project was created. Date-Time represented in ISO-8601 format.
	CreatedTimestamp pulumi.StringOutput `pulumi:"createdTimestamp"`
	// For optimistic concurrency control.
	ETag pulumi.StringPtrOutput `pulumi:"eTag"`
	// List of machine names that are part of this group.
	Machines pulumi.StringArrayOutput `pulumi:"machines"`
	// Name of the group.
	Name pulumi.StringOutput `pulumi:"name"`
	// Type of the object = [Microsoft.Migrate/projects/groups].
	Type pulumi.StringOutput `pulumi:"type"`
	// Time when this project was last updated. Date-Time represented in ISO-8601 format.
	UpdatedTimestamp pulumi.StringOutput `pulumi:"updatedTimestamp"`
}

// NewGroup registers a new resource with the given unique name, arguments, and options.
func NewGroup(ctx *pulumi.Context,
	name string, args *GroupArgs, opts ...pulumi.ResourceOption) (*Group, error) {
	if args == nil || args.GroupName == nil {
		return nil, errors.New("missing required argument 'GroupName'")
	}
	if args == nil || args.Machines == nil {
		return nil, errors.New("missing required argument 'Machines'")
	}
	if args == nil || args.ProjectName == nil {
		return nil, errors.New("missing required argument 'ProjectName'")
	}
	if args == nil || args.ResourceGroupName == nil {
		return nil, errors.New("missing required argument 'ResourceGroupName'")
	}
	if args == nil {
		args = &GroupArgs{}
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-nextgen:migrate/v20171111preview:Group"),
		},
	})
	opts = append(opts, aliases)
	var resource Group
	err := ctx.RegisterResource("azure-nextgen:migrate/v20180202:Group", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetGroup gets an existing Group resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetGroup(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *GroupState, opts ...pulumi.ResourceOption) (*Group, error) {
	var resource Group
	err := ctx.ReadResource("azure-nextgen:migrate/v20180202:Group", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Group resources.
type groupState struct {
	// List of References to Assessments created on this group.
	Assessments []string `pulumi:"assessments"`
	// Time when this project was created. Date-Time represented in ISO-8601 format.
	CreatedTimestamp *string `pulumi:"createdTimestamp"`
	// For optimistic concurrency control.
	ETag *string `pulumi:"eTag"`
	// List of machine names that are part of this group.
	Machines []string `pulumi:"machines"`
	// Name of the group.
	Name *string `pulumi:"name"`
	// Type of the object = [Microsoft.Migrate/projects/groups].
	Type *string `pulumi:"type"`
	// Time when this project was last updated. Date-Time represented in ISO-8601 format.
	UpdatedTimestamp *string `pulumi:"updatedTimestamp"`
}

type GroupState struct {
	// List of References to Assessments created on this group.
	Assessments pulumi.StringArrayInput
	// Time when this project was created. Date-Time represented in ISO-8601 format.
	CreatedTimestamp pulumi.StringPtrInput
	// For optimistic concurrency control.
	ETag pulumi.StringPtrInput
	// List of machine names that are part of this group.
	Machines pulumi.StringArrayInput
	// Name of the group.
	Name pulumi.StringPtrInput
	// Type of the object = [Microsoft.Migrate/projects/groups].
	Type pulumi.StringPtrInput
	// Time when this project was last updated. Date-Time represented in ISO-8601 format.
	UpdatedTimestamp pulumi.StringPtrInput
}

func (GroupState) ElementType() reflect.Type {
	return reflect.TypeOf((*groupState)(nil)).Elem()
}

type groupArgs struct {
	// For optimistic concurrency control.
	ETag *string `pulumi:"eTag"`
	// Unique name of a group within a project.
	GroupName string `pulumi:"groupName"`
	// List of machine names that are part of this group.
	Machines []string `pulumi:"machines"`
	// Name of the Azure Migrate project.
	ProjectName string `pulumi:"projectName"`
	// Name of the Azure Resource Group that project is part of.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// The set of arguments for constructing a Group resource.
type GroupArgs struct {
	// For optimistic concurrency control.
	ETag pulumi.StringPtrInput
	// Unique name of a group within a project.
	GroupName pulumi.StringInput
	// List of machine names that are part of this group.
	Machines pulumi.StringArrayInput
	// Name of the Azure Migrate project.
	ProjectName pulumi.StringInput
	// Name of the Azure Resource Group that project is part of.
	ResourceGroupName pulumi.StringInput
}

func (GroupArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*groupArgs)(nil)).Elem()
}
