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
    /// Provides required information, for the service to be able to mount Azure Blob Storage container on the cluster nodes.
    /// </summary>
    public sealed class AzureBlobFileSystemReferenceArgs : Pulumi.ResourceArgs
    {
        [Input("accountName", required: true)]
        public Input<string> AccountName { get; set; } = null!;

        [Input("containerName", required: true)]
        public Input<string> ContainerName { get; set; } = null!;

        /// <summary>
        /// Credentials to access Azure File Share.
        /// </summary>
        [Input("credentials", required: true)]
        public Input<Inputs.AzureStorageCredentialsInfoArgs> Credentials { get; set; } = null!;

        [Input("mountOptions")]
        public Input<string>? MountOptions { get; set; }

        /// <summary>
        /// Note that all blob file systems will be mounted under $AZ_BATCHAI_MOUNT_ROOT location.
        /// </summary>
        [Input("relativeMountPath", required: true)]
        public Input<string> RelativeMountPath { get; set; } = null!;

        public AzureBlobFileSystemReferenceArgs()
        {
        }
    }
}
