// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v20160801

import (
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

func ListWebAppConnectionStrings(ctx *pulumi.Context, args *ListWebAppConnectionStringsArgs, opts ...pulumi.InvokeOption) (*ListWebAppConnectionStringsResult, error) {
	var rv ListWebAppConnectionStringsResult
	err := ctx.Invoke("azure-nextgen:web/v20160801:listWebAppConnectionStrings", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type ListWebAppConnectionStringsArgs struct {
	// Name of the app.
	Name string `pulumi:"name"`
	// Name of the resource group to which the resource belongs.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// String dictionary resource.
type ListWebAppConnectionStringsResult struct {
	// Kind of resource.
	Kind *string `pulumi:"kind"`
	// Resource Name.
	Name string `pulumi:"name"`
	// Resource type.
	Type string `pulumi:"type"`
}