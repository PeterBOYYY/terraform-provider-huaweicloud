// ----------------------------------------------------------------------------
//
//     ***     AUTO GENERATED CODE    ***    AUTO GENERATED CODE     ***
//
// ----------------------------------------------------------------------------
//
//     This file is automatically generated by Magic Modules and manual
//     changes will be clobbered when the file is regenerated.
//
//     Please read more about how to change this file at
//     https://www.github.com/huaweicloud/magic-modules
//
// ----------------------------------------------------------------------------

package huaweicloud

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
	"github.com/huaweicloud/golangsdk"
)

func TestAccRdsInstanceV3_basic(t *testing.T) {
	name := acctest.RandString(10)
	resourceName := "huaweicloud_rds_instance.instance"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckRdsInstanceV3Destroy,
		Steps: []resource.TestStep{
			{
				Config: testAccRdsInstanceV3_basic(name),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRdsInstanceV3Exists(),
					resource.TestCheckResourceAttr(resourceName, "name", fmt.Sprintf("terraform_test_rds_instance%s", name)),
					resource.TestCheckResourceAttr(resourceName, "backup_strategy.0.keep_days", "1"),
					resource.TestCheckResourceAttr(resourceName, "flavor", "rds.pg.c2.large"),
					resource.TestCheckResourceAttr(resourceName, "volume.0.size", "50"),
					resource.TestCheckResourceAttr(resourceName, "tags.key", "value"),
					resource.TestCheckResourceAttr(resourceName, "tags.foo", "bar"),
					resource.TestCheckResourceAttr(resourceName, "time_zone", "UTC+10:00"),
					resource.TestCheckResourceAttr(resourceName, "fixed_ip", "192.168.0.58"),
				),
			},
			{
				ResourceName:      resourceName,
				ImportState:       true,
				ImportStateVerify: true,
				ImportStateVerifyIgnore: []string{
					"db",
				},
			},
			{
				Config: testAccRdsInstanceV3_update(name),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRdsInstanceV3Exists(),
					resource.TestCheckResourceAttr(resourceName, "name", fmt.Sprintf("terraform_test_rds_instance_update%s", name)),
					resource.TestCheckResourceAttr(resourceName, "backup_strategy.0.keep_days", "2"),
					resource.TestCheckResourceAttr(resourceName, "flavor", "rds.pg.c2.xlarge"),
					resource.TestCheckResourceAttr(resourceName, "volume.0.size", "100"),
					resource.TestCheckResourceAttr(resourceName, "tags.key1", "value"),
					resource.TestCheckResourceAttr(resourceName, "tags.foo", "bar_updated"),
				),
			},
		},
	})
}

func TestAccRdsInstanceV3_withEpsId(t *testing.T) {
	name := acctest.RandString(10)
	resourceName := "huaweicloud_rds_instance.instance"
	resource.ParallelTest(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheckEpsID(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckRdsInstanceV3Destroy,
		Steps: []resource.TestStep{
			{
				Config: testAccRdsInstanceV3_epsId(name),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRdsInstanceV3Exists(),
					resource.TestCheckResourceAttr(resourceName, "enterprise_project_id", HW_ENTERPRISE_PROJECT_ID_TEST),
				),
			},
		},
	})
}

func testAccRdsInstanceV3_base(val string) string {
	return fmt.Sprintf(`
resource "huaweicloud_vpc" "test" {
  name = "vpc-rds-test-%s"
  cidr = "192.168.0.0/16"
}

resource "huaweicloud_vpc_subnet" "test" {
  name          = "subnet-rds-test-%s"
  cidr          = "192.168.0.0/16"
  gateway_ip    = "192.168.0.1"
  primary_dns   = "100.125.1.250"
  secondary_dns = "100.125.21.250"
  vpc_id        = huaweicloud_vpc.test.id
}

resource "huaweicloud_networking_secgroup" "secgroup_1" {
  name = "sg-rds-test-%s"
}
`, val, val, val)
}

func testAccRdsInstanceV3_basic(val string) string {
	return fmt.Sprintf(`
%s

resource "huaweicloud_rds_instance" "instance" {
  name = "terraform_test_rds_instance%s"
  flavor = "rds.pg.c2.large"
  availability_zone = ["%s"]
  security_group_id = huaweicloud_networking_secgroup.secgroup_1.id
  subnet_id = huaweicloud_vpc_subnet.test.id
  vpc_id = huaweicloud_vpc.test.id
  time_zone = "UTC+10:00"
  fixed_ip = "192.168.0.58"

  db {
    password = "Huangwei!120521"
    type = "PostgreSQL"
    version = "10"
    port = "8635"
  }
  volume {
    type = "ULTRAHIGH"
    size = 50
  }
  backup_strategy {
    start_time = "08:00-09:00"
    keep_days = 1
  }

  tags = {
    key = "value"
    foo = "bar"
  }
}
	`, testAccRdsInstanceV3_base(val), val, HW_AVAILABILITY_ZONE)
}

// name, volume.size, backup_strategy, flavor and tags will be updated
func testAccRdsInstanceV3_update(val string) string {
	return fmt.Sprintf(`
%s

resource "huaweicloud_rds_instance" "instance" {
  name = "terraform_test_rds_instance_update%s"
  flavor = "rds.pg.c2.xlarge"
  availability_zone = ["%s"]
  security_group_id = huaweicloud_networking_secgroup.secgroup_1.id
  subnet_id = huaweicloud_vpc_subnet.test.id
  vpc_id = huaweicloud_vpc.test.id
  time_zone = "UTC+10:00"

  db {
    password = "Huangwei!120521"
    type = "PostgreSQL"
    version = "10"
    port = "8635"
  }
  volume {
    type = "ULTRAHIGH"
    size = 100
  }
  backup_strategy {
    start_time = "09:00-10:00"
    keep_days = 2
  }

  tags = {
    key1 = "value"
    foo = "bar_updated"
  }
}
	`, testAccRdsInstanceV3_base(val), val, HW_AVAILABILITY_ZONE)
}

func testAccRdsInstanceV3_epsId(val string) string {
	return fmt.Sprintf(`
resource "huaweicloud_vpc_v1" "test" {
  name = "vpc-rds-test-%s"
  cidr = "192.168.0.0/16"
}
	
resource "huaweicloud_vpc_subnet_v1" "test" {
  name          = "subnet-rds-test-%s"
  cidr          = "192.168.0.0/16"
  gateway_ip    = "192.168.0.1"
  primary_dns   = "100.125.1.250"
  secondary_dns = "100.125.21.250"
  vpc_id        = huaweicloud_vpc_v1.test.id
}

resource "huaweicloud_networking_secgroup_v2" "secgroup_1" {
  name = "sg-rds-test-%s"
}

resource "huaweicloud_rds_instance" "instance" {
  name = "terraform_test_rds_instance%s"
  flavor = "rds.pg.c2.medium"
  availability_zone = ["%s"]
  security_group_id = huaweicloud_networking_secgroup_v2.secgroup_1.id
  subnet_id = huaweicloud_vpc_subnet_v1.test.id
  vpc_id = huaweicloud_vpc_v1.test.id
  enterprise_project_id = "%s"

  db {
    password = "Huangwei!120521"
    type = "PostgreSQL"
    version = "10"
    port = "8635"
  }
  volume {
    type = "ULTRAHIGH"
    size = 50
  }
  backup_strategy {
    start_time = "08:00-09:00"
    keep_days = 1
  }
}
	`, val, val, val, val, HW_AVAILABILITY_ZONE, HW_ENTERPRISE_PROJECT_ID_TEST)
}

func testAccCheckRdsInstanceV3Destroy(s *terraform.State) error {
	config := testAccProvider.Meta().(*Config)
	client, err := config.RdsV3Client(HW_REGION_NAME)
	if err != nil {
		return fmt.Errorf("Error creating sdk client, err=%s", err)
	}

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "huaweicloud_rds_instance" {
			continue
		}

		v, err := fetchRdsInstanceV3ByListOnTest(rs, client)
		if err != nil {
			return err
		}
		if v != nil {
			return fmt.Errorf("huaweicloud rds instance still exists")
		}
	}

	return nil
}

func testAccCheckRdsInstanceV3Exists() resource.TestCheckFunc {
	return func(s *terraform.State) error {
		config := testAccProvider.Meta().(*Config)
		client, err := config.RdsV3Client(HW_REGION_NAME)
		if err != nil {
			return fmt.Errorf("Error creating sdk client, err=%s", err)
		}

		rs, ok := s.RootModule().Resources["huaweicloud_rds_instance.instance"]
		if !ok {
			return fmt.Errorf("Error checking huaweicloud_rds_instance.instance exist, err=not found this resource")
		}

		v, err := fetchRdsInstanceV3ByListOnTest(rs, client)
		if err != nil {
			return fmt.Errorf("Error checking huaweicloud_rds_instance.instance exist, err=%s", err)
		}
		if v == nil {
			return fmt.Errorf("huaweicloud rds instance is not exist")
		}
		return nil
	}
}

func fetchRdsInstanceV3ByListOnTest(rs *terraform.ResourceState,
	client *golangsdk.ServiceClient) (interface{}, error) {

	identity := map[string]interface{}{"id": rs.Primary.ID}

	queryLink := "?id=" + identity["id"].(string)

	link := client.ServiceURL("instances") + queryLink

	return findRdsInstanceV3ByList(client, link, identity)
}
