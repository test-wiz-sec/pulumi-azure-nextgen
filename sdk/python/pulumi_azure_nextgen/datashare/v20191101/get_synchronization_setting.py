# coding=utf-8
# *** WARNING: this file was generated by the Pulumi SDK Generator. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union
from ... import _utilities, _tables

__all__ = [
    'GetSynchronizationSettingResult',
    'AwaitableGetSynchronizationSettingResult',
    'get_synchronization_setting',
]

@pulumi.output_type
class GetSynchronizationSettingResult:
    """
    A Synchronization Setting data transfer object.
    """
    def __init__(__self__, kind=None, name=None, type=None):
        if kind and not isinstance(kind, str):
            raise TypeError("Expected argument 'kind' to be a str")
        pulumi.set(__self__, "kind", kind)
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        pulumi.set(__self__, "name", name)
        if type and not isinstance(type, str):
            raise TypeError("Expected argument 'type' to be a str")
        pulumi.set(__self__, "type", type)

    @property
    @pulumi.getter
    def kind(self) -> str:
        """
        Kind of synchronization
        """
        return pulumi.get(self, "kind")

    @property
    @pulumi.getter
    def name(self) -> str:
        """
        Name of the azure resource
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def type(self) -> str:
        """
        Type of the azure resource
        """
        return pulumi.get(self, "type")


class AwaitableGetSynchronizationSettingResult(GetSynchronizationSettingResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetSynchronizationSettingResult(
            kind=self.kind,
            name=self.name,
            type=self.type)


def get_synchronization_setting(account_name: Optional[str] = None,
                                resource_group_name: Optional[str] = None,
                                share_name: Optional[str] = None,
                                synchronization_setting_name: Optional[str] = None,
                                opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetSynchronizationSettingResult:
    """
    Use this data source to access information about an existing resource.

    :param str account_name: The name of the share account.
    :param str resource_group_name: The resource group name.
    :param str share_name: The name of the share.
    :param str synchronization_setting_name: The name of the synchronizationSetting.
    """
    __args__ = dict()
    __args__['accountName'] = account_name
    __args__['resourceGroupName'] = resource_group_name
    __args__['shareName'] = share_name
    __args__['synchronizationSettingName'] = synchronization_setting_name
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = _utilities.get_version()
    __ret__ = pulumi.runtime.invoke('azure-nextgen:datashare/v20191101:getSynchronizationSetting', __args__, opts=opts, typ=GetSynchronizationSettingResult).value

    return AwaitableGetSynchronizationSettingResult(
        kind=__ret__.kind,
        name=__ret__.name,
        type=__ret__.type)
