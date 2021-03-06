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

__all__ = ['ExpressRouteCircuit']


class ExpressRouteCircuit(pulumi.CustomResource):
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 authorizations: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['ExpressRouteCircuitAuthorizationArgs']]]]] = None,
                 circuit_name: Optional[pulumi.Input[str]] = None,
                 circuit_provisioning_state: Optional[pulumi.Input[str]] = None,
                 etag: Optional[pulumi.Input[str]] = None,
                 id: Optional[pulumi.Input[str]] = None,
                 location: Optional[pulumi.Input[str]] = None,
                 peerings: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['ExpressRouteCircuitPeeringArgs']]]]] = None,
                 provisioning_state: Optional[pulumi.Input[str]] = None,
                 resource_group_name: Optional[pulumi.Input[str]] = None,
                 service_key: Optional[pulumi.Input[str]] = None,
                 service_provider_notes: Optional[pulumi.Input[str]] = None,
                 service_provider_properties: Optional[pulumi.Input[pulumi.InputType['ExpressRouteCircuitServiceProviderPropertiesArgs']]] = None,
                 service_provider_provisioning_state: Optional[pulumi.Input[str]] = None,
                 sku: Optional[pulumi.Input[pulumi.InputType['ExpressRouteCircuitSkuArgs']]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 __props__=None,
                 __name__=None,
                 __opts__=None):
        """
        ExpressRouteCircuit resource

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['ExpressRouteCircuitAuthorizationArgs']]]] authorizations: The list of authorizations.
        :param pulumi.Input[str] circuit_name: The name of the circuit.
        :param pulumi.Input[str] circuit_provisioning_state: The CircuitProvisioningState state of the resource.
        :param pulumi.Input[str] etag: Gets a unique read-only string that changes whenever the resource is updated.
        :param pulumi.Input[str] id: Resource Identifier.
        :param pulumi.Input[str] location: Resource location.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['ExpressRouteCircuitPeeringArgs']]]] peerings: The list of peerings.
        :param pulumi.Input[str] provisioning_state: Gets the provisioning state of the public IP resource. Possible values are: 'Updating', 'Deleting', and 'Failed'.
        :param pulumi.Input[str] resource_group_name: The name of the resource group.
        :param pulumi.Input[str] service_key: The ServiceKey.
        :param pulumi.Input[str] service_provider_notes: The ServiceProviderNotes.
        :param pulumi.Input[pulumi.InputType['ExpressRouteCircuitServiceProviderPropertiesArgs']] service_provider_properties: The ServiceProviderProperties.
        :param pulumi.Input[str] service_provider_provisioning_state: The ServiceProviderProvisioningState state of the resource. Possible values are 'NotProvisioned', 'Provisioning', 'Provisioned', and 'Deprovisioning'.
        :param pulumi.Input[pulumi.InputType['ExpressRouteCircuitSkuArgs']] sku: The SKU.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] tags: Resource tags.
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

            __props__['authorizations'] = authorizations
            if circuit_name is None:
                raise TypeError("Missing required property 'circuit_name'")
            __props__['circuit_name'] = circuit_name
            __props__['circuit_provisioning_state'] = circuit_provisioning_state
            __props__['etag'] = etag
            __props__['id'] = id
            __props__['location'] = location
            __props__['peerings'] = peerings
            __props__['provisioning_state'] = provisioning_state
            if resource_group_name is None:
                raise TypeError("Missing required property 'resource_group_name'")
            __props__['resource_group_name'] = resource_group_name
            __props__['service_key'] = service_key
            __props__['service_provider_notes'] = service_provider_notes
            __props__['service_provider_properties'] = service_provider_properties
            __props__['service_provider_provisioning_state'] = service_provider_provisioning_state
            __props__['sku'] = sku
            __props__['tags'] = tags
            __props__['name'] = None
            __props__['type'] = None
        alias_opts = pulumi.ResourceOptions(aliases=[pulumi.Alias(type_="azure-nextgen:network/latest:ExpressRouteCircuit"), pulumi.Alias(type_="azure-nextgen:network/v20150501preview:ExpressRouteCircuit"), pulumi.Alias(type_="azure-nextgen:network/v20160330:ExpressRouteCircuit"), pulumi.Alias(type_="azure-nextgen:network/v20160601:ExpressRouteCircuit"), pulumi.Alias(type_="azure-nextgen:network/v20160901:ExpressRouteCircuit"), pulumi.Alias(type_="azure-nextgen:network/v20161201:ExpressRouteCircuit"), pulumi.Alias(type_="azure-nextgen:network/v20170301:ExpressRouteCircuit"), pulumi.Alias(type_="azure-nextgen:network/v20170601:ExpressRouteCircuit"), pulumi.Alias(type_="azure-nextgen:network/v20170801:ExpressRouteCircuit"), pulumi.Alias(type_="azure-nextgen:network/v20170901:ExpressRouteCircuit"), pulumi.Alias(type_="azure-nextgen:network/v20171001:ExpressRouteCircuit"), pulumi.Alias(type_="azure-nextgen:network/v20171101:ExpressRouteCircuit"), pulumi.Alias(type_="azure-nextgen:network/v20180101:ExpressRouteCircuit"), pulumi.Alias(type_="azure-nextgen:network/v20180201:ExpressRouteCircuit"), pulumi.Alias(type_="azure-nextgen:network/v20180401:ExpressRouteCircuit"), pulumi.Alias(type_="azure-nextgen:network/v20180601:ExpressRouteCircuit"), pulumi.Alias(type_="azure-nextgen:network/v20180701:ExpressRouteCircuit"), pulumi.Alias(type_="azure-nextgen:network/v20180801:ExpressRouteCircuit"), pulumi.Alias(type_="azure-nextgen:network/v20181001:ExpressRouteCircuit"), pulumi.Alias(type_="azure-nextgen:network/v20181101:ExpressRouteCircuit"), pulumi.Alias(type_="azure-nextgen:network/v20181201:ExpressRouteCircuit"), pulumi.Alias(type_="azure-nextgen:network/v20190201:ExpressRouteCircuit"), pulumi.Alias(type_="azure-nextgen:network/v20190401:ExpressRouteCircuit"), pulumi.Alias(type_="azure-nextgen:network/v20190601:ExpressRouteCircuit"), pulumi.Alias(type_="azure-nextgen:network/v20190701:ExpressRouteCircuit"), pulumi.Alias(type_="azure-nextgen:network/v20190801:ExpressRouteCircuit"), pulumi.Alias(type_="azure-nextgen:network/v20190901:ExpressRouteCircuit"), pulumi.Alias(type_="azure-nextgen:network/v20191101:ExpressRouteCircuit"), pulumi.Alias(type_="azure-nextgen:network/v20191201:ExpressRouteCircuit"), pulumi.Alias(type_="azure-nextgen:network/v20200301:ExpressRouteCircuit"), pulumi.Alias(type_="azure-nextgen:network/v20200401:ExpressRouteCircuit"), pulumi.Alias(type_="azure-nextgen:network/v20200501:ExpressRouteCircuit"), pulumi.Alias(type_="azure-nextgen:network/v20200601:ExpressRouteCircuit"), pulumi.Alias(type_="azure-nextgen:network/v20200701:ExpressRouteCircuit")])
        opts = pulumi.ResourceOptions.merge(opts, alias_opts)
        super(ExpressRouteCircuit, __self__).__init__(
            'azure-nextgen:network/v20150615:ExpressRouteCircuit',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None) -> 'ExpressRouteCircuit':
        """
        Get an existing ExpressRouteCircuit resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        return ExpressRouteCircuit(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def authorizations(self) -> pulumi.Output[Optional[Sequence['outputs.ExpressRouteCircuitAuthorizationResponse']]]:
        """
        The list of authorizations.
        """
        return pulumi.get(self, "authorizations")

    @property
    @pulumi.getter(name="circuitProvisioningState")
    def circuit_provisioning_state(self) -> pulumi.Output[Optional[str]]:
        """
        The CircuitProvisioningState state of the resource.
        """
        return pulumi.get(self, "circuit_provisioning_state")

    @property
    @pulumi.getter
    def etag(self) -> pulumi.Output[Optional[str]]:
        """
        Gets a unique read-only string that changes whenever the resource is updated.
        """
        return pulumi.get(self, "etag")

    @property
    @pulumi.getter
    def location(self) -> pulumi.Output[Optional[str]]:
        """
        Resource location.
        """
        return pulumi.get(self, "location")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        Resource name.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def peerings(self) -> pulumi.Output[Optional[Sequence['outputs.ExpressRouteCircuitPeeringResponse']]]:
        """
        The list of peerings.
        """
        return pulumi.get(self, "peerings")

    @property
    @pulumi.getter(name="provisioningState")
    def provisioning_state(self) -> pulumi.Output[Optional[str]]:
        """
        Gets the provisioning state of the public IP resource. Possible values are: 'Updating', 'Deleting', and 'Failed'.
        """
        return pulumi.get(self, "provisioning_state")

    @property
    @pulumi.getter(name="serviceKey")
    def service_key(self) -> pulumi.Output[Optional[str]]:
        """
        The ServiceKey.
        """
        return pulumi.get(self, "service_key")

    @property
    @pulumi.getter(name="serviceProviderNotes")
    def service_provider_notes(self) -> pulumi.Output[Optional[str]]:
        """
        The ServiceProviderNotes.
        """
        return pulumi.get(self, "service_provider_notes")

    @property
    @pulumi.getter(name="serviceProviderProperties")
    def service_provider_properties(self) -> pulumi.Output[Optional['outputs.ExpressRouteCircuitServiceProviderPropertiesResponse']]:
        """
        The ServiceProviderProperties.
        """
        return pulumi.get(self, "service_provider_properties")

    @property
    @pulumi.getter(name="serviceProviderProvisioningState")
    def service_provider_provisioning_state(self) -> pulumi.Output[Optional[str]]:
        """
        The ServiceProviderProvisioningState state of the resource. Possible values are 'NotProvisioned', 'Provisioning', 'Provisioned', and 'Deprovisioning'.
        """
        return pulumi.get(self, "service_provider_provisioning_state")

    @property
    @pulumi.getter
    def sku(self) -> pulumi.Output[Optional['outputs.ExpressRouteCircuitSkuResponse']]:
        """
        The SKU.
        """
        return pulumi.get(self, "sku")

    @property
    @pulumi.getter
    def tags(self) -> pulumi.Output[Optional[Mapping[str, str]]]:
        """
        Resource tags.
        """
        return pulumi.get(self, "tags")

    @property
    @pulumi.getter
    def type(self) -> pulumi.Output[str]:
        """
        Resource type.
        """
        return pulumi.get(self, "type")

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

