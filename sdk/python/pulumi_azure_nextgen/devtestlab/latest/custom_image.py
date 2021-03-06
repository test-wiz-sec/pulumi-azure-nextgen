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

__all__ = ['CustomImage']


class CustomImage(pulumi.CustomResource):
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 author: Optional[pulumi.Input[str]] = None,
                 custom_image_plan: Optional[pulumi.Input[pulumi.InputType['CustomImagePropertiesFromPlanArgs']]] = None,
                 data_disk_storage_info: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['DataDiskStorageTypeInfoArgs']]]]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 is_plan_authorized: Optional[pulumi.Input[bool]] = None,
                 lab_name: Optional[pulumi.Input[str]] = None,
                 location: Optional[pulumi.Input[str]] = None,
                 managed_image_id: Optional[pulumi.Input[str]] = None,
                 managed_snapshot_id: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 resource_group_name: Optional[pulumi.Input[str]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 vhd: Optional[pulumi.Input[pulumi.InputType['CustomImagePropertiesCustomArgs']]] = None,
                 vm: Optional[pulumi.Input[pulumi.InputType['CustomImagePropertiesFromVmArgs']]] = None,
                 __props__=None,
                 __name__=None,
                 __opts__=None):
        """
        A custom image.

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] author: The author of the custom image.
        :param pulumi.Input[pulumi.InputType['CustomImagePropertiesFromPlanArgs']] custom_image_plan: Storage information about the plan related to this custom image
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['DataDiskStorageTypeInfoArgs']]]] data_disk_storage_info: Storage information about the data disks present in the custom image
        :param pulumi.Input[str] description: The description of the custom image.
        :param pulumi.Input[bool] is_plan_authorized: Whether or not the custom images underlying offer/plan has been enabled for programmatic deployment
        :param pulumi.Input[str] lab_name: The name of the lab.
        :param pulumi.Input[str] location: The location of the resource.
        :param pulumi.Input[str] managed_image_id: The Managed Image Id backing the custom image.
        :param pulumi.Input[str] managed_snapshot_id: The Managed Snapshot Id backing the custom image.
        :param pulumi.Input[str] name: The name of the custom image.
        :param pulumi.Input[str] resource_group_name: The name of the resource group.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] tags: The tags of the resource.
        :param pulumi.Input[pulumi.InputType['CustomImagePropertiesCustomArgs']] vhd: The VHD from which the image is to be created.
        :param pulumi.Input[pulumi.InputType['CustomImagePropertiesFromVmArgs']] vm: The virtual machine from which the image is to be created.
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

            __props__['author'] = author
            __props__['custom_image_plan'] = custom_image_plan
            __props__['data_disk_storage_info'] = data_disk_storage_info
            __props__['description'] = description
            __props__['is_plan_authorized'] = is_plan_authorized
            if lab_name is None:
                raise TypeError("Missing required property 'lab_name'")
            __props__['lab_name'] = lab_name
            __props__['location'] = location
            __props__['managed_image_id'] = managed_image_id
            __props__['managed_snapshot_id'] = managed_snapshot_id
            if name is None:
                raise TypeError("Missing required property 'name'")
            __props__['name'] = name
            if resource_group_name is None:
                raise TypeError("Missing required property 'resource_group_name'")
            __props__['resource_group_name'] = resource_group_name
            __props__['tags'] = tags
            __props__['vhd'] = vhd
            __props__['vm'] = vm
            __props__['creation_date'] = None
            __props__['provisioning_state'] = None
            __props__['type'] = None
            __props__['unique_identifier'] = None
        alias_opts = pulumi.ResourceOptions(aliases=[pulumi.Alias(type_="azure-nextgen:devtestlab/v20150521preview:CustomImage"), pulumi.Alias(type_="azure-nextgen:devtestlab/v20160515:CustomImage"), pulumi.Alias(type_="azure-nextgen:devtestlab/v20180915:CustomImage")])
        opts = pulumi.ResourceOptions.merge(opts, alias_opts)
        super(CustomImage, __self__).__init__(
            'azure-nextgen:devtestlab/latest:CustomImage',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None) -> 'CustomImage':
        """
        Get an existing CustomImage resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        return CustomImage(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def author(self) -> pulumi.Output[Optional[str]]:
        """
        The author of the custom image.
        """
        return pulumi.get(self, "author")

    @property
    @pulumi.getter(name="creationDate")
    def creation_date(self) -> pulumi.Output[str]:
        """
        The creation date of the custom image.
        """
        return pulumi.get(self, "creation_date")

    @property
    @pulumi.getter(name="customImagePlan")
    def custom_image_plan(self) -> pulumi.Output[Optional['outputs.CustomImagePropertiesFromPlanResponse']]:
        """
        Storage information about the plan related to this custom image
        """
        return pulumi.get(self, "custom_image_plan")

    @property
    @pulumi.getter(name="dataDiskStorageInfo")
    def data_disk_storage_info(self) -> pulumi.Output[Optional[Sequence['outputs.DataDiskStorageTypeInfoResponse']]]:
        """
        Storage information about the data disks present in the custom image
        """
        return pulumi.get(self, "data_disk_storage_info")

    @property
    @pulumi.getter
    def description(self) -> pulumi.Output[Optional[str]]:
        """
        The description of the custom image.
        """
        return pulumi.get(self, "description")

    @property
    @pulumi.getter(name="isPlanAuthorized")
    def is_plan_authorized(self) -> pulumi.Output[Optional[bool]]:
        """
        Whether or not the custom images underlying offer/plan has been enabled for programmatic deployment
        """
        return pulumi.get(self, "is_plan_authorized")

    @property
    @pulumi.getter
    def location(self) -> pulumi.Output[Optional[str]]:
        """
        The location of the resource.
        """
        return pulumi.get(self, "location")

    @property
    @pulumi.getter(name="managedImageId")
    def managed_image_id(self) -> pulumi.Output[Optional[str]]:
        """
        The Managed Image Id backing the custom image.
        """
        return pulumi.get(self, "managed_image_id")

    @property
    @pulumi.getter(name="managedSnapshotId")
    def managed_snapshot_id(self) -> pulumi.Output[Optional[str]]:
        """
        The Managed Snapshot Id backing the custom image.
        """
        return pulumi.get(self, "managed_snapshot_id")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        The name of the resource.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="provisioningState")
    def provisioning_state(self) -> pulumi.Output[str]:
        """
        The provisioning status of the resource.
        """
        return pulumi.get(self, "provisioning_state")

    @property
    @pulumi.getter
    def tags(self) -> pulumi.Output[Optional[Mapping[str, str]]]:
        """
        The tags of the resource.
        """
        return pulumi.get(self, "tags")

    @property
    @pulumi.getter
    def type(self) -> pulumi.Output[str]:
        """
        The type of the resource.
        """
        return pulumi.get(self, "type")

    @property
    @pulumi.getter(name="uniqueIdentifier")
    def unique_identifier(self) -> pulumi.Output[str]:
        """
        The unique immutable identifier of a resource (Guid).
        """
        return pulumi.get(self, "unique_identifier")

    @property
    @pulumi.getter
    def vhd(self) -> pulumi.Output[Optional['outputs.CustomImagePropertiesCustomResponse']]:
        """
        The VHD from which the image is to be created.
        """
        return pulumi.get(self, "vhd")

    @property
    @pulumi.getter
    def vm(self) -> pulumi.Output[Optional['outputs.CustomImagePropertiesFromVmResponse']]:
        """
        The virtual machine from which the image is to be created.
        """
        return pulumi.get(self, "vm")

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

