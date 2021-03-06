// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../../types/input";
import * as outputs from "../../types/output";
import * as utilities from "../../utilities";

export function getPrivateDnsZoneGroup(args: GetPrivateDnsZoneGroupArgs, opts?: pulumi.InvokeOptions): Promise<GetPrivateDnsZoneGroupResult> {
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    return pulumi.runtime.invoke("azure-nextgen:network/latest:getPrivateDnsZoneGroup", {
        "privateDnsZoneGroupName": args.privateDnsZoneGroupName,
        "privateEndpointName": args.privateEndpointName,
        "resourceGroupName": args.resourceGroupName,
    }, opts);
}

export interface GetPrivateDnsZoneGroupArgs {
    /**
     * The name of the private dns zone group.
     */
    readonly privateDnsZoneGroupName: string;
    /**
     * The name of the private endpoint.
     */
    readonly privateEndpointName: string;
    /**
     * The name of the resource group.
     */
    readonly resourceGroupName: string;
}

/**
 * Private dns zone group resource.
 */
export interface GetPrivateDnsZoneGroupResult {
    /**
     * A unique read-only string that changes whenever the resource is updated.
     */
    readonly etag: string;
    /**
     * Name of the resource that is unique within a resource group. This name can be used to access the resource.
     */
    readonly name?: string;
    /**
     * A collection of private dns zone configurations of the private dns zone group.
     */
    readonly privateDnsZoneConfigs?: outputs.network.latest.PrivateDnsZoneConfigResponse[];
    /**
     * The provisioning state of the private dns zone group resource.
     */
    readonly provisioningState: string;
}
