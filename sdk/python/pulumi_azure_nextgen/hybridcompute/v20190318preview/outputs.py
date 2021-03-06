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
    'ErrorDetailResponse',
    'OSProfileResponse',
]

@pulumi.output_type
class ErrorDetailResponse(dict):
    def __init__(__self__, *,
                 code: str,
                 message: str,
                 details: Optional[Sequence['outputs.ErrorDetailResponse']] = None,
                 target: Optional[str] = None):
        """
        :param str code: The error's code.
        :param str message: A human readable error message.
        :param Sequence['ErrorDetailResponseArgs'] details: Additional error details.
        :param str target: Indicates which property in the request is responsible for the error.
        """
        pulumi.set(__self__, "code", code)
        pulumi.set(__self__, "message", message)
        if details is not None:
            pulumi.set(__self__, "details", details)
        if target is not None:
            pulumi.set(__self__, "target", target)

    @property
    @pulumi.getter
    def code(self) -> str:
        """
        The error's code.
        """
        return pulumi.get(self, "code")

    @property
    @pulumi.getter
    def message(self) -> str:
        """
        A human readable error message.
        """
        return pulumi.get(self, "message")

    @property
    @pulumi.getter
    def details(self) -> Optional[Sequence['outputs.ErrorDetailResponse']]:
        """
        Additional error details.
        """
        return pulumi.get(self, "details")

    @property
    @pulumi.getter
    def target(self) -> Optional[str]:
        """
        Indicates which property in the request is responsible for the error.
        """
        return pulumi.get(self, "target")

    def _translate_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop


@pulumi.output_type
class OSProfileResponse(dict):
    """
    Specifies the operating system settings for the hybrid machine.
    """
    def __init__(__self__, *,
                 computer_name: str):
        """
        Specifies the operating system settings for the hybrid machine.
        :param str computer_name: Specifies the host OS name of the hybrid machine.
        """
        pulumi.set(__self__, "computer_name", computer_name)

    @property
    @pulumi.getter(name="computerName")
    def computer_name(self) -> str:
        """
        Specifies the host OS name of the hybrid machine.
        """
        return pulumi.get(self, "computer_name")

    def _translate_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop


