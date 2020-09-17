// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNextGen.CustomerInsights.V20170426.Outputs
{

    [OutputType]
    public sealed class ConnectorMappingFormatResponse
    {
        /// <summary>
        /// The oData language.
        /// </summary>
        public readonly string? AcceptLanguage;
        /// <summary>
        /// Character separating array elements.
        /// </summary>
        public readonly string? ArraySeparator;
        /// <summary>
        /// The character that signifies a break between columns.
        /// </summary>
        public readonly string? ColumnDelimiter;
        /// <summary>
        /// The type mapping format.
        /// </summary>
        public readonly string FormatType;
        /// <summary>
        /// Quote character, used to indicate enquoted fields.
        /// </summary>
        public readonly string? QuoteCharacter;
        /// <summary>
        /// Escape character for quotes, can be the same as the quoteCharacter.
        /// </summary>
        public readonly string? QuoteEscapeCharacter;

        [OutputConstructor]
        private ConnectorMappingFormatResponse(
            string? acceptLanguage,

            string? arraySeparator,

            string? columnDelimiter,

            string formatType,

            string? quoteCharacter,

            string? quoteEscapeCharacter)
        {
            AcceptLanguage = acceptLanguage;
            ArraySeparator = arraySeparator;
            ColumnDelimiter = columnDelimiter;
            FormatType = formatType;
            QuoteCharacter = quoteCharacter;
            QuoteEscapeCharacter = quoteEscapeCharacter;
        }
    }
}