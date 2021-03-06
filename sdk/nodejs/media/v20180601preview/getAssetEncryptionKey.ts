// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../../types/input";
import * as outputs from "../../types/output";
import * as utilities from "../../utilities";

export function getAssetEncryptionKey(args: GetAssetEncryptionKeyArgs, opts?: pulumi.InvokeOptions): Promise<GetAssetEncryptionKeyResult> {
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    return pulumi.runtime.invoke("azure-nextgen:media/v20180601preview:getAssetEncryptionKey", {
        "accountName": args.accountName,
        "assetName": args.assetName,
        "resourceGroupName": args.resourceGroupName,
    }, opts);
}

export interface GetAssetEncryptionKeyArgs {
    /**
     * The Media Services account name.
     */
    readonly accountName: string;
    /**
     * The Asset name.
     */
    readonly assetName: string;
    /**
     * The name of the resource group within the Azure subscription.
     */
    readonly resourceGroupName: string;
}

/**
 * The Asset Storage encryption key.
 */
export interface GetAssetEncryptionKeyResult {
    /**
     * The Asset storage encryption key.
     */
    readonly storageEncryptionKey?: string;
}
