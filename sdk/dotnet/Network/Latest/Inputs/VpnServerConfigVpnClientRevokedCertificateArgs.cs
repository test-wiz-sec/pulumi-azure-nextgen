// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNextGen.Network.Latest.Inputs
{

    /// <summary>
    /// Properties of the revoked VPN client certificate of VpnServerConfiguration.
    /// </summary>
    public sealed class VpnServerConfigVpnClientRevokedCertificateArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The certificate name.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The revoked VPN client certificate thumbprint.
        /// </summary>
        [Input("thumbprint")]
        public Input<string>? Thumbprint { get; set; }

        public VpnServerConfigVpnClientRevokedCertificateArgs()
        {
        }
    }
}