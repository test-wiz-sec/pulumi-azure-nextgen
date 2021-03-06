// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../../types/input";
import * as outputs from "../../types/output";
import * as utilities from "../../utilities";

export function getJob(args: GetJobArgs, opts?: pulumi.InvokeOptions): Promise<GetJobResult> {
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    return pulumi.runtime.invoke("azure-nextgen:scheduler/latest:getJob", {
        "jobCollectionName": args.jobCollectionName,
        "jobName": args.jobName,
        "resourceGroupName": args.resourceGroupName,
    }, opts);
}

export interface GetJobArgs {
    /**
     * The job collection name.
     */
    readonly jobCollectionName: string;
    /**
     * The job name.
     */
    readonly jobName: string;
    /**
     * The resource group name.
     */
    readonly resourceGroupName: string;
}

export interface GetJobResult {
    /**
     * Gets the job resource name.
     */
    readonly name: string;
    /**
     * Gets or sets the job properties.
     */
    readonly properties: outputs.scheduler.latest.JobPropertiesResponse;
    /**
     * Gets the job resource type.
     */
    readonly type: string;
}
