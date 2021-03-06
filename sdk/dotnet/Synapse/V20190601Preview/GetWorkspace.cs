// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNextGen.Synapse.V20190601Preview
{
    public static class GetWorkspace
    {
        public static Task<GetWorkspaceResult> InvokeAsync(GetWorkspaceArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetWorkspaceResult>("azure-nextgen:synapse/v20190601preview:getWorkspace", args ?? new GetWorkspaceArgs(), options.WithVersion());
    }


    public sealed class GetWorkspaceArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// The name of the resource group. The name is case insensitive.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public string ResourceGroupName { get; set; } = null!;

        /// <summary>
        /// The name of the workspace
        /// </summary>
        [Input("workspaceName", required: true)]
        public string WorkspaceName { get; set; } = null!;

        public GetWorkspaceArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetWorkspaceResult
    {
        /// <summary>
        /// Babylon Configuration
        /// </summary>
        public readonly Outputs.BabylonConfigurationResponse? BabylonConfiguration;
        /// <summary>
        /// Connectivity endpoints
        /// </summary>
        public readonly ImmutableDictionary<string, string>? ConnectivityEndpoints;
        /// <summary>
        /// Workspace default data lake storage account details
        /// </summary>
        public readonly Outputs.DataLakeStorageAccountDetailsResponse? DefaultDataLakeStorage;
        /// <summary>
        /// The encryption details of the workspace
        /// </summary>
        public readonly Outputs.EncryptionDetailsResponse? Encryption;
        /// <summary>
        /// Workspace level configs and feature flags
        /// </summary>
        public readonly ImmutableDictionary<string, object> ExtraProperties;
        /// <summary>
        /// Identity of the workspace
        /// </summary>
        public readonly Outputs.ManagedIdentityResponse? Identity;
        /// <summary>
        /// The geo-location where the resource lives
        /// </summary>
        public readonly string Location;
        /// <summary>
        /// Workspace managed resource group. The resource group name uniquely identifies the resource group within the user subscriptionId. The resource group name must be no longer than 90 characters long, and must be alphanumeric characters (Char.IsLetterOrDigit()) and '-', '_', '(', ')' and'.'. Note that the name cannot end with '.'
        /// </summary>
        public readonly string? ManagedResourceGroupName;
        /// <summary>
        /// Setting this to 'default' will ensure that all compute for this workspace is in a virtual network managed on behalf of the user.
        /// </summary>
        public readonly string? ManagedVirtualNetwork;
        /// <summary>
        /// Managed Virtual Network Settings
        /// </summary>
        public readonly Outputs.ManagedVirtualNetworkSettingsResponse? ManagedVirtualNetworkSettings;
        /// <summary>
        /// The name of the resource
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// Private endpoint connections to the workspace
        /// </summary>
        public readonly ImmutableArray<Outputs.PrivateEndpointConnectionResponse> PrivateEndpointConnections;
        /// <summary>
        /// Resource provisioning state
        /// </summary>
        public readonly string ProvisioningState;
        /// <summary>
        /// Login for workspace SQL active directory administrator
        /// </summary>
        public readonly string? SqlAdministratorLogin;
        /// <summary>
        /// SQL administrator login password
        /// </summary>
        public readonly string? SqlAdministratorLoginPassword;
        /// <summary>
        /// Resource tags.
        /// </summary>
        public readonly ImmutableDictionary<string, string>? Tags;
        /// <summary>
        /// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
        /// </summary>
        public readonly string Type;
        /// <summary>
        /// Virtual Network profile
        /// </summary>
        public readonly Outputs.VirtualNetworkProfileResponse? VirtualNetworkProfile;
        /// <summary>
        /// The workspace unique identifier
        /// </summary>
        public readonly string WorkspaceUID;

        [OutputConstructor]
        private GetWorkspaceResult(
            Outputs.BabylonConfigurationResponse? babylonConfiguration,

            ImmutableDictionary<string, string>? connectivityEndpoints,

            Outputs.DataLakeStorageAccountDetailsResponse? defaultDataLakeStorage,

            Outputs.EncryptionDetailsResponse? encryption,

            ImmutableDictionary<string, object> extraProperties,

            Outputs.ManagedIdentityResponse? identity,

            string location,

            string? managedResourceGroupName,

            string? managedVirtualNetwork,

            Outputs.ManagedVirtualNetworkSettingsResponse? managedVirtualNetworkSettings,

            string name,

            ImmutableArray<Outputs.PrivateEndpointConnectionResponse> privateEndpointConnections,

            string provisioningState,

            string? sqlAdministratorLogin,

            string? sqlAdministratorLoginPassword,

            ImmutableDictionary<string, string>? tags,

            string type,

            Outputs.VirtualNetworkProfileResponse? virtualNetworkProfile,

            string workspaceUID)
        {
            BabylonConfiguration = babylonConfiguration;
            ConnectivityEndpoints = connectivityEndpoints;
            DefaultDataLakeStorage = defaultDataLakeStorage;
            Encryption = encryption;
            ExtraProperties = extraProperties;
            Identity = identity;
            Location = location;
            ManagedResourceGroupName = managedResourceGroupName;
            ManagedVirtualNetwork = managedVirtualNetwork;
            ManagedVirtualNetworkSettings = managedVirtualNetworkSettings;
            Name = name;
            PrivateEndpointConnections = privateEndpointConnections;
            ProvisioningState = provisioningState;
            SqlAdministratorLogin = sqlAdministratorLogin;
            SqlAdministratorLoginPassword = sqlAdministratorLoginPassword;
            Tags = tags;
            Type = type;
            VirtualNetworkProfile = virtualNetworkProfile;
            WorkspaceUID = workspaceUID;
        }
    }
}
