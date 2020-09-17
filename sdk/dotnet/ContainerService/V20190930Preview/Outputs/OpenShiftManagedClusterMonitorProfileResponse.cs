// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNextGen.ContainerService.V20190930Preview.Outputs
{

    [OutputType]
    public sealed class OpenShiftManagedClusterMonitorProfileResponse
    {
        /// <summary>
        /// If the Log analytics integration should be turned on or off
        /// </summary>
        public readonly bool? Enabled;
        /// <summary>
        /// Azure Resource Manager Resource ID for the Log Analytics workspace to integrate with.
        /// </summary>
        public readonly string? WorkspaceResourceID;

        [OutputConstructor]
        private OpenShiftManagedClusterMonitorProfileResponse(
            bool? enabled,

            string? workspaceResourceID)
        {
            Enabled = enabled;
            WorkspaceResourceID = workspaceResourceID;
        }
    }
}