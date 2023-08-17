// Copyright (c) 2022 F5, Inc. All rights reserved.
// Code generated by ves-gen-schema-go. DO NOT EDIT.

package forward_proxy_policy

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

// Create CustomDataAPI GRPC Client satisfying server.CustomClient
type CustomDataAPIGrpcClient struct {
	conn       *grpc.ClientConn
	grpcClient CustomDataAPIClient
	// map of rpc name to its invocation
	rpcFns map[string]func(context.Context, string, ...grpc.CallOption) (proto.Message, error)
}

func (c *CustomDataAPIGrpcClient) doRPCForwardProxyPolicyHits(ctx context.Context, yamlReq string, opts ...grpc.CallOption) (proto.Message, error) {
	req := &ForwardProxyPolicyHitsRequest{}
	if err := codec.FromYAML(yamlReq, req); err != nil {
		return nil, fmt.Errorf("YAML Request %s is not of type *ves.io.schema.views.forward_proxy_policy.ForwardProxyPolicyHitsRequest", yamlReq)
	}
	rsp, err := c.grpcClient.ForwardProxyPolicyHits(ctx, req, opts...)
	return rsp, err
}

func (c *CustomDataAPIGrpcClient) DoRPC(ctx context.Context, rpc string, opts ...server.CustomCallOpt) (proto.Message, error) {
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

func NewCustomDataAPIGrpcClient(cc *grpc.ClientConn) server.CustomClient {
	ccl := &CustomDataAPIGrpcClient{
		conn:       cc,
		grpcClient: NewCustomDataAPIClient(cc),
	}
	rpcFns := make(map[string]func(context.Context, string, ...grpc.CallOption) (proto.Message, error))
	rpcFns["ForwardProxyPolicyHits"] = ccl.doRPCForwardProxyPolicyHits

	ccl.rpcFns = rpcFns

	return ccl
}

// Create CustomDataAPI REST Client satisfying server.CustomClient
type CustomDataAPIRestClient struct {
	baseURL string
	client  http.Client
	// map of rpc name to its invocation
	rpcFns map[string]func(context.Context, *server.CustomCallOpts) (proto.Message, error)
}

func (c *CustomDataAPIRestClient) doRPCForwardProxyPolicyHits(ctx context.Context, callOpts *server.CustomCallOpts) (proto.Message, error) {
	if callOpts.URI == "" {
		return nil, fmt.Errorf("Error, URI should be specified, got empty")
	}
	url := fmt.Sprintf("%s%s", c.baseURL, callOpts.URI)

	yamlReq := callOpts.YAMLReq
	req := &ForwardProxyPolicyHitsRequest{}
	if err := codec.FromYAML(yamlReq, req); err != nil {
		return nil, fmt.Errorf("YAML Request %s is not of type *ves.io.schema.views.forward_proxy_policy.ForwardProxyPolicyHitsRequest: %s", yamlReq, err)
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
		q.Add("end_time", fmt.Sprintf("%v", req.EndTime))
		for _, item := range req.GroupBy {
			q.Add("group_by", fmt.Sprintf("%v", item))
		}
		for _, item := range req.LabelFilter {
			q.Add("label_filter", fmt.Sprintf("%v", item))
		}
		q.Add("namespace", fmt.Sprintf("%v", req.Namespace))
		q.Add("start_time", fmt.Sprintf("%v", req.StartTime))
		q.Add("step", fmt.Sprintf("%v", req.Step))

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
	pbRsp := &ForwardProxyPolicyHitsResponse{}
	if err := codec.FromJSON(string(body), pbRsp); err != nil {
		return nil, errors.Wrapf(err, "JSON Response %s is not of type *ves.io.schema.views.forward_proxy_policy.ForwardProxyPolicyHitsResponse", body)

	}
	if callOpts.OutCallResponse != nil {
		callOpts.OutCallResponse.ProtoMsg = pbRsp
		callOpts.OutCallResponse.JSON = string(body)
	}
	return pbRsp, nil
}

func (c *CustomDataAPIRestClient) DoRPC(ctx context.Context, rpc string, opts ...server.CustomCallOpt) (proto.Message, error) {
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

func NewCustomDataAPIRestClient(baseURL string, hc http.Client) server.CustomClient {
	ccl := &CustomDataAPIRestClient{
		baseURL: baseURL,
		client:  hc,
	}

	rpcFns := make(map[string]func(context.Context, *server.CustomCallOpts) (proto.Message, error))
	rpcFns["ForwardProxyPolicyHits"] = ccl.doRPCForwardProxyPolicyHits

	ccl.rpcFns = rpcFns

	return ccl
}

// Create customDataAPIInprocClient

// INPROC Client (satisfying CustomDataAPIClient interface)
type customDataAPIInprocClient struct {
	CustomDataAPIServer
}

func (c *customDataAPIInprocClient) ForwardProxyPolicyHits(ctx context.Context, in *ForwardProxyPolicyHitsRequest, opts ...grpc.CallOption) (*ForwardProxyPolicyHitsResponse, error) {
	return c.CustomDataAPIServer.ForwardProxyPolicyHits(ctx, in)
}

func NewCustomDataAPIInprocClient(svc svcfw.Service) CustomDataAPIClient {
	return &customDataAPIInprocClient{CustomDataAPIServer: NewCustomDataAPIServer(svc)}
}

// RegisterGwCustomDataAPIHandler registers with grpc-gw with an inproc-client backing so that
// rest to grpc happens without a grpc.Dial (thus avoiding additional certs for mTLS)
func RegisterGwCustomDataAPIHandler(ctx context.Context, mux *runtime.ServeMux, svc interface{}) error {
	s, ok := svc.(svcfw.Service)
	if !ok {
		return fmt.Errorf("svc is not svcfw.Service")
	}
	return RegisterCustomDataAPIHandlerClient(ctx, mux, NewCustomDataAPIInprocClient(s))
}

// Create customDataAPISrv

// SERVER (satisfying CustomDataAPIServer interface)
type customDataAPISrv struct {
	svc svcfw.Service
}

func (s *customDataAPISrv) ForwardProxyPolicyHits(ctx context.Context, in *ForwardProxyPolicyHitsRequest) (*ForwardProxyPolicyHitsResponse, error) {
	ah := s.svc.GetAPIHandler("ves.io.schema.views.forward_proxy_policy.CustomDataAPI")
	cah, ok := ah.(CustomDataAPIServer)
	if !ok {
		return nil, fmt.Errorf("ah %v is not of type *CustomDataAPIServer", ah)
	}

	var (
		rsp *ForwardProxyPolicyHitsResponse
		err error
	)

	bodyFields := svcfw.GenAuditReqBodyFields(ctx, s.svc, "ves.io.schema.views.forward_proxy_policy.ForwardProxyPolicyHitsRequest", in)
	defer func() {
		if len(bodyFields) > 0 {
			server.ExtendAPIAudit(ctx, svcfw.PublicAPIBodyLog.Uid, bodyFields)
		}
		userMsg := "The 'CustomDataAPI.ForwardProxyPolicyHits' operation on 'forward_proxy_policy'"
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
		if rvFn := s.svc.GetRPCValidator("ves.io.schema.views.forward_proxy_policy.CustomDataAPI.ForwardProxyPolicyHits"); rvFn != nil {
			if verr := rvFn(ctx, in); verr != nil {
				err = server.MaybePublicRestError(ctx, errors.Wrapf(verr, "Validating Request"))
				return nil, server.GRPCStatusFromError(err).Err()
			}
		}
	}

	rsp, err = cah.ForwardProxyPolicyHits(ctx, in)
	if err != nil {
		return rsp, server.GRPCStatusFromError(server.MaybePublicRestError(ctx, err)).Err()
	}

	bodyFields = append(bodyFields, svcfw.GenAuditRspBodyFields(ctx, s.svc, "ves.io.schema.views.forward_proxy_policy.ForwardProxyPolicyHitsResponse", rsp)...)

	return rsp, nil
}

func NewCustomDataAPIServer(svc svcfw.Service) CustomDataAPIServer {
	return &customDataAPISrv{svc: svc}
}

var CustomDataAPISwaggerJSON string = `{
    "swagger": "2.0",
    "info": {
        "title": "Forward Proxy Policy",
        "description": "Monitoring APIs for Forward Proxy Policy",
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
        "/public/namespaces/{namespace}/forward_proxy_policy/hits": {
            "post": {
                "summary": "Forward Proxy Policy Hits",
                "description": "Get the counter for Forward Proxy Policy hits for a given namespace.",
                "operationId": "ves.io.schema.views.forward_proxy_policy.CustomDataAPI.ForwardProxyPolicyHits",
                "responses": {
                    "200": {
                        "description": "A successful response.",
                        "schema": {
                            "$ref": "#/definitions/forward_proxy_policyForwardProxyPolicyHitsResponse"
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
                        "description": "Namespace\n\nx-example: \"ns1\"\nNamespace is used to scope forward proxy policy hits for the given namespace.",
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
                            "$ref": "#/definitions/forward_proxy_policyForwardProxyPolicyHitsRequest"
                        }
                    }
                ],
                "tags": [
                    "CustomDataAPI"
                ],
                "externalDocs": {
                    "description": "Examples of this operation",
                    "url": "https://www.volterra.io/docs/reference/api-ref/ves-io-schema-views-forward_proxy_policy-customdataapi-forwardproxypolicyhits"
                },
                "x-ves-proto-rpc": "ves.io.schema.views.forward_proxy_policy.CustomDataAPI.ForwardProxyPolicyHits"
            },
            "x-displayname": "Forward Proxy Policy Monitoring APIs",
            "x-ves-proto-service": "ves.io.schema.views.forward_proxy_policy.CustomDataAPI",
            "x-ves-proto-service-type": "CUSTOM_PUBLIC"
        }
    },
    "definitions": {
        "forward_proxy_policyForwardProxyPolicyHits": {
            "type": "object",
            "description": "ForwardProxyPolicyHits contains the timeseries data of forward proxy policy hits",
            "title": "Forward Proxy Policy Hits",
            "x-displayname": "Forward Proxy Policy Hits",
            "x-ves-proto-message": "ves.io.schema.views.forward_proxy_policy.ForwardProxyPolicyHits",
            "properties": {
                "id": {
                    "description": " ID identifies the unique combination of group_by label values in the response",
                    "title": "ID",
                    "$ref": "#/definitions/forward_proxy_policyForwardProxyPolicyHitsId",
                    "x-displayname": "ID"
                },
                "metric": {
                    "type": "array",
                    "description": " x-unit: \"count\"\n List of metric values",
                    "title": "Metric",
                    "items": {
                        "$ref": "#/definitions/schemaMetricValue"
                    },
                    "x-displayname": "Metric"
                }
            }
        },
        "forward_proxy_policyForwardProxyPolicyHitsId": {
            "type": "object",
            "description": "ForwardProxyPolicyHitsId uniquely identifies an entry in the response to forward proxy policy hits request.\nForward proxy policy hits counter is aggregated based on the labels specified in the group_by field in the request.\nTherefore, only the feields that corresponds to the MetricLabel in the group_by field will have non-empty\nvalue in the response.",
            "title": "Forward Proxy Policy Hits ID",
            "x-displayname": "Forward Proxy Policy Hits ID",
            "x-ves-proto-message": "ves.io.schema.views.forward_proxy_policy.ForwardProxyPolicyHitsId",
            "properties": {
                "action": {
                    "type": "string",
                    "description": " Action associated with the policy rule\n\nExample: - \"allow\"-",
                    "title": "Action",
                    "x-displayname": "Action",
                    "x-ves-example": "allow"
                },
                "namespace": {
                    "type": "string",
                    "description": " Namespace in which the policy rule was hit\n\nExample: - \"ns1\"-",
                    "title": "Namespace",
                    "x-displayname": "Namespace",
                    "x-ves-example": "ns1"
                },
                "policy": {
                    "type": "string",
                    "description": " Policy name\n\nExample: - \"policy1\"-",
                    "title": "Policy",
                    "x-displayname": "Policy",
                    "x-ves-example": "policy1"
                },
                "policy_rule": {
                    "type": "string",
                    "description": " Policy Rule name\n\nExample: - \"rule1\"-",
                    "title": "Policy Rule",
                    "x-displayname": "Policy Rule",
                    "x-ves-example": "rule1"
                },
                "site": {
                    "type": "string",
                    "description": " Site name\n\nExample: - \"ce1\"-",
                    "title": "Site",
                    "x-displayname": "Site",
                    "x-ves-example": "ce1"
                },
                "virtual_host": {
                    "type": "string",
                    "description": " Virtual Host name\n\nExample: - \"productpage\"-",
                    "title": "Virtual Host",
                    "x-displayname": "Virtual Host",
                    "x-ves-example": "productpage"
                }
            }
        },
        "forward_proxy_policyForwardProxyPolicyHitsRequest": {
            "type": "object",
            "description": "Request to get the forward proxy policy hits counter.",
            "title": "Forward Proxy Policy Hits Request",
            "x-displayname": "Forward Proxy Policy Hits Request",
            "x-ves-proto-message": "ves.io.schema.views.forward_proxy_policy.ForwardProxyPolicyHitsRequest",
            "properties": {
                "end_time": {
                    "type": "string",
                    "description": " end time of metric collection from which data will be considered.\n Format: unix_timestamp|rfc 3339\n\n Optional: If not specified, then the end_time will be evaluated to start_time+10m\n           If start_time is not specified, then the end_time will be evaluated to \u003ccurrent time\u003e\n\nExample: - \"1570007981\"-\n\nValidation Rules:\n  ves.io.schema.rules.string.query_time: true\n",
                    "title": "End time",
                    "x-displayname": "End Time",
                    "x-ves-example": "1570007981",
                    "x-ves-validation-rules": {
                        "ves.io.schema.rules.string.query_time": "true"
                    }
                },
                "group_by": {
                    "type": "array",
                    "description": " Aggregate data by one of more labels specified in group_by.\n\n Optional: If not specified, then the rule hits are aggregated/grouped by POLICY.",
                    "title": "Group by",
                    "items": {
                        "$ref": "#/definitions/forward_proxy_policyForwardProxyPolicyMetricLabel"
                    },
                    "x-displayname": "Group By"
                },
                "label_filter": {
                    "type": "array",
                    "description": " List of label filter expressions of the form \"label\" Op \"value\".\n Response will only contain data that matches all the conditions specified in the label_filter.\n\n Optional: If not specified, then the metrics will be filtered only based on the namespace in the request.",
                    "title": "Label Filter",
                    "items": {
                        "$ref": "#/definitions/forward_proxy_policyForwardProxyPolicyMetricLabelFilter"
                    },
                    "x-displayname": "Label Filter"
                },
                "namespace": {
                    "type": "string",
                    "description": " Namespace is used to scope forward proxy policy hits for the given namespace.\n\nExample: - \"ns1\"-",
                    "title": "Namespace",
                    "x-displayname": "Namespace",
                    "x-ves-example": "ns1"
                },
                "start_time": {
                    "type": "string",
                    "description": " start time of metric collection from which data will be considered.\n Format: unix_timestamp|rfc 3339\n\n Optional: If not specified, then the start_time will be evaluated to end_time-10m\n           If end_time is not specified, then the start_time will be evaluated to \u003ccurrent time\u003e-10m\n\nExample: - \"1570007981\"-\n\nValidation Rules:\n  ves.io.schema.rules.string.query_time: true\n",
                    "title": "Start time",
                    "x-displayname": "Start Time",
                    "x-ves-example": "1570007981",
                    "x-ves-validation-rules": {
                        "ves.io.schema.rules.string.query_time": "true"
                    }
                },
                "step": {
                    "type": "string",
                    "description": " step is the resolution width, which determines the number of the data points [x-axis (time)] to be returned in the response.\n The timestamps in the response will be t1=start_time, t2=t1+step, ... tn=tn-1+step, where tn \u003c= end_time.\n Format: [0-9][smhd], where s - seconds, m - minutes, h - hours, d - days\n\n Optional: If not specified, then step size is evaluated to \u003cend_time - start_time\u003e\n\nExample: - \"5m\"-\n\nValidation Rules:\n  ves.io.schema.rules.string.query_step: true\n",
                    "title": "Step",
                    "x-displayname": "Step",
                    "x-ves-example": "5m",
                    "x-ves-validation-rules": {
                        "ves.io.schema.rules.string.query_step": "true"
                    }
                }
            }
        },
        "forward_proxy_policyForwardProxyPolicyHitsResponse": {
            "type": "object",
            "description": "Number of forward proxy policy rule hits for each unique combination of group_by labels in the request.",
            "title": "Forward Proxy Policy Hits Response",
            "x-displayname": "Forward Proxy Policy Hits Response",
            "x-ves-proto-message": "ves.io.schema.views.forward_proxy_policy.ForwardProxyPolicyHitsResponse",
            "properties": {
                "data": {
                    "type": "array",
                    "description": " List of forward proxy policy hits data",
                    "title": "Forward Proxy Policy Hits",
                    "items": {
                        "$ref": "#/definitions/forward_proxy_policyForwardProxyPolicyHits"
                    },
                    "x-displayname": "Forward Proxy Policy Hits"
                },
                "step": {
                    "type": "string",
                    "description": " Actual step size used in the response. It could be higher than the requested step due to metric rollups and the query duration.\n Format: [0-9][smhd], where s - seconds, m - minutes, h - hours, d - days\n\nExample: - \"30m\"-\n\nValidation Rules:\n  ves.io.schema.rules.string.time_interval: true\n",
                    "title": "step",
                    "x-displayname": "Step",
                    "x-ves-example": "30m",
                    "x-ves-validation-rules": {
                        "ves.io.schema.rules.string.time_interval": "true"
                    }
                }
            }
        },
        "forward_proxy_policyForwardProxyPolicyMetricLabel": {
            "type": "string",
            "description": "Forward Proxy policy hits can be sliced and diced based on one or more labels listed below.\n",
            "title": "Forward Proxy Policy Metric Labels",
            "enum": [
                "NAMESPACE",
                "POLICY",
                "POLICY_RULE",
                "ACTION",
                "SITE",
                "VIRTUAL_HOST"
            ],
            "default": "NAMESPACE",
            "x-displayname": "Forward Proxy Policy Metric Labels",
            "x-ves-proto-enum": "ves.io.schema.views.forward_proxy_policy.ForwardProxyPolicyMetricLabel"
        },
        "forward_proxy_policyForwardProxyPolicyMetricLabelFilter": {
            "type": "object",
            "description": "Label filter can be specified to filter metrics based on label match",
            "title": "Forward Proxy Policy Metric Label Filter",
            "x-displayname": "Forward Proxy Policy Metric Label Filter",
            "x-ves-proto-message": "ves.io.schema.views.forward_proxy_policy.ForwardProxyPolicyMetricLabelFilter",
            "properties": {
                "label": {
                    "description": " Label associated with Forward proxy policy hits",
                    "title": "Label",
                    "$ref": "#/definitions/forward_proxy_policyForwardProxyPolicyMetricLabel",
                    "x-displayname": "Label"
                },
                "op": {
                    "description": " Operator to evaluate the label ",
                    "title": "Operator",
                    "$ref": "#/definitions/schemaMetricLabelOp",
                    "x-displayname": "Operator"
                },
                "value": {
                    "type": "string",
                    "description": " Value to be compared with\n\nExample: - \"policy1\"-",
                    "title": "Value",
                    "x-displayname": "Value",
                    "x-ves-example": "policy1"
                }
            }
        },
        "schemaMetricLabelOp": {
            "type": "string",
            "description": "The operator to use when filtering metrics based on label values.\n",
            "title": "Metric Label Operator",
            "enum": [
                "EQ",
                "NEQ"
            ],
            "default": "EQ",
            "x-displayname": "Metric Label Operator",
            "x-ves-proto-enum": "ves.io.schema.MetricLabelOp"
        },
        "schemaMetricValue": {
            "type": "object",
            "description": "Metric data contains timestamp and the value.",
            "title": "Metric Value",
            "x-displayname": "Metric Value",
            "x-ves-proto-message": "ves.io.schema.MetricValue",
            "properties": {
                "timestamp": {
                    "type": "number",
                    "description": " timestamp\n\nExample: - \"1570007981\"-",
                    "title": "Timestamp",
                    "format": "double",
                    "x-displayname": "Timestamp",
                    "x-ves-example": "1570007981"
                },
                "trend_value": {
                    "description": " trend value for the metric\n\nExample: - \"100.000000\"-",
                    "title": "Trend value",
                    "$ref": "#/definitions/schemaTrendValue",
                    "x-displayname": "Trend Value",
                    "x-ves-example": "100.000000"
                },
                "value": {
                    "type": "string",
                    "description": "\n\nExample: - \"15\"-",
                    "title": "Value",
                    "x-displayname": "Value",
                    "x-ves-example": "15"
                }
            }
        },
        "schemaTrendSentiment": {
            "type": "string",
            "description": "trend sentiment\n\nIndicates trend sentiment is positive\nIndicates trend sentiment is negative.",
            "title": "Trend Sentiment",
            "enum": [
                "TREND_SENTIMENT_NONE",
                "TREND_SENTIMENT_POSITIVE",
                "TREND_SENTIMENT_NEGATIVE"
            ],
            "default": "TREND_SENTIMENT_NONE",
            "x-displayname": "Trend Sentiment",
            "x-ves-proto-enum": "ves.io.schema.TrendSentiment"
        },
        "schemaTrendValue": {
            "type": "object",
            "description": "Trend value contains trend value, trend sentiment and trend calculation description and window size.",
            "title": "Trend Value",
            "x-displayname": "Trend Value",
            "x-ves-proto-message": "ves.io.schema.TrendValue",
            "properties": {
                "description": {
                    "type": "string",
                    "description": " description of the method used to calculate trend.\n\nExample: - \"Trend was calculated by comparing the avg of window size intervals of end-start Time and last window time interval\"-",
                    "title": "Description",
                    "x-displayname": "Description",
                    "x-ves-example": "Trend was calculated by comparing the avg of window size intervals of end-start Time and last window time interval"
                },
                "previous_value": {
                    "type": "string",
                    "description": "\n\nExample: - \"200.00\"-",
                    "title": "Previous Value",
                    "x-displayname": "Previous Value",
                    "x-ves-example": "200.00"
                },
                "sentiment": {
                    "description": "\n\nExample: - \"Positive\"-",
                    "title": "Sentiment",
                    "$ref": "#/definitions/schemaTrendSentiment",
                    "x-displayname": "Sentiment",
                    "x-ves-example": "Positive"
                },
                "value": {
                    "type": "string",
                    "description": "\n\nExample: - \"-15\"-",
                    "title": "Value",
                    "x-displayname": "Value",
                    "x-ves-example": "-15"
                }
            }
        }
    },
    "x-displayname": "Configure Forward Proxy Policy",
    "x-ves-proto-file": "ves.io/schema/views/forward_proxy_policy/public_custom_data_api.proto"
}`
