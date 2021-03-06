// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNextGen.Network.V20200401
{
    public static class GetNetworkVirtualAppliance
    {
        public static Task<GetNetworkVirtualApplianceResult> InvokeAsync(GetNetworkVirtualApplianceArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetNetworkVirtualApplianceResult>("azure-nextgen:network/v20200401:getNetworkVirtualAppliance", args ?? new GetNetworkVirtualApplianceArgs(), options.WithVersion());
    }


    public sealed class GetNetworkVirtualApplianceArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// Expands referenced resources.
        /// </summary>
        [Input("expand")]
        public string? Expand { get; set; }

        /// <summary>
        /// The name of Network Virtual Appliance.
        /// </summary>
        [Input("networkVirtualApplianceName", required: true)]
        public string NetworkVirtualApplianceName { get; set; } = null!;

        /// <summary>
        /// The name of the resource group.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public string ResourceGroupName { get; set; } = null!;

        public GetNetworkVirtualApplianceArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetNetworkVirtualApplianceResult
    {
        /// <summary>
        /// BootStrapConfigurationBlob storage URLs.
        /// </summary>
        public readonly ImmutableArray<string> BootStrapConfigurationBlob;
        /// <summary>
        /// CloudInitConfigurationBlob storage URLs.
        /// </summary>
        public readonly ImmutableArray<string> CloudInitConfigurationBlob;
        /// <summary>
        /// A unique read-only string that changes whenever the resource is updated.
        /// </summary>
        public readonly string Etag;
        /// <summary>
        /// The service principal that has read access to cloud-init and config blob.
        /// </summary>
        public readonly Outputs.ManagedServiceIdentityResponse? Identity;
        /// <summary>
        /// Resource location.
        /// </summary>
        public readonly string? Location;
        /// <summary>
        /// Resource name.
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// The provisioning state of the resource.
        /// </summary>
        public readonly string ProvisioningState;
        /// <summary>
        /// Network Virtual Appliance SKU.
        /// </summary>
        public readonly Outputs.VirtualApplianceSkuPropertiesResponse? Sku;
        /// <summary>
        /// Resource tags.
        /// </summary>
        public readonly ImmutableDictionary<string, string>? Tags;
        /// <summary>
        /// Resource type.
        /// </summary>
        public readonly string Type;
        /// <summary>
        /// VirtualAppliance ASN.
        /// </summary>
        public readonly int? VirtualApplianceAsn;
        /// <summary>
        /// List of Virtual Appliance Network Interfaces.
        /// </summary>
        public readonly ImmutableArray<Outputs.VirtualApplianceNicPropertiesResponse> VirtualApplianceNics;
        /// <summary>
        /// The Virtual Hub where Network Virtual Appliance is being deployed.
        /// </summary>
        public readonly Outputs.SubResourceResponse? VirtualHub;

        [OutputConstructor]
        private GetNetworkVirtualApplianceResult(
            ImmutableArray<string> bootStrapConfigurationBlob,

            ImmutableArray<string> cloudInitConfigurationBlob,

            string etag,

            Outputs.ManagedServiceIdentityResponse? identity,

            string? location,

            string name,

            string provisioningState,

            Outputs.VirtualApplianceSkuPropertiesResponse? sku,

            ImmutableDictionary<string, string>? tags,

            string type,

            int? virtualApplianceAsn,

            ImmutableArray<Outputs.VirtualApplianceNicPropertiesResponse> virtualApplianceNics,

            Outputs.SubResourceResponse? virtualHub)
        {
            BootStrapConfigurationBlob = bootStrapConfigurationBlob;
            CloudInitConfigurationBlob = cloudInitConfigurationBlob;
            Etag = etag;
            Identity = identity;
            Location = location;
            Name = name;
            ProvisioningState = provisioningState;
            Sku = sku;
            Tags = tags;
            Type = type;
            VirtualApplianceAsn = virtualApplianceAsn;
            VirtualApplianceNics = virtualApplianceNics;
            VirtualHub = virtualHub;
        }
    }
}
