// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNextGen.PowerPlatform.V20201030Preview.Inputs
{

    /// <summary>
    /// Identity for the EnterprisePolicy.
    /// </summary>
    public sealed class EnterprisePolicyIdentityArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The encryption settings for a configuration store.
        /// </summary>
        [Input("encryption")]
        public Input<Inputs.EnterprisePolicyIdentityEncryptionArgs>? Encryption { get; set; }

        /// <summary>
        /// Settings concerning lockbox.
        /// </summary>
        [Input("lockbox")]
        public Input<Inputs.EnterprisePolicyIdentityLockboxArgs>? Lockbox { get; set; }

        /// <summary>
        /// Metadata for the enterprisePolicy.
        /// </summary>
        [Input("systemData")]
        public Input<Inputs.SystemDataArgs>? SystemData { get; set; }

        /// <summary>
        /// The type of identity used for the EnterprisePolicy. Currently, the only supported type is 'SystemAssigned', which implicitly creates an identity.
        /// </summary>
        [Input("type")]
        public Input<string>? Type { get; set; }

        public EnterprisePolicyIdentityArgs()
        {
        }
    }
}