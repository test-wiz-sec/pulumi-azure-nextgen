// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v20190601preview

import (
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

func LookupWorkspace(ctx *pulumi.Context, args *LookupWorkspaceArgs, opts ...pulumi.InvokeOption) (*LookupWorkspaceResult, error) {
	var rv LookupWorkspaceResult
	err := ctx.Invoke("azure-nextgen:synapse/v20190601preview:getWorkspace", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupWorkspaceArgs struct {
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The name of the workspace
	WorkspaceName string `pulumi:"workspaceName"`
}

// A workspace
type LookupWorkspaceResult struct {
	// Connectivity endpoints
	ConnectivityEndpoints map[string]string `pulumi:"connectivityEndpoints"`
	// Workspace default data lake storage account details
	DefaultDataLakeStorage *DataLakeStorageAccountDetailsResponse `pulumi:"defaultDataLakeStorage"`
	// Workspace level configs and feature flags
	ExtraProperties map[string]interface{} `pulumi:"extraProperties"`
	// Identity of the workspace
	Identity *ManagedIdentityResponse `pulumi:"identity"`
	// The geo-location where the resource lives
	Location string `pulumi:"location"`
	// Workspace managed resource group. The resource group name uniquely identifies the resource group within the user subscriptionId. The resource group name must be no longer than 90 characters long, and must be alphanumeric characters (Char.IsLetterOrDigit()) and '-', '_', '(', ')' and'.'. Note that the name cannot end with '.'
	ManagedResourceGroupName *string `pulumi:"managedResourceGroupName"`
	// Setting this to 'default' will ensure that all compute for this workspace is in a virtual network managed on behalf of the user.
	ManagedVirtualNetwork *string `pulumi:"managedVirtualNetwork"`
	// Managed Virtual Network Settings
	ManagedVirtualNetworkSettings *ManagedVirtualNetworkSettingsResponse `pulumi:"managedVirtualNetworkSettings"`
	// The name of the resource
	Name string `pulumi:"name"`
	// Private endpoint connections to the workspace
	PrivateEndpointConnections []PrivateEndpointConnectionResponse `pulumi:"privateEndpointConnections"`
	// Resource provisioning state
	ProvisioningState string `pulumi:"provisioningState"`
	// Login for workspace SQL active directory administrator
	SqlAdministratorLogin *string `pulumi:"sqlAdministratorLogin"`
	// SQL administrator login password
	SqlAdministratorLoginPassword *string `pulumi:"sqlAdministratorLoginPassword"`
	// Resource tags.
	Tags map[string]string `pulumi:"tags"`
	// The type of the resource. Ex- Microsoft.Compute/virtualMachines or Microsoft.Storage/storageAccounts.
	Type string `pulumi:"type"`
	// Virtual Network profile
	VirtualNetworkProfile *VirtualNetworkProfileResponse `pulumi:"virtualNetworkProfile"`
}
