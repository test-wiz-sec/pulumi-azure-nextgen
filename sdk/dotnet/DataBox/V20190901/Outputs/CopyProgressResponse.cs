// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNextGen.DataBox.V20190901.Outputs
{

    [OutputType]
    public sealed class CopyProgressResponse
    {
        /// <summary>
        /// Id of the account where the data needs to be uploaded.
        /// </summary>
        public readonly string AccountId;
        /// <summary>
        /// Amount of data uploaded by the job as of now.
        /// </summary>
        public readonly int BytesSentToCloud;
        /// <summary>
        /// Data Destination Type.
        /// </summary>
        public readonly string DataDestinationType;
        /// <summary>
        /// Number of files which could not be copied
        /// </summary>
        public readonly int FilesErroredOut;
        /// <summary>
        /// Number of files processed by the job as of now.
        /// </summary>
        public readonly int FilesProcessed;
        /// <summary>
        /// Total amount of data not adhering to azure naming conventions which were processed by automatic renaming
        /// </summary>
        public readonly int InvalidFileBytesUploaded;
        /// <summary>
        /// Number of files not adhering to azure naming conventions which were processed by automatic renaming
        /// </summary>
        public readonly int InvalidFilesProcessed;
        /// <summary>
        /// Number of folders not adhering to azure naming conventions which were processed by automatic renaming
        /// </summary>
        public readonly int RenamedContainerCount;
        /// <summary>
        /// Name of the storage account where the data needs to be uploaded.
        /// </summary>
        public readonly string StorageAccountName;
        /// <summary>
        /// Total amount of data to be processed by the job.
        /// </summary>
        public readonly int TotalBytesToProcess;
        /// <summary>
        /// Total number of files to be processed by the job.
        /// </summary>
        public readonly int TotalFilesToProcess;

        [OutputConstructor]
        private CopyProgressResponse(
            string accountId,

            int bytesSentToCloud,

            string dataDestinationType,

            int filesErroredOut,

            int filesProcessed,

            int invalidFileBytesUploaded,

            int invalidFilesProcessed,

            int renamedContainerCount,

            string storageAccountName,

            int totalBytesToProcess,

            int totalFilesToProcess)
        {
            AccountId = accountId;
            BytesSentToCloud = bytesSentToCloud;
            DataDestinationType = dataDestinationType;
            FilesErroredOut = filesErroredOut;
            FilesProcessed = filesProcessed;
            InvalidFileBytesUploaded = invalidFileBytesUploaded;
            InvalidFilesProcessed = invalidFilesProcessed;
            RenamedContainerCount = renamedContainerCount;
            StorageAccountName = storageAccountName;
            TotalBytesToProcess = totalBytesToProcess;
            TotalFilesToProcess = totalFilesToProcess;
        }
    }
}
