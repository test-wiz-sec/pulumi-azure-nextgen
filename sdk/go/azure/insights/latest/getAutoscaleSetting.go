// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package latest

import (
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

func LookupAutoscaleSetting(ctx *pulumi.Context, args *LookupAutoscaleSettingArgs, opts ...pulumi.InvokeOption) (*LookupAutoscaleSettingResult, error) {
	var rv LookupAutoscaleSettingResult
	err := ctx.Invoke("azure-nextgen:insights/latest:getAutoscaleSetting", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupAutoscaleSettingArgs struct {
	// The autoscale setting name.
	AutoscaleSettingName string `pulumi:"autoscaleSettingName"`
	// The name of the resource group.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// The autoscale setting resource.
type LookupAutoscaleSettingResult struct {
	// the enabled flag. Specifies whether automatic scaling is enabled for the resource. The default value is 'true'.
	Enabled *bool `pulumi:"enabled"`
	// Resource location
	Location string `pulumi:"location"`
	// Azure resource name
	Name string `pulumi:"name"`
	// the collection of notifications.
	Notifications []AutoscaleNotificationResponse `pulumi:"notifications"`
	// the collection of automatic scaling profiles that specify different scaling parameters for different time periods. A maximum of 20 profiles can be specified.
	Profiles []AutoscaleProfileResponse `pulumi:"profiles"`
	// Resource tags
	Tags map[string]string `pulumi:"tags"`
	// the resource identifier of the resource that the autoscale setting should be added to.
	TargetResourceUri *string `pulumi:"targetResourceUri"`
	// Azure resource type
	Type string `pulumi:"type"`
}