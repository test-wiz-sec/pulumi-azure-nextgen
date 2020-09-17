// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNextGen.DBforMariaDB.V20180601Preview.Outputs
{

    [OutputType]
    public sealed class SkuResponse
    {
        /// <summary>
        /// The scale up/out capacity, representing server's compute units.
        /// </summary>
        public readonly int? Capacity;
        /// <summary>
        /// The family of hardware.
        /// </summary>
        public readonly string? Family;
        /// <summary>
        /// The name of the sku, typically, tier + family + cores, e.g. B_Gen4_1, GP_Gen5_8.
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// The size code, to be interpreted by resource as appropriate.
        /// </summary>
        public readonly string? Size;
        /// <summary>
        /// The tier of the particular SKU, e.g. Basic.
        /// </summary>
        public readonly string? Tier;

        [OutputConstructor]
        private SkuResponse(
            int? capacity,

            string? family,

            string name,

            string? size,

            string? tier)
        {
            Capacity = capacity;
            Family = family;
            Name = name;
            Size = size;
            Tier = tier;
        }
    }
}
