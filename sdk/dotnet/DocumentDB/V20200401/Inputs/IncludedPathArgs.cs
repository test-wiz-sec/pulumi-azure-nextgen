// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNextGen.DocumentDB.V20200401.Inputs
{

    /// <summary>
    /// The paths that are included in indexing
    /// </summary>
    public sealed class IncludedPathArgs : Pulumi.ResourceArgs
    {
        [Input("indexes")]
        private InputList<Inputs.IndexesArgs>? _indexes;

        /// <summary>
        /// List of indexes for this path
        /// </summary>
        public InputList<Inputs.IndexesArgs> Indexes
        {
            get => _indexes ?? (_indexes = new InputList<Inputs.IndexesArgs>());
            set => _indexes = value;
        }

        /// <summary>
        /// The path for which the indexing behavior applies to. Index paths typically start with root and end with wildcard (/path/*)
        /// </summary>
        [Input("path")]
        public Input<string>? Path { get; set; }

        public IncludedPathArgs()
        {
        }
    }
}