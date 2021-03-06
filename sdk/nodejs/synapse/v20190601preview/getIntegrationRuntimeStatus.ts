// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../../types/input";
import * as outputs from "../../types/output";
import * as utilities from "../../utilities";

export function getIntegrationRuntimeStatus(args: GetIntegrationRuntimeStatusArgs, opts?: pulumi.InvokeOptions): Promise<GetIntegrationRuntimeStatusResult> {
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    return pulumi.runtime.invoke("azure-nextgen:synapse/v20190601preview:getIntegrationRuntimeStatus", {
        "integrationRuntimeName": args.integrationRuntimeName,
        "resourceGroupName": args.resourceGroupName,
        "workspaceName": args.workspaceName,
    }, opts);
}

export interface GetIntegrationRuntimeStatusArgs {
    /**
     * Integration runtime name
     */
    readonly integrationRuntimeName: string;
    /**
     * The name of the resource group. The name is case insensitive.
     */
    readonly resourceGroupName: string;
    /**
     * The name of the workspace
     */
    readonly workspaceName: string;
}

/**
 * Integration runtime status response.
 */
export interface GetIntegrationRuntimeStatusResult {
    /**
     * The integration runtime name.
     */
    readonly name: string;
    /**
     * Integration runtime properties.
     */
    readonly properties: outputs.synapse.v20190601preview.ManagedIntegrationRuntimeStatusResponse | outputs.synapse.v20190601preview.SelfHostedIntegrationRuntimeStatusResponse;
}
