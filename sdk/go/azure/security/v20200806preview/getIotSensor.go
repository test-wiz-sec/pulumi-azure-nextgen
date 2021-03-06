// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v20200806preview

import (
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

func LookupIotSensor(ctx *pulumi.Context, args *LookupIotSensorArgs, opts ...pulumi.InvokeOption) (*LookupIotSensorResult, error) {
	var rv LookupIotSensorResult
	err := ctx.Invoke("azure-nextgen:security/v20200806preview:getIotSensor", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupIotSensorArgs struct {
	// Name of the IoT sensor
	IotSensorName string `pulumi:"iotSensorName"`
	// Scope of the query (IoT Hub, /providers/Microsoft.Devices/iotHubs/myHub)
	Scope string `pulumi:"scope"`
}

// IoT sensor
type LookupIotSensorResult struct {
	// Resource name
	Name string `pulumi:"name"`
	// Resource type
	Type string `pulumi:"type"`
}
