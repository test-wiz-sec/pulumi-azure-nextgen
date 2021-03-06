// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNextGen.BatchAI.V20170901Preview.Outputs
{

    [OutputType]
    public sealed class JobPreparationResponse
    {
        /// <summary>
        /// If containerSettings is specified on the job, this commandLine will be executed in the same container as job. Otherwise it will be executed on the node.
        /// </summary>
        public readonly string CommandLine;

        [OutputConstructor]
        private JobPreparationResponse(string commandLine)
        {
            CommandLine = commandLine;
        }
    }
}
