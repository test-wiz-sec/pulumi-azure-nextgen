// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNextGen.Datadog.V20200201Preview
{
    public static class GetMonitor
    {
        public static Task<GetMonitorResult> InvokeAsync(GetMonitorArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetMonitorResult>("azure-nextgen:datadog/v20200201preview:getMonitor", args ?? new GetMonitorArgs(), options.WithVersion());
    }


    public sealed class GetMonitorArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// Monitor resource name
        /// </summary>
        [Input("monitorName", required: true)]
        public string MonitorName { get; set; } = null!;

        /// <summary>
        /// The name of the resource group to which the Datadog resource belongs.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public string ResourceGroupName { get; set; } = null!;

        public GetMonitorArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetMonitorResult
    {
        public readonly Outputs.IdentityPropertiesResponse? Identity;
        public readonly string Location;
        /// <summary>
        /// Name of the monitor resource.
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// Properties specific to the monitor resource.
        /// </summary>
        public readonly Outputs.MonitorPropertiesResponse Properties;
        public readonly Outputs.ResourceSkuResponse? Sku;
        public readonly ImmutableDictionary<string, string>? Tags;
        /// <summary>
        /// The type of the monitor resource.
        /// </summary>
        public readonly string Type;

        [OutputConstructor]
        private GetMonitorResult(
            Outputs.IdentityPropertiesResponse? identity,

            string location,

            string name,

            Outputs.MonitorPropertiesResponse properties,

            Outputs.ResourceSkuResponse? sku,

            ImmutableDictionary<string, string>? tags,

            string type)
        {
            Identity = identity;
            Location = location;
            Name = name;
            Properties = properties;
            Sku = sku;
            Tags = tags;
            Type = type;
        }
    }
}