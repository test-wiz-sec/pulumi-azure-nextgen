// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNextGen.Sql.V20200801Preview
{
    public static class GetDatabase
    {
        public static Task<GetDatabaseResult> InvokeAsync(GetDatabaseArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetDatabaseResult>("azure-nextgen:sql/v20200801preview:getDatabase", args ?? new GetDatabaseArgs(), options.WithVersion());
    }


    public sealed class GetDatabaseArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// The name of the database.
        /// </summary>
        [Input("databaseName", required: true)]
        public string DatabaseName { get; set; } = null!;

        /// <summary>
        /// The name of the resource group that contains the resource. You can obtain this value from the Azure Resource Manager API or the portal.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public string ResourceGroupName { get; set; } = null!;

        /// <summary>
        /// The name of the server.
        /// </summary>
        [Input("serverName", required: true)]
        public string ServerName { get; set; } = null!;

        public GetDatabaseArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetDatabaseResult
    {
        /// <summary>
        /// Time in minutes after which database is automatically paused. A value of -1 means that automatic pause is disabled
        /// </summary>
        public readonly int? AutoPauseDelay;
        /// <summary>
        /// Collation of the metadata catalog.
        /// </summary>
        public readonly string? CatalogCollation;
        /// <summary>
        /// The collation of the database.
        /// </summary>
        public readonly string? Collation;
        /// <summary>
        /// Specifies the mode of database creation.
        /// 
        /// Default: regular database creation.
        /// 
        /// Copy: creates a database as a copy of an existing database. sourceDatabaseId must be specified as the resource ID of the source database.
        /// 
        /// Secondary: creates a database as a secondary replica of an existing database. sourceDatabaseId must be specified as the resource ID of the existing primary database.
        /// 
        /// PointInTimeRestore: Creates a database by restoring a point in time backup of an existing database. sourceDatabaseId must be specified as the resource ID of the existing database, and restorePointInTime must be specified.
        /// 
        /// Recovery: Creates a database by restoring a geo-replicated backup. sourceDatabaseId must be specified as the recoverable database resource ID to restore.
        /// 
        /// Restore: Creates a database by restoring a backup of a deleted database. sourceDatabaseId must be specified. If sourceDatabaseId is the database's original resource ID, then sourceDatabaseDeletionDate must be specified. Otherwise sourceDatabaseId must be the restorable dropped database resource ID and sourceDatabaseDeletionDate is ignored. restorePointInTime may also be specified to restore from an earlier point in time.
        /// 
        /// RestoreLongTermRetentionBackup: Creates a database by restoring from a long term retention vault. recoveryServicesRecoveryPointResourceId must be specified as the recovery point resource ID.
        /// 
        /// Copy, Secondary, and RestoreLongTermRetentionBackup are not supported for DataWarehouse edition.
        /// </summary>
        public readonly string? CreateMode;
        /// <summary>
        /// The creation date of the database (ISO8601 format).
        /// </summary>
        public readonly string CreationDate;
        /// <summary>
        /// The current service level objective name of the database.
        /// </summary>
        public readonly string CurrentServiceObjectiveName;
        /// <summary>
        /// The name and tier of the SKU.
        /// </summary>
        public readonly Outputs.SkuResponse CurrentSku;
        /// <summary>
        /// The ID of the database.
        /// </summary>
        public readonly string DatabaseId;
        /// <summary>
        /// The default secondary region for this database.
        /// </summary>
        public readonly string DefaultSecondaryLocation;
        /// <summary>
        /// This records the earliest start date and time that restore is available for this database (ISO8601 format).
        /// </summary>
        public readonly string EarliestRestoreDate;
        /// <summary>
        /// The resource identifier of the elastic pool containing this database.
        /// </summary>
        public readonly string? ElasticPoolId;
        /// <summary>
        /// Failover Group resource identifier that this database belongs to.
        /// </summary>
        public readonly string FailoverGroupId;
        /// <summary>
        /// The number of secondary replicas associated with the database that are used to provide high availability.
        /// </summary>
        public readonly int? HighAvailabilityReplicaCount;
        /// <summary>
        /// Kind of database. This is metadata used for the Azure portal experience.
        /// </summary>
        public readonly string Kind;
        /// <summary>
        /// The license type to apply for this database. `LicenseIncluded` if you need a license, or `BasePrice` if you have a license and are eligible for the Azure Hybrid Benefit.
        /// </summary>
        public readonly string? LicenseType;
        /// <summary>
        /// Resource location.
        /// </summary>
        public readonly string Location;
        /// <summary>
        /// The resource identifier of the long term retention backup associated with create operation of this database.
        /// </summary>
        public readonly string? LongTermRetentionBackupResourceId;
        /// <summary>
        /// Maintenance configuration id assigned to the database. This configuration defines the period when the maintenance updates will be rolled out.
        /// </summary>
        public readonly string? MaintenanceConfigurationId;
        /// <summary>
        /// Resource that manages the database.
        /// </summary>
        public readonly string ManagedBy;
        /// <summary>
        /// The max log size for this database.
        /// </summary>
        public readonly int MaxLogSizeBytes;
        /// <summary>
        /// The max size of the database expressed in bytes.
        /// </summary>
        public readonly int? MaxSizeBytes;
        /// <summary>
        /// Minimal capacity that database will always have allocated, if not paused
        /// </summary>
        public readonly double? MinCapacity;
        /// <summary>
        /// Resource name.
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// The date when database was paused by user configuration or action(ISO8601 format). Null if the database is ready.
        /// </summary>
        public readonly string PausedDate;
        /// <summary>
        /// The state of read-only routing. If enabled, connections that have application intent set to readonly in their connection string may be routed to a readonly secondary replica in the same region.
        /// </summary>
        public readonly string? ReadScale;
        /// <summary>
        /// The resource identifier of the recoverable database associated with create operation of this database.
        /// </summary>
        public readonly string? RecoverableDatabaseId;
        /// <summary>
        /// The resource identifier of the recovery point associated with create operation of this database.
        /// </summary>
        public readonly string? RecoveryServicesRecoveryPointId;
        /// <summary>
        /// The requested service level objective name of the database.
        /// </summary>
        public readonly string RequestedServiceObjectiveName;
        /// <summary>
        /// The resource identifier of the restorable dropped database associated with create operation of this database.
        /// </summary>
        public readonly string? RestorableDroppedDatabaseId;
        /// <summary>
        /// Specifies the point in time (ISO8601 format) of the source database that will be restored to create the new database.
        /// </summary>
        public readonly string? RestorePointInTime;
        /// <summary>
        /// The date when database was resumed by user action or database login (ISO8601 format). Null if the database is paused.
        /// </summary>
        public readonly string ResumedDate;
        /// <summary>
        /// The name of the sample schema to apply when creating this database.
        /// </summary>
        public readonly string? SampleName;
        /// <summary>
        /// The secondary type of the database if it is a secondary.  Valid values are Geo and Named.
        /// </summary>
        public readonly string? SecondaryType;
        /// <summary>
        /// The database SKU.
        /// 
        /// The list of SKUs may vary by region and support offer. To determine the SKUs (including the SKU name, tier/edition, family, and capacity) that are available to your subscription in an Azure region, use the `Capabilities_ListByLocation` REST API or one of the following commands:
        /// 
        /// ```azurecli
        /// az sql db list-editions -l &lt;location&gt; -o table
        /// ````
        /// 
        /// ```powershell
        /// Get-AzSqlServerServiceObjective -Location &lt;location&gt;
        /// ````
        /// </summary>
        public readonly Outputs.SkuResponse? Sku;
        /// <summary>
        /// Specifies the time that the database was deleted.
        /// </summary>
        public readonly string? SourceDatabaseDeletionDate;
        /// <summary>
        /// The resource identifier of the source database associated with create operation of this database.
        /// </summary>
        public readonly string? SourceDatabaseId;
        /// <summary>
        /// The status of the database.
        /// </summary>
        public readonly string Status;
        /// <summary>
        /// The storage account type used to store backups for this database. Currently the only supported option is GRS (GeoRedundantStorage).
        /// </summary>
        public readonly string? StorageAccountType;
        /// <summary>
        /// Resource tags.
        /// </summary>
        public readonly ImmutableDictionary<string, string>? Tags;
        /// <summary>
        /// Resource type.
        /// </summary>
        public readonly string Type;
        /// <summary>
        /// Whether or not this database is zone redundant, which means the replicas of this database will be spread across multiple availability zones.
        /// </summary>
        public readonly bool? ZoneRedundant;

        [OutputConstructor]
        private GetDatabaseResult(
            int? autoPauseDelay,

            string? catalogCollation,

            string? collation,

            string? createMode,

            string creationDate,

            string currentServiceObjectiveName,

            Outputs.SkuResponse currentSku,

            string databaseId,

            string defaultSecondaryLocation,

            string earliestRestoreDate,

            string? elasticPoolId,

            string failoverGroupId,

            int? highAvailabilityReplicaCount,

            string kind,

            string? licenseType,

            string location,

            string? longTermRetentionBackupResourceId,

            string? maintenanceConfigurationId,

            string managedBy,

            int maxLogSizeBytes,

            int? maxSizeBytes,

            double? minCapacity,

            string name,

            string pausedDate,

            string? readScale,

            string? recoverableDatabaseId,

            string? recoveryServicesRecoveryPointId,

            string requestedServiceObjectiveName,

            string? restorableDroppedDatabaseId,

            string? restorePointInTime,

            string resumedDate,

            string? sampleName,

            string? secondaryType,

            Outputs.SkuResponse? sku,

            string? sourceDatabaseDeletionDate,

            string? sourceDatabaseId,

            string status,

            string? storageAccountType,

            ImmutableDictionary<string, string>? tags,

            string type,

            bool? zoneRedundant)
        {
            AutoPauseDelay = autoPauseDelay;
            CatalogCollation = catalogCollation;
            Collation = collation;
            CreateMode = createMode;
            CreationDate = creationDate;
            CurrentServiceObjectiveName = currentServiceObjectiveName;
            CurrentSku = currentSku;
            DatabaseId = databaseId;
            DefaultSecondaryLocation = defaultSecondaryLocation;
            EarliestRestoreDate = earliestRestoreDate;
            ElasticPoolId = elasticPoolId;
            FailoverGroupId = failoverGroupId;
            HighAvailabilityReplicaCount = highAvailabilityReplicaCount;
            Kind = kind;
            LicenseType = licenseType;
            Location = location;
            LongTermRetentionBackupResourceId = longTermRetentionBackupResourceId;
            MaintenanceConfigurationId = maintenanceConfigurationId;
            ManagedBy = managedBy;
            MaxLogSizeBytes = maxLogSizeBytes;
            MaxSizeBytes = maxSizeBytes;
            MinCapacity = minCapacity;
            Name = name;
            PausedDate = pausedDate;
            ReadScale = readScale;
            RecoverableDatabaseId = recoverableDatabaseId;
            RecoveryServicesRecoveryPointId = recoveryServicesRecoveryPointId;
            RequestedServiceObjectiveName = requestedServiceObjectiveName;
            RestorableDroppedDatabaseId = restorableDroppedDatabaseId;
            RestorePointInTime = restorePointInTime;
            ResumedDate = resumedDate;
            SampleName = sampleName;
            SecondaryType = secondaryType;
            Sku = sku;
            SourceDatabaseDeletionDate = sourceDatabaseDeletionDate;
            SourceDatabaseId = sourceDatabaseId;
            Status = status;
            StorageAccountType = storageAccountType;
            Tags = tags;
            Type = type;
            ZoneRedundant = zoneRedundant;
        }
    }
}