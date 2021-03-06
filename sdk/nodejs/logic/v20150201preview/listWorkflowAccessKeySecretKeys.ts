// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../../types/input";
import * as outputs from "../../types/output";
import * as utilities from "../../utilities";

export function listWorkflowAccessKeySecretKeys(args: ListWorkflowAccessKeySecretKeysArgs, opts?: pulumi.InvokeOptions): Promise<ListWorkflowAccessKeySecretKeysResult> {
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    return pulumi.runtime.invoke("azure-nextgen:logic/v20150201preview:listWorkflowAccessKeySecretKeys", {
        "accessKeyName": args.accessKeyName,
        "resourceGroupName": args.resourceGroupName,
        "workflowName": args.workflowName,
    }, opts);
}

export interface ListWorkflowAccessKeySecretKeysArgs {
    /**
     * The workflow access key name.
     */
    readonly accessKeyName: string;
    /**
     * The resource group name.
     */
    readonly resourceGroupName: string;
    /**
     * The workflow name.
     */
    readonly workflowName: string;
}

export interface ListWorkflowAccessKeySecretKeysResult {
    /**
     * Gets the primary secret key.
     */
    readonly primarySecretKey: string;
    /**
     * Gets the secondary secret key.
     */
    readonly secondarySecretKey: string;
}
