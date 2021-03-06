// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v20160330

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Azure Data Catalog.
type ADCCatalog struct {
	pulumi.CustomResourceState

	// Azure data catalog admin list.
	Admins PrincipalsResponseArrayOutput `pulumi:"admins"`
	// Automatic unit adjustment enabled or not.
	EnableAutomaticUnitAdjustment pulumi.BoolPtrOutput `pulumi:"enableAutomaticUnitAdjustment"`
	// Resource etag
	Etag pulumi.StringPtrOutput `pulumi:"etag"`
	// Resource location
	Location pulumi.StringPtrOutput `pulumi:"location"`
	// Resource name
	Name pulumi.StringOutput `pulumi:"name"`
	// Azure data catalog SKU.
	Sku pulumi.StringPtrOutput `pulumi:"sku"`
	// Azure data catalog provision status.
	SuccessfullyProvisioned pulumi.BoolPtrOutput `pulumi:"successfullyProvisioned"`
	// Resource tags
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// Resource type
	Type pulumi.StringOutput `pulumi:"type"`
	// Azure data catalog units.
	Units pulumi.IntPtrOutput `pulumi:"units"`
	// Azure data catalog user list.
	Users PrincipalsResponseArrayOutput `pulumi:"users"`
}

// NewADCCatalog registers a new resource with the given unique name, arguments, and options.
func NewADCCatalog(ctx *pulumi.Context,
	name string, args *ADCCatalogArgs, opts ...pulumi.ResourceOption) (*ADCCatalog, error) {
	if args == nil || args.CatalogName == nil {
		return nil, errors.New("missing required argument 'CatalogName'")
	}
	if args == nil || args.ResourceGroupName == nil {
		return nil, errors.New("missing required argument 'ResourceGroupName'")
	}
	if args == nil {
		args = &ADCCatalogArgs{}
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-nextgen:datacatalog/latest:ADCCatalog"),
		},
	})
	opts = append(opts, aliases)
	var resource ADCCatalog
	err := ctx.RegisterResource("azure-nextgen:datacatalog/v20160330:ADCCatalog", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetADCCatalog gets an existing ADCCatalog resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetADCCatalog(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ADCCatalogState, opts ...pulumi.ResourceOption) (*ADCCatalog, error) {
	var resource ADCCatalog
	err := ctx.ReadResource("azure-nextgen:datacatalog/v20160330:ADCCatalog", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ADCCatalog resources.
type adccatalogState struct {
	// Azure data catalog admin list.
	Admins []PrincipalsResponse `pulumi:"admins"`
	// Automatic unit adjustment enabled or not.
	EnableAutomaticUnitAdjustment *bool `pulumi:"enableAutomaticUnitAdjustment"`
	// Resource etag
	Etag *string `pulumi:"etag"`
	// Resource location
	Location *string `pulumi:"location"`
	// Resource name
	Name *string `pulumi:"name"`
	// Azure data catalog SKU.
	Sku *string `pulumi:"sku"`
	// Azure data catalog provision status.
	SuccessfullyProvisioned *bool `pulumi:"successfullyProvisioned"`
	// Resource tags
	Tags map[string]string `pulumi:"tags"`
	// Resource type
	Type *string `pulumi:"type"`
	// Azure data catalog units.
	Units *int `pulumi:"units"`
	// Azure data catalog user list.
	Users []PrincipalsResponse `pulumi:"users"`
}

type ADCCatalogState struct {
	// Azure data catalog admin list.
	Admins PrincipalsResponseArrayInput
	// Automatic unit adjustment enabled or not.
	EnableAutomaticUnitAdjustment pulumi.BoolPtrInput
	// Resource etag
	Etag pulumi.StringPtrInput
	// Resource location
	Location pulumi.StringPtrInput
	// Resource name
	Name pulumi.StringPtrInput
	// Azure data catalog SKU.
	Sku pulumi.StringPtrInput
	// Azure data catalog provision status.
	SuccessfullyProvisioned pulumi.BoolPtrInput
	// Resource tags
	Tags pulumi.StringMapInput
	// Resource type
	Type pulumi.StringPtrInput
	// Azure data catalog units.
	Units pulumi.IntPtrInput
	// Azure data catalog user list.
	Users PrincipalsResponseArrayInput
}

func (ADCCatalogState) ElementType() reflect.Type {
	return reflect.TypeOf((*adccatalogState)(nil)).Elem()
}

type adccatalogArgs struct {
	// Azure data catalog admin list.
	Admins []Principals `pulumi:"admins"`
	// The name of the data catalog in the specified subscription and resource group.
	CatalogName string `pulumi:"catalogName"`
	// Automatic unit adjustment enabled or not.
	EnableAutomaticUnitAdjustment *bool `pulumi:"enableAutomaticUnitAdjustment"`
	// Resource etag
	Etag *string `pulumi:"etag"`
	// Resource location
	Location *string `pulumi:"location"`
	// The name of the resource group within the user's subscription. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// Azure data catalog SKU.
	Sku *string `pulumi:"sku"`
	// Azure data catalog provision status.
	SuccessfullyProvisioned *bool `pulumi:"successfullyProvisioned"`
	// Resource tags
	Tags map[string]string `pulumi:"tags"`
	// Azure data catalog units.
	Units *int `pulumi:"units"`
	// Azure data catalog user list.
	Users []Principals `pulumi:"users"`
}

// The set of arguments for constructing a ADCCatalog resource.
type ADCCatalogArgs struct {
	// Azure data catalog admin list.
	Admins PrincipalsArrayInput
	// The name of the data catalog in the specified subscription and resource group.
	CatalogName pulumi.StringInput
	// Automatic unit adjustment enabled or not.
	EnableAutomaticUnitAdjustment pulumi.BoolPtrInput
	// Resource etag
	Etag pulumi.StringPtrInput
	// Resource location
	Location pulumi.StringPtrInput
	// The name of the resource group within the user's subscription. The name is case insensitive.
	ResourceGroupName pulumi.StringInput
	// Azure data catalog SKU.
	Sku pulumi.StringPtrInput
	// Azure data catalog provision status.
	SuccessfullyProvisioned pulumi.BoolPtrInput
	// Resource tags
	Tags pulumi.StringMapInput
	// Azure data catalog units.
	Units pulumi.IntPtrInput
	// Azure data catalog user list.
	Users PrincipalsArrayInput
}

func (ADCCatalogArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*adccatalogArgs)(nil)).Elem()
}
