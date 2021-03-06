// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v20171103preview

import (
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

func LookupSapMonitor(ctx *pulumi.Context, args *LookupSapMonitorArgs, opts ...pulumi.InvokeOption) (*LookupSapMonitorResult, error) {
	var rv LookupSapMonitorResult
	err := ctx.Invoke("azure-nextgen:hanaonazure/v20171103preview:getSapMonitor", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupSapMonitorArgs struct {
	// Name of the resource group.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// Name of the SAP monitor resource.
	SapMonitorName string `pulumi:"sapMonitorName"`
}

// SAP monitor info on Azure (ARM properties and SAP monitor properties)
type LookupSapMonitorResult struct {
	// The value indicating whether to send analytics to Microsoft
	EnableCustomerAnalytics *bool `pulumi:"enableCustomerAnalytics"`
	// MSI ID passed by customer which has access to customer's KeyVault and to be assigned to the Collector VM.
	HanaDbCredentialsMsiId *string `pulumi:"hanaDbCredentialsMsiId"`
	// Database name of the HANA instance.
	HanaDbName *string `pulumi:"hanaDbName"`
	// Database password of the HANA instance.
	HanaDbPassword *string `pulumi:"hanaDbPassword"`
	// KeyVault URL link to the password for the HANA database.
	HanaDbPasswordKeyVaultUrl *string `pulumi:"hanaDbPasswordKeyVaultUrl"`
	// Database port of the HANA instance.
	HanaDbSqlPort *int `pulumi:"hanaDbSqlPort"`
	// Database username of the HANA instance.
	HanaDbUsername *string `pulumi:"hanaDbUsername"`
	// Hostname of the HANA instance.
	HanaHostname *string `pulumi:"hanaHostname"`
	// Specifies the SAP monitor unique ID.
	HanaSubnet *string `pulumi:"hanaSubnet"`
	// Key Vault ID containing customer's HANA credentials.
	KeyVaultId *string `pulumi:"keyVaultId"`
	// Resource location
	Location *string `pulumi:"location"`
	// The ARM ID of the Log Analytics Workspace that is used for monitoring
	LogAnalyticsWorkspaceArmId *string `pulumi:"logAnalyticsWorkspaceArmId"`
	// The workspace ID of the log analytics workspace to be used for monitoring
	LogAnalyticsWorkspaceId *string `pulumi:"logAnalyticsWorkspaceId"`
	// The shared key of the log analytics workspace that is used for monitoring
	LogAnalyticsWorkspaceSharedKey *string `pulumi:"logAnalyticsWorkspaceSharedKey"`
	// The name of the resource group the SAP Monitor resources get deployed into.
	ManagedResourceGroupName string `pulumi:"managedResourceGroupName"`
	// Resource name
	Name string `pulumi:"name"`
	// State of provisioning of the HanaInstance
	ProvisioningState string `pulumi:"provisioningState"`
	// Resource tags
	Tags map[string]string `pulumi:"tags"`
	// Resource type
	Type string `pulumi:"type"`
}
