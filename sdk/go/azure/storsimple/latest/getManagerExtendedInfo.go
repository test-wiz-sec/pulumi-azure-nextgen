// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package latest

import (
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

func LookupManagerExtendedInfo(ctx *pulumi.Context, args *LookupManagerExtendedInfoArgs, opts ...pulumi.InvokeOption) (*LookupManagerExtendedInfoResult, error) {
	var rv LookupManagerExtendedInfoResult
	err := ctx.Invoke("azure-nextgen:storsimple/latest:getManagerExtendedInfo", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupManagerExtendedInfoArgs struct {
	// The manager name
	ManagerName string `pulumi:"managerName"`
	// The resource group name
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// The extended info of the manager.
type LookupManagerExtendedInfoResult struct {
	// Represents the encryption algorithm used to encrypt the keys. None - if Key is saved in plain text format. Algorithm name - if key is encrypted
	Algorithm string `pulumi:"algorithm"`
	// Represents the CEK of the resource.
	EncryptionKey *string `pulumi:"encryptionKey"`
	// Represents the Cert thumbprint that was used to encrypt the CEK.
	EncryptionKeyThumbprint *string `pulumi:"encryptionKeyThumbprint"`
	// The etag of the resource.
	Etag *string `pulumi:"etag"`
	// Represents the CIK of the resource.
	IntegrityKey string `pulumi:"integrityKey"`
	// The Kind of the object. Currently only Series8000 is supported
	Kind *string `pulumi:"kind"`
	// The name of the object.
	Name string `pulumi:"name"`
	// Represents the portal thumbprint which can be used optionally to encrypt the entire data before storing it.
	PortalCertificateThumbprint *string `pulumi:"portalCertificateThumbprint"`
	// The hierarchical type of the object.
	Type string `pulumi:"type"`
	// The version of the extended info being persisted.
	Version *string `pulumi:"version"`
}
