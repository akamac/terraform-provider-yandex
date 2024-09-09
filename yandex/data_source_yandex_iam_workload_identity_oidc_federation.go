package yandex

import (
	"context"
	"fmt"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"

	"github.com/yandex-cloud/go-genproto/yandex/cloud/iam/v1/workload/oidc"
	"github.com/yandex-cloud/go-sdk/sdkresolvers"
)

func dataSourceYandexIAMWorkloadIdentityOidcFederation() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceYandexIAMWorkloadIdentityOidcFederationRead,

		SchemaVersion: 1,

		Schema: map[string]*schema.Schema{
			"federation_id": {
				Type:          schema.TypeString,
				Computed:      true,
				Optional:      true,
				ValidateFunc:  validation.StringLenBetween(0, 50),
				ConflictsWith: []string{"name"},
			},

			"folder_id": {
				Type:     schema.TypeString,
				Computed: true,
			},

			"name": {
				Type:          schema.TypeString,
				Computed:      true,
				Optional:      true,
				ConflictsWith: []string{"federation_id"},
			},

			"description": {
				Type:     schema.TypeString,
				Computed: true,
			},

			"enabled": {
				Type:     schema.TypeBool,
				Computed: true,
			},

			"audiences": {
				Type: schema.TypeList,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
				Computed: true,
			},

			"issuer": {
				Type:     schema.TypeString,
				Computed: true,
			},

			"jwks_url": {
				Type:     schema.TypeString,
				Computed: true,
			},

			"labels": {
				Type: schema.TypeMap,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
				Set:      schema.HashString,
				Computed: true,
			},

			"created_at": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func dataSourceYandexIAMWorkloadIdentityOidcFederationRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	config := meta.(*Config)

	err := checkOneOf(d, "federation_id", "name")
	if err != nil {
		return diag.FromErr(err)
	}

	federationID := d.Get("federation_id").(string)
	_, federationNameOk := d.GetOk("name")

	if federationNameOk {
		federationID, err = resolveObjectID(ctx, config, d, sdkresolvers.WliFederationResolver)
		if err != nil {
			return diag.FromErr(fmt.Errorf("failed to resolve workload identity federation by name: %v", err))
		}
	}

	req := &oidc.GetFederationRequest{
		FederationId: federationID,
	}

	resp, err := config.sdk.WorkloadOidc().Federation().Get(ctx, req)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId(resp.Id)

	if err := d.Set("federation_id", resp.GetId()); err != nil {
		log.Printf("[ERROR] failed set field federation_id: %s", err)
		return diag.FromErr(err)
	}
	if err := d.Set("folder_id", resp.GetFolderId()); err != nil {
		log.Printf("[ERROR] failed set field folder_id: %s", err)
		return diag.FromErr(err)
	}
	if err := d.Set("name", resp.GetName()); err != nil {
		log.Printf("[ERROR] failed set field name: %s", err)
		return diag.FromErr(err)
	}
	if err := d.Set("description", resp.GetDescription()); err != nil {
		log.Printf("[ERROR] failed set field description: %s", err)
		return diag.FromErr(err)
	}
	if err := d.Set("enabled", resp.GetEnabled()); err != nil {
		log.Printf("[ERROR] failed set field enabled: %s", err)
		return diag.FromErr(err)
	}
	if err := d.Set("audiences", resp.GetAudiences()); err != nil {
		log.Printf("[ERROR] failed set field audiences: %s", err)
		return diag.FromErr(err)
	}
	if err := d.Set("issuer", resp.GetIssuer()); err != nil {
		log.Printf("[ERROR] failed set field issuer: %s", err)
		return diag.FromErr(err)
	}
	if err := d.Set("jwks_url", resp.GetJwksUrl()); err != nil {
		log.Printf("[ERROR] failed set field jwks_url: %s", err)
		return diag.FromErr(err)
	}
	if err := d.Set("labels", resp.GetLabels()); err != nil {
		log.Printf("[ERROR] failed set field labels: %s", err)
		return diag.FromErr(err)
	}
	if err := d.Set("created_at", getTimestamp(resp.GetCreatedAt())); err != nil {
		log.Printf("[ERROR] failed set field created_at: %s", err)
		return diag.FromErr(err)
	}

	return nil
}