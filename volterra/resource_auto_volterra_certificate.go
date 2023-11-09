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

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"gopkg.volterra.us/stdlib/client/vesapi"

	ves_io_schema "github.com/volterraedge/terraform-provider-volterra/pbgo/extschema/schema"
	ves_io_schema_certificate "github.com/volterraedge/terraform-provider-volterra/pbgo/extschema/schema/certificate"
	ves_io_schema_views "github.com/volterraedge/terraform-provider-volterra/pbgo/extschema/schema/views"
)

// resourceVolterraCertificate is implementation of Volterra's Certificate resources
func resourceVolterraCertificate() *schema.Resource {
	return &schema.Resource{
		Create: resourceVolterraCertificateCreate,
		Read:   resourceVolterraCertificateRead,
		Update: resourceVolterraCertificateUpdate,
		Delete: resourceVolterraCertificateDelete,

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

			"certificate_chain": {

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

			"certificate_url": {
				Type:     schema.TypeString,
				Required: true,
			},

			"custom_hash_algorithms": {

				Type:     schema.TypeSet,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"hash_algorithms": {

							Type: schema.TypeList,

							Required: true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
					},
				},
			},

			"disable_ocsp_stapling": {

				Type:     schema.TypeBool,
				Optional: true,
			},

			"use_system_defaults": {

				Type:     schema.TypeBool,
				Optional: true,
			},

			"private_key": {

				Type:     schema.TypeSet,
				Required: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"blindfold_secret_info_internal": {

							Type:     schema.TypeSet,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"decryption_provider": {
										Type:     schema.TypeString,
										Optional: true,
									},

									"location": {
										Type:     schema.TypeString,
										Required: true,
									},

									"store_provider": {
										Type:     schema.TypeString,
										Optional: true,
									},
								},
							},
						},

						"secret_encoding_type": {
							Type:     schema.TypeString,
							Optional: true,
						},

						"blindfold_secret_info": {

							Type:     schema.TypeSet,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"decryption_provider": {
										Type:     schema.TypeString,
										Optional: true,
									},

									"location": {
										Type:     schema.TypeString,
										Required: true,
									},

									"store_provider": {
										Type:     schema.TypeString,
										Optional: true,
									},
								},
							},
						},

						"clear_secret_info": {

							Type:     schema.TypeSet,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"provider": {
										Type:     schema.TypeString,
										Optional: true,
									},

									"url": {
										Type:     schema.TypeString,
										Required: true,
									},
								},
							},
						},

						"vault_secret_info": {

							Type:     schema.TypeSet,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"key": {
										Type:     schema.TypeString,
										Optional: true,
									},

									"location": {
										Type:     schema.TypeString,
										Required: true,
									},

									"provider": {
										Type:     schema.TypeString,
										Required: true,
									},

									"secret_encoding": {
										Type:     schema.TypeString,
										Optional: true,
									},

									"version": {
										Type:     schema.TypeInt,
										Optional: true,
									},
								},
							},
						},

						"wingman_secret_info": {

							Type:     schema.TypeSet,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"name": {
										Type:     schema.TypeString,
										Required: true,
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

// resourceVolterraCertificateCreate creates Certificate resource
func resourceVolterraCertificateCreate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*APIClient)

	createMeta := &ves_io_schema.ObjectCreateMetaType{}
	createSpec := &ves_io_schema_certificate.CreateSpecType{}
	createReq := &ves_io_schema_certificate.CreateRequest{
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

	//certificate_chain
	if v, ok := d.GetOk("certificate_chain"); ok && !isIntfNil(v) {

		sl := v.(*schema.Set).List()
		certificateChainInt := &ves_io_schema_views.ObjectRefType{}
		createSpec.CertificateChain = certificateChainInt

		for _, set := range sl {
			ccMapToStrVal := set.(map[string]interface{})
			if val, ok := ccMapToStrVal["name"]; ok && !isIntfNil(v) {
				certificateChainInt.Name = val.(string)
			}
			if val, ok := ccMapToStrVal["namespace"]; ok && !isIntfNil(v) {
				certificateChainInt.Namespace = val.(string)
			}

			if val, ok := ccMapToStrVal["tenant"]; ok && !isIntfNil(v) {
				certificateChainInt.Tenant = val.(string)
			}
		}

	}

	//certificate_url
	if v, ok := d.GetOk("certificate_url"); ok && !isIntfNil(v) {

		createSpec.CertificateUrl =
			v.(string)

	}

	//ocsp_stapling_choice

	ocspStaplingChoiceTypeFound := false

	if v, ok := d.GetOk("custom_hash_algorithms"); ok && !ocspStaplingChoiceTypeFound {

		ocspStaplingChoiceTypeFound = true
		ocspStaplingChoiceInt := &ves_io_schema_certificate.CreateSpecType_CustomHashAlgorithms{}
		ocspStaplingChoiceInt.CustomHashAlgorithms = &ves_io_schema.HashAlgorithms{}
		createSpec.OcspStaplingChoice = ocspStaplingChoiceInt

		sl := v.(*schema.Set).List()
		for _, set := range sl {
			cs := set.(map[string]interface{})

			if v, ok := cs["hash_algorithms"]; ok && !isIntfNil(v) {

				hash_algorithmsList := []ves_io_schema.HashAlgorithm{}
				for _, j := range v.([]interface{}) {
					hash_algorithmsList = append(hash_algorithmsList, ves_io_schema.HashAlgorithm(ves_io_schema.HashAlgorithm_value[j.(string)]))
				}
				ocspStaplingChoiceInt.CustomHashAlgorithms.HashAlgorithms = hash_algorithmsList

			}

		}

	}

	if v, ok := d.GetOk("disable_ocsp_stapling"); ok && !ocspStaplingChoiceTypeFound {

		ocspStaplingChoiceTypeFound = true

		if v.(bool) {
			ocspStaplingChoiceInt := &ves_io_schema_certificate.CreateSpecType_DisableOcspStapling{}
			ocspStaplingChoiceInt.DisableOcspStapling = &ves_io_schema.Empty{}
			createSpec.OcspStaplingChoice = ocspStaplingChoiceInt
		}

	}

	if v, ok := d.GetOk("use_system_defaults"); ok && !ocspStaplingChoiceTypeFound {

		ocspStaplingChoiceTypeFound = true

		if v.(bool) {
			ocspStaplingChoiceInt := &ves_io_schema_certificate.CreateSpecType_UseSystemDefaults{}
			ocspStaplingChoiceInt.UseSystemDefaults = &ves_io_schema.Empty{}
			createSpec.OcspStaplingChoice = ocspStaplingChoiceInt
		}

	}

	//private_key
	if v, ok := d.GetOk("private_key"); ok && !isIntfNil(v) {

		sl := v.(*schema.Set).List()
		privateKey := &ves_io_schema.SecretType{}
		createSpec.PrivateKey = privateKey
		for _, set := range sl {
			privateKeyMapStrToI := set.(map[string]interface{})

			if v, ok := privateKeyMapStrToI["blindfold_secret_info_internal"]; ok && !isIntfNil(v) {

				sl := v.(*schema.Set).List()
				blindfoldSecretInfoInternal := &ves_io_schema.BlindfoldSecretInfoType{}
				privateKey.BlindfoldSecretInfoInternal = blindfoldSecretInfoInternal
				for _, set := range sl {
					blindfoldSecretInfoInternalMapStrToI := set.(map[string]interface{})

					if w, ok := blindfoldSecretInfoInternalMapStrToI["decryption_provider"]; ok && !isIntfNil(w) {
						blindfoldSecretInfoInternal.DecryptionProvider = w.(string)
					}

					if w, ok := blindfoldSecretInfoInternalMapStrToI["location"]; ok && !isIntfNil(w) {
						blindfoldSecretInfoInternal.Location = w.(string)
					}

					if w, ok := blindfoldSecretInfoInternalMapStrToI["store_provider"]; ok && !isIntfNil(w) {
						blindfoldSecretInfoInternal.StoreProvider = w.(string)
					}

				}

			}

			if v, ok := privateKeyMapStrToI["secret_encoding_type"]; ok && !isIntfNil(v) {

				privateKey.SecretEncodingType = ves_io_schema.SecretEncodingType(ves_io_schema.SecretEncodingType_value[v.(string)])

			}

			secretInfoOneofTypeFound := false

			if v, ok := privateKeyMapStrToI["blindfold_secret_info"]; ok && !isIntfNil(v) && !secretInfoOneofTypeFound {

				secretInfoOneofTypeFound = true
				secretInfoOneofInt := &ves_io_schema.SecretType_BlindfoldSecretInfo{}
				secretInfoOneofInt.BlindfoldSecretInfo = &ves_io_schema.BlindfoldSecretInfoType{}
				privateKey.SecretInfoOneof = secretInfoOneofInt

				sl := v.(*schema.Set).List()
				for _, set := range sl {
					cs := set.(map[string]interface{})

					if v, ok := cs["decryption_provider"]; ok && !isIntfNil(v) {

						secretInfoOneofInt.BlindfoldSecretInfo.DecryptionProvider = v.(string)

					}

					if v, ok := cs["location"]; ok && !isIntfNil(v) {

						secretInfoOneofInt.BlindfoldSecretInfo.Location = v.(string)

					}

					if v, ok := cs["store_provider"]; ok && !isIntfNil(v) {

						secretInfoOneofInt.BlindfoldSecretInfo.StoreProvider = v.(string)

					}

				}

			}

			if v, ok := privateKeyMapStrToI["clear_secret_info"]; ok && !isIntfNil(v) && !secretInfoOneofTypeFound {

				secretInfoOneofTypeFound = true
				secretInfoOneofInt := &ves_io_schema.SecretType_ClearSecretInfo{}
				secretInfoOneofInt.ClearSecretInfo = &ves_io_schema.ClearSecretInfoType{}
				privateKey.SecretInfoOneof = secretInfoOneofInt

				sl := v.(*schema.Set).List()
				for _, set := range sl {
					cs := set.(map[string]interface{})

					if v, ok := cs["provider"]; ok && !isIntfNil(v) {

						secretInfoOneofInt.ClearSecretInfo.Provider = v.(string)

					}

					if v, ok := cs["url"]; ok && !isIntfNil(v) {

						secretInfoOneofInt.ClearSecretInfo.Url = v.(string)

					}

				}

			}

			if v, ok := privateKeyMapStrToI["vault_secret_info"]; ok && !isIntfNil(v) && !secretInfoOneofTypeFound {

				secretInfoOneofTypeFound = true
				secretInfoOneofInt := &ves_io_schema.SecretType_VaultSecretInfo{}
				secretInfoOneofInt.VaultSecretInfo = &ves_io_schema.VaultSecretInfoType{}
				privateKey.SecretInfoOneof = secretInfoOneofInt

				sl := v.(*schema.Set).List()
				for _, set := range sl {
					cs := set.(map[string]interface{})

					if v, ok := cs["key"]; ok && !isIntfNil(v) {

						secretInfoOneofInt.VaultSecretInfo.Key = v.(string)

					}

					if v, ok := cs["location"]; ok && !isIntfNil(v) {

						secretInfoOneofInt.VaultSecretInfo.Location = v.(string)

					}

					if v, ok := cs["provider"]; ok && !isIntfNil(v) {

						secretInfoOneofInt.VaultSecretInfo.Provider = v.(string)

					}

					if v, ok := cs["secret_encoding"]; ok && !isIntfNil(v) {

						secretInfoOneofInt.VaultSecretInfo.SecretEncoding = ves_io_schema.SecretEncodingType(ves_io_schema.SecretEncodingType_value[v.(string)])

					}

					if v, ok := cs["version"]; ok && !isIntfNil(v) {

						secretInfoOneofInt.VaultSecretInfo.Version = uint32(v.(int))

					}

				}

			}

			if v, ok := privateKeyMapStrToI["wingman_secret_info"]; ok && !isIntfNil(v) && !secretInfoOneofTypeFound {

				secretInfoOneofTypeFound = true
				secretInfoOneofInt := &ves_io_schema.SecretType_WingmanSecretInfo{}
				secretInfoOneofInt.WingmanSecretInfo = &ves_io_schema.WingmanSecretInfoType{}
				privateKey.SecretInfoOneof = secretInfoOneofInt

				sl := v.(*schema.Set).List()
				for _, set := range sl {
					cs := set.(map[string]interface{})

					if v, ok := cs["name"]; ok && !isIntfNil(v) {

						secretInfoOneofInt.WingmanSecretInfo.Name = v.(string)

					}

				}

			}

		}

	}

	log.Printf("[DEBUG] Creating Volterra Certificate object with struct: %+v", createReq)

	createCertificateResp, err := client.CreateObject(context.Background(), ves_io_schema_certificate.ObjectType, createReq)
	if err != nil {
		return fmt.Errorf("error creating Certificate: %s", err)
	}
	d.SetId(createCertificateResp.GetObjSystemMetadata().GetUid())

	return resourceVolterraCertificateRead(d, meta)
}

func resourceVolterraCertificateRead(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*APIClient)
	name := d.Get("name").(string)
	namespace := d.Get("namespace").(string)

	resp, err := client.GetObject(context.Background(), ves_io_schema_certificate.ObjectType, namespace, name)
	if err != nil {
		if strings.Contains(err.Error(), "status code 404") {
			log.Printf("[INFO] Certificate %s no longer exists", d.Id())
			d.SetId("")
			return nil
		}
		return fmt.Errorf("Error finding Volterra Certificate %q: %s", d.Id(), err)
	}
	return setCertificateFields(client, d, resp)
}

func setCertificateFields(client *APIClient, d *schema.ResourceData, resp vesapi.GetObjectResponse) error {
	metadata := resp.GetObjMetadata()

	d.Set("annotations", metadata.GetAnnotations())

	d.Set("description", metadata.GetDescription())

	d.Set("disable", metadata.GetDisable())

	d.Set("labels", metadata.GetLabels())

	d.Set("name", metadata.GetName())

	d.Set("namespace", metadata.GetNamespace())

	return nil
}

// resourceVolterraCertificateUpdate updates Certificate resource
func resourceVolterraCertificateUpdate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*APIClient)

	updateMeta := &ves_io_schema.ObjectReplaceMetaType{}
	updateSpec := &ves_io_schema_certificate.ReplaceSpecType{}
	updateReq := &ves_io_schema_certificate.ReplaceRequest{
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

	if v, ok := d.GetOk("certificate_chain"); ok && !isIntfNil(v) {

		sl := v.(*schema.Set).List()
		certificateChainInt := &ves_io_schema_views.ObjectRefType{}
		updateSpec.CertificateChain = certificateChainInt

		for _, set := range sl {
			ccMapToStrVal := set.(map[string]interface{})
			if val, ok := ccMapToStrVal["name"]; ok && !isIntfNil(v) {
				certificateChainInt.Name = val.(string)
			}
			if val, ok := ccMapToStrVal["namespace"]; ok && !isIntfNil(v) {
				certificateChainInt.Namespace = val.(string)
			}

			if val, ok := ccMapToStrVal["tenant"]; ok && !isIntfNil(v) {
				certificateChainInt.Tenant = val.(string)
			}
		}

	}

	if v, ok := d.GetOk("certificate_url"); ok && !isIntfNil(v) {

		updateSpec.CertificateUrl =
			v.(string)

	}

	ocspStaplingChoiceTypeFound := false

	if v, ok := d.GetOk("custom_hash_algorithms"); ok && !ocspStaplingChoiceTypeFound {

		ocspStaplingChoiceTypeFound = true
		ocspStaplingChoiceInt := &ves_io_schema_certificate.ReplaceSpecType_CustomHashAlgorithms{}
		ocspStaplingChoiceInt.CustomHashAlgorithms = &ves_io_schema.HashAlgorithms{}
		updateSpec.OcspStaplingChoice = ocspStaplingChoiceInt

		sl := v.(*schema.Set).List()
		for _, set := range sl {
			cs := set.(map[string]interface{})

			if v, ok := cs["hash_algorithms"]; ok && !isIntfNil(v) {

				hash_algorithmsList := []ves_io_schema.HashAlgorithm{}
				for _, j := range v.([]interface{}) {
					hash_algorithmsList = append(hash_algorithmsList, ves_io_schema.HashAlgorithm(ves_io_schema.HashAlgorithm_value[j.(string)]))
				}
				ocspStaplingChoiceInt.CustomHashAlgorithms.HashAlgorithms = hash_algorithmsList

			}

		}

	}

	if v, ok := d.GetOk("disable_ocsp_stapling"); ok && !ocspStaplingChoiceTypeFound {

		ocspStaplingChoiceTypeFound = true

		if v.(bool) {
			ocspStaplingChoiceInt := &ves_io_schema_certificate.ReplaceSpecType_DisableOcspStapling{}
			ocspStaplingChoiceInt.DisableOcspStapling = &ves_io_schema.Empty{}
			updateSpec.OcspStaplingChoice = ocspStaplingChoiceInt
		}

	}

	if v, ok := d.GetOk("use_system_defaults"); ok && !ocspStaplingChoiceTypeFound {

		ocspStaplingChoiceTypeFound = true

		if v.(bool) {
			ocspStaplingChoiceInt := &ves_io_schema_certificate.ReplaceSpecType_UseSystemDefaults{}
			ocspStaplingChoiceInt.UseSystemDefaults = &ves_io_schema.Empty{}
			updateSpec.OcspStaplingChoice = ocspStaplingChoiceInt
		}

	}

	if v, ok := d.GetOk("private_key"); ok && !isIntfNil(v) {

		sl := v.(*schema.Set).List()
		privateKey := &ves_io_schema.SecretType{}
		updateSpec.PrivateKey = privateKey
		for _, set := range sl {
			privateKeyMapStrToI := set.(map[string]interface{})

			if v, ok := privateKeyMapStrToI["blindfold_secret_info_internal"]; ok && !isIntfNil(v) {

				sl := v.(*schema.Set).List()
				blindfoldSecretInfoInternal := &ves_io_schema.BlindfoldSecretInfoType{}
				privateKey.BlindfoldSecretInfoInternal = blindfoldSecretInfoInternal
				for _, set := range sl {
					blindfoldSecretInfoInternalMapStrToI := set.(map[string]interface{})

					if w, ok := blindfoldSecretInfoInternalMapStrToI["decryption_provider"]; ok && !isIntfNil(w) {
						blindfoldSecretInfoInternal.DecryptionProvider = w.(string)
					}

					if w, ok := blindfoldSecretInfoInternalMapStrToI["location"]; ok && !isIntfNil(w) {
						blindfoldSecretInfoInternal.Location = w.(string)
					}

					if w, ok := blindfoldSecretInfoInternalMapStrToI["store_provider"]; ok && !isIntfNil(w) {
						blindfoldSecretInfoInternal.StoreProvider = w.(string)
					}

				}

			}

			if v, ok := privateKeyMapStrToI["secret_encoding_type"]; ok && !isIntfNil(v) {

				privateKey.SecretEncodingType = ves_io_schema.SecretEncodingType(ves_io_schema.SecretEncodingType_value[v.(string)])

			}

			secretInfoOneofTypeFound := false

			if v, ok := privateKeyMapStrToI["blindfold_secret_info"]; ok && !isIntfNil(v) && !secretInfoOneofTypeFound {

				secretInfoOneofTypeFound = true
				secretInfoOneofInt := &ves_io_schema.SecretType_BlindfoldSecretInfo{}
				secretInfoOneofInt.BlindfoldSecretInfo = &ves_io_schema.BlindfoldSecretInfoType{}
				privateKey.SecretInfoOneof = secretInfoOneofInt

				sl := v.(*schema.Set).List()
				for _, set := range sl {
					cs := set.(map[string]interface{})

					if v, ok := cs["decryption_provider"]; ok && !isIntfNil(v) {

						secretInfoOneofInt.BlindfoldSecretInfo.DecryptionProvider = v.(string)

					}

					if v, ok := cs["location"]; ok && !isIntfNil(v) {

						secretInfoOneofInt.BlindfoldSecretInfo.Location = v.(string)

					}

					if v, ok := cs["store_provider"]; ok && !isIntfNil(v) {

						secretInfoOneofInt.BlindfoldSecretInfo.StoreProvider = v.(string)

					}

				}

			}

			if v, ok := privateKeyMapStrToI["clear_secret_info"]; ok && !isIntfNil(v) && !secretInfoOneofTypeFound {

				secretInfoOneofTypeFound = true
				secretInfoOneofInt := &ves_io_schema.SecretType_ClearSecretInfo{}
				secretInfoOneofInt.ClearSecretInfo = &ves_io_schema.ClearSecretInfoType{}
				privateKey.SecretInfoOneof = secretInfoOneofInt

				sl := v.(*schema.Set).List()
				for _, set := range sl {
					cs := set.(map[string]interface{})

					if v, ok := cs["provider"]; ok && !isIntfNil(v) {

						secretInfoOneofInt.ClearSecretInfo.Provider = v.(string)

					}

					if v, ok := cs["url"]; ok && !isIntfNil(v) {

						secretInfoOneofInt.ClearSecretInfo.Url = v.(string)

					}

				}

			}

			if v, ok := privateKeyMapStrToI["vault_secret_info"]; ok && !isIntfNil(v) && !secretInfoOneofTypeFound {

				secretInfoOneofTypeFound = true
				secretInfoOneofInt := &ves_io_schema.SecretType_VaultSecretInfo{}
				secretInfoOneofInt.VaultSecretInfo = &ves_io_schema.VaultSecretInfoType{}
				privateKey.SecretInfoOneof = secretInfoOneofInt

				sl := v.(*schema.Set).List()
				for _, set := range sl {
					cs := set.(map[string]interface{})

					if v, ok := cs["key"]; ok && !isIntfNil(v) {

						secretInfoOneofInt.VaultSecretInfo.Key = v.(string)

					}

					if v, ok := cs["location"]; ok && !isIntfNil(v) {

						secretInfoOneofInt.VaultSecretInfo.Location = v.(string)

					}

					if v, ok := cs["provider"]; ok && !isIntfNil(v) {

						secretInfoOneofInt.VaultSecretInfo.Provider = v.(string)

					}

					if v, ok := cs["secret_encoding"]; ok && !isIntfNil(v) {

						secretInfoOneofInt.VaultSecretInfo.SecretEncoding = ves_io_schema.SecretEncodingType(ves_io_schema.SecretEncodingType_value[v.(string)])

					}

					if v, ok := cs["version"]; ok && !isIntfNil(v) {

						secretInfoOneofInt.VaultSecretInfo.Version = uint32(v.(int))

					}

				}

			}

			if v, ok := privateKeyMapStrToI["wingman_secret_info"]; ok && !isIntfNil(v) && !secretInfoOneofTypeFound {

				secretInfoOneofTypeFound = true
				secretInfoOneofInt := &ves_io_schema.SecretType_WingmanSecretInfo{}
				secretInfoOneofInt.WingmanSecretInfo = &ves_io_schema.WingmanSecretInfoType{}
				privateKey.SecretInfoOneof = secretInfoOneofInt

				sl := v.(*schema.Set).List()
				for _, set := range sl {
					cs := set.(map[string]interface{})

					if v, ok := cs["name"]; ok && !isIntfNil(v) {

						secretInfoOneofInt.WingmanSecretInfo.Name = v.(string)

					}

				}

			}

		}

	}

	log.Printf("[DEBUG] Updating Volterra Certificate obj with struct: %+v", updateReq)

	err := client.ReplaceObject(context.Background(), ves_io_schema_certificate.ObjectType, updateReq)
	if err != nil {
		return fmt.Errorf("error updating Certificate: %s", err)
	}

	return resourceVolterraCertificateRead(d, meta)
}

func resourceVolterraCertificateDelete(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*APIClient)
	name := d.Get("name").(string)
	namespace := d.Get("namespace").(string)

	_, err := client.GetObject(context.Background(), ves_io_schema_certificate.ObjectType, namespace, name)
	if err != nil {
		if strings.Contains(err.Error(), "status code 404") {
			log.Printf("[INFO] Certificate %s no longer exists", d.Id())
			d.SetId("")
			return nil
		}
		return fmt.Errorf("Error finding Volterra Certificate before deleting %q: %s", d.Id(), err)
	}

	log.Printf("[DEBUG] Deleting Volterra Certificate obj with name %+v in namespace %+v", name, namespace)
	return client.DeleteObject(context.Background(), ves_io_schema_certificate.ObjectType, namespace, name)
}
