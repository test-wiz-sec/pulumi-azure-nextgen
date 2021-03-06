// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../../types/input";
import * as outputs from "../../types/output";
import * as utilities from "../../utilities";

export function getNamespaceIpFilterRule(args: GetNamespaceIpFilterRuleArgs, opts?: pulumi.InvokeOptions): Promise<GetNamespaceIpFilterRuleResult> {
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    return pulumi.runtime.invoke("azure-nextgen:servicebus/v20180101preview:getNamespaceIpFilterRule", {
        "ipFilterRuleName": args.ipFilterRuleName,
        "namespaceName": args.namespaceName,
        "resourceGroupName": args.resourceGroupName,
    }, opts);
}

export interface GetNamespaceIpFilterRuleArgs {
    /**
     * The IP Filter Rule name.
     */
    readonly ipFilterRuleName: string;
    /**
     * The namespace name
     */
    readonly namespaceName: string;
    /**
     * Name of the Resource group within the Azure subscription.
     */
    readonly resourceGroupName: string;
}

/**
 * Single item in a List or Get IpFilterRules operation
 */
export interface GetNamespaceIpFilterRuleResult {
    /**
     * The IP Filter Action
     */
    readonly action?: string;
    /**
     * IP Filter name
     */
    readonly filterName?: string;
    /**
     * IP Mask
     */
    readonly ipMask?: string;
    /**
     * Resource name
     */
    readonly name: string;
    /**
     * Resource type
     */
    readonly type: string;
}
