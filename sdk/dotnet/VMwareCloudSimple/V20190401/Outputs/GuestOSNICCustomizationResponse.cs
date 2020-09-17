// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNextGen.VMwareCloudSimple.V20190401.Outputs
{

    [OutputType]
    public sealed class GuestOSNICCustomizationResponse
    {
        /// <summary>
        /// IP address allocation method
        /// </summary>
        public readonly string? Allocation;
        /// <summary>
        /// List of dns servers to use
        /// </summary>
        public readonly ImmutableArray<string> DnsServers;
        /// <summary>
        /// Gateway addresses assigned to nic
        /// </summary>
        public readonly ImmutableArray<string> Gateway;
        /// <summary>
        /// Static ip address for nic
        /// </summary>
        public readonly string? IpAddress;
        /// <summary>
        /// Network mask for nic
        /// </summary>
        public readonly string? Mask;
        /// <summary>
        /// primary WINS server for Windows
        /// </summary>
        public readonly string? PrimaryWinsServer;
        /// <summary>
        /// secondary WINS server for Windows
        /// </summary>
        public readonly string? SecondaryWinsServer;

        [OutputConstructor]
        private GuestOSNICCustomizationResponse(
            string? allocation,

            ImmutableArray<string> dnsServers,

            ImmutableArray<string> gateway,

            string? ipAddress,

            string? mask,

            string? primaryWinsServer,

            string? secondaryWinsServer)
        {
            Allocation = allocation;
            DnsServers = dnsServers;
            Gateway = gateway;
            IpAddress = ipAddress;
            Mask = mask;
            PrimaryWinsServer = primaryWinsServer;
            SecondaryWinsServer = secondaryWinsServer;
        }
    }
}