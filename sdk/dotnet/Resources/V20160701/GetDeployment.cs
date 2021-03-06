// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNextGen.Resources.V20160701
{
    public static class GetDeployment
    {
        public static Task<GetDeploymentResult> InvokeAsync(GetDeploymentArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetDeploymentResult>("azure-nextgen:resources/v20160701:getDeployment", args ?? new GetDeploymentArgs(), options.WithVersion());
    }


    public sealed class GetDeploymentArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// The name of the deployment.
        /// </summary>
        [Input("deploymentName", required: true)]
        public string DeploymentName { get; set; } = null!;

        /// <summary>
        /// The name of the resource group to get. The name is case insensitive.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public string ResourceGroupName { get; set; } = null!;

        public GetDeploymentArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetDeploymentResult
    {
        /// <summary>
        /// The name of the deployment.
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// Deployment properties.
        /// </summary>
        public readonly Outputs.DeploymentPropertiesExtendedResponse Properties;

        [OutputConstructor]
        private GetDeploymentResult(
            string name,

            Outputs.DeploymentPropertiesExtendedResponse properties)
        {
            Name = name;
            Properties = properties;
        }
    }
}
