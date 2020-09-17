// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v20150504preview

import (
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

func LookupZone(ctx *pulumi.Context, args *LookupZoneArgs, opts ...pulumi.InvokeOption) (*LookupZoneResult, error) {
	var rv LookupZoneResult
	err := ctx.Invoke("azure-nextgen:network/v20150504preview:getZone", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupZoneArgs struct {
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The name of the zone without a terminating dot.
	ZoneName string `pulumi:"zoneName"`
}

// Describes a DNS zone.
type LookupZoneResult struct {
	// Gets or sets the ETag of the zone that is being updated, as received from a Get operation.
	Etag *string `pulumi:"etag"`
	// The geo-location where the resource lives
	Location string `pulumi:"location"`
	// The name of the resource
	Name string `pulumi:"name"`
	// Gets or sets the properties of the zone.
	Properties ZonePropertiesResponse `pulumi:"properties"`
	// Resource tags.
	Tags map[string]string `pulumi:"tags"`
	// The type of the resource. Ex- Microsoft.Compute/virtualMachines or Microsoft.Storage/storageAccounts.
	Type string `pulumi:"type"`
}