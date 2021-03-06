// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNextGen.RecoveryServices.V20160601
{
    /// <summary>
    /// The base class for backup policy. Workload-specific backup policies are derived from this class.
    /// </summary>
    public partial class ProtectionPolicy : Pulumi.CustomResource
    {
        /// <summary>
        /// Optional ETag.
        /// </summary>
        [Output("eTag")]
        public Output<string?> ETag { get; private set; } = null!;

        /// <summary>
        /// Resource location.
        /// </summary>
        [Output("location")]
        public Output<string?> Location { get; private set; } = null!;

        /// <summary>
        /// Resource name associated with the resource.
        /// </summary>
        [Output("name")]
        public Output<string?> Name { get; private set; } = null!;

        /// <summary>
        /// The base class for a backup policy. Workload-specific backup policies are derived from this class.
        /// </summary>
        [Output("properties")]
        public Output<Union<Outputs.AzureIaaSVMProtectionPolicyResponse, Union<Outputs.AzureSqlProtectionPolicyResponse, Outputs.MabProtectionPolicyResponse>>> Properties { get; private set; } = null!;

        /// <summary>
        /// Resource tags.
        /// </summary>
        [Output("tags")]
        public Output<ImmutableDictionary<string, string>?> Tags { get; private set; } = null!;

        /// <summary>
        /// Resource type represents the complete path of the form Namespace/ResourceType/ResourceType/...
        /// </summary>
        [Output("type")]
        public Output<string?> Type { get; private set; } = null!;


        /// <summary>
        /// Create a ProtectionPolicy resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public ProtectionPolicy(string name, ProtectionPolicyArgs args, CustomResourceOptions? options = null)
            : base("azure-nextgen:recoveryservices/v20160601:ProtectionPolicy", name, args ?? new ProtectionPolicyArgs(), MakeResourceOptions(options, ""))
        {
        }

        private ProtectionPolicy(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("azure-nextgen:recoveryservices/v20160601:ProtectionPolicy", name, null, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                Aliases =
                {
                    new Pulumi.Alias { Type = "azure-nextgen:recoveryservices/latest:ProtectionPolicy"},
                },
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing ProtectionPolicy resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static ProtectionPolicy Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new ProtectionPolicy(name, id, options);
        }
    }

    public sealed class ProtectionPolicyArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Optional ETag.
        /// </summary>
        [Input("eTag")]
        public Input<string>? ETag { get; set; }

        /// <summary>
        /// Resource ID represents the complete path to the resource.
        /// </summary>
        [Input("id")]
        public Input<string>? Id { get; set; }

        /// <summary>
        /// Resource location.
        /// </summary>
        [Input("location")]
        public Input<string>? Location { get; set; }

        /// <summary>
        /// Resource name associated with the resource.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The backup policy to be created.
        /// </summary>
        [Input("policyName", required: true)]
        public Input<string> PolicyName { get; set; } = null!;

        /// <summary>
        /// The base class for a backup policy. Workload-specific backup policies are derived from this class.
        /// </summary>
        [Input("properties")]
        public InputUnion<Inputs.AzureIaaSVMProtectionPolicyArgs, InputUnion<Inputs.AzureSqlProtectionPolicyArgs, Inputs.MabProtectionPolicyArgs>>? Properties { get; set; }

        /// <summary>
        /// The name of the resource group associated with the Recovery Services vault.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public Input<string> ResourceGroupName { get; set; } = null!;

        [Input("tags")]
        private InputMap<string>? _tags;

        /// <summary>
        /// Resource tags.
        /// </summary>
        public InputMap<string> Tags
        {
            get => _tags ?? (_tags = new InputMap<string>());
            set => _tags = value;
        }

        /// <summary>
        /// Resource type represents the complete path of the form Namespace/ResourceType/ResourceType/...
        /// </summary>
        [Input("type")]
        public Input<string>? Type { get; set; }

        /// <summary>
        /// The name of the Recovery Services vault.
        /// </summary>
        [Input("vaultName", required: true)]
        public Input<string> VaultName { get; set; } = null!;

        public ProtectionPolicyArgs()
        {
        }
    }
}
