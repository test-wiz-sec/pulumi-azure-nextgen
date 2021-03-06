// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../../types/input";
import * as outputs from "../../types/output";
import * as utilities from "../../utilities";

export function getServerKey(args: GetServerKeyArgs, opts?: pulumi.InvokeOptions): Promise<GetServerKeyResult> {
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    return pulumi.runtime.invoke("azure-nextgen:dbformysql/latest:getServerKey", {
        "keyName": args.keyName,
        "resourceGroupName": args.resourceGroupName,
        "serverName": args.serverName,
    }, opts);
}

export interface GetServerKeyArgs {
    /**
     * The name of the MySQL Server key to be retrieved.
     */
    readonly keyName: string;
    /**
     * The name of the resource group. The name is case insensitive.
     */
    readonly resourceGroupName: string;
    /**
     * The name of the server.
     */
    readonly serverName: string;
}

/**
 * A MySQL Server key.
 */
export interface GetServerKeyResult {
    /**
     * The key creation date.
     */
    readonly creationDate: string;
    /**
     * Kind of encryption protector used to protect the key.
     */
    readonly kind: string;
    /**
     * The name of the resource
     */
    readonly name: string;
    /**
     * The key type like 'AzureKeyVault'.
     */
    readonly serverKeyType: string;
    /**
     * The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
     */
    readonly type: string;
    /**
     * The URI of the key.
     */
    readonly uri?: string;
}
