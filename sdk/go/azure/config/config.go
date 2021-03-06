// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package config

import (
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi/config"
)

func GetAuxiliaryTenantIds(ctx *pulumi.Context) string {
	return config.Get(ctx, "azure-nextgen:auxiliaryTenantIds")
}

// The password associated with the Client Certificate. For use when authenticating as a Service Principal using a Client Certificate
func GetClientCertificatePassword(ctx *pulumi.Context) string {
	return config.Get(ctx, "azure-nextgen:clientCertificatePassword")
}

// The path to the Client Certificate associated with the Service Principal for use when authenticating as a Service Principal using a Client Certificate.
func GetClientCertificatePath(ctx *pulumi.Context) string {
	return config.Get(ctx, "azure-nextgen:clientCertificatePath")
}

// The Client ID which should be used.
func GetClientId(ctx *pulumi.Context) string {
	return config.Get(ctx, "azure-nextgen:clientId")
}

// The Client Secret which should be used. For use When authenticating as a Service Principal using a Client Secret.
func GetClientSecret(ctx *pulumi.Context) string {
	return config.Get(ctx, "azure-nextgen:clientSecret")
}

// This will disable the Pulumi Partner ID which is used if a custom `partnerId` isn't specified.
func GetDisablePulumiPartnerId(ctx *pulumi.Context) bool {
	return config.GetBool(ctx, "azure-nextgen:disablePulumiPartnerId")
}

// The Cloud Environment which should be used. Possible values are public, usgovernment, german, and china. Defaults to public.
func GetEnvironment(ctx *pulumi.Context) string {
	return config.Get(ctx, "azure-nextgen:environment")
}

// The path to a custom endpoint for Managed Service Identity - in most circumstances this should be detected automatically.
func GetMsiEndpoint(ctx *pulumi.Context) string {
	return config.Get(ctx, "azure-nextgen:msiEndpoint")
}

// A GUID/UUID that is registered with Microsoft to facilitate partner resource usage attribution.
func GetPartnerId(ctx *pulumi.Context) string {
	return config.Get(ctx, "azure-nextgen:partnerId")
}

// The Subscription ID which should be used.
func GetSubscriptionId(ctx *pulumi.Context) string {
	return config.Get(ctx, "azure-nextgen:subscriptionId")
}

// The Tenant ID which should be used.
func GetTenantId(ctx *pulumi.Context) string {
	return config.Get(ctx, "azure-nextgen:tenantId")
}

// Allowed Managed Service Identity be used for Authentication.
func GetUseMsi(ctx *pulumi.Context) bool {
	return config.GetBool(ctx, "azure-nextgen:useMsi")
}
