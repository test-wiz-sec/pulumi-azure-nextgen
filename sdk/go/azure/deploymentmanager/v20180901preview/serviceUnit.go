// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v20180901preview

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Represents the response of a service unit resource.
type ServiceUnit struct {
	pulumi.CustomResourceState

	// The artifacts for the service unit.
	Artifacts ServiceUnitArtifactsResponsePtrOutput `pulumi:"artifacts"`
	// Describes the type of ARM deployment to be performed on the resource.
	DeploymentMode pulumi.StringOutput `pulumi:"deploymentMode"`
	// The geo-location where the resource lives
	Location pulumi.StringOutput `pulumi:"location"`
	// The name of the resource
	Name pulumi.StringOutput `pulumi:"name"`
	// Resource tags.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// The Azure Resource Group to which the resources in the service unit belong to or should be deployed to.
	TargetResourceGroup pulumi.StringOutput `pulumi:"targetResourceGroup"`
	// The type of the resource. Ex- Microsoft.Compute/virtualMachines or Microsoft.Storage/storageAccounts.
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewServiceUnit registers a new resource with the given unique name, arguments, and options.
func NewServiceUnit(ctx *pulumi.Context,
	name string, args *ServiceUnitArgs, opts ...pulumi.ResourceOption) (*ServiceUnit, error) {
	if args == nil || args.DeploymentMode == nil {
		return nil, errors.New("missing required argument 'DeploymentMode'")
	}
	if args == nil || args.Location == nil {
		return nil, errors.New("missing required argument 'Location'")
	}
	if args == nil || args.ResourceGroupName == nil {
		return nil, errors.New("missing required argument 'ResourceGroupName'")
	}
	if args == nil || args.ServiceName == nil {
		return nil, errors.New("missing required argument 'ServiceName'")
	}
	if args == nil || args.ServiceTopologyName == nil {
		return nil, errors.New("missing required argument 'ServiceTopologyName'")
	}
	if args == nil || args.ServiceUnitName == nil {
		return nil, errors.New("missing required argument 'ServiceUnitName'")
	}
	if args == nil || args.TargetResourceGroup == nil {
		return nil, errors.New("missing required argument 'TargetResourceGroup'")
	}
	if args == nil {
		args = &ServiceUnitArgs{}
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-nextgen:deploymentmanager/v20191101preview:ServiceUnit"),
		},
	})
	opts = append(opts, aliases)
	var resource ServiceUnit
	err := ctx.RegisterResource("azure-nextgen:deploymentmanager/v20180901preview:ServiceUnit", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetServiceUnit gets an existing ServiceUnit resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetServiceUnit(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ServiceUnitState, opts ...pulumi.ResourceOption) (*ServiceUnit, error) {
	var resource ServiceUnit
	err := ctx.ReadResource("azure-nextgen:deploymentmanager/v20180901preview:ServiceUnit", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ServiceUnit resources.
type serviceUnitState struct {
	// The artifacts for the service unit.
	Artifacts *ServiceUnitArtifactsResponse `pulumi:"artifacts"`
	// Describes the type of ARM deployment to be performed on the resource.
	DeploymentMode *string `pulumi:"deploymentMode"`
	// The geo-location where the resource lives
	Location *string `pulumi:"location"`
	// The name of the resource
	Name *string `pulumi:"name"`
	// Resource tags.
	Tags map[string]string `pulumi:"tags"`
	// The Azure Resource Group to which the resources in the service unit belong to or should be deployed to.
	TargetResourceGroup *string `pulumi:"targetResourceGroup"`
	// The type of the resource. Ex- Microsoft.Compute/virtualMachines or Microsoft.Storage/storageAccounts.
	Type *string `pulumi:"type"`
}

type ServiceUnitState struct {
	// The artifacts for the service unit.
	Artifacts ServiceUnitArtifactsResponsePtrInput
	// Describes the type of ARM deployment to be performed on the resource.
	DeploymentMode pulumi.StringPtrInput
	// The geo-location where the resource lives
	Location pulumi.StringPtrInput
	// The name of the resource
	Name pulumi.StringPtrInput
	// Resource tags.
	Tags pulumi.StringMapInput
	// The Azure Resource Group to which the resources in the service unit belong to or should be deployed to.
	TargetResourceGroup pulumi.StringPtrInput
	// The type of the resource. Ex- Microsoft.Compute/virtualMachines or Microsoft.Storage/storageAccounts.
	Type pulumi.StringPtrInput
}

func (ServiceUnitState) ElementType() reflect.Type {
	return reflect.TypeOf((*serviceUnitState)(nil)).Elem()
}

type serviceUnitArgs struct {
	// The artifacts for the service unit.
	Artifacts *ServiceUnitArtifacts `pulumi:"artifacts"`
	// Describes the type of ARM deployment to be performed on the resource.
	DeploymentMode string `pulumi:"deploymentMode"`
	// The geo-location where the resource lives
	Location string `pulumi:"location"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The name of the service resource.
	ServiceName string `pulumi:"serviceName"`
	// The name of the service topology .
	ServiceTopologyName string `pulumi:"serviceTopologyName"`
	// The name of the service unit resource.
	ServiceUnitName string `pulumi:"serviceUnitName"`
	// Resource tags.
	Tags map[string]string `pulumi:"tags"`
	// The Azure Resource Group to which the resources in the service unit belong to or should be deployed to.
	TargetResourceGroup string `pulumi:"targetResourceGroup"`
}

// The set of arguments for constructing a ServiceUnit resource.
type ServiceUnitArgs struct {
	// The artifacts for the service unit.
	Artifacts ServiceUnitArtifactsPtrInput
	// Describes the type of ARM deployment to be performed on the resource.
	DeploymentMode pulumi.StringInput
	// The geo-location where the resource lives
	Location pulumi.StringInput
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput
	// The name of the service resource.
	ServiceName pulumi.StringInput
	// The name of the service topology .
	ServiceTopologyName pulumi.StringInput
	// The name of the service unit resource.
	ServiceUnitName pulumi.StringInput
	// Resource tags.
	Tags pulumi.StringMapInput
	// The Azure Resource Group to which the resources in the service unit belong to or should be deployed to.
	TargetResourceGroup pulumi.StringInput
}

func (ServiceUnitArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*serviceUnitArgs)(nil)).Elem()
}