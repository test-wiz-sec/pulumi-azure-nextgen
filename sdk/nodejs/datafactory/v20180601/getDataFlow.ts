// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../../types/input";
import * as outputs from "../../types/output";
import * as utilities from "../../utilities";

export function getDataFlow(args: GetDataFlowArgs, opts?: pulumi.InvokeOptions): Promise<GetDataFlowResult> {
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    return pulumi.runtime.invoke("azure-nextgen:datafactory/v20180601:getDataFlow", {
        "dataFlowName": args.dataFlowName,
        "factoryName": args.factoryName,
        "resourceGroupName": args.resourceGroupName,
    }, opts);
}

export interface GetDataFlowArgs {
    /**
     * The data flow name.
     */
    readonly dataFlowName: string;
    /**
     * The factory name.
     */
    readonly factoryName: string;
    /**
     * The resource group name.
     */
    readonly resourceGroupName: string;
}

/**
 * Data flow resource type.
 */
export interface GetDataFlowResult {
    /**
     * Etag identifies change in the resource.
     */
    readonly etag: string;
    /**
     * The resource name.
     */
    readonly name: string;
    /**
     * Data flow properties.
     */
    readonly properties: outputs.datafactory.v20180601.MappingDataFlowResponse;
    /**
     * The resource type.
     */
    readonly type: string;
}
