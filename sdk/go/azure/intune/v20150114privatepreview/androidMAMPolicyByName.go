// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v20150114privatepreview

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Android Policy entity for Intune MAM.
type AndroidMAMPolicyByName struct {
	pulumi.CustomResourceState

	AccessRecheckOfflineTimeout pulumi.StringPtrOutput `pulumi:"accessRecheckOfflineTimeout"`
	AccessRecheckOnlineTimeout  pulumi.StringPtrOutput `pulumi:"accessRecheckOnlineTimeout"`
	AppSharingFromLevel         pulumi.StringPtrOutput `pulumi:"appSharingFromLevel"`
	AppSharingToLevel           pulumi.StringPtrOutput `pulumi:"appSharingToLevel"`
	Authentication              pulumi.StringPtrOutput `pulumi:"authentication"`
	ClipboardSharingLevel       pulumi.StringPtrOutput `pulumi:"clipboardSharingLevel"`
	DataBackup                  pulumi.StringPtrOutput `pulumi:"dataBackup"`
	Description                 pulumi.StringPtrOutput `pulumi:"description"`
	DeviceCompliance            pulumi.StringPtrOutput `pulumi:"deviceCompliance"`
	FileEncryption              pulumi.StringPtrOutput `pulumi:"fileEncryption"`
	FileSharingSaveAs           pulumi.StringPtrOutput `pulumi:"fileSharingSaveAs"`
	FriendlyName                pulumi.StringOutput    `pulumi:"friendlyName"`
	GroupStatus                 pulumi.StringOutput    `pulumi:"groupStatus"`
	LastModifiedTime            pulumi.StringOutput    `pulumi:"lastModifiedTime"`
	// Resource Location
	Location       pulumi.StringPtrOutput `pulumi:"location"`
	ManagedBrowser pulumi.StringPtrOutput `pulumi:"managedBrowser"`
	// Resource name
	Name               pulumi.StringOutput    `pulumi:"name"`
	NumOfApps          pulumi.IntOutput       `pulumi:"numOfApps"`
	OfflineWipeTimeout pulumi.StringPtrOutput `pulumi:"offlineWipeTimeout"`
	Pin                pulumi.StringPtrOutput `pulumi:"pin"`
	PinNumRetry        pulumi.IntPtrOutput    `pulumi:"pinNumRetry"`
	ScreenCapture      pulumi.StringPtrOutput `pulumi:"screenCapture"`
	// Resource Tags
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// Resource type
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewAndroidMAMPolicyByName registers a new resource with the given unique name, arguments, and options.
func NewAndroidMAMPolicyByName(ctx *pulumi.Context,
	name string, args *AndroidMAMPolicyByNameArgs, opts ...pulumi.ResourceOption) (*AndroidMAMPolicyByName, error) {
	if args == nil || args.FriendlyName == nil {
		return nil, errors.New("missing required argument 'FriendlyName'")
	}
	if args == nil || args.HostName == nil {
		return nil, errors.New("missing required argument 'HostName'")
	}
	if args == nil || args.PolicyName == nil {
		return nil, errors.New("missing required argument 'PolicyName'")
	}
	if args == nil {
		args = &AndroidMAMPolicyByNameArgs{}
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-nextgen:intune/v20150114preview:AndroidMAMPolicyByName"),
		},
	})
	opts = append(opts, aliases)
	var resource AndroidMAMPolicyByName
	err := ctx.RegisterResource("azure-nextgen:intune/v20150114privatepreview:AndroidMAMPolicyByName", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetAndroidMAMPolicyByName gets an existing AndroidMAMPolicyByName resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetAndroidMAMPolicyByName(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *AndroidMAMPolicyByNameState, opts ...pulumi.ResourceOption) (*AndroidMAMPolicyByName, error) {
	var resource AndroidMAMPolicyByName
	err := ctx.ReadResource("azure-nextgen:intune/v20150114privatepreview:AndroidMAMPolicyByName", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering AndroidMAMPolicyByName resources.
type androidMAMPolicyByNameState struct {
	AccessRecheckOfflineTimeout *string `pulumi:"accessRecheckOfflineTimeout"`
	AccessRecheckOnlineTimeout  *string `pulumi:"accessRecheckOnlineTimeout"`
	AppSharingFromLevel         *string `pulumi:"appSharingFromLevel"`
	AppSharingToLevel           *string `pulumi:"appSharingToLevel"`
	Authentication              *string `pulumi:"authentication"`
	ClipboardSharingLevel       *string `pulumi:"clipboardSharingLevel"`
	DataBackup                  *string `pulumi:"dataBackup"`
	Description                 *string `pulumi:"description"`
	DeviceCompliance            *string `pulumi:"deviceCompliance"`
	FileEncryption              *string `pulumi:"fileEncryption"`
	FileSharingSaveAs           *string `pulumi:"fileSharingSaveAs"`
	FriendlyName                *string `pulumi:"friendlyName"`
	GroupStatus                 *string `pulumi:"groupStatus"`
	LastModifiedTime            *string `pulumi:"lastModifiedTime"`
	// Resource Location
	Location       *string `pulumi:"location"`
	ManagedBrowser *string `pulumi:"managedBrowser"`
	// Resource name
	Name               *string `pulumi:"name"`
	NumOfApps          *int    `pulumi:"numOfApps"`
	OfflineWipeTimeout *string `pulumi:"offlineWipeTimeout"`
	Pin                *string `pulumi:"pin"`
	PinNumRetry        *int    `pulumi:"pinNumRetry"`
	ScreenCapture      *string `pulumi:"screenCapture"`
	// Resource Tags
	Tags map[string]string `pulumi:"tags"`
	// Resource type
	Type *string `pulumi:"type"`
}

type AndroidMAMPolicyByNameState struct {
	AccessRecheckOfflineTimeout pulumi.StringPtrInput
	AccessRecheckOnlineTimeout  pulumi.StringPtrInput
	AppSharingFromLevel         pulumi.StringPtrInput
	AppSharingToLevel           pulumi.StringPtrInput
	Authentication              pulumi.StringPtrInput
	ClipboardSharingLevel       pulumi.StringPtrInput
	DataBackup                  pulumi.StringPtrInput
	Description                 pulumi.StringPtrInput
	DeviceCompliance            pulumi.StringPtrInput
	FileEncryption              pulumi.StringPtrInput
	FileSharingSaveAs           pulumi.StringPtrInput
	FriendlyName                pulumi.StringPtrInput
	GroupStatus                 pulumi.StringPtrInput
	LastModifiedTime            pulumi.StringPtrInput
	// Resource Location
	Location       pulumi.StringPtrInput
	ManagedBrowser pulumi.StringPtrInput
	// Resource name
	Name               pulumi.StringPtrInput
	NumOfApps          pulumi.IntPtrInput
	OfflineWipeTimeout pulumi.StringPtrInput
	Pin                pulumi.StringPtrInput
	PinNumRetry        pulumi.IntPtrInput
	ScreenCapture      pulumi.StringPtrInput
	// Resource Tags
	Tags pulumi.StringMapInput
	// Resource type
	Type pulumi.StringPtrInput
}

func (AndroidMAMPolicyByNameState) ElementType() reflect.Type {
	return reflect.TypeOf((*androidMAMPolicyByNameState)(nil)).Elem()
}

type androidMAMPolicyByNameArgs struct {
	AccessRecheckOfflineTimeout *string `pulumi:"accessRecheckOfflineTimeout"`
	AccessRecheckOnlineTimeout  *string `pulumi:"accessRecheckOnlineTimeout"`
	AppSharingFromLevel         *string `pulumi:"appSharingFromLevel"`
	AppSharingToLevel           *string `pulumi:"appSharingToLevel"`
	Authentication              *string `pulumi:"authentication"`
	ClipboardSharingLevel       *string `pulumi:"clipboardSharingLevel"`
	DataBackup                  *string `pulumi:"dataBackup"`
	Description                 *string `pulumi:"description"`
	DeviceCompliance            *string `pulumi:"deviceCompliance"`
	FileEncryption              *string `pulumi:"fileEncryption"`
	FileSharingSaveAs           *string `pulumi:"fileSharingSaveAs"`
	FriendlyName                string  `pulumi:"friendlyName"`
	// Location hostName for the tenant
	HostName string `pulumi:"hostName"`
	// Resource Location
	Location           *string `pulumi:"location"`
	ManagedBrowser     *string `pulumi:"managedBrowser"`
	OfflineWipeTimeout *string `pulumi:"offlineWipeTimeout"`
	Pin                *string `pulumi:"pin"`
	PinNumRetry        *int    `pulumi:"pinNumRetry"`
	// Unique name for the policy
	PolicyName    string  `pulumi:"policyName"`
	ScreenCapture *string `pulumi:"screenCapture"`
	// Resource Tags
	Tags map[string]string `pulumi:"tags"`
}

// The set of arguments for constructing a AndroidMAMPolicyByName resource.
type AndroidMAMPolicyByNameArgs struct {
	AccessRecheckOfflineTimeout pulumi.StringPtrInput
	AccessRecheckOnlineTimeout  pulumi.StringPtrInput
	AppSharingFromLevel         pulumi.StringPtrInput
	AppSharingToLevel           pulumi.StringPtrInput
	Authentication              pulumi.StringPtrInput
	ClipboardSharingLevel       pulumi.StringPtrInput
	DataBackup                  pulumi.StringPtrInput
	Description                 pulumi.StringPtrInput
	DeviceCompliance            pulumi.StringPtrInput
	FileEncryption              pulumi.StringPtrInput
	FileSharingSaveAs           pulumi.StringPtrInput
	FriendlyName                pulumi.StringInput
	// Location hostName for the tenant
	HostName pulumi.StringInput
	// Resource Location
	Location           pulumi.StringPtrInput
	ManagedBrowser     pulumi.StringPtrInput
	OfflineWipeTimeout pulumi.StringPtrInput
	Pin                pulumi.StringPtrInput
	PinNumRetry        pulumi.IntPtrInput
	// Unique name for the policy
	PolicyName    pulumi.StringInput
	ScreenCapture pulumi.StringPtrInput
	// Resource Tags
	Tags pulumi.StringMapInput
}

func (AndroidMAMPolicyByNameArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*androidMAMPolicyByNameArgs)(nil)).Elem()
}
