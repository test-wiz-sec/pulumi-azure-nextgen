// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNextGen.Logic.V20190501.Outputs
{

    [OutputType]
    public sealed class IntegrationServiceEnvironmentAccessEndpointResponse
    {
        /// <summary>
        /// The access endpoint type.
        /// </summary>
        public readonly string? Type;

        [OutputConstructor]
        private IntegrationServiceEnvironmentAccessEndpointResponse(string? type)
        {
            Type = type;
        }
    }
}
