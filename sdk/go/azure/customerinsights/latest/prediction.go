// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package latest

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// The prediction resource format.
type Prediction struct {
	pulumi.CustomResourceState

	// Whether do auto analyze.
	AutoAnalyze pulumi.BoolOutput `pulumi:"autoAnalyze"`
	// Description of the prediction.
	Description pulumi.StringMapOutput `pulumi:"description"`
	// Display name of the prediction.
	DisplayName pulumi.StringMapOutput `pulumi:"displayName"`
	// The prediction grades.
	Grades PredictionResponseGradesArrayOutput `pulumi:"grades"`
	// Interaction types involved in the prediction.
	InvolvedInteractionTypes pulumi.StringArrayOutput `pulumi:"involvedInteractionTypes"`
	// KPI types involved in the prediction.
	InvolvedKpiTypes pulumi.StringArrayOutput `pulumi:"involvedKpiTypes"`
	// Relationships involved in the prediction.
	InvolvedRelationships pulumi.StringArrayOutput `pulumi:"involvedRelationships"`
	// Definition of the link mapping of prediction.
	Mappings PredictionResponseMappingsOutput `pulumi:"mappings"`
	// Resource name.
	Name pulumi.StringOutput `pulumi:"name"`
	// Negative outcome expression.
	NegativeOutcomeExpression pulumi.StringOutput `pulumi:"negativeOutcomeExpression"`
	// Positive outcome expression.
	PositiveOutcomeExpression pulumi.StringOutput `pulumi:"positiveOutcomeExpression"`
	// Name of the prediction.
	PredictionName pulumi.StringPtrOutput `pulumi:"predictionName"`
	// Primary profile type.
	PrimaryProfileType pulumi.StringOutput `pulumi:"primaryProfileType"`
	// Provisioning state.
	ProvisioningState pulumi.StringOutput `pulumi:"provisioningState"`
	// Scope expression.
	ScopeExpression pulumi.StringOutput `pulumi:"scopeExpression"`
	// Score label.
	ScoreLabel pulumi.StringOutput `pulumi:"scoreLabel"`
	// System generated entities.
	SystemGeneratedEntities PredictionResponseSystemGeneratedEntitiesOutput `pulumi:"systemGeneratedEntities"`
	// The hub name.
	TenantId pulumi.StringOutput `pulumi:"tenantId"`
	// Resource type.
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewPrediction registers a new resource with the given unique name, arguments, and options.
func NewPrediction(ctx *pulumi.Context,
	name string, args *PredictionArgs, opts ...pulumi.ResourceOption) (*Prediction, error) {
	if args == nil || args.AutoAnalyze == nil {
		return nil, errors.New("missing required argument 'AutoAnalyze'")
	}
	if args == nil || args.HubName == nil {
		return nil, errors.New("missing required argument 'HubName'")
	}
	if args == nil || args.Mappings == nil {
		return nil, errors.New("missing required argument 'Mappings'")
	}
	if args == nil || args.NegativeOutcomeExpression == nil {
		return nil, errors.New("missing required argument 'NegativeOutcomeExpression'")
	}
	if args == nil || args.PositiveOutcomeExpression == nil {
		return nil, errors.New("missing required argument 'PositiveOutcomeExpression'")
	}
	if args == nil || args.PredictionName == nil {
		return nil, errors.New("missing required argument 'PredictionName'")
	}
	if args == nil || args.PrimaryProfileType == nil {
		return nil, errors.New("missing required argument 'PrimaryProfileType'")
	}
	if args == nil || args.ResourceGroupName == nil {
		return nil, errors.New("missing required argument 'ResourceGroupName'")
	}
	if args == nil || args.ScopeExpression == nil {
		return nil, errors.New("missing required argument 'ScopeExpression'")
	}
	if args == nil || args.ScoreLabel == nil {
		return nil, errors.New("missing required argument 'ScoreLabel'")
	}
	if args == nil {
		args = &PredictionArgs{}
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-nextgen:customerinsights/v20170426:Prediction"),
		},
	})
	opts = append(opts, aliases)
	var resource Prediction
	err := ctx.RegisterResource("azure-nextgen:customerinsights/latest:Prediction", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetPrediction gets an existing Prediction resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetPrediction(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *PredictionState, opts ...pulumi.ResourceOption) (*Prediction, error) {
	var resource Prediction
	err := ctx.ReadResource("azure-nextgen:customerinsights/latest:Prediction", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Prediction resources.
type predictionState struct {
	// Whether do auto analyze.
	AutoAnalyze *bool `pulumi:"autoAnalyze"`
	// Description of the prediction.
	Description map[string]string `pulumi:"description"`
	// Display name of the prediction.
	DisplayName map[string]string `pulumi:"displayName"`
	// The prediction grades.
	Grades []PredictionResponseGrades `pulumi:"grades"`
	// Interaction types involved in the prediction.
	InvolvedInteractionTypes []string `pulumi:"involvedInteractionTypes"`
	// KPI types involved in the prediction.
	InvolvedKpiTypes []string `pulumi:"involvedKpiTypes"`
	// Relationships involved in the prediction.
	InvolvedRelationships []string `pulumi:"involvedRelationships"`
	// Definition of the link mapping of prediction.
	Mappings *PredictionResponseMappings `pulumi:"mappings"`
	// Resource name.
	Name *string `pulumi:"name"`
	// Negative outcome expression.
	NegativeOutcomeExpression *string `pulumi:"negativeOutcomeExpression"`
	// Positive outcome expression.
	PositiveOutcomeExpression *string `pulumi:"positiveOutcomeExpression"`
	// Name of the prediction.
	PredictionName *string `pulumi:"predictionName"`
	// Primary profile type.
	PrimaryProfileType *string `pulumi:"primaryProfileType"`
	// Provisioning state.
	ProvisioningState *string `pulumi:"provisioningState"`
	// Scope expression.
	ScopeExpression *string `pulumi:"scopeExpression"`
	// Score label.
	ScoreLabel *string `pulumi:"scoreLabel"`
	// System generated entities.
	SystemGeneratedEntities *PredictionResponseSystemGeneratedEntities `pulumi:"systemGeneratedEntities"`
	// The hub name.
	TenantId *string `pulumi:"tenantId"`
	// Resource type.
	Type *string `pulumi:"type"`
}

type PredictionState struct {
	// Whether do auto analyze.
	AutoAnalyze pulumi.BoolPtrInput
	// Description of the prediction.
	Description pulumi.StringMapInput
	// Display name of the prediction.
	DisplayName pulumi.StringMapInput
	// The prediction grades.
	Grades PredictionResponseGradesArrayInput
	// Interaction types involved in the prediction.
	InvolvedInteractionTypes pulumi.StringArrayInput
	// KPI types involved in the prediction.
	InvolvedKpiTypes pulumi.StringArrayInput
	// Relationships involved in the prediction.
	InvolvedRelationships pulumi.StringArrayInput
	// Definition of the link mapping of prediction.
	Mappings PredictionResponseMappingsPtrInput
	// Resource name.
	Name pulumi.StringPtrInput
	// Negative outcome expression.
	NegativeOutcomeExpression pulumi.StringPtrInput
	// Positive outcome expression.
	PositiveOutcomeExpression pulumi.StringPtrInput
	// Name of the prediction.
	PredictionName pulumi.StringPtrInput
	// Primary profile type.
	PrimaryProfileType pulumi.StringPtrInput
	// Provisioning state.
	ProvisioningState pulumi.StringPtrInput
	// Scope expression.
	ScopeExpression pulumi.StringPtrInput
	// Score label.
	ScoreLabel pulumi.StringPtrInput
	// System generated entities.
	SystemGeneratedEntities PredictionResponseSystemGeneratedEntitiesPtrInput
	// The hub name.
	TenantId pulumi.StringPtrInput
	// Resource type.
	Type pulumi.StringPtrInput
}

func (PredictionState) ElementType() reflect.Type {
	return reflect.TypeOf((*predictionState)(nil)).Elem()
}

type predictionArgs struct {
	// Whether do auto analyze.
	AutoAnalyze bool `pulumi:"autoAnalyze"`
	// Description of the prediction.
	Description map[string]string `pulumi:"description"`
	// Display name of the prediction.
	DisplayName map[string]string `pulumi:"displayName"`
	// The prediction grades.
	Grades []PredictionGrades `pulumi:"grades"`
	// The name of the hub.
	HubName string `pulumi:"hubName"`
	// Interaction types involved in the prediction.
	InvolvedInteractionTypes []string `pulumi:"involvedInteractionTypes"`
	// KPI types involved in the prediction.
	InvolvedKpiTypes []string `pulumi:"involvedKpiTypes"`
	// Relationships involved in the prediction.
	InvolvedRelationships []string `pulumi:"involvedRelationships"`
	// Definition of the link mapping of prediction.
	Mappings PredictionMappings `pulumi:"mappings"`
	// Negative outcome expression.
	NegativeOutcomeExpression string `pulumi:"negativeOutcomeExpression"`
	// Positive outcome expression.
	PositiveOutcomeExpression string `pulumi:"positiveOutcomeExpression"`
	// Name of the prediction.
	PredictionName string `pulumi:"predictionName"`
	// Primary profile type.
	PrimaryProfileType string `pulumi:"primaryProfileType"`
	// The name of the resource group.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// Scope expression.
	ScopeExpression string `pulumi:"scopeExpression"`
	// Score label.
	ScoreLabel string `pulumi:"scoreLabel"`
}

// The set of arguments for constructing a Prediction resource.
type PredictionArgs struct {
	// Whether do auto analyze.
	AutoAnalyze pulumi.BoolInput
	// Description of the prediction.
	Description pulumi.StringMapInput
	// Display name of the prediction.
	DisplayName pulumi.StringMapInput
	// The prediction grades.
	Grades PredictionGradesArrayInput
	// The name of the hub.
	HubName pulumi.StringInput
	// Interaction types involved in the prediction.
	InvolvedInteractionTypes pulumi.StringArrayInput
	// KPI types involved in the prediction.
	InvolvedKpiTypes pulumi.StringArrayInput
	// Relationships involved in the prediction.
	InvolvedRelationships pulumi.StringArrayInput
	// Definition of the link mapping of prediction.
	Mappings PredictionMappingsInput
	// Negative outcome expression.
	NegativeOutcomeExpression pulumi.StringInput
	// Positive outcome expression.
	PositiveOutcomeExpression pulumi.StringInput
	// Name of the prediction.
	PredictionName pulumi.StringInput
	// Primary profile type.
	PrimaryProfileType pulumi.StringInput
	// The name of the resource group.
	ResourceGroupName pulumi.StringInput
	// Scope expression.
	ScopeExpression pulumi.StringInput
	// Score label.
	ScoreLabel pulumi.StringInput
}

func (PredictionArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*predictionArgs)(nil)).Elem()
}