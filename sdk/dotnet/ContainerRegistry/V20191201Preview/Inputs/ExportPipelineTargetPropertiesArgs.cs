// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNextGen.ContainerRegistry.V20191201Preview.Inputs
{

    /// <summary>
    /// The properties of the export pipeline target.
    /// </summary>
    public sealed class ExportPipelineTargetPropertiesArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// They key vault secret uri to obtain the target storage SAS token.
        /// </summary>
        [Input("keyVaultUri", required: true)]
        public Input<string> KeyVaultUri { get; set; } = null!;

        /// <summary>
        /// The type of target for the export pipeline.
        /// </summary>
        [Input("type")]
        public Input<string>? Type { get; set; }

        /// <summary>
        /// The target uri of the export pipeline.
        /// When 'AzureStorageBlob': "https://accountName.blob.core.windows.net/containerName/blobName"
        /// When 'AzureStorageBlobContainer':  "https://accountName.blob.core.windows.net/containerName"
        /// </summary>
        [Input("uri")]
        public Input<string>? Uri { get; set; }

        public ExportPipelineTargetPropertiesArgs()
        {
        }
    }
}