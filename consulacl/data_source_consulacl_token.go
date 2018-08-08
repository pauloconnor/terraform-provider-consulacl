package consulacl

import (
	consul "github.com/hashicorp/consul/api"
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/hashicorp/terraform/helper/validation"
)

func dataSourceConsulACLToken() *schema.Resource {
	var allScopes []string
	allScopes = append(allScopes, prefixedScopes...)
	allScopes = append(allScopes, singletonScopes...)
	return &schema.Resource{
		Read: dataSourceConsulACLRead,

		Schema: map[string]*schema.Schema{
			FieldName: {
				Type:     schema.TypeString,
				Required: true,
			},

			FieldRule: {
				Type:     schema.TypeSet,
				Optional: true,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						FieldScope: {
							Type: schema.TypeString,
							// it's required but we have to enforce otherwise due to a bug in terraform
							// when injecting rules as
							// rule = ["${data.null_data_source.policy.*.outputs}"]
							Computed:     true,
							Optional:     true,
							ValidateFunc: validation.StringInSlice(allScopes, true),
						},
						FieldPrefix: {
							Type:     schema.TypeString,
							Computed: true,
							Optional: true,
						},
						FieldPolicy: {
							Type: schema.TypeString,
							// it's required but we have to enforce otherwise due to a bug in terraform
							// when injecting rules as
							// rule = ["${data.null_data_source.policy.*.outputs}"]
							Computed: true,
							Optional: true,
						},
					},
				},
			},

			FieldToken: {
				Type:      schema.TypeString,
				Computed:  true,
				Optional:  true,
				Sensitive: true,
			},

			FieldType: {
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: validation.StringInSlice([]string{"client", "management"}, true),
			},

			FieldInherits: &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
		},
	}
}

func dataSourceConsulACLRead(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*consul.Client)

	acl, _, err := client.ACL().Info(d.Get(FieldToken).(string), nil)
	if err != nil {
		return err
	}

	if acl == nil {
		d.SetId("")
		return nil
	}

	rules, err := decodeRules(acl.Rules)
	if err != nil {
		return err
	}

	d.SetId(getSHA256(acl.ID))
	d.Set(FieldToken, acl.ID)
	d.Set(FieldName, acl.Name)
	d.Set(FieldType, acl.Type)
	d.Set(FieldRule, rules)

	return nil
}
