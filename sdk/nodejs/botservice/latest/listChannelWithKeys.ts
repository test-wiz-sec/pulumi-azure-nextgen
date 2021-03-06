// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../../types/input";
import * as outputs from "../../types/output";
import * as utilities from "../../utilities";

export function listChannelWithKeys(args: ListChannelWithKeysArgs, opts?: pulumi.InvokeOptions): Promise<ListChannelWithKeysResult> {
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    return pulumi.runtime.invoke("azure-nextgen:botservice/latest:listChannelWithKeys", {
        "channelName": args.channelName,
        "resourceGroupName": args.resourceGroupName,
        "resourceName": args.resourceName,
    }, opts);
}

export interface ListChannelWithKeysArgs {
    /**
     * The name of the Channel resource.
     */
    readonly channelName: string;
    /**
     * The name of the Bot resource group in the user subscription.
     */
    readonly resourceGroupName: string;
    /**
     * The name of the Bot resource.
     */
    readonly resourceName: string;
}

/**
 * Bot channel resource definition
 */
export interface ListChannelWithKeysResult {
    /**
     * Entity Tag
     */
    readonly etag?: string;
    /**
     * Required. Gets or sets the Kind of the resource.
     */
    readonly kind?: string;
    /**
     * Specifies the location of the resource.
     */
    readonly location?: string;
    /**
     * Specifies the name of the resource.
     */
    readonly name: string;
    /**
     * The set of properties specific to bot channel resource
     */
    readonly properties: outputs.botservice.latest.AlexaChannelResponse | outputs.botservice.latest.DirectLineChannelResponse | outputs.botservice.latest.DirectLineSpeechChannelResponse | outputs.botservice.latest.EmailChannelResponse | outputs.botservice.latest.FacebookChannelResponse | outputs.botservice.latest.KikChannelResponse | outputs.botservice.latest.LineChannelResponse | outputs.botservice.latest.MsTeamsChannelResponse | outputs.botservice.latest.SkypeChannelResponse | outputs.botservice.latest.SlackChannelResponse | outputs.botservice.latest.SmsChannelResponse | outputs.botservice.latest.TelegramChannelResponse | outputs.botservice.latest.WebChatChannelResponse;
    /**
     * Gets or sets the SKU of the resource.
     */
    readonly sku?: outputs.botservice.latest.SkuResponse;
    /**
     * Contains resource tags defined as key/value pairs.
     */
    readonly tags?: {[key: string]: string};
    /**
     * Specifies the type of the resource.
     */
    readonly type: string;
}
