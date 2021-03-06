// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../../types/input";
import * as outputs from "../../types/output";
import * as utilities from "../../utilities";

export function getNetworkProfile(args: GetNetworkProfileArgs, opts?: pulumi.InvokeOptions): Promise<GetNetworkProfileResult> {
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    return pulumi.runtime.invoke("azure-nextgen:network/v20181101:getNetworkProfile", {
        "expand": args.expand,
        "networkProfileName": args.networkProfileName,
        "resourceGroupName": args.resourceGroupName,
    }, opts);
}

export interface GetNetworkProfileArgs {
    /**
     * Expands referenced resources.
     */
    readonly expand?: string;
    /**
     * The name of the PublicIPPrefix.
     */
    readonly networkProfileName: string;
    /**
     * The name of the resource group.
     */
    readonly resourceGroupName: string;
}

/**
 * Network profile resource.
 */
export interface GetNetworkProfileResult {
    /**
     * List of chid container network interface configurations.
     */
    readonly containerNetworkInterfaceConfigurations?: outputs.network.v20181101.ContainerNetworkInterfaceConfigurationResponse[];
    /**
     * List of child container network interfaces.
     */
    readonly containerNetworkInterfaces?: outputs.network.v20181101.ContainerNetworkInterfaceResponse[];
    /**
     * A unique read-only string that changes whenever the resource is updated.
     */
    readonly etag?: string;
    /**
     * Resource location.
     */
    readonly location?: string;
    /**
     * Resource name.
     */
    readonly name: string;
    /**
     * The provisioning state of the resource.
     */
    readonly provisioningState: string;
    /**
     * The resource GUID property of the network interface resource.
     */
    readonly resourceGuid: string;
    /**
     * Resource tags.
     */
    readonly tags?: {[key: string]: string};
    /**
     * Resource type.
     */
    readonly type: string;
}
