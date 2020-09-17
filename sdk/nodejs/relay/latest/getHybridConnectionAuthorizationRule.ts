// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../../types/input";
import * as outputs from "../../types/output";
import * as utilities from "../../utilities";

export function getHybridConnectionAuthorizationRule(args: GetHybridConnectionAuthorizationRuleArgs, opts?: pulumi.InvokeOptions): Promise<GetHybridConnectionAuthorizationRuleResult> {
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    return pulumi.runtime.invoke("azure-nextgen:relay/latest:getHybridConnectionAuthorizationRule", {
        "authorizationRuleName": args.authorizationRuleName,
        "hybridConnectionName": args.hybridConnectionName,
        "namespaceName": args.namespaceName,
        "resourceGroupName": args.resourceGroupName,
    }, opts);
}

export interface GetHybridConnectionAuthorizationRuleArgs {
    /**
     * The authorization rule name.
     */
    readonly authorizationRuleName: string;
    /**
     * The hybrid connection name.
     */
    readonly hybridConnectionName: string;
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
 * Description of a namespace authorization rule.
 */
export interface GetHybridConnectionAuthorizationRuleResult {
    /**
     * Resource name.
     */
    readonly name: string;
    /**
     * The rights associated with the rule.
     */
    readonly rights: string[];
    /**
     * Resource type.
     */
    readonly type: string;
}