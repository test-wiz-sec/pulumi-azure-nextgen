// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../../types/input";
import * as outputs from "../../types/output";
import * as utilities from "../../utilities";

export function getFileServer(args: GetFileServerArgs, opts?: pulumi.InvokeOptions): Promise<GetFileServerResult> {
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    return pulumi.runtime.invoke("azure-nextgen:batchai/latest:getFileServer", {
        "fileServerName": args.fileServerName,
        "resourceGroupName": args.resourceGroupName,
        "workspaceName": args.workspaceName,
    }, opts);
}

export interface GetFileServerArgs {
    /**
     * The name of the file server within the specified resource group. File server names can only contain a combination of alphanumeric characters along with dash (-) and underscore (_). The name must be from 1 through 64 characters long.
     */
    readonly fileServerName: string;
    /**
     * Name of the resource group to which the resource belongs.
     */
    readonly resourceGroupName: string;
    /**
     * The name of the workspace. Workspace names can only contain a combination of alphanumeric characters along with dash (-) and underscore (_). The name must be from 1 through 64 characters long.
     */
    readonly workspaceName: string;
}

/**
 * File Server information.
 */
export interface GetFileServerResult {
    /**
     * Time when the FileServer was created.
     */
    readonly creationTime: string;
    /**
     * Information about disks attached to File Server VM.
     */
    readonly dataDisks?: outputs.batchai.latest.DataDisksResponse;
    /**
     * File Server mount settings.
     */
    readonly mountSettings: outputs.batchai.latest.MountSettingsResponse;
    /**
     * The name of the resource.
     */
    readonly name: string;
    /**
     * Provisioning state of the File Server. Possible values: creating - The File Server is getting created; updating - The File Server creation has been accepted and it is getting updated; deleting - The user has requested that the File Server be deleted, and it is in the process of being deleted; failed - The File Server creation has failed with the specified error code. Details about the error code are specified in the message field; succeeded - The File Server creation has succeeded.
     */
    readonly provisioningState: string;
    /**
     * Time when the provisioning state was changed.
     */
    readonly provisioningStateTransitionTime: string;
    /**
     * SSH configuration for accessing the File Server node.
     */
    readonly sshConfiguration?: outputs.batchai.latest.SshConfigurationResponse;
    /**
     * File Server virtual network subnet resource ID.
     */
    readonly subnet?: outputs.batchai.latest.ResourceIdResponse;
    /**
     * The type of the resource.
     */
    readonly type: string;
    /**
     * VM size of the File Server.
     */
    readonly vmSize?: string;
}
