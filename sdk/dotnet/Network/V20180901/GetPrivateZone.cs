// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNextGen.Network.V20180901
{
    public static class GetPrivateZone
    {
        public static Task<GetPrivateZoneResult> InvokeAsync(GetPrivateZoneArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetPrivateZoneResult>("azure-nextgen:network/v20180901:getPrivateZone", args ?? new GetPrivateZoneArgs(), options.WithVersion());
    }


    public sealed class GetPrivateZoneArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// The name of the Private DNS zone (without a terminating dot).
        /// </summary>
        [Input("privateZoneName", required: true)]
        public string PrivateZoneName { get; set; } = null!;

        /// <summary>
        /// The name of the resource group.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public string ResourceGroupName { get; set; } = null!;

        public GetPrivateZoneArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetPrivateZoneResult
    {
        /// <summary>
        /// The ETag of the zone.
        /// </summary>
        public readonly string? Etag;
        /// <summary>
        /// The Azure Region where the resource lives
        /// </summary>
        public readonly string? Location;
        /// <summary>
        /// The maximum number of record sets that can be created in this Private DNS zone. This is a read-only property and any attempt to set this value will be ignored.
        /// </summary>
        public readonly int MaxNumberOfRecordSets;
        /// <summary>
        /// The maximum number of virtual networks that can be linked to this Private DNS zone. This is a read-only property and any attempt to set this value will be ignored.
        /// </summary>
        public readonly int MaxNumberOfVirtualNetworkLinks;
        /// <summary>
        /// The maximum number of virtual networks that can be linked to this Private DNS zone with registration enabled. This is a read-only property and any attempt to set this value will be ignored.
        /// </summary>
        public readonly int MaxNumberOfVirtualNetworkLinksWithRegistration;
        /// <summary>
        /// The name of the resource
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// The current number of record sets in this Private DNS zone. This is a read-only property and any attempt to set this value will be ignored.
        /// </summary>
        public readonly int NumberOfRecordSets;
        /// <summary>
        /// The current number of virtual networks that are linked to this Private DNS zone. This is a read-only property and any attempt to set this value will be ignored.
        /// </summary>
        public readonly int NumberOfVirtualNetworkLinks;
        /// <summary>
        /// The current number of virtual networks that are linked to this Private DNS zone with registration enabled. This is a read-only property and any attempt to set this value will be ignored.
        /// </summary>
        public readonly int NumberOfVirtualNetworkLinksWithRegistration;
        /// <summary>
        /// The provisioning state of the resource. This is a read-only property and any attempt to set this value will be ignored.
        /// </summary>
        public readonly string ProvisioningState;
        /// <summary>
        /// Resource tags.
        /// </summary>
        public readonly ImmutableDictionary<string, string>? Tags;
        /// <summary>
        /// The type of the resource. Example - 'Microsoft.Network/privateDnsZones'.
        /// </summary>
        public readonly string Type;

        [OutputConstructor]
        private GetPrivateZoneResult(
            string? etag,

            string? location,

            int maxNumberOfRecordSets,

            int maxNumberOfVirtualNetworkLinks,

            int maxNumberOfVirtualNetworkLinksWithRegistration,

            string name,

            int numberOfRecordSets,

            int numberOfVirtualNetworkLinks,

            int numberOfVirtualNetworkLinksWithRegistration,

            string provisioningState,

            ImmutableDictionary<string, string>? tags,

            string type)
        {
            Etag = etag;
            Location = location;
            MaxNumberOfRecordSets = maxNumberOfRecordSets;
            MaxNumberOfVirtualNetworkLinks = maxNumberOfVirtualNetworkLinks;
            MaxNumberOfVirtualNetworkLinksWithRegistration = maxNumberOfVirtualNetworkLinksWithRegistration;
            Name = name;
            NumberOfRecordSets = numberOfRecordSets;
            NumberOfVirtualNetworkLinks = numberOfVirtualNetworkLinks;
            NumberOfVirtualNetworkLinksWithRegistration = numberOfVirtualNetworkLinksWithRegistration;
            ProvisioningState = provisioningState;
            Tags = tags;
            Type = type;
        }
    }
}
