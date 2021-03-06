// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v20190601

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// The replication policy between two storage accounts. Multiple rules can be defined in one policy.
type ObjectReplicationPolicy struct {
	pulumi.CustomResourceState

	// Required. Destination account name.
	DestinationAccount pulumi.StringOutput `pulumi:"destinationAccount"`
	// Indicates when the policy is enabled on the source account.
	EnabledTime pulumi.StringOutput `pulumi:"enabledTime"`
	// The name of the resource
	Name pulumi.StringOutput `pulumi:"name"`
	// A unique id for object replication policy.
	PolicyId pulumi.StringOutput `pulumi:"policyId"`
	// The storage account object replication rules.
	Rules ObjectReplicationPolicyRuleResponseArrayOutput `pulumi:"rules"`
	// Required. Source account name.
	SourceAccount pulumi.StringOutput `pulumi:"sourceAccount"`
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewObjectReplicationPolicy registers a new resource with the given unique name, arguments, and options.
func NewObjectReplicationPolicy(ctx *pulumi.Context,
	name string, args *ObjectReplicationPolicyArgs, opts ...pulumi.ResourceOption) (*ObjectReplicationPolicy, error) {
	if args == nil || args.AccountName == nil {
		return nil, errors.New("missing required argument 'AccountName'")
	}
	if args == nil || args.DestinationAccount == nil {
		return nil, errors.New("missing required argument 'DestinationAccount'")
	}
	if args == nil || args.ObjectReplicationPolicyId == nil {
		return nil, errors.New("missing required argument 'ObjectReplicationPolicyId'")
	}
	if args == nil || args.ResourceGroupName == nil {
		return nil, errors.New("missing required argument 'ResourceGroupName'")
	}
	if args == nil || args.SourceAccount == nil {
		return nil, errors.New("missing required argument 'SourceAccount'")
	}
	if args == nil {
		args = &ObjectReplicationPolicyArgs{}
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-nextgen:storage/latest:ObjectReplicationPolicy"),
		},
		{
			Type: pulumi.String("azure-nextgen:storage/v20200801preview:ObjectReplicationPolicy"),
		},
	})
	opts = append(opts, aliases)
	var resource ObjectReplicationPolicy
	err := ctx.RegisterResource("azure-nextgen:storage/v20190601:ObjectReplicationPolicy", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetObjectReplicationPolicy gets an existing ObjectReplicationPolicy resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetObjectReplicationPolicy(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ObjectReplicationPolicyState, opts ...pulumi.ResourceOption) (*ObjectReplicationPolicy, error) {
	var resource ObjectReplicationPolicy
	err := ctx.ReadResource("azure-nextgen:storage/v20190601:ObjectReplicationPolicy", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ObjectReplicationPolicy resources.
type objectReplicationPolicyState struct {
	// Required. Destination account name.
	DestinationAccount *string `pulumi:"destinationAccount"`
	// Indicates when the policy is enabled on the source account.
	EnabledTime *string `pulumi:"enabledTime"`
	// The name of the resource
	Name *string `pulumi:"name"`
	// A unique id for object replication policy.
	PolicyId *string `pulumi:"policyId"`
	// The storage account object replication rules.
	Rules []ObjectReplicationPolicyRuleResponse `pulumi:"rules"`
	// Required. Source account name.
	SourceAccount *string `pulumi:"sourceAccount"`
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type *string `pulumi:"type"`
}

type ObjectReplicationPolicyState struct {
	// Required. Destination account name.
	DestinationAccount pulumi.StringPtrInput
	// Indicates when the policy is enabled on the source account.
	EnabledTime pulumi.StringPtrInput
	// The name of the resource
	Name pulumi.StringPtrInput
	// A unique id for object replication policy.
	PolicyId pulumi.StringPtrInput
	// The storage account object replication rules.
	Rules ObjectReplicationPolicyRuleResponseArrayInput
	// Required. Source account name.
	SourceAccount pulumi.StringPtrInput
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type pulumi.StringPtrInput
}

func (ObjectReplicationPolicyState) ElementType() reflect.Type {
	return reflect.TypeOf((*objectReplicationPolicyState)(nil)).Elem()
}

type objectReplicationPolicyArgs struct {
	// The name of the storage account within the specified resource group. Storage account names must be between 3 and 24 characters in length and use numbers and lower-case letters only.
	AccountName string `pulumi:"accountName"`
	// Required. Destination account name.
	DestinationAccount string `pulumi:"destinationAccount"`
	// The ID of object replication policy or 'default' if the policy ID is unknown.
	ObjectReplicationPolicyId string `pulumi:"objectReplicationPolicyId"`
	// The name of the resource group within the user's subscription. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The storage account object replication rules.
	Rules []ObjectReplicationPolicyRule `pulumi:"rules"`
	// Required. Source account name.
	SourceAccount string `pulumi:"sourceAccount"`
}

// The set of arguments for constructing a ObjectReplicationPolicy resource.
type ObjectReplicationPolicyArgs struct {
	// The name of the storage account within the specified resource group. Storage account names must be between 3 and 24 characters in length and use numbers and lower-case letters only.
	AccountName pulumi.StringInput
	// Required. Destination account name.
	DestinationAccount pulumi.StringInput
	// The ID of object replication policy or 'default' if the policy ID is unknown.
	ObjectReplicationPolicyId pulumi.StringInput
	// The name of the resource group within the user's subscription. The name is case insensitive.
	ResourceGroupName pulumi.StringInput
	// The storage account object replication rules.
	Rules ObjectReplicationPolicyRuleArrayInput
	// Required. Source account name.
	SourceAccount pulumi.StringInput
}

func (ObjectReplicationPolicyArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*objectReplicationPolicyArgs)(nil)).Elem()
}
