// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNextGen.Automation.V20170515Preview
{
    /// <summary>
    /// Definition of the source control.
    /// </summary>
    public partial class SourceControl : Pulumi.CustomResource
    {
        /// <summary>
        /// The auto sync of the source control. Default is false.
        /// </summary>
        [Output("autoSync")]
        public Output<bool?> AutoSync { get; private set; } = null!;

        /// <summary>
        /// The repo branch of the source control. Include branch as empty string for VsoTfvc.
        /// </summary>
        [Output("branch")]
        public Output<string?> Branch { get; private set; } = null!;

        /// <summary>
        /// The creation time.
        /// </summary>
        [Output("creationTime")]
        public Output<string?> CreationTime { get; private set; } = null!;

        /// <summary>
        /// The description.
        /// </summary>
        [Output("description")]
        public Output<string?> Description { get; private set; } = null!;

        /// <summary>
        /// The folder path of the source control.
        /// </summary>
        [Output("folderPath")]
        public Output<string?> FolderPath { get; private set; } = null!;

        /// <summary>
        /// The last modified time.
        /// </summary>
        [Output("lastModifiedTime")]
        public Output<string?> LastModifiedTime { get; private set; } = null!;

        /// <summary>
        /// The name of the resource
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// The auto publish of the source control. Default is true.
        /// </summary>
        [Output("publishRunbook")]
        public Output<bool?> PublishRunbook { get; private set; } = null!;

        /// <summary>
        /// The repo url of the source control.
        /// </summary>
        [Output("repoUrl")]
        public Output<string?> RepoUrl { get; private set; } = null!;

        /// <summary>
        /// The source type. Must be one of VsoGit, VsoTfvc, GitHub.
        /// </summary>
        [Output("sourceType")]
        public Output<string?> SourceType { get; private set; } = null!;

        /// <summary>
        /// The type of the resource.
        /// </summary>
        [Output("type")]
        public Output<string> Type { get; private set; } = null!;


        /// <summary>
        /// Create a SourceControl resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public SourceControl(string name, SourceControlArgs args, CustomResourceOptions? options = null)
            : base("azure-nextgen:automation/v20170515preview:SourceControl", name, args ?? new SourceControlArgs(), MakeResourceOptions(options, ""))
        {
        }

        private SourceControl(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("azure-nextgen:automation/v20170515preview:SourceControl", name, null, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing SourceControl resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static SourceControl Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new SourceControl(name, id, options);
        }
    }

    public sealed class SourceControlArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The auto async of the source control. Default is false.
        /// </summary>
        [Input("autoSync")]
        public Input<bool>? AutoSync { get; set; }

        /// <summary>
        /// The name of the automation account.
        /// </summary>
        [Input("automationAccountName", required: true)]
        public Input<string> AutomationAccountName { get; set; } = null!;

        /// <summary>
        /// The repo branch of the source control. Include branch as empty string for VsoTfvc.
        /// </summary>
        [Input("branch")]
        public Input<string>? Branch { get; set; }

        /// <summary>
        /// The user description of the source control.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// The folder path of the source control. Path must be relative.
        /// </summary>
        [Input("folderPath")]
        public Input<string>? FolderPath { get; set; }

        /// <summary>
        /// The auto publish of the source control. Default is true.
        /// </summary>
        [Input("publishRunbook")]
        public Input<bool>? PublishRunbook { get; set; }

        /// <summary>
        /// The repo url of the source control.
        /// </summary>
        [Input("repoUrl")]
        public Input<string>? RepoUrl { get; set; }

        /// <summary>
        /// Name of an Azure Resource group.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public Input<string> ResourceGroupName { get; set; } = null!;

        /// <summary>
        /// The authorization token for the repo of the source control.
        /// </summary>
        [Input("securityToken")]
        public Input<Inputs.SourceControlSecurityTokenPropertiesArgs>? SecurityToken { get; set; }

        /// <summary>
        /// The source control name.
        /// </summary>
        [Input("sourceControlName", required: true)]
        public Input<string> SourceControlName { get; set; } = null!;

        /// <summary>
        /// The source type. Must be one of VsoGit, VsoTfvc, GitHub, case sensitive.
        /// </summary>
        [Input("sourceType")]
        public Input<string>? SourceType { get; set; }

        public SourceControlArgs()
        {
        }
    }
}