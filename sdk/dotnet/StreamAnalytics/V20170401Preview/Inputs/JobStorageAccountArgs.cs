// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNextGen.StreamAnalytics.V20170401Preview.Inputs
{

    /// <summary>
    /// The properties that are associated with an Azure Storage account with MSI
    /// </summary>
    public sealed class JobStorageAccountArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The account key for the Azure Storage account. Required on PUT (CreateOrReplace) requests.
        /// </summary>
        [Input("accountKey")]
        public Input<string>? AccountKey { get; set; }

        /// <summary>
        /// The name of the Azure Storage account. Required on PUT (CreateOrReplace) requests.
        /// </summary>
        [Input("accountName")]
        public Input<string>? AccountName { get; set; }

        /// <summary>
        /// Authentication Mode.
        /// </summary>
        [Input("authenticationMode")]
        public Input<string>? AuthenticationMode { get; set; }

        public JobStorageAccountArgs()
        {
        }
    }
}
