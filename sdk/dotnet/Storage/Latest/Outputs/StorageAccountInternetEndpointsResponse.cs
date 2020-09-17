// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNextGen.Storage.Latest.Outputs
{

    [OutputType]
    public sealed class StorageAccountInternetEndpointsResponse
    {
        /// <summary>
        /// Gets the blob endpoint.
        /// </summary>
        public readonly string Blob;
        /// <summary>
        /// Gets the dfs endpoint.
        /// </summary>
        public readonly string Dfs;
        /// <summary>
        /// Gets the file endpoint.
        /// </summary>
        public readonly string File;
        /// <summary>
        /// Gets the web endpoint.
        /// </summary>
        public readonly string Web;

        [OutputConstructor]
        private StorageAccountInternetEndpointsResponse(
            string blob,

            string dfs,

            string file,

            string web)
        {
            Blob = blob;
            Dfs = dfs;
            File = file;
            Web = web;
        }
    }
}