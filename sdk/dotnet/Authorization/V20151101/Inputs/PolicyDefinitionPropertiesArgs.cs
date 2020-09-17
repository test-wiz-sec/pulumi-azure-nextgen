// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNextGen.Authorization.V20151101.Inputs
{

    /// <summary>
    /// Policy definition properties.
    /// </summary>
    public sealed class PolicyDefinitionPropertiesArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Gets or sets the policy definition description.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// Gets or sets the policy definition display name.
        /// </summary>
        [Input("displayName")]
        public Input<string>? DisplayName { get; set; }

        [Input("policyRule")]
        private InputMap<object>? _policyRule;

        /// <summary>
        /// The policy rule json.
        /// </summary>
        public InputMap<object> PolicyRule
        {
            get => _policyRule ?? (_policyRule = new InputMap<object>());
            set => _policyRule = value;
        }

        public PolicyDefinitionPropertiesArgs()
        {
        }
    }
}