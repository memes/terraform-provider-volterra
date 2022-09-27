//
// Copyright (c) 2022 F5, Inc. All rights reserved.
// Code generated by ves-gen-schema-go. DO NOT EDIT.
//

package managed_tenant

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

func (c *CustomAPIGrpcClient) doRPCGetManagedTenantList(ctx context.Context, yamlReq string, opts ...grpc.CallOption) (proto.Message, error) {
	req := &GetManagedTenantListReq{}
	if err := codec.FromYAML(yamlReq, req); err != nil {
		return nil, fmt.Errorf("YAML Request %s is not of type *ves.io.schema.tenant_management.managed_tenant.GetManagedTenantListReq", yamlReq)
	}
	rsp, err := c.grpcClient.GetManagedTenantList(ctx, req, opts...)
	return rsp, err
}

func (c *CustomAPIGrpcClient) doRPCGetManagedTenantListByUser(ctx context.Context, yamlReq string, opts ...grpc.CallOption) (proto.Message, error) {
	req := &GetManagedTenantListReq{}
	if err := codec.FromYAML(yamlReq, req); err != nil {
		return nil, fmt.Errorf("YAML Request %s is not of type *ves.io.schema.tenant_management.managed_tenant.GetManagedTenantListReq", yamlReq)
	}
	rsp, err := c.grpcClient.GetManagedTenantListByUser(ctx, req, opts...)
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
	rpcFns["GetManagedTenantList"] = ccl.doRPCGetManagedTenantList

	rpcFns["GetManagedTenantListByUser"] = ccl.doRPCGetManagedTenantListByUser

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

func (c *CustomAPIRestClient) doRPCGetManagedTenantList(ctx context.Context, callOpts *server.CustomCallOpts) (proto.Message, error) {
	if callOpts.URI == "" {
		return nil, fmt.Errorf("Error, URI should be specified, got empty")
	}
	url := fmt.Sprintf("%s%s", c.baseURL, callOpts.URI)

	yamlReq := callOpts.YAMLReq
	req := &GetManagedTenantListReq{}
	if err := codec.FromYAML(yamlReq, req); err != nil {
		return nil, fmt.Errorf("YAML Request %s is not of type *ves.io.schema.tenant_management.managed_tenant.GetManagedTenantListReq: %s", yamlReq, err)
	}

	var hReq *http.Request
	hm := strings.ToLower(callOpts.HTTPMethod)
	switch hm {
	case "post", "put":
		jsn, err := codec.ToJSON(req, codec.ToWithUseProtoFieldName())
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
		q.Add("page_limit", fmt.Sprintf("%v", req.PageLimit))
		q.Add("page_start", fmt.Sprintf("%v", req.PageStart))
		q.Add("search_keyword", fmt.Sprintf("%v", req.SearchKeyword))

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

	// checking whether the status code is a successful status code (2xx series)
	if rsp.StatusCode < 200 || rsp.StatusCode > 299 {
		body, err := ioutil.ReadAll(rsp.Body)
		return nil, fmt.Errorf("Unsuccessful custom API %s on %s, status code %d, body %s, err %s", callOpts.HTTPMethod, callOpts.URI, rsp.StatusCode, body, err)
	}

	body, err := ioutil.ReadAll(rsp.Body)
	if err != nil {
		return nil, errors.Wrap(err, "Custom API RestClient read body")
	}
	pbRsp := &GetManagedTenantListResp{}
	if err := codec.FromJSON(string(body), pbRsp); err != nil {
		return nil, errors.Wrapf(err, "JSON Response %s is not of type *ves.io.schema.tenant_management.managed_tenant.GetManagedTenantListResp", body)

	}
	if callOpts.OutCallResponse != nil {
		callOpts.OutCallResponse.ProtoMsg = pbRsp
		callOpts.OutCallResponse.JSON = string(body)
	}
	return pbRsp, nil
}

func (c *CustomAPIRestClient) doRPCGetManagedTenantListByUser(ctx context.Context, callOpts *server.CustomCallOpts) (proto.Message, error) {
	if callOpts.URI == "" {
		return nil, fmt.Errorf("Error, URI should be specified, got empty")
	}
	url := fmt.Sprintf("%s%s", c.baseURL, callOpts.URI)

	yamlReq := callOpts.YAMLReq
	req := &GetManagedTenantListReq{}
	if err := codec.FromYAML(yamlReq, req); err != nil {
		return nil, fmt.Errorf("YAML Request %s is not of type *ves.io.schema.tenant_management.managed_tenant.GetManagedTenantListReq: %s", yamlReq, err)
	}

	var hReq *http.Request
	hm := strings.ToLower(callOpts.HTTPMethod)
	switch hm {
	case "post", "put":
		jsn, err := codec.ToJSON(req, codec.ToWithUseProtoFieldName())
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
		q.Add("page_limit", fmt.Sprintf("%v", req.PageLimit))
		q.Add("page_start", fmt.Sprintf("%v", req.PageStart))
		q.Add("search_keyword", fmt.Sprintf("%v", req.SearchKeyword))

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

	// checking whether the status code is a successful status code (2xx series)
	if rsp.StatusCode < 200 || rsp.StatusCode > 299 {
		body, err := ioutil.ReadAll(rsp.Body)
		return nil, fmt.Errorf("Unsuccessful custom API %s on %s, status code %d, body %s, err %s", callOpts.HTTPMethod, callOpts.URI, rsp.StatusCode, body, err)
	}

	body, err := ioutil.ReadAll(rsp.Body)
	if err != nil {
		return nil, errors.Wrap(err, "Custom API RestClient read body")
	}
	pbRsp := &GetManagedTenantListResp{}
	if err := codec.FromJSON(string(body), pbRsp); err != nil {
		return nil, errors.Wrapf(err, "JSON Response %s is not of type *ves.io.schema.tenant_management.managed_tenant.GetManagedTenantListResp", body)

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
	rpcFns["GetManagedTenantList"] = ccl.doRPCGetManagedTenantList

	rpcFns["GetManagedTenantListByUser"] = ccl.doRPCGetManagedTenantListByUser

	ccl.rpcFns = rpcFns

	return ccl
}

// Create customAPIInprocClient

// INPROC Client (satisfying CustomAPIClient interface)
type customAPIInprocClient struct {
	CustomAPIServer
}

func (c *customAPIInprocClient) GetManagedTenantList(ctx context.Context, in *GetManagedTenantListReq, opts ...grpc.CallOption) (*GetManagedTenantListResp, error) {
	return c.CustomAPIServer.GetManagedTenantList(ctx, in)
}
func (c *customAPIInprocClient) GetManagedTenantListByUser(ctx context.Context, in *GetManagedTenantListReq, opts ...grpc.CallOption) (*GetManagedTenantListResp, error) {
	return c.CustomAPIServer.GetManagedTenantListByUser(ctx, in)
}

func NewCustomAPIInprocClient(svc svcfw.Service) CustomAPIClient {
	return &customAPIInprocClient{CustomAPIServer: NewCustomAPIServer(svc)}
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

// Create customAPISrv

// SERVER (satisfying CustomAPIServer interface)
type customAPISrv struct {
	svc svcfw.Service
}

func (s *customAPISrv) GetManagedTenantList(ctx context.Context, in *GetManagedTenantListReq) (*GetManagedTenantListResp, error) {
	ah := s.svc.GetAPIHandler("ves.io.schema.tenant_management.managed_tenant.CustomAPI")
	cah, ok := ah.(CustomAPIServer)
	if !ok {
		return nil, fmt.Errorf("ah %v is not of type *CustomAPIServer", ah)
	}

	var (
		rsp *GetManagedTenantListResp
		err error
	)

	bodyFields := svcfw.GenAuditReqBodyFields(ctx, s.svc, "ves.io.schema.tenant_management.managed_tenant.GetManagedTenantListReq", in)
	defer func() {
		if len(bodyFields) > 0 {
			server.ExtendAPIAudit(ctx, svcfw.PublicAPIBodyLog.Uid, bodyFields)
		}
		userMsg := "The 'CustomAPI.GetManagedTenantList' operation on 'managed_tenant'"
		if err == nil {
			userMsg += " was successfully performed."
		} else {
			userMsg += " failed to be performed."
		}
		server.AddUserMsgToAPIAudit(ctx, userMsg)
	}()

	if err := svcfw.FillOneofDefaultChoice(ctx, s.svc, in); err != nil {
		err = server.MaybePublicRestError(ctx, errors.Wrapf(err, "Filling oneof default choice"))
		return nil, server.GRPCStatusFromError(err).Err()
	}

	if s.svc.Config().EnableAPIValidation {
		if rvFn := s.svc.GetRPCValidator("ves.io.schema.tenant_management.managed_tenant.CustomAPI.GetManagedTenantList"); rvFn != nil {
			if verr := rvFn(ctx, in); verr != nil {
				err = server.MaybePublicRestError(ctx, errors.Wrapf(verr, "Validating Request"))
				return nil, server.GRPCStatusFromError(err).Err()
			}
		}
	}

	rsp, err = cah.GetManagedTenantList(ctx, in)
	if err != nil {
		return rsp, server.GRPCStatusFromError(server.MaybePublicRestError(ctx, err)).Err()
	}

	bodyFields = append(bodyFields, svcfw.GenAuditRspBodyFields(ctx, s.svc, "ves.io.schema.tenant_management.managed_tenant.GetManagedTenantListResp", rsp)...)

	return rsp, nil
}
func (s *customAPISrv) GetManagedTenantListByUser(ctx context.Context, in *GetManagedTenantListReq) (*GetManagedTenantListResp, error) {
	ah := s.svc.GetAPIHandler("ves.io.schema.tenant_management.managed_tenant.CustomAPI")
	cah, ok := ah.(CustomAPIServer)
	if !ok {
		return nil, fmt.Errorf("ah %v is not of type *CustomAPIServer", ah)
	}

	var (
		rsp *GetManagedTenantListResp
		err error
	)

	bodyFields := svcfw.GenAuditReqBodyFields(ctx, s.svc, "ves.io.schema.tenant_management.managed_tenant.GetManagedTenantListReq", in)
	defer func() {
		if len(bodyFields) > 0 {
			server.ExtendAPIAudit(ctx, svcfw.PublicAPIBodyLog.Uid, bodyFields)
		}
		userMsg := "The 'CustomAPI.GetManagedTenantListByUser' operation on 'managed_tenant'"
		if err == nil {
			userMsg += " was successfully performed."
		} else {
			userMsg += " failed to be performed."
		}
		server.AddUserMsgToAPIAudit(ctx, userMsg)
	}()

	if err := svcfw.FillOneofDefaultChoice(ctx, s.svc, in); err != nil {
		err = server.MaybePublicRestError(ctx, errors.Wrapf(err, "Filling oneof default choice"))
		return nil, server.GRPCStatusFromError(err).Err()
	}

	if s.svc.Config().EnableAPIValidation {
		if rvFn := s.svc.GetRPCValidator("ves.io.schema.tenant_management.managed_tenant.CustomAPI.GetManagedTenantListByUser"); rvFn != nil {
			if verr := rvFn(ctx, in); verr != nil {
				err = server.MaybePublicRestError(ctx, errors.Wrapf(verr, "Validating Request"))
				return nil, server.GRPCStatusFromError(err).Err()
			}
		}
	}

	rsp, err = cah.GetManagedTenantListByUser(ctx, in)
	if err != nil {
		return rsp, server.GRPCStatusFromError(server.MaybePublicRestError(ctx, err)).Err()
	}

	bodyFields = append(bodyFields, svcfw.GenAuditRspBodyFields(ctx, s.svc, "ves.io.schema.tenant_management.managed_tenant.GetManagedTenantListResp", rsp)...)

	return rsp, nil
}

func NewCustomAPIServer(svc svcfw.Service) CustomAPIServer {
	return &customAPISrv{svc: svc}
}

var CustomAPISwaggerJSON string = `{
    "swagger": "2.0",
    "info": {
        "title": "Managed Tenant",
        "description": "Additional public APIs for managed_tenant config object.",
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
        "/public/namespaces/system/managed_tenants_by_user": {
            "get": {
                "summary": "Get List of Managed Tenant By User",
                "description": "Get list of managed tenants that user have access to based on assingned membership.\nThis is an optimized list generated based on the requesting user's current group assignments\nthat will allow access to managed tenant.",
                "operationId": "ves.io.schema.tenant_management.managed_tenant.CustomAPI.GetManagedTenantListByUser",
                "responses": {
                    "200": {
                        "description": "A successful response.",
                        "schema": {
                            "$ref": "#/definitions/managed_tenantGetManagedTenantListResp"
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
                        "name": "search_keyword",
                        "description": "x-example: \"value\"\nSearch Keyword for filtering the Managed Tenant List.",
                        "in": "query",
                        "required": false,
                        "type": "string",
                        "x-displayname": "SearchKeyword"
                    },
                    {
                        "name": "page_start",
                        "description": "x-example: \"c5776a8e-bcae-4392-98d3-3556f4b9df1b\"\nPageStart will hold the UUID of the first item in the requested page.\nResponse will contain items upto count specified in PageLimit starting from  PageStart.\nIf this is empty then first page will be served.",
                        "in": "query",
                        "required": false,
                        "type": "string",
                        "x-displayname": "PageStart"
                    },
                    {
                        "name": "page_limit",
                        "description": "x-example: \"100\"\nPageLimit will hold the limit of items required per query.\nDefault value is set as 100",
                        "in": "query",
                        "required": false,
                        "type": "integer",
                        "format": "int32",
                        "x-displayname": "PageLimit"
                    }
                ],
                "tags": [
                    "CustomAPI"
                ],
                "externalDocs": {
                    "description": "Examples of this operation",
                    "url": "https://www.volterra.io/docs/reference/api-ref/ves-io-schema-tenant_management-managed_tenant-customapi-getmanagedtenantlistbyuser"
                },
                "x-ves-proto-rpc": "ves.io.schema.tenant_management.managed_tenant.CustomAPI.GetManagedTenantListByUser"
            },
            "x-displayname": "Managed Tenant - UAM Manager Custom APIs",
            "x-ves-proto-service": "ves.io.schema.tenant_management.managed_tenant.CustomAPI",
            "x-ves-proto-service-type": "CUSTOM_PUBLIC"
        },
        "/public/namespaces/system/managed_tenants_list": {
            "get": {
                "summary": "Get List of Managed Tenant",
                "description": "Get full list of managed tenants access details.\nThis response will contain full list of managed tenant based on the configuration\nand is not filtered by requesting user's group membership that enable access.",
                "operationId": "ves.io.schema.tenant_management.managed_tenant.CustomAPI.GetManagedTenantList",
                "responses": {
                    "200": {
                        "description": "A successful response.",
                        "schema": {
                            "$ref": "#/definitions/managed_tenantGetManagedTenantListResp"
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
                        "name": "search_keyword",
                        "description": "x-example: \"value\"\nSearch Keyword for filtering the Managed Tenant List.",
                        "in": "query",
                        "required": false,
                        "type": "string",
                        "x-displayname": "SearchKeyword"
                    },
                    {
                        "name": "page_start",
                        "description": "x-example: \"c5776a8e-bcae-4392-98d3-3556f4b9df1b\"\nPageStart will hold the UUID of the first item in the requested page.\nResponse will contain items upto count specified in PageLimit starting from  PageStart.\nIf this is empty then first page will be served.",
                        "in": "query",
                        "required": false,
                        "type": "string",
                        "x-displayname": "PageStart"
                    },
                    {
                        "name": "page_limit",
                        "description": "x-example: \"100\"\nPageLimit will hold the limit of items required per query.\nDefault value is set as 100",
                        "in": "query",
                        "required": false,
                        "type": "integer",
                        "format": "int32",
                        "x-displayname": "PageLimit"
                    }
                ],
                "tags": [
                    "CustomAPI"
                ],
                "externalDocs": {
                    "description": "Examples of this operation",
                    "url": "https://www.volterra.io/docs/reference/api-ref/ves-io-schema-tenant_management-managed_tenant-customapi-getmanagedtenantlist"
                },
                "x-ves-proto-rpc": "ves.io.schema.tenant_management.managed_tenant.CustomAPI.GetManagedTenantList"
            },
            "x-displayname": "Managed Tenant - UAM Manager Custom APIs",
            "x-ves-proto-service": "ves.io.schema.tenant_management.managed_tenant.CustomAPI",
            "x-ves-proto-service-type": "CUSTOM_PUBLIC"
        }
    },
    "definitions": {
        "managed_tenantAccessInfo": {
            "type": "object",
            "description": "Access details of a managed tenant.",
            "title": "AccessInfo",
            "x-displayname": "Access Info",
            "x-ves-proto-message": "ves.io.schema.tenant_management.managed_tenant.AccessInfo",
            "properties": {
                "groups": {
                    "type": "array",
                    "description": " List of local user group association to user groups in the managed tenant.",
                    "title": "groups",
                    "items": {
                        "$ref": "#/definitions/managed_tenantGroupAssignmentType"
                    },
                    "x-displayname": "Groups"
                },
                "link": {
                    "description": " Info about hyperlink to access the managed tenant.",
                    "title": "Link",
                    "$ref": "#/definitions/viewsLinkRefType",
                    "x-displayname": "Link"
                },
                "name": {
                    "type": "string",
                    "description": " Name of the managed tenant config object. it can be used as known identifier.\n\nExample: - \"l1-support\"-",
                    "title": "Name",
                    "x-displayname": "Name",
                    "x-ves-example": "l1-support"
                }
            }
        },
        "managed_tenantGetManagedTenantListResp": {
            "type": "object",
            "description": "Response to get list of managed tenant access.",
            "title": "GetManagedTenantListResp",
            "x-displayname": "Get Managed Tenant Access Response",
            "x-ves-proto-message": "ves.io.schema.tenant_management.managed_tenant.GetManagedTenantListResp",
            "properties": {
                "access_config": {
                    "type": "array",
                    "description": " Allowed access configuration details for the tenant.",
                    "title": "access_config",
                    "items": {
                        "$ref": "#/definitions/managed_tenantAccessInfo"
                    },
                    "x-displayname": "Access Config"
                },
                "next_page": {
                    "type": "string",
                    "description": " NextPage contains the UUID of the first item of the next page.\n This value can be passed back as the PageStart in the next request.\n If this empty means current one is the last page.",
                    "title": "next_page",
                    "x-displayname": "NextPage"
                },
                "total_access_config_count": {
                    "type": "integer",
                    "description": " TotalAccessConfigCount holds total count of access config.",
                    "title": "total_access_config_count",
                    "format": "int32",
                    "x-displayname": "TotalAccessConfigCount"
                }
            }
        },
        "managed_tenantGroupAssignmentType": {
            "type": "object",
            "description": "Shape for specifying user group assosciation to user groups in a managed tenant.",
            "title": "GroupAssignmentType",
            "x-displayname": "Group to Assign",
            "x-ves-proto-message": "ves.io.schema.tenant_management.managed_tenant.GroupAssignmentType",
            "properties": {
                "group": {
                    "description": " Assosciate existing local user group which will be used to map groups in managed tenant.\n User should be member of this group to gain access into managed tenant.",
                    "title": "group",
                    "$ref": "#/definitions/schemaviewsObjectRefType",
                    "x-displayname": "Group"
                },
                "managed_tenant_groups": {
                    "type": "array",
                    "description": " List of group names in managed tenant (MT).\n Note - To properly establish access, admin of managed tenant need to create corresponding Allowed Tenant\n configuration object with access to use same group names. Once it's setup, when user from original tenant\n access managed tenant, underlying roles from managed tenant will be applied to user.\n\nExample: - \"user-group1\"-\n\nValidation Rules:\n  ves.io.schema.rules.repeated.max_items: 32\n  ves.io.schema.rules.repeated.unique: true\n",
                    "title": "managed_tenant_groups",
                    "maxItems": 32,
                    "items": {
                        "type": "string"
                    },
                    "x-displayname": "Managed Tenant Groups",
                    "x-ves-example": "user-group1",
                    "x-ves-validation-rules": {
                        "ves.io.schema.rules.repeated.max_items": "32",
                        "ves.io.schema.rules.repeated.unique": "true"
                    }
                }
            }
        },
        "schemaviewsObjectRefType": {
            "type": "object",
            "description": "This type establishes a direct reference from one object(the referrer) to another(the referred).\nSuch a reference is in form of tenant/namespace/name",
            "title": "ObjectRefType",
            "x-displayname": "Object reference",
            "x-ves-proto-message": "ves.io.schema.views.ObjectRefType",
            "properties": {
                "name": {
                    "type": "string",
                    "description": " When a configuration object(e.g. virtual_host) refers to another(e.g route)\n then name will hold the referred object's(e.g. route's) name.\n\nExample: - \"contacts-route\"-\n\nRequired: YES\n\nValidation Rules:\n  ves.io.schema.rules.message.required: true\n  ves.io.schema.rules.string.max_bytes: 64\n  ves.io.schema.rules.string.min_bytes: 1\n",
                    "title": "name",
                    "minLength": 1,
                    "maxLength": 64,
                    "x-displayname": "Name",
                    "x-ves-example": "contacts-route",
                    "x-ves-required": "true",
                    "x-ves-validation-rules": {
                        "ves.io.schema.rules.message.required": "true",
                        "ves.io.schema.rules.string.max_bytes": "64",
                        "ves.io.schema.rules.string.min_bytes": "1"
                    }
                },
                "namespace": {
                    "type": "string",
                    "description": " When a configuration object(e.g. virtual_host) refers to another(e.g route)\n then namespace will hold the referred object's(e.g. route's) namespace.\n\nExample: - \"ns1\"-\n\nValidation Rules:\n  ves.io.schema.rules.string.max_bytes: 64\n",
                    "title": "namespace",
                    "maxLength": 64,
                    "x-displayname": "Namespace",
                    "x-ves-example": "ns1",
                    "x-ves-validation-rules": {
                        "ves.io.schema.rules.string.max_bytes": "64"
                    }
                },
                "tenant": {
                    "type": "string",
                    "description": " When a configuration object(e.g. virtual_host) refers to another(e.g route)\n then tenant will hold the referred object's(e.g. route's) tenant.\n\nExample: - \"acmecorp\"-\n\nValidation Rules:\n  ves.io.schema.rules.string.max_bytes: 64\n",
                    "title": "tenant",
                    "maxLength": 64,
                    "x-displayname": "Tenant",
                    "x-ves-example": "acmecorp",
                    "x-ves-validation-rules": {
                        "ves.io.schema.rules.string.max_bytes": "64"
                    }
                }
            }
        },
        "viewsLinkRefType": {
            "type": "object",
            "description": "LinkRefType\nThis message defines a reference to hyperlink that can be accessed via web.",
            "x-ves-proto-message": "ves.io.schema.views.LinkRefType",
            "properties": {
                "href": {
                    "type": "string",
                    "description": " Referred link's location. This can be treated as equivalent of href in html.\n\nExample: - \"https://f5.com/link/resource_a\"-\n\nValidation Rules:\n  ves.io.schema.rules.string.max_len: 512\n  ves.io.schema.rules.string.min_len: 1\n",
                    "title": "HyperlinkRef",
                    "minLength": 1,
                    "maxLength": 512,
                    "x-displayname": "Hyperlink reference",
                    "x-ves-example": "https://f5.com/link/resource_a",
                    "x-ves-validation-rules": {
                        "ves.io.schema.rules.string.max_len": "512",
                        "ves.io.schema.rules.string.min_len": "1"
                    }
                },
                "name": {
                    "type": "string",
                    "description": " Name to use for displaying above link in href\n\nExample: - \"Resource A\"-\n\nValidation Rules:\n  ves.io.schema.rules.string.max_len: 256\n  ves.io.schema.rules.string.min_len: 1\n",
                    "title": "Name",
                    "minLength": 1,
                    "maxLength": 256,
                    "x-displayname": "Link Name",
                    "x-ves-example": "Resource A",
                    "x-ves-validation-rules": {
                        "ves.io.schema.rules.string.max_len": "256",
                        "ves.io.schema.rules.string.min_len": "1"
                    }
                }
            }
        }
    },
    "x-displayname": "Managed Tenant",
    "x-ves-proto-file": "ves.io/schema/tenant_management/managed_tenant/public_customapi.proto"
}`
