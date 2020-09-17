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
    public sealed class VirtualMachineConfigurationResponse
    {
        /// <summary>
        /// If specified, setup is performed on each node in the pool to allow tasks to run in containers. All regular tasks and job manager tasks run on this pool must specify the containerSettings property, and all other tasks may specify it.
        /// </summary>
        public readonly Outputs.ContainerConfigurationResponse? ContainerConfiguration;
        /// <summary>
        /// This property must be specified if the compute nodes in the pool need to have empty data disks attached to them.
        /// </summary>
        public readonly ImmutableArray<Outputs.DataDiskResponse> DataDisks;
        /// <summary>
        /// If specified, encryption is performed on each node in the pool during node provisioning.
        /// </summary>
        public readonly Outputs.DiskEncryptionConfigurationResponse? DiskEncryptionConfiguration;
        public readonly Outputs.ImageReferenceResponse ImageReference;
        /// <summary>
        /// This only applies to images that contain the Windows operating system, and should only be used when you hold valid on-premises licenses for the nodes which will be deployed. If omitted, no on-premises licensing discount is applied. Values are:
        /// 
        ///  Windows_Server - The on-premises license is for Windows Server.
        ///  Windows_Client - The on-premises license is for Windows Client.
        /// </summary>
        public readonly string? LicenseType;
        /// <summary>
        /// The Batch node agent is a program that runs on each node in the pool, and provides the command-and-control interface between the node and the Batch service. There are different implementations of the node agent, known as SKUs, for different operating systems. You must specify a node agent SKU which matches the selected image reference. To get the list of supported node agent SKUs along with their list of verified image references, see the 'List supported node agent SKUs' operation.
        /// </summary>
        public readonly string NodeAgentSkuId;
        /// <summary>
        /// This property must not be specified if the imageReference specifies a Linux OS image.
        /// </summary>
        public readonly Outputs.WindowsConfigurationResponse? WindowsConfiguration;

        [OutputConstructor]
        private VirtualMachineConfigurationResponse(
            Outputs.ContainerConfigurationResponse? containerConfiguration,

            ImmutableArray<Outputs.DataDiskResponse> dataDisks,

            Outputs.DiskEncryptionConfigurationResponse? diskEncryptionConfiguration,

            Outputs.ImageReferenceResponse imageReference,

            string? licenseType,

            string nodeAgentSkuId,

            Outputs.WindowsConfigurationResponse? windowsConfiguration)
        {
            ContainerConfiguration = containerConfiguration;
            DataDisks = dataDisks;
            DiskEncryptionConfiguration = diskEncryptionConfiguration;
            ImageReference = imageReference;
            LicenseType = licenseType;
            NodeAgentSkuId = nodeAgentSkuId;
            WindowsConfiguration = windowsConfiguration;
        }
    }
}
