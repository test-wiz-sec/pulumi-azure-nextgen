// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNextGen.Kubernetes.V20200101Preview
{
    public static class GetConnectedCluster
    {
        public static Task<GetConnectedClusterResult> InvokeAsync(GetConnectedClusterArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetConnectedClusterResult>("azure-nextgen:kubernetes/v20200101preview:getConnectedCluster", args ?? new GetConnectedClusterArgs(), options.WithVersion());
    }


    public sealed class GetConnectedClusterArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// The name of the Kubernetes cluster on which get is called.
        /// </summary>
        [Input("clusterName", required: true)]
        public string ClusterName { get; set; } = null!;

        /// <summary>
        /// The name of the resource group. The name is case insensitive.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public string ResourceGroupName { get; set; } = null!;

        public GetConnectedClusterArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetConnectedClusterResult
    {
        /// <summary>
        /// AAD profile of the connected cluster.
        /// </summary>
        public readonly Outputs.ConnectedClusterAADProfileResponse AadProfile;
        /// <summary>
        /// Base64 encoded public certificate used by the agent to do the initial handshake to the backend services in Azure.
        /// </summary>
        public readonly string AgentPublicKeyCertificate;
        /// <summary>
        /// Version of the agent running on the connected cluster resource
        /// </summary>
        public readonly string AgentVersion;
        /// <summary>
        /// The identity of the connected cluster.
        /// </summary>
        public readonly Outputs.ConnectedClusterIdentityResponse Identity;
        /// <summary>
        /// The Kubernetes version of the connected cluster resource
        /// </summary>
        public readonly string KubernetesVersion;
        /// <summary>
        /// The geo-location where the resource lives
        /// </summary>
        public readonly string Location;
        /// <summary>
        /// The name of the resource
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// Provisioning state of the connected cluster resource.
        /// </summary>
        public readonly string? ProvisioningState;
        /// <summary>
        /// Resource tags.
        /// </summary>
        public readonly ImmutableDictionary<string, string>? Tags;
        /// <summary>
        /// Number of nodes present in the connected cluster resource
        /// </summary>
        public readonly int TotalNodeCount;
        /// <summary>
        /// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
        /// </summary>
        public readonly string Type;

        [OutputConstructor]
        private GetConnectedClusterResult(
            Outputs.ConnectedClusterAADProfileResponse aadProfile,

            string agentPublicKeyCertificate,

            string agentVersion,

            Outputs.ConnectedClusterIdentityResponse identity,

            string kubernetesVersion,

            string location,

            string name,

            string? provisioningState,

            ImmutableDictionary<string, string>? tags,

            int totalNodeCount,

            string type)
        {
            AadProfile = aadProfile;
            AgentPublicKeyCertificate = agentPublicKeyCertificate;
            AgentVersion = agentVersion;
            Identity = identity;
            KubernetesVersion = kubernetesVersion;
            Location = location;
            Name = name;
            ProvisioningState = provisioningState;
            Tags = tags;
            TotalNodeCount = totalNodeCount;
            Type = type;
        }
    }
}
