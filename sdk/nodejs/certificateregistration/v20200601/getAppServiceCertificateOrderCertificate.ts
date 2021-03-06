// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../../types/input";
import * as outputs from "../../types/output";
import * as utilities from "../../utilities";

export function getAppServiceCertificateOrderCertificate(args: GetAppServiceCertificateOrderCertificateArgs, opts?: pulumi.InvokeOptions): Promise<GetAppServiceCertificateOrderCertificateResult> {
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    return pulumi.runtime.invoke("azure-nextgen:certificateregistration/v20200601:getAppServiceCertificateOrderCertificate", {
        "certificateOrderName": args.certificateOrderName,
        "name": args.name,
        "resourceGroupName": args.resourceGroupName,
    }, opts);
}

export interface GetAppServiceCertificateOrderCertificateArgs {
    /**
     * Name of the certificate order.
     */
    readonly certificateOrderName: string;
    /**
     * Name of the certificate.
     */
    readonly name: string;
    /**
     * Name of the resource group to which the resource belongs.
     */
    readonly resourceGroupName: string;
}

/**
 * Key Vault container ARM resource for a certificate that is purchased through Azure.
 */
export interface GetAppServiceCertificateOrderCertificateResult {
    /**
     * Key Vault resource Id.
     */
    readonly keyVaultId?: string;
    /**
     * Key Vault secret name.
     */
    readonly keyVaultSecretName?: string;
    /**
     * Kind of resource.
     */
    readonly kind?: string;
    /**
     * Resource Location.
     */
    readonly location: string;
    /**
     * Resource Name.
     */
    readonly name: string;
    /**
     * Status of the Key Vault secret.
     */
    readonly provisioningState: string;
    /**
     * Resource tags.
     */
    readonly tags?: {[key: string]: string};
    /**
     * Resource type.
     */
    readonly type: string;
}
