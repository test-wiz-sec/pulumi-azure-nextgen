// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../../types/input";
import * as outputs from "../../types/output";
import * as utilities from "../../utilities";

export function getManagementLockByScope(args: GetManagementLockByScopeArgs, opts?: pulumi.InvokeOptions): Promise<GetManagementLockByScopeResult> {
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    return pulumi.runtime.invoke("azure-nextgen:authorization/v20160901:getManagementLockByScope", {
        "lockName": args.lockName,
        "scope": args.scope,
    }, opts);
}

export interface GetManagementLockByScopeArgs {
    /**
     * The name of lock.
     */
    readonly lockName: string;
    /**
     * The scope for the lock. 
     */
    readonly scope: string;
}

/**
 * The lock information.
 */
export interface GetManagementLockByScopeResult {
    /**
     * The level of the lock. Possible values are: NotSpecified, CanNotDelete, ReadOnly. CanNotDelete means authorized users are able to read and modify the resources, but not delete. ReadOnly means authorized users can only read from a resource, but they can't modify or delete it.
     */
    readonly level: string;
    /**
     * The name of the lock.
     */
    readonly name: string;
    /**
     * Notes about the lock. Maximum of 512 characters.
     */
    readonly notes?: string;
    /**
     * The owners of the lock.
     */
    readonly owners?: outputs.authorization.v20160901.ManagementLockOwnerResponse[];
    /**
     * The resource type of the lock - Microsoft.Authorization/locks.
     */
    readonly type: string;
}
