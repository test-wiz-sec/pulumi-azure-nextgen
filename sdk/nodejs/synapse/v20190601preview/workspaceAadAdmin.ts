// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../../utilities";

/**
 * Workspace active directory administrator
 */
export class WorkspaceAadAdmin extends pulumi.CustomResource {
    /**
     * Get an existing WorkspaceAadAdmin resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): WorkspaceAadAdmin {
        return new WorkspaceAadAdmin(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'azure-nextgen:synapse/v20190601preview:WorkspaceAadAdmin';

    /**
     * Returns true if the given object is an instance of WorkspaceAadAdmin.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is WorkspaceAadAdmin {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === WorkspaceAadAdmin.__pulumiType;
    }

    /**
     * Workspace active directory administrator type
     */
    public readonly administratorType!: pulumi.Output<string | undefined>;
    /**
     * Login of the workspace active directory administrator
     */
    public readonly login!: pulumi.Output<string | undefined>;
    /**
     * The name of the resource
     */
    public /*out*/ readonly name!: pulumi.Output<string>;
    /**
     * Object ID of the workspace active directory administrator
     */
    public readonly sid!: pulumi.Output<string | undefined>;
    /**
     * Tenant ID of the workspace active directory administrator
     */
    public readonly tenantId!: pulumi.Output<string | undefined>;
    /**
     * The type of the resource. Ex- Microsoft.Compute/virtualMachines or Microsoft.Storage/storageAccounts.
     */
    public /*out*/ readonly type!: pulumi.Output<string>;

    /**
     * Create a WorkspaceAadAdmin resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: WorkspaceAadAdminArgs, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (!(opts && opts.id)) {
            if (!args || args.resourceGroupName === undefined) {
                throw new Error("Missing required property 'resourceGroupName'");
            }
            if (!args || args.workspaceName === undefined) {
                throw new Error("Missing required property 'workspaceName'");
            }
            inputs["administratorType"] = args ? args.administratorType : undefined;
            inputs["login"] = args ? args.login : undefined;
            inputs["resourceGroupName"] = args ? args.resourceGroupName : undefined;
            inputs["sid"] = args ? args.sid : undefined;
            inputs["tenantId"] = args ? args.tenantId : undefined;
            inputs["workspaceName"] = args ? args.workspaceName : undefined;
            inputs["name"] = undefined /*out*/;
            inputs["type"] = undefined /*out*/;
        } else {
            inputs["administratorType"] = undefined /*out*/;
            inputs["login"] = undefined /*out*/;
            inputs["name"] = undefined /*out*/;
            inputs["sid"] = undefined /*out*/;
            inputs["tenantId"] = undefined /*out*/;
            inputs["type"] = undefined /*out*/;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        super(WorkspaceAadAdmin.__pulumiType, name, inputs, opts);
    }
}

/**
 * The set of arguments for constructing a WorkspaceAadAdmin resource.
 */
export interface WorkspaceAadAdminArgs {
    /**
     * Workspace active directory administrator type
     */
    readonly administratorType?: pulumi.Input<string>;
    /**
     * Login of the workspace active directory administrator
     */
    readonly login?: pulumi.Input<string>;
    /**
     * The name of the resource group. The name is case insensitive.
     */
    readonly resourceGroupName: pulumi.Input<string>;
    /**
     * Object ID of the workspace active directory administrator
     */
    readonly sid?: pulumi.Input<string>;
    /**
     * Tenant ID of the workspace active directory administrator
     */
    readonly tenantId?: pulumi.Input<string>;
    /**
     * The name of the workspace
     */
    readonly workspaceName: pulumi.Input<string>;
}