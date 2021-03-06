// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v20190601preview

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Workspace active directory administrator
type WorkspaceAadAdmin struct {
	pulumi.CustomResourceState

	// Workspace active directory administrator type
	AdministratorType pulumi.StringPtrOutput `pulumi:"administratorType"`
	// Login of the workspace active directory administrator
	Login pulumi.StringPtrOutput `pulumi:"login"`
	// The name of the resource
	Name pulumi.StringOutput `pulumi:"name"`
	// Object ID of the workspace active directory administrator
	Sid pulumi.StringPtrOutput `pulumi:"sid"`
	// Tenant ID of the workspace active directory administrator
	TenantId pulumi.StringPtrOutput `pulumi:"tenantId"`
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewWorkspaceAadAdmin registers a new resource with the given unique name, arguments, and options.
func NewWorkspaceAadAdmin(ctx *pulumi.Context,
	name string, args *WorkspaceAadAdminArgs, opts ...pulumi.ResourceOption) (*WorkspaceAadAdmin, error) {
	if args == nil || args.ResourceGroupName == nil {
		return nil, errors.New("missing required argument 'ResourceGroupName'")
	}
	if args == nil || args.WorkspaceName == nil {
		return nil, errors.New("missing required argument 'WorkspaceName'")
	}
	if args == nil {
		args = &WorkspaceAadAdminArgs{}
	}
	var resource WorkspaceAadAdmin
	err := ctx.RegisterResource("azure-nextgen:synapse/v20190601preview:WorkspaceAadAdmin", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetWorkspaceAadAdmin gets an existing WorkspaceAadAdmin resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetWorkspaceAadAdmin(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *WorkspaceAadAdminState, opts ...pulumi.ResourceOption) (*WorkspaceAadAdmin, error) {
	var resource WorkspaceAadAdmin
	err := ctx.ReadResource("azure-nextgen:synapse/v20190601preview:WorkspaceAadAdmin", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering WorkspaceAadAdmin resources.
type workspaceAadAdminState struct {
	// Workspace active directory administrator type
	AdministratorType *string `pulumi:"administratorType"`
	// Login of the workspace active directory administrator
	Login *string `pulumi:"login"`
	// The name of the resource
	Name *string `pulumi:"name"`
	// Object ID of the workspace active directory administrator
	Sid *string `pulumi:"sid"`
	// Tenant ID of the workspace active directory administrator
	TenantId *string `pulumi:"tenantId"`
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type *string `pulumi:"type"`
}

type WorkspaceAadAdminState struct {
	// Workspace active directory administrator type
	AdministratorType pulumi.StringPtrInput
	// Login of the workspace active directory administrator
	Login pulumi.StringPtrInput
	// The name of the resource
	Name pulumi.StringPtrInput
	// Object ID of the workspace active directory administrator
	Sid pulumi.StringPtrInput
	// Tenant ID of the workspace active directory administrator
	TenantId pulumi.StringPtrInput
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type pulumi.StringPtrInput
}

func (WorkspaceAadAdminState) ElementType() reflect.Type {
	return reflect.TypeOf((*workspaceAadAdminState)(nil)).Elem()
}

type workspaceAadAdminArgs struct {
	// Workspace active directory administrator type
	AdministratorType *string `pulumi:"administratorType"`
	// Login of the workspace active directory administrator
	Login *string `pulumi:"login"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// Object ID of the workspace active directory administrator
	Sid *string `pulumi:"sid"`
	// Tenant ID of the workspace active directory administrator
	TenantId *string `pulumi:"tenantId"`
	// The name of the workspace
	WorkspaceName string `pulumi:"workspaceName"`
}

// The set of arguments for constructing a WorkspaceAadAdmin resource.
type WorkspaceAadAdminArgs struct {
	// Workspace active directory administrator type
	AdministratorType pulumi.StringPtrInput
	// Login of the workspace active directory administrator
	Login pulumi.StringPtrInput
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput
	// Object ID of the workspace active directory administrator
	Sid pulumi.StringPtrInput
	// Tenant ID of the workspace active directory administrator
	TenantId pulumi.StringPtrInput
	// The name of the workspace
	WorkspaceName pulumi.StringInput
}

func (WorkspaceAadAdminArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*workspaceAadAdminArgs)(nil)).Elem()
}
