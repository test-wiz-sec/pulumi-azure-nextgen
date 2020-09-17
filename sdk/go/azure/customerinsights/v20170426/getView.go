// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v20170426

import (
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

func LookupView(ctx *pulumi.Context, args *LookupViewArgs, opts ...pulumi.InvokeOption) (*LookupViewResult, error) {
	var rv LookupViewResult
	err := ctx.Invoke("azure-nextgen:customerinsights/v20170426:getView", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupViewArgs struct {
	// The name of the hub.
	HubName string `pulumi:"hubName"`
	// The name of the resource group.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The user ID. Use * to retrieve hub level view.
	UserId string `pulumi:"userId"`
	// The name of the view.
	ViewName string `pulumi:"viewName"`
}

// The view resource format.
type LookupViewResult struct {
	// Date time when view was last modified.
	Changed string `pulumi:"changed"`
	// Date time when view was created.
	Created string `pulumi:"created"`
	// View definition.
	Definition string `pulumi:"definition"`
	// Localized display name for the view.
	DisplayName map[string]string `pulumi:"displayName"`
	// Resource name.
	Name string `pulumi:"name"`
	// the hub name.
	TenantId string `pulumi:"tenantId"`
	// Resource type.
	Type string `pulumi:"type"`
	// the user ID.
	UserId *string `pulumi:"userId"`
	// Name of the view.
	ViewName string `pulumi:"viewName"`
}