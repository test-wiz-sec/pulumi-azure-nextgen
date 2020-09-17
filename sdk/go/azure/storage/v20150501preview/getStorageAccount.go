// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v20150501preview

import (
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

func LookupStorageAccount(ctx *pulumi.Context, args *LookupStorageAccountArgs, opts ...pulumi.InvokeOption) (*LookupStorageAccountResult, error) {
	var rv LookupStorageAccountResult
	err := ctx.Invoke("azure-nextgen:storage/v20150501preview:getStorageAccount", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupStorageAccountArgs struct {
	// The name of the storage account within the specified resource group. Storage account names must be between 3 and 24 characters in length and use numbers and lower-case letters only.
	AccountName string `pulumi:"accountName"`
	// The name of the resource group within the user's subscription.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// The storage account.
type LookupStorageAccountResult struct {
	// Gets the type of the storage account.
	AccountType *string `pulumi:"accountType"`
	// Gets the creation date and time of the storage account in UTC.
	CreationTime *string `pulumi:"creationTime"`
	// Gets the user assigned custom domain assigned to this storage account.
	CustomDomain *CustomDomainResponse `pulumi:"customDomain"`
	// Gets the timestamp of the most recent instance of a failover to the secondary location. Only the most recent timestamp is retained. This element is not returned if there has never been a failover instance. Only available if the accountType is StandardGRS or StandardRAGRS.
	LastGeoFailoverTime *string `pulumi:"lastGeoFailoverTime"`
	// Resource location
	Location string `pulumi:"location"`
	// Resource name
	Name string `pulumi:"name"`
	// Gets the URLs that are used to perform a retrieval of a public blob, queue or table object.Note that StandardZRS and PremiumLRS accounts only return the blob endpoint.
	PrimaryEndpoints *EndpointsResponse `pulumi:"primaryEndpoints"`
	// Gets the location of the primary for the storage account.
	PrimaryLocation *string `pulumi:"primaryLocation"`
	// Gets the status of the storage account at the time the operation was called.
	ProvisioningState *string `pulumi:"provisioningState"`
	// Gets the URLs that are used to perform a retrieval of a public blob, queue or table object from the secondary location of the storage account. Only available if the accountType is StandardRAGRS.
	SecondaryEndpoints *EndpointsResponse `pulumi:"secondaryEndpoints"`
	// Gets the location of the geo replicated secondary for the storage account. Only available if the accountType is StandardGRS or StandardRAGRS.
	SecondaryLocation *string `pulumi:"secondaryLocation"`
	// Gets the status indicating whether the primary location of the storage account is available or unavailable.
	StatusOfPrimary *string `pulumi:"statusOfPrimary"`
	// Gets the status indicating whether the secondary location of the storage account is available or unavailable. Only available if the accountType is StandardGRS or StandardRAGRS.
	StatusOfSecondary *string `pulumi:"statusOfSecondary"`
	// Resource tags
	Tags map[string]string `pulumi:"tags"`
	// Resource type
	Type string `pulumi:"type"`
}