// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNextGen.Migrate.Latest.Outputs
{

    [OutputType]
    public sealed class ProjectPropertiesResponse
    {
        /// <summary>
        /// Assessment solution ARM id tracked by Microsoft.Migrate/migrateProjects.
        /// </summary>
        public readonly string? AssessmentSolutionId;
        /// <summary>
        /// Time when this project was created. Date-Time represented in ISO-8601 format.
        /// </summary>
        public readonly string CreatedTimestamp;
        /// <summary>
        /// The ARM id of service map workspace created by customer.
        /// </summary>
        public readonly string? CustomerWorkspaceId;
        /// <summary>
        /// Location of service map workspace created by customer.
        /// </summary>
        public readonly string? CustomerWorkspaceLocation;
        /// <summary>
        /// Time when last assessment was created. Date-Time represented in ISO-8601 format. This value will be null until assessment is created.
        /// </summary>
        public readonly string LastAssessmentTimestamp;
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
        /// Assessment project status.
        /// </summary>
        public readonly string? ProjectStatus;
        /// <summary>
        /// Provisioning state of the project.
        /// </summary>
        public readonly string ProvisioningState;
        /// <summary>
        /// Endpoint at which the collector agent can call agent REST API.
        /// </summary>
        public readonly string ServiceEndpoint;
        /// <summary>
        /// Time when this project was last updated. Date-Time represented in ISO-8601 format.
        /// </summary>
        public readonly string UpdatedTimestamp;

        [OutputConstructor]
        private ProjectPropertiesResponse(
            string? assessmentSolutionId,

            string createdTimestamp,

            string? customerWorkspaceId,

            string? customerWorkspaceLocation,

            string lastAssessmentTimestamp,

            int numberOfAssessments,

            int numberOfGroups,

            int numberOfMachines,

            string? projectStatus,

            string provisioningState,

            string serviceEndpoint,

            string updatedTimestamp)
        {
            AssessmentSolutionId = assessmentSolutionId;
            CreatedTimestamp = createdTimestamp;
            CustomerWorkspaceId = customerWorkspaceId;
            CustomerWorkspaceLocation = customerWorkspaceLocation;
            LastAssessmentTimestamp = lastAssessmentTimestamp;
            NumberOfAssessments = numberOfAssessments;
            NumberOfGroups = numberOfGroups;
            NumberOfMachines = numberOfMachines;
            ProjectStatus = projectStatus;
            ProvisioningState = provisioningState;
            ServiceEndpoint = serviceEndpoint;
            UpdatedTimestamp = updatedTimestamp;
        }
    }
}
