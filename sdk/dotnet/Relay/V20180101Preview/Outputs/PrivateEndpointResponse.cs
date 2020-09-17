// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNextGen.Relay.V20180101Preview.Outputs
{

    [OutputType]
    public sealed class PrivateEndpointResponse
    {
        /// <summary>
        /// Full identifier of the private endpoint resource.
        /// </summary>
        public readonly string? Id;

        [OutputConstructor]
        private PrivateEndpointResponse(string? id)
        {
            Id = id;
        }
    }
}