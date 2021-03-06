// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../../types/input";
import * as outputs from "../../types/output";
import * as utilities from "../../utilities";

export function getSiteVNETConnectionSlot(args: GetSiteVNETConnectionSlotArgs, opts?: pulumi.InvokeOptions): Promise<GetSiteVNETConnectionSlotResult> {
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    return pulumi.runtime.invoke("azure-nextgen:web/v20150801:getSiteVNETConnectionSlot", {
        "name": args.name,
        "resourceGroupName": args.resourceGroupName,
        "slot": args.slot,
        "vnetName": args.vnetName,
    }, opts);
}

export interface GetSiteVNETConnectionSlotArgs {
    /**
     * The name of the web app
     */
    readonly name: string;
    /**
     * The resource group name
     */
    readonly resourceGroupName: string;
    /**
     * The name of the slot for this web app.
     */
    readonly slot: string;
    /**
     * The name of the Virtual Network
     */
    readonly vnetName: string;
}

/**
 * VNETInfo contract. This contract is public and is a stripped down version of VNETInfoInternal
 */
export interface GetSiteVNETConnectionSlotResult {
    /**
     * A certificate file (.cer) blob containing the public key of the private key used to authenticate a 
     *             Point-To-Site VPN connection.
     */
    readonly certBlob?: string;
    /**
     * The client certificate thumbprint
     */
    readonly certThumbprint?: string;
    /**
     * Dns servers to be used by this VNET. This should be a comma-separated list of IP addresses.
     */
    readonly dnsServers?: string;
    /**
     * Kind of resource
     */
    readonly kind?: string;
    /**
     * Resource Location
     */
    readonly location: string;
    /**
     * Resource Name
     */
    readonly name?: string;
    /**
     * Flag to determine if a resync is required
     */
    readonly resyncRequired?: boolean;
    /**
     * The routes that this virtual network connection uses.
     */
    readonly routes?: outputs.web.v20150801.VnetRouteResponse[];
    /**
     * Resource tags
     */
    readonly tags?: {[key: string]: string};
    /**
     * Resource type
     */
    readonly type?: string;
    /**
     * The vnet resource id
     */
    readonly vnetResourceId?: string;
}
