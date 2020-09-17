// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v20170821preview

import (
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

func ListIotDpsResourceKeys(ctx *pulumi.Context, args *ListIotDpsResourceKeysArgs, opts ...pulumi.InvokeOption) (*ListIotDpsResourceKeysResult, error) {
	var rv ListIotDpsResourceKeysResult
	err := ctx.Invoke("azure-nextgen:devices/v20170821preview:listIotDpsResourceKeys", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type ListIotDpsResourceKeysArgs struct {
	// The provisioning service name to get the shared access keys for.
	ProvisioningServiceName string `pulumi:"provisioningServiceName"`
	// resource group name
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// List of shared access keys.
type ListIotDpsResourceKeysResult struct {
	NextLink string                                                                  `pulumi:"nextLink"`
	Value    []SharedAccessSignatureAuthorizationRuleAccessRightsDescriptionResponse `pulumi:"value"`
}