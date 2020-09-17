// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNextGen.Devices.V20171115.Inputs
{

    /// <summary>
    /// Description of the shared access key.
    /// </summary>
    public sealed class SharedAccessSignatureAuthorizationRuleAccessRightsDescriptionArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Name of the key.
        /// </summary>
        [Input("keyName", required: true)]
        public Input<string> KeyName { get; set; } = null!;

        /// <summary>
        /// Primary SAS key value.
        /// </summary>
        [Input("primaryKey")]
        public Input<string>? PrimaryKey { get; set; }

        /// <summary>
        /// Rights that this key has.
        /// </summary>
        [Input("rights", required: true)]
        public Input<string> Rights { get; set; } = null!;

        /// <summary>
        /// Secondary SAS key value.
        /// </summary>
        [Input("secondaryKey")]
        public Input<string>? SecondaryKey { get; set; }

        public SharedAccessSignatureAuthorizationRuleAccessRightsDescriptionArgs()
        {
        }
    }
}