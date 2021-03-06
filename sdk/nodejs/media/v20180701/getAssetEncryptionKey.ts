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
    return pulumi.runtime.invoke("azure-nextgen:media/v20180701:getAssetEncryptionKey", {
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
 * Data needed to decrypt asset files encrypted with legacy storage encryption.
 */
export interface GetAssetEncryptionKeyResult {
    /**
     * Asset File encryption metadata.
     */
    readonly assetFileEncryptionMetadata?: outputs.media.v20180701.AssetFileEncryptionMetadataResponse[];
    /**
     * The Asset File storage encryption key.
     */
    readonly key?: string;
}
