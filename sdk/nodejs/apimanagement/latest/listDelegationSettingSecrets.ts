// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../../types/input";
import * as outputs from "../../types/output";
import * as utilities from "../../utilities";

export function listDelegationSettingSecrets(args: ListDelegationSettingSecretsArgs, opts?: pulumi.InvokeOptions): Promise<ListDelegationSettingSecretsResult> {
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    return pulumi.runtime.invoke("azure-nextgen:apimanagement/latest:listDelegationSettingSecrets", {
        "resourceGroupName": args.resourceGroupName,
        "serviceName": args.serviceName,
    }, opts);
}

export interface ListDelegationSettingSecretsArgs {
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
 * Client or app secret used in IdentityProviders, Aad, OpenID or OAuth.
 */
export interface ListDelegationSettingSecretsResult {
    /**
     * This is secret value of the validation key in portal settings.
     */
    readonly validationKey?: string;
}
