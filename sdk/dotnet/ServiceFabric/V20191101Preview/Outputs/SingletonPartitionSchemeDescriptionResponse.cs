// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNextGen.ServiceFabric.V20191101Preview.Outputs
{

    [OutputType]
    public sealed class SingletonPartitionSchemeDescriptionResponse
    {
        /// <summary>
        /// Specifies how the service is partitioned.
        /// </summary>
        public readonly string PartitionScheme;

        [OutputConstructor]
        private SingletonPartitionSchemeDescriptionResponse(string partitionScheme)
        {
            PartitionScheme = partitionScheme;
        }
    }
}
