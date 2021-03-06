# coding=utf-8
# *** WARNING: this file was generated by the Pulumi SDK Generator. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union
from ... import _utilities, _tables

__all__ = ['RoleAssignment']


class RoleAssignment(pulumi.CustomResource):
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 can_delegate: Optional[pulumi.Input[bool]] = None,
                 principal_id: Optional[pulumi.Input[str]] = None,
                 principal_type: Optional[pulumi.Input[str]] = None,
                 role_assignment_name: Optional[pulumi.Input[str]] = None,
                 role_definition_id: Optional[pulumi.Input[str]] = None,
                 scope: Optional[pulumi.Input[str]] = None,
                 __props__=None,
                 __name__=None,
                 __opts__=None):
        """
        Role Assignments

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[bool] can_delegate: The delegation flag used for creating a role assignment
        :param pulumi.Input[str] principal_id: The principal ID assigned to the role. This maps to the ID inside the Active Directory. It can point to a user, service principal, or security group.
        :param pulumi.Input[str] principal_type: The principal type of the assigned principal ID.
        :param pulumi.Input[str] role_assignment_name: The name of the role assignment to create. It can be any valid GUID.
        :param pulumi.Input[str] role_definition_id: The role definition ID used in the role assignment.
        :param pulumi.Input[str] scope: The scope of the role assignment to create. The scope can be any REST resource instance. For example, use '/subscriptions/{subscription-id}/' for a subscription, '/subscriptions/{subscription-id}/resourceGroups/{resource-group-name}' for a resource group, and '/subscriptions/{subscription-id}/resourceGroups/{resource-group-name}/providers/{resource-provider}/{resource-type}/{resource-name}' for a resource.
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

            __props__['can_delegate'] = can_delegate
            if principal_id is None:
                raise TypeError("Missing required property 'principal_id'")
            __props__['principal_id'] = principal_id
            __props__['principal_type'] = principal_type
            if role_assignment_name is None:
                raise TypeError("Missing required property 'role_assignment_name'")
            __props__['role_assignment_name'] = role_assignment_name
            if role_definition_id is None:
                raise TypeError("Missing required property 'role_definition_id'")
            __props__['role_definition_id'] = role_definition_id
            if scope is None:
                raise TypeError("Missing required property 'scope'")
            __props__['scope'] = scope
            __props__['name'] = None
            __props__['type'] = None
        alias_opts = pulumi.ResourceOptions(aliases=[pulumi.Alias(type_="azure-nextgen:authorization/latest:RoleAssignment"), pulumi.Alias(type_="azure-nextgen:authorization/v20150701:RoleAssignment"), pulumi.Alias(type_="azure-nextgen:authorization/v20171001preview:RoleAssignment"), pulumi.Alias(type_="azure-nextgen:authorization/v20180101preview:RoleAssignment"), pulumi.Alias(type_="azure-nextgen:authorization/v20200401preview:RoleAssignment")])
        opts = pulumi.ResourceOptions.merge(opts, alias_opts)
        super(RoleAssignment, __self__).__init__(
            'azure-nextgen:authorization/v20180901preview:RoleAssignment',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None) -> 'RoleAssignment':
        """
        Get an existing RoleAssignment resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        return RoleAssignment(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="canDelegate")
    def can_delegate(self) -> pulumi.Output[Optional[bool]]:
        """
        The Delegation flag for the role assignment
        """
        return pulumi.get(self, "can_delegate")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        The role assignment name.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="principalId")
    def principal_id(self) -> pulumi.Output[Optional[str]]:
        """
        The principal ID.
        """
        return pulumi.get(self, "principal_id")

    @property
    @pulumi.getter(name="principalType")
    def principal_type(self) -> pulumi.Output[Optional[str]]:
        """
        The principal type of the assigned principal ID.
        """
        return pulumi.get(self, "principal_type")

    @property
    @pulumi.getter(name="roleDefinitionId")
    def role_definition_id(self) -> pulumi.Output[Optional[str]]:
        """
        The role definition ID.
        """
        return pulumi.get(self, "role_definition_id")

    @property
    @pulumi.getter
    def scope(self) -> pulumi.Output[Optional[str]]:
        """
        The role assignment scope.
        """
        return pulumi.get(self, "scope")

    @property
    @pulumi.getter
    def type(self) -> pulumi.Output[str]:
        """
        The role assignment type.
        """
        return pulumi.get(self, "type")

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

