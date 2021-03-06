// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../../types/input";
import * as outputs from "../../types/output";
import * as utilities from "../../utilities";

export function getJob(args: GetJobArgs, opts?: pulumi.InvokeOptions): Promise<GetJobResult> {
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    return pulumi.runtime.invoke("azure-nextgen:batchai/v20180301:getJob", {
        "jobName": args.jobName,
        "resourceGroupName": args.resourceGroupName,
    }, opts);
}

export interface GetJobArgs {
    /**
     * The name of the job within the specified resource group. Job names can only contain a combination of alphanumeric characters along with dash (-) and underscore (_). The name must be from 1 through 64 characters long.
     */
    readonly jobName: string;
    /**
     * Name of the resource group to which the resource belongs.
     */
    readonly resourceGroupName: string;
}

/**
 * Contains information about the job.
 */
export interface GetJobResult {
    /**
     * Specifies the settings for Caffe job.
     */
    readonly caffeSettings?: outputs.batchai.v20180301.CaffeSettingsResponse;
    /**
     * Specifies the settings for Chainer job.
     */
    readonly chainerSettings?: outputs.batchai.v20180301.ChainerSettingsResponse;
    /**
     * Represents a resource ID. For example, for a subnet, it is the resource URL for the subnet.
     */
    readonly cluster?: outputs.batchai.v20180301.ResourceIdResponse;
    /**
     * Specifies the settings for CNTK (aka Microsoft Cognitive Toolkit) job.
     */
    readonly cntkSettings?: outputs.batchai.v20180301.CNTKsettingsResponse;
    /**
     * Constraints associated with the Job.
     */
    readonly constraints?: outputs.batchai.v20180301.JobPropertiesResponseConstraints;
    /**
     * If the container was downloaded as part of cluster setup then the same container image will be used. If not provided, the job will run on the VM.
     */
    readonly containerSettings?: outputs.batchai.v20180301.ContainerSettingsResponse;
    /**
     * The creation time of the job.
     */
    readonly creationTime: string;
    /**
     * Specifies the settings for a custom tool kit job.
     */
    readonly customToolkitSettings?: outputs.batchai.v20180301.CustomToolkitSettingsResponse;
    /**
     * Batch AI will setup these additional environment variables for the job.
     */
    readonly environmentVariables?: outputs.batchai.v20180301.EnvironmentVariableResponse[];
    /**
     * Contains information about the execution of a job in the Azure Batch service.
     */
    readonly executionInfo?: outputs.batchai.v20180301.JobPropertiesResponseExecutionInfo;
    /**
     * The current state of the job. Possible values are: queued - The job is queued and able to run. A job enters this state when it is created, or when it is awaiting a retry after a failed run. running - The job is running on a compute cluster. This includes job-level preparation such as downloading resource files or set up container specified on the job - it does not necessarily mean that the job command line has started executing. terminating - The job is terminated by the user, the terminate operation is in progress. succeeded - The job has completed running successfully and exited with exit code 0. failed - The job has finished unsuccessfully (failed with a non-zero exit code) and has exhausted its retry limit. A job is also marked as failed if an error occurred launching the job.
     */
    readonly executionState?: string;
    /**
     * The time at which the job entered its current execution state.
     */
    readonly executionStateTransitionTime: string;
    /**
     * Describe the experiment information of the job
     */
    readonly experimentName?: string;
    readonly inputDirectories?: outputs.batchai.v20180301.InputDirectoryResponse[];
    /**
     * Batch AI creates job's output directories under an unique path to avoid conflicts between jobs. This value contains a path segment generated by Batch AI to make the path unique and can be used to find the output directory on the node or mounted filesystem.
     */
    readonly jobOutputDirectoryPathSegment?: string;
    /**
     * The specified actions will run on all the nodes that are part of the job
     */
    readonly jobPreparation?: outputs.batchai.v20180301.JobPreparationResponse;
    /**
     * The location of the resource
     */
    readonly location: string;
    /**
     * These volumes will be mounted before the job execution and will be unmounted after the job completion. The volumes will be mounted at location specified by $AZ_BATCHAI_JOB_MOUNT_ROOT environment variable.
     */
    readonly mountVolumes?: outputs.batchai.v20180301.MountVolumesResponse;
    /**
     * The name of the resource
     */
    readonly name: string;
    /**
     * The job will be gang scheduled on that many compute nodes
     */
    readonly nodeCount?: number;
    readonly outputDirectories?: outputs.batchai.v20180301.OutputDirectoryResponse[];
    /**
     * Priority associated with the job. Priority values can range from -1000 to 1000, with -1000 being the lowest priority and 1000 being the highest priority. The default value is 0.
     */
    readonly priority?: number;
    /**
     * The provisioned state of the Batch AI job
     */
    readonly provisioningState: string;
    /**
     * The time at which the job entered its current provisioning state.
     */
    readonly provisioningStateTransitionTime: string;
    /**
     * Specifies the settings for pyTorch job.
     */
    readonly pyTorchSettings?: outputs.batchai.v20180301.PyTorchSettingsResponse;
    /**
     * Batch AI will setup these additional environment variables for the job. Server will never report values of these variables back.
     */
    readonly secrets?: outputs.batchai.v20180301.EnvironmentVariableWithSecretValueResponse[];
    /**
     * The path where the Batch AI service will upload stdout and stderror of the job.
     */
    readonly stdOutErrPathPrefix?: string;
    /**
     * The tags of the resource
     */
    readonly tags: {[key: string]: string};
    /**
     * Specifies the settings for TensorFlow job.
     */
    readonly tensorFlowSettings?: outputs.batchai.v20180301.TensorFlowSettingsResponse;
    /**
     * Possible values are: cntk, tensorflow, caffe, caffe2, chainer, pytorch, custom.
     */
    readonly toolType?: string;
    /**
     * The type of the resource
     */
    readonly type: string;
}
