// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v20200501preview

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// The bandwidth schedule details.
type BandwidthSchedule struct {
	pulumi.CustomResourceState

	// The days of the week when this schedule is applicable.
	Days pulumi.StringArrayOutput `pulumi:"days"`
	// The object name.
	Name pulumi.StringOutput `pulumi:"name"`
	// The bandwidth rate in Mbps.
	RateInMbps pulumi.IntOutput `pulumi:"rateInMbps"`
	// The start time of the schedule in UTC.
	Start pulumi.StringOutput `pulumi:"start"`
	// The stop time of the schedule in UTC.
	Stop pulumi.StringOutput `pulumi:"stop"`
	// The hierarchical type of the object.
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewBandwidthSchedule registers a new resource with the given unique name, arguments, and options.
func NewBandwidthSchedule(ctx *pulumi.Context,
	name string, args *BandwidthScheduleArgs, opts ...pulumi.ResourceOption) (*BandwidthSchedule, error) {
	if args == nil || args.Days == nil {
		return nil, errors.New("missing required argument 'Days'")
	}
	if args == nil || args.DeviceName == nil {
		return nil, errors.New("missing required argument 'DeviceName'")
	}
	if args == nil || args.Name == nil {
		return nil, errors.New("missing required argument 'Name'")
	}
	if args == nil || args.RateInMbps == nil {
		return nil, errors.New("missing required argument 'RateInMbps'")
	}
	if args == nil || args.ResourceGroupName == nil {
		return nil, errors.New("missing required argument 'ResourceGroupName'")
	}
	if args == nil || args.Start == nil {
		return nil, errors.New("missing required argument 'Start'")
	}
	if args == nil || args.Stop == nil {
		return nil, errors.New("missing required argument 'Stop'")
	}
	if args == nil {
		args = &BandwidthScheduleArgs{}
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-nextgen:databoxedge/latest:BandwidthSchedule"),
		},
		{
			Type: pulumi.String("azure-nextgen:databoxedge/v20190301:BandwidthSchedule"),
		},
		{
			Type: pulumi.String("azure-nextgen:databoxedge/v20190701:BandwidthSchedule"),
		},
		{
			Type: pulumi.String("azure-nextgen:databoxedge/v20190801:BandwidthSchedule"),
		},
	})
	opts = append(opts, aliases)
	var resource BandwidthSchedule
	err := ctx.RegisterResource("azure-nextgen:databoxedge/v20200501preview:BandwidthSchedule", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetBandwidthSchedule gets an existing BandwidthSchedule resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetBandwidthSchedule(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *BandwidthScheduleState, opts ...pulumi.ResourceOption) (*BandwidthSchedule, error) {
	var resource BandwidthSchedule
	err := ctx.ReadResource("azure-nextgen:databoxedge/v20200501preview:BandwidthSchedule", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering BandwidthSchedule resources.
type bandwidthScheduleState struct {
	// The days of the week when this schedule is applicable.
	Days []string `pulumi:"days"`
	// The object name.
	Name *string `pulumi:"name"`
	// The bandwidth rate in Mbps.
	RateInMbps *int `pulumi:"rateInMbps"`
	// The start time of the schedule in UTC.
	Start *string `pulumi:"start"`
	// The stop time of the schedule in UTC.
	Stop *string `pulumi:"stop"`
	// The hierarchical type of the object.
	Type *string `pulumi:"type"`
}

type BandwidthScheduleState struct {
	// The days of the week when this schedule is applicable.
	Days pulumi.StringArrayInput
	// The object name.
	Name pulumi.StringPtrInput
	// The bandwidth rate in Mbps.
	RateInMbps pulumi.IntPtrInput
	// The start time of the schedule in UTC.
	Start pulumi.StringPtrInput
	// The stop time of the schedule in UTC.
	Stop pulumi.StringPtrInput
	// The hierarchical type of the object.
	Type pulumi.StringPtrInput
}

func (BandwidthScheduleState) ElementType() reflect.Type {
	return reflect.TypeOf((*bandwidthScheduleState)(nil)).Elem()
}

type bandwidthScheduleArgs struct {
	// The days of the week when this schedule is applicable.
	Days []string `pulumi:"days"`
	// The device name.
	DeviceName string `pulumi:"deviceName"`
	// The bandwidth schedule name which needs to be added/updated.
	Name string `pulumi:"name"`
	// The bandwidth rate in Mbps.
	RateInMbps int `pulumi:"rateInMbps"`
	// The resource group name.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The start time of the schedule in UTC.
	Start string `pulumi:"start"`
	// The stop time of the schedule in UTC.
	Stop string `pulumi:"stop"`
}

// The set of arguments for constructing a BandwidthSchedule resource.
type BandwidthScheduleArgs struct {
	// The days of the week when this schedule is applicable.
	Days pulumi.StringArrayInput
	// The device name.
	DeviceName pulumi.StringInput
	// The bandwidth schedule name which needs to be added/updated.
	Name pulumi.StringInput
	// The bandwidth rate in Mbps.
	RateInMbps pulumi.IntInput
	// The resource group name.
	ResourceGroupName pulumi.StringInput
	// The start time of the schedule in UTC.
	Start pulumi.StringInput
	// The stop time of the schedule in UTC.
	Stop pulumi.StringInput
}

func (BandwidthScheduleArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*bandwidthScheduleArgs)(nil)).Elem()
}