// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNextGen.Batch.V20200901.Inputs
{

    public sealed class AutoUserSpecificationArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The default value is nonAdmin.
        /// </summary>
        [Input("elevationLevel")]
        public Input<string>? ElevationLevel { get; set; }

        /// <summary>
        /// The default value is Pool. If the pool is running Windows a value of Task should be specified if stricter isolation between tasks is required. For example, if the task mutates the registry in a way which could impact other tasks, or if certificates have been specified on the pool which should not be accessible by normal tasks but should be accessible by start tasks.
        /// </summary>
        [Input("scope")]
        public Input<string>? Scope { get; set; }

        public AutoUserSpecificationArgs()
        {
        }
    }
}
