// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNextGen.Compute.V20200601.Inputs
{

    /// <summary>
    /// Identity for the virtual machine.
    /// </summary>
    public sealed class VirtualMachineIdentityArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The type of identity used for the virtual machine. The type 'SystemAssigned, UserAssigned' includes both an implicitly created identity and a set of user assigned identities. The type 'None' will remove any identities from the virtual machine.
        /// </summary>
        [Input("type")]
        public Input<string>? Type { get; set; }

        public VirtualMachineIdentityArgs()
        {
        }
    }
}
