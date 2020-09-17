// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNextGen.CustomProviders.V20180901Preview.Outputs
{

    [OutputType]
    public sealed class CustomRPValidationsResponse
    {
        /// <summary>
        /// A link to the validation specification. The specification must be hosted on raw.githubusercontent.com.
        /// </summary>
        public readonly string Specification;
        /// <summary>
        /// The type of validation to run against a matching request.
        /// </summary>
        public readonly string? ValidationType;

        [OutputConstructor]
        private CustomRPValidationsResponse(
            string specification,

            string? validationType)
        {
            Specification = specification;
            ValidationType = validationType;
        }
    }
}