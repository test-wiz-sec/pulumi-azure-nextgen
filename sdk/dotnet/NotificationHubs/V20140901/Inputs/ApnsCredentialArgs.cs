// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNextGen.NotificationHubs.V20140901.Inputs
{

    /// <summary>
    /// Description of a NotificationHub ApnsCredential.
    /// </summary>
    public sealed class ApnsCredentialArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Gets or sets properties of NotificationHub ApnsCredential.
        /// </summary>
        [Input("properties")]
        public Input<Inputs.ApnsCredentialPropertiesArgs>? Properties { get; set; }

        public ApnsCredentialArgs()
        {
        }
    }
}
