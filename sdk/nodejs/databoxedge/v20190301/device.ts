// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../../types/input";
import * as outputs from "../../types/output";
import * as utilities from "../../utilities";

/**
 * The Data Box Edge/Gateway device.
 */
export class Device extends pulumi.CustomResource {
    /**
     * Get an existing Device resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): Device {
        return new Device(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'azure-nextgen:databoxedge/v20190301:Device';

    /**
     * Returns true if the given object is an instance of Device.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Device {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Device.__pulumiType;
    }

    /**
     * Type of compute roles configured.
     */
    public /*out*/ readonly configuredRoleTypes!: pulumi.Output<string[]>;
    /**
     * The Data Box Edge/Gateway device culture.
     */
    public /*out*/ readonly culture!: pulumi.Output<string>;
    /**
     * The status of the Data Box Edge/Gateway device.
     */
    public readonly dataBoxEdgeDeviceStatus!: pulumi.Output<string | undefined>;
    /**
     * The Description of the Data Box Edge/Gateway device.
     */
    public readonly description!: pulumi.Output<string | undefined>;
    /**
     * The device software version number of the device (eg: 1.2.18105.6).
     */
    public /*out*/ readonly deviceHcsVersion!: pulumi.Output<string>;
    /**
     * The Data Box Edge/Gateway device local capacity in MB.
     */
    public /*out*/ readonly deviceLocalCapacity!: pulumi.Output<number>;
    /**
     * The Data Box Edge/Gateway device model.
     */
    public /*out*/ readonly deviceModel!: pulumi.Output<string>;
    /**
     * The Data Box Edge/Gateway device software version.
     */
    public /*out*/ readonly deviceSoftwareVersion!: pulumi.Output<string>;
    /**
     * The type of the Data Box Edge/Gateway device.
     */
    public /*out*/ readonly deviceType!: pulumi.Output<string>;
    /**
     * The etag for the devices.
     */
    public readonly etag!: pulumi.Output<string | undefined>;
    /**
     * The Data Box Edge/Gateway device name.
     */
    public readonly friendlyName!: pulumi.Output<string | undefined>;
    /**
     * The location of the device. This is a supported and registered Azure geographical region (for example, West US, East US, or Southeast Asia). The geographical region of a device cannot be changed once it is created, but if an identical geographical region is specified on update, the request will succeed.
     */
    public readonly location!: pulumi.Output<string>;
    /**
     * The description of the Data Box Edge/Gateway device model.
     */
    public readonly modelDescription!: pulumi.Output<string | undefined>;
    /**
     * The object name.
     */
    public /*out*/ readonly name!: pulumi.Output<string>;
    /**
     * The Serial Number of Data Box Edge/Gateway device.
     */
    public /*out*/ readonly serialNumber!: pulumi.Output<string>;
    /**
     * The SKU type.
     */
    public readonly sku!: pulumi.Output<outputs.databoxedge.v20190301.SkuResponse | undefined>;
    /**
     * The list of tags that describe the device. These tags can be used to view and group this device (across resource groups).
     */
    public readonly tags!: pulumi.Output<{[key: string]: string} | undefined>;
    /**
     * The Data Box Edge/Gateway device timezone.
     */
    public /*out*/ readonly timeZone!: pulumi.Output<string>;
    /**
     * The hierarchical type of the object.
     */
    public /*out*/ readonly type!: pulumi.Output<string>;

    /**
     * Create a Device resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: DeviceArgs, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (!(opts && opts.id)) {
            if (!args || args.deviceName === undefined) {
                throw new Error("Missing required property 'deviceName'");
            }
            if (!args || args.location === undefined) {
                throw new Error("Missing required property 'location'");
            }
            if (!args || args.resourceGroupName === undefined) {
                throw new Error("Missing required property 'resourceGroupName'");
            }
            inputs["dataBoxEdgeDeviceStatus"] = args ? args.dataBoxEdgeDeviceStatus : undefined;
            inputs["description"] = args ? args.description : undefined;
            inputs["deviceName"] = args ? args.deviceName : undefined;
            inputs["etag"] = args ? args.etag : undefined;
            inputs["friendlyName"] = args ? args.friendlyName : undefined;
            inputs["location"] = args ? args.location : undefined;
            inputs["modelDescription"] = args ? args.modelDescription : undefined;
            inputs["resourceGroupName"] = args ? args.resourceGroupName : undefined;
            inputs["sku"] = args ? args.sku : undefined;
            inputs["tags"] = args ? args.tags : undefined;
            inputs["configuredRoleTypes"] = undefined /*out*/;
            inputs["culture"] = undefined /*out*/;
            inputs["deviceHcsVersion"] = undefined /*out*/;
            inputs["deviceLocalCapacity"] = undefined /*out*/;
            inputs["deviceModel"] = undefined /*out*/;
            inputs["deviceSoftwareVersion"] = undefined /*out*/;
            inputs["deviceType"] = undefined /*out*/;
            inputs["name"] = undefined /*out*/;
            inputs["serialNumber"] = undefined /*out*/;
            inputs["timeZone"] = undefined /*out*/;
            inputs["type"] = undefined /*out*/;
        } else {
            inputs["configuredRoleTypes"] = undefined /*out*/;
            inputs["culture"] = undefined /*out*/;
            inputs["dataBoxEdgeDeviceStatus"] = undefined /*out*/;
            inputs["description"] = undefined /*out*/;
            inputs["deviceHcsVersion"] = undefined /*out*/;
            inputs["deviceLocalCapacity"] = undefined /*out*/;
            inputs["deviceModel"] = undefined /*out*/;
            inputs["deviceSoftwareVersion"] = undefined /*out*/;
            inputs["deviceType"] = undefined /*out*/;
            inputs["etag"] = undefined /*out*/;
            inputs["friendlyName"] = undefined /*out*/;
            inputs["location"] = undefined /*out*/;
            inputs["modelDescription"] = undefined /*out*/;
            inputs["name"] = undefined /*out*/;
            inputs["serialNumber"] = undefined /*out*/;
            inputs["sku"] = undefined /*out*/;
            inputs["tags"] = undefined /*out*/;
            inputs["timeZone"] = undefined /*out*/;
            inputs["type"] = undefined /*out*/;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        const aliasOpts = { aliases: [{ type: "azure-nextgen:databoxedge/latest:Device" }, { type: "azure-nextgen:databoxedge/v20190701:Device" }, { type: "azure-nextgen:databoxedge/v20190801:Device" }, { type: "azure-nextgen:databoxedge/v20200501preview:Device" }] };
        opts = opts ? pulumi.mergeOptions(opts, aliasOpts) : aliasOpts;
        super(Device.__pulumiType, name, inputs, opts);
    }
}

/**
 * The set of arguments for constructing a Device resource.
 */
export interface DeviceArgs {
    /**
     * The status of the Data Box Edge/Gateway device.
     */
    readonly dataBoxEdgeDeviceStatus?: pulumi.Input<string>;
    /**
     * The Description of the Data Box Edge/Gateway device.
     */
    readonly description?: pulumi.Input<string>;
    /**
     * The device name.
     */
    readonly deviceName: pulumi.Input<string>;
    /**
     * The etag for the devices.
     */
    readonly etag?: pulumi.Input<string>;
    /**
     * The Data Box Edge/Gateway device name.
     */
    readonly friendlyName?: pulumi.Input<string>;
    /**
     * The location of the device. This is a supported and registered Azure geographical region (for example, West US, East US, or Southeast Asia). The geographical region of a device cannot be changed once it is created, but if an identical geographical region is specified on update, the request will succeed.
     */
    readonly location: pulumi.Input<string>;
    /**
     * The description of the Data Box Edge/Gateway device model.
     */
    readonly modelDescription?: pulumi.Input<string>;
    /**
     * The resource group name.
     */
    readonly resourceGroupName: pulumi.Input<string>;
    /**
     * The SKU type.
     */
    readonly sku?: pulumi.Input<inputs.databoxedge.v20190301.Sku>;
    /**
     * The list of tags that describe the device. These tags can be used to view and group this device (across resource groups).
     */
    readonly tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
}