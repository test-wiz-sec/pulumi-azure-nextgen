// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNextGen.Migrate.V20171111Preview
{
    public static class GetProject
    {
        public static Task<GetProjectResult> InvokeAsync(GetProjectArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetProjectResult>("azure-nextgen:migrate/v20171111preview:getProject", args ?? new GetProjectArgs(), options.WithVersion());
    }


    public sealed class GetProjectArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// Name of the Azure Migrate project.
        /// </summary>
        [Input("projectName", required: true)]
        public string ProjectName { get; set; } = null!;

        /// <summary>
        /// Name of the Azure Resource Group that project is part of.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public string ResourceGroupName { get; set; } = null!;

        public GetProjectArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetProjectResult
    {
        /// <summary>
        /// Time when this project was created. Date-Time represented in ISO-8601 format.
        /// </summary>
        public readonly string CreatedTimestamp;
        /// <summary>
        /// ARM ID of the Service Map workspace created by user.
        /// </summary>
        public readonly string? CustomerWorkspaceId;
        /// <summary>
        /// Reports whether project is under discovery.
        /// </summary>
        public readonly string DiscoveryStatus;
        /// <summary>
        /// For optimistic concurrency control.
        /// </summary>
        public readonly string? ETag;
        /// <summary>
        /// Azure location in which project is created.
        /// </summary>
        public readonly string? Location;
        /// <summary>
        /// Name of the project.
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// Number of assessments created in the project.
        /// </summary>
        public readonly int NumberOfAssessments;
        /// <summary>
        /// Number of groups created in the project.
        /// </summary>
        public readonly int NumberOfGroups;
        /// <summary>
        /// Number of machines in the project.
        /// </summary>
        public readonly int NumberOfMachines;
        /// <summary>
        /// Provisioning state of the project.
        /// </summary>
        public readonly string? ProvisioningState;
        /// <summary>
        /// Tags provided by Azure Tagging service.
        /// </summary>
        public readonly object? Tags;
        /// <summary>
        /// Type of the object = [Microsoft.Migrate/projects].
        /// </summary>
        public readonly string Type;
        /// <summary>
        /// Time when this project was last updated. Date-Time represented in ISO-8601 format.
        /// </summary>
        public readonly string UpdatedTimestamp;

        [OutputConstructor]
        private GetProjectResult(
            string createdTimestamp,

            string? customerWorkspaceId,

            string discoveryStatus,

            string? eTag,

            string? location,

            string name,

            int numberOfAssessments,

            int numberOfGroups,

            int numberOfMachines,

            string? provisioningState,

            object? tags,

            string type,

            string updatedTimestamp)
        {
            CreatedTimestamp = createdTimestamp;
            CustomerWorkspaceId = customerWorkspaceId;
            DiscoveryStatus = discoveryStatus;
            ETag = eTag;
            Location = location;
            Name = name;
            NumberOfAssessments = numberOfAssessments;
            NumberOfGroups = numberOfGroups;
            NumberOfMachines = numberOfMachines;
            ProvisioningState = provisioningState;
            Tags = tags;
            Type = type;
            UpdatedTimestamp = updatedTimestamp;
        }
    }
}
