// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNextGen.DBforMySQL.V20171201Preview.Outputs
{

    [OutputType]
    public sealed class PrivateEndpointConnectionPropertiesResponse
    {
        /// <summary>
        /// Private endpoint which the connection belongs to.
        /// </summary>
        public readonly Outputs.PrivateEndpointPropertyResponse? PrivateEndpoint;
        /// <summary>
        /// Connection state of the private endpoint connection.
        /// </summary>
        public readonly Outputs.PrivateLinkServiceConnectionStatePropertyResponse? PrivateLinkServiceConnectionState;
        /// <summary>
        /// State of the private endpoint connection.
        /// </summary>
        public readonly string ProvisioningState;

        [OutputConstructor]
        private PrivateEndpointConnectionPropertiesResponse(
            Outputs.PrivateEndpointPropertyResponse? privateEndpoint,

            Outputs.PrivateLinkServiceConnectionStatePropertyResponse? privateLinkServiceConnectionState,

            string provisioningState)
        {
            PrivateEndpoint = privateEndpoint;
            PrivateLinkServiceConnectionState = privateLinkServiceConnectionState;
            ProvisioningState = provisioningState;
        }
    }
}
