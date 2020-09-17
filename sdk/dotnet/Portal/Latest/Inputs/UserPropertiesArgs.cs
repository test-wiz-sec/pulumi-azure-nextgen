// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNextGen.Portal.Latest.Inputs
{

    /// <summary>
    /// The cloud shell user settings properties.
    /// </summary>
    public sealed class UserPropertiesArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The preferred location of the cloud shell.
        /// </summary>
        [Input("preferredLocation", required: true)]
        public Input<string> PreferredLocation { get; set; } = null!;

        /// <summary>
        /// The operating system type of the cloud shell. Deprecated, use preferredShellType.
        /// </summary>
        [Input("preferredOsType", required: true)]
        public Input<string> PreferredOsType { get; set; } = null!;

        /// <summary>
        /// The shell type of the cloud shell.
        /// </summary>
        [Input("preferredShellType", required: true)]
        public Input<string> PreferredShellType { get; set; } = null!;

        /// <summary>
        /// The storage profile of the user settings.
        /// </summary>
        [Input("storageProfile", required: true)]
        public Input<Inputs.StorageProfileArgs> StorageProfile { get; set; } = null!;

        /// <summary>
        /// Settings for terminal appearance.
        /// </summary>
        [Input("terminalSettings", required: true)]
        public Input<Inputs.TerminalSettingsArgs> TerminalSettings { get; set; } = null!;

        public UserPropertiesArgs()
        {
        }
    }
}