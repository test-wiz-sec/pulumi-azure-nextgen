// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNextGen.Compute.V20200601.Outputs
{

    [OutputType]
    public sealed class VirtualMachineAgentInstanceViewResponse
    {
        /// <summary>
        /// The virtual machine extension handler instance view.
        /// </summary>
        public readonly ImmutableArray<Outputs.VirtualMachineExtensionHandlerInstanceViewResponse> ExtensionHandlers;
        /// <summary>
        /// The resource status information.
        /// </summary>
        public readonly ImmutableArray<Outputs.InstanceViewStatusResponse> Statuses;
        /// <summary>
        /// The VM Agent full version.
        /// </summary>
        public readonly string? VmAgentVersion;

        [OutputConstructor]
        private VirtualMachineAgentInstanceViewResponse(
            ImmutableArray<Outputs.VirtualMachineExtensionHandlerInstanceViewResponse> extensionHandlers,

            ImmutableArray<Outputs.InstanceViewStatusResponse> statuses,

            string? vmAgentVersion)
        {
            ExtensionHandlers = extensionHandlers;
            Statuses = statuses;
            VmAgentVersion = vmAgentVersion;
        }
    }
}