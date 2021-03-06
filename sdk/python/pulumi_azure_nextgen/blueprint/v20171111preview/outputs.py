# coding=utf-8
# *** WARNING: this file was generated by the Pulumi SDK Generator. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union
from ... import _utilities, _tables

__all__ = [
    'AssignmentLockSettingsResponse',
    'AssignmentStatusResponse',
    'ManagedServiceIdentityResponse',
    'ParameterValueBaseResponse',
    'ResourceGroupValueResponse',
]

@pulumi.output_type
class AssignmentLockSettingsResponse(dict):
    """
    Defines how Blueprint-managed resources will be locked.
    """
    def __init__(__self__, *,
                 mode: Optional[str] = None):
        """
        Defines how Blueprint-managed resources will be locked.
        :param str mode: Lock mode.
        """
        if mode is not None:
            pulumi.set(__self__, "mode", mode)

    @property
    @pulumi.getter
    def mode(self) -> Optional[str]:
        """
        Lock mode.
        """
        return pulumi.get(self, "mode")

    def _translate_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop


@pulumi.output_type
class AssignmentStatusResponse(dict):
    """
    The status of Blueprint assignment. This field is readonly.
    """
    def __init__(__self__, *,
                 last_modified: str,
                 time_created: str):
        """
        The status of Blueprint assignment. This field is readonly.
        :param str last_modified: Last modified time of this blueprint.
        :param str time_created: Creation time of this blueprint.
        """
        pulumi.set(__self__, "last_modified", last_modified)
        pulumi.set(__self__, "time_created", time_created)

    @property
    @pulumi.getter(name="lastModified")
    def last_modified(self) -> str:
        """
        Last modified time of this blueprint.
        """
        return pulumi.get(self, "last_modified")

    @property
    @pulumi.getter(name="timeCreated")
    def time_created(self) -> str:
        """
        Creation time of this blueprint.
        """
        return pulumi.get(self, "time_created")

    def _translate_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop


@pulumi.output_type
class ManagedServiceIdentityResponse(dict):
    """
    Managed Service Identity
    """
    def __init__(__self__, *,
                 type: str,
                 principal_id: Optional[str] = None,
                 tenant_id: Optional[str] = None):
        """
        Managed Service Identity
        :param str type: Type of the Managed Service Identity.
        :param str principal_id: Azure Active Directory principal ID associated with this Identity.
        :param str tenant_id: ID of the Azure Active Directory.
        """
        pulumi.set(__self__, "type", type)
        if principal_id is not None:
            pulumi.set(__self__, "principal_id", principal_id)
        if tenant_id is not None:
            pulumi.set(__self__, "tenant_id", tenant_id)

    @property
    @pulumi.getter
    def type(self) -> str:
        """
        Type of the Managed Service Identity.
        """
        return pulumi.get(self, "type")

    @property
    @pulumi.getter(name="principalId")
    def principal_id(self) -> Optional[str]:
        """
        Azure Active Directory principal ID associated with this Identity.
        """
        return pulumi.get(self, "principal_id")

    @property
    @pulumi.getter(name="tenantId")
    def tenant_id(self) -> Optional[str]:
        """
        ID of the Azure Active Directory.
        """
        return pulumi.get(self, "tenant_id")

    def _translate_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop


@pulumi.output_type
class ParameterValueBaseResponse(dict):
    """
    Base class for ParameterValue.
    """
    def __init__(__self__, *,
                 description: Optional[str] = None):
        """
        Base class for ParameterValue.
        :param str description: Optional property, just to establish ParameterValueBase as a BaseClass.
        """
        if description is not None:
            pulumi.set(__self__, "description", description)

    @property
    @pulumi.getter
    def description(self) -> Optional[str]:
        """
        Optional property, just to establish ParameterValueBase as a BaseClass.
        """
        return pulumi.get(self, "description")

    def _translate_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop


@pulumi.output_type
class ResourceGroupValueResponse(dict):
    """
    Represents an Azure resource group.
    """
    def __init__(__self__, *,
                 location: Optional[str] = None,
                 name: Optional[str] = None):
        """
        Represents an Azure resource group.
        :param str location: Location of the resource group
        :param str name: Name of the resource group
        """
        if location is not None:
            pulumi.set(__self__, "location", location)
        if name is not None:
            pulumi.set(__self__, "name", name)

    @property
    @pulumi.getter
    def location(self) -> Optional[str]:
        """
        Location of the resource group
        """
        return pulumi.get(self, "location")

    @property
    @pulumi.getter
    def name(self) -> Optional[str]:
        """
        Name of the resource group
        """
        return pulumi.get(self, "name")

    def _translate_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop


