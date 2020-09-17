// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v20161001

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// The File Share.
type FileShare struct {
	pulumi.CustomResourceState

	// The user/group who will have full permission in this share. Active directory email address. Example: xyz@contoso.com or Contoso\xyz.
	AdminUser pulumi.StringOutput `pulumi:"adminUser"`
	// The data policy
	DataPolicy pulumi.StringOutput `pulumi:"dataPolicy"`
	// Description for file share
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// The local used capacity in Bytes.
	LocalUsedCapacityInBytes pulumi.IntOutput `pulumi:"localUsedCapacityInBytes"`
	// The monitoring status
	MonitoringStatus pulumi.StringOutput `pulumi:"monitoringStatus"`
	// The name.
	Name pulumi.StringOutput `pulumi:"name"`
	// The total provisioned capacity in Bytes
	ProvisionedCapacityInBytes pulumi.IntOutput `pulumi:"provisionedCapacityInBytes"`
	// The Share Status
	ShareStatus pulumi.StringOutput `pulumi:"shareStatus"`
	// The type.
	Type pulumi.StringOutput `pulumi:"type"`
	// The used capacity in Bytes.
	UsedCapacityInBytes pulumi.IntOutput `pulumi:"usedCapacityInBytes"`
}

// NewFileShare registers a new resource with the given unique name, arguments, and options.
func NewFileShare(ctx *pulumi.Context,
	name string, args *FileShareArgs, opts ...pulumi.ResourceOption) (*FileShare, error) {
	if args == nil || args.AdminUser == nil {
		return nil, errors.New("missing required argument 'AdminUser'")
	}
	if args == nil || args.DataPolicy == nil {
		return nil, errors.New("missing required argument 'DataPolicy'")
	}
	if args == nil || args.DeviceName == nil {
		return nil, errors.New("missing required argument 'DeviceName'")
	}
	if args == nil || args.FileServerName == nil {
		return nil, errors.New("missing required argument 'FileServerName'")
	}
	if args == nil || args.ManagerName == nil {
		return nil, errors.New("missing required argument 'ManagerName'")
	}
	if args == nil || args.MonitoringStatus == nil {
		return nil, errors.New("missing required argument 'MonitoringStatus'")
	}
	if args == nil || args.ProvisionedCapacityInBytes == nil {
		return nil, errors.New("missing required argument 'ProvisionedCapacityInBytes'")
	}
	if args == nil || args.ResourceGroupName == nil {
		return nil, errors.New("missing required argument 'ResourceGroupName'")
	}
	if args == nil || args.ShareName == nil {
		return nil, errors.New("missing required argument 'ShareName'")
	}
	if args == nil || args.ShareStatus == nil {
		return nil, errors.New("missing required argument 'ShareStatus'")
	}
	if args == nil {
		args = &FileShareArgs{}
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-nextgen:storsimple/latest:FileShare"),
		},
	})
	opts = append(opts, aliases)
	var resource FileShare
	err := ctx.RegisterResource("azure-nextgen:storsimple/v20161001:FileShare", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetFileShare gets an existing FileShare resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetFileShare(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *FileShareState, opts ...pulumi.ResourceOption) (*FileShare, error) {
	var resource FileShare
	err := ctx.ReadResource("azure-nextgen:storsimple/v20161001:FileShare", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering FileShare resources.
type fileShareState struct {
	// The user/group who will have full permission in this share. Active directory email address. Example: xyz@contoso.com or Contoso\xyz.
	AdminUser *string `pulumi:"adminUser"`
	// The data policy
	DataPolicy *string `pulumi:"dataPolicy"`
	// Description for file share
	Description *string `pulumi:"description"`
	// The local used capacity in Bytes.
	LocalUsedCapacityInBytes *int `pulumi:"localUsedCapacityInBytes"`
	// The monitoring status
	MonitoringStatus *string `pulumi:"monitoringStatus"`
	// The name.
	Name *string `pulumi:"name"`
	// The total provisioned capacity in Bytes
	ProvisionedCapacityInBytes *int `pulumi:"provisionedCapacityInBytes"`
	// The Share Status
	ShareStatus *string `pulumi:"shareStatus"`
	// The type.
	Type *string `pulumi:"type"`
	// The used capacity in Bytes.
	UsedCapacityInBytes *int `pulumi:"usedCapacityInBytes"`
}

type FileShareState struct {
	// The user/group who will have full permission in this share. Active directory email address. Example: xyz@contoso.com or Contoso\xyz.
	AdminUser pulumi.StringPtrInput
	// The data policy
	DataPolicy pulumi.StringPtrInput
	// Description for file share
	Description pulumi.StringPtrInput
	// The local used capacity in Bytes.
	LocalUsedCapacityInBytes pulumi.IntPtrInput
	// The monitoring status
	MonitoringStatus pulumi.StringPtrInput
	// The name.
	Name pulumi.StringPtrInput
	// The total provisioned capacity in Bytes
	ProvisionedCapacityInBytes pulumi.IntPtrInput
	// The Share Status
	ShareStatus pulumi.StringPtrInput
	// The type.
	Type pulumi.StringPtrInput
	// The used capacity in Bytes.
	UsedCapacityInBytes pulumi.IntPtrInput
}

func (FileShareState) ElementType() reflect.Type {
	return reflect.TypeOf((*fileShareState)(nil)).Elem()
}

type fileShareArgs struct {
	// The user/group who will have full permission in this share. Active directory email address. Example: xyz@contoso.com or Contoso\xyz.
	AdminUser string `pulumi:"adminUser"`
	// The data policy
	DataPolicy string `pulumi:"dataPolicy"`
	// Description for file share
	Description *string `pulumi:"description"`
	// The device name.
	DeviceName string `pulumi:"deviceName"`
	// The file server name.
	FileServerName string `pulumi:"fileServerName"`
	// The manager name
	ManagerName string `pulumi:"managerName"`
	// The monitoring status
	MonitoringStatus string `pulumi:"monitoringStatus"`
	// The total provisioned capacity in Bytes
	ProvisionedCapacityInBytes int `pulumi:"provisionedCapacityInBytes"`
	// The resource group name
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The file share name.
	ShareName string `pulumi:"shareName"`
	// The Share Status
	ShareStatus string `pulumi:"shareStatus"`
}

// The set of arguments for constructing a FileShare resource.
type FileShareArgs struct {
	// The user/group who will have full permission in this share. Active directory email address. Example: xyz@contoso.com or Contoso\xyz.
	AdminUser pulumi.StringInput
	// The data policy
	DataPolicy pulumi.StringInput
	// Description for file share
	Description pulumi.StringPtrInput
	// The device name.
	DeviceName pulumi.StringInput
	// The file server name.
	FileServerName pulumi.StringInput
	// The manager name
	ManagerName pulumi.StringInput
	// The monitoring status
	MonitoringStatus pulumi.StringInput
	// The total provisioned capacity in Bytes
	ProvisionedCapacityInBytes pulumi.IntInput
	// The resource group name
	ResourceGroupName pulumi.StringInput
	// The file share name.
	ShareName pulumi.StringInput
	// The Share Status
	ShareStatus pulumi.StringInput
}

func (FileShareArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*fileShareArgs)(nil)).Elem()
}