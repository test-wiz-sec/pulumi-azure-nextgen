// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNextGen.EventGrid.Latest.Inputs
{

    public sealed class PrivateEndpointConnectionArgs : Pulumi.ResourceArgs
    {
        [Input("groupIds")]
        private InputList<string>? _groupIds;

        /// <summary>
        /// GroupIds from the private link service resource.
        /// </summary>
        public InputList<string> GroupIds
        {
            get => _groupIds ?? (_groupIds = new InputList<string>());
            set => _groupIds = value;
        }

        /// <summary>
        /// The Private Endpoint resource for this Connection.
        /// </summary>
        [Input("privateEndpoint")]
        public Input<Inputs.PrivateEndpointArgs>? PrivateEndpoint { get; set; }

        /// <summary>
        /// Details about the state of the connection.
        /// </summary>
        [Input("privateLinkServiceConnectionState")]
        public Input<Inputs.ConnectionStateArgs>? PrivateLinkServiceConnectionState { get; set; }

        /// <summary>
        /// Provisioning state of the Private Endpoint Connection.
        /// </summary>
        [Input("provisioningState")]
        public Input<string>? ProvisioningState { get; set; }

        public PrivateEndpointConnectionArgs()
        {
        }
    }
}