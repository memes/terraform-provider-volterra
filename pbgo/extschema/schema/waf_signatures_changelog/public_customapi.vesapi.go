// Copyright (c) 2022 F5, Inc. All rights reserved.
// Code generated by ves-gen-schema-go. DO NOT EDIT.

package waf_signatures_changelog

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

// Create WafSignatureChangelogCustomApi GRPC Client satisfying server.CustomClient
type WafSignatureChangelogCustomApiGrpcClient struct {
	conn       *grpc.ClientConn
	grpcClient WafSignatureChangelogCustomApiClient
	// map of rpc name to its invocation
	rpcFns map[string]func(context.Context, string, ...grpc.CallOption) (proto.Message, error)
}

func (c *WafSignatureChangelogCustomApiGrpcClient) doRPCGetReleasedSignatures(ctx context.Context, yamlReq string, opts ...grpc.CallOption) (proto.Message, error) {
	req := &ReleasedSignaturesReq{}
	if err := codec.FromYAML(yamlReq, req); err != nil {
		return nil, fmt.Errorf("YAML Request %s is not of type *ves.io.schema.waf_signatures_changelog.ReleasedSignaturesReq", yamlReq)
	}
	rsp, err := c.grpcClient.GetReleasedSignatures(ctx, req, opts...)
	return rsp, err
}

func (c *WafSignatureChangelogCustomApiGrpcClient) DoRPC(ctx context.Context, rpc string, opts ...server.CustomCallOpt) (proto.Message, error) {
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

func NewWafSignatureChangelogCustomApiGrpcClient(cc *grpc.ClientConn) server.CustomClient {
	ccl := &WafSignatureChangelogCustomApiGrpcClient{
		conn:       cc,
		grpcClient: NewWafSignatureChangelogCustomApiClient(cc),
	}
	rpcFns := make(map[string]func(context.Context, string, ...grpc.CallOption) (proto.Message, error))
	rpcFns["GetReleasedSignatures"] = ccl.doRPCGetReleasedSignatures

	ccl.rpcFns = rpcFns

	return ccl
}

// Create WafSignatureChangelogCustomApi REST Client satisfying server.CustomClient
type WafSignatureChangelogCustomApiRestClient struct {
	baseURL string
	client  http.Client
	// map of rpc name to its invocation
	rpcFns map[string]func(context.Context, *server.CustomCallOpts) (proto.Message, error)
}

func (c *WafSignatureChangelogCustomApiRestClient) doRPCGetReleasedSignatures(ctx context.Context, callOpts *server.CustomCallOpts) (proto.Message, error) {
	if callOpts.URI == "" {
		return nil, fmt.Errorf("Error, URI should be specified, got empty")
	}
	url := fmt.Sprintf("%s%s", c.baseURL, callOpts.URI)

	yamlReq := callOpts.YAMLReq
	req := &ReleasedSignaturesReq{}
	if err := codec.FromYAML(yamlReq, req); err != nil {
		return nil, fmt.Errorf("YAML Request %s is not of type *ves.io.schema.waf_signatures_changelog.ReleasedSignaturesReq: %s", yamlReq, err)
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
		q.Add("namespace", fmt.Sprintf("%v", req.Namespace))
		q.Add("vh_name", fmt.Sprintf("%v", req.VhName))

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
	pbRsp := &ReleasedSignaturesRsp{}
	if err := codec.FromJSON(string(body), pbRsp); err != nil {
		return nil, errors.Wrapf(err, "JSON Response %s is not of type *ves.io.schema.waf_signatures_changelog.ReleasedSignaturesRsp", body)

	}
	if callOpts.OutCallResponse != nil {
		callOpts.OutCallResponse.ProtoMsg = pbRsp
		callOpts.OutCallResponse.JSON = string(body)
	}
	return pbRsp, nil
}

func (c *WafSignatureChangelogCustomApiRestClient) DoRPC(ctx context.Context, rpc string, opts ...server.CustomCallOpt) (proto.Message, error) {
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

func NewWafSignatureChangelogCustomApiRestClient(baseURL string, hc http.Client) server.CustomClient {
	ccl := &WafSignatureChangelogCustomApiRestClient{
		baseURL: baseURL,
		client:  hc,
	}

	rpcFns := make(map[string]func(context.Context, *server.CustomCallOpts) (proto.Message, error))
	rpcFns["GetReleasedSignatures"] = ccl.doRPCGetReleasedSignatures

	ccl.rpcFns = rpcFns

	return ccl
}

// Create wafSignatureChangelogCustomApiInprocClient

// INPROC Client (satisfying WafSignatureChangelogCustomApiClient interface)
type wafSignatureChangelogCustomApiInprocClient struct {
	WafSignatureChangelogCustomApiServer
}

func (c *wafSignatureChangelogCustomApiInprocClient) GetReleasedSignatures(ctx context.Context, in *ReleasedSignaturesReq, opts ...grpc.CallOption) (*ReleasedSignaturesRsp, error) {
	ctx = server.ContextFromInprocReq(ctx, "ves.io.schema.waf_signatures_changelog.WafSignatureChangelogCustomApi.GetReleasedSignatures", nil)
	return c.WafSignatureChangelogCustomApiServer.GetReleasedSignatures(ctx, in)
}

func NewWafSignatureChangelogCustomApiInprocClient(svc svcfw.Service) WafSignatureChangelogCustomApiClient {
	return &wafSignatureChangelogCustomApiInprocClient{WafSignatureChangelogCustomApiServer: NewWafSignatureChangelogCustomApiServer(svc)}
}

// RegisterGwWafSignatureChangelogCustomApiHandler registers with grpc-gw with an inproc-client backing so that
// rest to grpc happens without a grpc.Dial (thus avoiding additional certs for mTLS)
func RegisterGwWafSignatureChangelogCustomApiHandler(ctx context.Context, mux *runtime.ServeMux, svc interface{}) error {
	s, ok := svc.(svcfw.Service)
	if !ok {
		return fmt.Errorf("svc is not svcfw.Service")
	}
	return RegisterWafSignatureChangelogCustomApiHandlerClient(ctx, mux, NewWafSignatureChangelogCustomApiInprocClient(s))
}

// Create wafSignatureChangelogCustomApiSrv

// SERVER (satisfying WafSignatureChangelogCustomApiServer interface)
type wafSignatureChangelogCustomApiSrv struct {
	svc svcfw.Service
}

func (s *wafSignatureChangelogCustomApiSrv) GetReleasedSignatures(ctx context.Context, in *ReleasedSignaturesReq) (*ReleasedSignaturesRsp, error) {
	ah := s.svc.GetAPIHandler("ves.io.schema.waf_signatures_changelog.WafSignatureChangelogCustomApi")
	cah, ok := ah.(WafSignatureChangelogCustomApiServer)
	if !ok {
		return nil, fmt.Errorf("ah %v is not of type *WafSignatureChangelogCustomApiServer", ah)
	}

	var (
		rsp *ReleasedSignaturesRsp
		err error
	)

	bodyFields := svcfw.GenAuditReqBodyFields(ctx, s.svc, "ves.io.schema.waf_signatures_changelog.ReleasedSignaturesReq", in)
	defer func() {
		if len(bodyFields) > 0 {
			server.ExtendAPIAudit(ctx, svcfw.PublicAPIBodyLog.Uid, bodyFields)
		}
		userMsg := "The 'WafSignatureChangelogCustomApi.GetReleasedSignatures' operation on 'waf_signatures_changelog'"
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
		if rvFn := s.svc.GetRPCValidator("ves.io.schema.waf_signatures_changelog.WafSignatureChangelogCustomApi.GetReleasedSignatures"); rvFn != nil {
			if verr := rvFn(ctx, in); verr != nil {
				err = server.MaybePublicRestError(ctx, errors.Wrapf(verr, "Validating Request"))
				return nil, server.GRPCStatusFromError(err).Err()
			}
		}
	}

	rsp, err = cah.GetReleasedSignatures(ctx, in)
	if err != nil {
		return rsp, server.GRPCStatusFromError(server.MaybePublicRestError(ctx, err)).Err()
	}

	bodyFields = append(bodyFields, svcfw.GenAuditRspBodyFields(ctx, s.svc, "ves.io.schema.waf_signatures_changelog.ReleasedSignaturesRsp", rsp)...)

	return rsp, nil
}

func NewWafSignatureChangelogCustomApiServer(svc svcfw.Service) WafSignatureChangelogCustomApiServer {
	return &wafSignatureChangelogCustomApiSrv{svc: svc}
}

var WafSignatureChangelogCustomApiSwaggerJSON string = `{
    "swagger": "2.0",
    "info": {
        "title": "WAF signature related Custom APIs",
        "description": "WAF Signatures Changelog custom APIs",
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
        "/public/namespaces/{namespace}/virtual_hosts/{vh_name}/released_signatures": {
            "get": {
                "summary": "Released Signatures",
                "description": "API to get Released Signatures",
                "operationId": "ves.io.schema.waf_signatures_changelog.WafSignatureChangelogCustomApi.GetReleasedSignatures",
                "responses": {
                    "200": {
                        "description": "A successful response.",
                        "schema": {
                            "$ref": "#/definitions/waf_signatures_changelogReleasedSignaturesRsp"
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
                        "description": "namespace\n\nx-example: \"shared\"\nFetch waf signatures changelog for the given namespace",
                        "in": "path",
                        "required": true,
                        "type": "string",
                        "x-displayname": "Namespace"
                    },
                    {
                        "name": "vh_name",
                        "description": "vh_name\n\nx-example: \"blogging-app\"\nVirtual Host for current request",
                        "in": "path",
                        "required": true,
                        "type": "string",
                        "x-displayname": "Virtual Host Name"
                    }
                ],
                "tags": [
                    "WafSignatureChangelogCustomApi"
                ],
                "externalDocs": {
                    "description": "Examples of this operation",
                    "url": "https://www.volterra.io/docs/reference/api-ref/ves-io-schema-waf_signatures_changelog-wafsignaturechangelogcustomapi-getreleasedsignatures"
                },
                "x-ves-proto-rpc": "ves.io.schema.waf_signatures_changelog.WafSignatureChangelogCustomApi.GetReleasedSignatures"
            },
            "x-displayname": "WAF Signature Changelog Custom API",
            "x-ves-proto-service": "ves.io.schema.waf_signatures_changelog.WafSignatureChangelogCustomApi",
            "x-ves-proto-service-type": "CUSTOM_PUBLIC"
        }
    },
    "definitions": {
        "schemawaf_signatures_changelogGlobalSpecType": {
            "type": "object",
            "title": "GlobalSpecType",
            "x-displayname": "Signatures Changelog Specification",
            "x-ves-proto-message": "ves.io.schema.waf_signatures_changelog.GlobalSpecType",
            "properties": {
                "added_signatures": {
                    "type": "array",
                    "description": " A list of new signature ids in the release.\n\nExample: - [\"200101852\", \"200103290\"]-",
                    "title": "Added Signatures",
                    "items": {
                        "type": "string"
                    },
                    "x-displayname": "Added Signatures",
                    "x-ves-deprecated": "Replaced by new signatures"
                },
                "added_signatures_data": {
                    "type": "array",
                    "description": " A list of new signatures in the release.",
                    "title": "Added Signatures Data",
                    "items": {
                        "$ref": "#/definitions/waf_signatures_changelogSignature"
                    },
                    "x-displayname": "Added Signatures Data"
                },
                "release_date": {
                    "type": "string",
                    "description": " Changelog release date.\n\nExample: - \"2022-12-06-",
                    "title": "Release Date",
                    "x-displayname": "Release Date"
                },
                "updated_signatures": {
                    "type": "array",
                    "description": " A list of updated signature ids in the release.\n\nExample: - [\"200101852\", \"200103290\"]-",
                    "title": "Updated Signatures",
                    "items": {
                        "type": "string"
                    },
                    "x-displayname": "Updated Signatures",
                    "x-ves-deprecated": "Replaced by modified signatures"
                },
                "updated_signatures_data": {
                    "type": "array",
                    "description": " A list of updated signatures in the release.",
                    "title": "Updated Signatures Data",
                    "items": {
                        "$ref": "#/definitions/waf_signatures_changelogSignature"
                    },
                    "x-displayname": "Updated Signatures Data"
                }
            }
        },
        "waf_signatures_changelogReleasedSignaturesRsp": {
            "type": "object",
            "description": "Response to get the list of released signatures",
            "title": "ReleasedSignaturesRes",
            "x-displayname": "Released Signatures Response",
            "x-ves-proto-message": "ves.io.schema.waf_signatures_changelog.ReleasedSignaturesRsp",
            "properties": {
                "released_signatures": {
                    "type": "array",
                    "description": " Released Signatures",
                    "title": "released_signatures",
                    "items": {
                        "$ref": "#/definitions/schemawaf_signatures_changelogGlobalSpecType"
                    },
                    "x-displayname": "Released Signatures"
                }
            }
        },
        "waf_signatures_changelogSignature": {
            "type": "object",
            "x-ves-proto-message": "ves.io.schema.waf_signatures_changelog.Signature",
            "properties": {
                "accuracy": {
                    "type": "string",
                    "description": " The Signature Accuracy\n\nExample: - \"Medium\"-",
                    "title": "accuracy",
                    "x-displayname": "Accuracy",
                    "x-ves-example": "Medium"
                },
                "applies_to": {
                    "type": "string",
                    "description": " The Signature Applies to\n\nExample: - \"Request\"-",
                    "title": "applies_to",
                    "x-displayname": "Applies To",
                    "x-ves-example": "Request"
                },
                "attack_type": {
                    "type": "string",
                    "description": " The Signature Attack Type\n\nExample: - \"Server Side Code Injection\"-",
                    "title": "attack_type",
                    "x-displayname": "Attack Type",
                    "x-ves-example": "Server Side Code Injection"
                },
                "description": {
                    "type": "string",
                    "description": " The Signature Description\n\nExample: - \"Summary:\\nThis event is generated when an attempt is made to inject Server-side Include code. This is a general attack detection signature (i.e. it is not specific to any web application).\\n\\nImpact:\\nSerious. Execution of arbitrary commands may be possible.\\n\\nDetailed Information:\\nThis event indicates that an attempt has been made to inject Server-side Include (SSI) code. SSI Injection allows to attacker to send server side code that could be executed locally by the web server.\\n\\nAffected Systems:\\nAll systems.\\n\\nAttack Scenarios:\\nThere are many possible.\\n\\nEase Of Attack:\\nSimple to medium.\\n\\nFalse Positives:\\nSome applications may accept valid input which matches these signatures.\\n\\nFalse Negatives:\\nNone known.\\n\\nCorrective Action:\\nEnsure the system is using an up to date version of the software and has had all vendor supplied patches applied. Utilize \\\"Positive Security Model\\\" by accepting only known types of input in web application.\\n\\nAdditional References:\\nhttp://www.webappsec.org/projects/threat/classes/ssi_injection.shtml\\n\\n\"-",
                    "title": "description",
                    "x-displayname": "Description",
                    "x-ves-example": "Summary:\\nThis event is generated when an attempt is made to inject Server-side Include code. This is a general attack detection signature (i.e. it is not specific to any web application).\\n\\nImpact:\\nSerious. Execution of arbitrary commands may be possible.\\n\\nDetailed Information:\\nThis event indicates that an attempt has been made to inject Server-side Include (SSI) code. SSI Injection allows to attacker to send server side code that could be executed locally by the web server.\\n\\nAffected Systems:\\nAll systems.\\n\\nAttack Scenarios:\\nThere are many possible.\\n\\nEase Of Attack:\\nSimple to medium.\\n\\nFalse Positives:\\nSome applications may accept valid input which matches these signatures.\\n\\nFalse Negatives:\\nNone known.\\n\\nCorrective Action:\\nEnsure the system is using an up to date version of the software and has had all vendor supplied patches applied. Utilize \\\"Positive Security Model\\\" by accepting only known types of input in web application.\\n\\nAdditional References:\\nhttp://www.webappsec.org/projects/threat/classes/ssi_injection.shtml\\n\\n"
                },
                "id": {
                    "type": "string",
                    "description": " The Signature ID\n\nExample: - \"200104853\"-",
                    "title": "id",
                    "format": "int64",
                    "x-displayname": "ID",
                    "x-ves-example": "200104853"
                },
                "last_update": {
                    "type": "string",
                    "description": " The Signature last update time\n\nExample: - \"2022/11/29 20:19:17\"-\n\nRequired: YES\n\nValidation Rules:\n  ves.io.schema.rules.message.required: true\n",
                    "title": "last_update",
                    "x-displayname": "Last Update",
                    "x-ves-example": "2022/11/29 20:19:17",
                    "x-ves-required": "true",
                    "x-ves-validation-rules": {
                        "ves.io.schema.rules.message.required": "true"
                    }
                },
                "name": {
                    "type": "string",
                    "description": " The Signature Name\n\nExample: - \"Java code injection FreeMarker - objectWrapper (URI)\"-",
                    "title": "name",
                    "x-displayname": "Name",
                    "x-ves-example": "Java code injection FreeMarker - objectWrapper (URI)"
                },
                "references": {
                    "type": "array",
                    "description": " The Signature References\n\nExample: - \"['http://cve.mitre.org/cgi-bin/cvename.cgi?name=CVE-2022-26377']\"-",
                    "title": "references",
                    "items": {
                        "type": "string"
                    },
                    "x-displayname": "References",
                    "x-ves-example": "['http://cve.mitre.org/cgi-bin/cvename.cgi?name=CVE-2022-26377']"
                },
                "risk": {
                    "type": "string",
                    "description": " The Signature Risk\n\nExample: - \"High\"-",
                    "title": "risk",
                    "x-displayname": "Risk",
                    "x-ves-example": "High"
                },
                "systems": {
                    "type": "array",
                    "description": " The Signature Systems\n\nExample: - \"['Java Servlets/JSP', 'Apache Struts', 'JavaServer Faces (JSF)']\"-",
                    "title": "systems",
                    "items": {
                        "type": "string"
                    },
                    "x-displayname": "Systems",
                    "x-ves-example": "['Java Servlets/JSP', 'Apache Struts', 'JavaServer Faces (JSF)']"
                }
            }
        }
    },
    "x-displayname": "WAF Signatures Changelog",
    "x-ves-proto-file": "ves.io/schema/waf_signatures_changelog/public_customapi.proto"
}`
