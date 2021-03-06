// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNextGen.Compute.V20191201.Inputs
{

    /// <summary>
    /// Identity for the virtual machine scale set.
    /// </summary>
    public sealed class VirtualMachineScaleSetIdentityArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The type of identity used for the virtual machine scale set. The type 'SystemAssigned, UserAssigned' includes both an implicitly created identity and a set of user assigned identities. The type 'None' will remove any identities from the virtual machine scale set.
        /// </summary>
        [Input("type")]
        public Input<string>? Type { get; set; }

        public VirtualMachineScaleSetIdentityArgs()
        {
        }
    }
}
