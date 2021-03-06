// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../../types/input";
import * as outputs from "../../types/output";
import * as utilities from "../../utilities";

export function getStorageInsightConfig(args: GetStorageInsightConfigArgs, opts?: pulumi.InvokeOptions): Promise<GetStorageInsightConfigResult> {
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    return pulumi.runtime.invoke("azure-nextgen:operationalinsights/latest:getStorageInsightConfig", {
        "resourceGroupName": args.resourceGroupName,
        "storageInsightName": args.storageInsightName,
        "workspaceName": args.workspaceName,
    }, opts);
}

export interface GetStorageInsightConfigArgs {
    /**
     * The name of the resource group. The name is case insensitive.
     */
    readonly resourceGroupName: string;
    /**
     * Name of the storageInsightsConfigs resource
     */
    readonly storageInsightName: string;
    /**
     * The name of the workspace.
     */
    readonly workspaceName: string;
}

/**
 * The top level storage insight resource container.
 */
export interface GetStorageInsightConfigResult {
    /**
     * The names of the blob containers that the workspace should read
     */
    readonly containers?: string[];
    /**
     * The ETag of the storage insight.
     */
    readonly eTag?: string;
    /**
     * The name of the resource
     */
    readonly name: string;
    /**
     * The status of the storage insight
     */
    readonly status: outputs.operationalinsights.latest.StorageInsightStatusResponse;
    /**
     * The storage account connection details
     */
    readonly storageAccount: outputs.operationalinsights.latest.StorageAccountResponse;
    /**
     * The names of the Azure tables that the workspace should read
     */
    readonly tables?: string[];
    /**
     * Resource tags.
     */
    readonly tags?: {[key: string]: string};
    /**
     * The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
     */
    readonly type: string;
}
