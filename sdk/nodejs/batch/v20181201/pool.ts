// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../../types/input";
import * as outputs from "../../types/output";
import * as utilities from "../../utilities";

/**
 * Contains information about a pool.
 */
export class Pool extends pulumi.CustomResource {
    /**
     * Get an existing Pool resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): Pool {
        return new Pool(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'azure-nextgen:batch/v20181201:Pool';

    /**
     * Returns true if the given object is an instance of Pool.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Pool {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Pool.__pulumiType;
    }

    public /*out*/ readonly allocationState!: pulumi.Output<string>;
    public /*out*/ readonly allocationStateTransitionTime!: pulumi.Output<string>;
    /**
     * The list of application licenses must be a subset of available Batch service application licenses. If a license is requested which is not supported, pool creation will fail.
     */
    public readonly applicationLicenses!: pulumi.Output<string[] | undefined>;
    /**
     * Changes to application packages affect all new compute nodes joining the pool, but do not affect compute nodes that are already in the pool until they are rebooted or reimaged.
     */
    public readonly applicationPackages!: pulumi.Output<outputs.batch.v20181201.ApplicationPackageReferenceResponse[] | undefined>;
    /**
     * This property is set only if the pool automatically scales, i.e. autoScaleSettings are used.
     */
    public /*out*/ readonly autoScaleRun!: pulumi.Output<outputs.batch.v20181201.AutoScaleRunResponse>;
    /**
     * For Windows compute nodes, the Batch service installs the certificates to the specified certificate store and location. For Linux compute nodes, the certificates are stored in a directory inside the task working directory and an environment variable AZ_BATCH_CERTIFICATES_DIR is supplied to the task to query for this location. For certificates with visibility of 'remoteUser', a 'certs' directory is created in the user's home directory (e.g., /home/{user-name}/certs) and certificates are placed in that directory.
     */
    public readonly certificates!: pulumi.Output<outputs.batch.v20181201.CertificateReferenceResponse[] | undefined>;
    public /*out*/ readonly creationTime!: pulumi.Output<string>;
    public /*out*/ readonly currentDedicatedNodes!: pulumi.Output<number>;
    public /*out*/ readonly currentLowPriorityNodes!: pulumi.Output<number>;
    /**
     * Using CloudServiceConfiguration specifies that the nodes should be creating using Azure Cloud Services (PaaS), while VirtualMachineConfiguration uses Azure Virtual Machines (IaaS).
     */
    public readonly deploymentConfiguration!: pulumi.Output<outputs.batch.v20181201.DeploymentConfigurationResponse | undefined>;
    /**
     * The display name need not be unique and can contain any Unicode characters up to a maximum length of 1024.
     */
    public readonly displayName!: pulumi.Output<string | undefined>;
    /**
     * The ETag of the resource, used for concurrency statements.
     */
    public /*out*/ readonly etag!: pulumi.Output<string>;
    /**
     * This imposes restrictions on which nodes can be assigned to the pool. Enabling this value can reduce the chance of the requested number of nodes to be allocated in the pool. If not specified, this value defaults to 'Disabled'.
     */
    public readonly interNodeCommunication!: pulumi.Output<string | undefined>;
    /**
     * This is the last time at which the pool level data, such as the targetDedicatedNodes or autoScaleSettings, changed. It does not factor in node-level changes such as a compute node changing state.
     */
    public /*out*/ readonly lastModified!: pulumi.Output<string>;
    public readonly maxTasksPerNode!: pulumi.Output<number | undefined>;
    /**
     * The Batch service does not assign any meaning to metadata; it is solely for the use of user code.
     */
    public readonly metadata!: pulumi.Output<outputs.batch.v20181201.MetadataItemResponse[] | undefined>;
    /**
     * The name of the resource.
     */
    public /*out*/ readonly name!: pulumi.Output<string>;
    /**
     * The network configuration for a pool.
     */
    public readonly networkConfiguration!: pulumi.Output<outputs.batch.v20181201.NetworkConfigurationResponse | undefined>;
    public /*out*/ readonly provisioningState!: pulumi.Output<string>;
    public /*out*/ readonly provisioningStateTransitionTime!: pulumi.Output<string>;
    /**
     * Describes either the current operation (if the pool AllocationState is Resizing) or the previously completed operation (if the AllocationState is Steady).
     */
    public /*out*/ readonly resizeOperationStatus!: pulumi.Output<outputs.batch.v20181201.ResizeOperationStatusResponse>;
    /**
     * Defines the desired size of the pool. This can either be 'fixedScale' where the requested targetDedicatedNodes is specified, or 'autoScale' which defines a formula which is periodically reevaluated. If this property is not specified, the pool will have a fixed scale with 0 targetDedicatedNodes.
     */
    public readonly scaleSettings!: pulumi.Output<outputs.batch.v20181201.ScaleSettingsResponse | undefined>;
    /**
     * In an PATCH (update) operation, this property can be set to an empty object to remove the start task from the pool.
     */
    public readonly startTask!: pulumi.Output<outputs.batch.v20181201.StartTaskResponse | undefined>;
    public readonly taskSchedulingPolicy!: pulumi.Output<outputs.batch.v20181201.TaskSchedulingPolicyResponse | undefined>;
    /**
     * The type of the resource.
     */
    public /*out*/ readonly type!: pulumi.Output<string>;
    public readonly userAccounts!: pulumi.Output<outputs.batch.v20181201.UserAccountResponse[] | undefined>;
    /**
     * For information about available sizes of virtual machines for Cloud Services pools (pools created with cloudServiceConfiguration), see Sizes for Cloud Services (https://azure.microsoft.com/documentation/articles/cloud-services-sizes-specs/). Batch supports all Cloud Services VM sizes except ExtraSmall. For information about available VM sizes for pools using images from the Virtual Machines Marketplace (pools created with virtualMachineConfiguration) see Sizes for Virtual Machines (Linux) (https://azure.microsoft.com/documentation/articles/virtual-machines-linux-sizes/) or Sizes for Virtual Machines (Windows) (https://azure.microsoft.com/documentation/articles/virtual-machines-windows-sizes/). Batch supports all Azure VM sizes except STANDARD_A0 and those with premium storage (STANDARD_GS, STANDARD_DS, and STANDARD_DSV2 series).
     */
    public readonly vmSize!: pulumi.Output<string | undefined>;

    /**
     * Create a Pool resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: PoolArgs, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (!(opts && opts.id)) {
            if (!args || args.accountName === undefined) {
                throw new Error("Missing required property 'accountName'");
            }
            if (!args || args.poolName === undefined) {
                throw new Error("Missing required property 'poolName'");
            }
            if (!args || args.resourceGroupName === undefined) {
                throw new Error("Missing required property 'resourceGroupName'");
            }
            inputs["accountName"] = args ? args.accountName : undefined;
            inputs["applicationLicenses"] = args ? args.applicationLicenses : undefined;
            inputs["applicationPackages"] = args ? args.applicationPackages : undefined;
            inputs["certificates"] = args ? args.certificates : undefined;
            inputs["deploymentConfiguration"] = args ? args.deploymentConfiguration : undefined;
            inputs["displayName"] = args ? args.displayName : undefined;
            inputs["interNodeCommunication"] = args ? args.interNodeCommunication : undefined;
            inputs["maxTasksPerNode"] = args ? args.maxTasksPerNode : undefined;
            inputs["metadata"] = args ? args.metadata : undefined;
            inputs["networkConfiguration"] = args ? args.networkConfiguration : undefined;
            inputs["poolName"] = args ? args.poolName : undefined;
            inputs["resourceGroupName"] = args ? args.resourceGroupName : undefined;
            inputs["scaleSettings"] = args ? args.scaleSettings : undefined;
            inputs["startTask"] = args ? args.startTask : undefined;
            inputs["taskSchedulingPolicy"] = args ? args.taskSchedulingPolicy : undefined;
            inputs["userAccounts"] = args ? args.userAccounts : undefined;
            inputs["vmSize"] = args ? args.vmSize : undefined;
            inputs["allocationState"] = undefined /*out*/;
            inputs["allocationStateTransitionTime"] = undefined /*out*/;
            inputs["autoScaleRun"] = undefined /*out*/;
            inputs["creationTime"] = undefined /*out*/;
            inputs["currentDedicatedNodes"] = undefined /*out*/;
            inputs["currentLowPriorityNodes"] = undefined /*out*/;
            inputs["etag"] = undefined /*out*/;
            inputs["lastModified"] = undefined /*out*/;
            inputs["name"] = undefined /*out*/;
            inputs["provisioningState"] = undefined /*out*/;
            inputs["provisioningStateTransitionTime"] = undefined /*out*/;
            inputs["resizeOperationStatus"] = undefined /*out*/;
            inputs["type"] = undefined /*out*/;
        } else {
            inputs["allocationState"] = undefined /*out*/;
            inputs["allocationStateTransitionTime"] = undefined /*out*/;
            inputs["applicationLicenses"] = undefined /*out*/;
            inputs["applicationPackages"] = undefined /*out*/;
            inputs["autoScaleRun"] = undefined /*out*/;
            inputs["certificates"] = undefined /*out*/;
            inputs["creationTime"] = undefined /*out*/;
            inputs["currentDedicatedNodes"] = undefined /*out*/;
            inputs["currentLowPriorityNodes"] = undefined /*out*/;
            inputs["deploymentConfiguration"] = undefined /*out*/;
            inputs["displayName"] = undefined /*out*/;
            inputs["etag"] = undefined /*out*/;
            inputs["interNodeCommunication"] = undefined /*out*/;
            inputs["lastModified"] = undefined /*out*/;
            inputs["maxTasksPerNode"] = undefined /*out*/;
            inputs["metadata"] = undefined /*out*/;
            inputs["name"] = undefined /*out*/;
            inputs["networkConfiguration"] = undefined /*out*/;
            inputs["provisioningState"] = undefined /*out*/;
            inputs["provisioningStateTransitionTime"] = undefined /*out*/;
            inputs["resizeOperationStatus"] = undefined /*out*/;
            inputs["scaleSettings"] = undefined /*out*/;
            inputs["startTask"] = undefined /*out*/;
            inputs["taskSchedulingPolicy"] = undefined /*out*/;
            inputs["type"] = undefined /*out*/;
            inputs["userAccounts"] = undefined /*out*/;
            inputs["vmSize"] = undefined /*out*/;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        const aliasOpts = { aliases: [{ type: "azure-nextgen:batch/latest:Pool" }, { type: "azure-nextgen:batch/v20170901:Pool" }, { type: "azure-nextgen:batch/v20190401:Pool" }, { type: "azure-nextgen:batch/v20190801:Pool" }, { type: "azure-nextgen:batch/v20200301:Pool" }, { type: "azure-nextgen:batch/v20200501:Pool" }, { type: "azure-nextgen:batch/v20200901:Pool" }] };
        opts = opts ? pulumi.mergeOptions(opts, aliasOpts) : aliasOpts;
        super(Pool.__pulumiType, name, inputs, opts);
    }
}

/**
 * The set of arguments for constructing a Pool resource.
 */
export interface PoolArgs {
    /**
     * The name of the Batch account.
     */
    readonly accountName: pulumi.Input<string>;
    /**
     * The list of application licenses must be a subset of available Batch service application licenses. If a license is requested which is not supported, pool creation will fail.
     */
    readonly applicationLicenses?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Changes to application packages affect all new compute nodes joining the pool, but do not affect compute nodes that are already in the pool until they are rebooted or reimaged.
     */
    readonly applicationPackages?: pulumi.Input<pulumi.Input<inputs.batch.v20181201.ApplicationPackageReference>[]>;
    /**
     * For Windows compute nodes, the Batch service installs the certificates to the specified certificate store and location. For Linux compute nodes, the certificates are stored in a directory inside the task working directory and an environment variable AZ_BATCH_CERTIFICATES_DIR is supplied to the task to query for this location. For certificates with visibility of 'remoteUser', a 'certs' directory is created in the user's home directory (e.g., /home/{user-name}/certs) and certificates are placed in that directory.
     */
    readonly certificates?: pulumi.Input<pulumi.Input<inputs.batch.v20181201.CertificateReference>[]>;
    /**
     * Using CloudServiceConfiguration specifies that the nodes should be creating using Azure Cloud Services (PaaS), while VirtualMachineConfiguration uses Azure Virtual Machines (IaaS).
     */
    readonly deploymentConfiguration?: pulumi.Input<inputs.batch.v20181201.DeploymentConfiguration>;
    /**
     * The display name need not be unique and can contain any Unicode characters up to a maximum length of 1024.
     */
    readonly displayName?: pulumi.Input<string>;
    /**
     * This imposes restrictions on which nodes can be assigned to the pool. Enabling this value can reduce the chance of the requested number of nodes to be allocated in the pool. If not specified, this value defaults to 'Disabled'.
     */
    readonly interNodeCommunication?: pulumi.Input<string>;
    readonly maxTasksPerNode?: pulumi.Input<number>;
    /**
     * The Batch service does not assign any meaning to metadata; it is solely for the use of user code.
     */
    readonly metadata?: pulumi.Input<pulumi.Input<inputs.batch.v20181201.MetadataItem>[]>;
    /**
     * The network configuration for a pool.
     */
    readonly networkConfiguration?: pulumi.Input<inputs.batch.v20181201.NetworkConfiguration>;
    /**
     * The pool name. This must be unique within the account.
     */
    readonly poolName: pulumi.Input<string>;
    /**
     * The name of the resource group that contains the Batch account.
     */
    readonly resourceGroupName: pulumi.Input<string>;
    /**
     * Defines the desired size of the pool. This can either be 'fixedScale' where the requested targetDedicatedNodes is specified, or 'autoScale' which defines a formula which is periodically reevaluated. If this property is not specified, the pool will have a fixed scale with 0 targetDedicatedNodes.
     */
    readonly scaleSettings?: pulumi.Input<inputs.batch.v20181201.ScaleSettings>;
    /**
     * In an PATCH (update) operation, this property can be set to an empty object to remove the start task from the pool.
     */
    readonly startTask?: pulumi.Input<inputs.batch.v20181201.StartTask>;
    readonly taskSchedulingPolicy?: pulumi.Input<inputs.batch.v20181201.TaskSchedulingPolicy>;
    readonly userAccounts?: pulumi.Input<pulumi.Input<inputs.batch.v20181201.UserAccount>[]>;
    /**
     * For information about available sizes of virtual machines for Cloud Services pools (pools created with cloudServiceConfiguration), see Sizes for Cloud Services (https://azure.microsoft.com/documentation/articles/cloud-services-sizes-specs/). Batch supports all Cloud Services VM sizes except ExtraSmall. For information about available VM sizes for pools using images from the Virtual Machines Marketplace (pools created with virtualMachineConfiguration) see Sizes for Virtual Machines (Linux) (https://azure.microsoft.com/documentation/articles/virtual-machines-linux-sizes/) or Sizes for Virtual Machines (Windows) (https://azure.microsoft.com/documentation/articles/virtual-machines-windows-sizes/). Batch supports all Azure VM sizes except STANDARD_A0 and those with premium storage (STANDARD_GS, STANDARD_DS, and STANDARD_DSV2 series).
     */
    readonly vmSize?: pulumi.Input<string>;
}
