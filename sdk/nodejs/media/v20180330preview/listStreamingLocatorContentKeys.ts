// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../../types/input";
import * as outputs from "../../types/output";
import * as utilities from "../../utilities";

export function listStreamingLocatorContentKeys(args: ListStreamingLocatorContentKeysArgs, opts?: pulumi.InvokeOptions): Promise<ListStreamingLocatorContentKeysResult> {
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    return pulumi.runtime.invoke("azure-nextgen:media/v20180330preview:listStreamingLocatorContentKeys", {
        "accountName": args.accountName,
        "resourceGroupName": args.resourceGroupName,
        "streamingLocatorName": args.streamingLocatorName,
    }, opts);
}

export interface ListStreamingLocatorContentKeysArgs {
    /**
     * The Media Services account name.
     */
    readonly accountName: string;
    /**
     * The name of the resource group within the Azure subscription.
     */
    readonly resourceGroupName: string;
    /**
     * The Streaming Locator name.
     */
    readonly streamingLocatorName: string;
}

/**
 * Class of response for listContentKeys action
 */
export interface ListStreamingLocatorContentKeysResult {
    /**
     * ContentKeys used by current Streaming Locator
     */
    readonly contentKeys?: outputs.media.v20180330preview.StreamingLocatorContentKeyResponse[];
}
