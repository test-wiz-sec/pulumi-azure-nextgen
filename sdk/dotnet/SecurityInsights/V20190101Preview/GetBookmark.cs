// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNextGen.SecurityInsights.V20190101Preview
{
    public static class GetBookmark
    {
        public static Task<GetBookmarkResult> InvokeAsync(GetBookmarkArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetBookmarkResult>("azure-nextgen:securityinsights/v20190101preview:getBookmark", args ?? new GetBookmarkArgs(), options.WithVersion());
    }


    public sealed class GetBookmarkArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// Bookmark ID
        /// </summary>
        [Input("bookmarkId", required: true)]
        public string BookmarkId { get; set; } = null!;

        /// <summary>
        /// The namespace of workspaces resource provider- Microsoft.OperationalInsights.
        /// </summary>
        [Input("operationalInsightsResourceProvider", required: true)]
        public string OperationalInsightsResourceProvider { get; set; } = null!;

        /// <summary>
        /// The name of the resource group within the user's subscription. The name is case insensitive.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public string ResourceGroupName { get; set; } = null!;

        /// <summary>
        /// The name of the workspace.
        /// </summary>
        [Input("workspaceName", required: true)]
        public string WorkspaceName { get; set; } = null!;

        public GetBookmarkArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetBookmarkResult
    {
        /// <summary>
        /// The time the bookmark was created
        /// </summary>
        public readonly string? Created;
        /// <summary>
        /// Describes a user that created the bookmark
        /// </summary>
        public readonly Outputs.UserInfoResponse? CreatedBy;
        /// <summary>
        /// The display name of the bookmark
        /// </summary>
        public readonly string DisplayName;
        /// <summary>
        /// Etag of the azure resource
        /// </summary>
        public readonly string? Etag;
        /// <summary>
        /// Describes an incident that relates to bookmark
        /// </summary>
        public readonly Outputs.IncidentInfoResponse? IncidentInfo;
        /// <summary>
        /// List of labels relevant to this bookmark
        /// </summary>
        public readonly ImmutableArray<string> Labels;
        /// <summary>
        /// Azure resource name
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// The notes of the bookmark
        /// </summary>
        public readonly string? Notes;
        /// <summary>
        /// The query of the bookmark.
        /// </summary>
        public readonly string Query;
        /// <summary>
        /// The query result of the bookmark.
        /// </summary>
        public readonly string? QueryResult;
        /// <summary>
        /// Azure resource type
        /// </summary>
        public readonly string Type;
        /// <summary>
        /// The last time the bookmark was updated
        /// </summary>
        public readonly string? Updated;
        /// <summary>
        /// Describes a user that updated the bookmark
        /// </summary>
        public readonly Outputs.UserInfoResponse? UpdatedBy;

        [OutputConstructor]
        private GetBookmarkResult(
            string? created,

            Outputs.UserInfoResponse? createdBy,

            string displayName,

            string? etag,

            Outputs.IncidentInfoResponse? incidentInfo,

            ImmutableArray<string> labels,

            string name,

            string? notes,

            string query,

            string? queryResult,

            string type,

            string? updated,

            Outputs.UserInfoResponse? updatedBy)
        {
            Created = created;
            CreatedBy = createdBy;
            DisplayName = displayName;
            Etag = etag;
            IncidentInfo = incidentInfo;
            Labels = labels;
            Name = name;
            Notes = notes;
            Query = query;
            QueryResult = queryResult;
            Type = type;
            Updated = updated;
            UpdatedBy = updatedBy;
        }
    }
}
