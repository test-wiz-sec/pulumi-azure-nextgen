// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNextGen.Security.V20190101Preview.Inputs
{

    /// <summary>
    /// A single automation scope.
    /// </summary>
    public sealed class AutomationScopeArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The resources scope description.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// The resources scope path. Can be the subscription on which the automation is defined on or a resource group under that subscription (fully qualified Azure resource IDs).
        /// </summary>
        [Input("scopePath")]
        public Input<string>? ScopePath { get; set; }

        public AutomationScopeArgs()
        {
        }
    }
}