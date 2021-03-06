// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../../types/input";
import * as outputs from "../../types/output";
import * as utilities from "../../utilities";

export function listManagerPublicEncryptionKey(args: ListManagerPublicEncryptionKeyArgs, opts?: pulumi.InvokeOptions): Promise<ListManagerPublicEncryptionKeyResult> {
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    return pulumi.runtime.invoke("azure-nextgen:storsimple/latest:listManagerPublicEncryptionKey", {
        "managerName": args.managerName,
        "resourceGroupName": args.resourceGroupName,
    }, opts);
}

export interface ListManagerPublicEncryptionKeyArgs {
    /**
     * The manager name
     */
    readonly managerName: string;
    /**
     * The resource group name
     */
    readonly resourceGroupName: string;
}

/**
 * Represents the secrets encrypted using Symmetric Encryption Key.
 */
export interface ListManagerPublicEncryptionKeyResult {
    /**
     * The algorithm used to encrypt the "Value".
     */
    readonly encryptionAlgorithm: string;
    /**
     * The value of the secret itself. If the secret is in plaintext or null then EncryptionAlgorithm will be none.
     */
    readonly value: string;
    /**
     * The thumbprint of the cert that was used to encrypt "Value".
     */
    readonly valueCertificateThumbprint?: string;
}
