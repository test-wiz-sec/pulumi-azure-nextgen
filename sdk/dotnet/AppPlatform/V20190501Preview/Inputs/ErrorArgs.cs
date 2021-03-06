// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNextGen.AppPlatform.V20190501Preview.Inputs
{

    /// <summary>
    /// The error code compose of code and message.
    /// </summary>
    public sealed class ErrorArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The code of error.
        /// </summary>
        [Input("code")]
        public Input<string>? Code { get; set; }

        /// <summary>
        /// The message of error.
        /// </summary>
        [Input("message")]
        public Input<string>? Message { get; set; }

        public ErrorArgs()
        {
        }
    }
}
