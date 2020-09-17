// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNextGen.Web.V20150801Preview.Outputs
{

    [OutputType]
    public sealed class ApiOAuthSettingsResponse
    {
        /// <summary>
        /// Resource provider client id
        /// </summary>
        public readonly string? ClientId;
        /// <summary>
        /// Client Secret needed for OAuth
        /// </summary>
        public readonly string? ClientSecret;
        /// <summary>
        /// OAuth parameters key is the name of parameter
        /// </summary>
        public readonly ImmutableDictionary<string, Outputs.ApiOAuthSettingsParameterResponse>? CustomParameters;
        /// <summary>
        /// Identity provider
        /// </summary>
        public readonly string? IdentityProvider;
        /// <summary>
        /// Read only properties for this oauth setting.
        /// </summary>
        public readonly ImmutableDictionary<string, object>? Properties;
        /// <summary>
        /// Url
        /// </summary>
        public readonly string? RedirectUrl;
        /// <summary>
        /// OAuth scopes
        /// </summary>
        public readonly ImmutableArray<string> Scopes;

        [OutputConstructor]
        private ApiOAuthSettingsResponse(
            string? clientId,

            string? clientSecret,

            ImmutableDictionary<string, Outputs.ApiOAuthSettingsParameterResponse>? customParameters,

            string? identityProvider,

            ImmutableDictionary<string, object>? properties,

            string? redirectUrl,

            ImmutableArray<string> scopes)
        {
            ClientId = clientId;
            ClientSecret = clientSecret;
            CustomParameters = customParameters;
            IdentityProvider = identityProvider;
            Properties = properties;
            RedirectUrl = redirectUrl;
            Scopes = scopes;
        }
    }
}