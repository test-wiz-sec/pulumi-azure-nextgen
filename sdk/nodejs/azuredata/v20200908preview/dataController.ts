// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../../types/input";
import * as outputs from "../../types/output";
import * as utilities from "../../utilities";

/**
 * Data controller resource
 */
export class DataController extends pulumi.CustomResource {
    /**
     * Get an existing DataController resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): DataController {
        return new DataController(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'azure-nextgen:azuredata/v20200908preview:DataController';

    /**
     * Returns true if the given object is an instance of DataController.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is DataController {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === DataController.__pulumiType;
    }

    /**
     * The raw kubernetes information
     */
    public readonly k8sRaw!: pulumi.Output<{[key: string]: any} | undefined>;
    /**
     * Last uploaded date from on premise cluster. Defaults to current date time
     */
    public readonly lastUploadedDate!: pulumi.Output<string | undefined>;
    /**
     * The geo-location where the resource lives
     */
    public readonly location!: pulumi.Output<string>;
    /**
     * The name of the resource
     */
    public /*out*/ readonly name!: pulumi.Output<string>;
    /**
     * Properties from the on premise data controller
     */
    public readonly onPremiseProperty!: pulumi.Output<outputs.azuredata.v20200908preview.OnPremisePropertyResponse>;
    /**
     * Read only system data
     */
    public /*out*/ readonly systemData!: pulumi.Output<outputs.azuredata.v20200908preview.SystemDataResponse>;
    /**
     * Resource tags.
     */
    public readonly tags!: pulumi.Output<{[key: string]: string} | undefined>;
    /**
     * The type of the resource. Ex- Microsoft.Compute/virtualMachines or Microsoft.Storage/storageAccounts.
     */
    public /*out*/ readonly type!: pulumi.Output<string>;

    /**
     * Create a DataController resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: DataControllerArgs, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (!(opts && opts.id)) {
            if (!args || args.dataControllerName === undefined) {
                throw new Error("Missing required property 'dataControllerName'");
            }
            if (!args || args.location === undefined) {
                throw new Error("Missing required property 'location'");
            }
            if (!args || args.onPremiseProperty === undefined) {
                throw new Error("Missing required property 'onPremiseProperty'");
            }
            if (!args || args.resourceGroupName === undefined) {
                throw new Error("Missing required property 'resourceGroupName'");
            }
            inputs["dataControllerName"] = args ? args.dataControllerName : undefined;
            inputs["k8sRaw"] = args ? args.k8sRaw : undefined;
            inputs["lastUploadedDate"] = args ? args.lastUploadedDate : undefined;
            inputs["location"] = args ? args.location : undefined;
            inputs["onPremiseProperty"] = args ? args.onPremiseProperty : undefined;
            inputs["resourceGroupName"] = args ? args.resourceGroupName : undefined;
            inputs["tags"] = args ? args.tags : undefined;
            inputs["name"] = undefined /*out*/;
            inputs["systemData"] = undefined /*out*/;
            inputs["type"] = undefined /*out*/;
        } else {
            inputs["k8sRaw"] = undefined /*out*/;
            inputs["lastUploadedDate"] = undefined /*out*/;
            inputs["location"] = undefined /*out*/;
            inputs["name"] = undefined /*out*/;
            inputs["onPremiseProperty"] = undefined /*out*/;
            inputs["systemData"] = undefined /*out*/;
            inputs["tags"] = undefined /*out*/;
            inputs["type"] = undefined /*out*/;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        const aliasOpts = { aliases: [{ type: "azure-nextgen:azuredata/v20190724preview:DataController" }] };
        opts = opts ? pulumi.mergeOptions(opts, aliasOpts) : aliasOpts;
        super(DataController.__pulumiType, name, inputs, opts);
    }
}

/**
 * The set of arguments for constructing a DataController resource.
 */
export interface DataControllerArgs {
    readonly dataControllerName: pulumi.Input<string>;
    /**
     * The raw kubernetes information
     */
    readonly k8sRaw?: pulumi.Input<{[key: string]: any}>;
    /**
     * Last uploaded date from on premise cluster. Defaults to current date time
     */
    readonly lastUploadedDate?: pulumi.Input<string>;
    /**
     * The geo-location where the resource lives
     */
    readonly location: pulumi.Input<string>;
    /**
     * Properties from the on premise data controller
     */
    readonly onPremiseProperty: pulumi.Input<inputs.azuredata.v20200908preview.OnPremiseProperty>;
    /**
     * The name of the Azure resource group
     */
    readonly resourceGroupName: pulumi.Input<string>;
    /**
     * Resource tags.
     */
    readonly tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
}
