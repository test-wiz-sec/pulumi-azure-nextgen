// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNextGen.DevOps.V20190701Preview
{
    public static class GetPipeline
    {
        public static Task<GetPipelineResult> InvokeAsync(GetPipelineArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetPipelineResult>("azure-nextgen:devops/v20190701preview:getPipeline", args ?? new GetPipelineArgs(), options.WithVersion());
    }


    public sealed class GetPipelineArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// The name of the Azure Pipeline resource in ARM.
        /// </summary>
        [Input("pipelineName", required: true)]
        public string PipelineName { get; set; } = null!;

        /// <summary>
        /// Name of the resource group within the Azure subscription.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public string ResourceGroupName { get; set; } = null!;

        public GetPipelineArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetPipelineResult
    {
        /// <summary>
        /// Configuration used to bootstrap the Pipeline.
        /// </summary>
        public readonly Outputs.BootstrapConfigurationResponse BootstrapConfiguration;
        /// <summary>
        /// Resource Location
        /// </summary>
        public readonly string? Location;
        /// <summary>
        /// Resource Name
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// Reference to the Azure DevOps Organization containing the Pipeline.
        /// </summary>
        public readonly Outputs.OrganizationReferenceResponse Organization;
        /// <summary>
        /// Unique identifier of the Azure Pipeline within the Azure DevOps Project.
        /// </summary>
        public readonly int PipelineId;
        /// <summary>
        /// Reference to the Azure DevOps Project containing the Pipeline.
        /// </summary>
        public readonly Outputs.ProjectReferenceResponse Project;
        /// <summary>
        /// Resource Tags
        /// </summary>
        public readonly ImmutableDictionary<string, string>? Tags;
        /// <summary>
        /// Resource Type
        /// </summary>
        public readonly string Type;

        [OutputConstructor]
        private GetPipelineResult(
            Outputs.BootstrapConfigurationResponse bootstrapConfiguration,

            string? location,

            string name,

            Outputs.OrganizationReferenceResponse organization,

            int pipelineId,

            Outputs.ProjectReferenceResponse project,

            ImmutableDictionary<string, string>? tags,

            string type)
        {
            BootstrapConfiguration = bootstrapConfiguration;
            Location = location;
            Name = name;
            Organization = organization;
            PipelineId = pipelineId;
            Project = project;
            Tags = tags;
            Type = type;
        }
    }
}
