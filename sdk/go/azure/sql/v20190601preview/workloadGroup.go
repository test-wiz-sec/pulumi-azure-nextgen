// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v20190601preview

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Workload group operations for a data warehouse
type WorkloadGroup struct {
	pulumi.CustomResourceState

	// The workload group importance level.
	Importance pulumi.StringPtrOutput `pulumi:"importance"`
	// The workload group cap percentage resource.
	MaxResourcePercent pulumi.IntOutput `pulumi:"maxResourcePercent"`
	// The workload group request maximum grant percentage.
	MaxResourcePercentPerRequest pulumi.Float64PtrOutput `pulumi:"maxResourcePercentPerRequest"`
	// The workload group minimum percentage resource.
	MinResourcePercent pulumi.IntOutput `pulumi:"minResourcePercent"`
	// The workload group request minimum grant percentage.
	MinResourcePercentPerRequest pulumi.Float64Output `pulumi:"minResourcePercentPerRequest"`
	// Resource name.
	Name pulumi.StringOutput `pulumi:"name"`
	// The workload group query execution timeout.
	QueryExecutionTimeout pulumi.IntPtrOutput `pulumi:"queryExecutionTimeout"`
	// Resource type.
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewWorkloadGroup registers a new resource with the given unique name, arguments, and options.
func NewWorkloadGroup(ctx *pulumi.Context,
	name string, args *WorkloadGroupArgs, opts ...pulumi.ResourceOption) (*WorkloadGroup, error) {
	if args == nil || args.DatabaseName == nil {
		return nil, errors.New("missing required argument 'DatabaseName'")
	}
	if args == nil || args.MaxResourcePercent == nil {
		return nil, errors.New("missing required argument 'MaxResourcePercent'")
	}
	if args == nil || args.MinResourcePercent == nil {
		return nil, errors.New("missing required argument 'MinResourcePercent'")
	}
	if args == nil || args.MinResourcePercentPerRequest == nil {
		return nil, errors.New("missing required argument 'MinResourcePercentPerRequest'")
	}
	if args == nil || args.ResourceGroupName == nil {
		return nil, errors.New("missing required argument 'ResourceGroupName'")
	}
	if args == nil || args.ServerName == nil {
		return nil, errors.New("missing required argument 'ServerName'")
	}
	if args == nil || args.WorkloadGroupName == nil {
		return nil, errors.New("missing required argument 'WorkloadGroupName'")
	}
	if args == nil {
		args = &WorkloadGroupArgs{}
	}
	var resource WorkloadGroup
	err := ctx.RegisterResource("azure-nextgen:sql/v20190601preview:WorkloadGroup", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetWorkloadGroup gets an existing WorkloadGroup resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetWorkloadGroup(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *WorkloadGroupState, opts ...pulumi.ResourceOption) (*WorkloadGroup, error) {
	var resource WorkloadGroup
	err := ctx.ReadResource("azure-nextgen:sql/v20190601preview:WorkloadGroup", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering WorkloadGroup resources.
type workloadGroupState struct {
	// The workload group importance level.
	Importance *string `pulumi:"importance"`
	// The workload group cap percentage resource.
	MaxResourcePercent *int `pulumi:"maxResourcePercent"`
	// The workload group request maximum grant percentage.
	MaxResourcePercentPerRequest *float64 `pulumi:"maxResourcePercentPerRequest"`
	// The workload group minimum percentage resource.
	MinResourcePercent *int `pulumi:"minResourcePercent"`
	// The workload group request minimum grant percentage.
	MinResourcePercentPerRequest *float64 `pulumi:"minResourcePercentPerRequest"`
	// Resource name.
	Name *string `pulumi:"name"`
	// The workload group query execution timeout.
	QueryExecutionTimeout *int `pulumi:"queryExecutionTimeout"`
	// Resource type.
	Type *string `pulumi:"type"`
}

type WorkloadGroupState struct {
	// The workload group importance level.
	Importance pulumi.StringPtrInput
	// The workload group cap percentage resource.
	MaxResourcePercent pulumi.IntPtrInput
	// The workload group request maximum grant percentage.
	MaxResourcePercentPerRequest pulumi.Float64PtrInput
	// The workload group minimum percentage resource.
	MinResourcePercent pulumi.IntPtrInput
	// The workload group request minimum grant percentage.
	MinResourcePercentPerRequest pulumi.Float64PtrInput
	// Resource name.
	Name pulumi.StringPtrInput
	// The workload group query execution timeout.
	QueryExecutionTimeout pulumi.IntPtrInput
	// Resource type.
	Type pulumi.StringPtrInput
}

func (WorkloadGroupState) ElementType() reflect.Type {
	return reflect.TypeOf((*workloadGroupState)(nil)).Elem()
}

type workloadGroupArgs struct {
	// The name of the database.
	DatabaseName string `pulumi:"databaseName"`
	// The workload group importance level.
	Importance *string `pulumi:"importance"`
	// The workload group cap percentage resource.
	MaxResourcePercent int `pulumi:"maxResourcePercent"`
	// The workload group request maximum grant percentage.
	MaxResourcePercentPerRequest *float64 `pulumi:"maxResourcePercentPerRequest"`
	// The workload group minimum percentage resource.
	MinResourcePercent int `pulumi:"minResourcePercent"`
	// The workload group request minimum grant percentage.
	MinResourcePercentPerRequest float64 `pulumi:"minResourcePercentPerRequest"`
	// The workload group query execution timeout.
	QueryExecutionTimeout *int `pulumi:"queryExecutionTimeout"`
	// The name of the resource group that contains the resource. You can obtain this value from the Azure Resource Manager API or the portal.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The name of the server.
	ServerName string `pulumi:"serverName"`
	// The name of the workload group.
	WorkloadGroupName string `pulumi:"workloadGroupName"`
}

// The set of arguments for constructing a WorkloadGroup resource.
type WorkloadGroupArgs struct {
	// The name of the database.
	DatabaseName pulumi.StringInput
	// The workload group importance level.
	Importance pulumi.StringPtrInput
	// The workload group cap percentage resource.
	MaxResourcePercent pulumi.IntInput
	// The workload group request maximum grant percentage.
	MaxResourcePercentPerRequest pulumi.Float64PtrInput
	// The workload group minimum percentage resource.
	MinResourcePercent pulumi.IntInput
	// The workload group request minimum grant percentage.
	MinResourcePercentPerRequest pulumi.Float64Input
	// The workload group query execution timeout.
	QueryExecutionTimeout pulumi.IntPtrInput
	// The name of the resource group that contains the resource. You can obtain this value from the Azure Resource Manager API or the portal.
	ResourceGroupName pulumi.StringInput
	// The name of the server.
	ServerName pulumi.StringInput
	// The name of the workload group.
	WorkloadGroupName pulumi.StringInput
}

func (WorkloadGroupArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*workloadGroupArgs)(nil)).Elem()
}
