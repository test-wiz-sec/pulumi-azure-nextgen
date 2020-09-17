// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNextGen.BatchAI.V20170901Preview.Inputs
{

    /// <summary>
    /// Use this to prepare the VM. NOTE: The volumes specified in mountVolumes are mounted first and then the setupTask is run. Therefore the setup task can use local mountPaths in its execution.
    /// </summary>
    public sealed class NodeSetupArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Details of volumes to mount on the cluster.
        /// </summary>
        [Input("mountVolumes")]
        public Input<Inputs.MountVolumesArgs>? MountVolumes { get; set; }

        /// <summary>
        /// Specifies a setup task which can be used to customize the compute nodes of the cluster.
        /// </summary>
        [Input("setupTask")]
        public Input<Inputs.SetupTaskArgs>? SetupTask { get; set; }

        public NodeSetupArgs()
        {
        }
    }
}