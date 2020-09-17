// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../../types/input";
import * as outputs from "../../types/output";
import * as utilities from "../../utilities";

/**
 * A manifest file that defines the custom resource provider resources.
 */
export class CustomResourceProvider extends pulumi.CustomResource {
    /**
     * Get an existing CustomResourceProvider resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): CustomResourceProvider {
        return new CustomResourceProvider(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'azure-nextgen:customproviders/v20180901preview:CustomResourceProvider';

    /**
     * Returns true if the given object is an instance of CustomResourceProvider.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is CustomResourceProvider {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === CustomResourceProvider.__pulumiType;
    }

    /**
     * A list of actions that the custom resource provider implements.
     */
    public readonly actions!: pulumi.Output<outputs.customproviders.v20180901preview.CustomRPActionRouteDefinitionResponse[] | undefined>;
    /**
     * Resource location
     */
    public readonly location!: pulumi.Output<string>;
    /**
     * Resource name
     */
    public /*out*/ readonly name!: pulumi.Output<string>;
    /**
     * The provisioning state of the resource provider.
     */
    public /*out*/ readonly provisioningState!: pulumi.Output<string>;
    /**
     * A list of resource types that the custom resource provider implements.
     */
    public readonly resourceTypes!: pulumi.Output<outputs.customproviders.v20180901preview.CustomRPResourceTypeRouteDefinitionResponse[] | undefined>;
    /**
     * Resource tags
     */
    public readonly tags!: pulumi.Output<{[key: string]: string} | undefined>;
    /**
     * Resource type
     */
    public /*out*/ readonly type!: pulumi.Output<string>;
    /**
     * A list of validations to run on the custom resource provider's requests.
     */
    public readonly validations!: pulumi.Output<outputs.customproviders.v20180901preview.CustomRPValidationsResponse[] | undefined>;

    /**
     * Create a CustomResourceProvider resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: CustomResourceProviderArgs, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (!(opts && opts.id)) {
            if (!args || args.location === undefined) {
                throw new Error("Missing required property 'location'");
            }
            if (!args || args.resourceGroupName === undefined) {
                throw new Error("Missing required property 'resourceGroupName'");
            }
            if (!args || args.resourceProviderName === undefined) {
                throw new Error("Missing required property 'resourceProviderName'");
            }
            inputs["actions"] = args ? args.actions : undefined;
            inputs["location"] = args ? args.location : undefined;
            inputs["resourceGroupName"] = args ? args.resourceGroupName : undefined;
            inputs["resourceProviderName"] = args ? args.resourceProviderName : undefined;
            inputs["resourceTypes"] = args ? args.resourceTypes : undefined;
            inputs["tags"] = args ? args.tags : undefined;
            inputs["validations"] = args ? args.validations : undefined;
            inputs["name"] = undefined /*out*/;
            inputs["provisioningState"] = undefined /*out*/;
            inputs["type"] = undefined /*out*/;
        } else {
            inputs["actions"] = undefined /*out*/;
            inputs["location"] = undefined /*out*/;
            inputs["name"] = undefined /*out*/;
            inputs["provisioningState"] = undefined /*out*/;
            inputs["resourceTypes"] = undefined /*out*/;
            inputs["tags"] = undefined /*out*/;
            inputs["type"] = undefined /*out*/;
            inputs["validations"] = undefined /*out*/;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        super(CustomResourceProvider.__pulumiType, name, inputs, opts);
    }
}

/**
 * The set of arguments for constructing a CustomResourceProvider resource.
 */
export interface CustomResourceProviderArgs {
    /**
     * A list of actions that the custom resource provider implements.
     */
    readonly actions?: pulumi.Input<pulumi.Input<inputs.customproviders.v20180901preview.CustomRPActionRouteDefinition>[]>;
    /**
     * Resource location
     */
    readonly location: pulumi.Input<string>;
    /**
     * The name of the resource group.
     */
    readonly resourceGroupName: pulumi.Input<string>;
    /**
     * The name of the resource provider.
     */
    readonly resourceProviderName: pulumi.Input<string>;
    /**
     * A list of resource types that the custom resource provider implements.
     */
    readonly resourceTypes?: pulumi.Input<pulumi.Input<inputs.customproviders.v20180901preview.CustomRPResourceTypeRouteDefinition>[]>;
    /**
     * Resource tags
     */
    readonly tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * A list of validations to run on the custom resource provider's requests.
     */
    readonly validations?: pulumi.Input<pulumi.Input<inputs.customproviders.v20180901preview.CustomRPValidations>[]>;
}