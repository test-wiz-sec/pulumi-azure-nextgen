// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNextGen.Batch.V20200901.Outputs
{

    [OutputType]
    public sealed class CIFSMountConfigurationResponse
    {
        /// <summary>
        /// These are 'net use' options in Windows and 'mount' options in Linux.
        /// </summary>
        public readonly string? MountOptions;
        public readonly string Password;
        /// <summary>
        /// All file systems are mounted relative to the Batch mounts directory, accessible via the AZ_BATCH_NODE_MOUNTS_DIR environment variable.
        /// </summary>
        public readonly string RelativeMountPath;
        public readonly string Source;
        public readonly string Username;

        [OutputConstructor]
        private CIFSMountConfigurationResponse(
            string? mountOptions,

            string password,

            string relativeMountPath,

            string source,

            string username)
        {
            MountOptions = mountOptions;
            Password = password;
            RelativeMountPath = relativeMountPath;
            Source = source;
            Username = username;
        }
    }
}
