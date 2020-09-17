// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNextGen.AlertsManagement.V20181102PrivatePreview.Inputs
{

    /// <summary>
    /// condition to trigger an action rule
    /// </summary>
    public sealed class ConditionArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// operator for a given condition
        /// </summary>
        [Input("operator")]
        public Input<string>? Operator { get; set; }

        [Input("values")]
        private InputList<string>? _values;

        /// <summary>
        /// list of values to match for a given condition.
        /// </summary>
        public InputList<string> Values
        {
            get => _values ?? (_values = new InputList<string>());
            set => _values = value;
        }

        public ConditionArgs()
        {
        }
    }
}