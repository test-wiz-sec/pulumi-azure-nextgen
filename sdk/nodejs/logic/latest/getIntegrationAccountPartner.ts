// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../../types/input";
import * as outputs from "../../types/output";
import * as utilities from "../../utilities";

export function getIntegrationAccountPartner(args: GetIntegrationAccountPartnerArgs, opts?: pulumi.InvokeOptions): Promise<GetIntegrationAccountPartnerResult> {
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    return pulumi.runtime.invoke("azure-nextgen:logic/latest:getIntegrationAccountPartner", {
        "integrationAccountName": args.integrationAccountName,
        "partnerName": args.partnerName,
        "resourceGroupName": args.resourceGroupName,
    }, opts);
}

export interface GetIntegrationAccountPartnerArgs {
    /**
     * The integration account name.
     */
    readonly integrationAccountName: string;
    /**
     * The integration account partner name.
     */
    readonly partnerName: string;
    /**
     * The resource group name.
     */
    readonly resourceGroupName: string;
}

/**
 * The integration account partner.
 */
export interface GetIntegrationAccountPartnerResult {
    /**
     * The changed time.
     */
    readonly changedTime: string;
    /**
     * The partner content.
     */
    readonly content: outputs.logic.latest.PartnerContentResponse;
    /**
     * The created time.
     */
    readonly createdTime: string;
    /**
     * The resource location.
     */
    readonly location?: string;
    /**
     * The metadata.
     */
    readonly metadata?: any;
    /**
     * Gets the resource name.
     */
    readonly name: string;
    /**
     * The partner type.
     */
    readonly partnerType: string;
    /**
     * The resource tags.
     */
    readonly tags?: {[key: string]: string};
    /**
     * Gets the resource type.
     */
    readonly type: string;
}
