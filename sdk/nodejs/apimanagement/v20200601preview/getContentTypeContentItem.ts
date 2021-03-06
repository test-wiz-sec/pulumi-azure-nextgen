// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../../types/input";
import * as outputs from "../../types/output";
import * as utilities from "../../utilities";

export function getContentTypeContentItem(args: GetContentTypeContentItemArgs, opts?: pulumi.InvokeOptions): Promise<GetContentTypeContentItemResult> {
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    return pulumi.runtime.invoke("azure-nextgen:apimanagement/v20200601preview:getContentTypeContentItem", {
        "contentItemId": args.contentItemId,
        "contentTypeId": args.contentTypeId,
        "resourceGroupName": args.resourceGroupName,
        "serviceName": args.serviceName,
    }, opts);
}

export interface GetContentTypeContentItemArgs {
    /**
     * Content item identifier.
     */
    readonly contentItemId: string;
    /**
     * Content type identifier.
     */
    readonly contentTypeId: string;
    /**
     * The name of the resource group.
     */
    readonly resourceGroupName: string;
    /**
     * The name of the API Management service.
     */
    readonly serviceName: string;
}

/**
 * Content type contract details.
 */
export interface GetContentTypeContentItemResult {
    /**
     * Resource name.
     */
    readonly name: string;
    /**
     * Properties of the content item.
     */
    readonly properties: any;
    /**
     * Resource type for API Management resource.
     */
    readonly type: string;
}
