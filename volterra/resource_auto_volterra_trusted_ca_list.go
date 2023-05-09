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
	ves_io_schema_trusted_ca_list "github.com/volterraedge/terraform-provider-volterra/pbgo/extschema/schema/trusted_ca_list"
)

// resourceVolterraTrustedCaList is implementation of Volterra's TrustedCaList resources
func resourceVolterraTrustedCaList() *schema.Resource {
	return &schema.Resource{
		Create: resourceVolterraTrustedCaListCreate,
		Read:   resourceVolterraTrustedCaListRead,
		Update: resourceVolterraTrustedCaListUpdate,
		Delete: resourceVolterraTrustedCaListDelete,

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

			"trusted_ca_url": {
				Type:     schema.TypeString,
				Optional: true,
			},
		},
	}
}

// resourceVolterraTrustedCaListCreate creates TrustedCaList resource
func resourceVolterraTrustedCaListCreate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*APIClient)

	createMeta := &ves_io_schema.ObjectCreateMetaType{}
	createSpec := &ves_io_schema_trusted_ca_list.CreateSpecType{}
	createReq := &ves_io_schema_trusted_ca_list.CreateRequest{
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

	//trusted_ca_url
	if v, ok := d.GetOk("trusted_ca_url"); ok && !isIntfNil(v) {

		createSpec.TrustedCaUrl =
			v.(string)

	}

	log.Printf("[DEBUG] Creating Volterra TrustedCaList object with struct: %+v", createReq)

	createTrustedCaListResp, err := client.CreateObject(context.Background(), ves_io_schema_trusted_ca_list.ObjectType, createReq)
	if err != nil {
		return fmt.Errorf("error creating TrustedCaList: %s", err)
	}
	d.SetId(createTrustedCaListResp.GetObjSystemMetadata().GetUid())

	return resourceVolterraTrustedCaListRead(d, meta)
}

func resourceVolterraTrustedCaListRead(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*APIClient)
	name := d.Get("name").(string)
	namespace := d.Get("namespace").(string)

	resp, err := client.GetObject(context.Background(), ves_io_schema_trusted_ca_list.ObjectType, namespace, name)
	if err != nil {
		if strings.Contains(err.Error(), "status code 404") {
			log.Printf("[INFO] TrustedCaList %s no longer exists", d.Id())
			d.SetId("")
			return nil
		}
		return fmt.Errorf("Error finding Volterra TrustedCaList %q: %s", d.Id(), err)
	}
	return setTrustedCaListFields(client, d, resp)
}

func setTrustedCaListFields(client *APIClient, d *schema.ResourceData, resp vesapi.GetObjectResponse) error {
	metadata := resp.GetObjMetadata()

	d.Set("annotations", metadata.GetAnnotations())

	d.Set("description", metadata.GetDescription())

	d.Set("disable", metadata.GetDisable())

	d.Set("labels", metadata.GetLabels())

	d.Set("name", metadata.GetName())

	d.Set("namespace", metadata.GetNamespace())

	return nil
}

// resourceVolterraTrustedCaListUpdate updates TrustedCaList resource
func resourceVolterraTrustedCaListUpdate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*APIClient)

	updateMeta := &ves_io_schema.ObjectReplaceMetaType{}
	updateSpec := &ves_io_schema_trusted_ca_list.ReplaceSpecType{}
	updateReq := &ves_io_schema_trusted_ca_list.ReplaceRequest{
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

	if v, ok := d.GetOk("trusted_ca_url"); ok && !isIntfNil(v) {

		updateSpec.TrustedCaUrl =
			v.(string)

	}

	log.Printf("[DEBUG] Updating Volterra TrustedCaList obj with struct: %+v", updateReq)

	err := client.ReplaceObject(context.Background(), ves_io_schema_trusted_ca_list.ObjectType, updateReq)
	if err != nil {
		return fmt.Errorf("error updating TrustedCaList: %s", err)
	}

	return resourceVolterraTrustedCaListRead(d, meta)
}

func resourceVolterraTrustedCaListDelete(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*APIClient)
	name := d.Get("name").(string)
	namespace := d.Get("namespace").(string)

	_, err := client.GetObject(context.Background(), ves_io_schema_trusted_ca_list.ObjectType, namespace, name)
	if err != nil {
		if strings.Contains(err.Error(), "status code 404") {
			log.Printf("[INFO] TrustedCaList %s no longer exists", d.Id())
			d.SetId("")
			return nil
		}
		return fmt.Errorf("Error finding Volterra TrustedCaList before deleting %q: %s", d.Id(), err)
	}

	log.Printf("[DEBUG] Deleting Volterra TrustedCaList obj with name %+v in namespace %+v", name, namespace)
	return client.DeleteObject(context.Background(), ves_io_schema_trusted_ca_list.ObjectType, namespace, name)
}
