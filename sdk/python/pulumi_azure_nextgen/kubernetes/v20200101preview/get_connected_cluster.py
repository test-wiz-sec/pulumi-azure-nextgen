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
    'GetConnectedClusterResult',
    'AwaitableGetConnectedClusterResult',
    'get_connected_cluster',
]

@pulumi.output_type
class GetConnectedClusterResult:
    """
    Represents a connected cluster.
    """
    def __init__(__self__, aad_profile=None, agent_public_key_certificate=None, agent_version=None, identity=None, kubernetes_version=None, location=None, name=None, provisioning_state=None, tags=None, total_node_count=None, type=None):
        if aad_profile and not isinstance(aad_profile, dict):
            raise TypeError("Expected argument 'aad_profile' to be a dict")
        pulumi.set(__self__, "aad_profile", aad_profile)
        if agent_public_key_certificate and not isinstance(agent_public_key_certificate, str):
            raise TypeError("Expected argument 'agent_public_key_certificate' to be a str")
        pulumi.set(__self__, "agent_public_key_certificate", agent_public_key_certificate)
        if agent_version and not isinstance(agent_version, str):
            raise TypeError("Expected argument 'agent_version' to be a str")
        pulumi.set(__self__, "agent_version", agent_version)
        if identity and not isinstance(identity, dict):
            raise TypeError("Expected argument 'identity' to be a dict")
        pulumi.set(__self__, "identity", identity)
        if kubernetes_version and not isinstance(kubernetes_version, str):
            raise TypeError("Expected argument 'kubernetes_version' to be a str")
        pulumi.set(__self__, "kubernetes_version", kubernetes_version)
        if location and not isinstance(location, str):
            raise TypeError("Expected argument 'location' to be a str")
        pulumi.set(__self__, "location", location)
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        pulumi.set(__self__, "name", name)
        if provisioning_state and not isinstance(provisioning_state, str):
            raise TypeError("Expected argument 'provisioning_state' to be a str")
        pulumi.set(__self__, "provisioning_state", provisioning_state)
        if tags and not isinstance(tags, dict):
            raise TypeError("Expected argument 'tags' to be a dict")
        pulumi.set(__self__, "tags", tags)
        if total_node_count and not isinstance(total_node_count, int):
            raise TypeError("Expected argument 'total_node_count' to be a int")
        pulumi.set(__self__, "total_node_count", total_node_count)
        if type and not isinstance(type, str):
            raise TypeError("Expected argument 'type' to be a str")
        pulumi.set(__self__, "type", type)

    @property
    @pulumi.getter(name="aadProfile")
    def aad_profile(self) -> 'outputs.ConnectedClusterAADProfileResponse':
        """
        AAD profile of the connected cluster.
        """
        return pulumi.get(self, "aad_profile")

    @property
    @pulumi.getter(name="agentPublicKeyCertificate")
    def agent_public_key_certificate(self) -> str:
        """
        Base64 encoded public certificate used by the agent to do the initial handshake to the backend services in Azure.
        """
        return pulumi.get(self, "agent_public_key_certificate")

    @property
    @pulumi.getter(name="agentVersion")
    def agent_version(self) -> str:
        """
        Version of the agent running on the connected cluster resource
        """
        return pulumi.get(self, "agent_version")

    @property
    @pulumi.getter
    def identity(self) -> 'outputs.ConnectedClusterIdentityResponse':
        """
        The identity of the connected cluster.
        """
        return pulumi.get(self, "identity")

    @property
    @pulumi.getter(name="kubernetesVersion")
    def kubernetes_version(self) -> str:
        """
        The Kubernetes version of the connected cluster resource
        """
        return pulumi.get(self, "kubernetes_version")

    @property
    @pulumi.getter
    def location(self) -> str:
        """
        The geo-location where the resource lives
        """
        return pulumi.get(self, "location")

    @property
    @pulumi.getter
    def name(self) -> str:
        """
        The name of the resource
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="provisioningState")
    def provisioning_state(self) -> Optional[str]:
        """
        Provisioning state of the connected cluster resource.
        """
        return pulumi.get(self, "provisioning_state")

    @property
    @pulumi.getter
    def tags(self) -> Optional[Mapping[str, str]]:
        """
        Resource tags.
        """
        return pulumi.get(self, "tags")

    @property
    @pulumi.getter(name="totalNodeCount")
    def total_node_count(self) -> int:
        """
        Number of nodes present in the connected cluster resource
        """
        return pulumi.get(self, "total_node_count")

    @property
    @pulumi.getter
    def type(self) -> str:
        """
        The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
        """
        return pulumi.get(self, "type")


class AwaitableGetConnectedClusterResult(GetConnectedClusterResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetConnectedClusterResult(
            aad_profile=self.aad_profile,
            agent_public_key_certificate=self.agent_public_key_certificate,
            agent_version=self.agent_version,
            identity=self.identity,
            kubernetes_version=self.kubernetes_version,
            location=self.location,
            name=self.name,
            provisioning_state=self.provisioning_state,
            tags=self.tags,
            total_node_count=self.total_node_count,
            type=self.type)


def get_connected_cluster(cluster_name: Optional[str] = None,
                          resource_group_name: Optional[str] = None,
                          opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetConnectedClusterResult:
    """
    Use this data source to access information about an existing resource.

    :param str cluster_name: The name of the Kubernetes cluster on which get is called.
    :param str resource_group_name: The name of the resource group. The name is case insensitive.
    """
    __args__ = dict()
    __args__['clusterName'] = cluster_name
    __args__['resourceGroupName'] = resource_group_name
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = _utilities.get_version()
    __ret__ = pulumi.runtime.invoke('azure-nextgen:kubernetes/v20200101preview:getConnectedCluster', __args__, opts=opts, typ=GetConnectedClusterResult).value

    return AwaitableGetConnectedClusterResult(
        aad_profile=__ret__.aad_profile,
        agent_public_key_certificate=__ret__.agent_public_key_certificate,
        agent_version=__ret__.agent_version,
        identity=__ret__.identity,
        kubernetes_version=__ret__.kubernetes_version,
        location=__ret__.location,
        name=__ret__.name,
        provisioning_state=__ret__.provisioning_state,
        tags=__ret__.tags,
        total_node_count=__ret__.total_node_count,
        type=__ret__.type)
