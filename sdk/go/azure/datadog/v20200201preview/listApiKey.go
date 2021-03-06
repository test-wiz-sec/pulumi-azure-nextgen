// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v20200201preview

import (
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

func ListApiKey(ctx *pulumi.Context, args *ListApiKeyArgs, opts ...pulumi.InvokeOption) (*ListApiKeyResult, error) {
	var rv ListApiKeyResult
	err := ctx.Invoke("azure-nextgen:datadog/v20200201preview:listApiKey", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type ListApiKeyArgs struct {
	// Monitor resource name
	MonitorName string `pulumi:"monitorName"`
	// The name of the resource group to which the Datadog resource belongs.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// Response of a list operation.
type ListApiKeyResult struct {
	// Link to the next set of results, if any.
	NextLink *string `pulumi:"nextLink"`
	// Results of a list operation.
	Value []DatadogApiKeyResponse `pulumi:"value"`
}
