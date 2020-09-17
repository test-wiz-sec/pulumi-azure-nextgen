// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../../utilities";

/**
 * The extended info of the manager.
 */
export class ManagerExtendedInfo extends pulumi.CustomResource {
    /**
     * Get an existing ManagerExtendedInfo resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): ManagerExtendedInfo {
        return new ManagerExtendedInfo(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'azure-nextgen:storsimple/latest:ManagerExtendedInfo';

    /**
     * Returns true if the given object is an instance of ManagerExtendedInfo.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is ManagerExtendedInfo {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === ManagerExtendedInfo.__pulumiType;
    }

    /**
     * Represents the encryption algorithm used to encrypt the keys. None - if Key is saved in plain text format. Algorithm name - if key is encrypted
     */
    public readonly algorithm!: pulumi.Output<string>;
    /**
     * Represents the CEK of the resource.
     */
    public readonly encryptionKey!: pulumi.Output<string | undefined>;
    /**
     * Represents the Cert thumbprint that was used to encrypt the CEK.
     */
    public readonly encryptionKeyThumbprint!: pulumi.Output<string | undefined>;
    /**
     * The etag of the resource.
     */
    public readonly etag!: pulumi.Output<string | undefined>;
    /**
     * Represents the CIK of the resource.
     */
    public readonly integrityKey!: pulumi.Output<string>;
    /**
     * The Kind of the object. Currently only Series8000 is supported
     */
    public readonly kind!: pulumi.Output<string | undefined>;
    /**
     * The name of the object.
     */
    public /*out*/ readonly name!: pulumi.Output<string>;
    /**
     * Represents the portal thumbprint which can be used optionally to encrypt the entire data before storing it.
     */
    public readonly portalCertificateThumbprint!: pulumi.Output<string | undefined>;
    /**
     * The hierarchical type of the object.
     */
    public /*out*/ readonly type!: pulumi.Output<string>;
    /**
     * The version of the extended info being persisted.
     */
    public readonly version!: pulumi.Output<string | undefined>;

    /**
     * Create a ManagerExtendedInfo resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: ManagerExtendedInfoArgs, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (!(opts && opts.id)) {
            if (!args || args.algorithm === undefined) {
                throw new Error("Missing required property 'algorithm'");
            }
            if (!args || args.integrityKey === undefined) {
                throw new Error("Missing required property 'integrityKey'");
            }
            if (!args || args.managerName === undefined) {
                throw new Error("Missing required property 'managerName'");
            }
            if (!args || args.resourceGroupName === undefined) {
                throw new Error("Missing required property 'resourceGroupName'");
            }
            inputs["algorithm"] = args ? args.algorithm : undefined;
            inputs["encryptionKey"] = args ? args.encryptionKey : undefined;
            inputs["encryptionKeyThumbprint"] = args ? args.encryptionKeyThumbprint : undefined;
            inputs["etag"] = args ? args.etag : undefined;
            inputs["integrityKey"] = args ? args.integrityKey : undefined;
            inputs["kind"] = args ? args.kind : undefined;
            inputs["managerName"] = args ? args.managerName : undefined;
            inputs["portalCertificateThumbprint"] = args ? args.portalCertificateThumbprint : undefined;
            inputs["resourceGroupName"] = args ? args.resourceGroupName : undefined;
            inputs["version"] = args ? args.version : undefined;
            inputs["name"] = undefined /*out*/;
            inputs["type"] = undefined /*out*/;
        } else {
            inputs["algorithm"] = undefined /*out*/;
            inputs["encryptionKey"] = undefined /*out*/;
            inputs["encryptionKeyThumbprint"] = undefined /*out*/;
            inputs["etag"] = undefined /*out*/;
            inputs["integrityKey"] = undefined /*out*/;
            inputs["kind"] = undefined /*out*/;
            inputs["name"] = undefined /*out*/;
            inputs["portalCertificateThumbprint"] = undefined /*out*/;
            inputs["type"] = undefined /*out*/;
            inputs["version"] = undefined /*out*/;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        const aliasOpts = { aliases: [{ type: "azure-nextgen:storsimple/v20161001:ManagerExtendedInfo" }, { type: "azure-nextgen:storsimple/v20170601:ManagerExtendedInfo" }] };
        opts = opts ? pulumi.mergeOptions(opts, aliasOpts) : aliasOpts;
        super(ManagerExtendedInfo.__pulumiType, name, inputs, opts);
    }
}

/**
 * The set of arguments for constructing a ManagerExtendedInfo resource.
 */
export interface ManagerExtendedInfoArgs {
    /**
     * Represents the encryption algorithm used to encrypt the keys. None - if Key is saved in plain text format. Algorithm name - if key is encrypted
     */
    readonly algorithm: pulumi.Input<string>;
    /**
     * Represents the CEK of the resource.
     */
    readonly encryptionKey?: pulumi.Input<string>;
    /**
     * Represents the Cert thumbprint that was used to encrypt the CEK.
     */
    readonly encryptionKeyThumbprint?: pulumi.Input<string>;
    /**
     * The etag of the resource.
     */
    readonly etag?: pulumi.Input<string>;
    /**
     * Represents the CIK of the resource.
     */
    readonly integrityKey: pulumi.Input<string>;
    /**
     * The Kind of the object. Currently only Series8000 is supported
     */
    readonly kind?: pulumi.Input<string>;
    /**
     * The manager name
     */
    readonly managerName: pulumi.Input<string>;
    /**
     * Represents the portal thumbprint which can be used optionally to encrypt the entire data before storing it.
     */
    readonly portalCertificateThumbprint?: pulumi.Input<string>;
    /**
     * The resource group name
     */
    readonly resourceGroupName: pulumi.Input<string>;
    /**
     * The version of the extended info being persisted.
     */
    readonly version?: pulumi.Input<string>;
}