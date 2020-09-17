// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNextGen.Network.V20190201.Outputs
{

    [OutputType]
    public sealed class PolicySettingsResponse
    {
        /// <summary>
        /// Describes if the policy is in enabled state or disabled state
        /// </summary>
        public readonly string? EnabledState;
        /// <summary>
        /// Describes if it is in detection mode  or prevention mode at policy level
        /// </summary>
        public readonly string? Mode;

        [OutputConstructor]
        private PolicySettingsResponse(
            string? enabledState,

            string? mode)
        {
            EnabledState = enabledState;
            Mode = mode;
        }
    }
}