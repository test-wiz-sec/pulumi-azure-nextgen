// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../../types/input";
import * as outputs from "../../types/output";
import * as utilities from "../../utilities";

export function getExposureControlFeatureValueByFactory(args: GetExposureControlFeatureValueByFactoryArgs, opts?: pulumi.InvokeOptions): Promise<GetExposureControlFeatureValueByFactoryResult> {
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    return pulumi.runtime.invoke("azure-nextgen:datafactory/v20180601:getExposureControlFeatureValueByFactory", {
        "factoryName": args.factoryName,
        "featureName": args.featureName,
        "featureType": args.featureType,
        "resourceGroupName": args.resourceGroupName,
    }, opts);
}

export interface GetExposureControlFeatureValueByFactoryArgs {
    /**
     * The factory name.
     */
    readonly factoryName: string;
    /**
     * The feature name.
     */
    readonly featureName?: string;
    /**
     * The feature type.
     */
    readonly featureType?: string;
    /**
     * The resource group name.
     */
    readonly resourceGroupName: string;
}

/**
 * The exposure control response.
 */
export interface GetExposureControlFeatureValueByFactoryResult {
    /**
     * The feature name.
     */
    readonly featureName: string;
    /**
     * The feature value.
     */
    readonly value: string;
}
