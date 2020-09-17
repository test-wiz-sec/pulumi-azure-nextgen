// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNextGen.Security.V20190101Preview.Inputs
{

    /// <summary>
    /// A rule set which evaluates all its rules upon an event interception. Only when all the included rules in the rule set will be evaluated as 'true', will the event trigger the defined actions.
    /// </summary>
    public sealed class AutomationRuleSetArgs : Pulumi.ResourceArgs
    {
        [Input("rules")]
        private InputList<Inputs.AutomationTriggeringRuleArgs>? _rules;
        public InputList<Inputs.AutomationTriggeringRuleArgs> Rules
        {
            get => _rules ?? (_rules = new InputList<Inputs.AutomationTriggeringRuleArgs>());
            set => _rules = value;
        }

        public AutomationRuleSetArgs()
        {
        }
    }
}