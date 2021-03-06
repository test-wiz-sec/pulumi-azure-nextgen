// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../../types/input";
import * as outputs from "../../types/output";
import * as utilities from "../../utilities";

export function getManagedDatabase(args: GetManagedDatabaseArgs, opts?: pulumi.InvokeOptions): Promise<GetManagedDatabaseResult> {
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    return pulumi.runtime.invoke("azure-nextgen:sql/v20200202preview:getManagedDatabase", {
        "databaseName": args.databaseName,
        "managedInstanceName": args.managedInstanceName,
        "resourceGroupName": args.resourceGroupName,
    }, opts);
}

export interface GetManagedDatabaseArgs {
    /**
     * The name of the database.
     */
    readonly databaseName: string;
    /**
     * The name of the managed instance.
     */
    readonly managedInstanceName: string;
    /**
     * The name of the resource group. The name is case insensitive.
     */
    readonly resourceGroupName: string;
}

/**
 * A managed database resource.
 */
export interface GetManagedDatabaseResult {
    /**
     * Whether to auto complete restore of this managed database.
     */
    readonly autoCompleteRestore?: boolean;
    /**
     * Collation of the metadata catalog.
     */
    readonly catalogCollation?: string;
    /**
     * Collation of the managed database.
     */
    readonly collation?: string;
    /**
     * Managed database create mode. PointInTimeRestore: Create a database by restoring a point in time backup of an existing database. SourceDatabaseName, SourceManagedInstanceName and PointInTime must be specified. RestoreExternalBackup: Create a database by restoring from external backup files. Collation, StorageContainerUri and StorageContainerSasToken must be specified. Recovery: Creates a database by restoring a geo-replicated backup. RecoverableDatabaseId must be specified as the recoverable database resource ID to restore. RestoreLongTermRetentionBackup: Create a database by restoring from a long term retention backup (longTermRetentionBackupResourceId required).
     */
    readonly createMode?: string;
    /**
     * Creation date of the database.
     */
    readonly creationDate: string;
    /**
     * Geo paired region.
     */
    readonly defaultSecondaryLocation: string;
    /**
     * Earliest restore point in time for point in time restore.
     */
    readonly earliestRestorePoint: string;
    /**
     * Instance Failover Group resource identifier that this managed database belongs to.
     */
    readonly failoverGroupId: string;
    /**
     * Last backup file name for restore of this managed database.
     */
    readonly lastBackupName?: string;
    /**
     * Resource location.
     */
    readonly location: string;
    /**
     * The name of the Long Term Retention backup to be used for restore of this managed database.
     */
    readonly longTermRetentionBackupResourceId?: string;
    /**
     * Resource name.
     */
    readonly name: string;
    /**
     * The resource identifier of the recoverable database associated with create operation of this database.
     */
    readonly recoverableDatabaseId?: string;
    /**
     * The restorable dropped database resource id to restore when creating this database.
     */
    readonly restorableDroppedDatabaseId?: string;
    /**
     * Conditional. If createMode is PointInTimeRestore, this value is required. Specifies the point in time (ISO8601 format) of the source database that will be restored to create the new database.
     */
    readonly restorePointInTime?: string;
    /**
     * The resource identifier of the source database associated with create operation of this database.
     */
    readonly sourceDatabaseId?: string;
    /**
     * Status of the database.
     */
    readonly status: string;
    /**
     * Conditional. If createMode is RestoreExternalBackup, this value is required. Specifies the storage container sas token.
     */
    readonly storageContainerSasToken?: string;
    /**
     * Conditional. If createMode is RestoreExternalBackup, this value is required. Specifies the uri of the storage container where backups for this restore are stored.
     */
    readonly storageContainerUri?: string;
    /**
     * Resource tags.
     */
    readonly tags?: {[key: string]: string};
    /**
     * Resource type.
     */
    readonly type: string;
}
