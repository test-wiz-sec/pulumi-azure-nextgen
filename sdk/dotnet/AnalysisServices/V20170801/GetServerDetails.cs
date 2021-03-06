// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNextGen.AnalysisServices.V20170801
{
    public static class GetServerDetails
    {
        public static Task<GetServerDetailsResult> InvokeAsync(GetServerDetailsArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetServerDetailsResult>("azure-nextgen:analysisservices/v20170801:getServerDetails", args ?? new GetServerDetailsArgs(), options.WithVersion());
    }


    public sealed class GetServerDetailsArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// The name of the Azure Resource group of which a given Analysis Services server is part. This name must be at least 1 character in length, and no more than 90.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public string ResourceGroupName { get; set; } = null!;

        /// <summary>
        /// The name of the Analysis Services server. It must be a minimum of 3 characters, and a maximum of 63.
        /// </summary>
        [Input("serverName", required: true)]
        public string ServerName { get; set; } = null!;

        public GetServerDetailsArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetServerDetailsResult
    {
        /// <summary>
        /// A collection of AS server administrators
        /// </summary>
        public readonly Outputs.ServerAdministratorsResponse? AsAdministrators;
        /// <summary>
        /// The SAS container URI to the backup container.
        /// </summary>
        public readonly string? BackupBlobContainerUri;
        /// <summary>
        /// The gateway details configured for the AS server.
        /// </summary>
        public readonly Outputs.GatewayDetailsResponse? GatewayDetails;
        /// <summary>
        /// The firewall settings for the AS server.
        /// </summary>
        public readonly Outputs.IPv4FirewallSettingsResponse? IpV4FirewallSettings;
        /// <summary>
        /// Location of the Analysis Services resource.
        /// </summary>
        public readonly string Location;
        /// <summary>
        /// The name of the Analysis Services resource.
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// The current deployment state of Analysis Services resource. The provisioningState is to indicate states for resource provisioning.
        /// </summary>
        public readonly string ProvisioningState;
        /// <summary>
        /// How the read-write server's participation in the query pool is controlled.&lt;br/&gt;It can have the following values: &lt;ul&gt;&lt;li&gt;readOnly - indicates that the read-write server is intended not to participate in query operations&lt;/li&gt;&lt;li&gt;all - indicates that the read-write server can participate in query operations&lt;/li&gt;&lt;/ul&gt;Specifying readOnly when capacity is 1 results in error.
        /// </summary>
        public readonly string? QuerypoolConnectionMode;
        /// <summary>
        /// The full name of the Analysis Services resource.
        /// </summary>
        public readonly string ServerFullName;
        /// <summary>
        /// The SKU of the Analysis Services resource.
        /// </summary>
        public readonly Outputs.ResourceSkuResponse Sku;
        /// <summary>
        /// The current state of Analysis Services resource. The state is to indicate more states outside of resource provisioning.
        /// </summary>
        public readonly string State;
        /// <summary>
        /// Key-value pairs of additional resource provisioning properties.
        /// </summary>
        public readonly ImmutableDictionary<string, string>? Tags;
        /// <summary>
        /// The type of the Analysis Services resource.
        /// </summary>
        public readonly string Type;

        [OutputConstructor]
        private GetServerDetailsResult(
            Outputs.ServerAdministratorsResponse? asAdministrators,

            string? backupBlobContainerUri,

            Outputs.GatewayDetailsResponse? gatewayDetails,

            Outputs.IPv4FirewallSettingsResponse? ipV4FirewallSettings,

            string location,

            string name,

            string provisioningState,

            string? querypoolConnectionMode,

            string serverFullName,

            Outputs.ResourceSkuResponse sku,

            string state,

            ImmutableDictionary<string, string>? tags,

            string type)
        {
            AsAdministrators = asAdministrators;
            BackupBlobContainerUri = backupBlobContainerUri;
            GatewayDetails = gatewayDetails;
            IpV4FirewallSettings = ipV4FirewallSettings;
            Location = location;
            Name = name;
            ProvisioningState = provisioningState;
            QuerypoolConnectionMode = querypoolConnectionMode;
            ServerFullName = serverFullName;
            Sku = sku;
            State = state;
            Tags = tags;
            Type = type;
        }
    }
}
