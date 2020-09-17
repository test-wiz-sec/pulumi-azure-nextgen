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
    public sealed class GetUserTablesPostgreSqlTaskOutputResponse
    {
        /// <summary>
        /// The database this result is for
        /// </summary>
        public readonly string DatabaseName;
        /// <summary>
        /// List of valid tables found for this database
        /// </summary>
        public readonly ImmutableArray<Outputs.DatabaseTableResponse> Tables;
        /// <summary>
        /// Validation errors associated with the task
        /// </summary>
        public readonly ImmutableArray<Outputs.ReportableExceptionResponse> ValidationErrors;

        [OutputConstructor]
        private GetUserTablesPostgreSqlTaskOutputResponse(
            string databaseName,

            ImmutableArray<Outputs.DatabaseTableResponse> tables,

            ImmutableArray<Outputs.ReportableExceptionResponse> validationErrors)
        {
            DatabaseName = databaseName;
            Tables = tables;
            ValidationErrors = validationErrors;
        }
    }
}