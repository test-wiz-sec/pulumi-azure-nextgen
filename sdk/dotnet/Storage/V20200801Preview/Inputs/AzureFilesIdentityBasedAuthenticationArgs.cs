// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNextGen.Storage.V20200801Preview.Inputs
{

    /// <summary>
    /// Settings for Azure Files identity based authentication.
    /// </summary>
    public sealed class AzureFilesIdentityBasedAuthenticationArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Required if choose AD.
        /// </summary>
        [Input("activeDirectoryProperties")]
        public Input<Inputs.ActiveDirectoryPropertiesArgs>? ActiveDirectoryProperties { get; set; }

        /// <summary>
        /// Indicates the directory service used.
        /// </summary>
        [Input("directoryServiceOptions", required: true)]
        public Input<string> DirectoryServiceOptions { get; set; } = null!;

        public AzureFilesIdentityBasedAuthenticationArgs()
        {
        }
    }
}