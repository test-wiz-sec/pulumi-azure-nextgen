// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNextGen.Authorization.V20180101Preview.Inputs
{

    /// <summary>
    /// Role definition permissions.
    /// </summary>
    public sealed class PermissionArgs : Pulumi.ResourceArgs
    {
        [Input("actions")]
        private InputList<string>? _actions;

        /// <summary>
        /// Allowed actions.
        /// </summary>
        public InputList<string> Actions
        {
            get => _actions ?? (_actions = new InputList<string>());
            set => _actions = value;
        }

        [Input("dataActions")]
        private InputList<string>? _dataActions;

        /// <summary>
        /// Allowed Data actions.
        /// </summary>
        public InputList<string> DataActions
        {
            get => _dataActions ?? (_dataActions = new InputList<string>());
            set => _dataActions = value;
        }

        [Input("notActions")]
        private InputList<string>? _notActions;

        /// <summary>
        /// Denied actions.
        /// </summary>
        public InputList<string> NotActions
        {
            get => _notActions ?? (_notActions = new InputList<string>());
            set => _notActions = value;
        }

        [Input("notDataActions")]
        private InputList<string>? _notDataActions;

        /// <summary>
        /// Denied Data actions.
        /// </summary>
        public InputList<string> NotDataActions
        {
            get => _notDataActions ?? (_notDataActions = new InputList<string>());
            set => _notDataActions = value;
        }

        public PermissionArgs()
        {
        }
    }
}