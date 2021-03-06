// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNextGen.Batch.V20181201
{
    public static class GetApplication
    {
        public static Task<GetApplicationResult> InvokeAsync(GetApplicationArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetApplicationResult>("azure-nextgen:batch/v20181201:getApplication", args ?? new GetApplicationArgs(), options.WithVersion());
    }


    public sealed class GetApplicationArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// The name of the Batch account.
        /// </summary>
        [Input("accountName", required: true)]
        public string AccountName { get; set; } = null!;

        /// <summary>
        /// The name of the application. This must be unique within the account.
        /// </summary>
        [Input("applicationName", required: true)]
        public string ApplicationName { get; set; } = null!;

        /// <summary>
        /// The name of the resource group that contains the Batch account.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public string ResourceGroupName { get; set; } = null!;

        public GetApplicationArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetApplicationResult
    {
        /// <summary>
        /// A value indicating whether packages within the application may be overwritten using the same version string.
        /// </summary>
        public readonly bool? AllowUpdates;
        /// <summary>
        /// The package to use if a client requests the application but does not specify a version. This property can only be set to the name of an existing package.
        /// </summary>
        public readonly string? DefaultVersion;
        /// <summary>
        /// The display name for the application.
        /// </summary>
        public readonly string? DisplayName;
        /// <summary>
        /// The ETag of the resource, used for concurrency statements.
        /// </summary>
        public readonly string Etag;
        /// <summary>
        /// The name of the resource.
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// The type of the resource.
        /// </summary>
        public readonly string Type;

        [OutputConstructor]
        private GetApplicationResult(
            bool? allowUpdates,

            string? defaultVersion,

            string? displayName,

            string etag,

            string name,

            string type)
        {
            AllowUpdates = allowUpdates;
            DefaultVersion = defaultVersion;
            DisplayName = displayName;
            Etag = etag;
            Name = name;
            Type = type;
        }
    }
}
