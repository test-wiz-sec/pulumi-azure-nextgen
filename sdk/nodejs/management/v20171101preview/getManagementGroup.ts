// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../../types/input";
import * as outputs from "../../types/output";
import * as utilities from "../../utilities";

export function getManagementGroup(args: GetManagementGroupArgs, opts?: pulumi.InvokeOptions): Promise<GetManagementGroupResult> {
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    return pulumi.runtime.invoke("azure-nextgen:management/v20171101preview:getManagementGroup", {
        "expand": args.expand,
        "groupId": args.groupId,
        "recurse": args.recurse,
    }, opts);
}

export interface GetManagementGroupArgs {
    /**
     * The $expand=children query string parameter allows clients to request inclusion of children in the response payload.
     */
    readonly expand?: string;
    /**
     * Management Group ID.
     */
    readonly groupId: string;
    /**
     * The $recurse=true query string parameter allows clients to request inclusion of entire hierarchy in the response payload.
     */
    readonly recurse?: boolean;
}

/**
 * The management group details.
 */
export interface GetManagementGroupResult {
    /**
     * The list of children.
     */
    readonly children?: outputs.management.v20171101preview.ManagementGroupChildInfoResponse[];
    /**
     * The details of a management group.
     */
    readonly details?: outputs.management.v20171101preview.ManagementGroupDetailsResponse;
    /**
     * The friendly name of the management group.
     */
    readonly displayName?: string;
    /**
     * The name of the management group. For example, 00000000-0000-0000-0000-000000000000
     */
    readonly name: string;
    /**
     * The AAD Tenant ID associated with the management group. For example, 00000000-0000-0000-0000-000000000000
     */
    readonly tenantId?: string;
    /**
     * The type of the resource.  For example, /providers/Microsoft.Management/managementGroups
     */
    readonly type: string;
}
