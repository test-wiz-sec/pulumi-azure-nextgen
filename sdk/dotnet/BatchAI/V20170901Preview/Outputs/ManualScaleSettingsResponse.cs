// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNextGen.BatchAI.V20170901Preview.Outputs
{

    [OutputType]
    public sealed class ManualScaleSettingsResponse
    {
        /// <summary>
        /// The default value is requeue.
        /// </summary>
        public readonly string? NodeDeallocationOption;
        /// <summary>
        /// Default is 0. If autoScaleSettings are not specified, then the Cluster starts with this target.
        /// </summary>
        public readonly int TargetNodeCount;

        [OutputConstructor]
        private ManualScaleSettingsResponse(
            string? nodeDeallocationOption,

            int targetNodeCount)
        {
            NodeDeallocationOption = nodeDeallocationOption;
            TargetNodeCount = targetNodeCount;
        }
    }
}
