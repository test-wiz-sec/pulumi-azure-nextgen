// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../../types/input";
import * as outputs from "../../types/output";
import * as utilities from "../../utilities";

/**
 * A report config resource.
 */
export class ReportConfig extends pulumi.CustomResource {
    /**
     * Get an existing ReportConfig resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): ReportConfig {
        return new ReportConfig(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'azure-nextgen:costmanagement/v20180531:ReportConfig';

    /**
     * Returns true if the given object is an instance of ReportConfig.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is ReportConfig {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === ReportConfig.__pulumiType;
    }

    /**
     * Has definition for the report config.
     */
    public readonly definition!: pulumi.Output<outputs.costmanagement.v20180531.ReportConfigDefinitionResponse>;
    /**
     * Has delivery information for the report config.
     */
    public readonly deliveryInfo!: pulumi.Output<outputs.costmanagement.v20180531.ReportConfigDeliveryInfoResponse>;
    /**
     * The format of the report being delivered.
     */
    public readonly format!: pulumi.Output<string | undefined>;
    /**
     * Resource name.
     */
    public /*out*/ readonly name!: pulumi.Output<string>;
    /**
     * Has schedule information for the report config.
     */
    public readonly schedule!: pulumi.Output<outputs.costmanagement.v20180531.ReportConfigScheduleResponse | undefined>;
    /**
     * Resource tags.
     */
    public /*out*/ readonly tags!: pulumi.Output<{[key: string]: string}>;
    /**
     * Resource type.
     */
    public /*out*/ readonly type!: pulumi.Output<string>;

    /**
     * Create a ReportConfig resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: ReportConfigArgs, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (!(opts && opts.id)) {
            if (!args || args.definition === undefined) {
                throw new Error("Missing required property 'definition'");
            }
            if (!args || args.deliveryInfo === undefined) {
                throw new Error("Missing required property 'deliveryInfo'");
            }
            if (!args || args.reportConfigName === undefined) {
                throw new Error("Missing required property 'reportConfigName'");
            }
            inputs["definition"] = args ? args.definition : undefined;
            inputs["deliveryInfo"] = args ? args.deliveryInfo : undefined;
            inputs["format"] = args ? args.format : undefined;
            inputs["reportConfigName"] = args ? args.reportConfigName : undefined;
            inputs["schedule"] = args ? args.schedule : undefined;
            inputs["name"] = undefined /*out*/;
            inputs["tags"] = undefined /*out*/;
            inputs["type"] = undefined /*out*/;
        } else {
            inputs["definition"] = undefined /*out*/;
            inputs["deliveryInfo"] = undefined /*out*/;
            inputs["format"] = undefined /*out*/;
            inputs["name"] = undefined /*out*/;
            inputs["schedule"] = undefined /*out*/;
            inputs["tags"] = undefined /*out*/;
            inputs["type"] = undefined /*out*/;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        const aliasOpts = { aliases: [{ type: "azure-nextgen:costmanagement/latest:ReportConfig" }] };
        opts = opts ? pulumi.mergeOptions(opts, aliasOpts) : aliasOpts;
        super(ReportConfig.__pulumiType, name, inputs, opts);
    }
}

/**
 * The set of arguments for constructing a ReportConfig resource.
 */
export interface ReportConfigArgs {
    /**
     * Has definition for the report config.
     */
    readonly definition: pulumi.Input<inputs.costmanagement.v20180531.ReportConfigDefinition>;
    /**
     * Has delivery information for the report config.
     */
    readonly deliveryInfo: pulumi.Input<inputs.costmanagement.v20180531.ReportConfigDeliveryInfo>;
    /**
     * The format of the report being delivered.
     */
    readonly format?: pulumi.Input<string>;
    /**
     * Report Config Name.
     */
    readonly reportConfigName: pulumi.Input<string>;
    /**
     * Has schedule information for the report config.
     */
    readonly schedule?: pulumi.Input<inputs.costmanagement.v20180531.ReportConfigSchedule>;
}