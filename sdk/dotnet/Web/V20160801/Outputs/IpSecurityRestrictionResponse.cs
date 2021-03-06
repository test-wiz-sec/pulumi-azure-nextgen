// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNextGen.Web.V20160801.Outputs
{

    [OutputType]
    public sealed class IpSecurityRestrictionResponse
    {
        /// <summary>
        /// IP address the security restriction is valid for.
        /// </summary>
        public readonly string IpAddress;
        /// <summary>
        /// Subnet mask for the range of IP addresses the restriction is valid for.
        /// </summary>
        public readonly string? SubnetMask;

        [OutputConstructor]
        private IpSecurityRestrictionResponse(
            string ipAddress,

            string? subnetMask)
        {
            IpAddress = ipAddress;
            SubnetMask = subnetMask;
        }
    }
}
