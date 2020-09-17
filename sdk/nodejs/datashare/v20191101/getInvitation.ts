// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../../types/input";
import * as outputs from "../../types/output";
import * as utilities from "../../utilities";

export function getInvitation(args: GetInvitationArgs, opts?: pulumi.InvokeOptions): Promise<GetInvitationResult> {
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    return pulumi.runtime.invoke("azure-nextgen:datashare/v20191101:getInvitation", {
        "accountName": args.accountName,
        "invitationName": args.invitationName,
        "resourceGroupName": args.resourceGroupName,
        "shareName": args.shareName,
    }, opts);
}

export interface GetInvitationArgs {
    /**
     * The name of the share account.
     */
    readonly accountName: string;
    /**
     * The name of the invitation.
     */
    readonly invitationName: string;
    /**
     * The resource group name.
     */
    readonly resourceGroupName: string;
    /**
     * The name of the share.
     */
    readonly shareName: string;
}

/**
 * A Invitation data transfer object.
 */
export interface GetInvitationResult {
    /**
     * unique invitation id
     */
    readonly invitationId: string;
    /**
     * The status of the invitation.
     */
    readonly invitationStatus: string;
    /**
     * Name of the azure resource
     */
    readonly name: string;
    /**
     * The time the recipient responded to the invitation.
     */
    readonly respondedAt: string;
    /**
     * Gets the time at which the invitation was sent.
     */
    readonly sentAt: string;
    /**
     * The target Azure AD Id. Can't be combined with email.
     */
    readonly targetActiveDirectoryId?: string;
    /**
     * The email the invitation is directed to.
     */
    readonly targetEmail?: string;
    /**
     * The target user or application Id that invitation is being sent to.
     * Must be specified along TargetActiveDirectoryId. This enables sending
     * invitations to specific users or applications in an AD tenant.
     */
    readonly targetObjectId?: string;
    /**
     * Type of the azure resource
     */
    readonly type: string;
    /**
     * Email of the user who created the resource
     */
    readonly userEmail: string;
    /**
     * Name of the user who created the resource
     */
    readonly userName: string;
}