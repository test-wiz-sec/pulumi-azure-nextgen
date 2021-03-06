// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../../types/input";
import * as outputs from "../../types/output";
import * as utilities from "../../utilities";

export function getLinkedServer(args: GetLinkedServerArgs, opts?: pulumi.InvokeOptions): Promise<GetLinkedServerResult> {
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    return pulumi.runtime.invoke("azure-nextgen:cache/latest:getLinkedServer", {
        "linkedServerName": args.linkedServerName,
        "name": args.name,
        "resourceGroupName": args.resourceGroupName,
    }, opts);
}

export interface GetLinkedServerArgs {
    /**
     * The name of the linked server.
     */
    readonly linkedServerName: string;
    /**
     * The name of the redis cache.
     */
    readonly name: string;
    /**
     * The name of the resource group.
     */
    readonly resourceGroupName: string;
}

/**
 * Response to put/get linked server (with properties) for Redis cache.
 */
export interface GetLinkedServerResult {
    /**
     * Fully qualified resourceId of the linked redis cache.
     */
    readonly linkedRedisCacheId: string;
    /**
     * Location of the linked redis cache.
     */
    readonly linkedRedisCacheLocation: string;
    /**
     * Resource name.
     */
    readonly name: string;
    /**
     * Terminal state of the link between primary and secondary redis cache.
     */
    readonly provisioningState: string;
    /**
     * Role of the linked server.
     */
    readonly serverRole: string;
    /**
     * Resource type.
     */
    readonly type: string;
}
