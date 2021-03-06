// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNextGen.Automation.V20151031
{
    public static class GetCertificate
    {
        public static Task<GetCertificateResult> InvokeAsync(GetCertificateArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetCertificateResult>("azure-nextgen:automation/v20151031:getCertificate", args ?? new GetCertificateArgs(), options.WithVersion());
    }


    public sealed class GetCertificateArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// The name of the automation account.
        /// </summary>
        [Input("automationAccountName", required: true)]
        public string AutomationAccountName { get; set; } = null!;

        /// <summary>
        /// The name of certificate.
        /// </summary>
        [Input("certificateName", required: true)]
        public string CertificateName { get; set; } = null!;

        /// <summary>
        /// Name of an Azure Resource group.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public string ResourceGroupName { get; set; } = null!;

        public GetCertificateArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetCertificateResult
    {
        /// <summary>
        /// Gets the creation time.
        /// </summary>
        public readonly string CreationTime;
        /// <summary>
        /// Gets or sets the description.
        /// </summary>
        public readonly string? Description;
        /// <summary>
        /// Gets the expiry time of the certificate.
        /// </summary>
        public readonly string ExpiryTime;
        /// <summary>
        /// Gets the is exportable flag of the certificate.
        /// </summary>
        public readonly bool IsExportable;
        /// <summary>
        /// Gets the last modified time.
        /// </summary>
        public readonly string LastModifiedTime;
        /// <summary>
        /// The name of the resource
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// Gets the thumbprint of the certificate.
        /// </summary>
        public readonly string Thumbprint;
        /// <summary>
        /// The type of the resource.
        /// </summary>
        public readonly string Type;

        [OutputConstructor]
        private GetCertificateResult(
            string creationTime,

            string? description,

            string expiryTime,

            bool isExportable,

            string lastModifiedTime,

            string name,

            string thumbprint,

            string type)
        {
            CreationTime = creationTime;
            Description = description;
            ExpiryTime = expiryTime;
            IsExportable = isExportable;
            LastModifiedTime = lastModifiedTime;
            Name = name;
            Thumbprint = thumbprint;
            Type = type;
        }
    }
}
