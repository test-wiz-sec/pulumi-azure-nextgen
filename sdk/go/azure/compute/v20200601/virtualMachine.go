// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v20200601

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Describes a Virtual Machine.
type VirtualMachine struct {
	pulumi.CustomResourceState

	// Specifies additional capabilities enabled or disabled on the virtual machine.
	AdditionalCapabilities AdditionalCapabilitiesResponsePtrOutput `pulumi:"additionalCapabilities"`
	// Specifies information about the availability set that the virtual machine should be assigned to. Virtual machines specified in the same availability set are allocated to different nodes to maximize availability. For more information about availability sets, see [Manage the availability of virtual machines](https://docs.microsoft.com/azure/virtual-machines/virtual-machines-windows-manage-availability?toc=%2fazure%2fvirtual-machines%2fwindows%2ftoc.json). <br><br> For more information on Azure planned maintenance, see [Planned maintenance for virtual machines in Azure](https://docs.microsoft.com/azure/virtual-machines/virtual-machines-windows-planned-maintenance?toc=%2fazure%2fvirtual-machines%2fwindows%2ftoc.json) <br><br> Currently, a VM can only be added to availability set at creation time. The availability set to which the VM is being added should be under the same resource group as the availability set resource. An existing VM cannot be added to an availability set. <br><br>This property cannot exist along with a non-null properties.virtualMachineScaleSet reference.
	AvailabilitySet SubResourceResponsePtrOutput `pulumi:"availabilitySet"`
	// Specifies the billing related details of a Azure Spot virtual machine. <br><br>Minimum api-version: 2019-03-01.
	BillingProfile BillingProfileResponsePtrOutput `pulumi:"billingProfile"`
	// Specifies the boot diagnostic settings state. <br><br>Minimum api-version: 2015-06-15.
	DiagnosticsProfile DiagnosticsProfileResponsePtrOutput `pulumi:"diagnosticsProfile"`
	// Specifies the eviction policy for the Azure Spot virtual machine and Azure Spot scale set. <br><br>For Azure Spot virtual machines, both 'Deallocate' and 'Delete' are supported and the minimum api-version is 2019-03-01. <br><br>For Azure Spot scale sets, both 'Deallocate' and 'Delete' are supported and the minimum api-version is 2017-10-30-preview.
	EvictionPolicy pulumi.StringPtrOutput `pulumi:"evictionPolicy"`
	// Specifies the time alloted for all extensions to start. The time duration should be between 15 minutes and 120 minutes (inclusive) and should be specified in ISO 8601 format. The default value is 90 minutes (PT1H30M). <br><br> Minimum api-version: 2020-06-01
	ExtensionsTimeBudget pulumi.StringPtrOutput `pulumi:"extensionsTimeBudget"`
	// Specifies the hardware settings for the virtual machine.
	HardwareProfile HardwareProfileResponsePtrOutput `pulumi:"hardwareProfile"`
	// Specifies information about the dedicated host that the virtual machine resides in. <br><br>Minimum api-version: 2018-10-01.
	Host SubResourceResponsePtrOutput `pulumi:"host"`
	// Specifies information about the dedicated host group that the virtual machine resides in. <br><br>Minimum api-version: 2020-06-01. <br><br>NOTE: User cannot specify both host and hostGroup properties.
	HostGroup SubResourceResponsePtrOutput `pulumi:"hostGroup"`
	// The identity of the virtual machine, if configured.
	Identity VirtualMachineIdentityResponsePtrOutput `pulumi:"identity"`
	// The virtual machine instance view.
	InstanceView VirtualMachineInstanceViewResponseOutput `pulumi:"instanceView"`
	// Specifies that the image or disk that is being used was licensed on-premises. <br><br> Possible values for Windows Server operating system are: <br><br> Windows_Client <br><br> Windows_Server <br><br> Possible values for Linux Server operating system are: <br><br> RHEL_BYOS (for RHEL) <br><br> SLES_BYOS (for SUSE) <br><br> For more information, see [Azure Hybrid Use Benefit for Windows Server](https://docs.microsoft.com/azure/virtual-machines/windows/hybrid-use-benefit-licensing) <br><br> [Azure Hybrid Use Benefit for Linux Server](https://docs.microsoft.com/azure/virtual-machines/linux/azure-hybrid-benefit-linux) <br><br> Minimum api-version: 2015-06-15
	LicenseType pulumi.StringPtrOutput `pulumi:"licenseType"`
	// Resource location
	Location pulumi.StringOutput `pulumi:"location"`
	// Resource name
	Name pulumi.StringOutput `pulumi:"name"`
	// Specifies the network interfaces of the virtual machine.
	NetworkProfile NetworkProfileResponsePtrOutput `pulumi:"networkProfile"`
	// Specifies the operating system settings used while creating the virtual machine. Some of the settings cannot be changed once VM is provisioned.
	OsProfile OSProfileResponsePtrOutput `pulumi:"osProfile"`
	// Specifies information about the marketplace image used to create the virtual machine. This element is only used for marketplace images. Before you can use a marketplace image from an API, you must enable the image for programmatic use.  In the Azure portal, find the marketplace image that you want to use and then click **Want to deploy programmatically, Get Started ->**. Enter any required information and then click **Save**.
	Plan PlanResponsePtrOutput `pulumi:"plan"`
	// Specifies the priority for the virtual machine. <br><br>Minimum api-version: 2019-03-01
	Priority pulumi.StringPtrOutput `pulumi:"priority"`
	// The provisioning state, which only appears in the response.
	ProvisioningState pulumi.StringOutput `pulumi:"provisioningState"`
	// Specifies information about the proximity placement group that the virtual machine should be assigned to. <br><br>Minimum api-version: 2018-04-01.
	ProximityPlacementGroup SubResourceResponsePtrOutput `pulumi:"proximityPlacementGroup"`
	// The virtual machine child extension resources.
	Resources VirtualMachineExtensionResponseArrayOutput `pulumi:"resources"`
	// Specifies the Security related profile settings for the virtual machine.
	SecurityProfile SecurityProfileResponsePtrOutput `pulumi:"securityProfile"`
	// Specifies the storage settings for the virtual machine disks.
	StorageProfile StorageProfileResponsePtrOutput `pulumi:"storageProfile"`
	// Resource tags
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// Resource type
	Type pulumi.StringOutput `pulumi:"type"`
	// Specifies information about the virtual machine scale set that the virtual machine should be assigned to. Virtual machines specified in the same virtual machine scale set are allocated to different nodes to maximize availability. Currently, a VM can only be added to virtual machine scale set at creation time. An existing VM cannot be added to a virtual machine scale set. <br><br>This property cannot exist along with a non-null properties.availabilitySet reference. <br><br>Minimum api‐version: 2019‐03‐01
	VirtualMachineScaleSet SubResourceResponsePtrOutput `pulumi:"virtualMachineScaleSet"`
	// Specifies the VM unique ID which is a 128-bits identifier that is encoded and stored in all Azure IaaS VMs SMBIOS and can be read using platform BIOS commands.
	VmId pulumi.StringOutput `pulumi:"vmId"`
	// The virtual machine zones.
	Zones pulumi.StringArrayOutput `pulumi:"zones"`
}

// NewVirtualMachine registers a new resource with the given unique name, arguments, and options.
func NewVirtualMachine(ctx *pulumi.Context,
	name string, args *VirtualMachineArgs, opts ...pulumi.ResourceOption) (*VirtualMachine, error) {
	if args == nil || args.Location == nil {
		return nil, errors.New("missing required argument 'Location'")
	}
	if args == nil || args.ResourceGroupName == nil {
		return nil, errors.New("missing required argument 'ResourceGroupName'")
	}
	if args == nil || args.VmName == nil {
		return nil, errors.New("missing required argument 'VmName'")
	}
	if args == nil {
		args = &VirtualMachineArgs{}
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-nextgen:compute/latest:VirtualMachine"),
		},
		{
			Type: pulumi.String("azure-nextgen:compute/v20150615:VirtualMachine"),
		},
		{
			Type: pulumi.String("azure-nextgen:compute/v20160330:VirtualMachine"),
		},
		{
			Type: pulumi.String("azure-nextgen:compute/v20160430preview:VirtualMachine"),
		},
		{
			Type: pulumi.String("azure-nextgen:compute/v20170330:VirtualMachine"),
		},
		{
			Type: pulumi.String("azure-nextgen:compute/v20171201:VirtualMachine"),
		},
		{
			Type: pulumi.String("azure-nextgen:compute/v20180401:VirtualMachine"),
		},
		{
			Type: pulumi.String("azure-nextgen:compute/v20180601:VirtualMachine"),
		},
		{
			Type: pulumi.String("azure-nextgen:compute/v20181001:VirtualMachine"),
		},
		{
			Type: pulumi.String("azure-nextgen:compute/v20190301:VirtualMachine"),
		},
		{
			Type: pulumi.String("azure-nextgen:compute/v20190701:VirtualMachine"),
		},
		{
			Type: pulumi.String("azure-nextgen:compute/v20191201:VirtualMachine"),
		},
	})
	opts = append(opts, aliases)
	var resource VirtualMachine
	err := ctx.RegisterResource("azure-nextgen:compute/v20200601:VirtualMachine", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetVirtualMachine gets an existing VirtualMachine resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetVirtualMachine(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *VirtualMachineState, opts ...pulumi.ResourceOption) (*VirtualMachine, error) {
	var resource VirtualMachine
	err := ctx.ReadResource("azure-nextgen:compute/v20200601:VirtualMachine", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering VirtualMachine resources.
type virtualMachineState struct {
	// Specifies additional capabilities enabled or disabled on the virtual machine.
	AdditionalCapabilities *AdditionalCapabilitiesResponse `pulumi:"additionalCapabilities"`
	// Specifies information about the availability set that the virtual machine should be assigned to. Virtual machines specified in the same availability set are allocated to different nodes to maximize availability. For more information about availability sets, see [Manage the availability of virtual machines](https://docs.microsoft.com/azure/virtual-machines/virtual-machines-windows-manage-availability?toc=%2fazure%2fvirtual-machines%2fwindows%2ftoc.json). <br><br> For more information on Azure planned maintenance, see [Planned maintenance for virtual machines in Azure](https://docs.microsoft.com/azure/virtual-machines/virtual-machines-windows-planned-maintenance?toc=%2fazure%2fvirtual-machines%2fwindows%2ftoc.json) <br><br> Currently, a VM can only be added to availability set at creation time. The availability set to which the VM is being added should be under the same resource group as the availability set resource. An existing VM cannot be added to an availability set. <br><br>This property cannot exist along with a non-null properties.virtualMachineScaleSet reference.
	AvailabilitySet *SubResourceResponse `pulumi:"availabilitySet"`
	// Specifies the billing related details of a Azure Spot virtual machine. <br><br>Minimum api-version: 2019-03-01.
	BillingProfile *BillingProfileResponse `pulumi:"billingProfile"`
	// Specifies the boot diagnostic settings state. <br><br>Minimum api-version: 2015-06-15.
	DiagnosticsProfile *DiagnosticsProfileResponse `pulumi:"diagnosticsProfile"`
	// Specifies the eviction policy for the Azure Spot virtual machine and Azure Spot scale set. <br><br>For Azure Spot virtual machines, both 'Deallocate' and 'Delete' are supported and the minimum api-version is 2019-03-01. <br><br>For Azure Spot scale sets, both 'Deallocate' and 'Delete' are supported and the minimum api-version is 2017-10-30-preview.
	EvictionPolicy *string `pulumi:"evictionPolicy"`
	// Specifies the time alloted for all extensions to start. The time duration should be between 15 minutes and 120 minutes (inclusive) and should be specified in ISO 8601 format. The default value is 90 minutes (PT1H30M). <br><br> Minimum api-version: 2020-06-01
	ExtensionsTimeBudget *string `pulumi:"extensionsTimeBudget"`
	// Specifies the hardware settings for the virtual machine.
	HardwareProfile *HardwareProfileResponse `pulumi:"hardwareProfile"`
	// Specifies information about the dedicated host that the virtual machine resides in. <br><br>Minimum api-version: 2018-10-01.
	Host *SubResourceResponse `pulumi:"host"`
	// Specifies information about the dedicated host group that the virtual machine resides in. <br><br>Minimum api-version: 2020-06-01. <br><br>NOTE: User cannot specify both host and hostGroup properties.
	HostGroup *SubResourceResponse `pulumi:"hostGroup"`
	// The identity of the virtual machine, if configured.
	Identity *VirtualMachineIdentityResponse `pulumi:"identity"`
	// The virtual machine instance view.
	InstanceView *VirtualMachineInstanceViewResponse `pulumi:"instanceView"`
	// Specifies that the image or disk that is being used was licensed on-premises. <br><br> Possible values for Windows Server operating system are: <br><br> Windows_Client <br><br> Windows_Server <br><br> Possible values for Linux Server operating system are: <br><br> RHEL_BYOS (for RHEL) <br><br> SLES_BYOS (for SUSE) <br><br> For more information, see [Azure Hybrid Use Benefit for Windows Server](https://docs.microsoft.com/azure/virtual-machines/windows/hybrid-use-benefit-licensing) <br><br> [Azure Hybrid Use Benefit for Linux Server](https://docs.microsoft.com/azure/virtual-machines/linux/azure-hybrid-benefit-linux) <br><br> Minimum api-version: 2015-06-15
	LicenseType *string `pulumi:"licenseType"`
	// Resource location
	Location *string `pulumi:"location"`
	// Resource name
	Name *string `pulumi:"name"`
	// Specifies the network interfaces of the virtual machine.
	NetworkProfile *NetworkProfileResponse `pulumi:"networkProfile"`
	// Specifies the operating system settings used while creating the virtual machine. Some of the settings cannot be changed once VM is provisioned.
	OsProfile *OSProfileResponse `pulumi:"osProfile"`
	// Specifies information about the marketplace image used to create the virtual machine. This element is only used for marketplace images. Before you can use a marketplace image from an API, you must enable the image for programmatic use.  In the Azure portal, find the marketplace image that you want to use and then click **Want to deploy programmatically, Get Started ->**. Enter any required information and then click **Save**.
	Plan *PlanResponse `pulumi:"plan"`
	// Specifies the priority for the virtual machine. <br><br>Minimum api-version: 2019-03-01
	Priority *string `pulumi:"priority"`
	// The provisioning state, which only appears in the response.
	ProvisioningState *string `pulumi:"provisioningState"`
	// Specifies information about the proximity placement group that the virtual machine should be assigned to. <br><br>Minimum api-version: 2018-04-01.
	ProximityPlacementGroup *SubResourceResponse `pulumi:"proximityPlacementGroup"`
	// The virtual machine child extension resources.
	Resources []VirtualMachineExtensionResponse `pulumi:"resources"`
	// Specifies the Security related profile settings for the virtual machine.
	SecurityProfile *SecurityProfileResponse `pulumi:"securityProfile"`
	// Specifies the storage settings for the virtual machine disks.
	StorageProfile *StorageProfileResponse `pulumi:"storageProfile"`
	// Resource tags
	Tags map[string]string `pulumi:"tags"`
	// Resource type
	Type *string `pulumi:"type"`
	// Specifies information about the virtual machine scale set that the virtual machine should be assigned to. Virtual machines specified in the same virtual machine scale set are allocated to different nodes to maximize availability. Currently, a VM can only be added to virtual machine scale set at creation time. An existing VM cannot be added to a virtual machine scale set. <br><br>This property cannot exist along with a non-null properties.availabilitySet reference. <br><br>Minimum api‐version: 2019‐03‐01
	VirtualMachineScaleSet *SubResourceResponse `pulumi:"virtualMachineScaleSet"`
	// Specifies the VM unique ID which is a 128-bits identifier that is encoded and stored in all Azure IaaS VMs SMBIOS and can be read using platform BIOS commands.
	VmId *string `pulumi:"vmId"`
	// The virtual machine zones.
	Zones []string `pulumi:"zones"`
}

type VirtualMachineState struct {
	// Specifies additional capabilities enabled or disabled on the virtual machine.
	AdditionalCapabilities AdditionalCapabilitiesResponsePtrInput
	// Specifies information about the availability set that the virtual machine should be assigned to. Virtual machines specified in the same availability set are allocated to different nodes to maximize availability. For more information about availability sets, see [Manage the availability of virtual machines](https://docs.microsoft.com/azure/virtual-machines/virtual-machines-windows-manage-availability?toc=%2fazure%2fvirtual-machines%2fwindows%2ftoc.json). <br><br> For more information on Azure planned maintenance, see [Planned maintenance for virtual machines in Azure](https://docs.microsoft.com/azure/virtual-machines/virtual-machines-windows-planned-maintenance?toc=%2fazure%2fvirtual-machines%2fwindows%2ftoc.json) <br><br> Currently, a VM can only be added to availability set at creation time. The availability set to which the VM is being added should be under the same resource group as the availability set resource. An existing VM cannot be added to an availability set. <br><br>This property cannot exist along with a non-null properties.virtualMachineScaleSet reference.
	AvailabilitySet SubResourceResponsePtrInput
	// Specifies the billing related details of a Azure Spot virtual machine. <br><br>Minimum api-version: 2019-03-01.
	BillingProfile BillingProfileResponsePtrInput
	// Specifies the boot diagnostic settings state. <br><br>Minimum api-version: 2015-06-15.
	DiagnosticsProfile DiagnosticsProfileResponsePtrInput
	// Specifies the eviction policy for the Azure Spot virtual machine and Azure Spot scale set. <br><br>For Azure Spot virtual machines, both 'Deallocate' and 'Delete' are supported and the minimum api-version is 2019-03-01. <br><br>For Azure Spot scale sets, both 'Deallocate' and 'Delete' are supported and the minimum api-version is 2017-10-30-preview.
	EvictionPolicy pulumi.StringPtrInput
	// Specifies the time alloted for all extensions to start. The time duration should be between 15 minutes and 120 minutes (inclusive) and should be specified in ISO 8601 format. The default value is 90 minutes (PT1H30M). <br><br> Minimum api-version: 2020-06-01
	ExtensionsTimeBudget pulumi.StringPtrInput
	// Specifies the hardware settings for the virtual machine.
	HardwareProfile HardwareProfileResponsePtrInput
	// Specifies information about the dedicated host that the virtual machine resides in. <br><br>Minimum api-version: 2018-10-01.
	Host SubResourceResponsePtrInput
	// Specifies information about the dedicated host group that the virtual machine resides in. <br><br>Minimum api-version: 2020-06-01. <br><br>NOTE: User cannot specify both host and hostGroup properties.
	HostGroup SubResourceResponsePtrInput
	// The identity of the virtual machine, if configured.
	Identity VirtualMachineIdentityResponsePtrInput
	// The virtual machine instance view.
	InstanceView VirtualMachineInstanceViewResponsePtrInput
	// Specifies that the image or disk that is being used was licensed on-premises. <br><br> Possible values for Windows Server operating system are: <br><br> Windows_Client <br><br> Windows_Server <br><br> Possible values for Linux Server operating system are: <br><br> RHEL_BYOS (for RHEL) <br><br> SLES_BYOS (for SUSE) <br><br> For more information, see [Azure Hybrid Use Benefit for Windows Server](https://docs.microsoft.com/azure/virtual-machines/windows/hybrid-use-benefit-licensing) <br><br> [Azure Hybrid Use Benefit for Linux Server](https://docs.microsoft.com/azure/virtual-machines/linux/azure-hybrid-benefit-linux) <br><br> Minimum api-version: 2015-06-15
	LicenseType pulumi.StringPtrInput
	// Resource location
	Location pulumi.StringPtrInput
	// Resource name
	Name pulumi.StringPtrInput
	// Specifies the network interfaces of the virtual machine.
	NetworkProfile NetworkProfileResponsePtrInput
	// Specifies the operating system settings used while creating the virtual machine. Some of the settings cannot be changed once VM is provisioned.
	OsProfile OSProfileResponsePtrInput
	// Specifies information about the marketplace image used to create the virtual machine. This element is only used for marketplace images. Before you can use a marketplace image from an API, you must enable the image for programmatic use.  In the Azure portal, find the marketplace image that you want to use and then click **Want to deploy programmatically, Get Started ->**. Enter any required information and then click **Save**.
	Plan PlanResponsePtrInput
	// Specifies the priority for the virtual machine. <br><br>Minimum api-version: 2019-03-01
	Priority pulumi.StringPtrInput
	// The provisioning state, which only appears in the response.
	ProvisioningState pulumi.StringPtrInput
	// Specifies information about the proximity placement group that the virtual machine should be assigned to. <br><br>Minimum api-version: 2018-04-01.
	ProximityPlacementGroup SubResourceResponsePtrInput
	// The virtual machine child extension resources.
	Resources VirtualMachineExtensionResponseArrayInput
	// Specifies the Security related profile settings for the virtual machine.
	SecurityProfile SecurityProfileResponsePtrInput
	// Specifies the storage settings for the virtual machine disks.
	StorageProfile StorageProfileResponsePtrInput
	// Resource tags
	Tags pulumi.StringMapInput
	// Resource type
	Type pulumi.StringPtrInput
	// Specifies information about the virtual machine scale set that the virtual machine should be assigned to. Virtual machines specified in the same virtual machine scale set are allocated to different nodes to maximize availability. Currently, a VM can only be added to virtual machine scale set at creation time. An existing VM cannot be added to a virtual machine scale set. <br><br>This property cannot exist along with a non-null properties.availabilitySet reference. <br><br>Minimum api‐version: 2019‐03‐01
	VirtualMachineScaleSet SubResourceResponsePtrInput
	// Specifies the VM unique ID which is a 128-bits identifier that is encoded and stored in all Azure IaaS VMs SMBIOS and can be read using platform BIOS commands.
	VmId pulumi.StringPtrInput
	// The virtual machine zones.
	Zones pulumi.StringArrayInput
}

func (VirtualMachineState) ElementType() reflect.Type {
	return reflect.TypeOf((*virtualMachineState)(nil)).Elem()
}

type virtualMachineArgs struct {
	// Specifies additional capabilities enabled or disabled on the virtual machine.
	AdditionalCapabilities *AdditionalCapabilities `pulumi:"additionalCapabilities"`
	// Specifies information about the availability set that the virtual machine should be assigned to. Virtual machines specified in the same availability set are allocated to different nodes to maximize availability. For more information about availability sets, see [Manage the availability of virtual machines](https://docs.microsoft.com/azure/virtual-machines/virtual-machines-windows-manage-availability?toc=%2fazure%2fvirtual-machines%2fwindows%2ftoc.json). <br><br> For more information on Azure planned maintenance, see [Planned maintenance for virtual machines in Azure](https://docs.microsoft.com/azure/virtual-machines/virtual-machines-windows-planned-maintenance?toc=%2fazure%2fvirtual-machines%2fwindows%2ftoc.json) <br><br> Currently, a VM can only be added to availability set at creation time. The availability set to which the VM is being added should be under the same resource group as the availability set resource. An existing VM cannot be added to an availability set. <br><br>This property cannot exist along with a non-null properties.virtualMachineScaleSet reference.
	AvailabilitySet *SubResource `pulumi:"availabilitySet"`
	// Specifies the billing related details of a Azure Spot virtual machine. <br><br>Minimum api-version: 2019-03-01.
	BillingProfile *BillingProfile `pulumi:"billingProfile"`
	// Specifies the boot diagnostic settings state. <br><br>Minimum api-version: 2015-06-15.
	DiagnosticsProfile *DiagnosticsProfile `pulumi:"diagnosticsProfile"`
	// Specifies the eviction policy for the Azure Spot virtual machine and Azure Spot scale set. <br><br>For Azure Spot virtual machines, both 'Deallocate' and 'Delete' are supported and the minimum api-version is 2019-03-01. <br><br>For Azure Spot scale sets, both 'Deallocate' and 'Delete' are supported and the minimum api-version is 2017-10-30-preview.
	EvictionPolicy *string `pulumi:"evictionPolicy"`
	// Specifies the time alloted for all extensions to start. The time duration should be between 15 minutes and 120 minutes (inclusive) and should be specified in ISO 8601 format. The default value is 90 minutes (PT1H30M). <br><br> Minimum api-version: 2020-06-01
	ExtensionsTimeBudget *string `pulumi:"extensionsTimeBudget"`
	// Specifies the hardware settings for the virtual machine.
	HardwareProfile *HardwareProfile `pulumi:"hardwareProfile"`
	// Specifies information about the dedicated host that the virtual machine resides in. <br><br>Minimum api-version: 2018-10-01.
	Host *SubResource `pulumi:"host"`
	// Specifies information about the dedicated host group that the virtual machine resides in. <br><br>Minimum api-version: 2020-06-01. <br><br>NOTE: User cannot specify both host and hostGroup properties.
	HostGroup *SubResource `pulumi:"hostGroup"`
	// The identity of the virtual machine, if configured.
	Identity *VirtualMachineIdentity `pulumi:"identity"`
	// Specifies that the image or disk that is being used was licensed on-premises. <br><br> Possible values for Windows Server operating system are: <br><br> Windows_Client <br><br> Windows_Server <br><br> Possible values for Linux Server operating system are: <br><br> RHEL_BYOS (for RHEL) <br><br> SLES_BYOS (for SUSE) <br><br> For more information, see [Azure Hybrid Use Benefit for Windows Server](https://docs.microsoft.com/azure/virtual-machines/windows/hybrid-use-benefit-licensing) <br><br> [Azure Hybrid Use Benefit for Linux Server](https://docs.microsoft.com/azure/virtual-machines/linux/azure-hybrid-benefit-linux) <br><br> Minimum api-version: 2015-06-15
	LicenseType *string `pulumi:"licenseType"`
	// Resource location
	Location string `pulumi:"location"`
	// Specifies the network interfaces of the virtual machine.
	NetworkProfile *NetworkProfile `pulumi:"networkProfile"`
	// Specifies the operating system settings used while creating the virtual machine. Some of the settings cannot be changed once VM is provisioned.
	OsProfile *OSProfile `pulumi:"osProfile"`
	// Specifies information about the marketplace image used to create the virtual machine. This element is only used for marketplace images. Before you can use a marketplace image from an API, you must enable the image for programmatic use.  In the Azure portal, find the marketplace image that you want to use and then click **Want to deploy programmatically, Get Started ->**. Enter any required information and then click **Save**.
	Plan *Plan `pulumi:"plan"`
	// Specifies the priority for the virtual machine. <br><br>Minimum api-version: 2019-03-01
	Priority *string `pulumi:"priority"`
	// Specifies information about the proximity placement group that the virtual machine should be assigned to. <br><br>Minimum api-version: 2018-04-01.
	ProximityPlacementGroup *SubResource `pulumi:"proximityPlacementGroup"`
	// The name of the resource group.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// Specifies the Security related profile settings for the virtual machine.
	SecurityProfile *SecurityProfile `pulumi:"securityProfile"`
	// Specifies the storage settings for the virtual machine disks.
	StorageProfile *StorageProfile `pulumi:"storageProfile"`
	// Resource tags
	Tags map[string]string `pulumi:"tags"`
	// Specifies information about the virtual machine scale set that the virtual machine should be assigned to. Virtual machines specified in the same virtual machine scale set are allocated to different nodes to maximize availability. Currently, a VM can only be added to virtual machine scale set at creation time. An existing VM cannot be added to a virtual machine scale set. <br><br>This property cannot exist along with a non-null properties.availabilitySet reference. <br><br>Minimum api‐version: 2019‐03‐01
	VirtualMachineScaleSet *SubResource `pulumi:"virtualMachineScaleSet"`
	// The name of the virtual machine.
	VmName string `pulumi:"vmName"`
	// The virtual machine zones.
	Zones []string `pulumi:"zones"`
}

// The set of arguments for constructing a VirtualMachine resource.
type VirtualMachineArgs struct {
	// Specifies additional capabilities enabled or disabled on the virtual machine.
	AdditionalCapabilities AdditionalCapabilitiesPtrInput
	// Specifies information about the availability set that the virtual machine should be assigned to. Virtual machines specified in the same availability set are allocated to different nodes to maximize availability. For more information about availability sets, see [Manage the availability of virtual machines](https://docs.microsoft.com/azure/virtual-machines/virtual-machines-windows-manage-availability?toc=%2fazure%2fvirtual-machines%2fwindows%2ftoc.json). <br><br> For more information on Azure planned maintenance, see [Planned maintenance for virtual machines in Azure](https://docs.microsoft.com/azure/virtual-machines/virtual-machines-windows-planned-maintenance?toc=%2fazure%2fvirtual-machines%2fwindows%2ftoc.json) <br><br> Currently, a VM can only be added to availability set at creation time. The availability set to which the VM is being added should be under the same resource group as the availability set resource. An existing VM cannot be added to an availability set. <br><br>This property cannot exist along with a non-null properties.virtualMachineScaleSet reference.
	AvailabilitySet SubResourcePtrInput
	// Specifies the billing related details of a Azure Spot virtual machine. <br><br>Minimum api-version: 2019-03-01.
	BillingProfile BillingProfilePtrInput
	// Specifies the boot diagnostic settings state. <br><br>Minimum api-version: 2015-06-15.
	DiagnosticsProfile DiagnosticsProfilePtrInput
	// Specifies the eviction policy for the Azure Spot virtual machine and Azure Spot scale set. <br><br>For Azure Spot virtual machines, both 'Deallocate' and 'Delete' are supported and the minimum api-version is 2019-03-01. <br><br>For Azure Spot scale sets, both 'Deallocate' and 'Delete' are supported and the minimum api-version is 2017-10-30-preview.
	EvictionPolicy pulumi.StringPtrInput
	// Specifies the time alloted for all extensions to start. The time duration should be between 15 minutes and 120 minutes (inclusive) and should be specified in ISO 8601 format. The default value is 90 minutes (PT1H30M). <br><br> Minimum api-version: 2020-06-01
	ExtensionsTimeBudget pulumi.StringPtrInput
	// Specifies the hardware settings for the virtual machine.
	HardwareProfile HardwareProfilePtrInput
	// Specifies information about the dedicated host that the virtual machine resides in. <br><br>Minimum api-version: 2018-10-01.
	Host SubResourcePtrInput
	// Specifies information about the dedicated host group that the virtual machine resides in. <br><br>Minimum api-version: 2020-06-01. <br><br>NOTE: User cannot specify both host and hostGroup properties.
	HostGroup SubResourcePtrInput
	// The identity of the virtual machine, if configured.
	Identity VirtualMachineIdentityPtrInput
	// Specifies that the image or disk that is being used was licensed on-premises. <br><br> Possible values for Windows Server operating system are: <br><br> Windows_Client <br><br> Windows_Server <br><br> Possible values for Linux Server operating system are: <br><br> RHEL_BYOS (for RHEL) <br><br> SLES_BYOS (for SUSE) <br><br> For more information, see [Azure Hybrid Use Benefit for Windows Server](https://docs.microsoft.com/azure/virtual-machines/windows/hybrid-use-benefit-licensing) <br><br> [Azure Hybrid Use Benefit for Linux Server](https://docs.microsoft.com/azure/virtual-machines/linux/azure-hybrid-benefit-linux) <br><br> Minimum api-version: 2015-06-15
	LicenseType pulumi.StringPtrInput
	// Resource location
	Location pulumi.StringInput
	// Specifies the network interfaces of the virtual machine.
	NetworkProfile NetworkProfilePtrInput
	// Specifies the operating system settings used while creating the virtual machine. Some of the settings cannot be changed once VM is provisioned.
	OsProfile OSProfilePtrInput
	// Specifies information about the marketplace image used to create the virtual machine. This element is only used for marketplace images. Before you can use a marketplace image from an API, you must enable the image for programmatic use.  In the Azure portal, find the marketplace image that you want to use and then click **Want to deploy programmatically, Get Started ->**. Enter any required information and then click **Save**.
	Plan PlanPtrInput
	// Specifies the priority for the virtual machine. <br><br>Minimum api-version: 2019-03-01
	Priority pulumi.StringPtrInput
	// Specifies information about the proximity placement group that the virtual machine should be assigned to. <br><br>Minimum api-version: 2018-04-01.
	ProximityPlacementGroup SubResourcePtrInput
	// The name of the resource group.
	ResourceGroupName pulumi.StringInput
	// Specifies the Security related profile settings for the virtual machine.
	SecurityProfile SecurityProfilePtrInput
	// Specifies the storage settings for the virtual machine disks.
	StorageProfile StorageProfilePtrInput
	// Resource tags
	Tags pulumi.StringMapInput
	// Specifies information about the virtual machine scale set that the virtual machine should be assigned to. Virtual machines specified in the same virtual machine scale set are allocated to different nodes to maximize availability. Currently, a VM can only be added to virtual machine scale set at creation time. An existing VM cannot be added to a virtual machine scale set. <br><br>This property cannot exist along with a non-null properties.availabilitySet reference. <br><br>Minimum api‐version: 2019‐03‐01
	VirtualMachineScaleSet SubResourcePtrInput
	// The name of the virtual machine.
	VmName pulumi.StringInput
	// The virtual machine zones.
	Zones pulumi.StringArrayInput
}

func (VirtualMachineArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*virtualMachineArgs)(nil)).Elem()
}
