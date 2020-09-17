// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../../types/input";
import * as outputs from "../../types/output";
import * as utilities from "../../utilities";

export function getSapMonitor(args: GetSapMonitorArgs, opts?: pulumi.InvokeOptions): Promise<GetSapMonitorResult> {
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    return pulumi.runtime.invoke("azure-nextgen:hanaonazure/v20171103preview:getSapMonitor", {
        "resourceGroupName": args.resourceGroupName,
        "sapMonitorName": args.sapMonitorName,
    }, opts);
}

export interface GetSapMonitorArgs {
    /**
     * Name of the resource group.
     */
    readonly resourceGroupName: string;
    /**
     * Name of the SAP monitor resource.
     */
    readonly sapMonitorName: string;
}

/**
 * SAP monitor info on Azure (ARM properties and SAP monitor properties)
 */
export interface GetSapMonitorResult {
    /**
     * The value indicating whether to send analytics to Microsoft
     */
    readonly enableCustomerAnalytics?: boolean;
    /**
     * MSI ID passed by customer which has access to customer's KeyVault and to be assigned to the Collector VM.
     */
    readonly hanaDbCredentialsMsiId?: string;
    /**
     * Database name of the HANA instance.
     */
    readonly hanaDbName?: string;
    /**
     * Database password of the HANA instance.
     */
    readonly hanaDbPassword?: string;
    /**
     * KeyVault URL link to the password for the HANA database.
     */
    readonly hanaDbPasswordKeyVaultUrl?: string;
    /**
     * Database port of the HANA instance.
     */
    readonly hanaDbSqlPort?: number;
    /**
     * Database username of the HANA instance.
     */
    readonly hanaDbUsername?: string;
    /**
     * Hostname of the HANA instance.
     */
    readonly hanaHostname?: string;
    /**
     * Specifies the SAP monitor unique ID.
     */
    readonly hanaSubnet?: string;
    /**
     * Key Vault ID containing customer's HANA credentials.
     */
    readonly keyVaultId?: string;
    /**
     * Resource location
     */
    readonly location?: string;
    /**
     * The ARM ID of the Log Analytics Workspace that is used for monitoring
     */
    readonly logAnalyticsWorkspaceArmId?: string;
    /**
     * The workspace ID of the log analytics workspace to be used for monitoring
     */
    readonly logAnalyticsWorkspaceId?: string;
    /**
     * The shared key of the log analytics workspace that is used for monitoring
     */
    readonly logAnalyticsWorkspaceSharedKey?: string;
    /**
     * The name of the resource group the SAP Monitor resources get deployed into.
     */
    readonly managedResourceGroupName: string;
    /**
     * Resource name
     */
    readonly name: string;
    /**
     * State of provisioning of the HanaInstance
     */
    readonly provisioningState: string;
    /**
     * Resource tags
     */
    readonly tags: {[key: string]: string};
    /**
     * Resource type
     */
    readonly type: string;
}