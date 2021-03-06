// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../../types/input";
import * as outputs from "../../types/output";
import * as utilities from "../../utilities";

export function getProtectedItem(args: GetProtectedItemArgs, opts?: pulumi.InvokeOptions): Promise<GetProtectedItemResult> {
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    return pulumi.runtime.invoke("azure-nextgen:recoveryservices/latest:getProtectedItem", {
        "containerName": args.containerName,
        "fabricName": args.fabricName,
        "filter": args.filter,
        "protectedItemName": args.protectedItemName,
        "resourceGroupName": args.resourceGroupName,
        "vaultName": args.vaultName,
    }, opts);
}

export interface GetProtectedItemArgs {
    /**
     * Container name associated with the backed up item.
     */
    readonly containerName: string;
    /**
     * Fabric name associated with the backed up item.
     */
    readonly fabricName: string;
    /**
     * OData filter options.
     */
    readonly filter?: string;
    /**
     * Backed up item name whose details are to be fetched.
     */
    readonly protectedItemName: string;
    /**
     * The name of the resource group where the recovery services vault is present.
     */
    readonly resourceGroupName: string;
    /**
     * The name of the recovery services vault.
     */
    readonly vaultName: string;
}

/**
 * Base class for backup items.
 */
export interface GetProtectedItemResult {
    /**
     * Optional ETag.
     */
    readonly eTag?: string;
    /**
     * Resource location.
     */
    readonly location?: string;
    /**
     * Resource name associated with the resource.
     */
    readonly name: string;
    /**
     * ProtectedItemResource properties
     */
    readonly properties: outputs.recoveryservices.latest.AzureFileshareProtectedItemResponse | outputs.recoveryservices.latest.AzureIaaSVMProtectedItemResponse | outputs.recoveryservices.latest.AzureSqlProtectedItemResponse | outputs.recoveryservices.latest.AzureVmWorkloadProtectedItemResponse | outputs.recoveryservices.latest.DPMProtectedItemResponse | outputs.recoveryservices.latest.GenericProtectedItemResponse | outputs.recoveryservices.latest.MabFileFolderProtectedItemResponse;
    /**
     * Resource tags.
     */
    readonly tags?: {[key: string]: string};
    /**
     * Resource type represents the complete path of the form Namespace/ResourceType/ResourceType/...
     */
    readonly type: string;
}
