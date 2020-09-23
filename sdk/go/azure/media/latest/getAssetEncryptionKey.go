// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package latest

import (
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

func GetAssetEncryptionKey(ctx *pulumi.Context, args *GetAssetEncryptionKeyArgs, opts ...pulumi.InvokeOption) (*GetAssetEncryptionKeyResult, error) {
	var rv GetAssetEncryptionKeyResult
	err := ctx.Invoke("azure-nextgen:media/latest:getAssetEncryptionKey", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type GetAssetEncryptionKeyArgs struct {
	// The Media Services account name.
	AccountName string `pulumi:"accountName"`
	// The Asset name.
	AssetName string `pulumi:"assetName"`
	// The name of the resource group within the Azure subscription.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// Data needed to decrypt asset files encrypted with legacy storage encryption.
type GetAssetEncryptionKeyResult struct {
	// Asset File encryption metadata.
	AssetFileEncryptionMetadata []AssetFileEncryptionMetadataResponse `pulumi:"assetFileEncryptionMetadata"`
	// The Asset File storage encryption key.
	Key *string `pulumi:"key"`
}