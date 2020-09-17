// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../../types/input";
import * as outputs from "../../types/output";
import * as utilities from "../../utilities";

export function getIoTSpace(args: GetIoTSpaceArgs, opts?: pulumi.InvokeOptions): Promise<GetIoTSpaceResult> {
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    return pulumi.runtime.invoke("azure-nextgen:iotspaces/v20171001preview:getIoTSpace", {
        "resourceGroupName": args.resourceGroupName,
        "resourceName": args.resourceName,
    }, opts);
}

export interface GetIoTSpaceArgs {
    /**
     * The name of the resource group that contains the IoTSpaces instance.
     */
    readonly resourceGroupName: string;
    /**
     * The name of the IoTSpaces instance.
     */
    readonly resourceName: string;
}

/**
 * The description of the IoTSpaces service.
 */
export interface GetIoTSpaceResult {
    /**
     * The resource location.
     */
    readonly location: string;
    /**
     * The resource name.
     */
    readonly name: string;
    /**
     * The common properties of a IoTSpaces service.
     */
    readonly properties: outputs.iotspaces.v20171001preview.IoTSpacesPropertiesResponse;
    /**
     * A valid instance SKU.
     */
    readonly sku: outputs.iotspaces.v20171001preview.IoTSpacesSkuInfoResponse;
    /**
     * The resource tags.
     */
    readonly tags?: {[key: string]: string};
    /**
     * The resource type.
     */
    readonly type: string;
}