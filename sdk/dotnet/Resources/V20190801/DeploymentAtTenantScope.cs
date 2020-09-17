// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNextGen.Resources.V20190801
{
    /// <summary>
    /// Deployment information.
    /// </summary>
    public partial class DeploymentAtTenantScope : Pulumi.CustomResource
    {
        /// <summary>
        /// the location of the deployment.
        /// </summary>
        [Output("location")]
        public Output<string?> Location { get; private set; } = null!;

        /// <summary>
        /// The name of the deployment.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// Deployment properties.
        /// </summary>
        [Output("properties")]
        public Output<Outputs.DeploymentPropertiesExtendedResponse> Properties { get; private set; } = null!;

        /// <summary>
        /// The type of the deployment.
        /// </summary>
        [Output("type")]
        public Output<string> Type { get; private set; } = null!;


        /// <summary>
        /// Create a DeploymentAtTenantScope resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public DeploymentAtTenantScope(string name, DeploymentAtTenantScopeArgs args, CustomResourceOptions? options = null)
            : base("azure-nextgen:resources/v20190801:DeploymentAtTenantScope", name, args ?? new DeploymentAtTenantScopeArgs(), MakeResourceOptions(options, ""))
        {
        }

        private DeploymentAtTenantScope(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("azure-nextgen:resources/v20190801:DeploymentAtTenantScope", name, null, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                Aliases =
                {
                    new Pulumi.Alias { Type = "azure-nextgen:resources/latest:DeploymentAtTenantScope"},
                    new Pulumi.Alias { Type = "azure-nextgen:resources/v20190701:DeploymentAtTenantScope"},
                    new Pulumi.Alias { Type = "azure-nextgen:resources/v20191001:DeploymentAtTenantScope"},
                    new Pulumi.Alias { Type = "azure-nextgen:resources/v20200601:DeploymentAtTenantScope"},
                },
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing DeploymentAtTenantScope resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static DeploymentAtTenantScope Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new DeploymentAtTenantScope(name, id, options);
        }
    }

    public sealed class DeploymentAtTenantScopeArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The name of the deployment.
        /// </summary>
        [Input("deploymentName", required: true)]
        public Input<string> DeploymentName { get; set; } = null!;

        /// <summary>
        /// The location to store the deployment data.
        /// </summary>
        [Input("location", required: true)]
        public Input<string> Location { get; set; } = null!;

        /// <summary>
        /// The deployment properties.
        /// </summary>
        [Input("properties", required: true)]
        public Input<Inputs.DeploymentPropertiesArgs> Properties { get; set; } = null!;

        public DeploymentAtTenantScopeArgs()
        {
        }
    }
}