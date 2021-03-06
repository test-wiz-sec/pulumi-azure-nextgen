// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package latest

import (
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

func ListWebAppMetadata(ctx *pulumi.Context, args *ListWebAppMetadataArgs, opts ...pulumi.InvokeOption) (*ListWebAppMetadataResult, error) {
	var rv ListWebAppMetadataResult
	err := ctx.Invoke("azure-nextgen:web/latest:listWebAppMetadata", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type ListWebAppMetadataArgs struct {
	// Name of the app.
	Name string `pulumi:"name"`
	// Name of the resource group to which the resource belongs.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// String dictionary resource.
type ListWebAppMetadataResult struct {
	// Kind of resource.
	Kind *string `pulumi:"kind"`
	// Resource Name.
	Name string `pulumi:"name"`
	// Settings.
	Properties map[string]string `pulumi:"properties"`
	// Resource type.
	Type string `pulumi:"type"`
}
