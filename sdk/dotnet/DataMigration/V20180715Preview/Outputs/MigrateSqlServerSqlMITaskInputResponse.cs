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
    public sealed class MigrateSqlServerSqlMITaskInputResponse
    {
        /// <summary>
        /// Azure Active Directory domain name in the format of 'contoso.com' for federated Azure AD or 'contoso.onmicrosoft.com' for managed domain, required if and only if Windows logins are selected
        /// </summary>
        public readonly string? AadDomainName;
        /// <summary>
        /// SAS URI of Azure Storage Account Container to be used for storing backup files.
        /// </summary>
        public readonly Outputs.BlobShareResponse BackupBlobShare;
        /// <summary>
        /// Backup file share information for all selected databases.
        /// </summary>
        public readonly Outputs.FileShareResponse? BackupFileShare;
        /// <summary>
        /// Backup Mode to specify whether to use existing backup or create new backup. If using existing backups, backup file paths are required to be provided in selectedDatabases.
        /// </summary>
        public readonly string? BackupMode;
        /// <summary>
        /// Agent Jobs to migrate.
        /// </summary>
        public readonly ImmutableArray<string> SelectedAgentJobs;
        /// <summary>
        /// Databases to migrate
        /// </summary>
        public readonly ImmutableArray<Outputs.MigrateSqlServerSqlMIDatabaseInputResponse> SelectedDatabases;
        /// <summary>
        /// Logins to migrate.
        /// </summary>
        public readonly ImmutableArray<string> SelectedLogins;
        /// <summary>
        /// Information for connecting to source
        /// </summary>
        public readonly Outputs.SqlConnectionInfoResponse SourceConnectionInfo;
        /// <summary>
        /// Information for connecting to target
        /// </summary>
        public readonly Outputs.SqlConnectionInfoResponse TargetConnectionInfo;

        [OutputConstructor]
        private MigrateSqlServerSqlMITaskInputResponse(
            string? aadDomainName,

            Outputs.BlobShareResponse backupBlobShare,

            Outputs.FileShareResponse? backupFileShare,

            string? backupMode,

            ImmutableArray<string> selectedAgentJobs,

            ImmutableArray<Outputs.MigrateSqlServerSqlMIDatabaseInputResponse> selectedDatabases,

            ImmutableArray<string> selectedLogins,

            Outputs.SqlConnectionInfoResponse sourceConnectionInfo,

            Outputs.SqlConnectionInfoResponse targetConnectionInfo)
        {
            AadDomainName = aadDomainName;
            BackupBlobShare = backupBlobShare;
            BackupFileShare = backupFileShare;
            BackupMode = backupMode;
            SelectedAgentJobs = selectedAgentJobs;
            SelectedDatabases = selectedDatabases;
            SelectedLogins = selectedLogins;
            SourceConnectionInfo = sourceConnectionInfo;
            TargetConnectionInfo = targetConnectionInfo;
        }
    }
}