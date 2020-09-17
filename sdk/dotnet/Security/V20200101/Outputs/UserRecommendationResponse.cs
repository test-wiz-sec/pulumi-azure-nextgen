// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNextGen.Security.V20200101.Outputs
{

    [OutputType]
    public sealed class UserRecommendationResponse
    {
        /// <summary>
        /// The recommendation action of the machine or rule
        /// </summary>
        public readonly string? RecommendationAction;
        /// <summary>
        /// Represents a user that is recommended to be allowed for a certain rule
        /// </summary>
        public readonly string? Username;

        [OutputConstructor]
        private UserRecommendationResponse(
            string? recommendationAction,

            string? username)
        {
            RecommendationAction = recommendationAction;
            Username = username;
        }
    }
}