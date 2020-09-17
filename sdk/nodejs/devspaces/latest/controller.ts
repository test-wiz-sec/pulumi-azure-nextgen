// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../../types/input";
import * as outputs from "../../types/output";
import * as utilities from "../../utilities";

export class Controller extends pulumi.CustomResource {
    /**
     * Get an existing Controller resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): Controller {
        return new Controller(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'azure-nextgen:devspaces/latest:Controller';

    /**
     * Returns true if the given object is an instance of Controller.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Controller {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Controller.__pulumiType;
    }

    /**
     * DNS name for accessing DataPlane services
     */
    public /*out*/ readonly dataPlaneFqdn!: pulumi.Output<string>;
    /**
     * DNS suffix for public endpoints running in the Azure Dev Spaces Controller.
     */
    public /*out*/ readonly hostSuffix!: pulumi.Output<string>;
    /**
     * Region where the Azure resource is located.
     */
    public readonly location!: pulumi.Output<string>;
    /**
     * The name of the resource.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * Provisioning state of the Azure Dev Spaces Controller.
     */
    public /*out*/ readonly provisioningState!: pulumi.Output<string>;
    /**
     * Model representing SKU for Azure Dev Spaces Controller.
     */
    public readonly sku!: pulumi.Output<outputs.devspaces.latest.SkuResponse>;
    /**
     * Tags for the Azure resource.
     */
    public readonly tags!: pulumi.Output<{[key: string]: string} | undefined>;
    /**
     * DNS of the target container host's API server
     */
    public /*out*/ readonly targetContainerHostApiServerFqdn!: pulumi.Output<string>;
    /**
     * Credentials of the target container host (base64).
     */
    public readonly targetContainerHostCredentialsBase64!: pulumi.Output<string>;
    /**
     * Resource ID of the target container host
     */
    public readonly targetContainerHostResourceId!: pulumi.Output<string>;
    /**
     * The type of the resource.
     */
    public /*out*/ readonly type!: pulumi.Output<string>;

    /**
     * Create a Controller resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: ControllerArgs, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (!(opts && opts.id)) {
            if (!args || args.location === undefined) {
                throw new Error("Missing required property 'location'");
            }
            if (!args || args.name === undefined) {
                throw new Error("Missing required property 'name'");
            }
            if (!args || args.resourceGroupName === undefined) {
                throw new Error("Missing required property 'resourceGroupName'");
            }
            if (!args || args.sku === undefined) {
                throw new Error("Missing required property 'sku'");
            }
            if (!args || args.targetContainerHostCredentialsBase64 === undefined) {
                throw new Error("Missing required property 'targetContainerHostCredentialsBase64'");
            }
            if (!args || args.targetContainerHostResourceId === undefined) {
                throw new Error("Missing required property 'targetContainerHostResourceId'");
            }
            inputs["location"] = args ? args.location : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["resourceGroupName"] = args ? args.resourceGroupName : undefined;
            inputs["sku"] = args ? args.sku : undefined;
            inputs["tags"] = args ? args.tags : undefined;
            inputs["targetContainerHostCredentialsBase64"] = args ? args.targetContainerHostCredentialsBase64 : undefined;
            inputs["targetContainerHostResourceId"] = args ? args.targetContainerHostResourceId : undefined;
            inputs["dataPlaneFqdn"] = undefined /*out*/;
            inputs["hostSuffix"] = undefined /*out*/;
            inputs["provisioningState"] = undefined /*out*/;
            inputs["targetContainerHostApiServerFqdn"] = undefined /*out*/;
            inputs["type"] = undefined /*out*/;
        } else {
            inputs["dataPlaneFqdn"] = undefined /*out*/;
            inputs["hostSuffix"] = undefined /*out*/;
            inputs["location"] = undefined /*out*/;
            inputs["name"] = undefined /*out*/;
            inputs["provisioningState"] = undefined /*out*/;
            inputs["sku"] = undefined /*out*/;
            inputs["tags"] = undefined /*out*/;
            inputs["targetContainerHostApiServerFqdn"] = undefined /*out*/;
            inputs["targetContainerHostCredentialsBase64"] = undefined /*out*/;
            inputs["targetContainerHostResourceId"] = undefined /*out*/;
            inputs["type"] = undefined /*out*/;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        const aliasOpts = { aliases: [{ type: "azure-nextgen:devspaces/v20190401:Controller" }] };
        opts = opts ? pulumi.mergeOptions(opts, aliasOpts) : aliasOpts;
        super(Controller.__pulumiType, name, inputs, opts);
    }
}

/**
 * The set of arguments for constructing a Controller resource.
 */
export interface ControllerArgs {
    /**
     * Region where the Azure resource is located.
     */
    readonly location: pulumi.Input<string>;
    /**
     * Name of the resource.
     */
    readonly name: pulumi.Input<string>;
    /**
     * Resource group to which the resource belongs.
     */
    readonly resourceGroupName: pulumi.Input<string>;
    /**
     * Model representing SKU for Azure Dev Spaces Controller.
     */
    readonly sku: pulumi.Input<inputs.devspaces.latest.Sku>;
    /**
     * Tags for the Azure resource.
     */
    readonly tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * Credentials of the target container host (base64).
     */
    readonly targetContainerHostCredentialsBase64: pulumi.Input<string>;
    /**
     * Resource ID of the target container host
     */
    readonly targetContainerHostResourceId: pulumi.Input<string>;
}