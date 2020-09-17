// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNextGen.StreamAnalytics.Latest.Outputs
{

    [OutputType]
    public sealed class DiagnosticsResponse
    {
        /// <summary>
        /// A collection of zero or more conditions applicable to the resource, or to the job overall, that warrant customer attention.
        /// </summary>
        public readonly ImmutableArray<Outputs.DiagnosticConditionResponse> Conditions;

        [OutputConstructor]
        private DiagnosticsResponse(ImmutableArray<Outputs.DiagnosticConditionResponse> conditions)
        {
            Conditions = conditions;
        }
    }
}