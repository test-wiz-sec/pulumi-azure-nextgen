// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../../types/input";
import * as outputs from "../../types/output";
import * as utilities from "../../utilities";

export function getEnterprisePolicy(args: GetEnterprisePolicyArgs, opts?: pulumi.InvokeOptions): Promise<GetEnterprisePolicyResult> {
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    return pulumi.runtime.invoke("azure-nextgen:powerplatform/v20201030preview:getEnterprisePolicy", {
        "enterprisePolicyName": args.enterprisePolicyName,
        "resourceGroupName": args.resourceGroupName,
    }, opts);
}

export interface GetEnterprisePolicyArgs {
    /**
     * The EnterprisePolicy name.
     */
    readonly enterprisePolicyName: string;
    /**
     * The name of the resource group. The name is case insensitive.
     */
    readonly resourceGroupName: string;
}

/**
 * Definition of the EnterprisePolicy.
 */
export interface GetEnterprisePolicyResult {
    /**
     * The identity of the EnterprisePolicy.
     */
    readonly identity?: outputs.powerplatform.v20201030preview.EnterprisePolicyIdentityResponse;
    /**
     * Region where the EnterprisePolicy is located.
     */
    readonly location: string;
    /**
     * Name of the EnterprisePolicy.
     */
    readonly name: string;
    /**
     * Resource tags.
     */
    readonly tags?: {[key: string]: string};
    /**
     * The type of the resource.
     */
    readonly type: string;
}