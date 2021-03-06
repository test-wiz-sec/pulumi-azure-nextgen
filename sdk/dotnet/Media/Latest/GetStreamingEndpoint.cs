// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNextGen.Media.Latest
{
    public static class GetStreamingEndpoint
    {
        public static Task<GetStreamingEndpointResult> InvokeAsync(GetStreamingEndpointArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetStreamingEndpointResult>("azure-nextgen:media/latest:getStreamingEndpoint", args ?? new GetStreamingEndpointArgs(), options.WithVersion());
    }


    public sealed class GetStreamingEndpointArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// The Media Services account name.
        /// </summary>
        [Input("accountName", required: true)]
        public string AccountName { get; set; } = null!;

        /// <summary>
        /// The name of the resource group within the Azure subscription.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public string ResourceGroupName { get; set; } = null!;

        /// <summary>
        /// The name of the streaming endpoint, maximum length is 24.
        /// </summary>
        [Input("streamingEndpointName", required: true)]
        public string StreamingEndpointName { get; set; } = null!;

        public GetStreamingEndpointArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetStreamingEndpointResult
    {
        /// <summary>
        /// The access control definition of the streaming endpoint.
        /// </summary>
        public readonly Outputs.StreamingEndpointAccessControlResponse? AccessControl;
        /// <summary>
        /// This feature is deprecated, do not set a value for this property.
        /// </summary>
        public readonly string? AvailabilitySetName;
        /// <summary>
        /// The CDN enabled flag.
        /// </summary>
        public readonly bool? CdnEnabled;
        /// <summary>
        /// The CDN profile name.
        /// </summary>
        public readonly string? CdnProfile;
        /// <summary>
        /// The CDN provider name.
        /// </summary>
        public readonly string? CdnProvider;
        /// <summary>
        /// The exact time the streaming endpoint was created.
        /// </summary>
        public readonly string Created;
        /// <summary>
        /// The streaming endpoint access policies.
        /// </summary>
        public readonly Outputs.CrossSiteAccessPoliciesResponse? CrossSiteAccessPolicies;
        /// <summary>
        /// The custom host names of the streaming endpoint
        /// </summary>
        public readonly ImmutableArray<string> CustomHostNames;
        /// <summary>
        /// The streaming endpoint description.
        /// </summary>
        public readonly string? Description;
        /// <summary>
        /// The free trial expiration time.
        /// </summary>
        public readonly string FreeTrialEndTime;
        /// <summary>
        /// The streaming endpoint host name.
        /// </summary>
        public readonly string HostName;
        /// <summary>
        /// The exact time the streaming endpoint was last modified.
        /// </summary>
        public readonly string LastModified;
        /// <summary>
        /// The geo-location where the resource lives
        /// </summary>
        public readonly string Location;
        /// <summary>
        /// Max cache age
        /// </summary>
        public readonly int? MaxCacheAge;
        /// <summary>
        /// The name of the resource
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// The provisioning state of the streaming endpoint.
        /// </summary>
        public readonly string ProvisioningState;
        /// <summary>
        /// The resource state of the streaming endpoint.
        /// </summary>
        public readonly string ResourceState;
        /// <summary>
        /// The number of scale units. Use the Scale operation to adjust this value.
        /// </summary>
        public readonly int ScaleUnits;
        /// <summary>
        /// Resource tags.
        /// </summary>
        public readonly ImmutableDictionary<string, string>? Tags;
        /// <summary>
        /// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
        /// </summary>
        public readonly string Type;

        [OutputConstructor]
        private GetStreamingEndpointResult(
            Outputs.StreamingEndpointAccessControlResponse? accessControl,

            string? availabilitySetName,

            bool? cdnEnabled,

            string? cdnProfile,

            string? cdnProvider,

            string created,

            Outputs.CrossSiteAccessPoliciesResponse? crossSiteAccessPolicies,

            ImmutableArray<string> customHostNames,

            string? description,

            string freeTrialEndTime,

            string hostName,

            string lastModified,

            string location,

            int? maxCacheAge,

            string name,

            string provisioningState,

            string resourceState,

            int scaleUnits,

            ImmutableDictionary<string, string>? tags,

            string type)
        {
            AccessControl = accessControl;
            AvailabilitySetName = availabilitySetName;
            CdnEnabled = cdnEnabled;
            CdnProfile = cdnProfile;
            CdnProvider = cdnProvider;
            Created = created;
            CrossSiteAccessPolicies = crossSiteAccessPolicies;
            CustomHostNames = customHostNames;
            Description = description;
            FreeTrialEndTime = freeTrialEndTime;
            HostName = hostName;
            LastModified = lastModified;
            Location = location;
            MaxCacheAge = maxCacheAge;
            Name = name;
            ProvisioningState = provisioningState;
            ResourceState = resourceState;
            ScaleUnits = scaleUnits;
            Tags = tags;
            Type = type;
        }
    }
}
