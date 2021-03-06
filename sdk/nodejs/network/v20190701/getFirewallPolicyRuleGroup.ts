// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../../types/input";
import * as outputs from "../../types/output";
import * as utilities from "../../utilities";

export function getFirewallPolicyRuleGroup(args: GetFirewallPolicyRuleGroupArgs, opts?: pulumi.InvokeOptions): Promise<GetFirewallPolicyRuleGroupResult> {
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    return pulumi.runtime.invoke("azure-nextgen:network/v20190701:getFirewallPolicyRuleGroup", {
        "firewallPolicyName": args.firewallPolicyName,
        "resourceGroupName": args.resourceGroupName,
        "ruleGroupName": args.ruleGroupName,
    }, opts);
}

export interface GetFirewallPolicyRuleGroupArgs {
    /**
     * The name of the Firewall Policy.
     */
    readonly firewallPolicyName: string;
    /**
     * The name of the resource group.
     */
    readonly resourceGroupName: string;
    /**
     * The name of the FirewallPolicyRuleGroup.
     */
    readonly ruleGroupName: string;
}

/**
 * Rule Group resource.
 */
export interface GetFirewallPolicyRuleGroupResult {
    /**
     * A unique read-only string that changes whenever the resource is updated.
     */
    readonly etag: string;
    /**
     * The name of the resource that is unique within a resource group. This name can be used to access the resource.
     */
    readonly name?: string;
    /**
     * Priority of the Firewall Policy Rule Group resource.
     */
    readonly priority?: number;
    /**
     * The provisioning state of the firewall policy rule group resource.
     */
    readonly provisioningState: string;
    /**
     * Group of Firewall Policy rules.
     */
    readonly rules?: outputs.network.v20190701.FirewallPolicyFilterRuleResponse | outputs.network.v20190701.FirewallPolicyNatRuleResponse[];
    /**
     * Rule Group type.
     */
    readonly type: string;
}
