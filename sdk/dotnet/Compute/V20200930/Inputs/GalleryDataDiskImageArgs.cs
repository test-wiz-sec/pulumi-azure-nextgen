// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNextGen.Compute.V20200930.Inputs
{

    /// <summary>
    /// This is the data disk image.
    /// </summary>
    public sealed class GalleryDataDiskImageArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The host caching of the disk. Valid values are 'None', 'ReadOnly', and 'ReadWrite'
        /// </summary>
        [Input("hostCaching")]
        public Input<string>? HostCaching { get; set; }

        /// <summary>
        /// This property specifies the logical unit number of the data disk. This value is used to identify data disks within the Virtual Machine and therefore must be unique for each data disk attached to the Virtual Machine.
        /// </summary>
        [Input("lun", required: true)]
        public Input<int> Lun { get; set; } = null!;

        /// <summary>
        /// The gallery artifact version source.
        /// </summary>
        [Input("source")]
        public Input<Inputs.GalleryArtifactVersionSourceArgs>? Source { get; set; }

        public GalleryDataDiskImageArgs()
        {
        }
    }
}