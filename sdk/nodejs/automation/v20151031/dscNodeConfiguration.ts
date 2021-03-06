// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../../types/input";
import * as outputs from "../../types/output";
import * as utilities from "../../utilities";

/**
 * Definition of the dsc node configuration.
 */
export class DscNodeConfiguration extends pulumi.CustomResource {
    /**
     * Get an existing DscNodeConfiguration resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): DscNodeConfiguration {
        return new DscNodeConfiguration(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'azure-nextgen:automation/v20151031:DscNodeConfiguration';

    /**
     * Returns true if the given object is an instance of DscNodeConfiguration.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is DscNodeConfiguration {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === DscNodeConfiguration.__pulumiType;
    }

    /**
     * Gets or sets the configuration of the node.
     */
    public readonly configuration!: pulumi.Output<outputs.automation.v20151031.DscConfigurationAssociationPropertyResponse | undefined>;
    /**
     * Gets or sets creation time.
     */
    public /*out*/ readonly creationTime!: pulumi.Output<string | undefined>;
    /**
     * Gets or sets the last modified time.
     */
    public /*out*/ readonly lastModifiedTime!: pulumi.Output<string | undefined>;
    /**
     * The name of the resource
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * The type of the resource.
     */
    public /*out*/ readonly type!: pulumi.Output<string>;

    /**
     * Create a DscNodeConfiguration resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: DscNodeConfigurationArgs, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (!(opts && opts.id)) {
            if (!args || args.automationAccountName === undefined) {
                throw new Error("Missing required property 'automationAccountName'");
            }
            if (!args || args.configuration === undefined) {
                throw new Error("Missing required property 'configuration'");
            }
            if (!args || args.name === undefined) {
                throw new Error("Missing required property 'name'");
            }
            if (!args || args.nodeConfigurationName === undefined) {
                throw new Error("Missing required property 'nodeConfigurationName'");
            }
            if (!args || args.resourceGroupName === undefined) {
                throw new Error("Missing required property 'resourceGroupName'");
            }
            if (!args || args.source === undefined) {
                throw new Error("Missing required property 'source'");
            }
            inputs["automationAccountName"] = args ? args.automationAccountName : undefined;
            inputs["configuration"] = args ? args.configuration : undefined;
            inputs["incrementNodeConfigurationBuild"] = args ? args.incrementNodeConfigurationBuild : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["nodeConfigurationName"] = args ? args.nodeConfigurationName : undefined;
            inputs["resourceGroupName"] = args ? args.resourceGroupName : undefined;
            inputs["source"] = args ? args.source : undefined;
            inputs["creationTime"] = undefined /*out*/;
            inputs["lastModifiedTime"] = undefined /*out*/;
            inputs["type"] = undefined /*out*/;
        } else {
            inputs["configuration"] = undefined /*out*/;
            inputs["creationTime"] = undefined /*out*/;
            inputs["lastModifiedTime"] = undefined /*out*/;
            inputs["name"] = undefined /*out*/;
            inputs["type"] = undefined /*out*/;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        const aliasOpts = { aliases: [{ type: "azure-nextgen:automation/latest:DscNodeConfiguration" }, { type: "azure-nextgen:automation/v20180115:DscNodeConfiguration" }] };
        opts = opts ? pulumi.mergeOptions(opts, aliasOpts) : aliasOpts;
        super(DscNodeConfiguration.__pulumiType, name, inputs, opts);
    }
}

/**
 * The set of arguments for constructing a DscNodeConfiguration resource.
 */
export interface DscNodeConfigurationArgs {
    /**
     * The name of the automation account.
     */
    readonly automationAccountName: pulumi.Input<string>;
    /**
     * Gets or sets the configuration of the node.
     */
    readonly configuration: pulumi.Input<inputs.automation.v20151031.DscConfigurationAssociationProperty>;
    /**
     * If a new build version of NodeConfiguration is required.
     */
    readonly incrementNodeConfigurationBuild?: pulumi.Input<boolean>;
    /**
     * Name of the node configuration.
     */
    readonly name: pulumi.Input<string>;
    /**
     * The create or update parameters for configuration.
     */
    readonly nodeConfigurationName: pulumi.Input<string>;
    /**
     * Name of an Azure Resource group.
     */
    readonly resourceGroupName: pulumi.Input<string>;
    /**
     * Gets or sets the source.
     */
    readonly source: pulumi.Input<inputs.automation.v20151031.ContentSource>;
}
