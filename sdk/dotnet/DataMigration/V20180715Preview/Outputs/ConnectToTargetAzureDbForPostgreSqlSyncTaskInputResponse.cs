// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNextGen.DataMigration.V20180715Preview.Outputs
{

    [OutputType]
    public sealed class ConnectToTargetAzureDbForPostgreSqlSyncTaskInputResponse
    {
        /// <summary>
        /// Connection information for source PostgreSQL server
        /// </summary>
        public readonly Outputs.PostgreSqlConnectionInfoResponse SourceConnectionInfo;
        /// <summary>
        /// Connection information for target Azure Database for PostgreSQL server
        /// </summary>
        public readonly Outputs.PostgreSqlConnectionInfoResponse TargetConnectionInfo;

        [OutputConstructor]
        private ConnectToTargetAzureDbForPostgreSqlSyncTaskInputResponse(
            Outputs.PostgreSqlConnectionInfoResponse sourceConnectionInfo,

            Outputs.PostgreSqlConnectionInfoResponse targetConnectionInfo)
        {
            SourceConnectionInfo = sourceConnectionInfo;
            TargetConnectionInfo = targetConnectionInfo;
        }
    }
}