// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../../types/input";
import * as outputs from "../../types/output";
import * as utilities from "../../utilities";

export function getApplicationGatewayPrivateEndpointConnection(args: GetApplicationGatewayPrivateEndpointConnectionArgs, opts?: pulumi.InvokeOptions): Promise<GetApplicationGatewayPrivateEndpointConnectionResult> {
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    return pulumi.runtime.invoke("azure-nextgen:network/latest:getApplicationGatewayPrivateEndpointConnection", {
        "applicationGatewayName": args.applicationGatewayName,
        "connectionName": args.connectionName,
        "resourceGroupName": args.resourceGroupName,
    }, opts);
}

export interface GetApplicationGatewayPrivateEndpointConnectionArgs {
    /**
     * The name of the application gateway.
     */
    readonly applicationGatewayName: string;
    /**
     * The name of the application gateway private endpoint connection.
     */
    readonly connectionName: string;
    /**
     * The name of the resource group.
     */
    readonly resourceGroupName: string;
}

/**
 * Private Endpoint connection on an application gateway.
 */
export interface GetApplicationGatewayPrivateEndpointConnectionResult {
    /**
     * A unique read-only string that changes whenever the resource is updated.
     */
    readonly etag: string;
    /**
     * The consumer link id.
     */
    readonly linkIdentifier: string;
    /**
     * Name of the private endpoint connection on an application gateway.
     */
    readonly name?: string;
    /**
     * The resource of private end point.
     */
    readonly privateEndpoint: outputs.network.latest.PrivateEndpointResponse;
    /**
     * A collection of information about the state of the connection between service consumer and provider.
     */
    readonly privateLinkServiceConnectionState?: outputs.network.latest.PrivateLinkServiceConnectionStateResponse;
    /**
     * The provisioning state of the application gateway private endpoint connection resource.
     */
    readonly provisioningState: string;
    /**
     * Type of the resource.
     */
    readonly type: string;
}
