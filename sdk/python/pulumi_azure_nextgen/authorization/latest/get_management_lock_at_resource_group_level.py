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
    'GetManagementLockAtResourceGroupLevelResult',
    'AwaitableGetManagementLockAtResourceGroupLevelResult',
    'get_management_lock_at_resource_group_level',
]

@pulumi.output_type
class GetManagementLockAtResourceGroupLevelResult:
    """
    The lock information.
    """
    def __init__(__self__, level=None, name=None, notes=None, owners=None, type=None):
        if level and not isinstance(level, str):
            raise TypeError("Expected argument 'level' to be a str")
        pulumi.set(__self__, "level", level)
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        pulumi.set(__self__, "name", name)
        if notes and not isinstance(notes, str):
            raise TypeError("Expected argument 'notes' to be a str")
        pulumi.set(__self__, "notes", notes)
        if owners and not isinstance(owners, list):
            raise TypeError("Expected argument 'owners' to be a list")
        pulumi.set(__self__, "owners", owners)
        if type and not isinstance(type, str):
            raise TypeError("Expected argument 'type' to be a str")
        pulumi.set(__self__, "type", type)

    @property
    @pulumi.getter
    def level(self) -> str:
        """
        The level of the lock. Possible values are: NotSpecified, CanNotDelete, ReadOnly. CanNotDelete means authorized users are able to read and modify the resources, but not delete. ReadOnly means authorized users can only read from a resource, but they can't modify or delete it.
        """
        return pulumi.get(self, "level")

    @property
    @pulumi.getter
    def name(self) -> str:
        """
        The name of the lock.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def notes(self) -> Optional[str]:
        """
        Notes about the lock. Maximum of 512 characters.
        """
        return pulumi.get(self, "notes")

    @property
    @pulumi.getter
    def owners(self) -> Optional[Sequence['outputs.ManagementLockOwnerResponse']]:
        """
        The owners of the lock.
        """
        return pulumi.get(self, "owners")

    @property
    @pulumi.getter
    def type(self) -> str:
        """
        The resource type of the lock - Microsoft.Authorization/locks.
        """
        return pulumi.get(self, "type")


class AwaitableGetManagementLockAtResourceGroupLevelResult(GetManagementLockAtResourceGroupLevelResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetManagementLockAtResourceGroupLevelResult(
            level=self.level,
            name=self.name,
            notes=self.notes,
            owners=self.owners,
            type=self.type)


def get_management_lock_at_resource_group_level(lock_name: Optional[str] = None,
                                                resource_group_name: Optional[str] = None,
                                                opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetManagementLockAtResourceGroupLevelResult:
    """
    Use this data source to access information about an existing resource.

    :param str lock_name: The name of the lock to get.
    :param str resource_group_name: The name of the locked resource group.
    """
    __args__ = dict()
    __args__['lockName'] = lock_name
    __args__['resourceGroupName'] = resource_group_name
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = _utilities.get_version()
    __ret__ = pulumi.runtime.invoke('azure-nextgen:authorization/latest:getManagementLockAtResourceGroupLevel', __args__, opts=opts, typ=GetManagementLockAtResourceGroupLevelResult).value

    return AwaitableGetManagementLockAtResourceGroupLevelResult(
        level=__ret__.level,
        name=__ret__.name,
        notes=__ret__.notes,
        owners=__ret__.owners,
        type=__ret__.type)