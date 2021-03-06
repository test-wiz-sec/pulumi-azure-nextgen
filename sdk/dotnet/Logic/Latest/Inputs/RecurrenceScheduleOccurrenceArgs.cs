// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNextGen.Logic.Latest.Inputs
{

    /// <summary>
    /// The recurrence schedule occurrence.
    /// </summary>
    public sealed class RecurrenceScheduleOccurrenceArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The day of the week.
        /// </summary>
        [Input("day")]
        public Input<string>? Day { get; set; }

        /// <summary>
        /// The occurrence.
        /// </summary>
        [Input("occurrence")]
        public Input<int>? Occurrence { get; set; }

        public RecurrenceScheduleOccurrenceArgs()
        {
        }
    }
}
