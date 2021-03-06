// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../../types/input";
import * as outputs from "../../types/output";
import * as utilities from "../../utilities";

export function listTopLevelDomainAgreements(args: ListTopLevelDomainAgreementsArgs, opts?: pulumi.InvokeOptions): Promise<ListTopLevelDomainAgreementsResult> {
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    return pulumi.runtime.invoke("azure-nextgen:domainregistration/v20180201:listTopLevelDomainAgreements", {
        "forTransfer": args.forTransfer,
        "includePrivacy": args.includePrivacy,
        "name": args.name,
    }, opts);
}

export interface ListTopLevelDomainAgreementsArgs {
    /**
     * If <code>true</code>, then the list of agreements will include agreements for domain transfer as well; otherwise, <code>false</code>.
     */
    readonly forTransfer?: boolean;
    /**
     * If <code>true</code>, then the list of agreements will include agreements for domain privacy as well; otherwise, <code>false</code>.
     */
    readonly includePrivacy?: boolean;
    /**
     * Name of the top-level domain.
     */
    readonly name: string;
}

/**
 * Collection of top-level domain legal agreements.
 */
export interface ListTopLevelDomainAgreementsResult {
    /**
     * Link to next page of resources.
     */
    readonly nextLink: string;
    /**
     * Collection of resources.
     */
    readonly value: outputs.domainregistration.v20180201.TldLegalAgreementResponse[];
}
