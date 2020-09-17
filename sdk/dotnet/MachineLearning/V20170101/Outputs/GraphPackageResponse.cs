// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNextGen.MachineLearning.V20170101.Outputs
{

    [OutputType]
    public sealed class GraphPackageResponse
    {
        /// <summary>
        /// The list of edges making up the graph.
        /// </summary>
        public readonly ImmutableArray<Outputs.GraphEdgeResponse> Edges;
        /// <summary>
        /// The collection of global parameters for the graph, given as a global parameter name to GraphParameter map. Each parameter here has a 1:1 match with the global parameters values map declared at the WebServiceProperties level.
        /// </summary>
        public readonly ImmutableDictionary<string, Outputs.GraphParameterResponse>? GraphParameters;
        /// <summary>
        /// The set of nodes making up the graph, provided as a nodeId to GraphNode map
        /// </summary>
        public readonly ImmutableDictionary<string, Outputs.GraphNodeResponse>? Nodes;

        [OutputConstructor]
        private GraphPackageResponse(
            ImmutableArray<Outputs.GraphEdgeResponse> edges,

            ImmutableDictionary<string, Outputs.GraphParameterResponse>? graphParameters,

            ImmutableDictionary<string, Outputs.GraphNodeResponse>? nodes)
        {
            Edges = edges;
            GraphParameters = graphParameters;
            Nodes = nodes;
        }
    }
}