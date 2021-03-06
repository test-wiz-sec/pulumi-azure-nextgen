// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNextGen.StorageSync.V20200901
{
    public static class GetSyncGroup
    {
        public static Task<GetSyncGroupResult> InvokeAsync(GetSyncGroupArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetSyncGroupResult>("azure-nextgen:storagesync/v20200901:getSyncGroup", args ?? new GetSyncGroupArgs(), options.WithVersion());
    }


    public sealed class GetSyncGroupArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// The name of the resource group. The name is case insensitive.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public string ResourceGroupName { get; set; } = null!;

        /// <summary>
        /// Name of Storage Sync Service resource.
        /// </summary>
        [Input("storageSyncServiceName", required: true)]
        public string StorageSyncServiceName { get; set; } = null!;

        /// <summary>
        /// Name of Sync Group resource.
        /// </summary>
        [Input("syncGroupName", required: true)]
        public string SyncGroupName { get; set; } = null!;

        public GetSyncGroupArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetSyncGroupResult
    {
        /// <summary>
        /// The name of the resource
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// Sync group status
        /// </summary>
        public readonly string SyncGroupStatus;
        /// <summary>
        /// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
        /// </summary>
        public readonly string Type;
        /// <summary>
        /// Unique Id
        /// </summary>
        public readonly string UniqueId;

        [OutputConstructor]
        private GetSyncGroupResult(
            string name,

            string syncGroupStatus,

            string type,

            string uniqueId)
        {
            Name = name;
            SyncGroupStatus = syncGroupStatus;
            Type = type;
            UniqueId = uniqueId;
        }
    }
}
