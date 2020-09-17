// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNextGen.Batch.V20200501
{
    /// <summary>
    /// Contains information about an application in a Batch account.
    /// </summary>
    public partial class Application : Pulumi.CustomResource
    {
        /// <summary>
        /// A value indicating whether packages within the application may be overwritten using the same version string.
        /// </summary>
        [Output("allowUpdates")]
        public Output<bool?> AllowUpdates { get; private set; } = null!;

        /// <summary>
        /// The package to use if a client requests the application but does not specify a version. This property can only be set to the name of an existing package.
        /// </summary>
        [Output("defaultVersion")]
        public Output<string?> DefaultVersion { get; private set; } = null!;

        /// <summary>
        /// The display name for the application.
        /// </summary>
        [Output("displayName")]
        public Output<string?> DisplayName { get; private set; } = null!;

        /// <summary>
        /// The ETag of the resource, used for concurrency statements.
        /// </summary>
        [Output("etag")]
        public Output<string> Etag { get; private set; } = null!;

        /// <summary>
        /// The name of the resource.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// The type of the resource.
        /// </summary>
        [Output("type")]
        public Output<string> Type { get; private set; } = null!;


        /// <summary>
        /// Create a Application resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Application(string name, ApplicationArgs args, CustomResourceOptions? options = null)
            : base("azure-nextgen:batch/v20200501:Application", name, args ?? new ApplicationArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Application(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("azure-nextgen:batch/v20200501:Application", name, null, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                Aliases =
                {
                    new Pulumi.Alias { Type = "azure-nextgen:batch/latest:Application"},
                    new Pulumi.Alias { Type = "azure-nextgen:batch/v20151201:Application"},
                    new Pulumi.Alias { Type = "azure-nextgen:batch/v20170101:Application"},
                    new Pulumi.Alias { Type = "azure-nextgen:batch/v20170501:Application"},
                    new Pulumi.Alias { Type = "azure-nextgen:batch/v20170901:Application"},
                    new Pulumi.Alias { Type = "azure-nextgen:batch/v20181201:Application"},
                    new Pulumi.Alias { Type = "azure-nextgen:batch/v20190401:Application"},
                    new Pulumi.Alias { Type = "azure-nextgen:batch/v20190801:Application"},
                    new Pulumi.Alias { Type = "azure-nextgen:batch/v20200301:Application"},
                    new Pulumi.Alias { Type = "azure-nextgen:batch/v20200901:Application"},
                },
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing Application resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Application Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new Application(name, id, options);
        }
    }

    public sealed class ApplicationArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The name of the Batch account.
        /// </summary>
        [Input("accountName", required: true)]
        public Input<string> AccountName { get; set; } = null!;

        /// <summary>
        /// A value indicating whether packages within the application may be overwritten using the same version string.
        /// </summary>
        [Input("allowUpdates")]
        public Input<bool>? AllowUpdates { get; set; }

        /// <summary>
        /// The name of the application. This must be unique within the account.
        /// </summary>
        [Input("applicationName", required: true)]
        public Input<string> ApplicationName { get; set; } = null!;

        /// <summary>
        /// The package to use if a client requests the application but does not specify a version. This property can only be set to the name of an existing package.
        /// </summary>
        [Input("defaultVersion")]
        public Input<string>? DefaultVersion { get; set; }

        /// <summary>
        /// The display name for the application.
        /// </summary>
        [Input("displayName")]
        public Input<string>? DisplayName { get; set; }

        /// <summary>
        /// The name of the resource group that contains the Batch account.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public Input<string> ResourceGroupName { get; set; } = null!;

        public ApplicationArgs()
        {
        }
    }
}
