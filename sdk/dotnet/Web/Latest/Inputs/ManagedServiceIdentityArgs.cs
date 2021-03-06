// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNextGen.Web.Latest.Inputs
{

    /// <summary>
    /// Managed service identity.
    /// </summary>
    public sealed class ManagedServiceIdentityArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Type of managed service identity.
        /// </summary>
        [Input("type")]
        public Input<string>? Type { get; set; }

        public ManagedServiceIdentityArgs()
        {
        }
    }
}
