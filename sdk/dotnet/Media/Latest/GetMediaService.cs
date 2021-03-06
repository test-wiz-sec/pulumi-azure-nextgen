// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNextGen.Media.Latest
{
    public static class GetMediaService
    {
        public static Task<GetMediaServiceResult> InvokeAsync(GetMediaServiceArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetMediaServiceResult>("azure-nextgen:media/latest:getMediaService", args ?? new GetMediaServiceArgs(), options.WithVersion());
    }


    public sealed class GetMediaServiceArgs : Pulumi.InvokeArgs
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

        public GetMediaServiceArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetMediaServiceResult
    {
        /// <summary>
        /// The account encryption properties.
        /// </summary>
        public readonly Outputs.AccountEncryptionResponse? Encryption;
        /// <summary>
        /// The Managed Identity for the Media Services account.
        /// </summary>
        public readonly Outputs.MediaServiceIdentityResponse? Identity;
        /// <summary>
        /// The geo-location where the resource lives
        /// </summary>
        public readonly string Location;
        /// <summary>
        /// The Media Services account ID.
        /// </summary>
        public readonly string MediaServiceId;
        /// <summary>
        /// The name of the resource
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// The storage accounts for this resource.
        /// </summary>
        public readonly ImmutableArray<Outputs.StorageAccountResponse> StorageAccounts;
        public readonly string? StorageAuthentication;
        /// <summary>
        /// Resource tags.
        /// </summary>
        public readonly ImmutableDictionary<string, string>? Tags;
        /// <summary>
        /// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
        /// </summary>
        public readonly string Type;

        [OutputConstructor]
        private GetMediaServiceResult(
            Outputs.AccountEncryptionResponse? encryption,

            Outputs.MediaServiceIdentityResponse? identity,

            string location,

            string mediaServiceId,

            string name,

            ImmutableArray<Outputs.StorageAccountResponse> storageAccounts,

            string? storageAuthentication,

            ImmutableDictionary<string, string>? tags,

            string type)
        {
            Encryption = encryption;
            Identity = identity;
            Location = location;
            MediaServiceId = mediaServiceId;
            Name = name;
            StorageAccounts = storageAccounts;
            StorageAuthentication = storageAuthentication;
            Tags = tags;
            Type = type;
        }
    }
}
