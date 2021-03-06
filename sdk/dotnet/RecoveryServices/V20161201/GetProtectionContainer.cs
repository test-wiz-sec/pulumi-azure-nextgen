// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNextGen.RecoveryServices.V20161201
{
    public static class GetProtectionContainer
    {
        public static Task<GetProtectionContainerResult> InvokeAsync(GetProtectionContainerArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetProtectionContainerResult>("azure-nextgen:recoveryservices/v20161201:getProtectionContainer", args ?? new GetProtectionContainerArgs(), options.WithVersion());
    }


    public sealed class GetProtectionContainerArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// Name of the container whose details need to be fetched.
        /// </summary>
        [Input("containerName", required: true)]
        public string ContainerName { get; set; } = null!;

        /// <summary>
        /// Name of the fabric where the container belongs.
        /// </summary>
        [Input("fabricName", required: true)]
        public string FabricName { get; set; } = null!;

        /// <summary>
        /// The name of the resource group where the recovery services vault is present.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public string ResourceGroupName { get; set; } = null!;

        /// <summary>
        /// The name of the recovery services vault.
        /// </summary>
        [Input("vaultName", required: true)]
        public string VaultName { get; set; } = null!;

        public GetProtectionContainerArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetProtectionContainerResult
    {
        /// <summary>
        /// Optional ETag.
        /// </summary>
        public readonly string? ETag;
        /// <summary>
        /// Resource location.
        /// </summary>
        public readonly string? Location;
        /// <summary>
        /// Resource name associated with the resource.
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// ProtectionContainerResource properties
        /// </summary>
        public readonly Union<Outputs.AzureSqlContainerResponse, Union<Outputs.AzureStorageContainerResponse, Union<Outputs.AzureWorkloadContainerResponse, Union<Outputs.DpmContainerResponse, Union<Outputs.GenericContainerResponse, Union<Outputs.IaaSVMContainerResponse, Outputs.MabContainerResponse>>>>>> Properties;
        /// <summary>
        /// Resource tags.
        /// </summary>
        public readonly ImmutableDictionary<string, string>? Tags;
        /// <summary>
        /// Resource type represents the complete path of the form Namespace/ResourceType/ResourceType/...
        /// </summary>
        public readonly string Type;

        [OutputConstructor]
        private GetProtectionContainerResult(
            string? eTag,

            string? location,

            string name,

            Union<Outputs.AzureSqlContainerResponse, Union<Outputs.AzureStorageContainerResponse, Union<Outputs.AzureWorkloadContainerResponse, Union<Outputs.DpmContainerResponse, Union<Outputs.GenericContainerResponse, Union<Outputs.IaaSVMContainerResponse, Outputs.MabContainerResponse>>>>>> properties,

            ImmutableDictionary<string, string>? tags,

            string type)
        {
            ETag = eTag;
            Location = location;
            Name = name;
            Properties = properties;
            Tags = tags;
            Type = type;
        }
    }
}
