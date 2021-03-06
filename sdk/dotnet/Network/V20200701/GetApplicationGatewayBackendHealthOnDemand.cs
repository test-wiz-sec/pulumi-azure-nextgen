// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNextGen.Network.V20200701
{
    public static class GetApplicationGatewayBackendHealthOnDemand
    {
        public static Task<GetApplicationGatewayBackendHealthOnDemandResult> InvokeAsync(GetApplicationGatewayBackendHealthOnDemandArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetApplicationGatewayBackendHealthOnDemandResult>("azure-nextgen:network/v20200701:getApplicationGatewayBackendHealthOnDemand", args ?? new GetApplicationGatewayBackendHealthOnDemandArgs(), options.WithVersion());
    }


    public sealed class GetApplicationGatewayBackendHealthOnDemandArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// The name of the application gateway.
        /// </summary>
        [Input("applicationGatewayName", required: true)]
        public string ApplicationGatewayName { get; set; } = null!;

        /// <summary>
        /// Reference to backend pool of application gateway to which probe request will be sent.
        /// </summary>
        [Input("backendAddressPool")]
        public Inputs.SubResourceArgs? BackendAddressPool { get; set; }

        /// <summary>
        /// Reference to backend http setting of application gateway to be used for test probe.
        /// </summary>
        [Input("backendHttpSettings")]
        public Inputs.SubResourceArgs? BackendHttpSettings { get; set; }

        /// <summary>
        /// Expands BackendAddressPool and BackendHttpSettings referenced in backend health.
        /// </summary>
        [Input("expand")]
        public string? Expand { get; set; }

        /// <summary>
        /// Host name to send the probe to.
        /// </summary>
        [Input("host")]
        public string? Host { get; set; }

        /// <summary>
        /// Criterion for classifying a healthy probe response.
        /// </summary>
        [Input("match")]
        public Inputs.ApplicationGatewayProbeHealthResponseMatchArgs? Match { get; set; }

        /// <summary>
        /// Relative path of probe. Valid path starts from '/'. Probe is sent to &lt;Protocol&gt;://&lt;host&gt;:&lt;port&gt;&lt;path&gt;.
        /// </summary>
        [Input("path")]
        public string? Path { get; set; }

        /// <summary>
        /// Whether the host header should be picked from the backend http settings. Default value is false.
        /// </summary>
        [Input("pickHostNameFromBackendHttpSettings")]
        public bool? PickHostNameFromBackendHttpSettings { get; set; }

        /// <summary>
        /// The protocol used for the probe.
        /// </summary>
        [Input("protocol")]
        public string? Protocol { get; set; }

        /// <summary>
        /// The name of the resource group.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public string ResourceGroupName { get; set; } = null!;

        /// <summary>
        /// The probe timeout in seconds. Probe marked as failed if valid response is not received with this timeout period. Acceptable values are from 1 second to 86400 seconds.
        /// </summary>
        [Input("timeout")]
        public int? Timeout { get; set; }

        public GetApplicationGatewayBackendHealthOnDemandArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetApplicationGatewayBackendHealthOnDemandResult
    {
        /// <summary>
        /// Reference to an ApplicationGatewayBackendAddressPool resource.
        /// </summary>
        public readonly Outputs.ApplicationGatewayBackendAddressPoolResponse? BackendAddressPool;
        /// <summary>
        /// Application gateway BackendHealthHttp settings.
        /// </summary>
        public readonly Outputs.ApplicationGatewayBackendHealthHttpSettingsResponseResult? BackendHealthHttpSettings;

        [OutputConstructor]
        private GetApplicationGatewayBackendHealthOnDemandResult(
            Outputs.ApplicationGatewayBackendAddressPoolResponse? backendAddressPool,

            Outputs.ApplicationGatewayBackendHealthHttpSettingsResponseResult? backendHealthHttpSettings)
        {
            BackendAddressPool = backendAddressPool;
            BackendHealthHttpSettings = backendHealthHttpSettings;
        }
    }
}
