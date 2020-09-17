// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNextGen.CostManagement.V20191101.Inputs
{

    /// <summary>
    /// The order by expression to be used in the report.
    /// </summary>
    public sealed class ReportConfigSortingArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Direction of sort.
        /// </summary>
        [Input("direction")]
        public Input<string>? Direction { get; set; }

        /// <summary>
        /// The name of the column to sort.
        /// </summary>
        [Input("name", required: true)]
        public Input<string> Name { get; set; } = null!;

        public ReportConfigSortingArgs()
        {
        }
    }
}