// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNextGen.Network.V20190801.Outputs
{

    [OutputType]
    public sealed class HubVirtualNetworkConnectionResponse
    {
        /// <summary>
        /// VirtualHub to RemoteVnet transit to enabled or not.
        /// </summary>
        public readonly bool? AllowHubToRemoteVnetTransit;
        /// <summary>
        /// Allow RemoteVnet to use Virtual Hub's gateways.
        /// </summary>
        public readonly bool? AllowRemoteVnetToUseHubVnetGateways;
        /// <summary>
        /// Enable internet security.
        /// </summary>
        public readonly bool? EnableInternetSecurity;
        /// <summary>
        /// A unique read-only string that changes whenever the resource is updated.
        /// </summary>
        public readonly string Etag;
        /// <summary>
        /// Resource ID.
        /// </summary>
        public readonly string? Id;
        /// <summary>
        /// The name of the resource that is unique within a resource group. This name can be used to access the resource.
        /// </summary>
        public readonly string? Name;
        /// <summary>
        /// The provisioning state of the hub virtual network connection resource.
        /// </summary>
        public readonly string? ProvisioningState;
        /// <summary>
        /// Reference to the remote virtual network.
        /// </summary>
        public readonly Outputs.SubResourceResponse? RemoteVirtualNetwork;

        [OutputConstructor]
        private HubVirtualNetworkConnectionResponse(
            bool? allowHubToRemoteVnetTransit,

            bool? allowRemoteVnetToUseHubVnetGateways,

            bool? enableInternetSecurity,

            string etag,

            string? id,

            string? name,

            string? provisioningState,

            Outputs.SubResourceResponse? remoteVirtualNetwork)
        {
            AllowHubToRemoteVnetTransit = allowHubToRemoteVnetTransit;
            AllowRemoteVnetToUseHubVnetGateways = allowRemoteVnetToUseHubVnetGateways;
            EnableInternetSecurity = enableInternetSecurity;
            Etag = etag;
            Id = id;
            Name = name;
            ProvisioningState = provisioningState;
            RemoteVirtualNetwork = remoteVirtualNetwork;
        }
    }
}