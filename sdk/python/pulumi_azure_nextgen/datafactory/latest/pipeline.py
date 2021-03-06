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

__all__ = ['Pipeline']


class Pipeline(pulumi.CustomResource):
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 activities: Optional[pulumi.Input[Sequence[pulumi.Input[Union[pulumi.InputType['ControlActivityArgs'], pulumi.InputType['ExecutionActivityArgs']]]]]] = None,
                 annotations: Optional[pulumi.Input[Sequence[Any]]] = None,
                 concurrency: Optional[pulumi.Input[int]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 factory_name: Optional[pulumi.Input[str]] = None,
                 folder: Optional[pulumi.Input[pulumi.InputType['PipelineFolderArgs']]] = None,
                 parameters: Optional[pulumi.Input[Mapping[str, pulumi.Input[pulumi.InputType['ParameterSpecificationArgs']]]]] = None,
                 pipeline_name: Optional[pulumi.Input[str]] = None,
                 resource_group_name: Optional[pulumi.Input[str]] = None,
                 run_dimensions: Optional[pulumi.Input[Mapping[str, Any]]] = None,
                 variables: Optional[pulumi.Input[Mapping[str, pulumi.Input[pulumi.InputType['VariableSpecificationArgs']]]]] = None,
                 __props__=None,
                 __name__=None,
                 __opts__=None):
        """
        Pipeline resource type.

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[Sequence[pulumi.Input[Union[pulumi.InputType['ControlActivityArgs'], pulumi.InputType['ExecutionActivityArgs']]]]] activities: List of activities in pipeline.
        :param pulumi.Input[Sequence[Any]] annotations: List of tags that can be used for describing the Pipeline.
        :param pulumi.Input[int] concurrency: The max number of concurrent runs for the pipeline.
        :param pulumi.Input[str] description: The description of the pipeline.
        :param pulumi.Input[str] factory_name: The factory name.
        :param pulumi.Input[pulumi.InputType['PipelineFolderArgs']] folder: The folder that this Pipeline is in. If not specified, Pipeline will appear at the root level.
        :param pulumi.Input[Mapping[str, pulumi.Input[pulumi.InputType['ParameterSpecificationArgs']]]] parameters: List of parameters for pipeline.
        :param pulumi.Input[str] pipeline_name: The pipeline name.
        :param pulumi.Input[str] resource_group_name: The resource group name.
        :param pulumi.Input[Mapping[str, Any]] run_dimensions: Dimensions emitted by Pipeline.
        :param pulumi.Input[Mapping[str, pulumi.Input[pulumi.InputType['VariableSpecificationArgs']]]] variables: List of variables for pipeline.
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

            __props__['activities'] = activities
            __props__['annotations'] = annotations
            __props__['concurrency'] = concurrency
            __props__['description'] = description
            if factory_name is None:
                raise TypeError("Missing required property 'factory_name'")
            __props__['factory_name'] = factory_name
            __props__['folder'] = folder
            __props__['parameters'] = parameters
            if pipeline_name is None:
                raise TypeError("Missing required property 'pipeline_name'")
            __props__['pipeline_name'] = pipeline_name
            if resource_group_name is None:
                raise TypeError("Missing required property 'resource_group_name'")
            __props__['resource_group_name'] = resource_group_name
            __props__['run_dimensions'] = run_dimensions
            __props__['variables'] = variables
            __props__['etag'] = None
            __props__['name'] = None
            __props__['type'] = None
        alias_opts = pulumi.ResourceOptions(aliases=[pulumi.Alias(type_="azure-nextgen:datafactory/v20170901preview:Pipeline"), pulumi.Alias(type_="azure-nextgen:datafactory/v20180601:Pipeline")])
        opts = pulumi.ResourceOptions.merge(opts, alias_opts)
        super(Pipeline, __self__).__init__(
            'azure-nextgen:datafactory/latest:Pipeline',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None) -> 'Pipeline':
        """
        Get an existing Pipeline resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        return Pipeline(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def activities(self) -> pulumi.Output[Optional[Sequence[Any]]]:
        """
        List of activities in pipeline.
        """
        return pulumi.get(self, "activities")

    @property
    @pulumi.getter
    def annotations(self) -> pulumi.Output[Optional[Sequence[Any]]]:
        """
        List of tags that can be used for describing the Pipeline.
        """
        return pulumi.get(self, "annotations")

    @property
    @pulumi.getter
    def concurrency(self) -> pulumi.Output[Optional[int]]:
        """
        The max number of concurrent runs for the pipeline.
        """
        return pulumi.get(self, "concurrency")

    @property
    @pulumi.getter
    def description(self) -> pulumi.Output[Optional[str]]:
        """
        The description of the pipeline.
        """
        return pulumi.get(self, "description")

    @property
    @pulumi.getter
    def etag(self) -> pulumi.Output[str]:
        """
        Etag identifies change in the resource.
        """
        return pulumi.get(self, "etag")

    @property
    @pulumi.getter
    def folder(self) -> pulumi.Output[Optional['outputs.PipelineResponseFolder']]:
        """
        The folder that this Pipeline is in. If not specified, Pipeline will appear at the root level.
        """
        return pulumi.get(self, "folder")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        The resource name.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def parameters(self) -> pulumi.Output[Optional[Mapping[str, 'outputs.ParameterSpecificationResponse']]]:
        """
        List of parameters for pipeline.
        """
        return pulumi.get(self, "parameters")

    @property
    @pulumi.getter(name="runDimensions")
    def run_dimensions(self) -> pulumi.Output[Optional[Mapping[str, Any]]]:
        """
        Dimensions emitted by Pipeline.
        """
        return pulumi.get(self, "run_dimensions")

    @property
    @pulumi.getter
    def type(self) -> pulumi.Output[str]:
        """
        The resource type.
        """
        return pulumi.get(self, "type")

    @property
    @pulumi.getter
    def variables(self) -> pulumi.Output[Optional[Mapping[str, 'outputs.VariableSpecificationResponse']]]:
        """
        List of variables for pipeline.
        """
        return pulumi.get(self, "variables")

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

