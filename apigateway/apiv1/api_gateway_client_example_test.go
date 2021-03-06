// Copyright 2021 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     https://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by protoc-gen-go_gapic. DO NOT EDIT.

package apigateway_test

import (
	"context"

	apigateway "cloud.google.com/go/apigateway/apiv1"
	"google.golang.org/api/iterator"
	apigatewaypb "google.golang.org/genproto/googleapis/cloud/apigateway/v1"
)

func ExampleNewClient() {
	ctx := context.Background()
	c, err := apigateway.NewClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use client.
	_ = c
}

func ExampleClient_ListGateways() {
	// import apigatewaypb "google.golang.org/genproto/googleapis/cloud/apigateway/v1"
	// import "google.golang.org/api/iterator"

	ctx := context.Background()
	c, err := apigateway.NewClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}

	req := &apigatewaypb.ListGatewaysRequest{
		// TODO: Fill request struct fields.
	}
	it := c.ListGateways(ctx, req)
	for {
		resp, err := it.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			// TODO: Handle error.
		}
		// TODO: Use resp.
		_ = resp
	}
}

func ExampleClient_GetGateway() {
	// import apigatewaypb "google.golang.org/genproto/googleapis/cloud/apigateway/v1"

	ctx := context.Background()
	c, err := apigateway.NewClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}

	req := &apigatewaypb.GetGatewayRequest{
		// TODO: Fill request struct fields.
	}
	resp, err := c.GetGateway(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}

func ExampleClient_CreateGateway() {
	// import apigatewaypb "google.golang.org/genproto/googleapis/cloud/apigateway/v1"

	ctx := context.Background()
	c, err := apigateway.NewClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}

	req := &apigatewaypb.CreateGatewayRequest{
		// TODO: Fill request struct fields.
	}
	op, err := c.CreateGateway(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}

	resp, err := op.Wait(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}

func ExampleClient_UpdateGateway() {
	// import apigatewaypb "google.golang.org/genproto/googleapis/cloud/apigateway/v1"

	ctx := context.Background()
	c, err := apigateway.NewClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}

	req := &apigatewaypb.UpdateGatewayRequest{
		// TODO: Fill request struct fields.
	}
	op, err := c.UpdateGateway(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}

	resp, err := op.Wait(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}

func ExampleClient_DeleteGateway() {
	// import apigatewaypb "google.golang.org/genproto/googleapis/cloud/apigateway/v1"

	ctx := context.Background()
	c, err := apigateway.NewClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}

	req := &apigatewaypb.DeleteGatewayRequest{
		// TODO: Fill request struct fields.
	}
	op, err := c.DeleteGateway(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}

	err = op.Wait(ctx)
	if err != nil {
		// TODO: Handle error.
	}
}

func ExampleClient_ListApis() {
	// import apigatewaypb "google.golang.org/genproto/googleapis/cloud/apigateway/v1"
	// import "google.golang.org/api/iterator"

	ctx := context.Background()
	c, err := apigateway.NewClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}

	req := &apigatewaypb.ListApisRequest{
		// TODO: Fill request struct fields.
	}
	it := c.ListApis(ctx, req)
	for {
		resp, err := it.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			// TODO: Handle error.
		}
		// TODO: Use resp.
		_ = resp
	}
}

func ExampleClient_GetApi() {
	// import apigatewaypb "google.golang.org/genproto/googleapis/cloud/apigateway/v1"

	ctx := context.Background()
	c, err := apigateway.NewClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}

	req := &apigatewaypb.GetApiRequest{
		// TODO: Fill request struct fields.
	}
	resp, err := c.GetApi(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}

func ExampleClient_CreateApi() {
	// import apigatewaypb "google.golang.org/genproto/googleapis/cloud/apigateway/v1"

	ctx := context.Background()
	c, err := apigateway.NewClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}

	req := &apigatewaypb.CreateApiRequest{
		// TODO: Fill request struct fields.
	}
	op, err := c.CreateApi(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}

	resp, err := op.Wait(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}

func ExampleClient_UpdateApi() {
	// import apigatewaypb "google.golang.org/genproto/googleapis/cloud/apigateway/v1"

	ctx := context.Background()
	c, err := apigateway.NewClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}

	req := &apigatewaypb.UpdateApiRequest{
		// TODO: Fill request struct fields.
	}
	op, err := c.UpdateApi(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}

	resp, err := op.Wait(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}

func ExampleClient_DeleteApi() {
	// import apigatewaypb "google.golang.org/genproto/googleapis/cloud/apigateway/v1"

	ctx := context.Background()
	c, err := apigateway.NewClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}

	req := &apigatewaypb.DeleteApiRequest{
		// TODO: Fill request struct fields.
	}
	op, err := c.DeleteApi(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}

	err = op.Wait(ctx)
	if err != nil {
		// TODO: Handle error.
	}
}

func ExampleClient_ListApiConfigs() {
	// import apigatewaypb "google.golang.org/genproto/googleapis/cloud/apigateway/v1"
	// import "google.golang.org/api/iterator"

	ctx := context.Background()
	c, err := apigateway.NewClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}

	req := &apigatewaypb.ListApiConfigsRequest{
		// TODO: Fill request struct fields.
	}
	it := c.ListApiConfigs(ctx, req)
	for {
		resp, err := it.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			// TODO: Handle error.
		}
		// TODO: Use resp.
		_ = resp
	}
}

func ExampleClient_GetApiConfig() {
	// import apigatewaypb "google.golang.org/genproto/googleapis/cloud/apigateway/v1"

	ctx := context.Background()
	c, err := apigateway.NewClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}

	req := &apigatewaypb.GetApiConfigRequest{
		// TODO: Fill request struct fields.
	}
	resp, err := c.GetApiConfig(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}

func ExampleClient_CreateApiConfig() {
	// import apigatewaypb "google.golang.org/genproto/googleapis/cloud/apigateway/v1"

	ctx := context.Background()
	c, err := apigateway.NewClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}

	req := &apigatewaypb.CreateApiConfigRequest{
		// TODO: Fill request struct fields.
	}
	op, err := c.CreateApiConfig(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}

	resp, err := op.Wait(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}

func ExampleClient_UpdateApiConfig() {
	// import apigatewaypb "google.golang.org/genproto/googleapis/cloud/apigateway/v1"

	ctx := context.Background()
	c, err := apigateway.NewClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}

	req := &apigatewaypb.UpdateApiConfigRequest{
		// TODO: Fill request struct fields.
	}
	op, err := c.UpdateApiConfig(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}

	resp, err := op.Wait(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}

func ExampleClient_DeleteApiConfig() {
	// import apigatewaypb "google.golang.org/genproto/googleapis/cloud/apigateway/v1"

	ctx := context.Background()
	c, err := apigateway.NewClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}

	req := &apigatewaypb.DeleteApiConfigRequest{
		// TODO: Fill request struct fields.
	}
	op, err := c.DeleteApiConfig(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}

	err = op.Wait(ctx)
	if err != nil {
		// TODO: Handle error.
	}
}
