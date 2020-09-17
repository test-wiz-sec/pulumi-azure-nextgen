// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNextGen.VMwareCloudSimple.Latest.Inputs
{

    /// <summary>
    /// Virtual network model
    /// </summary>
    public sealed class VirtualNetworkArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// virtual network id (privateCloudId:vsphereId)
        /// </summary>
        [Input("id", required: true)]
        public Input<string> Id { get; set; } = null!;

        public VirtualNetworkArgs()
        {
        }
    }
}