// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../../types/input";
import * as outputs from "../../types/output";
import * as utilities from "../../utilities";

export function getReportByResourceGroupName(args: GetReportByResourceGroupNameArgs, opts?: pulumi.InvokeOptions): Promise<GetReportByResourceGroupNameResult> {
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    return pulumi.runtime.invoke("azure-nextgen:costmanagement/v20180801preview:getReportByResourceGroupName", {
        "reportName": args.reportName,
        "resourceGroupName": args.resourceGroupName,
    }, opts);
}

export interface GetReportByResourceGroupNameArgs {
    /**
     * Report Name.
     */
    readonly reportName: string;
    /**
     * Azure Resource Group Name.
     */
    readonly resourceGroupName: string;
}

/**
 * A report resource.
 */
export interface GetReportByResourceGroupNameResult {
    /**
     * Has definition for the report.
     */
    readonly definition: outputs.costmanagement.v20180801preview.ReportDefinitionResponse;
    /**
     * Has delivery information for the report.
     */
    readonly deliveryInfo: outputs.costmanagement.v20180801preview.ReportDeliveryInfoResponse;
    /**
     * The format of the report being delivered.
     */
    readonly format?: string;
    /**
     * Resource name.
     */
    readonly name: string;
    /**
     * Has schedule information for the report.
     */
    readonly schedule?: outputs.costmanagement.v20180801preview.ReportScheduleResponse;
    /**
     * Resource tags.
     */
    readonly tags: {[key: string]: string};
    /**
     * Resource type.
     */
    readonly type: string;
}
