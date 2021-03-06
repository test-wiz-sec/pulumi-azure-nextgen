// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNextGen.Network.Latest
{
    public static class GetNetworkExperimentProfile
    {
        public static Task<GetNetworkExperimentProfileResult> InvokeAsync(GetNetworkExperimentProfileArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetNetworkExperimentProfileResult>("azure-nextgen:network/latest:getNetworkExperimentProfile", args ?? new GetNetworkExperimentProfileArgs(), options.WithVersion());
    }


    public sealed class GetNetworkExperimentProfileArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// The Profile identifier associated with the Tenant and Partner
        /// </summary>
        [Input("profileName", required: true)]
        public string ProfileName { get; set; } = null!;

        /// <summary>
        /// Name of the Resource group within the Azure subscription.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public string ResourceGroupName { get; set; } = null!;

        public GetNetworkExperimentProfileArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetNetworkExperimentProfileResult
    {
        /// <summary>
        /// The state of the Experiment
        /// </summary>
        public readonly string? EnabledState;
        /// <summary>
        /// Gets a unique read-only string that changes whenever the resource is updated.
        /// </summary>
        public readonly string? Etag;
        /// <summary>
        /// Resource location.
        /// </summary>
        public readonly string? Location;
        /// <summary>
        /// Resource name.
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// Resource status.
        /// </summary>
        public readonly string ResourceState;
        /// <summary>
        /// Resource tags.
        /// </summary>
        public readonly ImmutableDictionary<string, string>? Tags;
        /// <summary>
        /// Resource type.
        /// </summary>
        public readonly string Type;

        [OutputConstructor]
        private GetNetworkExperimentProfileResult(
            string? enabledState,

            string? etag,

            string? location,

            string name,

            string resourceState,

            ImmutableDictionary<string, string>? tags,

            string type)
        {
            EnabledState = enabledState;
            Etag = etag;
            Location = location;
            Name = name;
            ResourceState = resourceState;
            Tags = tags;
            Type = type;
        }
    }
}
