// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNextGen.PolicyInsights.V20180701Preview.Outputs
{

    [OutputType]
    public sealed class TypedErrorInfoResponseResult
    {
        /// <summary>
        /// The scenario specific error details.
        /// </summary>
        public readonly ImmutableDictionary<string, object> Info;
        /// <summary>
        /// The type of included error details.
        /// </summary>
        public readonly string Type;

        [OutputConstructor]
        private TypedErrorInfoResponseResult(
            ImmutableDictionary<string, object> info,

            string type)
        {
            Info = info;
            Type = type;
        }
    }
}