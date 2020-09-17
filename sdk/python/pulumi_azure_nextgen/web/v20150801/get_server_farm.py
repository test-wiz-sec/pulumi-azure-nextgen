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
    'GetServerFarmResult',
    'AwaitableGetServerFarmResult',
    'get_server_farm',
]

@pulumi.output_type
class GetServerFarmResult:
    """
    App Service Plan Model
    """
    def __init__(__self__, admin_site_name=None, geo_region=None, hosting_environment_profile=None, kind=None, location=None, maximum_number_of_workers=None, name=None, number_of_sites=None, per_site_scaling=None, reserved=None, resource_group=None, sku=None, status=None, subscription=None, tags=None, type=None, worker_tier_name=None):
        if admin_site_name and not isinstance(admin_site_name, str):
            raise TypeError("Expected argument 'admin_site_name' to be a str")
        pulumi.set(__self__, "admin_site_name", admin_site_name)
        if geo_region and not isinstance(geo_region, str):
            raise TypeError("Expected argument 'geo_region' to be a str")
        pulumi.set(__self__, "geo_region", geo_region)
        if hosting_environment_profile and not isinstance(hosting_environment_profile, dict):
            raise TypeError("Expected argument 'hosting_environment_profile' to be a dict")
        pulumi.set(__self__, "hosting_environment_profile", hosting_environment_profile)
        if kind and not isinstance(kind, str):
            raise TypeError("Expected argument 'kind' to be a str")
        pulumi.set(__self__, "kind", kind)
        if location and not isinstance(location, str):
            raise TypeError("Expected argument 'location' to be a str")
        pulumi.set(__self__, "location", location)
        if maximum_number_of_workers and not isinstance(maximum_number_of_workers, int):
            raise TypeError("Expected argument 'maximum_number_of_workers' to be a int")
        pulumi.set(__self__, "maximum_number_of_workers", maximum_number_of_workers)
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        pulumi.set(__self__, "name", name)
        if number_of_sites and not isinstance(number_of_sites, int):
            raise TypeError("Expected argument 'number_of_sites' to be a int")
        pulumi.set(__self__, "number_of_sites", number_of_sites)
        if per_site_scaling and not isinstance(per_site_scaling, bool):
            raise TypeError("Expected argument 'per_site_scaling' to be a bool")
        pulumi.set(__self__, "per_site_scaling", per_site_scaling)
        if reserved and not isinstance(reserved, bool):
            raise TypeError("Expected argument 'reserved' to be a bool")
        pulumi.set(__self__, "reserved", reserved)
        if resource_group and not isinstance(resource_group, str):
            raise TypeError("Expected argument 'resource_group' to be a str")
        pulumi.set(__self__, "resource_group", resource_group)
        if sku and not isinstance(sku, dict):
            raise TypeError("Expected argument 'sku' to be a dict")
        pulumi.set(__self__, "sku", sku)
        if status and not isinstance(status, str):
            raise TypeError("Expected argument 'status' to be a str")
        pulumi.set(__self__, "status", status)
        if subscription and not isinstance(subscription, str):
            raise TypeError("Expected argument 'subscription' to be a str")
        pulumi.set(__self__, "subscription", subscription)
        if tags and not isinstance(tags, dict):
            raise TypeError("Expected argument 'tags' to be a dict")
        pulumi.set(__self__, "tags", tags)
        if type and not isinstance(type, str):
            raise TypeError("Expected argument 'type' to be a str")
        pulumi.set(__self__, "type", type)
        if worker_tier_name and not isinstance(worker_tier_name, str):
            raise TypeError("Expected argument 'worker_tier_name' to be a str")
        pulumi.set(__self__, "worker_tier_name", worker_tier_name)

    @property
    @pulumi.getter(name="adminSiteName")
    def admin_site_name(self) -> Optional[str]:
        """
        App Service Plan administration site
        """
        return pulumi.get(self, "admin_site_name")

    @property
    @pulumi.getter(name="geoRegion")
    def geo_region(self) -> str:
        """
        Geographical location for the App Service Plan
        """
        return pulumi.get(self, "geo_region")

    @property
    @pulumi.getter(name="hostingEnvironmentProfile")
    def hosting_environment_profile(self) -> Optional['outputs.HostingEnvironmentProfileResponse']:
        """
        Specification for the hosting environment (App Service Environment) to use for the App Service Plan
        """
        return pulumi.get(self, "hosting_environment_profile")

    @property
    @pulumi.getter
    def kind(self) -> Optional[str]:
        """
        Kind of resource
        """
        return pulumi.get(self, "kind")

    @property
    @pulumi.getter
    def location(self) -> str:
        """
        Resource Location
        """
        return pulumi.get(self, "location")

    @property
    @pulumi.getter(name="maximumNumberOfWorkers")
    def maximum_number_of_workers(self) -> Optional[int]:
        """
        Maximum number of instances that can be assigned to this App Service Plan
        """
        return pulumi.get(self, "maximum_number_of_workers")

    @property
    @pulumi.getter
    def name(self) -> Optional[str]:
        """
        Resource Name
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="numberOfSites")
    def number_of_sites(self) -> int:
        """
        Number of web apps assigned to this App Service Plan
        """
        return pulumi.get(self, "number_of_sites")

    @property
    @pulumi.getter(name="perSiteScaling")
    def per_site_scaling(self) -> Optional[bool]:
        """
        If True apps assigned to this App Service Plan can be scaled independently
                    If False apps assigned to this App Service Plan will scale to all instances of the plan
        """
        return pulumi.get(self, "per_site_scaling")

    @property
    @pulumi.getter
    def reserved(self) -> Optional[bool]:
        """
        Enables creation of a Linux App Service Plan
        """
        return pulumi.get(self, "reserved")

    @property
    @pulumi.getter(name="resourceGroup")
    def resource_group(self) -> str:
        """
        Resource group of the server farm
        """
        return pulumi.get(self, "resource_group")

    @property
    @pulumi.getter
    def sku(self) -> Optional['outputs.SkuDescriptionResponse']:
        """
        Describes a sku for a scalable resource
        """
        return pulumi.get(self, "sku")

    @property
    @pulumi.getter
    def status(self) -> str:
        """
        App Service Plan Status
        """
        return pulumi.get(self, "status")

    @property
    @pulumi.getter
    def subscription(self) -> str:
        """
        App Service Plan Subscription
        """
        return pulumi.get(self, "subscription")

    @property
    @pulumi.getter
    def tags(self) -> Optional[Mapping[str, str]]:
        """
        Resource tags
        """
        return pulumi.get(self, "tags")

    @property
    @pulumi.getter
    def type(self) -> Optional[str]:
        """
        Resource type
        """
        return pulumi.get(self, "type")

    @property
    @pulumi.getter(name="workerTierName")
    def worker_tier_name(self) -> Optional[str]:
        """
        Target worker tier assigned to the App Service Plan
        """
        return pulumi.get(self, "worker_tier_name")


class AwaitableGetServerFarmResult(GetServerFarmResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetServerFarmResult(
            admin_site_name=self.admin_site_name,
            geo_region=self.geo_region,
            hosting_environment_profile=self.hosting_environment_profile,
            kind=self.kind,
            location=self.location,
            maximum_number_of_workers=self.maximum_number_of_workers,
            name=self.name,
            number_of_sites=self.number_of_sites,
            per_site_scaling=self.per_site_scaling,
            reserved=self.reserved,
            resource_group=self.resource_group,
            sku=self.sku,
            status=self.status,
            subscription=self.subscription,
            tags=self.tags,
            type=self.type,
            worker_tier_name=self.worker_tier_name)


def get_server_farm(name: Optional[str] = None,
                    resource_group_name: Optional[str] = None,
                    opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetServerFarmResult:
    """
    Use this data source to access information about an existing resource.

    :param str name: Name of App Service Plan
    :param str resource_group_name: Name of resource group
    """
    __args__ = dict()
    __args__['name'] = name
    __args__['resourceGroupName'] = resource_group_name
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = _utilities.get_version()
    __ret__ = pulumi.runtime.invoke('azure-nextgen:web/v20150801:getServerFarm', __args__, opts=opts, typ=GetServerFarmResult).value

    return AwaitableGetServerFarmResult(
        admin_site_name=__ret__.admin_site_name,
        geo_region=__ret__.geo_region,
        hosting_environment_profile=__ret__.hosting_environment_profile,
        kind=__ret__.kind,
        location=__ret__.location,
        maximum_number_of_workers=__ret__.maximum_number_of_workers,
        name=__ret__.name,
        number_of_sites=__ret__.number_of_sites,
        per_site_scaling=__ret__.per_site_scaling,
        reserved=__ret__.reserved,
        resource_group=__ret__.resource_group,
        sku=__ret__.sku,
        status=__ret__.status,
        subscription=__ret__.subscription,
        tags=__ret__.tags,
        type=__ret__.type,
        worker_tier_name=__ret__.worker_tier_name)