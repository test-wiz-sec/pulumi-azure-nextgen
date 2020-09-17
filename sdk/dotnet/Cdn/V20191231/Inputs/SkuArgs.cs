// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNextGen.Cdn.V20191231.Inputs
{

    /// <summary>
    /// The pricing tier (defines a CDN provider, feature list and rate) of the CDN profile.
    /// </summary>
    public sealed class SkuArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Name of the pricing tier.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        public SkuArgs()
        {
        }
    }
}