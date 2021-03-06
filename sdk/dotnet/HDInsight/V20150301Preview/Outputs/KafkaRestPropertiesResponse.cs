// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNextGen.HDInsight.V20150301Preview.Outputs
{

    [OutputType]
    public sealed class KafkaRestPropertiesResponse
    {
        /// <summary>
        /// The information of AAD security group.
        /// </summary>
        public readonly Outputs.ClientGroupInfoResponse? ClientGroupInfo;

        [OutputConstructor]
        private KafkaRestPropertiesResponse(Outputs.ClientGroupInfoResponse? clientGroupInfo)
        {
            ClientGroupInfo = clientGroupInfo;
        }
    }
}
