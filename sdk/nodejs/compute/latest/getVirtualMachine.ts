// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../../types/input";
import * as outputs from "../../types/output";
import * as utilities from "../../utilities";

export function getVirtualMachine(args: GetVirtualMachineArgs, opts?: pulumi.InvokeOptions): Promise<GetVirtualMachineResult> {
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    return pulumi.runtime.invoke("azure-nextgen:compute/latest:getVirtualMachine", {
        "expand": args.expand,
        "resourceGroupName": args.resourceGroupName,
        "vmName": args.vmName,
    }, opts);
}

export interface GetVirtualMachineArgs {
    /**
     * The expand expression to apply on the operation.
     */
    readonly expand?: string;
    /**
     * The name of the resource group.
     */
    readonly resourceGroupName: string;
    /**
     * The name of the virtual machine.
     */
    readonly vmName: string;
}

/**
 * Describes a Virtual Machine.
 */
export interface GetVirtualMachineResult {
    /**
     * Specifies additional capabilities enabled or disabled on the virtual machine.
     */
    readonly additionalCapabilities?: outputs.compute.latest.AdditionalCapabilitiesResponse;
    /**
     * Specifies information about the availability set that the virtual machine should be assigned to. Virtual machines specified in the same availability set are allocated to different nodes to maximize availability. For more information about availability sets, see [Manage the availability of virtual machines](https://docs.microsoft.com/azure/virtual-machines/virtual-machines-windows-manage-availability?toc=%2fazure%2fvirtual-machines%2fwindows%2ftoc.json). <br><br> For more information on Azure planned maintenance, see [Planned maintenance for virtual machines in Azure](https://docs.microsoft.com/azure/virtual-machines/virtual-machines-windows-planned-maintenance?toc=%2fazure%2fvirtual-machines%2fwindows%2ftoc.json) <br><br> Currently, a VM can only be added to availability set at creation time. The availability set to which the VM is being added should be under the same resource group as the availability set resource. An existing VM cannot be added to an availability set. <br><br>This property cannot exist along with a non-null properties.virtualMachineScaleSet reference.
     */
    readonly availabilitySet?: outputs.compute.latest.SubResourceResponse;
    /**
     * Specifies the billing related details of a Azure Spot virtual machine. <br><br>Minimum api-version: 2019-03-01.
     */
    readonly billingProfile?: outputs.compute.latest.BillingProfileResponse;
    /**
     * Specifies the boot diagnostic settings state. <br><br>Minimum api-version: 2015-06-15.
     */
    readonly diagnosticsProfile?: outputs.compute.latest.DiagnosticsProfileResponse;
    /**
     * Specifies the eviction policy for the Azure Spot virtual machine and Azure Spot scale set. <br><br>For Azure Spot virtual machines, both 'Deallocate' and 'Delete' are supported and the minimum api-version is 2019-03-01. <br><br>For Azure Spot scale sets, both 'Deallocate' and 'Delete' are supported and the minimum api-version is 2017-10-30-preview.
     */
    readonly evictionPolicy?: string;
    /**
     * Specifies the time alloted for all extensions to start. The time duration should be between 15 minutes and 120 minutes (inclusive) and should be specified in ISO 8601 format. The default value is 90 minutes (PT1H30M). <br><br> Minimum api-version: 2020-06-01
     */
    readonly extensionsTimeBudget?: string;
    /**
     * Specifies the hardware settings for the virtual machine.
     */
    readonly hardwareProfile?: outputs.compute.latest.HardwareProfileResponse;
    /**
     * Specifies information about the dedicated host that the virtual machine resides in. <br><br>Minimum api-version: 2018-10-01.
     */
    readonly host?: outputs.compute.latest.SubResourceResponse;
    /**
     * Specifies information about the dedicated host group that the virtual machine resides in. <br><br>Minimum api-version: 2020-06-01. <br><br>NOTE: User cannot specify both host and hostGroup properties.
     */
    readonly hostGroup?: outputs.compute.latest.SubResourceResponse;
    /**
     * The identity of the virtual machine, if configured.
     */
    readonly identity?: outputs.compute.latest.VirtualMachineIdentityResponse;
    /**
     * The virtual machine instance view.
     */
    readonly instanceView: outputs.compute.latest.VirtualMachineInstanceViewResponse;
    /**
     * Specifies that the image or disk that is being used was licensed on-premises. This element is only used for images that contain the Windows Server operating system. <br><br> Possible values are: <br><br> Windows_Client <br><br> Windows_Server <br><br> If this element is included in a request for an update, the value must match the initial value. This value cannot be updated. <br><br> For more information, see [Azure Hybrid Use Benefit for Windows Server](https://docs.microsoft.com/azure/virtual-machines/virtual-machines-windows-hybrid-use-benefit-licensing?toc=%2fazure%2fvirtual-machines%2fwindows%2ftoc.json) <br><br> Minimum api-version: 2015-06-15
     */
    readonly licenseType?: string;
    /**
     * Resource location
     */
    readonly location: string;
    /**
     * Resource name
     */
    readonly name: string;
    /**
     * Specifies the network interfaces of the virtual machine.
     */
    readonly networkProfile?: outputs.compute.latest.NetworkProfileResponse;
    /**
     * Specifies the operating system settings used while creating the virtual machine. Some of the settings cannot be changed once VM is provisioned.
     */
    readonly osProfile?: outputs.compute.latest.OSProfileResponse;
    /**
     * Specifies information about the marketplace image used to create the virtual machine. This element is only used for marketplace images. Before you can use a marketplace image from an API, you must enable the image for programmatic use.  In the Azure portal, find the marketplace image that you want to use and then click **Want to deploy programmatically, Get Started ->**. Enter any required information and then click **Save**.
     */
    readonly plan?: outputs.compute.latest.PlanResponse;
    /**
     * Specifies the priority for the virtual machine. <br><br>Minimum api-version: 2019-03-01
     */
    readonly priority?: string;
    /**
     * The provisioning state, which only appears in the response.
     */
    readonly provisioningState: string;
    /**
     * Specifies information about the proximity placement group that the virtual machine should be assigned to. <br><br>Minimum api-version: 2018-04-01.
     */
    readonly proximityPlacementGroup?: outputs.compute.latest.SubResourceResponse;
    /**
     * The virtual machine child extension resources.
     */
    readonly resources: outputs.compute.latest.VirtualMachineExtensionResponse[];
    /**
     * Specifies the Security related profile settings for the virtual machine.
     */
    readonly securityProfile?: outputs.compute.latest.SecurityProfileResponse;
    /**
     * Specifies the storage settings for the virtual machine disks.
     */
    readonly storageProfile?: outputs.compute.latest.StorageProfileResponse;
    /**
     * Resource tags
     */
    readonly tags?: {[key: string]: string};
    /**
     * Resource type
     */
    readonly type: string;
    /**
     * Specifies information about the virtual machine scale set that the virtual machine should be assigned to. Virtual machines specified in the same virtual machine scale set are allocated to different nodes to maximize availability. Currently, a VM can only be added to virtual machine scale set at creation time. An existing VM cannot be added to a virtual machine scale set. <br><br>This property cannot exist along with a non-null properties.availabilitySet reference. <br><br>Minimum api‐version: 2019‐03‐01
     */
    readonly virtualMachineScaleSet?: outputs.compute.latest.SubResourceResponse;
    /**
     * Specifies the VM unique ID which is a 128-bits identifier that is encoded and stored in all Azure IaaS VMs SMBIOS and can be read using platform BIOS commands.
     */
    readonly vmId: string;
    /**
     * The virtual machine zones.
     */
    readonly zones?: string[];
}