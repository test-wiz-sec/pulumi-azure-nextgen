// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNextGen.Scheduler.V20140801Preview.Inputs
{

    public sealed class JobMaxRecurrenceArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Gets or sets the frequency of recurrence (second, minute, hour, day, week, month).
        /// </summary>
        [Input("frequency")]
        public Input<string>? Frequency { get; set; }

        /// <summary>
        /// Gets or sets the interval between retries.
        /// </summary>
        [Input("interval")]
        public Input<int>? Interval { get; set; }

        public JobMaxRecurrenceArgs()
        {
        }
    }
}