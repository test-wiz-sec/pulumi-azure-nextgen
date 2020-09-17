# coding=utf-8
# *** WARNING: this file was generated by the Pulumi SDK Generator. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union
from ... import _utilities, _tables

__all__ = [
    'PermissionResponse',
]

@pulumi.output_type
class PermissionResponse(dict):
    """
    Role definition permissions.
    """
    def __init__(__self__, *,
                 actions: Optional[Sequence[str]] = None,
                 data_actions: Optional[Sequence[str]] = None,
                 not_actions: Optional[Sequence[str]] = None,
                 not_data_actions: Optional[Sequence[str]] = None):
        """
        Role definition permissions.
        :param Sequence[str] actions: Allowed actions.
        :param Sequence[str] data_actions: Allowed Data actions.
        :param Sequence[str] not_actions: Denied actions.
        :param Sequence[str] not_data_actions: Denied Data actions.
        """
        if actions is not None:
            pulumi.set(__self__, "actions", actions)
        if data_actions is not None:
            pulumi.set(__self__, "data_actions", data_actions)
        if not_actions is not None:
            pulumi.set(__self__, "not_actions", not_actions)
        if not_data_actions is not None:
            pulumi.set(__self__, "not_data_actions", not_data_actions)

    @property
    @pulumi.getter
    def actions(self) -> Optional[Sequence[str]]:
        """
        Allowed actions.
        """
        return pulumi.get(self, "actions")

    @property
    @pulumi.getter(name="dataActions")
    def data_actions(self) -> Optional[Sequence[str]]:
        """
        Allowed Data actions.
        """
        return pulumi.get(self, "data_actions")

    @property
    @pulumi.getter(name="notActions")
    def not_actions(self) -> Optional[Sequence[str]]:
        """
        Denied actions.
        """
        return pulumi.get(self, "not_actions")

    @property
    @pulumi.getter(name="notDataActions")
    def not_data_actions(self) -> Optional[Sequence[str]]:
        """
        Denied Data actions.
        """
        return pulumi.get(self, "not_data_actions")

    def _translate_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

