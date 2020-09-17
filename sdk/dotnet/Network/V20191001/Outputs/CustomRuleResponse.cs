// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNextGen.Network.V20191001.Outputs
{

    [OutputType]
    public sealed class CustomRuleResponse
    {
        /// <summary>
        /// Describes what action to be applied when rule matches.
        /// </summary>
        public readonly string Action;
        /// <summary>
        /// Describes if the custom rule is in enabled or disabled state. Defaults to Enabled if not specified.
        /// </summary>
        public readonly string? EnabledState;
        /// <summary>
        /// List of match conditions.
        /// </summary>
        public readonly ImmutableArray<Outputs.MatchConditionResponse> MatchConditions;
        /// <summary>
        /// Describes the name of the rule.
        /// </summary>
        public readonly string? Name;
        /// <summary>
        /// Describes priority of the rule. Rules with a lower value will be evaluated before rules with a higher value.
        /// </summary>
        public readonly int Priority;
        /// <summary>
        /// Time window for resetting the rate limit count. Default is 1 minute.
        /// </summary>
        public readonly int? RateLimitDurationInMinutes;
        /// <summary>
        /// Number of allowed requests per client within the time window.
        /// </summary>
        public readonly int? RateLimitThreshold;
        /// <summary>
        /// Describes type of rule.
        /// </summary>
        public readonly string RuleType;

        [OutputConstructor]
        private CustomRuleResponse(
            string action,

            string? enabledState,

            ImmutableArray<Outputs.MatchConditionResponse> matchConditions,

            string? name,

            int priority,

            int? rateLimitDurationInMinutes,

            int? rateLimitThreshold,

            string ruleType)
        {
            Action = action;
            EnabledState = enabledState;
            MatchConditions = matchConditions;
            Name = name;
            Priority = priority;
            RateLimitDurationInMinutes = rateLimitDurationInMinutes;
            RateLimitThreshold = rateLimitThreshold;
            RuleType = ruleType;
        }
    }
}