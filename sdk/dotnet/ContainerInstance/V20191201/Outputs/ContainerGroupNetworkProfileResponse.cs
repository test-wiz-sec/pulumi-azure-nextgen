// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNextGen.ContainerInstance.V20191201.Outputs
{

    [OutputType]
    public sealed class ContainerGroupNetworkProfileResponse
    {
        /// <summary>
        /// The identifier for a network profile.
        /// </summary>
        public readonly string Id;

        [OutputConstructor]
        private ContainerGroupNetworkProfileResponse(string id)
        {
            Id = id;
        }
    }
}