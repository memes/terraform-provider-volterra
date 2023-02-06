//
// Copyright (c) 2018 Volterra, Inc. All rights reserved.
// Code generated by ves-gen-tf-provider. DO NOT EDIT.
//

package volterra

import (
	"context"
	"fmt"
	"log"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"gopkg.volterra.us/stdlib/client/vesapi"

	ves_io_schema "github.com/volterraedge/terraform-provider-volterra/pbgo/extschema/schema"
	ves_io_schema_api_group_element "github.com/volterraedge/terraform-provider-volterra/pbgo/extschema/schema/api_group_element"
	ves_io_schema_views "github.com/volterraedge/terraform-provider-volterra/pbgo/extschema/schema/views"
	ves_io_schema_views_api_definition "github.com/volterraedge/terraform-provider-volterra/pbgo/extschema/schema/views/api_definition"
	ves_io_schema_views_app_api_group "github.com/volterraedge/terraform-provider-volterra/pbgo/extschema/schema/views/app_api_group"
)

// resourceVolterraAppApiGroup is implementation of Volterra's AppApiGroup resources
func resourceVolterraAppApiGroup() *schema.Resource {
	return &schema.Resource{
		Create: resourceVolterraAppApiGroupCreate,
		Read:   resourceVolterraAppApiGroupRead,
		Update: resourceVolterraAppApiGroupUpdate,
		Delete: resourceVolterraAppApiGroupDelete,

		Schema: map[string]*schema.Schema{

			"annotations": {
				Type:     schema.TypeMap,
				Optional: true,
			},

			"description": {
				Type:     schema.TypeString,
				Optional: true,
			},

			"disable": {
				Type:     schema.TypeBool,
				Optional: true,
			},

			"labels": {
				Type:     schema.TypeMap,
				Optional: true,
			},

			"name": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},

			"namespace": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},

			"api_group_builder": {

				Type:     schema.TypeSet,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"excluded_operations": {

							Type:     schema.TypeList,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"method": {
										Type:     schema.TypeString,
										Optional: true,
									},

									"path": {
										Type:     schema.TypeString,
										Optional: true,
									},
								},
							},
						},

						"included_operations": {

							Type:     schema.TypeList,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"method": {
										Type:     schema.TypeString,
										Optional: true,
									},

									"path": {
										Type:     schema.TypeString,
										Optional: true,
									},
								},
							},
						},

						"label_filter": {

							Type:     schema.TypeSet,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"expressions": {

										Type: schema.TypeList,

										Required: true,
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},
								},
							},
						},

						"metadata": {

							Type:     schema.TypeSet,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"description": {
										Type:     schema.TypeString,
										Optional: true,
									},

									"disable": {
										Type:     schema.TypeBool,
										Optional: true,
									},

									"name": {
										Type:     schema.TypeString,
										Optional: true,
									},
								},
							},
						},

						"path_filter": {
							Type:     schema.TypeString,
							Optional: true,
						},
					},
				},
			},

			"elements": {

				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"methods": {

							Type: schema.TypeList,

							Required: true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},

						"path_regex": {
							Type:     schema.TypeString,
							Optional: true,
						},
					},
				},
			},

			"api_definition": {

				Type:     schema.TypeSet,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"api_definition": {

							Type:     schema.TypeSet,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"name": {
										Type:     schema.TypeString,
										Optional: true,
									},
									"namespace": {
										Type:     schema.TypeString,
										Optional: true,
									},
									"tenant": {
										Type:     schema.TypeString,
										Optional: true,
									},
								},
							},
						},
					},
				},
			},

			"generic": {

				Type:     schema.TypeBool,
				Optional: true,
			},

			"http_loadbalancer": {

				Type:     schema.TypeSet,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"http_loadbalancer": {

							Type:     schema.TypeSet,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"name": {
										Type:     schema.TypeString,
										Optional: true,
									},
									"namespace": {
										Type:     schema.TypeString,
										Optional: true,
									},
									"tenant": {
										Type:     schema.TypeString,
										Optional: true,
									},
								},
							},
						},
					},
				},
			},
		},
	}
}

// resourceVolterraAppApiGroupCreate creates AppApiGroup resource
func resourceVolterraAppApiGroupCreate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*APIClient)

	createMeta := &ves_io_schema.ObjectCreateMetaType{}
	createSpec := &ves_io_schema_views_app_api_group.CreateSpecType{}
	createReq := &ves_io_schema_views_app_api_group.CreateRequest{
		Metadata: createMeta,
		Spec:     createSpec,
	}

	if v, ok := d.GetOk("annotations"); ok && !isIntfNil(v) {

		ms := map[string]string{}

		for k, v := range v.(map[string]interface{}) {
			val := v.(string)
			ms[k] = val
		}
		createMeta.Annotations = ms
	}

	if v, ok := d.GetOk("description"); ok && !isIntfNil(v) {
		createMeta.Description =
			v.(string)
	}

	if v, ok := d.GetOk("disable"); ok && !isIntfNil(v) {
		createMeta.Disable =
			v.(bool)
	}

	if v, ok := d.GetOk("labels"); ok && !isIntfNil(v) {

		ms := map[string]string{}

		for k, v := range v.(map[string]interface{}) {
			val := v.(string)
			ms[k] = val
		}
		createMeta.Labels = ms
	}

	if v, ok := d.GetOk("name"); ok && !isIntfNil(v) {
		createMeta.Name =
			v.(string)
	}

	if v, ok := d.GetOk("namespace"); ok && !isIntfNil(v) {
		createMeta.Namespace =
			v.(string)
	}

	//api_group_builder
	if v, ok := d.GetOk("api_group_builder"); ok && !isIntfNil(v) {

		sl := v.(*schema.Set).List()
		apiGroupBuilder := &ves_io_schema_views_api_definition.ApiGroupBuilder{}
		createSpec.ApiGroupBuilder = apiGroupBuilder
		for _, set := range sl {
			apiGroupBuilderMapStrToI := set.(map[string]interface{})

			if v, ok := apiGroupBuilderMapStrToI["excluded_operations"]; ok && !isIntfNil(v) {

				sl := v.([]interface{})
				excludedOperations := make([]*ves_io_schema_views_api_definition.ApiOperation, len(sl))
				apiGroupBuilder.ExcludedOperations = excludedOperations
				for i, set := range sl {
					excludedOperations[i] = &ves_io_schema_views_api_definition.ApiOperation{}
					excludedOperationsMapStrToI := set.(map[string]interface{})

					if v, ok := excludedOperationsMapStrToI["method"]; ok && !isIntfNil(v) {

						excludedOperations[i].Method = ves_io_schema.HttpMethod(ves_io_schema.HttpMethod_value[v.(string)])

					}

					if w, ok := excludedOperationsMapStrToI["path"]; ok && !isIntfNil(w) {
						excludedOperations[i].Path = w.(string)
					}

				}

			}

			if v, ok := apiGroupBuilderMapStrToI["included_operations"]; ok && !isIntfNil(v) {

				sl := v.([]interface{})
				includedOperations := make([]*ves_io_schema_views_api_definition.ApiOperation, len(sl))
				apiGroupBuilder.IncludedOperations = includedOperations
				for i, set := range sl {
					includedOperations[i] = &ves_io_schema_views_api_definition.ApiOperation{}
					includedOperationsMapStrToI := set.(map[string]interface{})

					if v, ok := includedOperationsMapStrToI["method"]; ok && !isIntfNil(v) {

						includedOperations[i].Method = ves_io_schema.HttpMethod(ves_io_schema.HttpMethod_value[v.(string)])

					}

					if w, ok := includedOperationsMapStrToI["path"]; ok && !isIntfNil(w) {
						includedOperations[i].Path = w.(string)
					}

				}

			}

			if v, ok := apiGroupBuilderMapStrToI["label_filter"]; ok && !isIntfNil(v) {

				sl := v.(*schema.Set).List()
				labelFilter := &ves_io_schema.LabelSelectorType{}
				apiGroupBuilder.LabelFilter = labelFilter
				for _, set := range sl {
					labelFilterMapStrToI := set.(map[string]interface{})

					if w, ok := labelFilterMapStrToI["expressions"]; ok && !isIntfNil(w) {
						ls := make([]string, len(w.([]interface{})))
						for i, v := range w.([]interface{}) {
							ls[i] = v.(string)
						}
						labelFilter.Expressions = ls
					}

				}

			}

			if v, ok := apiGroupBuilderMapStrToI["metadata"]; ok && !isIntfNil(v) {

				sl := v.(*schema.Set).List()
				metadata := &ves_io_schema.MessageMetaType{}
				apiGroupBuilder.Metadata = metadata
				for _, set := range sl {
					metadataMapStrToI := set.(map[string]interface{})

					if w, ok := metadataMapStrToI["description"]; ok && !isIntfNil(w) {
						metadata.Description = w.(string)
					}

					if w, ok := metadataMapStrToI["disable"]; ok && !isIntfNil(w) {
						metadata.Disable = w.(bool)
					}

					if w, ok := metadataMapStrToI["name"]; ok && !isIntfNil(w) {
						metadata.Name = w.(string)
					}

				}

			}

			if w, ok := apiGroupBuilderMapStrToI["path_filter"]; ok && !isIntfNil(w) {
				apiGroupBuilder.PathFilter = w.(string)
			}

		}

	}

	//elements
	if v, ok := d.GetOk("elements"); ok && !isIntfNil(v) {

		sl := v.([]interface{})
		elements := make([]*ves_io_schema_api_group_element.GlobalSpecType, len(sl))
		createSpec.Elements = elements
		for i, set := range sl {
			elements[i] = &ves_io_schema_api_group_element.GlobalSpecType{}
			elementsMapStrToI := set.(map[string]interface{})

			if v, ok := elementsMapStrToI["methods"]; ok && !isIntfNil(v) {

				methodsList := []ves_io_schema.HttpMethod{}
				for _, j := range v.([]interface{}) {
					methodsList = append(methodsList, ves_io_schema.HttpMethod(ves_io_schema.HttpMethod_value[j.(string)]))
				}
				elements[i].Methods = methodsList

			}

			if w, ok := elementsMapStrToI["path_regex"]; ok && !isIntfNil(w) {
				elements[i].PathRegex = w.(string)
			}

		}

	}

	//scope_choice

	scopeChoiceTypeFound := false

	if v, ok := d.GetOk("api_definition"); ok && !scopeChoiceTypeFound {

		scopeChoiceTypeFound = true
		scopeChoiceInt := &ves_io_schema_views_app_api_group.CreateSpecType_ApiDefinition{}
		scopeChoiceInt.ApiDefinition = &ves_io_schema_views_app_api_group.ApiGroupScopeApiDefinition{}
		createSpec.ScopeChoice = scopeChoiceInt

		sl := v.(*schema.Set).List()
		for _, set := range sl {
			cs := set.(map[string]interface{})

			if v, ok := cs["api_definition"]; ok && !isIntfNil(v) {

				sl := v.(*schema.Set).List()
				apiDefinitionInt := &ves_io_schema_views.ObjectRefType{}
				scopeChoiceInt.ApiDefinition.ApiDefinition = apiDefinitionInt

				for _, set := range sl {
					adMapToStrVal := set.(map[string]interface{})
					if val, ok := adMapToStrVal["name"]; ok && !isIntfNil(v) {
						apiDefinitionInt.Name = val.(string)
					}
					if val, ok := adMapToStrVal["namespace"]; ok && !isIntfNil(v) {
						apiDefinitionInt.Namespace = val.(string)
					}

					if val, ok := adMapToStrVal["tenant"]; ok && !isIntfNil(v) {
						apiDefinitionInt.Tenant = val.(string)
					}
				}

			}

		}

	}

	if v, ok := d.GetOk("generic"); ok && !scopeChoiceTypeFound {

		scopeChoiceTypeFound = true

		if v.(bool) {
			scopeChoiceInt := &ves_io_schema_views_app_api_group.CreateSpecType_Generic{}
			scopeChoiceInt.Generic = &ves_io_schema.Empty{}
			createSpec.ScopeChoice = scopeChoiceInt
		}

	}

	if v, ok := d.GetOk("http_loadbalancer"); ok && !scopeChoiceTypeFound {

		scopeChoiceTypeFound = true
		scopeChoiceInt := &ves_io_schema_views_app_api_group.CreateSpecType_HttpLoadbalancer{}
		scopeChoiceInt.HttpLoadbalancer = &ves_io_schema_views_app_api_group.ApiGroupScopeHttpLoadbalancer{}
		createSpec.ScopeChoice = scopeChoiceInt

		sl := v.(*schema.Set).List()
		for _, set := range sl {
			cs := set.(map[string]interface{})

			if v, ok := cs["http_loadbalancer"]; ok && !isIntfNil(v) {

				sl := v.(*schema.Set).List()
				httpLoadbalancerInt := &ves_io_schema_views.ObjectRefType{}
				scopeChoiceInt.HttpLoadbalancer.HttpLoadbalancer = httpLoadbalancerInt

				for _, set := range sl {
					hlMapToStrVal := set.(map[string]interface{})
					if val, ok := hlMapToStrVal["name"]; ok && !isIntfNil(v) {
						httpLoadbalancerInt.Name = val.(string)
					}
					if val, ok := hlMapToStrVal["namespace"]; ok && !isIntfNil(v) {
						httpLoadbalancerInt.Namespace = val.(string)
					}

					if val, ok := hlMapToStrVal["tenant"]; ok && !isIntfNil(v) {
						httpLoadbalancerInt.Tenant = val.(string)
					}
				}

			}

		}

	}

	log.Printf("[DEBUG] Creating Volterra AppApiGroup object with struct: %+v", createReq)

	createAppApiGroupResp, err := client.CreateObject(context.Background(), ves_io_schema_views_app_api_group.ObjectType, createReq)
	if err != nil {
		return fmt.Errorf("error creating AppApiGroup: %s", err)
	}
	d.SetId(createAppApiGroupResp.GetObjSystemMetadata().GetUid())

	return resourceVolterraAppApiGroupRead(d, meta)
}

func resourceVolterraAppApiGroupRead(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*APIClient)
	name := d.Get("name").(string)
	namespace := d.Get("namespace").(string)

	resp, err := client.GetObject(context.Background(), ves_io_schema_views_app_api_group.ObjectType, namespace, name)
	if err != nil {
		if strings.Contains(err.Error(), "status code 404") {
			log.Printf("[INFO] AppApiGroup %s no longer exists", d.Id())
			d.SetId("")
			return nil
		}
		return fmt.Errorf("Error finding Volterra AppApiGroup %q: %s", d.Id(), err)
	}
	return setAppApiGroupFields(client, d, resp)
}

func setAppApiGroupFields(client *APIClient, d *schema.ResourceData, resp vesapi.GetObjectResponse) error {
	metadata := resp.GetObjMetadata()

	d.Set("annotations", metadata.GetAnnotations())

	d.Set("description", metadata.GetDescription())

	d.Set("disable", metadata.GetDisable())

	d.Set("labels", metadata.GetLabels())

	d.Set("name", metadata.GetName())

	d.Set("namespace", metadata.GetNamespace())

	return nil
}

// resourceVolterraAppApiGroupUpdate updates AppApiGroup resource
func resourceVolterraAppApiGroupUpdate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*APIClient)

	updateMeta := &ves_io_schema.ObjectReplaceMetaType{}
	updateSpec := &ves_io_schema_views_app_api_group.ReplaceSpecType{}
	updateReq := &ves_io_schema_views_app_api_group.ReplaceRequest{
		Metadata: updateMeta,
		Spec:     updateSpec,
	}
	if v, ok := d.GetOk("annotations"); ok && !isIntfNil(v) {

		ms := map[string]string{}

		for k, v := range v.(map[string]interface{}) {
			val := v.(string)
			ms[k] = val
		}
		updateMeta.Annotations = ms
	}

	if v, ok := d.GetOk("description"); ok && !isIntfNil(v) {
		updateMeta.Description =
			v.(string)
	}

	if v, ok := d.GetOk("disable"); ok && !isIntfNil(v) {
		updateMeta.Disable =
			v.(bool)
	}

	if v, ok := d.GetOk("labels"); ok && !isIntfNil(v) {

		ms := map[string]string{}

		for k, v := range v.(map[string]interface{}) {
			val := v.(string)
			ms[k] = val
		}
		updateMeta.Labels = ms
	}

	if v, ok := d.GetOk("name"); ok && !isIntfNil(v) {
		updateMeta.Name =
			v.(string)
	}

	if v, ok := d.GetOk("namespace"); ok && !isIntfNil(v) {
		updateMeta.Namespace =
			v.(string)
	}

	if v, ok := d.GetOk("api_group_builder"); ok && !isIntfNil(v) {

		sl := v.(*schema.Set).List()
		apiGroupBuilder := &ves_io_schema_views_api_definition.ApiGroupBuilder{}
		updateSpec.ApiGroupBuilder = apiGroupBuilder
		for _, set := range sl {
			apiGroupBuilderMapStrToI := set.(map[string]interface{})

			if v, ok := apiGroupBuilderMapStrToI["excluded_operations"]; ok && !isIntfNil(v) {

				sl := v.([]interface{})
				excludedOperations := make([]*ves_io_schema_views_api_definition.ApiOperation, len(sl))
				apiGroupBuilder.ExcludedOperations = excludedOperations
				for i, set := range sl {
					excludedOperations[i] = &ves_io_schema_views_api_definition.ApiOperation{}
					excludedOperationsMapStrToI := set.(map[string]interface{})

					if v, ok := excludedOperationsMapStrToI["method"]; ok && !isIntfNil(v) {

						excludedOperations[i].Method = ves_io_schema.HttpMethod(ves_io_schema.HttpMethod_value[v.(string)])

					}

					if w, ok := excludedOperationsMapStrToI["path"]; ok && !isIntfNil(w) {
						excludedOperations[i].Path = w.(string)
					}

				}

			}

			if v, ok := apiGroupBuilderMapStrToI["included_operations"]; ok && !isIntfNil(v) {

				sl := v.([]interface{})
				includedOperations := make([]*ves_io_schema_views_api_definition.ApiOperation, len(sl))
				apiGroupBuilder.IncludedOperations = includedOperations
				for i, set := range sl {
					includedOperations[i] = &ves_io_schema_views_api_definition.ApiOperation{}
					includedOperationsMapStrToI := set.(map[string]interface{})

					if v, ok := includedOperationsMapStrToI["method"]; ok && !isIntfNil(v) {

						includedOperations[i].Method = ves_io_schema.HttpMethod(ves_io_schema.HttpMethod_value[v.(string)])

					}

					if w, ok := includedOperationsMapStrToI["path"]; ok && !isIntfNil(w) {
						includedOperations[i].Path = w.(string)
					}

				}

			}

			if v, ok := apiGroupBuilderMapStrToI["label_filter"]; ok && !isIntfNil(v) {

				sl := v.(*schema.Set).List()
				labelFilter := &ves_io_schema.LabelSelectorType{}
				apiGroupBuilder.LabelFilter = labelFilter
				for _, set := range sl {
					labelFilterMapStrToI := set.(map[string]interface{})

					if w, ok := labelFilterMapStrToI["expressions"]; ok && !isIntfNil(w) {
						ls := make([]string, len(w.([]interface{})))
						for i, v := range w.([]interface{}) {
							ls[i] = v.(string)
						}
						labelFilter.Expressions = ls
					}

				}

			}

			if v, ok := apiGroupBuilderMapStrToI["metadata"]; ok && !isIntfNil(v) {

				sl := v.(*schema.Set).List()
				metadata := &ves_io_schema.MessageMetaType{}
				apiGroupBuilder.Metadata = metadata
				for _, set := range sl {
					metadataMapStrToI := set.(map[string]interface{})

					if w, ok := metadataMapStrToI["description"]; ok && !isIntfNil(w) {
						metadata.Description = w.(string)
					}

					if w, ok := metadataMapStrToI["disable"]; ok && !isIntfNil(w) {
						metadata.Disable = w.(bool)
					}

					if w, ok := metadataMapStrToI["name"]; ok && !isIntfNil(w) {
						metadata.Name = w.(string)
					}

				}

			}

			if w, ok := apiGroupBuilderMapStrToI["path_filter"]; ok && !isIntfNil(w) {
				apiGroupBuilder.PathFilter = w.(string)
			}

		}

	}

	if v, ok := d.GetOk("elements"); ok && !isIntfNil(v) {

		sl := v.([]interface{})
		elements := make([]*ves_io_schema_api_group_element.GlobalSpecType, len(sl))
		updateSpec.Elements = elements
		for i, set := range sl {
			elements[i] = &ves_io_schema_api_group_element.GlobalSpecType{}
			elementsMapStrToI := set.(map[string]interface{})

			if v, ok := elementsMapStrToI["methods"]; ok && !isIntfNil(v) {

				methodsList := []ves_io_schema.HttpMethod{}
				for _, j := range v.([]interface{}) {
					methodsList = append(methodsList, ves_io_schema.HttpMethod(ves_io_schema.HttpMethod_value[j.(string)]))
				}
				elements[i].Methods = methodsList

			}

			if w, ok := elementsMapStrToI["path_regex"]; ok && !isIntfNil(w) {
				elements[i].PathRegex = w.(string)
			}

		}

	}

	scopeChoiceTypeFound := false

	if v, ok := d.GetOk("api_definition"); ok && !scopeChoiceTypeFound {

		scopeChoiceTypeFound = true
		scopeChoiceInt := &ves_io_schema_views_app_api_group.ReplaceSpecType_ApiDefinition{}
		scopeChoiceInt.ApiDefinition = &ves_io_schema_views_app_api_group.ApiGroupScopeApiDefinition{}
		updateSpec.ScopeChoice = scopeChoiceInt

		sl := v.(*schema.Set).List()
		for _, set := range sl {
			cs := set.(map[string]interface{})

			if v, ok := cs["api_definition"]; ok && !isIntfNil(v) {

				sl := v.(*schema.Set).List()
				apiDefinitionInt := &ves_io_schema_views.ObjectRefType{}
				scopeChoiceInt.ApiDefinition.ApiDefinition = apiDefinitionInt

				for _, set := range sl {
					adMapToStrVal := set.(map[string]interface{})
					if val, ok := adMapToStrVal["name"]; ok && !isIntfNil(v) {
						apiDefinitionInt.Name = val.(string)
					}
					if val, ok := adMapToStrVal["namespace"]; ok && !isIntfNil(v) {
						apiDefinitionInt.Namespace = val.(string)
					}

					if val, ok := adMapToStrVal["tenant"]; ok && !isIntfNil(v) {
						apiDefinitionInt.Tenant = val.(string)
					}
				}

			}

		}

	}

	if v, ok := d.GetOk("generic"); ok && !scopeChoiceTypeFound {

		scopeChoiceTypeFound = true

		if v.(bool) {
			scopeChoiceInt := &ves_io_schema_views_app_api_group.ReplaceSpecType_Generic{}
			scopeChoiceInt.Generic = &ves_io_schema.Empty{}
			updateSpec.ScopeChoice = scopeChoiceInt
		}

	}

	if v, ok := d.GetOk("http_loadbalancer"); ok && !scopeChoiceTypeFound {

		scopeChoiceTypeFound = true
		scopeChoiceInt := &ves_io_schema_views_app_api_group.ReplaceSpecType_HttpLoadbalancer{}
		scopeChoiceInt.HttpLoadbalancer = &ves_io_schema_views_app_api_group.ApiGroupScopeHttpLoadbalancer{}
		updateSpec.ScopeChoice = scopeChoiceInt

		sl := v.(*schema.Set).List()
		for _, set := range sl {
			cs := set.(map[string]interface{})

			if v, ok := cs["http_loadbalancer"]; ok && !isIntfNil(v) {

				sl := v.(*schema.Set).List()
				httpLoadbalancerInt := &ves_io_schema_views.ObjectRefType{}
				scopeChoiceInt.HttpLoadbalancer.HttpLoadbalancer = httpLoadbalancerInt

				for _, set := range sl {
					hlMapToStrVal := set.(map[string]interface{})
					if val, ok := hlMapToStrVal["name"]; ok && !isIntfNil(v) {
						httpLoadbalancerInt.Name = val.(string)
					}
					if val, ok := hlMapToStrVal["namespace"]; ok && !isIntfNil(v) {
						httpLoadbalancerInt.Namespace = val.(string)
					}

					if val, ok := hlMapToStrVal["tenant"]; ok && !isIntfNil(v) {
						httpLoadbalancerInt.Tenant = val.(string)
					}
				}

			}

		}

	}

	log.Printf("[DEBUG] Updating Volterra AppApiGroup obj with struct: %+v", updateReq)

	err := client.ReplaceObject(context.Background(), ves_io_schema_views_app_api_group.ObjectType, updateReq)
	if err != nil {
		return fmt.Errorf("error updating AppApiGroup: %s", err)
	}

	return resourceVolterraAppApiGroupRead(d, meta)
}

func resourceVolterraAppApiGroupDelete(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*APIClient)
	name := d.Get("name").(string)
	namespace := d.Get("namespace").(string)

	_, err := client.GetObject(context.Background(), ves_io_schema_views_app_api_group.ObjectType, namespace, name)
	if err != nil {
		if strings.Contains(err.Error(), "status code 404") {
			log.Printf("[INFO] AppApiGroup %s no longer exists", d.Id())
			d.SetId("")
			return nil
		}
		return fmt.Errorf("Error finding Volterra AppApiGroup before deleting %q: %s", d.Id(), err)
	}

	log.Printf("[DEBUG] Deleting Volterra AppApiGroup obj with name %+v in namespace %+v", name, namespace)
	return client.DeleteObject(context.Background(), ves_io_schema_views_app_api_group.ObjectType, namespace, name)
}
