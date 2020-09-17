// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v20200301preview

import (
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

func LookupCluster(ctx *pulumi.Context, args *LookupClusterArgs, opts ...pulumi.InvokeOption) (*LookupClusterResult, error) {
	var rv LookupClusterResult
	err := ctx.Invoke("azure-nextgen:azurestackhci/v20200301preview:getCluster", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupClusterArgs struct {
	// The name of the cluster.
	ClusterName string `pulumi:"clusterName"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// Cluster details.
type LookupClusterResult struct {
	// App id of cluster AAD identity.
	AadClientId string `pulumi:"aadClientId"`
	// Tenant id of cluster AAD identity.
	AadTenantId string `pulumi:"aadTenantId"`
	// Type of billing applied to the resource.
	BillingModel string `pulumi:"billingModel"`
	// Unique, immutable resource id.
	CloudId string `pulumi:"cloudId"`
	// Most recent billing meter timestamp.
	LastBillingTimestamp string `pulumi:"lastBillingTimestamp"`
	// Most recent cluster sync timestamp.
	LastSyncTimestamp string `pulumi:"lastSyncTimestamp"`
	// The geo-location where the resource lives
	Location string `pulumi:"location"`
	// The name of the resource
	Name string `pulumi:"name"`
	// Provisioning state.
	ProvisioningState string `pulumi:"provisioningState"`
	// First cluster sync timestamp.
	RegistrationTimestamp string `pulumi:"registrationTimestamp"`
	// Properties reported by cluster agent.
	ReportedProperties *ClusterReportedPropertiesResponse `pulumi:"reportedProperties"`
	// Status of the cluster agent.
	Status string `pulumi:"status"`
	// Resource tags.
	Tags map[string]string `pulumi:"tags"`
	// Number of days remaining in the trial period.
	TrialDaysRemaining float64 `pulumi:"trialDaysRemaining"`
	// The type of the resource. Ex- Microsoft.Compute/virtualMachines or Microsoft.Storage/storageAccounts.
	Type string `pulumi:"type"`
}