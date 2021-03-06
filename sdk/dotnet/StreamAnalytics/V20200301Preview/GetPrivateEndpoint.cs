// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNextGen.StreamAnalytics.V20200301Preview
{
    public static class GetPrivateEndpoint
    {
        public static Task<GetPrivateEndpointResult> InvokeAsync(GetPrivateEndpointArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetPrivateEndpointResult>("azure-nextgen:streamanalytics/v20200301preview:getPrivateEndpoint", args ?? new GetPrivateEndpointArgs(), options.WithVersion());
    }


    public sealed class GetPrivateEndpointArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// The name of the cluster.
        /// </summary>
        [Input("clusterName", required: true)]
        public string ClusterName { get; set; } = null!;

        /// <summary>
        /// The name of the private endpoint.
        /// </summary>
        [Input("privateEndpointName", required: true)]
        public string PrivateEndpointName { get; set; } = null!;

        /// <summary>
        /// The name of the resource group. The name is case insensitive.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public string ResourceGroupName { get; set; } = null!;

        public GetPrivateEndpointArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetPrivateEndpointResult
    {
        /// <summary>
        /// Unique opaque string (generally a GUID) that represents the metadata state of the resource (private endpoint) and changes whenever the resource is updated. Required on PUT (CreateOrUpdate) requests.
        /// </summary>
        public readonly string Etag;
        /// <summary>
        /// The name of the resource
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// The properties associated with a private endpoint.
        /// </summary>
        public readonly Outputs.PrivateEndpointPropertiesResponse Properties;
        /// <summary>
        /// The type of the resource. Ex- Microsoft.Compute/virtualMachines or Microsoft.Storage/storageAccounts.
        /// </summary>
        public readonly string Type;

        [OutputConstructor]
        private GetPrivateEndpointResult(
            string etag,

            string name,

            Outputs.PrivateEndpointPropertiesResponse properties,

            string type)
        {
            Etag = etag;
            Name = name;
            Properties = properties;
            Type = type;
        }
    }
}
