// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../../types/input";
import * as outputs from "../../types/output";
import * as utilities from "../../utilities";

export function getProtectionContainer(args: GetProtectionContainerArgs, opts?: pulumi.InvokeOptions): Promise<GetProtectionContainerResult> {
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    return pulumi.runtime.invoke("azure-nextgen:recoveryservices/latest:getProtectionContainer", {
        "containerName": args.containerName,
        "fabricName": args.fabricName,
        "resourceGroupName": args.resourceGroupName,
        "vaultName": args.vaultName,
    }, opts);
}

export interface GetProtectionContainerArgs {
    /**
     * Name of the container whose details need to be fetched.
     */
    readonly containerName: string;
    /**
     * Name of the fabric where the container belongs.
     */
    readonly fabricName: string;
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
 * Base class for container with backup items. Containers with specific workloads are derived from this class.
 */
export interface GetProtectionContainerResult {
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
     * ProtectionContainerResource properties
     */
    readonly properties: outputs.recoveryservices.latest.AzureSqlContainerResponse | outputs.recoveryservices.latest.AzureStorageContainerResponse | outputs.recoveryservices.latest.AzureWorkloadContainerResponse | outputs.recoveryservices.latest.DpmContainerResponse | outputs.recoveryservices.latest.GenericContainerResponse | outputs.recoveryservices.latest.IaaSVMContainerResponse | outputs.recoveryservices.latest.MabContainerResponse;
    /**
     * Resource tags.
     */
    readonly tags?: {[key: string]: string};
    /**
     * Resource type represents the complete path of the form Namespace/ResourceType/ResourceType/...
     */
    readonly type: string;
}
