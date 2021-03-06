// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNextGen.Batch.V20190401.Inputs
{

    public sealed class AutoUserSpecificationArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// nonAdmin - The auto user is a standard user without elevated access. admin - The auto user is a user with elevated access and operates with full Administrator permissions. The default value is nonAdmin.
        /// </summary>
        [Input("elevationLevel")]
        public Input<string>? ElevationLevel { get; set; }

        /// <summary>
        /// The default value is task.
        /// </summary>
        [Input("scope")]
        public Input<string>? Scope { get; set; }

        public AutoUserSpecificationArgs()
        {
        }
    }
}
