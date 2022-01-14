//
// Copyright (c) 2018 Volterra, Inc. All rights reserved.
// Code generated by ves-gen-schema-go. DO NOT EDIT.
//

package user

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

// Create CustomPrivateAPI GRPC Client satisfying server.CustomClient
type CustomPrivateAPIGrpcClient struct {
	conn       *grpc.ClientConn
	grpcClient CustomPrivateAPIClient
	// map of rpc name to its invocation
	rpcFns map[string]func(context.Context, string, ...grpc.CallOption) (proto.Message, error)
}

func (c *CustomPrivateAPIGrpcClient) doRPCCascadeDelete(ctx context.Context, yamlReq string, opts ...grpc.CallOption) (proto.Message, error) {
	req := &PrivateCascadeDeleteRequest{}
	if err := codec.FromYAML(yamlReq, req); err != nil {
		return nil, fmt.Errorf("YAML Request %s is not of type *ves.io.schema.user.PrivateCascadeDeleteRequest", yamlReq)
	}
	rsp, err := c.grpcClient.CascadeDelete(ctx, req, opts...)
	return rsp, err
}

func (c *CustomPrivateAPIGrpcClient) doRPCUpdateLastLogin(ctx context.Context, yamlReq string, opts ...grpc.CallOption) (proto.Message, error) {
	req := &LastLoginUpdateRequest{}
	if err := codec.FromYAML(yamlReq, req); err != nil {
		return nil, fmt.Errorf("YAML Request %s is not of type *ves.io.schema.user.LastLoginUpdateRequest", yamlReq)
	}
	rsp, err := c.grpcClient.UpdateLastLogin(ctx, req, opts...)
	return rsp, err
}

func (c *CustomPrivateAPIGrpcClient) DoRPC(ctx context.Context, rpc string, opts ...server.CustomCallOpt) (proto.Message, error) {
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

func NewCustomPrivateAPIGrpcClient(cc *grpc.ClientConn) server.CustomClient {
	ccl := &CustomPrivateAPIGrpcClient{
		conn:       cc,
		grpcClient: NewCustomPrivateAPIClient(cc),
	}
	rpcFns := make(map[string]func(context.Context, string, ...grpc.CallOption) (proto.Message, error))
	rpcFns["CascadeDelete"] = ccl.doRPCCascadeDelete

	rpcFns["UpdateLastLogin"] = ccl.doRPCUpdateLastLogin

	ccl.rpcFns = rpcFns

	return ccl
}

// Create CustomPrivateAPI REST Client satisfying server.CustomClient
type CustomPrivateAPIRestClient struct {
	baseURL string
	client  http.Client
	// map of rpc name to its invocation
	rpcFns map[string]func(context.Context, *server.CustomCallOpts) (proto.Message, error)
}

func (c *CustomPrivateAPIRestClient) doRPCCascadeDelete(ctx context.Context, callOpts *server.CustomCallOpts) (proto.Message, error) {
	if callOpts.URI == "" {
		return nil, fmt.Errorf("Error, URI should be specified, got empty")
	}
	url := fmt.Sprintf("%s%s", c.baseURL, callOpts.URI)

	yamlReq := callOpts.YAMLReq
	req := &PrivateCascadeDeleteRequest{}
	if err := codec.FromYAML(yamlReq, req); err != nil {
		return nil, fmt.Errorf("YAML Request %s is not of type *ves.io.schema.user.PrivateCascadeDeleteRequest: %s", yamlReq, err)
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
		q.Add("email", fmt.Sprintf("%v", req.Email))
		q.Add("namespace", fmt.Sprintf("%v", req.Namespace))
		q.Add("tenant_name", fmt.Sprintf("%v", req.TenantName))

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
	pbRsp := &CascadeDeleteResponse{}
	if err := codec.FromJSON(string(body), pbRsp); err != nil {
		return nil, fmt.Errorf("JSON Response %s is not of type *ves.io.schema.user.CascadeDeleteResponse", body)

	}
	if callOpts.OutCallResponse != nil {
		callOpts.OutCallResponse.ProtoMsg = pbRsp
		callOpts.OutCallResponse.JSON = string(body)
	}
	return pbRsp, nil
}

func (c *CustomPrivateAPIRestClient) doRPCUpdateLastLogin(ctx context.Context, callOpts *server.CustomCallOpts) (proto.Message, error) {
	if callOpts.URI == "" {
		return nil, fmt.Errorf("Error, URI should be specified, got empty")
	}
	url := fmt.Sprintf("%s%s", c.baseURL, callOpts.URI)

	yamlReq := callOpts.YAMLReq
	req := &LastLoginUpdateRequest{}
	if err := codec.FromYAML(yamlReq, req); err != nil {
		return nil, fmt.Errorf("YAML Request %s is not of type *ves.io.schema.user.LastLoginUpdateRequest: %s", yamlReq, err)
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
		q.Add("last_login_timestamp", fmt.Sprintf("%v", req.LastLoginTimestamp))
		q.Add("tenant", fmt.Sprintf("%v", req.Tenant))
		q.Add("user", fmt.Sprintf("%v", req.User))

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
	pbRsp := &LastLoginUpdateResponse{}
	if err := codec.FromJSON(string(body), pbRsp); err != nil {
		return nil, fmt.Errorf("JSON Response %s is not of type *ves.io.schema.user.LastLoginUpdateResponse", body)

	}
	if callOpts.OutCallResponse != nil {
		callOpts.OutCallResponse.ProtoMsg = pbRsp
		callOpts.OutCallResponse.JSON = string(body)
	}
	return pbRsp, nil
}

func (c *CustomPrivateAPIRestClient) DoRPC(ctx context.Context, rpc string, opts ...server.CustomCallOpt) (proto.Message, error) {
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

func NewCustomPrivateAPIRestClient(baseURL string, hc http.Client) server.CustomClient {
	ccl := &CustomPrivateAPIRestClient{
		baseURL: baseURL,
		client:  hc,
	}

	rpcFns := make(map[string]func(context.Context, *server.CustomCallOpts) (proto.Message, error))
	rpcFns["CascadeDelete"] = ccl.doRPCCascadeDelete

	rpcFns["UpdateLastLogin"] = ccl.doRPCUpdateLastLogin

	ccl.rpcFns = rpcFns

	return ccl
}

// Create CustomPrivateAPIInprocClient

// INPROC Client (satisfying CustomPrivateAPIClient interface)
type CustomPrivateAPIInprocClient struct {
	svc svcfw.Service
}

func (c *CustomPrivateAPIInprocClient) CascadeDelete(ctx context.Context, in *PrivateCascadeDeleteRequest, opts ...grpc.CallOption) (*CascadeDeleteResponse, error) {
	ah := c.svc.GetAPIHandler("ves.io.schema.user.CustomPrivateAPI")
	cah, ok := ah.(CustomPrivateAPIServer)
	if !ok {
		return nil, fmt.Errorf("ah %v is not of type *CustomPrivateAPISrv", ah)
	}

	var (
		rsp *CascadeDeleteResponse
		err error
	)

	if c.svc.Config().EnableAPIValidation {
		if rvFn := c.svc.GetRPCValidator("ves.io.schema.user.CustomPrivateAPI.CascadeDelete"); rvFn != nil {
			if verr := rvFn(ctx, in); verr != nil {
				err = server.MaybePublicRestError(ctx, errors.Wrapf(verr, "Validating Request"))
				return nil, server.GRPCStatusFromError(err).Err()
			}
		}
	}

	rsp, err = cah.CascadeDelete(ctx, in)
	if err != nil {
		return rsp, server.GRPCStatusFromError(server.MaybePublicRestError(ctx, err)).Err()
	}

	return rsp, nil
}
func (c *CustomPrivateAPIInprocClient) UpdateLastLogin(ctx context.Context, in *LastLoginUpdateRequest, opts ...grpc.CallOption) (*LastLoginUpdateResponse, error) {
	ah := c.svc.GetAPIHandler("ves.io.schema.user.CustomPrivateAPI")
	cah, ok := ah.(CustomPrivateAPIServer)
	if !ok {
		return nil, fmt.Errorf("ah %v is not of type *CustomPrivateAPISrv", ah)
	}

	var (
		rsp *LastLoginUpdateResponse
		err error
	)

	if c.svc.Config().EnableAPIValidation {
		if rvFn := c.svc.GetRPCValidator("ves.io.schema.user.CustomPrivateAPI.UpdateLastLogin"); rvFn != nil {
			if verr := rvFn(ctx, in); verr != nil {
				err = server.MaybePublicRestError(ctx, errors.Wrapf(verr, "Validating Request"))
				return nil, server.GRPCStatusFromError(err).Err()
			}
		}
	}

	rsp, err = cah.UpdateLastLogin(ctx, in)
	if err != nil {
		return rsp, server.GRPCStatusFromError(server.MaybePublicRestError(ctx, err)).Err()
	}

	return rsp, nil
}

func NewCustomPrivateAPIInprocClient(svc svcfw.Service) CustomPrivateAPIClient {
	return &CustomPrivateAPIInprocClient{svc: svc}
}

// RegisterGwCustomPrivateAPIHandler registers with grpc-gw with an inproc-client backing so that
// rest to grpc happens without a grpc.Dial (thus avoiding additional certs for mTLS)
func RegisterGwCustomPrivateAPIHandler(ctx context.Context, mux *runtime.ServeMux, svc interface{}) error {
	s, ok := svc.(svcfw.Service)
	if !ok {
		return fmt.Errorf("svc is not svcfw.Service")
	}
	return RegisterCustomPrivateAPIHandlerClient(ctx, mux, NewCustomPrivateAPIInprocClient(s))
}

var CustomPrivateAPISwaggerJSON string = `{
    "swagger": "2.0",
    "info": {
        "title": "User Custom Private API",
        "description": "Custom private APIs for user management",
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
        "/private/custom/namespaces/system/users/update_last_login": {
            "post": {
                "summary": "Update Last Login",
                "description": "API to update last login timestamp of user",
                "operationId": "ves.io.schema.user.CustomPrivateAPI.UpdateLastLogin",
                "responses": {
                    "200": {
                        "description": "A successful response.",
                        "schema": {
                            "$ref": "#/definitions/userLastLoginUpdateResponse"
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
                        "name": "body",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/userLastLoginUpdateRequest"
                        }
                    }
                ],
                "tags": [
                    "CustomPrivateAPI"
                ],
                "externalDocs": {
                    "description": "Examples of this operation",
                    "url": "https://www.volterra.io/docs/reference/api-ref/ves-io-schema-user-CustomPrivateAPI-UpdateLastLogin"
                },
                "x-ves-proto-rpc": "ves.io.schema.user.CustomPrivateAPI.UpdateLastLogin"
            },
            "x-displayname": "Custom Private API",
            "x-ves-proto-service": "ves.io.schema.user.CustomPrivateAPI",
            "x-ves-proto-service-type": "CUSTOM_PRIVATE"
        },
        "/ves.io.schema/introspect/write/namespaces/{namespace}/users/cascade_delete": {
            "post": {
                "summary": "CascadeDelete",
                "description": "CascadeDelete deletes the user and associated namespace roles for this user.\nUse this only if the user and its referenced objects need to be wiped out altogether.\nNote: users will always be in the system namespace.",
                "operationId": "ves.io.schema.user.CustomPrivateAPI.CascadeDelete",
                "responses": {
                    "200": {
                        "description": "A successful response.",
                        "schema": {
                            "$ref": "#/definitions/userCascadeDeleteResponse"
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
                        "description": "Namespace\n\nx-example: \"value\"\nValue of namespace is always \"system\"",
                        "in": "path",
                        "required": true,
                        "type": "string",
                        "x-displayname": "Namespace"
                    },
                    {
                        "name": "body",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/userPrivateCascadeDeleteRequest"
                        }
                    }
                ],
                "tags": [
                    "CustomPrivateAPI"
                ],
                "externalDocs": {
                    "description": "Examples of this operation",
                    "url": "https://www.volterra.io/docs/reference/api-ref/ves-io-schema-user-CustomPrivateAPI-CascadeDelete"
                },
                "x-ves-proto-rpc": "ves.io.schema.user.CustomPrivateAPI.CascadeDelete"
            },
            "x-displayname": "Custom Private API",
            "x-ves-proto-service": "ves.io.schema.user.CustomPrivateAPI",
            "x-ves-proto-service-type": "CUSTOM_PRIVATE"
        },
        "/ves.io.schema/introspect/write/private/custom/namespaces/system/users/update_last_login": {
            "post": {
                "summary": "Update Last Login",
                "description": "API to update last login timestamp of user",
                "operationId": "ves.io.schema.user.CustomPrivateAPI.UpdateLastLogin",
                "responses": {
                    "200": {
                        "description": "A successful response.",
                        "schema": {
                            "$ref": "#/definitions/userLastLoginUpdateResponse"
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
                        "name": "body",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/userLastLoginUpdateRequest"
                        }
                    }
                ],
                "tags": [
                    "CustomPrivateAPI"
                ],
                "externalDocs": {
                    "description": "Examples of this operation",
                    "url": "https://www.volterra.io/docs/reference/api-ref/ves-io-schema-user-CustomPrivateAPI-UpdateLastLogin"
                },
                "x-ves-proto-rpc": "ves.io.schema.user.CustomPrivateAPI.UpdateLastLogin"
            },
            "x-displayname": "Custom Private API",
            "x-ves-proto-service": "ves.io.schema.user.CustomPrivateAPI",
            "x-ves-proto-service-type": "CUSTOM_PRIVATE"
        }
    },
    "definitions": {
        "userCascadeDeleteItemType": {
            "type": "object",
            "description": "CascadeDeleteItemType contains details of object that was handled as part of cascade delete\nof user and whether it was successfully deleted",
            "title": "CascadeDeleteItemType",
            "x-displayname": "Cascade Deletion of User and Associated Namespace Roles",
            "x-ves-proto-message": "ves.io.schema.user.CascadeDeleteItemType",
            "properties": {
                "error_message": {
                    "type": "string",
                    "description": " informative error message about the success or failure of the object's deletion response\n\nExample: - \"value\"-",
                    "title": "error message",
                    "x-displayname": "Error Message",
                    "x-ves-example": "value"
                },
                "object_name": {
                    "type": "string",
                    "description": " Name of the object\n\nExample: - \"value\"-",
                    "title": "object's name",
                    "x-displayname": "Object Name",
                    "x-ves-example": "value"
                },
                "object_type": {
                    "type": "string",
                    "description": " The type of the object\n\nExample: - \"value\"-",
                    "title": "object's type",
                    "x-displayname": "Object Type",
                    "x-ves-example": "value"
                },
                "object_uid": {
                    "type": "string",
                    "description": " The uid of the object\n\nExample: - \"value\"-",
                    "title": "object's uid",
                    "x-displayname": "Object Uid",
                    "x-ves-example": "value"
                }
            }
        },
        "userCascadeDeleteResponse": {
            "type": "object",
            "description": "CascadeDeleteResponse contains a list of user objects that were deleted\nand possibly any errors when attempting to delete those objects.",
            "title": "CascadeDeleteResponse",
            "x-displayname": "Delete Response for the User and Associated Namespace Roles",
            "x-ves-proto-message": "ves.io.schema.user.CascadeDeleteResponse",
            "properties": {
                "delete_ok": {
                    "type": "boolean",
                    "description": " status of the deleted objects. \n \"true\" value indicates that the operation had been successful for all the objects.\n \"false\" value indicates that at least one of the delete operations had been unsuccessful.\n\nExample: - \"true\"-",
                    "title": "delete_ok",
                    "format": "boolean",
                    "x-displayname": "Delete Ok",
                    "x-ves-example": "true"
                },
                "items": {
                    "type": "array",
                    "description": " The objects deleted for the specific user",
                    "title": "items",
                    "items": {
                        "$ref": "#/definitions/userCascadeDeleteItemType"
                    },
                    "x-displayname": "Items"
                }
            }
        },
        "userLastLoginUpdateRequest": {
            "type": "object",
            "description": "Request to update user login timestamp.",
            "title": "LastLoginUpdateRequest",
            "x-displayname": "Last Login Update Request",
            "x-ves-proto-message": "ves.io.schema.user.LastLoginUpdateRequest",
            "properties": {
                "last_login_timestamp": {
                    "type": "string",
                    "description": " Last successful login timestamp of the user .",
                    "title": "last_login_timestamp",
                    "format": "date-time",
                    "x-displayname": "Last Login Timestamp"
                },
                "tenant": {
                    "type": "string",
                    "description": " Tenant ID of the tenant user belongs to.\n\nExample: - \"company1-as432s\"-",
                    "title": "tenant",
                    "x-displayname": "Tenant",
                    "x-ves-example": "company1-as432s"
                },
                "user": {
                    "type": "string",
                    "description": " User ID of the user. typically email id\n\nExample: - \"user@company1.com\"-",
                    "title": "user",
                    "x-displayname": "User",
                    "x-ves-example": "user@company1.com"
                }
            }
        },
        "userLastLoginUpdateResponse": {
            "type": "object",
            "title": "LastLoginUpdateResponse",
            "x-displayname": "Last Login Update Response",
            "x-ves-proto-message": "ves.io.schema.user.LastLoginUpdateResponse"
        },
        "userPrivateCascadeDeleteRequest": {
            "type": "object",
            "description": "PrivateCascadeDeleteRequest is the request to delete the user along with the associated namespace role objects.",
            "title": "PrivateCascadeDeleteRequest",
            "x-displayname": "Delete the User and Associated Namespace Roles",
            "x-ves-proto-message": "ves.io.schema.user.PrivateCascadeDeleteRequest",
            "properties": {
                "email": {
                    "type": "string",
                    "description": " email of the user requesting for\n\nExample: - \"value\"-\n\nRequired: YES\n\nValidation Rules:\n  ves.io.schema.rules.message.required: true\n",
                    "title": "email of the user",
                    "x-displayname": "Email",
                    "x-ves-example": "value",
                    "x-ves-required": "true",
                    "x-ves-validation-rules": {
                        "ves.io.schema.rules.message.required": "true"
                    }
                },
                "namespace": {
                    "type": "string",
                    "description": " Value of namespace is always \"system\"\n\nExample: - \"value\"-",
                    "title": "Namespace",
                    "x-displayname": "Namespace",
                    "x-ves-example": "value"
                },
                "tenant_name": {
                    "type": "string",
                    "description": " User deletion will be executed within this tenant.\n\nRequired: YES\n\nValidation Rules:\n  ves.io.schema.rules.message.required: true\n",
                    "title": "Tenant name",
                    "x-displayname": "Tenant name",
                    "x-ves-required": "true",
                    "x-ves-validation-rules": {
                        "ves.io.schema.rules.message.required": "true"
                    }
                }
            }
        }
    },
    "x-displayname": "User custom private API",
    "x-ves-proto-file": "ves.io/schema/user/private_customapi.proto"
}`
