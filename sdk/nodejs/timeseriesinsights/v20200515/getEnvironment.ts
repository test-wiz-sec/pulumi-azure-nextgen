// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../../types/input";
import * as outputs from "../../types/output";
import * as utilities from "../../utilities";

export function getEnvironment(args: GetEnvironmentArgs, opts?: pulumi.InvokeOptions): Promise<GetEnvironmentResult> {
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    return pulumi.runtime.invoke("azure-nextgen:timeseriesinsights/v20200515:getEnvironment", {
        "environmentName": args.environmentName,
        "expand": args.expand,
        "resourceGroupName": args.resourceGroupName,
    }, opts);
}

export interface GetEnvironmentArgs {
    /**
     * The name of the Time Series Insights environment associated with the specified resource group.
     */
    readonly environmentName: string;
    /**
     * Setting $expand=status will include the status of the internal services of the environment in the Time Series Insights service.
     */
    readonly expand?: string;
    /**
     * Name of an Azure Resource group.
     */
    readonly resourceGroupName: string;
}

/**
 * An environment is a set of time-series data available for query, and is the top level Azure Time Series Insights resource.
 */
export interface GetEnvironmentResult {
    /**
     * The kind of the environment.
     */
    readonly kind: string;
    /**
     * Resource location
     */
    readonly location: string;
    /**
     * Resource name
     */
    readonly name: string;
    /**
     * The sku determines the type of environment, either Gen1 (S1 or S2) or Gen2 (L1). For Gen1 environments the sku determines the capacity of the environment, the ingress rate, and the billing rate.
     */
    readonly sku: outputs.timeseriesinsights.v20200515.SkuResponse;
    /**
     * Resource tags
     */
    readonly tags?: {[key: string]: string};
    /**
     * Resource type
     */
    readonly type: string;
}
