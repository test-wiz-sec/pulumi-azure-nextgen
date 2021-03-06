// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../../types/input";
import * as outputs from "../../types/output";
import * as utilities from "../../utilities";

export function getWebAppSiteExtension(args: GetWebAppSiteExtensionArgs, opts?: pulumi.InvokeOptions): Promise<GetWebAppSiteExtensionResult> {
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    return pulumi.runtime.invoke("azure-nextgen:web/v20160801:getWebAppSiteExtension", {
        "name": args.name,
        "resourceGroupName": args.resourceGroupName,
        "siteExtensionId": args.siteExtensionId,
    }, opts);
}

export interface GetWebAppSiteExtensionArgs {
    /**
     * Site name.
     */
    readonly name: string;
    /**
     * Name of the resource group to which the resource belongs.
     */
    readonly resourceGroupName: string;
    /**
     * Site extension name.
     */
    readonly siteExtensionId: string;
}

/**
 * Site Extension Information.
 */
export interface GetWebAppSiteExtensionResult {
    /**
     * List of authors.
     */
    readonly authors?: string[];
    /**
     * Site Extension comment.
     */
    readonly comment?: string;
    /**
     * Detailed description.
     */
    readonly description?: string;
    /**
     * Count of downloads.
     */
    readonly downloadCount?: number;
    /**
     * Extension URL.
     */
    readonly extensionUrl?: string;
    /**
     * Feed URL.
     */
    readonly feedUrl?: string;
    /**
     * Icon URL.
     */
    readonly iconUrl?: string;
    /**
     * Installer command line parameters.
     */
    readonly installationArgs?: string;
    /**
     * Installed timestamp.
     */
    readonly installedDateTime?: string;
    /**
     * Kind of resource.
     */
    readonly kind?: string;
    /**
     * License URL.
     */
    readonly licenseUrl?: string;
    /**
     * <code>true</code> if the local version is the latest version; <code>false</code> otherwise.
     */
    readonly localIsLatestVersion?: boolean;
    /**
     * Local path.
     */
    readonly localPath?: string;
    /**
     * Resource Name.
     */
    readonly name: string;
    /**
     * Project URL.
     */
    readonly projectUrl?: string;
    /**
     * Provisioning state.
     */
    readonly provisioningState?: string;
    /**
     * Published timestamp.
     */
    readonly publishedDateTime?: string;
    /**
     * Summary description.
     */
    readonly summary?: string;
    /**
     * Site extension title.
     */
    readonly title?: string;
    /**
     * Resource type.
     */
    readonly type: string;
    /**
     * Version information.
     */
    readonly version?: string;
}
