// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNextGen.DocumentDB.V20160331.Outputs
{

    [OutputType]
    public sealed class CassandraSchemaResponse
    {
        /// <summary>
        /// List of cluster key.
        /// </summary>
        public readonly ImmutableArray<Outputs.ClusterKeyResponse> ClusterKeys;
        /// <summary>
        /// List of Cassandra table columns.
        /// </summary>
        public readonly ImmutableArray<Outputs.ColumnResponse> Columns;
        /// <summary>
        /// List of partition key.
        /// </summary>
        public readonly ImmutableArray<Outputs.CassandraPartitionKeyResponse> PartitionKeys;

        [OutputConstructor]
        private CassandraSchemaResponse(
            ImmutableArray<Outputs.ClusterKeyResponse> clusterKeys,

            ImmutableArray<Outputs.ColumnResponse> columns,

            ImmutableArray<Outputs.CassandraPartitionKeyResponse> partitionKeys)
        {
            ClusterKeys = clusterKeys;
            Columns = columns;
            PartitionKeys = partitionKeys;
        }
    }
}