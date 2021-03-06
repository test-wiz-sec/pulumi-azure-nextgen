// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v20200401preview

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Event Channel.
type EventChannel struct {
	pulumi.CustomResourceState

	// Represents the destination of an event channel.
	Destination EventChannelDestinationResponsePtrOutput `pulumi:"destination"`
	// Expiration time of the event channel. If this timer expires while the corresponding partner topic is never activated,
	// the event channel and corresponding partner topic are deleted.
	ExpirationTimeIfNotActivatedUtc pulumi.StringPtrOutput `pulumi:"expirationTimeIfNotActivatedUtc"`
	// Information about the filter for the event channel.
	Filter EventChannelFilterResponsePtrOutput `pulumi:"filter"`
	// Name of the resource
	Name pulumi.StringOutput `pulumi:"name"`
	// Friendly description about the topic. This can be set by the publisher/partner to show custom description for the customer partner topic.
	// This will be helpful to remove any ambiguity of the origin of creation of the partner topic for the customer.
	PartnerTopicFriendlyDescription pulumi.StringPtrOutput `pulumi:"partnerTopicFriendlyDescription"`
	// The readiness state of the corresponding partner topic.
	PartnerTopicReadinessState pulumi.StringOutput `pulumi:"partnerTopicReadinessState"`
	// Provisioning state of the event channel.
	ProvisioningState pulumi.StringOutput `pulumi:"provisioningState"`
	// Source of the event channel. This represents a unique resource in the partner's resource model.
	Source EventChannelSourceResponsePtrOutput `pulumi:"source"`
	// Type of the resource
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewEventChannel registers a new resource with the given unique name, arguments, and options.
func NewEventChannel(ctx *pulumi.Context,
	name string, args *EventChannelArgs, opts ...pulumi.ResourceOption) (*EventChannel, error) {
	if args == nil || args.EventChannelName == nil {
		return nil, errors.New("missing required argument 'EventChannelName'")
	}
	if args == nil || args.PartnerNamespaceName == nil {
		return nil, errors.New("missing required argument 'PartnerNamespaceName'")
	}
	if args == nil || args.ResourceGroupName == nil {
		return nil, errors.New("missing required argument 'ResourceGroupName'")
	}
	if args == nil {
		args = &EventChannelArgs{}
	}
	var resource EventChannel
	err := ctx.RegisterResource("azure-nextgen:eventgrid/v20200401preview:EventChannel", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetEventChannel gets an existing EventChannel resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetEventChannel(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *EventChannelState, opts ...pulumi.ResourceOption) (*EventChannel, error) {
	var resource EventChannel
	err := ctx.ReadResource("azure-nextgen:eventgrid/v20200401preview:EventChannel", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering EventChannel resources.
type eventChannelState struct {
	// Represents the destination of an event channel.
	Destination *EventChannelDestinationResponse `pulumi:"destination"`
	// Expiration time of the event channel. If this timer expires while the corresponding partner topic is never activated,
	// the event channel and corresponding partner topic are deleted.
	ExpirationTimeIfNotActivatedUtc *string `pulumi:"expirationTimeIfNotActivatedUtc"`
	// Information about the filter for the event channel.
	Filter *EventChannelFilterResponse `pulumi:"filter"`
	// Name of the resource
	Name *string `pulumi:"name"`
	// Friendly description about the topic. This can be set by the publisher/partner to show custom description for the customer partner topic.
	// This will be helpful to remove any ambiguity of the origin of creation of the partner topic for the customer.
	PartnerTopicFriendlyDescription *string `pulumi:"partnerTopicFriendlyDescription"`
	// The readiness state of the corresponding partner topic.
	PartnerTopicReadinessState *string `pulumi:"partnerTopicReadinessState"`
	// Provisioning state of the event channel.
	ProvisioningState *string `pulumi:"provisioningState"`
	// Source of the event channel. This represents a unique resource in the partner's resource model.
	Source *EventChannelSourceResponse `pulumi:"source"`
	// Type of the resource
	Type *string `pulumi:"type"`
}

type EventChannelState struct {
	// Represents the destination of an event channel.
	Destination EventChannelDestinationResponsePtrInput
	// Expiration time of the event channel. If this timer expires while the corresponding partner topic is never activated,
	// the event channel and corresponding partner topic are deleted.
	ExpirationTimeIfNotActivatedUtc pulumi.StringPtrInput
	// Information about the filter for the event channel.
	Filter EventChannelFilterResponsePtrInput
	// Name of the resource
	Name pulumi.StringPtrInput
	// Friendly description about the topic. This can be set by the publisher/partner to show custom description for the customer partner topic.
	// This will be helpful to remove any ambiguity of the origin of creation of the partner topic for the customer.
	PartnerTopicFriendlyDescription pulumi.StringPtrInput
	// The readiness state of the corresponding partner topic.
	PartnerTopicReadinessState pulumi.StringPtrInput
	// Provisioning state of the event channel.
	ProvisioningState pulumi.StringPtrInput
	// Source of the event channel. This represents a unique resource in the partner's resource model.
	Source EventChannelSourceResponsePtrInput
	// Type of the resource
	Type pulumi.StringPtrInput
}

func (EventChannelState) ElementType() reflect.Type {
	return reflect.TypeOf((*eventChannelState)(nil)).Elem()
}

type eventChannelArgs struct {
	// Represents the destination of an event channel.
	Destination *EventChannelDestination `pulumi:"destination"`
	// Name of the event channel.
	EventChannelName string `pulumi:"eventChannelName"`
	// Expiration time of the event channel. If this timer expires while the corresponding partner topic is never activated,
	// the event channel and corresponding partner topic are deleted.
	ExpirationTimeIfNotActivatedUtc *string `pulumi:"expirationTimeIfNotActivatedUtc"`
	// Information about the filter for the event channel.
	Filter *EventChannelFilter `pulumi:"filter"`
	// Name of the partner namespace.
	PartnerNamespaceName string `pulumi:"partnerNamespaceName"`
	// Friendly description about the topic. This can be set by the publisher/partner to show custom description for the customer partner topic.
	// This will be helpful to remove any ambiguity of the origin of creation of the partner topic for the customer.
	PartnerTopicFriendlyDescription *string `pulumi:"partnerTopicFriendlyDescription"`
	// The name of the resource group within the user's subscription.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// Source of the event channel. This represents a unique resource in the partner's resource model.
	Source *EventChannelSource `pulumi:"source"`
}

// The set of arguments for constructing a EventChannel resource.
type EventChannelArgs struct {
	// Represents the destination of an event channel.
	Destination EventChannelDestinationPtrInput
	// Name of the event channel.
	EventChannelName pulumi.StringInput
	// Expiration time of the event channel. If this timer expires while the corresponding partner topic is never activated,
	// the event channel and corresponding partner topic are deleted.
	ExpirationTimeIfNotActivatedUtc pulumi.StringPtrInput
	// Information about the filter for the event channel.
	Filter EventChannelFilterPtrInput
	// Name of the partner namespace.
	PartnerNamespaceName pulumi.StringInput
	// Friendly description about the topic. This can be set by the publisher/partner to show custom description for the customer partner topic.
	// This will be helpful to remove any ambiguity of the origin of creation of the partner topic for the customer.
	PartnerTopicFriendlyDescription pulumi.StringPtrInput
	// The name of the resource group within the user's subscription.
	ResourceGroupName pulumi.StringInput
	// Source of the event channel. This represents a unique resource in the partner's resource model.
	Source EventChannelSourcePtrInput
}

func (EventChannelArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*eventChannelArgs)(nil)).Elem()
}
