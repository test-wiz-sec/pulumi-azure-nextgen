// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNextGen.DBforPostgreSQL.V20171201Preview.Inputs
{

    /// <summary>
    /// The properties used to create a new server by restoring from a backup.
    /// </summary>
    public sealed class ServerPropertiesForRestoreArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The mode to create a new server.
        /// </summary>
        [Input("createMode", required: true)]
        public Input<string> CreateMode { get; set; } = null!;

        /// <summary>
        /// Enforce a minimal Tls version for the server.
        /// </summary>
        [Input("minimalTlsVersion")]
        public Input<string>? MinimalTlsVersion { get; set; }

        /// <summary>
        /// Restore point creation time (ISO8601 format), specifying the time to restore from.
        /// </summary>
        [Input("restorePointInTime", required: true)]
        public Input<string> RestorePointInTime { get; set; } = null!;

        /// <summary>
        /// The source server id to restore from.
        /// </summary>
        [Input("sourceServerId", required: true)]
        public Input<string> SourceServerId { get; set; } = null!;

        /// <summary>
        /// Enable ssl enforcement or not when connect to server.
        /// </summary>
        [Input("sslEnforcement")]
        public Input<string>? SslEnforcement { get; set; }

        /// <summary>
        /// Storage profile of a server.
        /// </summary>
        [Input("storageProfile")]
        public Input<Inputs.StorageProfileArgs>? StorageProfile { get; set; }

        /// <summary>
        /// Server version.
        /// </summary>
        [Input("version")]
        public Input<string>? Version { get; set; }

        public ServerPropertiesForRestoreArgs()
        {
        }
    }
}
