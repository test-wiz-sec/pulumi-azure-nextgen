// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNextGen.Network.V20170801.Inputs
{

    /// <summary>
    /// SKU of a load balancer
    /// </summary>
    public sealed class LoadBalancerSkuArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Name of a load balancer SKU.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        public LoadBalancerSkuArgs()
        {
        }
    }
}
