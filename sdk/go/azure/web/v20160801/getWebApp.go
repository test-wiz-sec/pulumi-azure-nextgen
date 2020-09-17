// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v20160801

import (
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

func LookupWebApp(ctx *pulumi.Context, args *LookupWebAppArgs, opts ...pulumi.InvokeOption) (*LookupWebAppResult, error) {
	var rv LookupWebAppResult
	err := ctx.Invoke("azure-nextgen:web/v20160801:getWebApp", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupWebAppArgs struct {
	// Name of the app.
	Name string `pulumi:"name"`
	// Name of the resource group to which the resource belongs.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// A web app, a mobile app backend, or an API app.
type LookupWebAppResult struct {
	// Management information availability state for the app.
	AvailabilityState string `pulumi:"availabilityState"`
	// <code>true</code> to enable client affinity; <code>false</code> to stop sending session affinity cookies, which route client requests in the same session to the same instance. Default is <code>true</code>.
	ClientAffinityEnabled *bool `pulumi:"clientAffinityEnabled"`
	// <code>true</code> to enable client certificate authentication (TLS mutual authentication); otherwise, <code>false</code>. Default is <code>false</code>.
	ClientCertEnabled *bool `pulumi:"clientCertEnabled"`
	// If specified during app creation, the app is cloned from a source app.
	CloningInfo *CloningInfoResponse `pulumi:"cloningInfo"`
	// Size of the function container.
	ContainerSize *int `pulumi:"containerSize"`
	// Maximum allowed daily memory-time quota (applicable on dynamic apps only).
	DailyMemoryTimeQuota *int `pulumi:"dailyMemoryTimeQuota"`
	// Default hostname of the app. Read-only.
	DefaultHostName string `pulumi:"defaultHostName"`
	// <code>true</code> if the app is enabled; otherwise, <code>false</code>. Setting this value to false disables the app (takes the app offline).
	Enabled *bool `pulumi:"enabled"`
	// Enabled hostnames for the app.Hostnames need to be assigned (see HostNames) AND enabled. Otherwise,
	// the app is not served on those hostnames.
	EnabledHostNames []string `pulumi:"enabledHostNames"`
	// Hostname SSL states are used to manage the SSL bindings for app's hostnames.
	HostNameSslStates []HostNameSslStateResponse `pulumi:"hostNameSslStates"`
	// Hostnames associated with the app.
	HostNames []string `pulumi:"hostNames"`
	// <code>true</code> to disable the public hostnames of the app; otherwise, <code>false</code>.
	//  If <code>true</code>, the app is only accessible via API management process.
	HostNamesDisabled *bool `pulumi:"hostNamesDisabled"`
	// App Service Environment to use for the app.
	HostingEnvironmentProfile *HostingEnvironmentProfileResponse `pulumi:"hostingEnvironmentProfile"`
	// HttpsOnly: configures a web site to accept only https requests. Issues redirect for
	// http requests
	HttpsOnly *bool `pulumi:"httpsOnly"`
	// Managed service identity.
	Identity *ManagedServiceIdentityResponse `pulumi:"identity"`
	// <code>true</code> if the app is a default container; otherwise, <code>false</code>.
	IsDefaultContainer bool `pulumi:"isDefaultContainer"`
	// Kind of resource.
	Kind *string `pulumi:"kind"`
	// Last time the app was modified, in UTC. Read-only.
	LastModifiedTimeUtc string `pulumi:"lastModifiedTimeUtc"`
	// Resource Location.
	Location string `pulumi:"location"`
	// Maximum number of workers.
	// This only applies to Functions container.
	MaxNumberOfWorkers int `pulumi:"maxNumberOfWorkers"`
	// Resource Name.
	Name string `pulumi:"name"`
	// List of IP addresses that the app uses for outbound connections (e.g. database access). Includes VIPs from tenants that site can be hosted with current settings. Read-only.
	OutboundIpAddresses string `pulumi:"outboundIpAddresses"`
	// List of IP addresses that the app uses for outbound connections (e.g. database access). Includes VIPs from all tenants. Read-only.
	PossibleOutboundIpAddresses string `pulumi:"possibleOutboundIpAddresses"`
	// Name of the repository site.
	RepositorySiteName string `pulumi:"repositorySiteName"`
	// <code>true</code> if reserved; otherwise, <code>false</code>.
	Reserved *bool `pulumi:"reserved"`
	// Name of the resource group the app belongs to. Read-only.
	ResourceGroup string `pulumi:"resourceGroup"`
	// <code>true</code> to stop SCM (KUDU) site when the app is stopped; otherwise, <code>false</code>. The default is <code>false</code>.
	ScmSiteAlsoStopped *bool `pulumi:"scmSiteAlsoStopped"`
	// Resource ID of the associated App Service plan, formatted as: "/subscriptions/{subscriptionID}/resourceGroups/{groupName}/providers/Microsoft.Web/serverfarms/{appServicePlanName}".
	ServerFarmId *string `pulumi:"serverFarmId"`
	// Configuration of the app.
	SiteConfig *SiteConfigResponse `pulumi:"siteConfig"`
	// Status of the last deployment slot swap operation.
	SlotSwapStatus SlotSwapStatusResponse `pulumi:"slotSwapStatus"`
	// If specified during app creation, the app is created from a previous snapshot.
	SnapshotInfo *SnapshotRecoveryRequestResponse `pulumi:"snapshotInfo"`
	// Current state of the app.
	State string `pulumi:"state"`
	// App suspended till in case memory-time quota is exceeded.
	SuspendedTill string `pulumi:"suspendedTill"`
	// Resource tags.
	Tags map[string]string `pulumi:"tags"`
	// Specifies which deployment slot this app will swap into. Read-only.
	TargetSwapSlot string `pulumi:"targetSwapSlot"`
	// Azure Traffic Manager hostnames associated with the app. Read-only.
	TrafficManagerHostNames []string `pulumi:"trafficManagerHostNames"`
	// Resource type.
	Type string `pulumi:"type"`
	// State indicating whether the app has exceeded its quota usage. Read-only.
	UsageState string `pulumi:"usageState"`
}