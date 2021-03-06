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
    public sealed class ManagementPolicyActionResponse
    {
        /// <summary>
        /// The management policy action for base blob
        /// </summary>
        public readonly Outputs.ManagementPolicyBaseBlobResponse? BaseBlob;
        /// <summary>
        /// The management policy action for snapshot
        /// </summary>
        public readonly Outputs.ManagementPolicySnapShotResponse? Snapshot;

        [OutputConstructor]
        private ManagementPolicyActionResponse(
            Outputs.ManagementPolicyBaseBlobResponse? baseBlob,

            Outputs.ManagementPolicySnapShotResponse? snapshot)
        {
            BaseBlob = baseBlob;
            Snapshot = snapshot;
        }
    }
}
