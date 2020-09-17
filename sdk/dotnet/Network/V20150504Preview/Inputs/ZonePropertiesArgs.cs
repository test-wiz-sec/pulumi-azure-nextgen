// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNextGen.Network.V20150504Preview.Inputs
{

    /// <summary>
    /// Represents the properties of the zone.
    /// </summary>
    public sealed class ZonePropertiesArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Gets or sets the maximum number of record sets that can be created in this zone.
        /// </summary>
        [Input("maxNumberOfRecordSets")]
        public Input<int>? MaxNumberOfRecordSets { get; set; }

        /// <summary>
        /// Gets or sets the current number of record sets in this zone.
        /// </summary>
        [Input("numberOfRecordSets")]
        public Input<int>? NumberOfRecordSets { get; set; }

        public ZonePropertiesArgs()
        {
        }
    }
}