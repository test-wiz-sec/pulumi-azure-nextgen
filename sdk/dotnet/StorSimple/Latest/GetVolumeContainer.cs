// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNextGen.StorSimple.Latest
{
    public static class GetVolumeContainer
    {
        public static Task<GetVolumeContainerResult> InvokeAsync(GetVolumeContainerArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetVolumeContainerResult>("azure-nextgen:storsimple/latest:getVolumeContainer", args ?? new GetVolumeContainerArgs(), options.WithVersion());
    }


    public sealed class GetVolumeContainerArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// The device name
        /// </summary>
        [Input("deviceName", required: true)]
        public string DeviceName { get; set; } = null!;

        /// <summary>
        /// The manager name
        /// </summary>
        [Input("managerName", required: true)]
        public string ManagerName { get; set; } = null!;

        /// <summary>
        /// The resource group name
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public string ResourceGroupName { get; set; } = null!;

        /// <summary>
        /// The name of the volume container.
        /// </summary>
        [Input("volumeContainerName", required: true)]
        public string VolumeContainerName { get; set; } = null!;

        public GetVolumeContainerArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetVolumeContainerResult
    {
        /// <summary>
        /// The bandwidth-rate set on the volume container.
        /// </summary>
        public readonly int? BandWidthRateInMbps;
        /// <summary>
        /// The ID of the bandwidth setting associated with the volume container.
        /// </summary>
        public readonly string? BandwidthSettingId;
        /// <summary>
        /// The key used to encrypt data in the volume container. It is required when property 'EncryptionStatus' is "Enabled".
        /// </summary>
        public readonly Outputs.AsymmetricEncryptedSecretResponse? EncryptionKey;
        /// <summary>
        /// The flag to denote whether encryption is enabled or not.
        /// </summary>
        public readonly string EncryptionStatus;
        /// <summary>
        /// The Kind of the object. Currently only Series8000 is supported
        /// </summary>
        public readonly string? Kind;
        /// <summary>
        /// The name of the object.
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// The owner ship status of the volume container. Only when the status is "NotOwned", the delete operation on the volume container is permitted.
        /// </summary>
        public readonly string OwnerShipStatus;
        /// <summary>
        /// The path ID of storage account associated with the volume container.
        /// </summary>
        public readonly string StorageAccountCredentialId;
        /// <summary>
        /// The total cloud storage for the volume container.
        /// </summary>
        public readonly int TotalCloudStorageUsageInBytes;
        /// <summary>
        /// The hierarchical type of the object.
        /// </summary>
        public readonly string Type;
        /// <summary>
        /// The number of volumes in the volume Container.
        /// </summary>
        public readonly int VolumeCount;

        [OutputConstructor]
        private GetVolumeContainerResult(
            int? bandWidthRateInMbps,

            string? bandwidthSettingId,

            Outputs.AsymmetricEncryptedSecretResponse? encryptionKey,

            string encryptionStatus,

            string? kind,

            string name,

            string ownerShipStatus,

            string storageAccountCredentialId,

            int totalCloudStorageUsageInBytes,

            string type,

            int volumeCount)
        {
            BandWidthRateInMbps = bandWidthRateInMbps;
            BandwidthSettingId = bandwidthSettingId;
            EncryptionKey = encryptionKey;
            EncryptionStatus = encryptionStatus;
            Kind = kind;
            Name = name;
            OwnerShipStatus = ownerShipStatus;
            StorageAccountCredentialId = storageAccountCredentialId;
            TotalCloudStorageUsageInBytes = totalCloudStorageUsageInBytes;
            Type = type;
            VolumeCount = volumeCount;
        }
    }
}