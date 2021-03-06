// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../../types/input";
import * as outputs from "../../types/output";
import * as utilities from "../../utilities";

export function getPeering(args: GetPeeringArgs, opts?: pulumi.InvokeOptions): Promise<GetPeeringResult> {
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    return pulumi.runtime.invoke("azure-nextgen:peering/v20190801preview:getPeering", {
        "peeringName": args.peeringName,
        "resourceGroupName": args.resourceGroupName,
    }, opts);
}

export interface GetPeeringArgs {
    /**
     * The name of the peering.
     */
    readonly peeringName: string;
    /**
     * The name of the resource group.
     */
    readonly resourceGroupName: string;
}

/**
 * Peering is a logical representation of a set of connections to the Microsoft Cloud Edge at a location.
 */
export interface GetPeeringResult {
    /**
     * The properties that define a direct peering.
     */
    readonly direct?: outputs.peering.v20190801preview.PeeringPropertiesDirectResponse;
    /**
     * The properties that define an exchange peering.
     */
    readonly exchange?: outputs.peering.v20190801preview.PeeringPropertiesExchangeResponse;
    /**
     * The kind of the peering.
     */
    readonly kind: string;
    /**
     * The location of the resource.
     */
    readonly location: string;
    /**
     * The name of the resource.
     */
    readonly name: string;
    /**
     * The location of the peering.
     */
    readonly peeringLocation?: string;
    /**
     * The provisioning state of the resource.
     */
    readonly provisioningState: string;
    /**
     * The SKU that defines the tier and kind of the peering.
     */
    readonly sku: outputs.peering.v20190801preview.PeeringSkuResponse;
    /**
     * The resource tags.
     */
    readonly tags?: {[key: string]: string};
    /**
     * The type of the resource.
     */
    readonly type: string;
}
