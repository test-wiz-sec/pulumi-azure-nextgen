// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNextGen.Security.V20170801Preview
{
    /// <summary>
    /// Contact details for security issues
    /// </summary>
    public partial class SecurityContact : Pulumi.CustomResource
    {
        /// <summary>
        /// Whether to send security alerts notifications to the security contact
        /// </summary>
        [Output("alertNotifications")]
        public Output<string> AlertNotifications { get; private set; } = null!;

        /// <summary>
        /// Whether to send security alerts notifications to subscription admins
        /// </summary>
        [Output("alertsToAdmins")]
        public Output<string> AlertsToAdmins { get; private set; } = null!;

        /// <summary>
        /// The email of this security contact
        /// </summary>
        [Output("email")]
        public Output<string> Email { get; private set; } = null!;

        /// <summary>
        /// Resource name
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// The phone number of this security contact
        /// </summary>
        [Output("phone")]
        public Output<string?> Phone { get; private set; } = null!;

        /// <summary>
        /// Resource type
        /// </summary>
        [Output("type")]
        public Output<string> Type { get; private set; } = null!;


        /// <summary>
        /// Create a SecurityContact resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public SecurityContact(string name, SecurityContactArgs args, CustomResourceOptions? options = null)
            : base("azure-nextgen:security/v20170801preview:SecurityContact", name, args ?? new SecurityContactArgs(), MakeResourceOptions(options, ""))
        {
        }

        private SecurityContact(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("azure-nextgen:security/v20170801preview:SecurityContact", name, null, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                Aliases =
                {
                    new Pulumi.Alias { Type = "azure-nextgen:security/v20200101preview:SecurityContact"},
                },
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing SecurityContact resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static SecurityContact Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new SecurityContact(name, id, options);
        }
    }

    public sealed class SecurityContactArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Whether to send security alerts notifications to the security contact
        /// </summary>
        [Input("alertNotifications", required: true)]
        public Input<string> AlertNotifications { get; set; } = null!;

        /// <summary>
        /// Whether to send security alerts notifications to subscription admins
        /// </summary>
        [Input("alertsToAdmins", required: true)]
        public Input<string> AlertsToAdmins { get; set; } = null!;

        /// <summary>
        /// The email of this security contact
        /// </summary>
        [Input("email", required: true)]
        public Input<string> Email { get; set; } = null!;

        /// <summary>
        /// The phone number of this security contact
        /// </summary>
        [Input("phone")]
        public Input<string>? Phone { get; set; }

        /// <summary>
        /// Name of the security contact object
        /// </summary>
        [Input("securityContactName", required: true)]
        public Input<string> SecurityContactName { get; set; } = null!;

        public SecurityContactArgs()
        {
        }
    }
}