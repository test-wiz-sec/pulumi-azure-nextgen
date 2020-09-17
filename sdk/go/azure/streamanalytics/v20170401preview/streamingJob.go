// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v20170401preview

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// A streaming job object, containing all information associated with the named streaming job.
type StreamingJob struct {
	pulumi.CustomResourceState

	// The cluster which streaming jobs will run on.
	Cluster ClusterInfoResponsePtrOutput `pulumi:"cluster"`
	// Controls certain runtime behaviors of the streaming job.
	CompatibilityLevel pulumi.StringPtrOutput `pulumi:"compatibilityLevel"`
	// Valid values are JobStorageAccount and SystemAccount. If set to JobStorageAccount, this requires the user to also specify jobStorageAccount property. .
	ContentStoragePolicy pulumi.StringOutput `pulumi:"contentStoragePolicy"`
	// Value is an ISO-8601 formatted UTC timestamp indicating when the streaming job was created.
	CreatedDate pulumi.StringOutput `pulumi:"createdDate"`
	// The data locale of the stream analytics job. Value should be the name of a supported .NET Culture from the set https://msdn.microsoft.com/en-us/library/system.globalization.culturetypes(v=vs.110).aspx. Defaults to 'en-US' if none specified.
	DataLocale pulumi.StringPtrOutput `pulumi:"dataLocale"`
	// The current entity tag for the streaming job. This is an opaque string. You can use it to detect whether the resource has changed between requests. You can also use it in the If-Match or If-None-Match headers for write operations for optimistic concurrency.
	Etag pulumi.StringOutput `pulumi:"etag"`
	// The maximum tolerable delay in seconds where events arriving late could be included.  Supported range is -1 to 1814399 (20.23:59:59 days) and -1 is used to specify wait indefinitely. If the property is absent, it is interpreted to have a value of -1.
	EventsLateArrivalMaxDelayInSeconds pulumi.IntPtrOutput `pulumi:"eventsLateArrivalMaxDelayInSeconds"`
	// The maximum tolerable delay in seconds where out-of-order events can be adjusted to be back in order.
	EventsOutOfOrderMaxDelayInSeconds pulumi.IntPtrOutput `pulumi:"eventsOutOfOrderMaxDelayInSeconds"`
	// Indicates the policy to apply to events that arrive out of order in the input event stream.
	EventsOutOfOrderPolicy pulumi.StringPtrOutput `pulumi:"eventsOutOfOrderPolicy"`
	// The storage account where the custom code artifacts are located.
	Externals ExternalResponsePtrOutput `pulumi:"externals"`
	// A list of one or more functions for the streaming job. The name property for each function is required when specifying this property in a PUT request. This property cannot be modify via a PATCH operation. You must use the PATCH API available for the individual transformation.
	Functions FunctionResponseArrayOutput `pulumi:"functions"`
	// Describes the system-assigned managed identity assigned to this job that can be used to authenticate with inputs and outputs.
	Identity IdentityResponsePtrOutput `pulumi:"identity"`
	// A list of one or more inputs to the streaming job. The name property for each input is required when specifying this property in a PUT request. This property cannot be modify via a PATCH operation. You must use the PATCH API available for the individual input.
	Inputs InputResponseArrayOutput `pulumi:"inputs"`
	// A GUID uniquely identifying the streaming job. This GUID is generated upon creation of the streaming job.
	JobId pulumi.StringOutput `pulumi:"jobId"`
	// Describes the state of the streaming job.
	JobState pulumi.StringOutput `pulumi:"jobState"`
	// The properties that are associated with an Azure Storage account with MSI
	JobStorageAccount JobStorageAccountResponsePtrOutput `pulumi:"jobStorageAccount"`
	// Describes the type of the job. Valid modes are `Cloud` and 'Edge'.
	JobType pulumi.StringPtrOutput `pulumi:"jobType"`
	// Value is either an ISO-8601 formatted timestamp indicating the last output event time of the streaming job or null indicating that output has not yet been produced. In case of multiple outputs or multiple streams, this shows the latest value in that set.
	LastOutputEventTime pulumi.StringOutput `pulumi:"lastOutputEventTime"`
	// The geo-location where the resource lives
	Location pulumi.StringPtrOutput `pulumi:"location"`
	// The name of the resource
	Name pulumi.StringOutput `pulumi:"name"`
	// Indicates the policy to apply to events that arrive at the output and cannot be written to the external storage due to being malformed (missing column values, column values of wrong type or size).
	OutputErrorPolicy pulumi.StringPtrOutput `pulumi:"outputErrorPolicy"`
	// This property should only be utilized when it is desired that the job be started immediately upon creation. Value may be JobStartTime, CustomTime, or LastOutputEventTime to indicate whether the starting point of the output event stream should start whenever the job is started, start at a custom user time stamp specified via the outputStartTime property, or start from the last event output time.
	OutputStartMode pulumi.StringPtrOutput `pulumi:"outputStartMode"`
	// Value is either an ISO-8601 formatted time stamp that indicates the starting point of the output event stream, or null to indicate that the output event stream will start whenever the streaming job is started. This property must have a value if outputStartMode is set to CustomTime.
	OutputStartTime pulumi.StringPtrOutput `pulumi:"outputStartTime"`
	// A list of one or more outputs for the streaming job. The name property for each output is required when specifying this property in a PUT request. This property cannot be modify via a PATCH operation. You must use the PATCH API available for the individual output.
	Outputs OutputResponseArrayOutput `pulumi:"outputs"`
	// Describes the provisioning status of the streaming job.
	ProvisioningState pulumi.StringOutput `pulumi:"provisioningState"`
	// Describes the SKU of the streaming job. Required on PUT (CreateOrReplace) requests.
	Sku StreamingJobSkuResponsePtrOutput `pulumi:"sku"`
	// Resource tags.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// Indicates the query and the number of streaming units to use for the streaming job. The name property of the transformation is required when specifying this property in a PUT request. This property cannot be modify via a PATCH operation. You must use the PATCH API available for the individual transformation.
	Transformation TransformationResponsePtrOutput `pulumi:"transformation"`
	// The type of the resource. Ex- Microsoft.Compute/virtualMachines or Microsoft.Storage/storageAccounts.
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewStreamingJob registers a new resource with the given unique name, arguments, and options.
func NewStreamingJob(ctx *pulumi.Context,
	name string, args *StreamingJobArgs, opts ...pulumi.ResourceOption) (*StreamingJob, error) {
	if args == nil || args.JobName == nil {
		return nil, errors.New("missing required argument 'JobName'")
	}
	if args == nil || args.ResourceGroupName == nil {
		return nil, errors.New("missing required argument 'ResourceGroupName'")
	}
	if args == nil {
		args = &StreamingJobArgs{}
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-nextgen:streamanalytics/latest:StreamingJob"),
		},
		{
			Type: pulumi.String("azure-nextgen:streamanalytics/v20160301:StreamingJob"),
		},
	})
	opts = append(opts, aliases)
	var resource StreamingJob
	err := ctx.RegisterResource("azure-nextgen:streamanalytics/v20170401preview:StreamingJob", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetStreamingJob gets an existing StreamingJob resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetStreamingJob(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *StreamingJobState, opts ...pulumi.ResourceOption) (*StreamingJob, error) {
	var resource StreamingJob
	err := ctx.ReadResource("azure-nextgen:streamanalytics/v20170401preview:StreamingJob", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering StreamingJob resources.
type streamingJobState struct {
	// The cluster which streaming jobs will run on.
	Cluster *ClusterInfoResponse `pulumi:"cluster"`
	// Controls certain runtime behaviors of the streaming job.
	CompatibilityLevel *string `pulumi:"compatibilityLevel"`
	// Valid values are JobStorageAccount and SystemAccount. If set to JobStorageAccount, this requires the user to also specify jobStorageAccount property. .
	ContentStoragePolicy *string `pulumi:"contentStoragePolicy"`
	// Value is an ISO-8601 formatted UTC timestamp indicating when the streaming job was created.
	CreatedDate *string `pulumi:"createdDate"`
	// The data locale of the stream analytics job. Value should be the name of a supported .NET Culture from the set https://msdn.microsoft.com/en-us/library/system.globalization.culturetypes(v=vs.110).aspx. Defaults to 'en-US' if none specified.
	DataLocale *string `pulumi:"dataLocale"`
	// The current entity tag for the streaming job. This is an opaque string. You can use it to detect whether the resource has changed between requests. You can also use it in the If-Match or If-None-Match headers for write operations for optimistic concurrency.
	Etag *string `pulumi:"etag"`
	// The maximum tolerable delay in seconds where events arriving late could be included.  Supported range is -1 to 1814399 (20.23:59:59 days) and -1 is used to specify wait indefinitely. If the property is absent, it is interpreted to have a value of -1.
	EventsLateArrivalMaxDelayInSeconds *int `pulumi:"eventsLateArrivalMaxDelayInSeconds"`
	// The maximum tolerable delay in seconds where out-of-order events can be adjusted to be back in order.
	EventsOutOfOrderMaxDelayInSeconds *int `pulumi:"eventsOutOfOrderMaxDelayInSeconds"`
	// Indicates the policy to apply to events that arrive out of order in the input event stream.
	EventsOutOfOrderPolicy *string `pulumi:"eventsOutOfOrderPolicy"`
	// The storage account where the custom code artifacts are located.
	Externals *ExternalResponse `pulumi:"externals"`
	// A list of one or more functions for the streaming job. The name property for each function is required when specifying this property in a PUT request. This property cannot be modify via a PATCH operation. You must use the PATCH API available for the individual transformation.
	Functions []FunctionResponse `pulumi:"functions"`
	// Describes the system-assigned managed identity assigned to this job that can be used to authenticate with inputs and outputs.
	Identity *IdentityResponse `pulumi:"identity"`
	// A list of one or more inputs to the streaming job. The name property for each input is required when specifying this property in a PUT request. This property cannot be modify via a PATCH operation. You must use the PATCH API available for the individual input.
	Inputs []InputResponse `pulumi:"inputs"`
	// A GUID uniquely identifying the streaming job. This GUID is generated upon creation of the streaming job.
	JobId *string `pulumi:"jobId"`
	// Describes the state of the streaming job.
	JobState *string `pulumi:"jobState"`
	// The properties that are associated with an Azure Storage account with MSI
	JobStorageAccount *JobStorageAccountResponse `pulumi:"jobStorageAccount"`
	// Describes the type of the job. Valid modes are `Cloud` and 'Edge'.
	JobType *string `pulumi:"jobType"`
	// Value is either an ISO-8601 formatted timestamp indicating the last output event time of the streaming job or null indicating that output has not yet been produced. In case of multiple outputs or multiple streams, this shows the latest value in that set.
	LastOutputEventTime *string `pulumi:"lastOutputEventTime"`
	// The geo-location where the resource lives
	Location *string `pulumi:"location"`
	// The name of the resource
	Name *string `pulumi:"name"`
	// Indicates the policy to apply to events that arrive at the output and cannot be written to the external storage due to being malformed (missing column values, column values of wrong type or size).
	OutputErrorPolicy *string `pulumi:"outputErrorPolicy"`
	// This property should only be utilized when it is desired that the job be started immediately upon creation. Value may be JobStartTime, CustomTime, or LastOutputEventTime to indicate whether the starting point of the output event stream should start whenever the job is started, start at a custom user time stamp specified via the outputStartTime property, or start from the last event output time.
	OutputStartMode *string `pulumi:"outputStartMode"`
	// Value is either an ISO-8601 formatted time stamp that indicates the starting point of the output event stream, or null to indicate that the output event stream will start whenever the streaming job is started. This property must have a value if outputStartMode is set to CustomTime.
	OutputStartTime *string `pulumi:"outputStartTime"`
	// A list of one or more outputs for the streaming job. The name property for each output is required when specifying this property in a PUT request. This property cannot be modify via a PATCH operation. You must use the PATCH API available for the individual output.
	Outputs []OutputResponse `pulumi:"outputs"`
	// Describes the provisioning status of the streaming job.
	ProvisioningState *string `pulumi:"provisioningState"`
	// Describes the SKU of the streaming job. Required on PUT (CreateOrReplace) requests.
	Sku *StreamingJobSkuResponse `pulumi:"sku"`
	// Resource tags.
	Tags map[string]string `pulumi:"tags"`
	// Indicates the query and the number of streaming units to use for the streaming job. The name property of the transformation is required when specifying this property in a PUT request. This property cannot be modify via a PATCH operation. You must use the PATCH API available for the individual transformation.
	Transformation *TransformationResponse `pulumi:"transformation"`
	// The type of the resource. Ex- Microsoft.Compute/virtualMachines or Microsoft.Storage/storageAccounts.
	Type *string `pulumi:"type"`
}

type StreamingJobState struct {
	// The cluster which streaming jobs will run on.
	Cluster ClusterInfoResponsePtrInput
	// Controls certain runtime behaviors of the streaming job.
	CompatibilityLevel pulumi.StringPtrInput
	// Valid values are JobStorageAccount and SystemAccount. If set to JobStorageAccount, this requires the user to also specify jobStorageAccount property. .
	ContentStoragePolicy pulumi.StringPtrInput
	// Value is an ISO-8601 formatted UTC timestamp indicating when the streaming job was created.
	CreatedDate pulumi.StringPtrInput
	// The data locale of the stream analytics job. Value should be the name of a supported .NET Culture from the set https://msdn.microsoft.com/en-us/library/system.globalization.culturetypes(v=vs.110).aspx. Defaults to 'en-US' if none specified.
	DataLocale pulumi.StringPtrInput
	// The current entity tag for the streaming job. This is an opaque string. You can use it to detect whether the resource has changed between requests. You can also use it in the If-Match or If-None-Match headers for write operations for optimistic concurrency.
	Etag pulumi.StringPtrInput
	// The maximum tolerable delay in seconds where events arriving late could be included.  Supported range is -1 to 1814399 (20.23:59:59 days) and -1 is used to specify wait indefinitely. If the property is absent, it is interpreted to have a value of -1.
	EventsLateArrivalMaxDelayInSeconds pulumi.IntPtrInput
	// The maximum tolerable delay in seconds where out-of-order events can be adjusted to be back in order.
	EventsOutOfOrderMaxDelayInSeconds pulumi.IntPtrInput
	// Indicates the policy to apply to events that arrive out of order in the input event stream.
	EventsOutOfOrderPolicy pulumi.StringPtrInput
	// The storage account where the custom code artifacts are located.
	Externals ExternalResponsePtrInput
	// A list of one or more functions for the streaming job. The name property for each function is required when specifying this property in a PUT request. This property cannot be modify via a PATCH operation. You must use the PATCH API available for the individual transformation.
	Functions FunctionResponseArrayInput
	// Describes the system-assigned managed identity assigned to this job that can be used to authenticate with inputs and outputs.
	Identity IdentityResponsePtrInput
	// A list of one or more inputs to the streaming job. The name property for each input is required when specifying this property in a PUT request. This property cannot be modify via a PATCH operation. You must use the PATCH API available for the individual input.
	Inputs InputResponseArrayInput
	// A GUID uniquely identifying the streaming job. This GUID is generated upon creation of the streaming job.
	JobId pulumi.StringPtrInput
	// Describes the state of the streaming job.
	JobState pulumi.StringPtrInput
	// The properties that are associated with an Azure Storage account with MSI
	JobStorageAccount JobStorageAccountResponsePtrInput
	// Describes the type of the job. Valid modes are `Cloud` and 'Edge'.
	JobType pulumi.StringPtrInput
	// Value is either an ISO-8601 formatted timestamp indicating the last output event time of the streaming job or null indicating that output has not yet been produced. In case of multiple outputs or multiple streams, this shows the latest value in that set.
	LastOutputEventTime pulumi.StringPtrInput
	// The geo-location where the resource lives
	Location pulumi.StringPtrInput
	// The name of the resource
	Name pulumi.StringPtrInput
	// Indicates the policy to apply to events that arrive at the output and cannot be written to the external storage due to being malformed (missing column values, column values of wrong type or size).
	OutputErrorPolicy pulumi.StringPtrInput
	// This property should only be utilized when it is desired that the job be started immediately upon creation. Value may be JobStartTime, CustomTime, or LastOutputEventTime to indicate whether the starting point of the output event stream should start whenever the job is started, start at a custom user time stamp specified via the outputStartTime property, or start from the last event output time.
	OutputStartMode pulumi.StringPtrInput
	// Value is either an ISO-8601 formatted time stamp that indicates the starting point of the output event stream, or null to indicate that the output event stream will start whenever the streaming job is started. This property must have a value if outputStartMode is set to CustomTime.
	OutputStartTime pulumi.StringPtrInput
	// A list of one or more outputs for the streaming job. The name property for each output is required when specifying this property in a PUT request. This property cannot be modify via a PATCH operation. You must use the PATCH API available for the individual output.
	Outputs OutputResponseArrayInput
	// Describes the provisioning status of the streaming job.
	ProvisioningState pulumi.StringPtrInput
	// Describes the SKU of the streaming job. Required on PUT (CreateOrReplace) requests.
	Sku StreamingJobSkuResponsePtrInput
	// Resource tags.
	Tags pulumi.StringMapInput
	// Indicates the query and the number of streaming units to use for the streaming job. The name property of the transformation is required when specifying this property in a PUT request. This property cannot be modify via a PATCH operation. You must use the PATCH API available for the individual transformation.
	Transformation TransformationResponsePtrInput
	// The type of the resource. Ex- Microsoft.Compute/virtualMachines or Microsoft.Storage/storageAccounts.
	Type pulumi.StringPtrInput
}

func (StreamingJobState) ElementType() reflect.Type {
	return reflect.TypeOf((*streamingJobState)(nil)).Elem()
}

type streamingJobArgs struct {
	// The cluster which streaming jobs will run on.
	Cluster *ClusterInfo `pulumi:"cluster"`
	// Controls certain runtime behaviors of the streaming job.
	CompatibilityLevel *string `pulumi:"compatibilityLevel"`
	// The data locale of the stream analytics job. Value should be the name of a supported .NET Culture from the set https://msdn.microsoft.com/en-us/library/system.globalization.culturetypes(v=vs.110).aspx. Defaults to 'en-US' if none specified.
	DataLocale *string `pulumi:"dataLocale"`
	// The maximum tolerable delay in seconds where events arriving late could be included.  Supported range is -1 to 1814399 (20.23:59:59 days) and -1 is used to specify wait indefinitely. If the property is absent, it is interpreted to have a value of -1.
	EventsLateArrivalMaxDelayInSeconds *int `pulumi:"eventsLateArrivalMaxDelayInSeconds"`
	// The maximum tolerable delay in seconds where out-of-order events can be adjusted to be back in order.
	EventsOutOfOrderMaxDelayInSeconds *int `pulumi:"eventsOutOfOrderMaxDelayInSeconds"`
	// Indicates the policy to apply to events that arrive out of order in the input event stream.
	EventsOutOfOrderPolicy *string `pulumi:"eventsOutOfOrderPolicy"`
	// The storage account where the custom code artifacts are located.
	Externals *External `pulumi:"externals"`
	// A list of one or more functions for the streaming job. The name property for each function is required when specifying this property in a PUT request. This property cannot be modify via a PATCH operation. You must use the PATCH API available for the individual transformation.
	Functions []FunctionType `pulumi:"functions"`
	// Describes the system-assigned managed identity assigned to this job that can be used to authenticate with inputs and outputs.
	Identity *Identity `pulumi:"identity"`
	// A list of one or more inputs to the streaming job. The name property for each input is required when specifying this property in a PUT request. This property cannot be modify via a PATCH operation. You must use the PATCH API available for the individual input.
	Inputs []InputType `pulumi:"inputs"`
	// The name of the streaming job.
	JobName string `pulumi:"jobName"`
	// The properties that are associated with an Azure Storage account with MSI
	JobStorageAccount *JobStorageAccount `pulumi:"jobStorageAccount"`
	// Describes the type of the job. Valid modes are `Cloud` and 'Edge'.
	JobType *string `pulumi:"jobType"`
	// The geo-location where the resource lives
	Location *string `pulumi:"location"`
	// Indicates the policy to apply to events that arrive at the output and cannot be written to the external storage due to being malformed (missing column values, column values of wrong type or size).
	OutputErrorPolicy *string `pulumi:"outputErrorPolicy"`
	// This property should only be utilized when it is desired that the job be started immediately upon creation. Value may be JobStartTime, CustomTime, or LastOutputEventTime to indicate whether the starting point of the output event stream should start whenever the job is started, start at a custom user time stamp specified via the outputStartTime property, or start from the last event output time.
	OutputStartMode *string `pulumi:"outputStartMode"`
	// Value is either an ISO-8601 formatted time stamp that indicates the starting point of the output event stream, or null to indicate that the output event stream will start whenever the streaming job is started. This property must have a value if outputStartMode is set to CustomTime.
	OutputStartTime *string `pulumi:"outputStartTime"`
	// A list of one or more outputs for the streaming job. The name property for each output is required when specifying this property in a PUT request. This property cannot be modify via a PATCH operation. You must use the PATCH API available for the individual output.
	Outputs []OutputType `pulumi:"outputs"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// Describes the SKU of the streaming job. Required on PUT (CreateOrReplace) requests.
	Sku *StreamingJobSku `pulumi:"sku"`
	// Resource tags.
	Tags map[string]string `pulumi:"tags"`
	// Indicates the query and the number of streaming units to use for the streaming job. The name property of the transformation is required when specifying this property in a PUT request. This property cannot be modify via a PATCH operation. You must use the PATCH API available for the individual transformation.
	Transformation *Transformation `pulumi:"transformation"`
}

// The set of arguments for constructing a StreamingJob resource.
type StreamingJobArgs struct {
	// The cluster which streaming jobs will run on.
	Cluster ClusterInfoPtrInput
	// Controls certain runtime behaviors of the streaming job.
	CompatibilityLevel pulumi.StringPtrInput
	// The data locale of the stream analytics job. Value should be the name of a supported .NET Culture from the set https://msdn.microsoft.com/en-us/library/system.globalization.culturetypes(v=vs.110).aspx. Defaults to 'en-US' if none specified.
	DataLocale pulumi.StringPtrInput
	// The maximum tolerable delay in seconds where events arriving late could be included.  Supported range is -1 to 1814399 (20.23:59:59 days) and -1 is used to specify wait indefinitely. If the property is absent, it is interpreted to have a value of -1.
	EventsLateArrivalMaxDelayInSeconds pulumi.IntPtrInput
	// The maximum tolerable delay in seconds where out-of-order events can be adjusted to be back in order.
	EventsOutOfOrderMaxDelayInSeconds pulumi.IntPtrInput
	// Indicates the policy to apply to events that arrive out of order in the input event stream.
	EventsOutOfOrderPolicy pulumi.StringPtrInput
	// The storage account where the custom code artifacts are located.
	Externals ExternalPtrInput
	// A list of one or more functions for the streaming job. The name property for each function is required when specifying this property in a PUT request. This property cannot be modify via a PATCH operation. You must use the PATCH API available for the individual transformation.
	Functions FunctionTypeArrayInput
	// Describes the system-assigned managed identity assigned to this job that can be used to authenticate with inputs and outputs.
	Identity IdentityPtrInput
	// A list of one or more inputs to the streaming job. The name property for each input is required when specifying this property in a PUT request. This property cannot be modify via a PATCH operation. You must use the PATCH API available for the individual input.
	Inputs InputTypeArrayInput
	// The name of the streaming job.
	JobName pulumi.StringInput
	// The properties that are associated with an Azure Storage account with MSI
	JobStorageAccount JobStorageAccountPtrInput
	// Describes the type of the job. Valid modes are `Cloud` and 'Edge'.
	JobType pulumi.StringPtrInput
	// The geo-location where the resource lives
	Location pulumi.StringPtrInput
	// Indicates the policy to apply to events that arrive at the output and cannot be written to the external storage due to being malformed (missing column values, column values of wrong type or size).
	OutputErrorPolicy pulumi.StringPtrInput
	// This property should only be utilized when it is desired that the job be started immediately upon creation. Value may be JobStartTime, CustomTime, or LastOutputEventTime to indicate whether the starting point of the output event stream should start whenever the job is started, start at a custom user time stamp specified via the outputStartTime property, or start from the last event output time.
	OutputStartMode pulumi.StringPtrInput
	// Value is either an ISO-8601 formatted time stamp that indicates the starting point of the output event stream, or null to indicate that the output event stream will start whenever the streaming job is started. This property must have a value if outputStartMode is set to CustomTime.
	OutputStartTime pulumi.StringPtrInput
	// A list of one or more outputs for the streaming job. The name property for each output is required when specifying this property in a PUT request. This property cannot be modify via a PATCH operation. You must use the PATCH API available for the individual output.
	Outputs OutputTypeArrayInput
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput
	// Describes the SKU of the streaming job. Required on PUT (CreateOrReplace) requests.
	Sku StreamingJobSkuPtrInput
	// Resource tags.
	Tags pulumi.StringMapInput
	// Indicates the query and the number of streaming units to use for the streaming job. The name property of the transformation is required when specifying this property in a PUT request. This property cannot be modify via a PATCH operation. You must use the PATCH API available for the individual transformation.
	Transformation TransformationPtrInput
}

func (StreamingJobArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*streamingJobArgs)(nil)).Elem()
}