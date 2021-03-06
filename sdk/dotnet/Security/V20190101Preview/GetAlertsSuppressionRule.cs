// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNextGen.Security.V20190101Preview
{
    public static class GetAlertsSuppressionRule
    {
        public static Task<GetAlertsSuppressionRuleResult> InvokeAsync(GetAlertsSuppressionRuleArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetAlertsSuppressionRuleResult>("azure-nextgen:security/v20190101preview:getAlertsSuppressionRule", args ?? new GetAlertsSuppressionRuleArgs(), options.WithVersion());
    }


    public sealed class GetAlertsSuppressionRuleArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// The unique name of the suppression alert rule
        /// </summary>
        [Input("alertsSuppressionRuleName", required: true)]
        public string AlertsSuppressionRuleName { get; set; } = null!;

        public GetAlertsSuppressionRuleArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetAlertsSuppressionRuleResult
    {
        /// <summary>
        /// Type of the alert to automatically suppress. For all alert types, use '*'
        /// </summary>
        public readonly string AlertType;
        /// <summary>
        /// Any comment regarding the rule
        /// </summary>
        public readonly string? Comment;
        /// <summary>
        /// Expiration date of the rule, if value is not provided or provided as null this field will default to the maximum allowed expiration date.
        /// </summary>
        public readonly string? ExpirationDateUtc;
        /// <summary>
        /// The last time this rule was modified
        /// </summary>
        public readonly string LastModifiedUtc;
        /// <summary>
        /// Resource name
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// The reason for dismissing the alert
        /// </summary>
        public readonly string Reason;
        /// <summary>
        /// Possible states of the rule
        /// </summary>
        public readonly string State;
        /// <summary>
        /// The suppression conditions
        /// </summary>
        public readonly Outputs.SuppressionAlertsScopeResponse? SuppressionAlertsScope;
        /// <summary>
        /// Resource type
        /// </summary>
        public readonly string Type;

        [OutputConstructor]
        private GetAlertsSuppressionRuleResult(
            string alertType,

            string? comment,

            string? expirationDateUtc,

            string lastModifiedUtc,

            string name,

            string reason,

            string state,

            Outputs.SuppressionAlertsScopeResponse? suppressionAlertsScope,

            string type)
        {
            AlertType = alertType;
            Comment = comment;
            ExpirationDateUtc = expirationDateUtc;
            LastModifiedUtc = lastModifiedUtc;
            Name = name;
            Reason = reason;
            State = state;
            SuppressionAlertsScope = suppressionAlertsScope;
            Type = type;
        }
    }
}
