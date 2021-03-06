// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../../types/input";
import * as outputs from "../../types/output";
import * as utilities from "../../utilities";

export function listBlockchainMemberApiKeys(args: ListBlockchainMemberApiKeysArgs, opts?: pulumi.InvokeOptions): Promise<ListBlockchainMemberApiKeysResult> {
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    return pulumi.runtime.invoke("azure-nextgen:blockchain/v20180601preview:listBlockchainMemberApiKeys", {
        "blockchainMemberName": args.blockchainMemberName,
        "resourceGroupName": args.resourceGroupName,
    }, opts);
}

export interface ListBlockchainMemberApiKeysArgs {
    /**
     * Blockchain member name.
     */
    readonly blockchainMemberName: string;
    /**
     * The name of the resource group that contains the resource. You can obtain this value from the Azure Resource Manager API or the portal.
     */
    readonly resourceGroupName: string;
}

/**
 * Collection of the API key payload which is exposed in the response of the resource provider.
 */
export interface ListBlockchainMemberApiKeysResult {
    /**
     * Gets or sets the collection of API key.
     */
    readonly keys?: outputs.blockchain.v20180601preview.ApiKeyResponse[];
}
