// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../../types/input";
import * as outputs from "../../types/output";
import * as utilities from "../../utilities";

export function getDeviceExtendedInformation(args: GetDeviceExtendedInformationArgs, opts?: pulumi.InvokeOptions): Promise<GetDeviceExtendedInformationResult> {
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    return pulumi.runtime.invoke("azure-nextgen:databoxedge/latest:getDeviceExtendedInformation", {
        "deviceName": args.deviceName,
        "resourceGroupName": args.resourceGroupName,
    }, opts);
}

export interface GetDeviceExtendedInformationArgs {
    /**
     * The device name.
     */
    readonly deviceName: string;
    /**
     * The resource group name.
     */
    readonly resourceGroupName: string;
}

/**
 * The extended Info of the Data Box Edge/Gateway device.
 */
export interface GetDeviceExtendedInformationResult {
    /**
     * The public part of the encryption certificate. Client uses this to encrypt any secret.
     */
    readonly encryptionKey?: string;
    /**
     * The digital signature of encrypted certificate.
     */
    readonly encryptionKeyThumbprint?: string;
    /**
     * The object name.
     */
    readonly name: string;
    /**
     * The Resource ID of the Resource.
     */
    readonly resourceKey: string;
    /**
     * The hierarchical type of the object.
     */
    readonly type: string;
}
