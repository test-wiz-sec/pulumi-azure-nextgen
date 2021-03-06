// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../../types/input";
import * as outputs from "../../types/output";
import * as utilities from "../../utilities";

export function getService(args: GetServiceArgs, opts?: pulumi.InvokeOptions): Promise<GetServiceResult> {
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    return pulumi.runtime.invoke("azure-nextgen:windowsiot/v20190601:getService", {
        "deviceName": args.deviceName,
        "resourceGroupName": args.resourceGroupName,
    }, opts);
}

export interface GetServiceArgs {
    /**
     * The name of the Windows IoT Device Service.
     */
    readonly deviceName: string;
    /**
     * The name of the resource group that contains the Windows IoT Device Service.
     */
    readonly resourceGroupName: string;
}

/**
 * The description of the Windows IoT Device Service.
 */
export interface GetServiceResult {
    /**
     * Windows IoT Device Service OEM AAD domain
     */
    readonly adminDomainName?: string;
    /**
     * Windows IoT Device Service ODM AAD domain
     */
    readonly billingDomainName?: string;
    /**
     * The Etag field is *not* required. If it is provided in the response body, it must also be provided as a header per the normal ETag convention.
     */
    readonly etag?: string;
    /**
     * The Azure Region where the resource lives
     */
    readonly location?: string;
    /**
     * The name of the resource
     */
    readonly name: string;
    /**
     * Windows IoT Device Service notes.
     */
    readonly notes?: string;
    /**
     * Windows IoT Device Service device allocation,
     */
    readonly quantity?: number;
    /**
     * Windows IoT Device Service start date,
     */
    readonly startDate: string;
    /**
     * Resource tags.
     */
    readonly tags?: {[key: string]: string};
    /**
     * The type of the resource.
     */
    readonly type: string;
}
