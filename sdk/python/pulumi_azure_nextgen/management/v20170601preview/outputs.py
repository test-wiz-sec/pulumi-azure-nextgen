# coding=utf-8
# *** WARNING: this file was generated by the Pulumi SDK Generator. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union
from ... import _utilities, _tables

__all__ = [
    'PolicyDefinitionReferenceResponse',
]

@pulumi.output_type
class PolicyDefinitionReferenceResponse(dict):
    """
    The policy definition reference.
    """
    def __init__(__self__, *,
                 parameters: Optional[Any] = None,
                 policy_definition_id: Optional[str] = None):
        """
        The policy definition reference.
        :param Any parameters: Required if a parameter is used in policy rule.
        :param str policy_definition_id: The ID of the policy definition or policy set definition.
        """
        if parameters is not None:
            pulumi.set(__self__, "parameters", parameters)
        if policy_definition_id is not None:
            pulumi.set(__self__, "policy_definition_id", policy_definition_id)

    @property
    @pulumi.getter
    def parameters(self) -> Optional[Any]:
        """
        Required if a parameter is used in policy rule.
        """
        return pulumi.get(self, "parameters")

    @property
    @pulumi.getter(name="policyDefinitionId")
    def policy_definition_id(self) -> Optional[str]:
        """
        The ID of the policy definition or policy set definition.
        """
        return pulumi.get(self, "policy_definition_id")

    def _translate_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop


