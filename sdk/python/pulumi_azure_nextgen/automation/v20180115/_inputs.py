# coding=utf-8
# *** WARNING: this file was generated by the Pulumi SDK Generator. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union
from ... import _utilities, _tables

__all__ = [
    'ContentHashArgs',
    'ContentSourceArgs',
    'DscConfigurationAssociationPropertyArgs',
]

@pulumi.input_type
class ContentHashArgs:
    def __init__(__self__, *,
                 algorithm: pulumi.Input[str],
                 value: pulumi.Input[str]):
        """
        Definition of the runbook property type.
        :param pulumi.Input[str] algorithm: Gets or sets the content hash algorithm used to hash the content.
        :param pulumi.Input[str] value: Gets or sets expected hash value of the content.
        """
        pulumi.set(__self__, "algorithm", algorithm)
        pulumi.set(__self__, "value", value)

    @property
    @pulumi.getter
    def algorithm(self) -> pulumi.Input[str]:
        """
        Gets or sets the content hash algorithm used to hash the content.
        """
        return pulumi.get(self, "algorithm")

    @algorithm.setter
    def algorithm(self, value: pulumi.Input[str]):
        pulumi.set(self, "algorithm", value)

    @property
    @pulumi.getter
    def value(self) -> pulumi.Input[str]:
        """
        Gets or sets expected hash value of the content.
        """
        return pulumi.get(self, "value")

    @value.setter
    def value(self, value: pulumi.Input[str]):
        pulumi.set(self, "value", value)


@pulumi.input_type
class ContentSourceArgs:
    def __init__(__self__, *,
                 hash: Optional[pulumi.Input['ContentHashArgs']] = None,
                 type: Optional[pulumi.Input[str]] = None,
                 value: Optional[pulumi.Input[str]] = None,
                 version: Optional[pulumi.Input[str]] = None):
        """
        Definition of the content source.
        :param pulumi.Input['ContentHashArgs'] hash: Gets or sets the hash.
        :param pulumi.Input[str] type: Gets or sets the content source type.
        :param pulumi.Input[str] value: Gets or sets the value of the content. This is based on the content source type.
        :param pulumi.Input[str] version: Gets or sets the version of the content.
        """
        if hash is not None:
            pulumi.set(__self__, "hash", hash)
        if type is not None:
            pulumi.set(__self__, "type", type)
        if value is not None:
            pulumi.set(__self__, "value", value)
        if version is not None:
            pulumi.set(__self__, "version", version)

    @property
    @pulumi.getter
    def hash(self) -> Optional[pulumi.Input['ContentHashArgs']]:
        """
        Gets or sets the hash.
        """
        return pulumi.get(self, "hash")

    @hash.setter
    def hash(self, value: Optional[pulumi.Input['ContentHashArgs']]):
        pulumi.set(self, "hash", value)

    @property
    @pulumi.getter
    def type(self) -> Optional[pulumi.Input[str]]:
        """
        Gets or sets the content source type.
        """
        return pulumi.get(self, "type")

    @type.setter
    def type(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "type", value)

    @property
    @pulumi.getter
    def value(self) -> Optional[pulumi.Input[str]]:
        """
        Gets or sets the value of the content. This is based on the content source type.
        """
        return pulumi.get(self, "value")

    @value.setter
    def value(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "value", value)

    @property
    @pulumi.getter
    def version(self) -> Optional[pulumi.Input[str]]:
        """
        Gets or sets the version of the content.
        """
        return pulumi.get(self, "version")

    @version.setter
    def version(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "version", value)


@pulumi.input_type
class DscConfigurationAssociationPropertyArgs:
    def __init__(__self__, *,
                 name: Optional[pulumi.Input[str]] = None):
        """
        The Dsc configuration property associated with the entity.
        :param pulumi.Input[str] name: Gets or sets the name of the Dsc configuration.
        """
        if name is not None:
            pulumi.set(__self__, "name", name)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        Gets or sets the name of the Dsc configuration.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

