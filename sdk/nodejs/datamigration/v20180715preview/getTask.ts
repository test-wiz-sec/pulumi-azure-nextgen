// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../../types/input";
import * as outputs from "../../types/output";
import * as utilities from "../../utilities";

export function getTask(args: GetTaskArgs, opts?: pulumi.InvokeOptions): Promise<GetTaskResult> {
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    return pulumi.runtime.invoke("azure-nextgen:datamigration/v20180715preview:getTask", {
        "expand": args.expand,
        "groupName": args.groupName,
        "projectName": args.projectName,
        "serviceName": args.serviceName,
        "taskName": args.taskName,
    }, opts);
}

export interface GetTaskArgs {
    /**
     * Expand the response
     */
    readonly expand?: string;
    /**
     * Name of the resource group
     */
    readonly groupName: string;
    /**
     * Name of the project
     */
    readonly projectName: string;
    /**
     * Name of the service
     */
    readonly serviceName: string;
    /**
     * Name of the Task
     */
    readonly taskName: string;
}

/**
 * A task resource
 */
export interface GetTaskResult {
    /**
     * HTTP strong entity tag value. This is ignored if submitted.
     */
    readonly etag?: string;
    /**
     * Resource name.
     */
    readonly name: string;
    /**
     * Custom task properties
     */
    readonly properties: outputs.datamigration.v20180715preview.ConnectToMongoDbTaskPropertiesResponse | outputs.datamigration.v20180715preview.ConnectToSourceOracleSyncTaskPropertiesResponse | outputs.datamigration.v20180715preview.ConnectToSourcePostgreSqlSyncTaskPropertiesResponse | outputs.datamigration.v20180715preview.ConnectToSourceSqlServerSyncTaskPropertiesResponse | outputs.datamigration.v20180715preview.ConnectToSourceSqlServerTaskPropertiesResponse | outputs.datamigration.v20180715preview.ConnectToTargetAzureDbForMySqlTaskPropertiesResponse | outputs.datamigration.v20180715preview.ConnectToTargetAzureDbForPostgreSqlSyncTaskPropertiesResponse | outputs.datamigration.v20180715preview.ConnectToTargetOracleAzureDbForPostgreSqlSyncTaskPropertiesResponse | outputs.datamigration.v20180715preview.ConnectToTargetSqlDbTaskPropertiesResponse | outputs.datamigration.v20180715preview.ConnectToTargetSqlMISyncTaskPropertiesResponse | outputs.datamigration.v20180715preview.ConnectToTargetSqlMITaskPropertiesResponse | outputs.datamigration.v20180715preview.ConnectToTargetSqlSqlDbSyncTaskPropertiesResponse | outputs.datamigration.v20180715preview.GetTdeCertificatesSqlTaskPropertiesResponse | outputs.datamigration.v20180715preview.GetUserTablesOracleTaskPropertiesResponse | outputs.datamigration.v20180715preview.GetUserTablesPostgreSqlTaskPropertiesResponse | outputs.datamigration.v20180715preview.GetUserTablesSqlSyncTaskPropertiesResponse | outputs.datamigration.v20180715preview.GetUserTablesSqlTaskPropertiesResponse | outputs.datamigration.v20180715preview.MigrateMongoDbTaskPropertiesResponse | outputs.datamigration.v20180715preview.MigrateMySqlAzureDbForMySqlSyncTaskPropertiesResponse | outputs.datamigration.v20180715preview.MigrateOracleAzureDbForPostgreSqlSyncTaskPropertiesResponse | outputs.datamigration.v20180715preview.MigratePostgreSqlAzureDbForPostgreSqlSyncTaskPropertiesResponse | outputs.datamigration.v20180715preview.MigrateSqlServerSqlDbSyncTaskPropertiesResponse | outputs.datamigration.v20180715preview.MigrateSqlServerSqlDbTaskPropertiesResponse | outputs.datamigration.v20180715preview.MigrateSqlServerSqlMISyncTaskPropertiesResponse | outputs.datamigration.v20180715preview.MigrateSqlServerSqlMITaskPropertiesResponse | outputs.datamigration.v20180715preview.MigrateSsisTaskPropertiesResponse | outputs.datamigration.v20180715preview.ValidateMigrationInputSqlServerSqlDbSyncTaskPropertiesResponse | outputs.datamigration.v20180715preview.ValidateMigrationInputSqlServerSqlMISyncTaskPropertiesResponse | outputs.datamigration.v20180715preview.ValidateMigrationInputSqlServerSqlMITaskPropertiesResponse | outputs.datamigration.v20180715preview.ValidateMongoDbTaskPropertiesResponse | outputs.datamigration.v20180715preview.ValidateOracleAzureDbForPostgreSqlSyncTaskPropertiesResponse;
    /**
     * Resource type.
     */
    readonly type: string;
}
