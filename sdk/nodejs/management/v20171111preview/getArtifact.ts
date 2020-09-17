// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../../types/input";
import * as outputs from "../../types/output";
import * as utilities from "../../utilities";

export function getArtifact(args: GetArtifactArgs, opts?: pulumi.InvokeOptions): Promise<GetArtifactResult> {
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    return pulumi.runtime.invoke("azure-nextgen:management/v20171111preview:getArtifact", {
        "artifactName": args.artifactName,
        "blueprintName": args.blueprintName,
        "managementGroupName": args.managementGroupName,
    }, opts);
}

export interface GetArtifactArgs {
    /**
     * name of the artifact.
     */
    readonly artifactName: string;
    /**
     * name of the blueprint.
     */
    readonly blueprintName: string;
    /**
     * ManagementGroup where blueprint stores.
     */
    readonly managementGroupName: string;
}

/**
 * Represents a Blueprint artifact.
 */
export interface GetArtifactResult {
    /**
     * Specifies the kind of Blueprint artifact.
     */
    readonly kind: string;
    /**
     * Name of this resource.
     */
    readonly name: string;
    /**
     * Type of this resource.
     */
    readonly type: string;
}