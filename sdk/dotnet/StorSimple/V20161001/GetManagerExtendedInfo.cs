// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNextGen.StorSimple.V20161001
{
    public static class GetManagerExtendedInfo
    {
        public static Task<GetManagerExtendedInfoResult> InvokeAsync(GetManagerExtendedInfoArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetManagerExtendedInfoResult>("azure-nextgen:storsimple/v20161001:getManagerExtendedInfo", args ?? new GetManagerExtendedInfoArgs(), options.WithVersion());
    }


    public sealed class GetManagerExtendedInfoArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// The manager name
        /// </summary>
        [Input("managerName", required: true)]
        public string ManagerName { get; set; } = null!;

        /// <summary>
        /// The resource group name
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public string ResourceGroupName { get; set; } = null!;

        public GetManagerExtendedInfoArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetManagerExtendedInfoResult
    {
        /// <summary>
        /// Represents the encryption algorithm used to encrypt the other keys. None - if EncryptionKey is saved in plain text format. AlgorithmName - if encryption is used
        /// </summary>
        public readonly string Algorithm;
        /// <summary>
        /// Represents the CEK of the resource
        /// </summary>
        public readonly string? EncryptionKey;
        /// <summary>
        /// Represents the Cert thumbprint that was used to encrypt the CEK
        /// </summary>
        public readonly string? EncryptionKeyThumbprint;
        /// <summary>
        /// ETag of the Resource
        /// </summary>
        public readonly string? Etag;
        /// <summary>
        /// Represents the CIK of the resource
        /// </summary>
        public readonly string IntegrityKey;
        /// <summary>
        /// The name.
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// Represents the portal thumbprint which can be used optionally to encrypt the entire data before storing it.
        /// </summary>
        public readonly string? PortalCertificateThumbprint;
        /// <summary>
        /// The type.
        /// </summary>
        public readonly string Type;
        /// <summary>
        /// Represents the version of the ExtendedInfo object being persisted
        /// </summary>
        public readonly string? Version;

        [OutputConstructor]
        private GetManagerExtendedInfoResult(
            string algorithm,

            string? encryptionKey,

            string? encryptionKeyThumbprint,

            string? etag,

            string integrityKey,

            string name,

            string? portalCertificateThumbprint,

            string type,

            string? version)
        {
            Algorithm = algorithm;
            EncryptionKey = encryptionKey;
            EncryptionKeyThumbprint = encryptionKeyThumbprint;
            Etag = etag;
            IntegrityKey = integrityKey;
            Name = name;
            PortalCertificateThumbprint = portalCertificateThumbprint;
            Type = type;
            Version = version;
        }
    }
}