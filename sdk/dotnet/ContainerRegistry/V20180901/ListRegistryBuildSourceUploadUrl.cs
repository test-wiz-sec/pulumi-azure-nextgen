// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNextGen.ContainerRegistry.V20180901
{
    public static class ListRegistryBuildSourceUploadUrl
    {
        public static Task<ListRegistryBuildSourceUploadUrlResult> InvokeAsync(ListRegistryBuildSourceUploadUrlArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<ListRegistryBuildSourceUploadUrlResult>("azure-nextgen:containerregistry/v20180901:listRegistryBuildSourceUploadUrl", args ?? new ListRegistryBuildSourceUploadUrlArgs(), options.WithVersion());
    }


    public sealed class ListRegistryBuildSourceUploadUrlArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// The name of the container registry.
        /// </summary>
        [Input("registryName", required: true)]
        public string RegistryName { get; set; } = null!;

        /// <summary>
        /// The name of the resource group to which the container registry belongs.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public string ResourceGroupName { get; set; } = null!;

        public ListRegistryBuildSourceUploadUrlArgs()
        {
        }
    }


    [OutputType]
    public sealed class ListRegistryBuildSourceUploadUrlResult
    {
        /// <summary>
        /// The relative path to the source. This is used to submit the subsequent queue build request.
        /// </summary>
        public readonly string? RelativePath;
        /// <summary>
        /// The URL where the client can upload the source.
        /// </summary>
        public readonly string? UploadUrl;

        [OutputConstructor]
        private ListRegistryBuildSourceUploadUrlResult(
            string? relativePath,

            string? uploadUrl)
        {
            RelativePath = relativePath;
            UploadUrl = uploadUrl;
        }
    }
}