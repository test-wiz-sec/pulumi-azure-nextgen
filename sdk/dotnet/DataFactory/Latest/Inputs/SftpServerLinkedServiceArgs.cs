// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNextGen.DataFactory.Latest.Inputs
{

    /// <summary>
    /// A linked service for an SSH File Transfer Protocol (SFTP) server. 
    /// </summary>
    public sealed class SftpServerLinkedServiceArgs : Pulumi.ResourceArgs
    {
        [Input("annotations")]
        private InputList<ImmutableDictionary<string, object>>? _annotations;

        /// <summary>
        /// List of tags that can be used for describing the linked service.
        /// </summary>
        public InputList<ImmutableDictionary<string, object>> Annotations
        {
            get => _annotations ?? (_annotations = new InputList<ImmutableDictionary<string, object>>());
            set => _annotations = value;
        }

        /// <summary>
        /// The authentication type to be used to connect to the FTP server.
        /// </summary>
        [Input("authenticationType")]
        public Input<string>? AuthenticationType { get; set; }

        /// <summary>
        /// The integration runtime reference.
        /// </summary>
        [Input("connectVia")]
        public Input<Inputs.IntegrationRuntimeReferenceArgs>? ConnectVia { get; set; }

        /// <summary>
        /// Linked service description.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        [Input("encryptedCredential")]
        private InputMap<object>? _encryptedCredential;

        /// <summary>
        /// The encrypted credential used for authentication. Credentials are encrypted using the integration runtime credential manager. Type: string (or Expression with resultType string).
        /// </summary>
        public InputMap<object> EncryptedCredential
        {
            get => _encryptedCredential ?? (_encryptedCredential = new InputMap<object>());
            set => _encryptedCredential = value;
        }

        [Input("host", required: true)]
        private InputMap<object>? _host;

        /// <summary>
        /// The SFTP server host name. Type: string (or Expression with resultType string).
        /// </summary>
        public InputMap<object> Host
        {
            get => _host ?? (_host = new InputMap<object>());
            set => _host = value;
        }

        [Input("hostKeyFingerprint")]
        private InputMap<object>? _hostKeyFingerprint;

        /// <summary>
        /// The host key finger-print of the SFTP server. When SkipHostKeyValidation is false, HostKeyFingerprint should be specified. Type: string (or Expression with resultType string).
        /// </summary>
        public InputMap<object> HostKeyFingerprint
        {
            get => _hostKeyFingerprint ?? (_hostKeyFingerprint = new InputMap<object>());
            set => _hostKeyFingerprint = value;
        }

        [Input("parameters")]
        private InputMap<Inputs.ParameterSpecificationArgs>? _parameters;

        /// <summary>
        /// Parameters for linked service.
        /// </summary>
        public InputMap<Inputs.ParameterSpecificationArgs> Parameters
        {
            get => _parameters ?? (_parameters = new InputMap<Inputs.ParameterSpecificationArgs>());
            set => _parameters = value;
        }

        /// <summary>
        /// The password to decrypt the SSH private key if the SSH private key is encrypted.
        /// </summary>
        [Input("passPhrase")]
        public InputUnion<Inputs.AzureKeyVaultSecretReferenceArgs, Inputs.SecureStringArgs>? PassPhrase { get; set; }

        /// <summary>
        /// Password to logon the SFTP server for Basic authentication.
        /// </summary>
        [Input("password")]
        public InputUnion<Inputs.AzureKeyVaultSecretReferenceArgs, Inputs.SecureStringArgs>? Password { get; set; }

        [Input("port")]
        private InputMap<object>? _port;

        /// <summary>
        /// The TCP port number that the SFTP server uses to listen for client connections. Default value is 22. Type: integer (or Expression with resultType integer), minimum: 0.
        /// </summary>
        public InputMap<object> Port
        {
            get => _port ?? (_port = new InputMap<object>());
            set => _port = value;
        }

        /// <summary>
        /// Base64 encoded SSH private key content for SshPublicKey authentication. For on-premises copy with SshPublicKey authentication, either PrivateKeyPath or PrivateKeyContent should be specified. SSH private key should be OpenSSH format.
        /// </summary>
        [Input("privateKeyContent")]
        public InputUnion<Inputs.AzureKeyVaultSecretReferenceArgs, Inputs.SecureStringArgs>? PrivateKeyContent { get; set; }

        [Input("privateKeyPath")]
        private InputMap<object>? _privateKeyPath;

        /// <summary>
        /// The SSH private key file path for SshPublicKey authentication. Only valid for on-premises copy. For on-premises copy with SshPublicKey authentication, either PrivateKeyPath or PrivateKeyContent should be specified. SSH private key should be OpenSSH format. Type: string (or Expression with resultType string).
        /// </summary>
        public InputMap<object> PrivateKeyPath
        {
            get => _privateKeyPath ?? (_privateKeyPath = new InputMap<object>());
            set => _privateKeyPath = value;
        }

        [Input("skipHostKeyValidation")]
        private InputMap<object>? _skipHostKeyValidation;

        /// <summary>
        /// If true, skip the SSH host key validation. Default value is false. Type: boolean (or Expression with resultType boolean).
        /// </summary>
        public InputMap<object> SkipHostKeyValidation
        {
            get => _skipHostKeyValidation ?? (_skipHostKeyValidation = new InputMap<object>());
            set => _skipHostKeyValidation = value;
        }

        /// <summary>
        /// Type of linked service.
        /// </summary>
        [Input("type", required: true)]
        public Input<string> Type { get; set; } = null!;

        [Input("userName")]
        private InputMap<object>? _userName;

        /// <summary>
        /// The username used to log on to the SFTP server. Type: string (or Expression with resultType string).
        /// </summary>
        public InputMap<object> UserName
        {
            get => _userName ?? (_userName = new InputMap<object>());
            set => _userName = value;
        }

        public SftpServerLinkedServiceArgs()
        {
        }
    }
}