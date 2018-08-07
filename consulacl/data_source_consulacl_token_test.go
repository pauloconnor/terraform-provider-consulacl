package consulacl

/*
import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"
)

func TestAccDataSourceConsulACL_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccDataSourceConsulACLConfig,
				Check: resource.ComposeTestCheckFunc(
					testAccDataSourceConsulACLCheck("data.consulacl.by_token"),
					testAccDataSourceConsulACLCheck("data.consulacl.by_name"),
				),
			},
		},
	})
}

func testAccDataSourceConsulACLCheck(name string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[name]
		if !ok {
			return fmt.Errorf("root module has no resource called %s", name)
		}

		eipRs, ok := s.RootModule().Resources["aws_eip.test"]
		if !ok {
			return fmt.Errorf("can't find aws_eip.test in state")
		}

		attr := rs.Primary.Attributes

		if attr["id"] != eipRs.Primary.Attributes["id"] {
			return fmt.Errorf(
				"id is %s; want %s",
				attr["id"],
				eipRs.Primary.Attributes["id"],
			)
		}

		if attr["public_ip"] != eipRs.Primary.Attributes["public_ip"] {
			return fmt.Errorf(
				"public_ip is %s; want %s",
				attr["public_ip"],
				eipRs.Primary.Attributes["public_ip"],
			)
		}

		return nil
	}
}

const testAccDataSourceConsulACLConfig = `
provider "aws" {
  region = "us-west-2"
}

resource "consulacl_token" "wrong1" {}
resource "consulacl_token" "test" {}
resource "consulacl_token" "wrong2" {}

data "consulacl_token" "by_token" {
  tokemn = "${consulacl_token.test.token}"
}

data "consulacl_token" "by_name" {
  name = "${consulacl_token.test.name}"
}
`*/
