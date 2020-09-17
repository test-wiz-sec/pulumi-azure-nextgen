// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNextGen.Search.V20200313.Inputs
{

    /// <summary>
    /// Describes the current state of an existing Private Link Service connection to the Azure Private Endpoint.
    /// </summary>
    public sealed class PrivateEndpointConnectionPropertiesPrivateLinkServiceConnectionStateArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// A description of any extra actions that may be required.
        /// </summary>
        [Input("actionsRequired")]
        public Input<string>? ActionsRequired { get; set; }

        /// <summary>
        /// The description for the private link service connection state.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// Status of the the private link service connection. Can be Pending, Approved, Rejected, or Disconnected.
        /// </summary>
        [Input("status")]
        public Input<string>? Status { get; set; }

        public PrivateEndpointConnectionPropertiesPrivateLinkServiceConnectionStateArgs()
        {
        }
    }
}