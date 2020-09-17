// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNextGen.Network.Latest
{
    /// <summary>
    /// NetworkVirtualAppliance Resource.
    /// </summary>
    public partial class NetworkVirtualAppliance : Pulumi.CustomResource
    {
        /// <summary>
        /// Address Prefix.
        /// </summary>
        [Output("addressPrefix")]
        public Output<string> AddressPrefix { get; private set; } = null!;

        /// <summary>
        /// BootStrapConfigurationBlobs storage URLs.
        /// </summary>
        [Output("bootStrapConfigurationBlobs")]
        public Output<ImmutableArray<string>> BootStrapConfigurationBlobs { get; private set; } = null!;

        /// <summary>
        /// CloudInitConfiguration string in plain text.
        /// </summary>
        [Output("cloudInitConfiguration")]
        public Output<string?> CloudInitConfiguration { get; private set; } = null!;

        /// <summary>
        /// CloudInitConfigurationBlob storage URLs.
        /// </summary>
        [Output("cloudInitConfigurationBlobs")]
        public Output<ImmutableArray<string>> CloudInitConfigurationBlobs { get; private set; } = null!;

        /// <summary>
        /// A unique read-only string that changes whenever the resource is updated.
        /// </summary>
        [Output("etag")]
        public Output<string> Etag { get; private set; } = null!;

        /// <summary>
        /// The service principal that has read access to cloud-init and config blob.
        /// </summary>
        [Output("identity")]
        public Output<Outputs.ManagedServiceIdentityResponse?> Identity { get; private set; } = null!;

        /// <summary>
        /// List of references to InboundSecurityRules.
        /// </summary>
        [Output("inboundSecurityRules")]
        public Output<ImmutableArray<Outputs.SubResourceResponse>> InboundSecurityRules { get; private set; } = null!;

        /// <summary>
        /// Resource location.
        /// </summary>
        [Output("location")]
        public Output<string?> Location { get; private set; } = null!;

        /// <summary>
        /// Resource name.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// Network Virtual Appliance SKU.
        /// </summary>
        [Output("nvaSku")]
        public Output<Outputs.VirtualApplianceSkuPropertiesResponse?> NvaSku { get; private set; } = null!;

        /// <summary>
        /// The provisioning state of the resource.
        /// </summary>
        [Output("provisioningState")]
        public Output<string> ProvisioningState { get; private set; } = null!;

        /// <summary>
        /// Resource tags.
        /// </summary>
        [Output("tags")]
        public Output<ImmutableDictionary<string, string>?> Tags { get; private set; } = null!;

        /// <summary>
        /// Resource type.
        /// </summary>
        [Output("type")]
        public Output<string> Type { get; private set; } = null!;

        /// <summary>
        /// VirtualAppliance ASN.
        /// </summary>
        [Output("virtualApplianceAsn")]
        public Output<int?> VirtualApplianceAsn { get; private set; } = null!;

        /// <summary>
        /// List of Virtual Appliance Network Interfaces.
        /// </summary>
        [Output("virtualApplianceNics")]
        public Output<ImmutableArray<Outputs.VirtualApplianceNicPropertiesResponse>> VirtualApplianceNics { get; private set; } = null!;

        /// <summary>
        /// List of references to VirtualApplianceSite.
        /// </summary>
        [Output("virtualApplianceSites")]
        public Output<ImmutableArray<Outputs.SubResourceResponse>> VirtualApplianceSites { get; private set; } = null!;

        /// <summary>
        /// The Virtual Hub where Network Virtual Appliance is being deployed.
        /// </summary>
        [Output("virtualHub")]
        public Output<Outputs.SubResourceResponse?> VirtualHub { get; private set; } = null!;


        /// <summary>
        /// Create a NetworkVirtualAppliance resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public NetworkVirtualAppliance(string name, NetworkVirtualApplianceArgs args, CustomResourceOptions? options = null)
            : base("azure-nextgen:network/latest:NetworkVirtualAppliance", name, args ?? new NetworkVirtualApplianceArgs(), MakeResourceOptions(options, ""))
        {
        }

        private NetworkVirtualAppliance(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("azure-nextgen:network/latest:NetworkVirtualAppliance", name, null, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                Aliases =
                {
                    new Pulumi.Alias { Type = "azure-nextgen:network/v20191201:NetworkVirtualAppliance"},
                    new Pulumi.Alias { Type = "azure-nextgen:network/v20200301:NetworkVirtualAppliance"},
                    new Pulumi.Alias { Type = "azure-nextgen:network/v20200401:NetworkVirtualAppliance"},
                    new Pulumi.Alias { Type = "azure-nextgen:network/v20200501:NetworkVirtualAppliance"},
                    new Pulumi.Alias { Type = "azure-nextgen:network/v20200601:NetworkVirtualAppliance"},
                },
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing NetworkVirtualAppliance resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static NetworkVirtualAppliance Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new NetworkVirtualAppliance(name, id, options);
        }
    }

    public sealed class NetworkVirtualApplianceArgs : Pulumi.ResourceArgs
    {
        [Input("bootStrapConfigurationBlobs")]
        private InputList<string>? _bootStrapConfigurationBlobs;

        /// <summary>
        /// BootStrapConfigurationBlobs storage URLs.
        /// </summary>
        public InputList<string> BootStrapConfigurationBlobs
        {
            get => _bootStrapConfigurationBlobs ?? (_bootStrapConfigurationBlobs = new InputList<string>());
            set => _bootStrapConfigurationBlobs = value;
        }

        /// <summary>
        /// CloudInitConfiguration string in plain text.
        /// </summary>
        [Input("cloudInitConfiguration")]
        public Input<string>? CloudInitConfiguration { get; set; }

        [Input("cloudInitConfigurationBlobs")]
        private InputList<string>? _cloudInitConfigurationBlobs;

        /// <summary>
        /// CloudInitConfigurationBlob storage URLs.
        /// </summary>
        public InputList<string> CloudInitConfigurationBlobs
        {
            get => _cloudInitConfigurationBlobs ?? (_cloudInitConfigurationBlobs = new InputList<string>());
            set => _cloudInitConfigurationBlobs = value;
        }

        /// <summary>
        /// Resource ID.
        /// </summary>
        [Input("id")]
        public Input<string>? Id { get; set; }

        /// <summary>
        /// The service principal that has read access to cloud-init and config blob.
        /// </summary>
        [Input("identity")]
        public Input<Inputs.ManagedServiceIdentityArgs>? Identity { get; set; }

        /// <summary>
        /// Resource location.
        /// </summary>
        [Input("location")]
        public Input<string>? Location { get; set; }

        /// <summary>
        /// The name of Network Virtual Appliance.
        /// </summary>
        [Input("networkVirtualApplianceName", required: true)]
        public Input<string> NetworkVirtualApplianceName { get; set; } = null!;

        /// <summary>
        /// Network Virtual Appliance SKU.
        /// </summary>
        [Input("nvaSku")]
        public Input<Inputs.VirtualApplianceSkuPropertiesArgs>? NvaSku { get; set; }

        /// <summary>
        /// The name of the resource group.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public Input<string> ResourceGroupName { get; set; } = null!;

        [Input("tags")]
        private InputMap<string>? _tags;

        /// <summary>
        /// Resource tags.
        /// </summary>
        public InputMap<string> Tags
        {
            get => _tags ?? (_tags = new InputMap<string>());
            set => _tags = value;
        }

        /// <summary>
        /// VirtualAppliance ASN.
        /// </summary>
        [Input("virtualApplianceAsn")]
        public Input<int>? VirtualApplianceAsn { get; set; }

        /// <summary>
        /// The Virtual Hub where Network Virtual Appliance is being deployed.
        /// </summary>
        [Input("virtualHub")]
        public Input<Inputs.SubResourceArgs>? VirtualHub { get; set; }

        public NetworkVirtualApplianceArgs()
        {
        }
    }
}