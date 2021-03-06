// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNextGen.ContainerRegistry.V20191201Preview
{
    public static class GetExportPipeline
    {
        public static Task<GetExportPipelineResult> InvokeAsync(GetExportPipelineArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetExportPipelineResult>("azure-nextgen:containerregistry/v20191201preview:getExportPipeline", args ?? new GetExportPipelineArgs(), options.WithVersion());
    }


    public sealed class GetExportPipelineArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// The name of the export pipeline.
        /// </summary>
        [Input("exportPipelineName", required: true)]
        public string ExportPipelineName { get; set; } = null!;

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

        public GetExportPipelineArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetExportPipelineResult
    {
        /// <summary>
        /// The identity of the export pipeline.
        /// </summary>
        public readonly Outputs.IdentityPropertiesResponse? Identity;
        /// <summary>
        /// The location of the export pipeline.
        /// </summary>
        public readonly string? Location;
        /// <summary>
        /// The name of the resource.
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// The list of all options configured for the pipeline.
        /// </summary>
        public readonly ImmutableArray<string> Options;
        /// <summary>
        /// The provisioning state of the pipeline at the time the operation was called.
        /// </summary>
        public readonly string ProvisioningState;
        /// <summary>
        /// The target properties of the export pipeline.
        /// </summary>
        public readonly Outputs.ExportPipelineTargetPropertiesResponse Target;
        /// <summary>
        /// The type of the resource.
        /// </summary>
        public readonly string Type;

        [OutputConstructor]
        private GetExportPipelineResult(
            Outputs.IdentityPropertiesResponse? identity,

            string? location,

            string name,

            ImmutableArray<string> options,

            string provisioningState,

            Outputs.ExportPipelineTargetPropertiesResponse target,

            string type)
        {
            Identity = identity;
            Location = location;
            Name = name;
            Options = options;
            ProvisioningState = provisioningState;
            Target = target;
            Type = type;
        }
    }
}
