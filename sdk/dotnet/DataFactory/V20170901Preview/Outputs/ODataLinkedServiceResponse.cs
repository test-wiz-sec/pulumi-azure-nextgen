// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNextGen.DataFactory.V20170901Preview.Outputs
{

    [OutputType]
    public sealed class ODataLinkedServiceResponse
    {
        /// <summary>
        /// List of tags that can be used for describing the Dataset.
        /// </summary>
        public readonly ImmutableArray<ImmutableDictionary<string, object>> Annotations;
        /// <summary>
        /// Type of authentication used to connect to the OData service.
        /// </summary>
        public readonly string? AuthenticationType;
        /// <summary>
        /// The integration runtime reference.
        /// </summary>
        public readonly Outputs.IntegrationRuntimeReferenceResponse? ConnectVia;
        /// <summary>
        /// Linked service description.
        /// </summary>
        public readonly string? Description;
        /// <summary>
        /// The encrypted credential used for authentication. Credentials are encrypted using the integration runtime credential manager. Type: string (or Expression with resultType string).
        /// </summary>
        public readonly ImmutableDictionary<string, object>? EncryptedCredential;
        /// <summary>
        /// Parameters for linked service.
        /// </summary>
        public readonly ImmutableDictionary<string, Outputs.ParameterSpecificationResponse>? Parameters;
        /// <summary>
        /// Password of the OData service.
        /// </summary>
        public readonly Union<Outputs.AzureKeyVaultSecretReferenceResponse, Outputs.SecureStringResponse>? Password;
        /// <summary>
        /// Type of linked service.
        /// </summary>
        public readonly string Type;
        /// <summary>
        /// The URL of the OData service endpoint. Type: string (or Expression with resultType string).
        /// </summary>
        public readonly ImmutableDictionary<string, object> Url;
        /// <summary>
        /// User name of the OData service. Type: string (or Expression with resultType string).
        /// </summary>
        public readonly ImmutableDictionary<string, object>? UserName;

        [OutputConstructor]
        private ODataLinkedServiceResponse(
            ImmutableArray<ImmutableDictionary<string, object>> annotations,

            string? authenticationType,

            Outputs.IntegrationRuntimeReferenceResponse? connectVia,

            string? description,

            ImmutableDictionary<string, object>? encryptedCredential,

            ImmutableDictionary<string, Outputs.ParameterSpecificationResponse>? parameters,

            Union<Outputs.AzureKeyVaultSecretReferenceResponse, Outputs.SecureStringResponse>? password,

            string type,

            ImmutableDictionary<string, object> url,

            ImmutableDictionary<string, object>? userName)
        {
            Annotations = annotations;
            AuthenticationType = authenticationType;
            ConnectVia = connectVia;
            Description = description;
            EncryptedCredential = encryptedCredential;
            Parameters = parameters;
            Password = password;
            Type = type;
            Url = url;
            UserName = userName;
        }
    }
}