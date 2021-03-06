// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../../types/input";
import * as outputs from "../../types/output";
import * as utilities from "../../utilities";

export function getSnapshotPolicy(args: GetSnapshotPolicyArgs, opts?: pulumi.InvokeOptions): Promise<GetSnapshotPolicyResult> {
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    return pulumi.runtime.invoke("azure-nextgen:netapp/v20200601:getSnapshotPolicy", {
        "accountName": args.accountName,
        "resourceGroupName": args.resourceGroupName,
        "snapshotPolicyName": args.snapshotPolicyName,
    }, opts);
}

export interface GetSnapshotPolicyArgs {
    /**
     * The name of the NetApp account
     */
    readonly accountName: string;
    /**
     * The name of the resource group.
     */
    readonly resourceGroupName: string;
    /**
     * The name of the snapshot policy target
     */
    readonly snapshotPolicyName: string;
}

/**
 * Snapshot policy information
 */
export interface GetSnapshotPolicyResult {
    /**
     * Schedule for daily snapshots
     */
    readonly dailySchedule?: outputs.netapp.v20200601.DailyScheduleResponse;
    /**
     * The property to decide policy is enabled or not
     */
    readonly enabled?: boolean;
    /**
     * Schedule for hourly snapshots
     */
    readonly hourlySchedule?: outputs.netapp.v20200601.HourlyScheduleResponse;
    /**
     * Resource location
     */
    readonly location: string;
    /**
     * Schedule for monthly snapshots
     */
    readonly monthlySchedule?: outputs.netapp.v20200601.MonthlyScheduleResponse;
    /**
     * Snapshot policy name
     */
    readonly name: string;
    /**
     * Azure lifecycle management
     */
    readonly provisioningState: string;
    /**
     * Resource tags
     */
    readonly tags?: {[key: string]: string};
    /**
     * Resource type
     */
    readonly type: string;
    /**
     * Schedule for weekly snapshots
     */
    readonly weeklySchedule?: outputs.netapp.v20200601.WeeklyScheduleResponse;
}
