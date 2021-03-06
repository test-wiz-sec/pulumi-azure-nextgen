// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNextGen.DBforMySQL.V20200701PrivatePreview
{
    public static class GetServer
    {
        public static Task<GetServerResult> InvokeAsync(GetServerArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetServerResult>("azure-nextgen:dbformysql/v20200701privatepreview:getServer", args ?? new GetServerArgs(), options.WithVersion());
    }


    public sealed class GetServerArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// The name of the resource group. The name is case insensitive.
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
        /// The password of the administrator login (required for server creation).
        /// </summary>
        public readonly string? AdministratorLoginPassword;
        /// <summary>
        /// availability Zone information of the server.
        /// </summary>
        public readonly string? AvailabilityZone;
        /// <summary>
        /// Status showing whether the data encryption is enabled with customer-managed keys.
        /// </summary>
        public readonly string ByokEnforcement;
        /// <summary>
        /// The mode to create a new MySQL server.
        /// </summary>
        public readonly string? CreateMode;
        /// <summary>
        /// Delegated subnet arguments.
        /// </summary>
        public readonly Outputs.DelegatedSubnetArgumentsResponse? DelegatedSubnetArguments;
        /// <summary>
        /// Earliest restore point creation time (ISO8601 format)
        /// </summary>
        public readonly string EarliestRestoreDate;
        /// <summary>
        /// The fully qualified domain name of a server.
        /// </summary>
        public readonly string FullyQualifiedDomainName;
        /// <summary>
        /// Enable HA or not for a server.
        /// </summary>
        public readonly string? HaEnabled;
        /// <summary>
        /// The state of a HA server.
        /// </summary>
        public readonly string HaState;
        /// <summary>
        /// The Azure Active Directory identity of the server.
        /// </summary>
        public readonly Outputs.IdentityResponse? Identity;
        /// <summary>
        /// Status showing whether the server enabled infrastructure encryption.
        /// </summary>
        public readonly string? InfrastructureEncryption;
        /// <summary>
        /// The geo-location where the resource lives
        /// </summary>
        public readonly string Location;
        /// <summary>
        /// Maintenance window of a server.
        /// </summary>
        public readonly Outputs.MaintenanceWindowResponse? MaintenanceWindow;
        /// <summary>
        /// The name of the resource
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// Whether or not public network access is allowed for this server. Value is optional but if passed in, must be 'Enabled' or 'Disabled'
        /// </summary>
        public readonly string PublicNetworkAccess;
        /// <summary>
        /// The maximum number of replicas that a primary server can have.
        /// </summary>
        public readonly int ReplicaCapacity;
        /// <summary>
        /// The replication role.
        /// </summary>
        public readonly string? ReplicationRole;
        /// <summary>
        /// Restore point creation time (ISO8601 format), specifying the time to restore from.
        /// </summary>
        public readonly string? RestorePointInTime;
        /// <summary>
        /// The SKU (pricing tier) of the server.
        /// </summary>
        public readonly Outputs.SkuResponse? Sku;
        /// <summary>
        /// The source MySQL server id.
        /// </summary>
        public readonly string? SourceServerId;
        /// <summary>
        /// Enable ssl enforcement or not when connect to server.
        /// </summary>
        public readonly string? SslEnforcement;
        /// <summary>
        /// availability Zone information of the server.
        /// </summary>
        public readonly string StandbyAvailabilityZone;
        /// <summary>
        /// The state of a server.
        /// </summary>
        public readonly string State;
        /// <summary>
        /// Storage profile of a server.
        /// </summary>
        public readonly Outputs.StorageProfileResponse? StorageProfile;
        /// <summary>
        /// Resource tags.
        /// </summary>
        public readonly ImmutableDictionary<string, string>? Tags;
        /// <summary>
        /// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
        /// </summary>
        public readonly string Type;
        /// <summary>
        /// Server version.
        /// </summary>
        public readonly string? Version;

        [OutputConstructor]
        private GetServerResult(
            string? administratorLogin,

            string? administratorLoginPassword,

            string? availabilityZone,

            string byokEnforcement,

            string? createMode,

            Outputs.DelegatedSubnetArgumentsResponse? delegatedSubnetArguments,

            string earliestRestoreDate,

            string fullyQualifiedDomainName,

            string? haEnabled,

            string haState,

            Outputs.IdentityResponse? identity,

            string? infrastructureEncryption,

            string location,

            Outputs.MaintenanceWindowResponse? maintenanceWindow,

            string name,

            string publicNetworkAccess,

            int replicaCapacity,

            string? replicationRole,

            string? restorePointInTime,

            Outputs.SkuResponse? sku,

            string? sourceServerId,

            string? sslEnforcement,

            string standbyAvailabilityZone,

            string state,

            Outputs.StorageProfileResponse? storageProfile,

            ImmutableDictionary<string, string>? tags,

            string type,

            string? version)
        {
            AdministratorLogin = administratorLogin;
            AdministratorLoginPassword = administratorLoginPassword;
            AvailabilityZone = availabilityZone;
            ByokEnforcement = byokEnforcement;
            CreateMode = createMode;
            DelegatedSubnetArguments = delegatedSubnetArguments;
            EarliestRestoreDate = earliestRestoreDate;
            FullyQualifiedDomainName = fullyQualifiedDomainName;
            HaEnabled = haEnabled;
            HaState = haState;
            Identity = identity;
            InfrastructureEncryption = infrastructureEncryption;
            Location = location;
            MaintenanceWindow = maintenanceWindow;
            Name = name;
            PublicNetworkAccess = publicNetworkAccess;
            ReplicaCapacity = replicaCapacity;
            ReplicationRole = replicationRole;
            RestorePointInTime = restorePointInTime;
            Sku = sku;
            SourceServerId = sourceServerId;
            SslEnforcement = sslEnforcement;
            StandbyAvailabilityZone = standbyAvailabilityZone;
            State = state;
            StorageProfile = storageProfile;
            Tags = tags;
            Type = type;
            Version = version;
        }
    }
}
