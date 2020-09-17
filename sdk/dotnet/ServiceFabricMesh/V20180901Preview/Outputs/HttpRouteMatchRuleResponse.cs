// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNextGen.ServiceFabricMesh.V20180901Preview.Outputs
{

    [OutputType]
    public sealed class HttpRouteMatchRuleResponse
    {
        /// <summary>
        /// headers and their values to match in request.
        /// </summary>
        public readonly ImmutableArray<Outputs.HttpRouteMatchHeaderResponse> Headers;
        /// <summary>
        /// Path to match for routing.
        /// </summary>
        public readonly Outputs.HttpRouteMatchPathResponse Path;

        [OutputConstructor]
        private HttpRouteMatchRuleResponse(
            ImmutableArray<Outputs.HttpRouteMatchHeaderResponse> headers,

            Outputs.HttpRouteMatchPathResponse path)
        {
            Headers = headers;
            Path = path;
        }
    }
}