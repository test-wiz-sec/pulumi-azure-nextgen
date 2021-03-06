// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNextGen.Web.V20150801.Inputs
{

    /// <summary>
    /// Class containing stamp capacity information
    /// </summary>
    public sealed class StampCapacityArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Available capacity (# of machines, bytes of storage etc...)
        /// </summary>
        [Input("availableCapacity")]
        public Input<int>? AvailableCapacity { get; set; }

        /// <summary>
        /// Shared/Dedicated workers
        /// </summary>
        [Input("computeMode")]
        public Input<string>? ComputeMode { get; set; }

        /// <summary>
        /// If true it includes basic sites
        ///             Basic sites are not used for capacity allocation.
        /// </summary>
        [Input("excludeFromCapacityAllocation")]
        public Input<bool>? ExcludeFromCapacityAllocation { get; set; }

        /// <summary>
        /// Is capacity applicable for all sites?
        /// </summary>
        [Input("isApplicableForAllComputeModes")]
        public Input<bool>? IsApplicableForAllComputeModes { get; set; }

        /// <summary>
        /// Name of the stamp
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// Shared or Dedicated
        /// </summary>
        [Input("siteMode")]
        public Input<string>? SiteMode { get; set; }

        /// <summary>
        /// Total capacity (# of machines, bytes of storage etc...)
        /// </summary>
        [Input("totalCapacity")]
        public Input<int>? TotalCapacity { get; set; }

        /// <summary>
        /// Name of the unit
        /// </summary>
        [Input("unit")]
        public Input<string>? Unit { get; set; }

        /// <summary>
        /// Size of the machines
        /// </summary>
        [Input("workerSize")]
        public Input<string>? WorkerSize { get; set; }

        /// <summary>
        /// Size Id of machines: 
        ///             0 - Small
        ///             1 - Medium
        ///             2 - Large
        /// </summary>
        [Input("workerSizeId")]
        public Input<int>? WorkerSizeId { get; set; }

        public StampCapacityArgs()
        {
        }
    }
}
