# coding=utf-8
# *** WARNING: this file was generated by the Pulumi SDK Generator. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union
from ... import _utilities, _tables
from . import outputs
from ._inputs import *

__all__ = ['ChapSetting']


class ChapSetting(pulumi.CustomResource):
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 chap_user_name: Optional[pulumi.Input[str]] = None,
                 device_name: Optional[pulumi.Input[str]] = None,
                 manager_name: Optional[pulumi.Input[str]] = None,
                 password: Optional[pulumi.Input[pulumi.InputType['AsymmetricEncryptedSecretArgs']]] = None,
                 resource_group_name: Optional[pulumi.Input[str]] = None,
                 __props__=None,
                 __name__=None,
                 __opts__=None):
        """
        Challenge-Handshake Authentication Protocol (CHAP) setting

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] chap_user_name: The chap user name.
        :param pulumi.Input[str] device_name: The device name.
        :param pulumi.Input[str] manager_name: The manager name
        :param pulumi.Input[pulumi.InputType['AsymmetricEncryptedSecretArgs']] password: The chap password.
        :param pulumi.Input[str] resource_group_name: The resource group name
        """
        if __name__ is not None:
            warnings.warn("explicit use of __name__ is deprecated", DeprecationWarning)
            resource_name = __name__
        if __opts__ is not None:
            warnings.warn("explicit use of __opts__ is deprecated, use 'opts' instead", DeprecationWarning)
            opts = __opts__
        if opts is None:
            opts = pulumi.ResourceOptions()
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.version is None:
            opts.version = _utilities.get_version()
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = dict()

            if chap_user_name is None:
                raise TypeError("Missing required property 'chap_user_name'")
            __props__['chap_user_name'] = chap_user_name
            if device_name is None:
                raise TypeError("Missing required property 'device_name'")
            __props__['device_name'] = device_name
            if manager_name is None:
                raise TypeError("Missing required property 'manager_name'")
            __props__['manager_name'] = manager_name
            if password is None:
                raise TypeError("Missing required property 'password'")
            __props__['password'] = password
            if resource_group_name is None:
                raise TypeError("Missing required property 'resource_group_name'")
            __props__['resource_group_name'] = resource_group_name
            __props__['name'] = None
            __props__['type'] = None
        alias_opts = pulumi.ResourceOptions(aliases=[pulumi.Alias(type_="azure-nextgen:storsimple/v20161001:ChapSetting")])
        opts = pulumi.ResourceOptions.merge(opts, alias_opts)
        super(ChapSetting, __self__).__init__(
            'azure-nextgen:storsimple/latest:ChapSetting',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None) -> 'ChapSetting':
        """
        Get an existing ChapSetting resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        return ChapSetting(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        The name.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def password(self) -> pulumi.Output['outputs.AsymmetricEncryptedSecretResponse']:
        """
        The chap password.
        """
        return pulumi.get(self, "password")

    @property
    @pulumi.getter
    def type(self) -> pulumi.Output[str]:
        """
        The type.
        """
        return pulumi.get(self, "type")

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

