// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../../types/input";
import * as outputs from "../../types/output";
import * as utilities from "../../utilities";

export function getAssessment(args: GetAssessmentArgs, opts?: pulumi.InvokeOptions): Promise<GetAssessmentResult> {
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    return pulumi.runtime.invoke("azure-nextgen:security/v20190101preview:getAssessment", {
        "assessmentName": args.assessmentName,
        "expand": args.expand,
        "resourceId": args.resourceId,
    }, opts);
}

export interface GetAssessmentArgs {
    /**
     * The Assessment Key - Unique key for the assessment type
     */
    readonly assessmentName: string;
    /**
     * OData expand. Optional.
     */
    readonly expand?: string;
    /**
     * The identifier of the resource.
     */
    readonly resourceId: string;
}

/**
 * Security assessment on a resource
 */
export interface GetAssessmentResult {
    /**
     * Additional data regarding the assessment
     */
    readonly additionalData?: {[key: string]: string};
    /**
     * User friendly display name of the assessment
     */
    readonly displayName: string;
    /**
     * Links relevant to the assessment
     */
    readonly links?: outputs.security.v20190101preview.AssessmentLinksResponse;
    /**
     * Resource name
     */
    readonly name: string;
    /**
     * Details of the resource that was assessed
     */
    readonly resourceDetails: outputs.security.v20190101preview.AzureResourceDetailsResponse | outputs.security.v20190101preview.OnPremiseResourceDetailsResponse;
    /**
     * The result of the assessment
     */
    readonly status: outputs.security.v20190101preview.AssessmentStatusResponse;
    /**
     * Resource type
     */
    readonly type: string;
}