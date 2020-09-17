// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNextGen.ManagedNetwork.V20190601Preview.Outputs
{

    [OutputType]
    public sealed class ScopeResponse
    {
        /// <summary>
        /// The collection of management groups covered by the Managed Network
        /// </summary>
        public readonly ImmutableArray<Outputs.ResourceIdResponse> ManagementGroups;
        /// <summary>
        /// The collection of  subnets covered by the Managed Network
        /// </summary>
        public readonly ImmutableArray<Outputs.ResourceIdResponse> Subnets;
        /// <summary>
        /// The collection of subscriptions covered by the Managed Network
        /// </summary>
        public readonly ImmutableArray<Outputs.ResourceIdResponse> Subscriptions;
        /// <summary>
        /// The collection of virtual nets covered by the Managed Network
        /// </summary>
        public readonly ImmutableArray<Outputs.ResourceIdResponse> VirtualNetworks;

        [OutputConstructor]
        private ScopeResponse(
            ImmutableArray<Outputs.ResourceIdResponse> managementGroups,

            ImmutableArray<Outputs.ResourceIdResponse> subnets,

            ImmutableArray<Outputs.ResourceIdResponse> subscriptions,

            ImmutableArray<Outputs.ResourceIdResponse> virtualNetworks)
        {
            ManagementGroups = managementGroups;
            Subnets = subnets;
            Subscriptions = subscriptions;
            VirtualNetworks = virtualNetworks;
        }
    }
}