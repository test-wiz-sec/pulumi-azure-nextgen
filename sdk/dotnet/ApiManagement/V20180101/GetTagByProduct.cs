// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNextGen.ApiManagement.V20180101
{
    public static class GetTagByProduct
    {
        public static Task<GetTagByProductResult> InvokeAsync(GetTagByProductArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetTagByProductResult>("azure-nextgen:apimanagement/v20180101:getTagByProduct", args ?? new GetTagByProductArgs(), options.WithVersion());
    }


    public sealed class GetTagByProductArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// Product identifier. Must be unique in the current API Management service instance.
        /// </summary>
        [Input("productId", required: true)]
        public string ProductId { get; set; } = null!;

        /// <summary>
        /// The name of the resource group.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public string ResourceGroupName { get; set; } = null!;

        /// <summary>
        /// The name of the API Management service.
        /// </summary>
        [Input("serviceName", required: true)]
        public string ServiceName { get; set; } = null!;

        /// <summary>
        /// Tag identifier. Must be unique in the current API Management service instance.
        /// </summary>
        [Input("tagId", required: true)]
        public string TagId { get; set; } = null!;

        public GetTagByProductArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetTagByProductResult
    {
        /// <summary>
        /// Tag name.
        /// </summary>
        public readonly string DisplayName;
        /// <summary>
        /// Resource name.
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// Resource type for API Management resource.
        /// </summary>
        public readonly string Type;

        [OutputConstructor]
        private GetTagByProductResult(
            string displayName,

            string name,

            string type)
        {
            DisplayName = displayName;
            Name = name;
            Type = type;
        }
    }
}
