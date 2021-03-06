// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNextGen.Media.V20200501.Outputs
{

    [OutputType]
    public sealed class ContentKeyPolicyTokenRestrictionResponse
    {
        /// <summary>
        /// A list of alternative verification keys.
        /// </summary>
        public readonly ImmutableArray<Union<Outputs.ContentKeyPolicyRsaTokenKeyResponse, Union<Outputs.ContentKeyPolicySymmetricTokenKeyResponse, Outputs.ContentKeyPolicyX509CertificateTokenKeyResponse>>> AlternateVerificationKeys;
        /// <summary>
        /// The audience for the token.
        /// </summary>
        public readonly string Audience;
        /// <summary>
        /// The token issuer.
        /// </summary>
        public readonly string Issuer;
        /// <summary>
        /// The discriminator for derived types.
        /// </summary>
        public readonly string OdataType;
        /// <summary>
        /// The OpenID connect discovery document.
        /// </summary>
        public readonly string? OpenIdConnectDiscoveryDocument;
        /// <summary>
        /// The primary verification key.
        /// </summary>
        public readonly Union<Outputs.ContentKeyPolicyRsaTokenKeyResponse, Union<Outputs.ContentKeyPolicySymmetricTokenKeyResponse, Outputs.ContentKeyPolicyX509CertificateTokenKeyResponse>> PrimaryVerificationKey;
        /// <summary>
        /// A list of required token claims.
        /// </summary>
        public readonly ImmutableArray<Outputs.ContentKeyPolicyTokenClaimResponse> RequiredClaims;
        /// <summary>
        /// The type of token.
        /// </summary>
        public readonly string RestrictionTokenType;

        [OutputConstructor]
        private ContentKeyPolicyTokenRestrictionResponse(
            ImmutableArray<Union<Outputs.ContentKeyPolicyRsaTokenKeyResponse, Union<Outputs.ContentKeyPolicySymmetricTokenKeyResponse, Outputs.ContentKeyPolicyX509CertificateTokenKeyResponse>>> alternateVerificationKeys,

            string audience,

            string issuer,

            string odataType,

            string? openIdConnectDiscoveryDocument,

            Union<Outputs.ContentKeyPolicyRsaTokenKeyResponse, Union<Outputs.ContentKeyPolicySymmetricTokenKeyResponse, Outputs.ContentKeyPolicyX509CertificateTokenKeyResponse>> primaryVerificationKey,

            ImmutableArray<Outputs.ContentKeyPolicyTokenClaimResponse> requiredClaims,

            string restrictionTokenType)
        {
            AlternateVerificationKeys = alternateVerificationKeys;
            Audience = audience;
            Issuer = issuer;
            OdataType = odataType;
            OpenIdConnectDiscoveryDocument = openIdConnectDiscoveryDocument;
            PrimaryVerificationKey = primaryVerificationKey;
            RequiredClaims = requiredClaims;
            RestrictionTokenType = restrictionTokenType;
        }
    }
}
