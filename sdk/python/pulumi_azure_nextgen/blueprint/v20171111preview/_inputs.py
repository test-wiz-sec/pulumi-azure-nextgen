# coding=utf-8
# *** WARNING: this file was generated by the Pulumi SDK Generator. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union
from ... import _utilities, _tables

__all__ = [
    'AssignmentLockSettingsArgs',
    'ManagedServiceIdentityArgs',
    'ParameterValueBaseArgs',
    'ResourceGroupValueArgs',
]

@pulumi.input_type
class AssignmentLockSettingsArgs:
    def __init__(__self__, *,
                 mode: Optional[pulumi.Input[str]] = None):
        """
        Defines how Blueprint-managed resources will be locked.
        :param pulumi.Input[str] mode: Lock mode.
        """
        if mode is not None:
            pulumi.set(__self__, "mode", mode)

    @property
    @pulumi.getter
    def mode(self) -> Optional[pulumi.Input[str]]:
        """
        Lock mode.
        """
        return pulumi.get(self, "mode")

    @mode.setter
    def mode(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "mode", value)


@pulumi.input_type
class ManagedServiceIdentityArgs:
    def __init__(__self__, *,
                 type: pulumi.Input[str],
                 principal_id: Optional[pulumi.Input[str]] = None,
                 tenant_id: Optional[pulumi.Input[str]] = None):
        """
        Managed Service Identity
        :param pulumi.Input[str] type: Type of the Managed Service Identity.
        :param pulumi.Input[str] principal_id: Azure Active Directory principal ID associated with this Identity.
        :param pulumi.Input[str] tenant_id: ID of the Azure Active Directory.
        """
        pulumi.set(__self__, "type", type)
        if principal_id is not None:
            pulumi.set(__self__, "principal_id", principal_id)
        if tenant_id is not None:
            pulumi.set(__self__, "tenant_id", tenant_id)

    @property
    @pulumi.getter
    def type(self) -> pulumi.Input[str]:
        """
        Type of the Managed Service Identity.
        """
        return pulumi.get(self, "type")

    @type.setter
    def type(self, value: pulumi.Input[str]):
        pulumi.set(self, "type", value)

    @property
    @pulumi.getter(name="principalId")
    def principal_id(self) -> Optional[pulumi.Input[str]]:
        """
        Azure Active Directory principal ID associated with this Identity.
        """
        return pulumi.get(self, "principal_id")

    @principal_id.setter
    def principal_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "principal_id", value)

    @property
    @pulumi.getter(name="tenantId")
    def tenant_id(self) -> Optional[pulumi.Input[str]]:
        """
        ID of the Azure Active Directory.
        """
        return pulumi.get(self, "tenant_id")

    @tenant_id.setter
    def tenant_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "tenant_id", value)


@pulumi.input_type
class ParameterValueBaseArgs:
    def __init__(__self__, *,
                 description: Optional[pulumi.Input[str]] = None):
        """
        Base class for ParameterValue.
        :param pulumi.Input[str] description: Optional property, just to establish ParameterValueBase as a BaseClass.
        """
        if description is not None:
            pulumi.set(__self__, "description", description)

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[str]]:
        """
        Optional property, just to establish ParameterValueBase as a BaseClass.
        """
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "description", value)


@pulumi.input_type
class ResourceGroupValueArgs:
    def __init__(__self__, *,
                 location: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None):
        """
        Represents an Azure resource group.
        :param pulumi.Input[str] location: Location of the resource group
        :param pulumi.Input[str] name: Name of the resource group
        """
        if location is not None:
            pulumi.set(__self__, "location", location)
        if name is not None:
            pulumi.set(__self__, "name", name)

    @property
    @pulumi.getter
    def location(self) -> Optional[pulumi.Input[str]]:
        """
        Location of the resource group
        """
        return pulumi.get(self, "location")

    @location.setter
    def location(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "location", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        Name of the resource group
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)


