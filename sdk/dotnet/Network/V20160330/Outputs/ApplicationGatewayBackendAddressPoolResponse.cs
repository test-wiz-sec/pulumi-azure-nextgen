// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNextGen.Network.V20160330.Outputs
{

    [OutputType]
    public sealed class ApplicationGatewayBackendAddressPoolResponse
    {
        /// <summary>
        /// Gets or sets the backend addresses
        /// </summary>
        public readonly ImmutableArray<Outputs.ApplicationGatewayBackendAddressResponse> BackendAddresses;
        /// <summary>
        /// Gets collection of references to IPs defined in NICs
        /// </summary>
        public readonly ImmutableArray<Outputs.NetworkInterfaceIPConfigurationResponse> BackendIPConfigurations;
        /// <summary>
        /// A unique read-only string that changes whenever the resource is updated
        /// </summary>
        public readonly string? Etag;
        /// <summary>
        /// Resource Id
        /// </summary>
        public readonly string? Id;
        /// <summary>
        /// Gets name of the resource that is unique within a resource group. This name can be used to access the resource
        /// </summary>
        public readonly string? Name;
        /// <summary>
        /// Gets or sets Provisioning state of the backend address pool resource Updating/Deleting/Failed
        /// </summary>
        public readonly string? ProvisioningState;

        [OutputConstructor]
        private ApplicationGatewayBackendAddressPoolResponse(
            ImmutableArray<Outputs.ApplicationGatewayBackendAddressResponse> backendAddresses,

            ImmutableArray<Outputs.NetworkInterfaceIPConfigurationResponse> backendIPConfigurations,

            string? etag,

            string? id,

            string? name,

            string? provisioningState)
        {
            BackendAddresses = backendAddresses;
            BackendIPConfigurations = backendIPConfigurations;
            Etag = etag;
            Id = id;
            Name = name;
            ProvisioningState = provisioningState;
        }
    }
}