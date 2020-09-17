// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../../types/input";
import * as outputs from "../../types/output";
import * as utilities from "../../utilities";

export function getConfigurationProfilePreference(args: GetConfigurationProfilePreferenceArgs, opts?: pulumi.InvokeOptions): Promise<GetConfigurationProfilePreferenceResult> {
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    return pulumi.runtime.invoke("azure-nextgen:automanage/v20200630preview:getConfigurationProfilePreference", {
        "configurationProfilePreferenceName": args.configurationProfilePreferenceName,
        "resourceGroupName": args.resourceGroupName,
    }, opts);
}

export interface GetConfigurationProfilePreferenceArgs {
    /**
     * The configuration profile preference name.
     */
    readonly configurationProfilePreferenceName: string;
    /**
     * The name of the resource group. The name is case insensitive.
     */
    readonly resourceGroupName: string;
}

/**
 * Definition of the configuration profile preference.
 */
export interface GetConfigurationProfilePreferenceResult {
    /**
     * The geo-location where the resource lives
     */
    readonly location: string;
    /**
     * The name of the resource
     */
    readonly name: string;
    /**
     * Properties of the configuration profile preference.
     */
    readonly properties: outputs.automanage.v20200630preview.ConfigurationProfilePreferencePropertiesResponse;
    /**
     * Resource tags.
     */
    readonly tags?: {[key: string]: string};
    /**
     * The type of the resource. Ex- Microsoft.Compute/virtualMachines or Microsoft.Storage/storageAccounts.
     */
    readonly type: string;
}