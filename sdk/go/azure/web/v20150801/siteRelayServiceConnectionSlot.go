// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v20150801

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Class that represents a BizTalk Hybrid Connection
type SiteRelayServiceConnectionSlot struct {
	pulumi.CustomResourceState

	BiztalkUri             pulumi.StringPtrOutput `pulumi:"biztalkUri"`
	EntityConnectionString pulumi.StringPtrOutput `pulumi:"entityConnectionString"`
	EntityName             pulumi.StringPtrOutput `pulumi:"entityName"`
	Hostname               pulumi.StringPtrOutput `pulumi:"hostname"`
	// Kind of resource
	Kind pulumi.StringPtrOutput `pulumi:"kind"`
	// Resource Location
	Location pulumi.StringOutput `pulumi:"location"`
	// Resource Name
	Name                     pulumi.StringPtrOutput `pulumi:"name"`
	Port                     pulumi.IntPtrOutput    `pulumi:"port"`
	ResourceConnectionString pulumi.StringPtrOutput `pulumi:"resourceConnectionString"`
	ResourceType             pulumi.StringPtrOutput `pulumi:"resourceType"`
	// Resource tags
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// Resource type
	Type pulumi.StringPtrOutput `pulumi:"type"`
}

// NewSiteRelayServiceConnectionSlot registers a new resource with the given unique name, arguments, and options.
func NewSiteRelayServiceConnectionSlot(ctx *pulumi.Context,
	name string, args *SiteRelayServiceConnectionSlotArgs, opts ...pulumi.ResourceOption) (*SiteRelayServiceConnectionSlot, error) {
	if args == nil || args.EntityName == nil {
		return nil, errors.New("missing required argument 'EntityName'")
	}
	if args == nil || args.Location == nil {
		return nil, errors.New("missing required argument 'Location'")
	}
	if args == nil || args.Name == nil {
		return nil, errors.New("missing required argument 'Name'")
	}
	if args == nil || args.ResourceGroupName == nil {
		return nil, errors.New("missing required argument 'ResourceGroupName'")
	}
	if args == nil || args.Slot == nil {
		return nil, errors.New("missing required argument 'Slot'")
	}
	if args == nil {
		args = &SiteRelayServiceConnectionSlotArgs{}
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-nextgen:web/latest:SiteRelayServiceConnectionSlot"),
		},
		{
			Type: pulumi.String("azure-nextgen:web/v20160801:SiteRelayServiceConnectionSlot"),
		},
		{
			Type: pulumi.String("azure-nextgen:web/v20180201:SiteRelayServiceConnectionSlot"),
		},
		{
			Type: pulumi.String("azure-nextgen:web/v20181101:SiteRelayServiceConnectionSlot"),
		},
		{
			Type: pulumi.String("azure-nextgen:web/v20190801:SiteRelayServiceConnectionSlot"),
		},
		{
			Type: pulumi.String("azure-nextgen:web/v20200601:SiteRelayServiceConnectionSlot"),
		},
	})
	opts = append(opts, aliases)
	var resource SiteRelayServiceConnectionSlot
	err := ctx.RegisterResource("azure-nextgen:web/v20150801:SiteRelayServiceConnectionSlot", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetSiteRelayServiceConnectionSlot gets an existing SiteRelayServiceConnectionSlot resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetSiteRelayServiceConnectionSlot(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SiteRelayServiceConnectionSlotState, opts ...pulumi.ResourceOption) (*SiteRelayServiceConnectionSlot, error) {
	var resource SiteRelayServiceConnectionSlot
	err := ctx.ReadResource("azure-nextgen:web/v20150801:SiteRelayServiceConnectionSlot", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering SiteRelayServiceConnectionSlot resources.
type siteRelayServiceConnectionSlotState struct {
	BiztalkUri             *string `pulumi:"biztalkUri"`
	EntityConnectionString *string `pulumi:"entityConnectionString"`
	EntityName             *string `pulumi:"entityName"`
	Hostname               *string `pulumi:"hostname"`
	// Kind of resource
	Kind *string `pulumi:"kind"`
	// Resource Location
	Location *string `pulumi:"location"`
	// Resource Name
	Name                     *string `pulumi:"name"`
	Port                     *int    `pulumi:"port"`
	ResourceConnectionString *string `pulumi:"resourceConnectionString"`
	ResourceType             *string `pulumi:"resourceType"`
	// Resource tags
	Tags map[string]string `pulumi:"tags"`
	// Resource type
	Type *string `pulumi:"type"`
}

type SiteRelayServiceConnectionSlotState struct {
	BiztalkUri             pulumi.StringPtrInput
	EntityConnectionString pulumi.StringPtrInput
	EntityName             pulumi.StringPtrInput
	Hostname               pulumi.StringPtrInput
	// Kind of resource
	Kind pulumi.StringPtrInput
	// Resource Location
	Location pulumi.StringPtrInput
	// Resource Name
	Name                     pulumi.StringPtrInput
	Port                     pulumi.IntPtrInput
	ResourceConnectionString pulumi.StringPtrInput
	ResourceType             pulumi.StringPtrInput
	// Resource tags
	Tags pulumi.StringMapInput
	// Resource type
	Type pulumi.StringPtrInput
}

func (SiteRelayServiceConnectionSlotState) ElementType() reflect.Type {
	return reflect.TypeOf((*siteRelayServiceConnectionSlotState)(nil)).Elem()
}

type siteRelayServiceConnectionSlotArgs struct {
	BiztalkUri             *string `pulumi:"biztalkUri"`
	EntityConnectionString *string `pulumi:"entityConnectionString"`
	EntityName             string  `pulumi:"entityName"`
	Hostname               *string `pulumi:"hostname"`
	// Resource Id
	Id *string `pulumi:"id"`
	// Kind of resource
	Kind *string `pulumi:"kind"`
	// Resource Location
	Location string `pulumi:"location"`
	// Resource Name
	Name                     string  `pulumi:"name"`
	Port                     *int    `pulumi:"port"`
	ResourceConnectionString *string `pulumi:"resourceConnectionString"`
	// The resource group name
	ResourceGroupName string  `pulumi:"resourceGroupName"`
	ResourceType      *string `pulumi:"resourceType"`
	// The name of the slot for the web app.
	Slot string `pulumi:"slot"`
	// Resource tags
	Tags map[string]string `pulumi:"tags"`
	// Resource type
	Type *string `pulumi:"type"`
}

// The set of arguments for constructing a SiteRelayServiceConnectionSlot resource.
type SiteRelayServiceConnectionSlotArgs struct {
	BiztalkUri             pulumi.StringPtrInput
	EntityConnectionString pulumi.StringPtrInput
	EntityName             pulumi.StringInput
	Hostname               pulumi.StringPtrInput
	// Resource Id
	Id pulumi.StringPtrInput
	// Kind of resource
	Kind pulumi.StringPtrInput
	// Resource Location
	Location pulumi.StringInput
	// Resource Name
	Name                     pulumi.StringInput
	Port                     pulumi.IntPtrInput
	ResourceConnectionString pulumi.StringPtrInput
	// The resource group name
	ResourceGroupName pulumi.StringInput
	ResourceType      pulumi.StringPtrInput
	// The name of the slot for the web app.
	Slot pulumi.StringInput
	// Resource tags
	Tags pulumi.StringMapInput
	// Resource type
	Type pulumi.StringPtrInput
}

func (SiteRelayServiceConnectionSlotArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*siteRelayServiceConnectionSlotArgs)(nil)).Elem()
}
