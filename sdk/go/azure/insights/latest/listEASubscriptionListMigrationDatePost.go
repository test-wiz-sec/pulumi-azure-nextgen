// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package latest

import (
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

func ListEASubscriptionListMigrationDatePost(ctx *pulumi.Context, args *ListEASubscriptionListMigrationDatePostArgs, opts ...pulumi.InvokeOption) (*ListEASubscriptionListMigrationDatePostResult, error) {
	var rv ListEASubscriptionListMigrationDatePostResult
	err := ctx.Invoke("azure-nextgen:insights/latest:listEASubscriptionListMigrationDatePost", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type ListEASubscriptionListMigrationDatePostArgs struct {
}

// Subscription migrate date information properties
type ListEASubscriptionListMigrationDatePostResult struct {
	// Is subscription in the grand fatherable subscription list.
	IsGrandFatherableSubscription *bool `pulumi:"isGrandFatherableSubscription"`
	// Time to start using new pricing model.
	OptedInDate *string `pulumi:"optedInDate"`
}
