// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../../types/input";
import * as outputs from "../../types/output";
import * as utilities from "../../utilities";

export function getHybridUseBenefit(args: GetHybridUseBenefitArgs, opts?: pulumi.InvokeOptions): Promise<GetHybridUseBenefitResult> {
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    return pulumi.runtime.invoke("azure-nextgen:softwareplan/v20191201:getHybridUseBenefit", {
        "planId": args.planId,
        "scope": args.scope,
    }, opts);
}

export interface GetHybridUseBenefitArgs {
    /**
     * This is a unique identifier for a plan. Should be a guid.
     */
    readonly planId: string;
    /**
     * The scope at which the operation is performed. This is limited to Microsoft.Compute/virtualMachines and Microsoft.Compute/hostGroups/hosts for now
     */
    readonly scope: string;
}

/**
 * Response on GET of a hybrid use benefit
 */
export interface GetHybridUseBenefitResult {
    /**
     * Created date
     */
    readonly createdDate: string;
    /**
     * Indicates the revision of the hybrid use benefit
     */
    readonly etag: number;
    /**
     * Last updated date
     */
    readonly lastUpdatedDate: string;
    /**
     * The name of the resource
     */
    readonly name: string;
    /**
     * Provisioning state
     */
    readonly provisioningState: string;
    /**
     * Hybrid use benefit SKU
     */
    readonly sku: outputs.softwareplan.v20191201.SkuResponse;
    /**
     * The type of the resource. Ex- Microsoft.Compute/virtualMachines or Microsoft.Storage/storageAccounts.
     */
    readonly type: string;
}