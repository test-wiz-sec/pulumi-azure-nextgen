// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v20150504preview

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Describes a DNS RecordSet (a set of DNS records with the same name and type).
type RecordSet struct {
	pulumi.CustomResourceState

	// Gets or sets the ETag of the RecordSet.
	Etag pulumi.StringPtrOutput `pulumi:"etag"`
	// The geo-location where the resource lives
	Location pulumi.StringOutput `pulumi:"location"`
	// The name of the resource
	Name pulumi.StringOutput `pulumi:"name"`
	// Gets or sets the properties of the RecordSet.
	Properties RecordSetPropertiesResponseOutput `pulumi:"properties"`
	// Resource tags.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewRecordSet registers a new resource with the given unique name, arguments, and options.
func NewRecordSet(ctx *pulumi.Context,
	name string, args *RecordSetArgs, opts ...pulumi.ResourceOption) (*RecordSet, error) {
	if args == nil || args.Location == nil {
		return nil, errors.New("missing required argument 'Location'")
	}
	if args == nil || args.RecordType == nil {
		return nil, errors.New("missing required argument 'RecordType'")
	}
	if args == nil || args.RelativeRecordSetName == nil {
		return nil, errors.New("missing required argument 'RelativeRecordSetName'")
	}
	if args == nil || args.ResourceGroupName == nil {
		return nil, errors.New("missing required argument 'ResourceGroupName'")
	}
	if args == nil || args.ZoneName == nil {
		return nil, errors.New("missing required argument 'ZoneName'")
	}
	if args == nil {
		args = &RecordSetArgs{}
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-nextgen:network/v20160401:RecordSet"),
		},
		{
			Type: pulumi.String("azure-nextgen:network/v20170901:RecordSet"),
		},
		{
			Type: pulumi.String("azure-nextgen:network/v20171001:RecordSet"),
		},
		{
			Type: pulumi.String("azure-nextgen:network/v20180301preview:RecordSet"),
		},
		{
			Type: pulumi.String("azure-nextgen:network/v20180501:RecordSet"),
		},
	})
	opts = append(opts, aliases)
	var resource RecordSet
	err := ctx.RegisterResource("azure-nextgen:network/v20150504preview:RecordSet", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetRecordSet gets an existing RecordSet resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetRecordSet(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *RecordSetState, opts ...pulumi.ResourceOption) (*RecordSet, error) {
	var resource RecordSet
	err := ctx.ReadResource("azure-nextgen:network/v20150504preview:RecordSet", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering RecordSet resources.
type recordSetState struct {
	// Gets or sets the ETag of the RecordSet.
	Etag *string `pulumi:"etag"`
	// The geo-location where the resource lives
	Location *string `pulumi:"location"`
	// The name of the resource
	Name *string `pulumi:"name"`
	// Gets or sets the properties of the RecordSet.
	Properties *RecordSetPropertiesResponse `pulumi:"properties"`
	// Resource tags.
	Tags map[string]string `pulumi:"tags"`
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type *string `pulumi:"type"`
}

type RecordSetState struct {
	// Gets or sets the ETag of the RecordSet.
	Etag pulumi.StringPtrInput
	// The geo-location where the resource lives
	Location pulumi.StringPtrInput
	// The name of the resource
	Name pulumi.StringPtrInput
	// Gets or sets the properties of the RecordSet.
	Properties RecordSetPropertiesResponsePtrInput
	// Resource tags.
	Tags pulumi.StringMapInput
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type pulumi.StringPtrInput
}

func (RecordSetState) ElementType() reflect.Type {
	return reflect.TypeOf((*recordSetState)(nil)).Elem()
}

type recordSetArgs struct {
	// Gets or sets the ETag of the RecordSet.
	Etag *string `pulumi:"etag"`
	// The geo-location where the resource lives
	Location string `pulumi:"location"`
	// Gets or sets the properties of the RecordSet.
	Properties *RecordSetProperties `pulumi:"properties"`
	// The type of DNS record.
	RecordType string `pulumi:"recordType"`
	// The name of the RecordSet, relative to the name of the zone.
	RelativeRecordSetName string `pulumi:"relativeRecordSetName"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// Resource tags.
	Tags map[string]string `pulumi:"tags"`
	// The name of the zone without a terminating dot.
	ZoneName string `pulumi:"zoneName"`
}

// The set of arguments for constructing a RecordSet resource.
type RecordSetArgs struct {
	// Gets or sets the ETag of the RecordSet.
	Etag pulumi.StringPtrInput
	// The geo-location where the resource lives
	Location pulumi.StringInput
	// Gets or sets the properties of the RecordSet.
	Properties RecordSetPropertiesPtrInput
	// The type of DNS record.
	RecordType pulumi.StringInput
	// The name of the RecordSet, relative to the name of the zone.
	RelativeRecordSetName pulumi.StringInput
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput
	// Resource tags.
	Tags pulumi.StringMapInput
	// The name of the zone without a terminating dot.
	ZoneName pulumi.StringInput
}

func (RecordSetArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*recordSetArgs)(nil)).Elem()
}
