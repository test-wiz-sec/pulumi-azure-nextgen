// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNextGen.Devices.V20180122.Inputs
{

    /// <summary>
    /// The operations monitoring properties for the IoT hub. The possible keys to the dictionary are Connections, DeviceTelemetry, C2DCommands, DeviceIdentityOperations, FileUploadOperations, Routes, D2CTwinOperations, C2DTwinOperations, TwinQueries, JobsOperations, DirectMethods.
    /// </summary>
    public sealed class OperationsMonitoringPropertiesArgs : Pulumi.ResourceArgs
    {
        [Input("events")]
        private InputMap<string>? _events;
        public InputMap<string> Events
        {
            get => _events ?? (_events = new InputMap<string>());
            set => _events = value;
        }

        public OperationsMonitoringPropertiesArgs()
        {
        }
    }
}
