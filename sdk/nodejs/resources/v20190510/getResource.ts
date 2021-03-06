// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../../types/input";
import * as outputs from "../../types/output";
import * as utilities from "../../utilities";

export function getResource(args: GetResourceArgs, opts?: pulumi.InvokeOptions): Promise<GetResourceResult> {
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    return pulumi.runtime.invoke("azure-nextgen:resources/v20190510:getResource", {
        "parentResourcePath": args.parentResourcePath,
        "resourceGroupName": args.resourceGroupName,
        "resourceName": args.resourceName,
        "resourceProviderNamespace": args.resourceProviderNamespace,
        "resourceType": args.resourceType,
    }, opts);
}

export interface GetResourceArgs {
    /**
     * The parent resource identity.
     */
    readonly parentResourcePath: string;
    /**
     * The name of the resource group containing the resource to get. The name is case insensitive.
     */
    readonly resourceGroupName: string;
    /**
     * The name of the resource to get.
     */
    readonly resourceName: string;
    /**
     * The namespace of the resource provider.
     */
    readonly resourceProviderNamespace: string;
    /**
     * The resource type of the resource.
     */
    readonly resourceType: string;
}

/**
 * Resource information.
 */
export interface GetResourceResult {
    /**
     * The identity of the resource.
     */
    readonly identity?: outputs.resources.v20190510.IdentityResponse;
    /**
     * The kind of the resource.
     */
    readonly kind?: string;
    /**
     * Resource location
     */
    readonly location?: string;
    /**
     * ID of the resource that manages this resource.
     */
    readonly managedBy?: string;
    /**
     * Resource name
     */
    readonly name: string;
    /**
     * The plan of the resource.
     */
    readonly plan?: outputs.resources.v20190510.PlanResponse;
    /**
     * The resource properties.
     */
    readonly properties: any;
    /**
     * The SKU of the resource.
     */
    readonly sku?: outputs.resources.v20190510.SkuResponse;
    /**
     * Resource tags
     */
    readonly tags?: {[key: string]: string};
    /**
     * Resource type
     */
    readonly type: string;
}
