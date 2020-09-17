// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNextGen.Management.V20191001.Outputs
{

    [OutputType]
    public sealed class AliasPathResponse
    {
        /// <summary>
        /// The API versions.
        /// </summary>
        public readonly ImmutableArray<string> ApiVersions;
        /// <summary>
        /// The path of an alias.
        /// </summary>
        public readonly string? Path;
        /// <summary>
        /// The pattern for an alias path.
        /// </summary>
        public readonly Outputs.AliasPatternResponse? Pattern;

        [OutputConstructor]
        private AliasPathResponse(
            ImmutableArray<string> apiVersions,

            string? path,

            Outputs.AliasPatternResponse? pattern)
        {
            ApiVersions = apiVersions;
            Path = path;
            Pattern = pattern;
        }
    }
}