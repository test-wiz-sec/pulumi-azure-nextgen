// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v20180401

import (
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

func LookupRouteFilter(ctx *pulumi.Context, args *LookupRouteFilterArgs, opts ...pulumi.InvokeOption) (*LookupRouteFilterResult, error) {
	var rv LookupRouteFilterResult
	err := ctx.Invoke("azure-nextgen:network/v20180401:getRouteFilter", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupRouteFilterArgs struct {
	// Expands referenced express route bgp peering resources.
	Expand *string `pulumi:"expand"`
	// The name of the resource group.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The name of the route filter.
	RouteFilterName string `pulumi:"routeFilterName"`
}

// Route Filter Resource.
type LookupRouteFilterResult struct {
	// Gets a unique read-only string that changes whenever the resource is updated.
	Etag string `pulumi:"etag"`
	// Resource location.
	Location string `pulumi:"location"`
	// Resource name.
	Name string `pulumi:"name"`
	// A collection of references to express route circuit peerings.
	Peerings []ExpressRouteCircuitPeeringResponse `pulumi:"peerings"`
	// The provisioning state of the resource. Possible values are: 'Updating', 'Deleting', 'Succeeded' and 'Failed'.
	ProvisioningState string `pulumi:"provisioningState"`
	// Collection of RouteFilterRules contained within a route filter.
	Rules []RouteFilterRuleResponse `pulumi:"rules"`
	// Resource tags.
	Tags map[string]string `pulumi:"tags"`
	// Resource type.
	Type string `pulumi:"type"`
}