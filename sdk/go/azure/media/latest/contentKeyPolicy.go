// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package latest

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// A Content Key Policy resource.
type ContentKeyPolicy struct {
	pulumi.CustomResourceState

	// The creation date of the Policy
	Created pulumi.StringOutput `pulumi:"created"`
	// A description for the Policy.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// The last modified date of the Policy
	LastModified pulumi.StringOutput `pulumi:"lastModified"`
	// The name of the resource
	Name pulumi.StringOutput `pulumi:"name"`
	// The Key Policy options.
	Options ContentKeyPolicyOptionResponseArrayOutput `pulumi:"options"`
	// The legacy Policy ID.
	PolicyId pulumi.StringOutput `pulumi:"policyId"`
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewContentKeyPolicy registers a new resource with the given unique name, arguments, and options.
func NewContentKeyPolicy(ctx *pulumi.Context,
	name string, args *ContentKeyPolicyArgs, opts ...pulumi.ResourceOption) (*ContentKeyPolicy, error) {
	if args == nil || args.AccountName == nil {
		return nil, errors.New("missing required argument 'AccountName'")
	}
	if args == nil || args.ContentKeyPolicyName == nil {
		return nil, errors.New("missing required argument 'ContentKeyPolicyName'")
	}
	if args == nil || args.Options == nil {
		return nil, errors.New("missing required argument 'Options'")
	}
	if args == nil || args.ResourceGroupName == nil {
		return nil, errors.New("missing required argument 'ResourceGroupName'")
	}
	if args == nil {
		args = &ContentKeyPolicyArgs{}
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-nextgen:media/v20180330preview:ContentKeyPolicy"),
		},
		{
			Type: pulumi.String("azure-nextgen:media/v20180601preview:ContentKeyPolicy"),
		},
		{
			Type: pulumi.String("azure-nextgen:media/v20180701:ContentKeyPolicy"),
		},
		{
			Type: pulumi.String("azure-nextgen:media/v20200501:ContentKeyPolicy"),
		},
	})
	opts = append(opts, aliases)
	var resource ContentKeyPolicy
	err := ctx.RegisterResource("azure-nextgen:media/latest:ContentKeyPolicy", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetContentKeyPolicy gets an existing ContentKeyPolicy resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetContentKeyPolicy(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ContentKeyPolicyState, opts ...pulumi.ResourceOption) (*ContentKeyPolicy, error) {
	var resource ContentKeyPolicy
	err := ctx.ReadResource("azure-nextgen:media/latest:ContentKeyPolicy", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ContentKeyPolicy resources.
type contentKeyPolicyState struct {
	// The creation date of the Policy
	Created *string `pulumi:"created"`
	// A description for the Policy.
	Description *string `pulumi:"description"`
	// The last modified date of the Policy
	LastModified *string `pulumi:"lastModified"`
	// The name of the resource
	Name *string `pulumi:"name"`
	// The Key Policy options.
	Options []ContentKeyPolicyOptionResponse `pulumi:"options"`
	// The legacy Policy ID.
	PolicyId *string `pulumi:"policyId"`
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type *string `pulumi:"type"`
}

type ContentKeyPolicyState struct {
	// The creation date of the Policy
	Created pulumi.StringPtrInput
	// A description for the Policy.
	Description pulumi.StringPtrInput
	// The last modified date of the Policy
	LastModified pulumi.StringPtrInput
	// The name of the resource
	Name pulumi.StringPtrInput
	// The Key Policy options.
	Options ContentKeyPolicyOptionResponseArrayInput
	// The legacy Policy ID.
	PolicyId pulumi.StringPtrInput
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type pulumi.StringPtrInput
}

func (ContentKeyPolicyState) ElementType() reflect.Type {
	return reflect.TypeOf((*contentKeyPolicyState)(nil)).Elem()
}

type contentKeyPolicyArgs struct {
	// The Media Services account name.
	AccountName string `pulumi:"accountName"`
	// The Content Key Policy name.
	ContentKeyPolicyName string `pulumi:"contentKeyPolicyName"`
	// A description for the Policy.
	Description *string `pulumi:"description"`
	// The Key Policy options.
	Options []ContentKeyPolicyOption `pulumi:"options"`
	// The name of the resource group within the Azure subscription.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// The set of arguments for constructing a ContentKeyPolicy resource.
type ContentKeyPolicyArgs struct {
	// The Media Services account name.
	AccountName pulumi.StringInput
	// The Content Key Policy name.
	ContentKeyPolicyName pulumi.StringInput
	// A description for the Policy.
	Description pulumi.StringPtrInput
	// The Key Policy options.
	Options ContentKeyPolicyOptionArrayInput
	// The name of the resource group within the Azure subscription.
	ResourceGroupName pulumi.StringInput
}

func (ContentKeyPolicyArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*contentKeyPolicyArgs)(nil)).Elem()
}
