// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNextGen.DataFactory.V20180601.Outputs
{

    [OutputType]
    public sealed class CassandraLinkedServiceResponse
    {
        /// <summary>
        /// List of tags that can be used for describing the linked service.
        /// </summary>
        public readonly ImmutableArray<ImmutableDictionary<string, object>> Annotations;
        /// <summary>
        /// AuthenticationType to be used for connection. Type: string (or Expression with resultType string).
        /// </summary>
        public readonly ImmutableDictionary<string, object>? AuthenticationType;
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
        /// Host name for connection. Type: string (or Expression with resultType string).
        /// </summary>
        public readonly ImmutableDictionary<string, object> Host;
        /// <summary>
        /// Parameters for linked service.
        /// </summary>
        public readonly ImmutableDictionary<string, Outputs.ParameterSpecificationResponse>? Parameters;
        /// <summary>
        /// Password for authentication.
        /// </summary>
        public readonly Union<Outputs.AzureKeyVaultSecretReferenceResponse, Outputs.SecureStringResponse>? Password;
        /// <summary>
        /// The port for the connection. Type: integer (or Expression with resultType integer).
        /// </summary>
        public readonly ImmutableDictionary<string, object>? Port;
        /// <summary>
        /// Type of linked service.
        /// </summary>
        public readonly string Type;
        /// <summary>
        /// Username for authentication. Type: string (or Expression with resultType string).
        /// </summary>
        public readonly ImmutableDictionary<string, object>? Username;

        [OutputConstructor]
        private CassandraLinkedServiceResponse(
            ImmutableArray<ImmutableDictionary<string, object>> annotations,

            ImmutableDictionary<string, object>? authenticationType,

            Outputs.IntegrationRuntimeReferenceResponse? connectVia,

            string? description,

            ImmutableDictionary<string, object>? encryptedCredential,

            ImmutableDictionary<string, object> host,

            ImmutableDictionary<string, Outputs.ParameterSpecificationResponse>? parameters,

            Union<Outputs.AzureKeyVaultSecretReferenceResponse, Outputs.SecureStringResponse>? password,

            ImmutableDictionary<string, object>? port,

            string type,

            ImmutableDictionary<string, object>? username)
        {
            Annotations = annotations;
            AuthenticationType = authenticationType;
            ConnectVia = connectVia;
            Description = description;
            EncryptedCredential = encryptedCredential;
            Host = host;
            Parameters = parameters;
            Password = password;
            Port = port;
            Type = type;
            Username = username;
        }
    }
}