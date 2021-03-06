// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNextGen.EngagementFabric.V20180901Preview
{
    public static class ListAccountKeys
    {
        public static Task<ListAccountKeysResult> InvokeAsync(ListAccountKeysArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<ListAccountKeysResult>("azure-nextgen:engagementfabric/v20180901preview:listAccountKeys", args ?? new ListAccountKeysArgs(), options.WithVersion());
    }


    public sealed class ListAccountKeysArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// Account Name
        /// </summary>
        [Input("accountName", required: true)]
        public string AccountName { get; set; } = null!;

        /// <summary>
        /// Resource Group Name
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public string ResourceGroupName { get; set; } = null!;

        public ListAccountKeysArgs()
        {
        }
    }


    [OutputType]
    public sealed class ListAccountKeysResult
    {
        /// <summary>
        /// Account keys
        /// </summary>
        public readonly ImmutableArray<Outputs.KeyDescriptionResponseResult> Value;

        [OutputConstructor]
        private ListAccountKeysResult(ImmutableArray<Outputs.KeyDescriptionResponseResult> value)
        {
            Value = value;
        }
    }
}
