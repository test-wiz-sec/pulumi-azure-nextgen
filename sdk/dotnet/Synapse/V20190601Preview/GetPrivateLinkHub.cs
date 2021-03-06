// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNextGen.Synapse.V20190601Preview
{
    public static class GetPrivateLinkHub
    {
        public static Task<GetPrivateLinkHubResult> InvokeAsync(GetPrivateLinkHubArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetPrivateLinkHubResult>("azure-nextgen:synapse/v20190601preview:getPrivateLinkHub", args ?? new GetPrivateLinkHubArgs(), options.WithVersion());
    }


    public sealed class GetPrivateLinkHubArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// Name of the privateLinkHub
        /// </summary>
        [Input("privateLinkHubName", required: true)]
        public string PrivateLinkHubName { get; set; } = null!;

        /// <summary>
        /// The name of the resource group. The name is case insensitive.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public string ResourceGroupName { get; set; } = null!;

        public GetPrivateLinkHubArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetPrivateLinkHubResult
    {
        /// <summary>
        /// The geo-location where the resource lives
        /// </summary>
        public readonly string Location;
        /// <summary>
        /// The name of the resource
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// List of private endpoint connections
        /// </summary>
        public readonly ImmutableArray<Outputs.PrivateEndpointConnectionForPrivateLinkHubBasicResponse> PrivateEndpointConnections;
        /// <summary>
        /// PrivateLinkHub provisioning state
        /// </summary>
        public readonly string? ProvisioningState;
        /// <summary>
        /// Resource tags.
        /// </summary>
        public readonly ImmutableDictionary<string, string>? Tags;
        /// <summary>
        /// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
        /// </summary>
        public readonly string Type;

        [OutputConstructor]
        private GetPrivateLinkHubResult(
            string location,

            string name,

            ImmutableArray<Outputs.PrivateEndpointConnectionForPrivateLinkHubBasicResponse> privateEndpointConnections,

            string? provisioningState,

            ImmutableDictionary<string, string>? tags,

            string type)
        {
            Location = location;
            Name = name;
            PrivateEndpointConnections = privateEndpointConnections;
            ProvisioningState = provisioningState;
            Tags = tags;
            Type = type;
        }
    }
}
