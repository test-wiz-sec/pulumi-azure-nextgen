// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../../types/input";
import * as outputs from "../../types/output";
import * as utilities from "../../utilities";

/**
 * Represents a connected cluster.
 */
export class ConnectedCluster extends pulumi.CustomResource {
    /**
     * Get an existing ConnectedCluster resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): ConnectedCluster {
        return new ConnectedCluster(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'azure-nextgen:kubernetes/v20200101preview:ConnectedCluster';

    /**
     * Returns true if the given object is an instance of ConnectedCluster.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is ConnectedCluster {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === ConnectedCluster.__pulumiType;
    }

    public readonly aadProfile!: pulumi.Output<outputs.kubernetes.v20200101preview.ConnectedClusterAADProfileResponse>;
    /**
     * Base64 encoded public certificate used by the agent to do the initial handshake to the backend services in Azure.
     */
    public readonly agentPublicKeyCertificate!: pulumi.Output<string>;
    /**
     * Version of the agent running on the connected cluster resource
     */
    public /*out*/ readonly agentVersion!: pulumi.Output<string>;
    /**
     * The identity of the connected cluster.
     */
    public readonly identity!: pulumi.Output<outputs.kubernetes.v20200101preview.ConnectedClusterIdentityResponse>;
    /**
     * The Kubernetes version of the connected cluster resource
     */
    public /*out*/ readonly kubernetesVersion!: pulumi.Output<string>;
    /**
     * The geo-location where the resource lives
     */
    public readonly location!: pulumi.Output<string>;
    /**
     * The name of the resource
     */
    public /*out*/ readonly name!: pulumi.Output<string>;
    /**
     * The current deployment state of connectedClusters.
     */
    public readonly provisioningState!: pulumi.Output<string | undefined>;
    /**
     * Resource tags.
     */
    public readonly tags!: pulumi.Output<{[key: string]: string} | undefined>;
    /**
     * Number of nodes present in the connected cluster resource
     */
    public /*out*/ readonly totalNodeCount!: pulumi.Output<number>;
    /**
     * The type of the resource. Ex- Microsoft.Compute/virtualMachines or Microsoft.Storage/storageAccounts.
     */
    public /*out*/ readonly type!: pulumi.Output<string>;

    /**
     * Create a ConnectedCluster resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: ConnectedClusterArgs, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (!(opts && opts.id)) {
            if (!args || args.aadProfile === undefined) {
                throw new Error("Missing required property 'aadProfile'");
            }
            if (!args || args.agentPublicKeyCertificate === undefined) {
                throw new Error("Missing required property 'agentPublicKeyCertificate'");
            }
            if (!args || args.clusterName === undefined) {
                throw new Error("Missing required property 'clusterName'");
            }
            if (!args || args.identity === undefined) {
                throw new Error("Missing required property 'identity'");
            }
            if (!args || args.location === undefined) {
                throw new Error("Missing required property 'location'");
            }
            if (!args || args.resourceGroupName === undefined) {
                throw new Error("Missing required property 'resourceGroupName'");
            }
            inputs["aadProfile"] = args ? args.aadProfile : undefined;
            inputs["agentPublicKeyCertificate"] = args ? args.agentPublicKeyCertificate : undefined;
            inputs["clusterName"] = args ? args.clusterName : undefined;
            inputs["identity"] = args ? args.identity : undefined;
            inputs["location"] = args ? args.location : undefined;
            inputs["provisioningState"] = args ? args.provisioningState : undefined;
            inputs["resourceGroupName"] = args ? args.resourceGroupName : undefined;
            inputs["tags"] = args ? args.tags : undefined;
            inputs["agentVersion"] = undefined /*out*/;
            inputs["kubernetesVersion"] = undefined /*out*/;
            inputs["name"] = undefined /*out*/;
            inputs["totalNodeCount"] = undefined /*out*/;
            inputs["type"] = undefined /*out*/;
        } else {
            inputs["aadProfile"] = undefined /*out*/;
            inputs["agentPublicKeyCertificate"] = undefined /*out*/;
            inputs["agentVersion"] = undefined /*out*/;
            inputs["identity"] = undefined /*out*/;
            inputs["kubernetesVersion"] = undefined /*out*/;
            inputs["location"] = undefined /*out*/;
            inputs["name"] = undefined /*out*/;
            inputs["provisioningState"] = undefined /*out*/;
            inputs["tags"] = undefined /*out*/;
            inputs["totalNodeCount"] = undefined /*out*/;
            inputs["type"] = undefined /*out*/;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        super(ConnectedCluster.__pulumiType, name, inputs, opts);
    }
}

/**
 * The set of arguments for constructing a ConnectedCluster resource.
 */
export interface ConnectedClusterArgs {
    readonly aadProfile: pulumi.Input<inputs.kubernetes.v20200101preview.ConnectedClusterAADProfile>;
    /**
     * Base64 encoded public certificate used by the agent to do the initial handshake to the backend services in Azure.
     */
    readonly agentPublicKeyCertificate: pulumi.Input<string>;
    /**
     * The name of the Kubernetes cluster on which get is called.
     */
    readonly clusterName: pulumi.Input<string>;
    /**
     * The identity of the connected cluster.
     */
    readonly identity: pulumi.Input<inputs.kubernetes.v20200101preview.ConnectedClusterIdentity>;
    /**
     * The geo-location where the resource lives
     */
    readonly location: pulumi.Input<string>;
    /**
     * The current deployment state of connectedClusters.
     */
    readonly provisioningState?: pulumi.Input<string>;
    /**
     * The name of the resource group. The name is case insensitive.
     */
    readonly resourceGroupName: pulumi.Input<string>;
    /**
     * Resource tags.
     */
    readonly tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
}