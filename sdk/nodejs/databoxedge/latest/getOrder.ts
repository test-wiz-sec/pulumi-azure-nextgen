// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../../types/input";
import * as outputs from "../../types/output";
import * as utilities from "../../utilities";

export function getOrder(args: GetOrderArgs, opts?: pulumi.InvokeOptions): Promise<GetOrderResult> {
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    return pulumi.runtime.invoke("azure-nextgen:databoxedge/latest:getOrder", {
        "deviceName": args.deviceName,
        "resourceGroupName": args.resourceGroupName,
    }, opts);
}

export interface GetOrderArgs {
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
 * The order details.
 */
export interface GetOrderResult {
    /**
     * The contact details.
     */
    readonly contactInformation: outputs.databoxedge.latest.ContactDetailsResponse;
    /**
     * Current status of the order.
     */
    readonly currentStatus?: outputs.databoxedge.latest.OrderStatusResponse;
    /**
     * Tracking information for the package delivered to the customer whether it has an original or a replacement device.
     */
    readonly deliveryTrackingInfo: outputs.databoxedge.latest.TrackingInfoResponse[];
    /**
     * The object name.
     */
    readonly name: string;
    /**
     * List of status changes in the order.
     */
    readonly orderHistory: outputs.databoxedge.latest.OrderStatusResponse[];
    /**
     * Tracking information for the package returned from the customer whether it has an original or a replacement device.
     */
    readonly returnTrackingInfo: outputs.databoxedge.latest.TrackingInfoResponse[];
    /**
     * Serial number of the device.
     */
    readonly serialNumber: string;
    /**
     * The shipping address.
     */
    readonly shippingAddress: outputs.databoxedge.latest.AddressResponse;
    /**
     * The hierarchical type of the object.
     */
    readonly type: string;
}
