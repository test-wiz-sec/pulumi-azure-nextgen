// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../../types/input";
import * as outputs from "../../types/output";
import * as utilities from "../../utilities";

export function getGlobalUserOperationStatus(args: GetGlobalUserOperationStatusArgs, opts?: pulumi.InvokeOptions): Promise<GetGlobalUserOperationStatusResult> {
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    return pulumi.runtime.invoke("azure-nextgen:labservices/latest:getGlobalUserOperationStatus", {
        "operationUrl": args.operationUrl,
        "userName": args.userName,
    }, opts);
}

export interface GetGlobalUserOperationStatusArgs {
    /**
     * The operation url of long running operation
     */
    readonly operationUrl: string;
    /**
     * The name of the user.
     */
    readonly userName: string;
}

/**
 * Status Details of the long running operation for an environment
 */
export interface GetGlobalUserOperationStatusResult {
    /**
     * status of the long running operation for an environment
     */
    readonly status: string;
}
