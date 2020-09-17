// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNextGen.Compute.V20160330.Inputs
{

    /// <summary>
    /// Specifies the Linux operating system settings on the virtual machine. &lt;br&gt;&lt;br&gt;For a list of supported Linux distributions, see [Linux on Azure-Endorsed Distributions](https://docs.microsoft.com/azure/virtual-machines/virtual-machines-linux-endorsed-distros?toc=%2fazure%2fvirtual-machines%2flinux%2ftoc.json) &lt;br&gt;&lt;br&gt; For running non-endorsed distributions, see [Information for Non-Endorsed Distributions](https://docs.microsoft.com/azure/virtual-machines/virtual-machines-linux-create-upload-generic?toc=%2fazure%2fvirtual-machines%2flinux%2ftoc.json).
    /// </summary>
    public sealed class LinuxConfigurationArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Specifies whether password authentication should be disabled.
        /// </summary>
        [Input("disablePasswordAuthentication")]
        public Input<bool>? DisablePasswordAuthentication { get; set; }

        /// <summary>
        /// Specifies the ssh key configuration for a Linux OS.
        /// </summary>
        [Input("ssh")]
        public Input<Inputs.SshConfigurationArgs>? Ssh { get; set; }

        public LinuxConfigurationArgs()
        {
        }
    }
}