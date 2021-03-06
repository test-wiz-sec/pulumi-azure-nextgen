# coding=utf-8
# *** WARNING: this file was generated by the Pulumi SDK Generator. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union
from ... import _utilities, _tables
from . import outputs

__all__ = [
    'ContainerServiceAgentPoolProfileResponse',
    'ContainerServiceDiagnosticsProfileResponse',
    'ContainerServiceLinuxProfileResponse',
    'ContainerServiceMasterProfileResponse',
    'ContainerServiceOrchestratorProfileResponse',
    'ContainerServiceSshConfigurationResponse',
    'ContainerServiceSshPublicKeyResponse',
    'ContainerServiceVMDiagnosticsResponse',
    'ContainerServiceWindowsProfileResponse',
]

@pulumi.output_type
class ContainerServiceAgentPoolProfileResponse(dict):
    """
    Profile for container service agent pool
    """
    def __init__(__self__, *,
                 dns_prefix: str,
                 fqdn: str,
                 name: str,
                 count: Optional[int] = None,
                 vm_size: Optional[str] = None):
        """
        Profile for container service agent pool
        :param str dns_prefix: DNS prefix to be used to create FQDN for this agent pool
        :param str fqdn: FQDN for the agent pool
        :param str name: Unique name of the agent pool profile within the context of the subscription and resource group
        :param int count: No. of agents (VMs) that will host docker containers
        :param str vm_size: Size of agent VMs
        """
        pulumi.set(__self__, "dns_prefix", dns_prefix)
        pulumi.set(__self__, "fqdn", fqdn)
        pulumi.set(__self__, "name", name)
        if count is not None:
            pulumi.set(__self__, "count", count)
        if vm_size is not None:
            pulumi.set(__self__, "vm_size", vm_size)

    @property
    @pulumi.getter(name="dnsPrefix")
    def dns_prefix(self) -> str:
        """
        DNS prefix to be used to create FQDN for this agent pool
        """
        return pulumi.get(self, "dns_prefix")

    @property
    @pulumi.getter
    def fqdn(self) -> str:
        """
        FQDN for the agent pool
        """
        return pulumi.get(self, "fqdn")

    @property
    @pulumi.getter
    def name(self) -> str:
        """
        Unique name of the agent pool profile within the context of the subscription and resource group
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def count(self) -> Optional[int]:
        """
        No. of agents (VMs) that will host docker containers
        """
        return pulumi.get(self, "count")

    @property
    @pulumi.getter(name="vmSize")
    def vm_size(self) -> Optional[str]:
        """
        Size of agent VMs
        """
        return pulumi.get(self, "vm_size")

    def _translate_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop


@pulumi.output_type
class ContainerServiceDiagnosticsProfileResponse(dict):
    def __init__(__self__, *,
                 vm_diagnostics: Optional['outputs.ContainerServiceVMDiagnosticsResponse'] = None):
        """
        :param 'ContainerServiceVMDiagnosticsResponseArgs' vm_diagnostics: Profile for container service VM diagnostic agent
        """
        if vm_diagnostics is not None:
            pulumi.set(__self__, "vm_diagnostics", vm_diagnostics)

    @property
    @pulumi.getter(name="vmDiagnostics")
    def vm_diagnostics(self) -> Optional['outputs.ContainerServiceVMDiagnosticsResponse']:
        """
        Profile for container service VM diagnostic agent
        """
        return pulumi.get(self, "vm_diagnostics")

    def _translate_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop


@pulumi.output_type
class ContainerServiceLinuxProfileResponse(dict):
    """
    Profile for Linux VM
    """
    def __init__(__self__, *,
                 admin_username: str,
                 ssh: 'outputs.ContainerServiceSshConfigurationResponse'):
        """
        Profile for Linux VM
        :param str admin_username: The administrator username to use for all Linux VMs
        :param 'ContainerServiceSshConfigurationResponseArgs' ssh: Specifies the ssh key configuration for Linux VMs
        """
        pulumi.set(__self__, "admin_username", admin_username)
        pulumi.set(__self__, "ssh", ssh)

    @property
    @pulumi.getter(name="adminUsername")
    def admin_username(self) -> str:
        """
        The administrator username to use for all Linux VMs
        """
        return pulumi.get(self, "admin_username")

    @property
    @pulumi.getter
    def ssh(self) -> 'outputs.ContainerServiceSshConfigurationResponse':
        """
        Specifies the ssh key configuration for Linux VMs
        """
        return pulumi.get(self, "ssh")

    def _translate_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop


@pulumi.output_type
class ContainerServiceMasterProfileResponse(dict):
    """
    Profile for container service master
    """
    def __init__(__self__, *,
                 dns_prefix: str,
                 fqdn: str,
                 count: Optional[int] = None):
        """
        Profile for container service master
        :param str dns_prefix: DNS prefix to be used to create FQDN for master
        :param str fqdn: FQDN for the master
        :param int count: Number of masters (VMs) in the container cluster
        """
        pulumi.set(__self__, "dns_prefix", dns_prefix)
        pulumi.set(__self__, "fqdn", fqdn)
        if count is not None:
            pulumi.set(__self__, "count", count)

    @property
    @pulumi.getter(name="dnsPrefix")
    def dns_prefix(self) -> str:
        """
        DNS prefix to be used to create FQDN for master
        """
        return pulumi.get(self, "dns_prefix")

    @property
    @pulumi.getter
    def fqdn(self) -> str:
        """
        FQDN for the master
        """
        return pulumi.get(self, "fqdn")

    @property
    @pulumi.getter
    def count(self) -> Optional[int]:
        """
        Number of masters (VMs) in the container cluster
        """
        return pulumi.get(self, "count")

    def _translate_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop


@pulumi.output_type
class ContainerServiceOrchestratorProfileResponse(dict):
    """
    Profile for Orchestrator
    """
    def __init__(__self__, *,
                 orchestrator_type: Optional[str] = None):
        """
        Profile for Orchestrator
        :param str orchestrator_type: Specifies what orchestrator will be used to manage container cluster resources.
        """
        if orchestrator_type is not None:
            pulumi.set(__self__, "orchestrator_type", orchestrator_type)

    @property
    @pulumi.getter(name="orchestratorType")
    def orchestrator_type(self) -> Optional[str]:
        """
        Specifies what orchestrator will be used to manage container cluster resources.
        """
        return pulumi.get(self, "orchestrator_type")

    def _translate_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop


@pulumi.output_type
class ContainerServiceSshConfigurationResponse(dict):
    """
    SSH configuration for Linux based VMs running on Azure
    """
    def __init__(__self__, *,
                 public_keys: Optional[Sequence['outputs.ContainerServiceSshPublicKeyResponse']] = None):
        """
        SSH configuration for Linux based VMs running on Azure
        :param Sequence['ContainerServiceSshPublicKeyResponseArgs'] public_keys: Gets or sets the list of SSH public keys used to authenticate with Linux based VMs
        """
        if public_keys is not None:
            pulumi.set(__self__, "public_keys", public_keys)

    @property
    @pulumi.getter(name="publicKeys")
    def public_keys(self) -> Optional[Sequence['outputs.ContainerServiceSshPublicKeyResponse']]:
        """
        Gets or sets the list of SSH public keys used to authenticate with Linux based VMs
        """
        return pulumi.get(self, "public_keys")

    def _translate_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop


@pulumi.output_type
class ContainerServiceSshPublicKeyResponse(dict):
    """
    Contains information about SSH certificate public key data.
    """
    def __init__(__self__, *,
                 key_data: str):
        """
        Contains information about SSH certificate public key data.
        :param str key_data: Gets or sets Certificate public key used to authenticate with VM through SSH. The certificate must be in Pem format with or without headers.
        """
        pulumi.set(__self__, "key_data", key_data)

    @property
    @pulumi.getter(name="keyData")
    def key_data(self) -> str:
        """
        Gets or sets Certificate public key used to authenticate with VM through SSH. The certificate must be in Pem format with or without headers.
        """
        return pulumi.get(self, "key_data")

    def _translate_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop


@pulumi.output_type
class ContainerServiceVMDiagnosticsResponse(dict):
    """
    Describes VM Diagnostics.
    """
    def __init__(__self__, *,
                 storage_uri: str,
                 enabled: Optional[bool] = None):
        """
        Describes VM Diagnostics.
        :param str storage_uri: Gets or sets whether VM Diagnostic Agent should be provisioned on the Virtual Machine.
        :param bool enabled: Gets or sets whether VM Diagnostic Agent should be provisioned on the Virtual Machine.
        """
        pulumi.set(__self__, "storage_uri", storage_uri)
        if enabled is not None:
            pulumi.set(__self__, "enabled", enabled)

    @property
    @pulumi.getter(name="storageUri")
    def storage_uri(self) -> str:
        """
        Gets or sets whether VM Diagnostic Agent should be provisioned on the Virtual Machine.
        """
        return pulumi.get(self, "storage_uri")

    @property
    @pulumi.getter
    def enabled(self) -> Optional[bool]:
        """
        Gets or sets whether VM Diagnostic Agent should be provisioned on the Virtual Machine.
        """
        return pulumi.get(self, "enabled")

    def _translate_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop


@pulumi.output_type
class ContainerServiceWindowsProfileResponse(dict):
    """
    Profile for Windows jumpbox
    """
    def __init__(__self__, *,
                 admin_password: str,
                 admin_username: str):
        """
        Profile for Windows jumpbox
        :param str admin_password: The administrator password to use for Windows jumpbox
        :param str admin_username: The administrator username to use for Windows jumpbox
        """
        pulumi.set(__self__, "admin_password", admin_password)
        pulumi.set(__self__, "admin_username", admin_username)

    @property
    @pulumi.getter(name="adminPassword")
    def admin_password(self) -> str:
        """
        The administrator password to use for Windows jumpbox
        """
        return pulumi.get(self, "admin_password")

    @property
    @pulumi.getter(name="adminUsername")
    def admin_username(self) -> str:
        """
        The administrator username to use for Windows jumpbox
        """
        return pulumi.get(self, "admin_username")

    def _translate_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop


