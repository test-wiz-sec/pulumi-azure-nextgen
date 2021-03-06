// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNextGen.ContainerRegistry.V20191201Preview
{
    public static class GetRegistry
    {
        public static Task<GetRegistryResult> InvokeAsync(GetRegistryArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetRegistryResult>("azure-nextgen:containerregistry/v20191201preview:getRegistry", args ?? new GetRegistryArgs(), options.WithVersion());
    }


    public sealed class GetRegistryArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// The name of the container registry.
        /// </summary>
        [Input("registryName", required: true)]
        public string RegistryName { get; set; } = null!;

        /// <summary>
        /// The name of the resource group to which the container registry belongs.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public string ResourceGroupName { get; set; } = null!;

        public GetRegistryArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetRegistryResult
    {
        /// <summary>
        /// The value that indicates whether the admin user is enabled.
        /// </summary>
        public readonly bool? AdminUserEnabled;
        /// <summary>
        /// The creation date of the container registry in ISO8601 format.
        /// </summary>
        public readonly string CreationDate;
        /// <summary>
        /// Enable a single data endpoint per region for serving data.
        /// </summary>
        public readonly bool? DataEndpointEnabled;
        /// <summary>
        /// List of host names that will serve data when dataEndpointEnabled is true.
        /// </summary>
        public readonly ImmutableArray<string> DataEndpointHostNames;
        /// <summary>
        /// The encryption settings of container registry.
        /// </summary>
        public readonly Outputs.EncryptionPropertyResponse? Encryption;
        /// <summary>
        /// The identity of the container registry.
        /// </summary>
        public readonly Outputs.IdentityPropertiesResponse? Identity;
        /// <summary>
        /// The location of the resource. This cannot be changed after the resource is created.
        /// </summary>
        public readonly string Location;
        /// <summary>
        /// The URL that can be used to log into the container registry.
        /// </summary>
        public readonly string LoginServer;
        /// <summary>
        /// The name of the resource.
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// The network rule set for a container registry.
        /// </summary>
        public readonly Outputs.NetworkRuleSetResponse? NetworkRuleSet;
        /// <summary>
        /// The policies for a container registry.
        /// </summary>
        public readonly Outputs.PoliciesResponse? Policies;
        /// <summary>
        /// List of private endpoint connections for a container registry.
        /// </summary>
        public readonly ImmutableArray<Outputs.PrivateEndpointConnectionResponse> PrivateEndpointConnections;
        /// <summary>
        /// The provisioning state of the container registry at the time the operation was called.
        /// </summary>
        public readonly string ProvisioningState;
        /// <summary>
        /// Whether or not public network access is allowed for the container registry.
        /// </summary>
        public readonly string? PublicNetworkAccess;
        /// <summary>
        /// The SKU of the container registry.
        /// </summary>
        public readonly Outputs.SkuResponse Sku;
        /// <summary>
        /// The status of the container registry at the time the operation was called.
        /// </summary>
        public readonly Outputs.StatusResponse Status;
        /// <summary>
        /// The properties of the storage account for the container registry. Only applicable to Classic SKU.
        /// </summary>
        public readonly Outputs.StorageAccountPropertiesResponse? StorageAccount;
        /// <summary>
        /// The tags of the resource.
        /// </summary>
        public readonly ImmutableDictionary<string, string>? Tags;
        /// <summary>
        /// The type of the resource.
        /// </summary>
        public readonly string Type;

        [OutputConstructor]
        private GetRegistryResult(
            bool? adminUserEnabled,

            string creationDate,

            bool? dataEndpointEnabled,

            ImmutableArray<string> dataEndpointHostNames,

            Outputs.EncryptionPropertyResponse? encryption,

            Outputs.IdentityPropertiesResponse? identity,

            string location,

            string loginServer,

            string name,

            Outputs.NetworkRuleSetResponse? networkRuleSet,

            Outputs.PoliciesResponse? policies,

            ImmutableArray<Outputs.PrivateEndpointConnectionResponse> privateEndpointConnections,

            string provisioningState,

            string? publicNetworkAccess,

            Outputs.SkuResponse sku,

            Outputs.StatusResponse status,

            Outputs.StorageAccountPropertiesResponse? storageAccount,

            ImmutableDictionary<string, string>? tags,

            string type)
        {
            AdminUserEnabled = adminUserEnabled;
            CreationDate = creationDate;
            DataEndpointEnabled = dataEndpointEnabled;
            DataEndpointHostNames = dataEndpointHostNames;
            Encryption = encryption;
            Identity = identity;
            Location = location;
            LoginServer = loginServer;
            Name = name;
            NetworkRuleSet = networkRuleSet;
            Policies = policies;
            PrivateEndpointConnections = privateEndpointConnections;
            ProvisioningState = provisioningState;
            PublicNetworkAccess = publicNetworkAccess;
            Sku = sku;
            Status = status;
            StorageAccount = storageAccount;
            Tags = tags;
            Type = type;
        }
    }
}
