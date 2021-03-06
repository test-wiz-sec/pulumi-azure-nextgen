// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNextGen.AVS.V20190809Preview.Inputs
{

    /// <summary>
    /// The properties of a default cluster
    /// </summary>
    public sealed class DefaultClusterPropertiesArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The cluster size
        /// </summary>
        [Input("clusterSize")]
        public Input<int>? ClusterSize { get; set; }

        public DefaultClusterPropertiesArgs()
        {
        }
    }
}
