// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNextGen.Network.V20160601.Outputs
{

    [OutputType]
    public sealed class NetworkInterfaceDnsSettingsResponse
    {
        /// <summary>
        /// Gets or sets list of Applied DNS servers IP addresses
        /// </summary>
        public readonly ImmutableArray<string> AppliedDnsServers;
        /// <summary>
        /// Gets or sets list of DNS servers IP addresses
        /// </summary>
        public readonly ImmutableArray<string> DnsServers;
        /// <summary>
        /// Gets or sets the internal DNS name
        /// </summary>
        public readonly string? InternalDnsNameLabel;
        /// <summary>
        /// Gets or sets internal domain name suffix of the NIC.
        /// </summary>
        public readonly string? InternalDomainNameSuffix;
        /// <summary>
        /// Gets or sets the internal fqdn.
        /// </summary>
        public readonly string? InternalFqdn;

        [OutputConstructor]
        private NetworkInterfaceDnsSettingsResponse(
            ImmutableArray<string> appliedDnsServers,

            ImmutableArray<string> dnsServers,

            string? internalDnsNameLabel,

            string? internalDomainNameSuffix,

            string? internalFqdn)
        {
            AppliedDnsServers = appliedDnsServers;
            DnsServers = dnsServers;
            InternalDnsNameLabel = internalDnsNameLabel;
            InternalDomainNameSuffix = internalDomainNameSuffix;
            InternalFqdn = internalFqdn;
        }
    }
}