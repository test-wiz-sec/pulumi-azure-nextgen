// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNextGen.Compute.V20181001.Inputs
{

    /// <summary>
    /// Describes the parameters of ephemeral disk settings that can be specified for operating system disk. &lt;br&gt;&lt;br&gt; NOTE: The ephemeral disk settings can only be specified for managed disk.
    /// </summary>
    public sealed class DiffDiskSettingsArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Specifies the ephemeral disk settings for operating system disk.
        /// </summary>
        [Input("option")]
        public Input<string>? Option { get; set; }

        public DiffDiskSettingsArgs()
        {
        }
    }
}
