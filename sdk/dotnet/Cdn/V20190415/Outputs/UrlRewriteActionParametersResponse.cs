// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNextGen.Cdn.V20190415.Outputs
{

    [OutputType]
    public sealed class UrlRewriteActionParametersResponse
    {
        /// <summary>
        /// Define the destination path for be used in the rewrite. This will overwrite the source pattern 
        /// </summary>
        public readonly string Destination;
        public readonly string OdataType;
        /// <summary>
        /// If True, the remaining path after the source pattern will be appended to the new destination path.  
        /// </summary>
        public readonly bool? PreserveUnmatchedPath;
        /// <summary>
        /// define a request URI pattern that identifies the type of requests that may be rewritten. Currently, source pattern uses a prefix-based match. To match all URL paths, use "/" as the source pattern value. To match only the root directory and re-write this path, use the origin path field
        /// </summary>
        public readonly string SourcePattern;

        [OutputConstructor]
        private UrlRewriteActionParametersResponse(
            string destination,

            string odataType,

            bool? preserveUnmatchedPath,

            string sourcePattern)
        {
            Destination = destination;
            OdataType = odataType;
            PreserveUnmatchedPath = preserveUnmatchedPath;
            SourcePattern = sourcePattern;
        }
    }
}