//
// Copyright (c) 2018 Volterra, Inc. All rights reserved.
// Code generated by ves-gen-schema-go. DO NOT EDIT.
//

package azure_vnet_site

import (
	"bytes"
	"context"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/gogo/protobuf/proto"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"google.golang.org/grpc"

	"gopkg.volterra.us/stdlib/client"
	"gopkg.volterra.us/stdlib/codec"
	"gopkg.volterra.us/stdlib/errors"
	"gopkg.volterra.us/stdlib/server"
	"gopkg.volterra.us/stdlib/svcfw"
)

var (
	_ = fmt.Sprintf("dummy for fmt import use")
)

// Create CustomAPI GRPC Client satisfying server.CustomClient
type CustomAPIGrpcClient struct {
	conn       *grpc.ClientConn
	grpcClient CustomAPIClient
	// map of rpc name to its invocation
	rpcFns map[string]func(context.Context, string, ...grpc.CallOption) (proto.Message, error)
}

func (c *CustomAPIGrpcClient) doRPCSetVIPInfo(ctx context.Context, yamlReq string, opts ...grpc.CallOption) (proto.Message, error) {
	req := &SetVIPInfoRequest{}
	if err := codec.FromYAML(yamlReq, req); err != nil {
		return nil, fmt.Errorf("YAML Request %s is not of type *ves.io.schema.views.azure_vnet_site.SetVIPInfoRequest", yamlReq)
	}
	rsp, err := c.grpcClient.SetVIPInfo(ctx, req, opts...)
	return rsp, err
}

func (c *CustomAPIGrpcClient) DoRPC(ctx context.Context, rpc string, opts ...server.CustomCallOpt) (proto.Message, error) {
	rpcFn, exists := c.rpcFns[rpc]
	if !exists {
		return nil, fmt.Errorf("Error, no such rpc %s", rpc)
	}
	cco := server.NewCustomCallOpts()
	for _, opt := range opts {
		opt(cco)
	}
	if cco.YAMLReq == "" {
		return nil, fmt.Errorf("Error, empty request body")
	}
	ctx = client.AddHdrsToCtx(cco.Headers, ctx)

	rsp, err := rpcFn(ctx, cco.YAMLReq, cco.GrpcCallOpts...)
	if err != nil {
		return nil, errors.Wrap(err, "Doing custom RPC using GRPC")
	}
	if cco.OutCallResponse != nil {
		cco.OutCallResponse.ProtoMsg = rsp
	}
	return rsp, nil
}

func NewCustomAPIGrpcClient(cc *grpc.ClientConn) server.CustomClient {
	ccl := &CustomAPIGrpcClient{
		conn:       cc,
		grpcClient: NewCustomAPIClient(cc),
	}
	rpcFns := make(map[string]func(context.Context, string, ...grpc.CallOption) (proto.Message, error))
	rpcFns["SetVIPInfo"] = ccl.doRPCSetVIPInfo

	ccl.rpcFns = rpcFns

	return ccl
}

// Create CustomAPI REST Client satisfying server.CustomClient
type CustomAPIRestClient struct {
	baseURL string
	client  http.Client
	// map of rpc name to its invocation
	rpcFns map[string]func(context.Context, *server.CustomCallOpts) (proto.Message, error)
}

func (c *CustomAPIRestClient) doRPCSetVIPInfo(ctx context.Context, callOpts *server.CustomCallOpts) (proto.Message, error) {
	if callOpts.URI == "" {
		return nil, fmt.Errorf("Error, URI should be specified, got empty")
	}
	url := fmt.Sprintf("%s%s", c.baseURL, callOpts.URI)

	yamlReq := callOpts.YAMLReq
	req := &SetVIPInfoRequest{}
	if err := codec.FromYAML(yamlReq, req); err != nil {
		return nil, fmt.Errorf("YAML Request %s is not of type *ves.io.schema.views.azure_vnet_site.SetVIPInfoRequest: %s", yamlReq, err)
	}

	var hReq *http.Request
	hm := strings.ToLower(callOpts.HTTPMethod)
	switch hm {
	case "post", "put":
		jsn, err := req.ToJSON()
		if err != nil {
			return nil, errors.Wrap(err, "Custom RestClient converting YAML to JSON")
		}
		var op string
		if hm == "post" {
			op = http.MethodPost
		} else {
			op = http.MethodPut
		}
		newReq, err := http.NewRequest(op, url, bytes.NewBuffer([]byte(jsn)))
		if err != nil {
			return nil, errors.Wrapf(err, "Creating new HTTP %s request for custom API", op)
		}
		hReq = newReq
	case "get":
		newReq, err := http.NewRequest(http.MethodGet, url, nil)
		if err != nil {
			return nil, errors.Wrap(err, "Creating new HTTP GET request for custom API")
		}
		hReq = newReq
		q := hReq.URL.Query()
		_ = q
		q.Add("name", fmt.Sprintf("%v", req.Name))
		q.Add("namespace", fmt.Sprintf("%v", req.Namespace))
		q.Add("vip_params_per_az", fmt.Sprintf("%v", req.VipParamsPerAz))

		hReq.URL.RawQuery += q.Encode()
	case "delete":
		newReq, err := http.NewRequest(http.MethodDelete, url, nil)
		if err != nil {
			return nil, errors.Wrap(err, "Creating new HTTP DELETE request for custom API")
		}
		hReq = newReq
	default:
		return nil, fmt.Errorf("Error, invalid/empty HTTPMethod(%s) specified, should be POST|DELETE|GET", callOpts.HTTPMethod)
	}
	hReq = hReq.WithContext(ctx)
	hReq.Header.Set("Content-Type", "application/json")
	client.AddHdrsToReq(callOpts.Headers, hReq)

	rsp, err := c.client.Do(hReq)
	if err != nil {
		return nil, errors.Wrap(err, "Custom API RestClient")
	}
	defer rsp.Body.Close()

	if rsp.StatusCode != http.StatusOK {
		body, err := ioutil.ReadAll(rsp.Body)
		return nil, fmt.Errorf("Unsuccessful custom API %s on %s, status code %d, body %s, err %s", callOpts.HTTPMethod, callOpts.URI, rsp.StatusCode, body, err)
	}

	body, err := ioutil.ReadAll(rsp.Body)
	if err != nil {
		return nil, errors.Wrap(err, "Custom API RestClient read body")
	}
	pbRsp := &SetVIPInfoResponse{}
	if err := codec.FromJSON(string(body), pbRsp); err != nil {
		return nil, fmt.Errorf("JSON Response %s is not of type *ves.io.schema.views.azure_vnet_site.SetVIPInfoResponse", body)

	}
	if callOpts.OutCallResponse != nil {
		callOpts.OutCallResponse.ProtoMsg = pbRsp
		callOpts.OutCallResponse.JSON = string(body)
	}
	return pbRsp, nil
}

func (c *CustomAPIRestClient) DoRPC(ctx context.Context, rpc string, opts ...server.CustomCallOpt) (proto.Message, error) {
	rpcFn, exists := c.rpcFns[rpc]
	if !exists {
		return nil, fmt.Errorf("Error, no such rpc %s", rpc)
	}
	cco := server.NewCustomCallOpts()
	for _, opt := range opts {
		opt(cco)
	}

	rsp, err := rpcFn(ctx, cco)
	if err != nil {
		return nil, errors.Wrap(err, "Doing custom RPC using Rest")
	}
	return rsp, nil
}

func NewCustomAPIRestClient(baseURL string, hc http.Client) server.CustomClient {
	ccl := &CustomAPIRestClient{
		baseURL: baseURL,
		client:  hc,
	}

	rpcFns := make(map[string]func(context.Context, *server.CustomCallOpts) (proto.Message, error))
	rpcFns["SetVIPInfo"] = ccl.doRPCSetVIPInfo

	ccl.rpcFns = rpcFns

	return ccl
}

// Create CustomAPIInprocClient

// INPROC Client (satisfying CustomAPIClient interface)
type CustomAPIInprocClient struct {
	svc svcfw.Service
}

func (c *CustomAPIInprocClient) SetVIPInfo(ctx context.Context, in *SetVIPInfoRequest, opts ...grpc.CallOption) (*SetVIPInfoResponse, error) {
	ah := c.svc.GetAPIHandler("ves.io.schema.views.azure_vnet_site.CustomAPI")
	cah, ok := ah.(CustomAPIServer)
	if !ok {
		return nil, fmt.Errorf("ah %v is not of type *CustomAPISrv", ah)
	}

	var (
		rsp *SetVIPInfoResponse
		err error
	)

	bodyFields := svcfw.GenAuditReqBodyFields(ctx, c.svc, "ves.io.schema.views.azure_vnet_site.SetVIPInfoRequest", in)
	defer func() {
		if len(bodyFields) > 0 {
			server.ExtendAPIAudit(ctx, svcfw.PublicAPIBodyLog.Uid, bodyFields)
		}
		userMsg := "The 'CustomAPI.SetVIPInfo' operation on 'azure_vnet_site'"
		if err == nil {
			userMsg += " was successfully performed."
		} else {
			userMsg += " failed to be performed."
		}
		server.AddUserMsgToAPIAudit(ctx, userMsg)
	}()

	if c.svc.Config().EnableAPIValidation {
		if rvFn := c.svc.GetRPCValidator("ves.io.schema.views.azure_vnet_site.CustomAPI.SetVIPInfo"); rvFn != nil {
			if verr := rvFn(ctx, in); verr != nil {
				err = server.MaybePublicRestError(ctx, errors.Wrapf(verr, "Validating Request"))
				return nil, server.GRPCStatusFromError(err).Err()
			}
		}
	}

	rsp, err = cah.SetVIPInfo(ctx, in)
	if err != nil {
		return rsp, server.GRPCStatusFromError(server.MaybePublicRestError(ctx, err)).Err()
	}

	bodyFields = append(bodyFields, svcfw.GenAuditRspBodyFields(ctx, c.svc, "ves.io.schema.views.azure_vnet_site.SetVIPInfoResponse", rsp)...)

	return rsp, nil
}

func NewCustomAPIInprocClient(svc svcfw.Service) CustomAPIClient {
	return &CustomAPIInprocClient{svc: svc}
}

// RegisterGwCustomAPIHandler registers with grpc-gw with an inproc-client backing so that
// rest to grpc happens without a grpc.Dial (thus avoiding additional certs for mTLS)
func RegisterGwCustomAPIHandler(ctx context.Context, mux *runtime.ServeMux, svc interface{}) error {
	s, ok := svc.(svcfw.Service)
	if !ok {
		return fmt.Errorf("svc is not svcfw.Service")
	}
	return RegisterCustomAPIHandlerClient(ctx, mux, NewCustomAPIInprocClient(s))
}

var CustomAPISwaggerJSON string = `{
    "swagger": "2.0",
    "info": {
        "title": "Azure Vnet site",
        "description": "Azure Vnet site view defines a required parameters that can be used in CRUD, to create and manage a volterra site in Azure Vnet.\nIt can be used to either automatically create or Manually assisted site creation in Azure Vnet.\n\nView will create following child objects.\n\n* Site",
        "version": "version not set"
    },
    "schemes": [
        "http",
        "https"
    ],
    "consumes": [
        "application/json"
    ],
    "produces": [
        "application/json"
    ],
    "tags": [],
    "paths": {
        "/public/namespaces/{namespace}/azure_vnet_site/{name}/set_vip_info": {
            "post": {
                "summary": "Configure Azure Vnet Site VIP Information",
                "description": "Configure Azure Vnet Site VIP Information",
                "operationId": "ves.io.schema.views.azure_vnet_site.CustomAPI.SetVIPInfo",
                "responses": {
                    "200": {
                        "description": "A successful response.",
                        "schema": {
                            "$ref": "#/definitions/azure_vnet_siteSetVIPInfoResponse"
                        }
                    },
                    "401": {
                        "description": "Returned when operation is not authorized",
                        "schema": {
                            "format": "string"
                        }
                    },
                    "403": {
                        "description": "Returned when there is no permission to access resource",
                        "schema": {
                            "format": "string"
                        }
                    },
                    "404": {
                        "description": "Returned when resource is not found",
                        "schema": {
                            "format": "string"
                        }
                    },
                    "409": {
                        "description": "Returned when operation on resource is conflicting with current value",
                        "schema": {
                            "format": "string"
                        }
                    },
                    "429": {
                        "description": "Returned when operation has been rejected as it is happening too frequently",
                        "schema": {
                            "format": "string"
                        }
                    },
                    "500": {
                        "description": "Returned when server encountered an error in processing API",
                        "schema": {
                            "format": "string"
                        }
                    },
                    "503": {
                        "description": "Returned when service is unavailable temporarily",
                        "schema": {
                            "format": "string"
                        }
                    },
                    "504": {
                        "description": "Returned when server timed out processing request",
                        "schema": {
                            "format": "string"
                        }
                    }
                },
                "parameters": [
                    {
                        "name": "namespace",
                        "description": "Namespace\n\nx-example: \"default\"\nNamespace for the object to be configured",
                        "in": "path",
                        "required": true,
                        "type": "string",
                        "x-displayname": "Namespace"
                    },
                    {
                        "name": "name",
                        "description": "Name\n\nx-example: \"aws-vpc-site-1\"\nName of the object to be configured",
                        "in": "path",
                        "required": true,
                        "type": "string",
                        "x-displayname": "Name"
                    },
                    {
                        "name": "body",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/azure_vnet_siteSetVIPInfoRequest"
                        }
                    }
                ],
                "tags": [
                    "CustomAPI"
                ],
                "externalDocs": {
                    "description": "Examples of this operation",
                    "url": "https://www.volterra.io/docs/reference/api-ref/ves-io-schema-views-azure_vnet_site-CustomAPI-SetVIPInfo"
                },
                "x-ves-proto-rpc": "ves.io.schema.views.azure_vnet_site.CustomAPI.SetVIPInfo"
            },
            "x-displayname": "Custom API for Azure Vnet site",
            "x-ves-proto-service": "ves.io.schema.views.azure_vnet_site.CustomAPI",
            "x-ves-proto-service-type": "CUSTOM_PUBLIC"
        }
    },
    "definitions": {
        "azure_vnet_siteSetVIPInfoRequest": {
            "type": "object",
            "description": "Request to configure Azure Vnet Site VIP information",
            "title": "Request to configure Azure Vnet Site VIP information",
            "x-displayname": "Request to configure Azure Vnet Site VIP information",
            "x-ves-proto-message": "ves.io.schema.views.azure_vnet_site.SetVIPInfoRequest",
            "properties": {
                "name": {
                    "type": "string",
                    "description": " Name of the object to be configured\n\nExample: - \"aws-vpc-site-1\"-",
                    "title": "Name",
                    "x-displayname": "Name",
                    "x-ves-example": "aws-vpc-site-1"
                },
                "namespace": {
                    "type": "string",
                    "description": " Namespace for the object to be configured\n\nExample: - \"default\"-",
                    "title": "Namespace",
                    "x-displayname": "Namespace",
                    "x-ves-example": "default"
                },
                "vip_params_per_az": {
                    "type": "array",
                    "description": " VIP Parameters per AZ\n\nExample: - \"master-0\"-\n\nValidation Rules:\n  ves.io.schema.rules.repeated.num_items: 1,3\n",
                    "title": "VIP Params Per AZ",
                    "items": {
                        "$ref": "#/definitions/sitePublishVIPParamsPerAz"
                    },
                    "x-displayname": "VIP Params Per AZ",
                    "x-ves-example": "master-0",
                    "x-ves-validation-rules": {
                        "ves.io.schema.rules.repeated.num_items": "1,3"
                    }
                }
            }
        },
        "azure_vnet_siteSetVIPInfoResponse": {
            "type": "object",
            "title": "Response to configure Azure Vnet Site VIP Information",
            "x-displayname": "Response to configure Azure Vnet Site VIP Information",
            "x-ves-proto-message": "ves.io.schema.views.azure_vnet_site.SetVIPInfoResponse"
        },
        "sitePublishVIPParamsPerAz": {
            "type": "object",
            "description": "Per AZ parameters needed to publish VIP for public cloud sites",
            "title": "Publish VIP Params Per AZ",
            "x-displayname": "Publish VIP Params Per AZ",
            "x-ves-proto-message": "ves.io.schema.site.PublishVIPParamsPerAz",
            "properties": {
                "az_name": {
                    "type": "string",
                    "description": " Name of the Availability zone\n\nExample: - \"us-east-2a\"-\n\nValidation Rules:\n  ves.io.schema.rules.string.in: [\\\"ap-northeast-1a\\\",\\\"ap-northeast-1c\\\",\\\"ap-northeast-1d\\\",\\\"ap-southeast-1a\\\",\\\"ap-southeast-1b\\\",\\\"ap-southeast-1c\\\",\\\"eu-central-1a\\\",\\\"eu-central-1b\\\",\\\"eu-central-1c\\\",\\\"eu-west-1a\\\",\\\"eu-west-1b\\\",\\\"eu-west-1c\\\",\\\"eu-west-3a\\\",\\\"eu-west-3b\\\",\\\"eu-west-3c\\\",\\\"sa-east-1a\\\",\\\"sa-east-1b\\\",\\\"sa-east-1c\\\",\\\"us-east-1a\\\",\\\"us-east-1b\\\",\\\"us-east-1c\\\",\\\"us-east-1d\\\",\\\"us-east-1e\\\",\\\"us-east-1f\\\",\\\"us-east-2a\\\",\\\"us-east-2b\\\",\\\"us-east-2c\\\",\\\"us-west-2a\\\",\\\"us-west-2b\\\",\\\"us-west-2c\\\",\\\"us-west-2d\\\",\\\"ca-central-1a\\\",\\\"ca-central-1b\\\",\\\"ca-central-1d\\\",\\\"af-south-1a\\\",\\\"af-south-1b\\\",\\\"af-south-1c\\\",\\\"ap-east-1a\\\",\\\"ap-east-1b\\\",\\\"ap-east-1c\\\",\\\"ap-south-1a\\\",\\\"ap-south-1b\\\",\\\"ap-south-1c\\\",\\\"ap-northeast-2a\\\",\\\"ap-northeast-2b\\\",\\\"ap-northeast-2c\\\",\\\"ap-northeast-2d\\\",\\\"ap-southeast-2a\\\",\\\"ap-southeast-2b\\\",\\\"ap-southeast-2c\\\",\\\"eu-south-1a\\\",\\\"eu-south-1b\\\",\\\"eu-south-1c\\\",\\\"eu-north-1a\\\",\\\"eu-north-1b\\\",\\\"eu-north-1c\\\",\\\"eu-west-2a\\\",\\\"eu-west-2b\\\",\\\"eu-west-2c\\\",\\\"me-south-1a\\\",\\\"me-south-1b\\\",\\\"me-south-1c\\\",\\\"us-west-1a\\\",\\\"us-west-1c\\\",\\\"1\\\",\\\"2\\\",\\\"3\\\",\\\"asia-east1-a\\\",\\\"asia-east1-b\\\",\\\"asia-east1-c\\\",\\\"asia-east2-a\\\",\\\"asia-east2-b\\\",\\\"asia-east2-c\\\",\\\"asia-northeast1-a\\\",\\\"asia-northeast1-b\\\",\\\"asia-northeast1-c\\\",\\\"asia-northeast2-a\\\",\\\"asia-northeast2-b\\\",\\\"asia-northeast2-c\\\",\\\"asia-northeast3-a\\\",\\\"asia-northeast3-b\\\",\\\"asia-northeast3-c\\\",\\\"asia-south1-a\\\",\\\"asia-south1-b\\\",\\\"asia-south1-c\\\",\\\"asia-southeast1-a\\\",\\\"asia-southeast1-b\\\",\\\"asia-southeast1-c\\\",\\\"asia-southeast2-a\\\",\\\"asia-southeast2-b\\\",\\\"asia-southeast2-c\\\",\\\"australia-southeast1-a\\\",\\\"australia-southeast1-b\\\",\\\"australia-southeast1-c\\\",\\\"europe-north1-a\\\",\\\"europe-north1-b\\\",\\\"europe-north1-c\\\",\\\"europe-west1-b\\\",\\\"europe-west1-c\\\",\\\"europe-west1-d\\\",\\\"europe-west2-a\\\",\\\"europe-west2-b\\\",\\\"europe-west2-c\\\",\\\"europe-west3-a\\\",\\\"europe-west3-b\\\",\\\"europe-west3-c\\\",\\\"europe-west4-a\\\",\\\"europe-west4-b\\\",\\\"europe-west4-c\\\",\\\"europe-west6-a\\\",\\\"europe-west6-b\\\",\\\"europe-west6-c\\\",\\\"northamerica-northeast1-a\\\",\\\"northamerica-northeast1-b\\\",\\\"northamerica-northeast1-c\\\",\\\"southamerica-east1-a\\\",\\\"southamerica-east1-b\\\",\\\"southamerica-east1-c\\\",\\\"us-central1-a\\\",\\\"us-central1-b\\\",\\\"us-central1-c\\\",\\\"us-central1-f\\\",\\\"us-east1-b\\\",\\\"us-east1-c\\\",\\\"us-east1-d\\\",\\\"us-east4-a\\\",\\\"us-east4-b\\\",\\\"us-east4-c\\\",\\\"us-west1-a\\\",\\\"us-west1-b\\\",\\\"us-west1-c\\\",\\\"us-west2-a\\\",\\\"us-west2-b\\\",\\\"us-west2-c\\\",\\\"us-west3-a\\\",\\\"us-west3-b\\\",\\\"us-west3-c\\\",\\\"us-west4-a\\\",\\\"us-west4-b\\\",\\\"us-west4-c\\\"]\n",
                    "title": "AZ Name",
                    "x-displayname": "AZ Name",
                    "x-ves-example": "us-east-2a",
                    "x-ves-validation-rules": {
                        "ves.io.schema.rules.string.in": "[\\\"ap-northeast-1a\\\",\\\"ap-northeast-1c\\\",\\\"ap-northeast-1d\\\",\\\"ap-southeast-1a\\\",\\\"ap-southeast-1b\\\",\\\"ap-southeast-1c\\\",\\\"eu-central-1a\\\",\\\"eu-central-1b\\\",\\\"eu-central-1c\\\",\\\"eu-west-1a\\\",\\\"eu-west-1b\\\",\\\"eu-west-1c\\\",\\\"eu-west-3a\\\",\\\"eu-west-3b\\\",\\\"eu-west-3c\\\",\\\"sa-east-1a\\\",\\\"sa-east-1b\\\",\\\"sa-east-1c\\\",\\\"us-east-1a\\\",\\\"us-east-1b\\\",\\\"us-east-1c\\\",\\\"us-east-1d\\\",\\\"us-east-1e\\\",\\\"us-east-1f\\\",\\\"us-east-2a\\\",\\\"us-east-2b\\\",\\\"us-east-2c\\\",\\\"us-west-2a\\\",\\\"us-west-2b\\\",\\\"us-west-2c\\\",\\\"us-west-2d\\\",\\\"ca-central-1a\\\",\\\"ca-central-1b\\\",\\\"ca-central-1d\\\",\\\"af-south-1a\\\",\\\"af-south-1b\\\",\\\"af-south-1c\\\",\\\"ap-east-1a\\\",\\\"ap-east-1b\\\",\\\"ap-east-1c\\\",\\\"ap-south-1a\\\",\\\"ap-south-1b\\\",\\\"ap-south-1c\\\",\\\"ap-northeast-2a\\\",\\\"ap-northeast-2b\\\",\\\"ap-northeast-2c\\\",\\\"ap-northeast-2d\\\",\\\"ap-southeast-2a\\\",\\\"ap-southeast-2b\\\",\\\"ap-southeast-2c\\\",\\\"eu-south-1a\\\",\\\"eu-south-1b\\\",\\\"eu-south-1c\\\",\\\"eu-north-1a\\\",\\\"eu-north-1b\\\",\\\"eu-north-1c\\\",\\\"eu-west-2a\\\",\\\"eu-west-2b\\\",\\\"eu-west-2c\\\",\\\"me-south-1a\\\",\\\"me-south-1b\\\",\\\"me-south-1c\\\",\\\"us-west-1a\\\",\\\"us-west-1c\\\",\\\"1\\\",\\\"2\\\",\\\"3\\\",\\\"asia-east1-a\\\",\\\"asia-east1-b\\\",\\\"asia-east1-c\\\",\\\"asia-east2-a\\\",\\\"asia-east2-b\\\",\\\"asia-east2-c\\\",\\\"asia-northeast1-a\\\",\\\"asia-northeast1-b\\\",\\\"asia-northeast1-c\\\",\\\"asia-northeast2-a\\\",\\\"asia-northeast2-b\\\",\\\"asia-northeast2-c\\\",\\\"asia-northeast3-a\\\",\\\"asia-northeast3-b\\\",\\\"asia-northeast3-c\\\",\\\"asia-south1-a\\\",\\\"asia-south1-b\\\",\\\"asia-south1-c\\\",\\\"asia-southeast1-a\\\",\\\"asia-southeast1-b\\\",\\\"asia-southeast1-c\\\",\\\"asia-southeast2-a\\\",\\\"asia-southeast2-b\\\",\\\"asia-southeast2-c\\\",\\\"australia-southeast1-a\\\",\\\"australia-southeast1-b\\\",\\\"australia-southeast1-c\\\",\\\"europe-north1-a\\\",\\\"europe-north1-b\\\",\\\"europe-north1-c\\\",\\\"europe-west1-b\\\",\\\"europe-west1-c\\\",\\\"europe-west1-d\\\",\\\"europe-west2-a\\\",\\\"europe-west2-b\\\",\\\"europe-west2-c\\\",\\\"europe-west3-a\\\",\\\"europe-west3-b\\\",\\\"europe-west3-c\\\",\\\"europe-west4-a\\\",\\\"europe-west4-b\\\",\\\"europe-west4-c\\\",\\\"europe-west6-a\\\",\\\"europe-west6-b\\\",\\\"europe-west6-c\\\",\\\"northamerica-northeast1-a\\\",\\\"northamerica-northeast1-b\\\",\\\"northamerica-northeast1-c\\\",\\\"southamerica-east1-a\\\",\\\"southamerica-east1-b\\\",\\\"southamerica-east1-c\\\",\\\"us-central1-a\\\",\\\"us-central1-b\\\",\\\"us-central1-c\\\",\\\"us-central1-f\\\",\\\"us-east1-b\\\",\\\"us-east1-c\\\",\\\"us-east1-d\\\",\\\"us-east4-a\\\",\\\"us-east4-b\\\",\\\"us-east4-c\\\",\\\"us-west1-a\\\",\\\"us-west1-b\\\",\\\"us-west1-c\\\",\\\"us-west2-a\\\",\\\"us-west2-b\\\",\\\"us-west2-c\\\",\\\"us-west3-a\\\",\\\"us-west3-b\\\",\\\"us-west3-c\\\",\\\"us-west4-a\\\",\\\"us-west4-b\\\",\\\"us-west4-c\\\"]"
                    }
                },
                "inside_vip": {
                    "type": "array",
                    "description": " List of Inside VIPs for an AZ\n\nExample: - \"192.168.0.156\"-\n\nValidation Rules:\n  ves.io.schema.rules.repeated.items.string.ipv4: true\n  ves.io.schema.rules.repeated.max_items: 3\n  ves.io.schema.rules.repeated.unique: true\n",
                    "title": "Inside VIP(s)",
                    "maxItems": 3,
                    "items": {
                        "type": "string"
                    },
                    "x-displayname": "Inside VIP(s)",
                    "x-ves-example": "192.168.0.156",
                    "x-ves-validation-rules": {
                        "ves.io.schema.rules.repeated.items.string.ipv4": "true",
                        "ves.io.schema.rules.repeated.max_items": "3",
                        "ves.io.schema.rules.repeated.unique": "true"
                    }
                },
                "inside_vip_cname": {
                    "type": "string",
                    "description": " CNAME value for the inside VIP,\n These are usually public cloud generated CNAME\n\nExample: - \"test.56670-387196482.useast2.ves.io\"-\n\nValidation Rules:\n  ves.io.schema.rules.string.max_len: 256\n",
                    "title": "Inside VIP CNAME",
                    "maxLength": 256,
                    "x-displayname": "Inside VIP CNAME",
                    "x-ves-example": "test.56670-387196482.useast2.ves.io",
                    "x-ves-validation-rules": {
                        "ves.io.schema.rules.string.max_len": "256"
                    }
                },
                "outside_vip": {
                    "type": "array",
                    "description": " List of Outside VIPs for an AZ\n\nExample: - \"192.168.0.156\"-\n\nRequired: YES\n\nValidation Rules:\n  ves.io.schema.rules.message.required: true\n  ves.io.schema.rules.repeated.items.string.ipv4: true\n  ves.io.schema.rules.repeated.max_items: 3\n  ves.io.schema.rules.repeated.min_items: 1\n  ves.io.schema.rules.repeated.unique: true\n",
                    "title": "Outside VIP(s)",
                    "minItems": 1,
                    "maxItems": 3,
                    "items": {
                        "type": "string"
                    },
                    "x-displayname": "Outside VIP(s)",
                    "x-ves-example": "192.168.0.156",
                    "x-ves-required": "true",
                    "x-ves-validation-rules": {
                        "ves.io.schema.rules.message.required": "true",
                        "ves.io.schema.rules.repeated.items.string.ipv4": "true",
                        "ves.io.schema.rules.repeated.max_items": "3",
                        "ves.io.schema.rules.repeated.min_items": "1",
                        "ves.io.schema.rules.repeated.unique": "true"
                    }
                },
                "outside_vip_cname": {
                    "type": "string",
                    "description": " CNAME value for the outside VIP\n These are usually public cloud generated CNAME\n\nExample: - \"test.56670-387196482.useast2.ves.io\"-\n\nValidation Rules:\n  ves.io.schema.rules.string.max_len: 256\n",
                    "title": "Outside VIP CNAME",
                    "maxLength": 256,
                    "x-displayname": "Outside VIP CNAME",
                    "x-ves-example": "test.56670-387196482.useast2.ves.io",
                    "x-ves-validation-rules": {
                        "ves.io.schema.rules.string.max_len": "256"
                    }
                }
            }
        }
    },
    "x-displayname": "Configure Azure Vnet Site",
    "x-ves-proto-file": "ves.io/schema/views/azure_vnet_site/public_customapi.proto"
}`
