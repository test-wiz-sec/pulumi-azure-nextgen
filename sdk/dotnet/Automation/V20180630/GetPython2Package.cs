// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNextGen.Automation.V20180630
{
    public static class GetPython2Package
    {
        public static Task<GetPython2PackageResult> InvokeAsync(GetPython2PackageArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetPython2PackageResult>("azure-nextgen:automation/v20180630:getPython2Package", args ?? new GetPython2PackageArgs(), options.WithVersion());
    }


    public sealed class GetPython2PackageArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// The name of the automation account.
        /// </summary>
        [Input("automationAccountName", required: true)]
        public string AutomationAccountName { get; set; } = null!;

        /// <summary>
        /// The python package name.
        /// </summary>
        [Input("packageName", required: true)]
        public string PackageName { get; set; } = null!;

        /// <summary>
        /// Name of an Azure Resource group.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public string ResourceGroupName { get; set; } = null!;

        public GetPython2PackageArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetPython2PackageResult
    {
        /// <summary>
        /// Gets or sets the activity count of the module.
        /// </summary>
        public readonly int? ActivityCount;
        /// <summary>
        /// Gets or sets the contentLink of the module.
        /// </summary>
        public readonly Outputs.ContentLinkResponse? ContentLink;
        /// <summary>
        /// Gets or sets the creation time.
        /// </summary>
        public readonly string? CreationTime;
        /// <summary>
        /// Gets or sets the description.
        /// </summary>
        public readonly string? Description;
        /// <summary>
        /// Gets or sets the error info of the module.
        /// </summary>
        public readonly Outputs.ModuleErrorInfoResponse? Error;
        /// <summary>
        /// Gets or sets the etag of the resource.
        /// </summary>
        public readonly string? Etag;
        /// <summary>
        /// Gets or sets type of module, if its composite or not.
        /// </summary>
        public readonly bool? IsComposite;
        /// <summary>
        /// Gets or sets the isGlobal flag of the module.
        /// </summary>
        public readonly bool? IsGlobal;
        /// <summary>
        /// Gets or sets the last modified time.
        /// </summary>
        public readonly string? LastModifiedTime;
        /// <summary>
        /// The Azure Region where the resource lives
        /// </summary>
        public readonly string? Location;
        /// <summary>
        /// The name of the resource
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// Gets or sets the provisioning state of the module.
        /// </summary>
        public readonly string? ProvisioningState;
        /// <summary>
        /// Gets or sets the size in bytes of the module.
        /// </summary>
        public readonly int? SizeInBytes;
        /// <summary>
        /// Resource tags.
        /// </summary>
        public readonly ImmutableDictionary<string, string>? Tags;
        /// <summary>
        /// The type of the resource.
        /// </summary>
        public readonly string Type;
        /// <summary>
        /// Gets or sets the version of the module.
        /// </summary>
        public readonly string? Version;

        [OutputConstructor]
        private GetPython2PackageResult(
            int? activityCount,

            Outputs.ContentLinkResponse? contentLink,

            string? creationTime,

            string? description,

            Outputs.ModuleErrorInfoResponse? error,

            string? etag,

            bool? isComposite,

            bool? isGlobal,

            string? lastModifiedTime,

            string? location,

            string name,

            string? provisioningState,

            int? sizeInBytes,

            ImmutableDictionary<string, string>? tags,

            string type,

            string? version)
        {
            ActivityCount = activityCount;
            ContentLink = contentLink;
            CreationTime = creationTime;
            Description = description;
            Error = error;
            Etag = etag;
            IsComposite = isComposite;
            IsGlobal = isGlobal;
            LastModifiedTime = lastModifiedTime;
            Location = location;
            Name = name;
            ProvisioningState = provisioningState;
            SizeInBytes = sizeInBytes;
            Tags = tags;
            Type = type;
            Version = version;
        }
    }
}