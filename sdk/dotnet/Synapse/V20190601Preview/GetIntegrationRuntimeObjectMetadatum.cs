// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNextGen.Synapse.V20190601Preview
{
    public static class GetIntegrationRuntimeObjectMetadatum
    {
        public static Task<GetIntegrationRuntimeObjectMetadatumResult> InvokeAsync(GetIntegrationRuntimeObjectMetadatumArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetIntegrationRuntimeObjectMetadatumResult>("azure-nextgen:synapse/v20190601preview:getIntegrationRuntimeObjectMetadatum", args ?? new GetIntegrationRuntimeObjectMetadatumArgs(), options.WithVersion());
    }


    public sealed class GetIntegrationRuntimeObjectMetadatumArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// Integration runtime name
        /// </summary>
        [Input("integrationRuntimeName", required: true)]
        public string IntegrationRuntimeName { get; set; } = null!;

        /// <summary>
        /// Metadata path.
        /// </summary>
        [Input("metadataPath")]
        public string? MetadataPath { get; set; }

        /// <summary>
        /// The name of the resource group. The name is case insensitive.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public string ResourceGroupName { get; set; } = null!;

        /// <summary>
        /// The name of the workspace
        /// </summary>
        [Input("workspaceName", required: true)]
        public string WorkspaceName { get; set; } = null!;

        public GetIntegrationRuntimeObjectMetadatumArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetIntegrationRuntimeObjectMetadatumResult
    {
        /// <summary>
        /// The link to the next page of results, if any remaining results exist.
        /// </summary>
        public readonly string? NextLink;
        /// <summary>
        /// List of SSIS object metadata.
        /// </summary>
        public readonly ImmutableArray<Union<Outputs.SsisEnvironmentResponseResult, Union<Outputs.SsisFolderResponseResult, Union<Outputs.SsisPackageResponseResult, Outputs.SsisProjectResponseResult>>>> Value;

        [OutputConstructor]
        private GetIntegrationRuntimeObjectMetadatumResult(
            string? nextLink,

            ImmutableArray<Union<Outputs.SsisEnvironmentResponseResult, Union<Outputs.SsisFolderResponseResult, Union<Outputs.SsisPackageResponseResult, Outputs.SsisProjectResponseResult>>>> value)
        {
            NextLink = nextLink;
            Value = value;
        }
    }
}