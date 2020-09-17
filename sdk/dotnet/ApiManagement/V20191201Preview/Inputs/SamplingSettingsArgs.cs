// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNextGen.ApiManagement.V20191201Preview.Inputs
{

    /// <summary>
    /// Sampling settings for Diagnostic.
    /// </summary>
    public sealed class SamplingSettingsArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Rate of sampling for fixed-rate sampling.
        /// </summary>
        [Input("percentage")]
        public Input<double>? Percentage { get; set; }

        /// <summary>
        /// Sampling type.
        /// </summary>
        [Input("samplingType")]
        public Input<string>? SamplingType { get; set; }

        public SamplingSettingsArgs()
        {
        }
    }
}