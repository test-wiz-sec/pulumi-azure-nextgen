// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../../types/input";
import * as outputs from "../../types/output";
import * as utilities from "../../utilities";

export function listConnectionConsentLinks(args: ListConnectionConsentLinksArgs, opts?: pulumi.InvokeOptions): Promise<ListConnectionConsentLinksResult> {
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    return pulumi.runtime.invoke("azure-nextgen:web/v20160601:listConnectionConsentLinks", {
        "connectionName": args.connectionName,
        "parameters": args.parameters,
        "resourceGroupName": args.resourceGroupName,
    }, opts);
}

export interface ListConnectionConsentLinksArgs {
    /**
     * Connection name
     */
    readonly connectionName: string;
    /**
     * Collection of resources
     */
    readonly parameters?: inputs.web.v20160601.ConsentLinkParameterDefinition[];
    /**
     * The resource group
     */
    readonly resourceGroupName: string;
}

/**
 * Collection of consent links
 */
export interface ListConnectionConsentLinksResult {
    /**
     * Collection of resources
     */
    readonly value?: outputs.web.v20160601.ConsentLinkDefinitionResponse[];
}
