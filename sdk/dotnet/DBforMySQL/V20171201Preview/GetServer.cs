// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNextGen.DBforMySQL.V20171201Preview
{
    public static class GetServer
    {
        public static Task<GetServerResult> InvokeAsync(GetServerArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetServerResult>("azure-nextgen:dbformysql/v20171201preview:getServer", args ?? new GetServerArgs(), options.WithVersion());
    }


    public sealed class GetServerArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// The name of the resource group that contains the resource. You can obtain this value from the Azure Resource Manager API or the portal.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public string ResourceGroupName { get; set; } = null!;

        /// <summary>
        /// The name of the server.
        /// </summary>
        [Input("serverName", required: true)]
        public string ServerName { get; set; } = null!;

        public GetServerArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetServerResult
    {
        /// <summary>
        /// The administrator's login name of a server. Can only be specified when the server is being created (and is required for creation).
        /// </summary>
        public readonly string? AdministratorLogin;
        /// <summary>
        /// Status showing whether the server data encryption is enabled with customer-managed keys.
        /// </summary>
        public readonly string ByokEnforcement;
        /// <summary>
        /// Earliest restore point creation time (ISO8601 format)
        /// </summary>
        public readonly string? EarliestRestoreDate;
        /// <summary>
        /// The fully qualified domain name of a server.
        /// </summary>
        public readonly string? FullyQualifiedDomainName;
        /// <summary>
        /// The Azure Active Directory identity of the server.
        /// </summary>
        public readonly Outputs.ResourceIdentityResponse? Identity;
        /// <summary>
        /// Status showing whether the server enabled infrastructure encryption.
        /// </summary>
        public readonly string? InfrastructureEncryption;
        /// <summary>
        /// The geo-location where the resource lives
        /// </summary>
        public readonly string Location;
        /// <summary>
        /// The master server id of a replica server.
        /// </summary>
        public readonly string? MasterServerId;
        /// <summary>
        /// Enforce a minimal Tls version for the server.
        /// </summary>
        public readonly string? MinimalTlsVersion;
        /// <summary>
        /// The name of the resource
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// List of private endpoint connections on a server
        /// </summary>
        public readonly ImmutableArray<Outputs.ServerPrivateEndpointConnectionResponse> PrivateEndpointConnections;
        /// <summary>
        /// Whether or not public endpoint access is allowed for this server.  Value is optional but if passed in, must be 'Enabled' or 'Disabled'
        /// </summary>
        public readonly string? PublicNetworkAccess;
        /// <summary>
        /// The maximum number of replicas that a master server can have.
        /// </summary>
        public readonly int? ReplicaCapacity;
        /// <summary>
        /// The replication role of the server.
        /// </summary>
        public readonly string? ReplicationRole;
        /// <summary>
        /// The SKU (pricing tier) of the server.
        /// </summary>
        public readonly Outputs.SkuResponse? Sku;
        /// <summary>
        /// Enable ssl enforcement or not when connect to server.
        /// </summary>
        public readonly string? SslEnforcement;
        /// <summary>
        /// Storage profile of a server.
        /// </summary>
        public readonly Outputs.StorageProfileResponse? StorageProfile;
        /// <summary>
        /// Resource tags.
        /// </summary>
        public readonly ImmutableDictionary<string, string>? Tags;
        /// <summary>
        /// The type of the resource. Ex- Microsoft.Compute/virtualMachines or Microsoft.Storage/storageAccounts.
        /// </summary>
        public readonly string Type;
        /// <summary>
        /// A state of a server that is visible to user.
        /// </summary>
        public readonly string? UserVisibleState;
        /// <summary>
        /// Server version.
        /// </summary>
        public readonly string? Version;

        [OutputConstructor]
        private GetServerResult(
            string? administratorLogin,

            string byokEnforcement,

            string? earliestRestoreDate,

            string? fullyQualifiedDomainName,

            Outputs.ResourceIdentityResponse? identity,

            string? infrastructureEncryption,

            string location,

            string? masterServerId,

            string? minimalTlsVersion,

            string name,

            ImmutableArray<Outputs.ServerPrivateEndpointConnectionResponse> privateEndpointConnections,

            string? publicNetworkAccess,

            int? replicaCapacity,

            string? replicationRole,

            Outputs.SkuResponse? sku,

            string? sslEnforcement,

            Outputs.StorageProfileResponse? storageProfile,

            ImmutableDictionary<string, string>? tags,

            string type,

            string? userVisibleState,

            string? version)
        {
            AdministratorLogin = administratorLogin;
            ByokEnforcement = byokEnforcement;
            EarliestRestoreDate = earliestRestoreDate;
            FullyQualifiedDomainName = fullyQualifiedDomainName;
            Identity = identity;
            InfrastructureEncryption = infrastructureEncryption;
            Location = location;
            MasterServerId = masterServerId;
            MinimalTlsVersion = minimalTlsVersion;
            Name = name;
            PrivateEndpointConnections = privateEndpointConnections;
            PublicNetworkAccess = publicNetworkAccess;
            ReplicaCapacity = replicaCapacity;
            ReplicationRole = replicationRole;
            Sku = sku;
            SslEnforcement = sslEnforcement;
            StorageProfile = storageProfile;
            Tags = tags;
            Type = type;
            UserVisibleState = userVisibleState;
            Version = version;
        }
    }
}
