// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNextGen.AppPlatform.V20200701.Inputs
{

    /// <summary>
    /// Service properties payload
    /// </summary>
    public sealed class ClusterResourcePropertiesArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Network profile of the Service
        /// </summary>
        [Input("networkProfile")]
        public Input<Inputs.NetworkProfileArgs>? NetworkProfile { get; set; }

        public ClusterResourcePropertiesArgs()
        {
        }
    }
}