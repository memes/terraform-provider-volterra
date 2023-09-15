// Copyright (c) 2022 F5, Inc. All rights reserved.
// Code generated by ves-gen-schema-go. DO NOT EDIT.

package nfv_service

import (
	"bytes"
	"context"
	"fmt"
	io "io"
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

func (c *CustomDataAPIGrpcClient) doRPCMetrics(ctx context.Context, yamlReq string, opts ...grpc.CallOption) (proto.Message, error) {
	req := &MetricsRequest{}
	if err := codec.FromYAML(yamlReq, req); err != nil {
		return nil, fmt.Errorf("YAML Request %s is not of type *ves.io.schema.nfv_service.MetricsRequest", yamlReq)
	}
	rsp, err := c.grpcClient.Metrics(ctx, req, opts...)
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
	rpcFns["Metrics"] = ccl.doRPCMetrics

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

func (c *CustomDataAPIRestClient) doRPCMetrics(ctx context.Context, callOpts *server.CustomCallOpts) (proto.Message, error) {
	if callOpts.URI == "" {
		return nil, fmt.Errorf("Error, URI should be specified, got empty")
	}
	url := fmt.Sprintf("%s%s", c.baseURL, callOpts.URI)

	yamlReq := callOpts.YAMLReq
	req := &MetricsRequest{}
	if err := codec.FromYAML(yamlReq, req); err != nil {
		return nil, fmt.Errorf("YAML Request %s is not of type *ves.io.schema.nfv_service.MetricsRequest: %s", yamlReq, err)
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
		for _, item := range req.FieldSelector {
			q.Add("field_selector", fmt.Sprintf("%v", item))
		}
		q.Add("filter", fmt.Sprintf("%v", req.Filter))
		for _, item := range req.GroupBy {
			q.Add("group_by", fmt.Sprintf("%v", item))
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
		body, err := io.ReadAll(rsp.Body)
		return nil, fmt.Errorf("Unsuccessful custom API %s on %s, status code %d, body %s, err %s", callOpts.HTTPMethod, callOpts.URI, rsp.StatusCode, body, err)
	}

	body, err := io.ReadAll(rsp.Body)
	if err != nil {
		return nil, errors.Wrap(err, "Custom API RestClient read body")
	}
	pbRsp := &MetricsResponse{}
	if err := codec.FromJSON(string(body), pbRsp); err != nil {
		return nil, errors.Wrapf(err, "JSON Response %s is not of type *ves.io.schema.nfv_service.MetricsResponse", body)

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
	rpcFns["Metrics"] = ccl.doRPCMetrics

	ccl.rpcFns = rpcFns

	return ccl
}

// Create customDataAPIInprocClient

// INPROC Client (satisfying CustomDataAPIClient interface)
type customDataAPIInprocClient struct {
	CustomDataAPIServer
}

func (c *customDataAPIInprocClient) Metrics(ctx context.Context, in *MetricsRequest, opts ...grpc.CallOption) (*MetricsResponse, error) {
	return c.CustomDataAPIServer.Metrics(ctx, in)
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

func (s *customDataAPISrv) Metrics(ctx context.Context, in *MetricsRequest) (*MetricsResponse, error) {
	ah := s.svc.GetAPIHandler("ves.io.schema.nfv_service.CustomDataAPI")
	cah, ok := ah.(CustomDataAPIServer)
	if !ok {
		return nil, fmt.Errorf("ah %v is not of type *CustomDataAPIServer", ah)
	}

	var (
		rsp *MetricsResponse
		err error
	)

	bodyFields := svcfw.GenAuditReqBodyFields(ctx, s.svc, "ves.io.schema.nfv_service.MetricsRequest", in)
	defer func() {
		if len(bodyFields) > 0 {
			server.ExtendAPIAudit(ctx, svcfw.PublicAPIBodyLog.Uid, bodyFields)
		}
		userMsg := "The 'CustomDataAPI.Metrics' operation on 'nfv_service'"
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
		if rvFn := s.svc.GetRPCValidator("ves.io.schema.nfv_service.CustomDataAPI.Metrics"); rvFn != nil {
			if verr := rvFn(ctx, in); verr != nil {
				err = server.MaybePublicRestError(ctx, errors.Wrapf(verr, "Validating Request"))
				return nil, server.GRPCStatusFromError(err).Err()
			}
		}
	}

	rsp, err = cah.Metrics(ctx, in)
	if err != nil {
		return rsp, server.GRPCStatusFromError(server.MaybePublicRestError(ctx, err)).Err()
	}

	bodyFields = append(bodyFields, svcfw.GenAuditRspBodyFields(ctx, s.svc, "ves.io.schema.nfv_service.MetricsResponse", rsp)...)

	return rsp, nil
}

func NewCustomDataAPIServer(svc svcfw.Service) CustomDataAPIServer {
	return &customDataAPISrv{svc: svc}
}

var CustomDataAPISwaggerJSON string = `{
    "swagger": "2.0",
    "info": {
        "title": "NFV Service",
        "description": "APIs to get metrics for NFV services",
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
        "/public/namespaces/{namespace}/nfv_services/metrics": {
            "post": {
                "summary": "Metrics",
                "description": "Nfv Service metrics",
                "operationId": "ves.io.schema.nfv_service.CustomDataAPI.Metrics",
                "responses": {
                    "200": {
                        "description": "A successful response.",
                        "schema": {
                            "$ref": "#/definitions/nfv_serviceMetricsResponse"
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
                        "description": "Namespace\n\nx-example: \"system\"\nOnly -system- namespace is supported.",
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
                            "$ref": "#/definitions/nfv_serviceMetricsRequest"
                        }
                    }
                ],
                "tags": [
                    "CustomDataAPI"
                ],
                "externalDocs": {
                    "description": "Examples of this operation",
                    "url": "https://www.volterra.io/docs/reference/api-ref/ves-io-schema-nfv_service-customdataapi-metrics"
                },
                "x-ves-proto-rpc": "ves.io.schema.nfv_service.CustomDataAPI.Metrics"
            },
            "x-displayname": "NFV Service Custom Data API",
            "x-ves-proto-service": "ves.io.schema.nfv_service.CustomDataAPI",
            "x-ves-proto-service-type": "CUSTOM_PUBLIC"
        }
    },
    "definitions": {
        "nfv_serviceMetricData": {
            "type": "object",
            "description": "Metric data contains the metric type and the corresponding metric value",
            "title": "Metric Data",
            "x-displayname": "Metric Data",
            "x-ves-proto-message": "ves.io.schema.nfv_service.MetricData",
            "properties": {
                "data": {
                    "type": "array",
                    "description": " List of metric data",
                    "title": "Data",
                    "items": {
                        "$ref": "#/definitions/schemanfv_serviceMetricTypeData"
                    },
                    "x-displayname": "Data"
                },
                "type": {
                    "description": " Metric Type",
                    "title": "Type",
                    "$ref": "#/definitions/schemanfv_serviceMetricType",
                    "x-displayname": "Type"
                }
            }
        },
        "nfv_serviceMetricLabel": {
            "type": "string",
            "description": "Labels in Nfv Service Metrics\n",
            "title": "Metric Label",
            "enum": [
                "SITE",
                "SERVICE",
                "SERVICE_INSTANCE",
                "SERVICE_INSTANCE_HOSTNAME"
            ],
            "default": "SITE",
            "x-displayname": "Metric Label",
            "x-ves-proto-enum": "ves.io.schema.nfv_service.MetricLabel"
        },
        "nfv_serviceMetricsRequest": {
            "type": "object",
            "description": "Request to get the metrics for Nfv services",
            "title": "Metrics Request",
            "x-displayname": "Metrics Request",
            "x-ves-proto-message": "ves.io.schema.nfv_service.MetricsRequest",
            "properties": {
                "end_time": {
                    "type": "string",
                    "description": " end time of metric collection from which data will be considered to build graph.\n Format: unix_timestamp|rfc 3339\n\n Optional: If not specified, then the end_time will be evaluated to start_time+10m\n           If start_time is not specified, then the end_time will be evaluated to \u003ccurrent time\u003e\n\nExample: - \"1570197600\"-\n\nValidation Rules:\n  ves.io.schema.rules.string.query_time: true\n",
                    "title": "End time",
                    "x-displayname": "End Time",
                    "x-ves-example": "1570197600",
                    "x-ves-validation-rules": {
                        "ves.io.schema.rules.string.query_time": "true"
                    }
                },
                "field_selector": {
                    "type": "array",
                    "description": " Select fields to be returned in the response.\n One or more fields in {TX_THROUGHPUT_BYTES, CONNECTION_RATE, TX_THROUGHPUT_PACKETS, RX_THROUGHPUT_BYTES, RX_THROUGHPUT_PACKETS} can be specified.\n\nRequired: YES\n\nValidation Rules:\n  ves.io.schema.rules.message.required: true\n  ves.io.schema.rules.repeated.max_items: 8\n  ves.io.schema.rules.repeated.min_items: 1\n  ves.io.schema.rules.repeated.unique: true\n",
                    "title": "Field Selector",
                    "minItems": 1,
                    "maxItems": 8,
                    "items": {
                        "$ref": "#/definitions/schemanfv_serviceMetricType"
                    },
                    "x-displayname": "Field Selector",
                    "x-ves-required": "true",
                    "x-ves-validation-rules": {
                        "ves.io.schema.rules.message.required": "true",
                        "ves.io.schema.rules.repeated.max_items": "8",
                        "ves.io.schema.rules.repeated.min_items": "1",
                        "ves.io.schema.rules.repeated.unique": "true"
                    }
                },
                "filter": {
                    "type": "string",
                    "description": " filter is used to specify the list of matchers\n syntax for filter := {[\u003cmatcher\u003e]}\n \u003cmatcher\u003e := \u003clabel\u003e\u003coperator\u003e\"\u003cvalue\u003e\"\n   \u003clabel\u003e := string\n     One or more labels defined in MetricLabel {SITE, SERVICE, SERVICE_INSTANCE} can be specified in the filter.\n   \u003cvalue\u003e := string\n   \u003coperator\u003e := [\"=\"|\"!=\"|\"=~\"|\"!~\"]\n     =  : equal to\n     != : not equal to\n     =~ : regex match\n     !~ : not regex match\n\n Optional: If not specified, metric will be aggregated based on the group_by labels.\n\nExample: - \"{HOSTNAME=\\\"WAF-1\\\"}\"-",
                    "title": "Label Filter",
                    "x-displayname": "Filter",
                    "x-ves-example": "{HOSTNAME=\\\"WAF-1\\\"}"
                },
                "group_by": {
                    "type": "array",
                    "description": " Aggregate data by zero or more labels {SITE, SERVICE, SERVICE_INSTANCE}",
                    "title": "Group by",
                    "items": {
                        "$ref": "#/definitions/nfv_serviceMetricLabel"
                    },
                    "x-displayname": "Group By"
                },
                "namespace": {
                    "type": "string",
                    "description": " Only -system- namespace is supported.\n\nExample: - \"system\"-",
                    "title": "Namespace",
                    "x-displayname": "Namespace",
                    "x-ves-example": "system"
                },
                "start_time": {
                    "type": "string",
                    "description": "\n start time of metric collection from which data will be considered to build graph.\n Format: unix_timestamp|rfc 3339\n\n Optional: If not specified, then the start_time will be evaluated to end_time-10m\n           If end_time is not specified, then the start_time will be evaluated to \u003ccurrent time\u003e-10m\n\nExample: - \"1570194000\"-\n\nValidation Rules:\n  ves.io.schema.rules.string.query_time: true\n",
                    "title": "Start time",
                    "x-displayname": "Start Time",
                    "x-ves-example": "1570194000",
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
        "nfv_serviceMetricsResponse": {
            "type": "object",
            "description": "Metrics for NFV Services",
            "title": "Metrics Response",
            "x-displayname": "Metrics Response",
            "x-ves-proto-message": "ves.io.schema.nfv_service.MetricsResponse",
            "properties": {
                "data": {
                    "type": "array",
                    "description": " Data contains time-series metric data for Nfv services",
                    "title": "Data",
                    "items": {
                        "$ref": "#/definitions/nfv_serviceMetricData"
                    },
                    "x-displayname": "Data"
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
        },
        "schemanfv_serviceMetricType": {
            "type": "string",
            "description": "List of metric types\n\nx-unit: \"Bps\"\nx-unit: \"Cps\"\nx-unit: \"pps\"\nx-unit: \"Bps\"\nx-unit: \"pps\"\nx-unit: \"percent\"\nx-unit: \"count\"\nx-unit: \"count\"\nx-unit: \"bytes\"\nx-unit: \"bytes\"\nx-unit: \"bytes\"\nx-unit: \"bytes\"\nx-unit: \"count\"\nx-unit: \"count\"\nx-unit: \"percentage\"",
            "title": "Metric Type",
            "enum": [
                "TX_THROUGHPUT_BYTES",
                "CONNECTION_RATE",
                "TX_THROUGHPUT_PACKETS",
                "RX_THROUGHPUT_BYTES",
                "RX_THROUGHPUT_PACKETS",
                "CPU_UTILIZATION",
                "DISK_READ_OPS",
                "DISK_WRITE_OPS",
                "DISK_READ_BYTES",
                "DISK_WRITE_BYTES",
                "NETWORK_IN_BYTES",
                "NETWORK_OUT_BYTES",
                "NETWORK_IN_PACKETS",
                "NETWORK_OUT_PACKETS",
                "HEALTHSCORE"
            ],
            "default": "TX_THROUGHPUT_BYTES",
            "x-displayname": "Metric Type",
            "x-ves-proto-enum": "ves.io.schema.nfv_service.MetricType"
        },
        "schemanfv_serviceMetricTypeData": {
            "type": "object",
            "description": "Metric Type Data contains key/value pair that uniquely identifies the group_by labels specified in the request.",
            "title": "Metric Type Data",
            "x-displayname": "Metric Type Data",
            "x-ves-proto-message": "ves.io.schema.nfv_service.MetricTypeData",
            "properties": {
                "key": {
                    "type": "object",
                    "description": " Key contains the name/value pair.\n \"name\" is the label name defined in \"MetricLabel\"",
                    "title": "Key",
                    "x-displayname": "Key"
                },
                "value": {
                    "type": "array",
                    "description": " List of metric values.",
                    "title": "Value",
                    "items": {
                        "$ref": "#/definitions/schemaMetricValue"
                    },
                    "x-displayname": "Value"
                }
            }
        }
    },
    "x-displayname": "NFV Service",
    "x-ves-proto-file": "ves.io/schema/nfv_service/public_custom_data_api.proto"
}`
