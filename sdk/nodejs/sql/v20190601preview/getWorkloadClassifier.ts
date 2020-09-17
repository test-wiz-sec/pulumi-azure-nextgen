// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../../types/input";
import * as outputs from "../../types/output";
import * as utilities from "../../utilities";

export function getWorkloadClassifier(args: GetWorkloadClassifierArgs, opts?: pulumi.InvokeOptions): Promise<GetWorkloadClassifierResult> {
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    return pulumi.runtime.invoke("azure-nextgen:sql/v20190601preview:getWorkloadClassifier", {
        "databaseName": args.databaseName,
        "resourceGroupName": args.resourceGroupName,
        "serverName": args.serverName,
        "workloadClassifierName": args.workloadClassifierName,
        "workloadGroupName": args.workloadGroupName,
    }, opts);
}

export interface GetWorkloadClassifierArgs {
    /**
     * The name of the database.
     */
    readonly databaseName: string;
    /**
     * The name of the resource group that contains the resource. You can obtain this value from the Azure Resource Manager API or the portal.
     */
    readonly resourceGroupName: string;
    /**
     * The name of the server.
     */
    readonly serverName: string;
    /**
     * The name of the workload classifier.
     */
    readonly workloadClassifierName: string;
    /**
     * The name of the workload group from which to receive the classifier from.
     */
    readonly workloadGroupName: string;
}

/**
 * Workload classifier operations for a data warehouse
 */
export interface GetWorkloadClassifierResult {
    /**
     * The workload classifier context.
     */
    readonly context?: string;
    /**
     * The workload classifier end time for classification.
     */
    readonly endTime?: string;
    /**
     * The workload classifier importance.
     */
    readonly importance?: string;
    /**
     * The workload classifier label.
     */
    readonly label?: string;
    /**
     * The workload classifier member name.
     */
    readonly memberName: string;
    /**
     * Resource name.
     */
    readonly name: string;
    /**
     * The workload classifier start time for classification.
     */
    readonly startTime?: string;
    /**
     * Resource type.
     */
    readonly type: string;
}