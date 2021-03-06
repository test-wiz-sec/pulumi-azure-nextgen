// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v20180601preview

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Diagnostic details.
type Diagnostic struct {
	pulumi.CustomResourceState

	// Specifies for what type of messages sampling settings should not apply.
	AlwaysLog pulumi.StringPtrOutput `pulumi:"alwaysLog"`
	// Diagnostic settings for incoming/outgoing HTTP messages to the Backend
	Backend PipelineDiagnosticSettingsResponsePtrOutput `pulumi:"backend"`
	// Whether to process Correlation Headers coming to Api Management Service. Only applicable to Application Insights diagnostics. Default is true.
	EnableHttpCorrelationHeaders pulumi.BoolPtrOutput `pulumi:"enableHttpCorrelationHeaders"`
	// Diagnostic settings for incoming/outgoing HTTP messages to the Gateway.
	Frontend PipelineDiagnosticSettingsResponsePtrOutput `pulumi:"frontend"`
	// Resource Id of a target logger.
	LoggerId pulumi.StringOutput `pulumi:"loggerId"`
	// Resource name.
	Name pulumi.StringOutput `pulumi:"name"`
	// Sampling settings for Diagnostic.
	Sampling SamplingSettingsResponsePtrOutput `pulumi:"sampling"`
	// Resource type for API Management resource.
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewDiagnostic registers a new resource with the given unique name, arguments, and options.
func NewDiagnostic(ctx *pulumi.Context,
	name string, args *DiagnosticArgs, opts ...pulumi.ResourceOption) (*Diagnostic, error) {
	if args == nil || args.DiagnosticId == nil {
		return nil, errors.New("missing required argument 'DiagnosticId'")
	}
	if args == nil || args.LoggerId == nil {
		return nil, errors.New("missing required argument 'LoggerId'")
	}
	if args == nil || args.ResourceGroupName == nil {
		return nil, errors.New("missing required argument 'ResourceGroupName'")
	}
	if args == nil || args.ServiceName == nil {
		return nil, errors.New("missing required argument 'ServiceName'")
	}
	if args == nil {
		args = &DiagnosticArgs{}
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-nextgen:apimanagement/latest:Diagnostic"),
		},
		{
			Type: pulumi.String("azure-nextgen:apimanagement/v20170301:Diagnostic"),
		},
		{
			Type: pulumi.String("azure-nextgen:apimanagement/v20180101:Diagnostic"),
		},
		{
			Type: pulumi.String("azure-nextgen:apimanagement/v20190101:Diagnostic"),
		},
		{
			Type: pulumi.String("azure-nextgen:apimanagement/v20191201:Diagnostic"),
		},
		{
			Type: pulumi.String("azure-nextgen:apimanagement/v20191201preview:Diagnostic"),
		},
		{
			Type: pulumi.String("azure-nextgen:apimanagement/v20200601preview:Diagnostic"),
		},
	})
	opts = append(opts, aliases)
	var resource Diagnostic
	err := ctx.RegisterResource("azure-nextgen:apimanagement/v20180601preview:Diagnostic", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetDiagnostic gets an existing Diagnostic resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetDiagnostic(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *DiagnosticState, opts ...pulumi.ResourceOption) (*Diagnostic, error) {
	var resource Diagnostic
	err := ctx.ReadResource("azure-nextgen:apimanagement/v20180601preview:Diagnostic", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Diagnostic resources.
type diagnosticState struct {
	// Specifies for what type of messages sampling settings should not apply.
	AlwaysLog *string `pulumi:"alwaysLog"`
	// Diagnostic settings for incoming/outgoing HTTP messages to the Backend
	Backend *PipelineDiagnosticSettingsResponse `pulumi:"backend"`
	// Whether to process Correlation Headers coming to Api Management Service. Only applicable to Application Insights diagnostics. Default is true.
	EnableHttpCorrelationHeaders *bool `pulumi:"enableHttpCorrelationHeaders"`
	// Diagnostic settings for incoming/outgoing HTTP messages to the Gateway.
	Frontend *PipelineDiagnosticSettingsResponse `pulumi:"frontend"`
	// Resource Id of a target logger.
	LoggerId *string `pulumi:"loggerId"`
	// Resource name.
	Name *string `pulumi:"name"`
	// Sampling settings for Diagnostic.
	Sampling *SamplingSettingsResponse `pulumi:"sampling"`
	// Resource type for API Management resource.
	Type *string `pulumi:"type"`
}

type DiagnosticState struct {
	// Specifies for what type of messages sampling settings should not apply.
	AlwaysLog pulumi.StringPtrInput
	// Diagnostic settings for incoming/outgoing HTTP messages to the Backend
	Backend PipelineDiagnosticSettingsResponsePtrInput
	// Whether to process Correlation Headers coming to Api Management Service. Only applicable to Application Insights diagnostics. Default is true.
	EnableHttpCorrelationHeaders pulumi.BoolPtrInput
	// Diagnostic settings for incoming/outgoing HTTP messages to the Gateway.
	Frontend PipelineDiagnosticSettingsResponsePtrInput
	// Resource Id of a target logger.
	LoggerId pulumi.StringPtrInput
	// Resource name.
	Name pulumi.StringPtrInput
	// Sampling settings for Diagnostic.
	Sampling SamplingSettingsResponsePtrInput
	// Resource type for API Management resource.
	Type pulumi.StringPtrInput
}

func (DiagnosticState) ElementType() reflect.Type {
	return reflect.TypeOf((*diagnosticState)(nil)).Elem()
}

type diagnosticArgs struct {
	// Specifies for what type of messages sampling settings should not apply.
	AlwaysLog *string `pulumi:"alwaysLog"`
	// Diagnostic settings for incoming/outgoing HTTP messages to the Backend
	Backend *PipelineDiagnosticSettings `pulumi:"backend"`
	// Diagnostic identifier. Must be unique in the current API Management service instance.
	DiagnosticId string `pulumi:"diagnosticId"`
	// Whether to process Correlation Headers coming to Api Management Service. Only applicable to Application Insights diagnostics. Default is true.
	EnableHttpCorrelationHeaders *bool `pulumi:"enableHttpCorrelationHeaders"`
	// Diagnostic settings for incoming/outgoing HTTP messages to the Gateway.
	Frontend *PipelineDiagnosticSettings `pulumi:"frontend"`
	// Resource Id of a target logger.
	LoggerId string `pulumi:"loggerId"`
	// The name of the resource group.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// Sampling settings for Diagnostic.
	Sampling *SamplingSettings `pulumi:"sampling"`
	// The name of the API Management service.
	ServiceName string `pulumi:"serviceName"`
}

// The set of arguments for constructing a Diagnostic resource.
type DiagnosticArgs struct {
	// Specifies for what type of messages sampling settings should not apply.
	AlwaysLog pulumi.StringPtrInput
	// Diagnostic settings for incoming/outgoing HTTP messages to the Backend
	Backend PipelineDiagnosticSettingsPtrInput
	// Diagnostic identifier. Must be unique in the current API Management service instance.
	DiagnosticId pulumi.StringInput
	// Whether to process Correlation Headers coming to Api Management Service. Only applicable to Application Insights diagnostics. Default is true.
	EnableHttpCorrelationHeaders pulumi.BoolPtrInput
	// Diagnostic settings for incoming/outgoing HTTP messages to the Gateway.
	Frontend PipelineDiagnosticSettingsPtrInput
	// Resource Id of a target logger.
	LoggerId pulumi.StringInput
	// The name of the resource group.
	ResourceGroupName pulumi.StringInput
	// Sampling settings for Diagnostic.
	Sampling SamplingSettingsPtrInput
	// The name of the API Management service.
	ServiceName pulumi.StringInput
}

func (DiagnosticArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*diagnosticArgs)(nil)).Elem()
}
