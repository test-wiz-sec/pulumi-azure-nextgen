// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNextGen.BatchAI.V20170901Preview.Inputs
{

    /// <summary>
    /// Settings for the data disk which would be created for the File Server.
    /// </summary>
    public sealed class DataDisksArgs : Pulumi.ResourceArgs
    {
        [Input("diskCount", required: true)]
        public Input<int> DiskCount { get; set; } = null!;

        [Input("diskSizeInGB", required: true)]
        public Input<int> DiskSizeInGB { get; set; } = null!;

        [Input("storageAccountType", required: true)]
        public Input<string> StorageAccountType { get; set; } = null!;

        public DataDisksArgs()
        {
        }
    }
}
