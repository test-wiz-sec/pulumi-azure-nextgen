// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../../types/input";
import * as outputs from "../../types/output";
import * as utilities from "../../utilities";

export function getProject(args: GetProjectArgs, opts?: pulumi.InvokeOptions): Promise<GetProjectResult> {
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    return pulumi.runtime.invoke("azure-nextgen:migrate/v20171111preview:getProject", {
        "projectName": args.projectName,
        "resourceGroupName": args.resourceGroupName,
    }, opts);
}

export interface GetProjectArgs {
    /**
     * Name of the Azure Migrate project.
     */
    readonly projectName: string;
    /**
     * Name of the Azure Resource Group that project is part of.
     */
    readonly resourceGroupName: string;
}

/**
 * Azure Migrate Project.
 */
export interface GetProjectResult {
    /**
     * Time when this project was created. Date-Time represented in ISO-8601 format.
     */
    readonly createdTimestamp: string;
    /**
     * ARM ID of the Service Map workspace created by user.
     */
    readonly customerWorkspaceId?: string;
    /**
     * Reports whether project is under discovery.
     */
    readonly discoveryStatus: string;
    /**
     * For optimistic concurrency control.
     */
    readonly eTag?: string;
    /**
     * Azure location in which project is created.
     */
    readonly location?: string;
    /**
     * Name of the project.
     */
    readonly name: string;
    /**
     * Number of assessments created in the project.
     */
    readonly numberOfAssessments: number;
    /**
     * Number of groups created in the project.
     */
    readonly numberOfGroups: number;
    /**
     * Number of machines in the project.
     */
    readonly numberOfMachines: number;
    /**
     * Provisioning state of the project.
     */
    readonly provisioningState?: string;
    /**
     * Tags provided by Azure Tagging service.
     */
    readonly tags?: any;
    /**
     * Type of the object = [Microsoft.Migrate/projects].
     */
    readonly type: string;
    /**
     * Time when this project was last updated. Date-Time represented in ISO-8601 format.
     */
    readonly updatedTimestamp: string;
}
