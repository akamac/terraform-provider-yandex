package project

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/hashicorp/go-multierror"
	"github.com/hashicorp/terraform-plugin-testing/helper/acctest"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/yandex-cloud/go-genproto/yandex/cloud/datasphere/v2"
	provider_config "github.com/yandex-cloud/terraform-provider-yandex/yandex-framework/provider/config"
	"github.com/yandex-cloud/terraform-provider-yandex/yandex-framework/test"
	dataspheretest "github.com/yandex-cloud/terraform-provider-yandex/yandex-framework/test/datasphere"
	"github.com/yandex-cloud/terraform-provider-yandex/yandex-framework/validate"
	"google.golang.org/grpc/codes"
)

func init() {
	resource.AddTestSweepers("yandex_datasphere_project", &resource.Sweeper{
		Name:         "yandex_datasphere_project",
		F:            testSweepProject,
		Dependencies: []string{},
	})
}

func testSweepProject(_ string) error {
	conf, err := test.ConfigForSweepers()
	if err != nil {
		return fmt.Errorf("error getting client: %w", err)
	}

	// get all communities for user
	communities, err := getAllCommunityIDs(conf)
	if err != nil {
		return fmt.Errorf("get communities: %w", err)
	}

	result := &multierror.Error{}

	for _, communityID := range communities {
		it := conf.SDK.Datasphere().Project().ProjectIterator(
			context.Background(),
			&datasphere.ListProjectsRequest{
				CommunityId: communityID,
				OwnedById:   test.AccTestsUser,
			},
		)

		for it.Next() {
			projectId := it.Value().GetId()
			if !test.IsTestResourceName(it.Value().GetName()) {
				continue
			}
			if !sweepProject(conf, projectId) {
				result = multierror.Append(
					result,
					fmt.Errorf("failed to sweep project id %q", projectId),
				)
			}
		}

		if err := it.Error(); err != nil {
			result = multierror.Append(
				result,
				fmt.Errorf("iterator error: %w", err),
			)
		}
	}

	return result.ErrorOrNil()
}

func getAllCommunityIDs(conf *provider_config.Config) ([]string, error) {
	var (
		it = conf.SDK.Datasphere().Community().CommunityIterator(
			context.Background(),
			&datasphere.ListCommunitiesRequest{OrganizationId: test.GetExampleOrganizationID()},
		)
		ids []string
	)

	for it.Next() {
		if !test.IsTestResourceName(it.Value().GetName()) {
			continue
		}
		ids = append(ids, it.Value().GetId())
	}

	if err := it.Error(); err != nil {
		if validate.IsStatusWithCode(err, codes.NotFound) {
			return nil, nil
		}
		return nil, fmt.Errorf("iterator err: %w", err)
	}

	return ids, nil
}

func sweepProject(conf *provider_config.Config, cloudId string) bool {
	return test.SweepWithRetry(sweepProjectOnce, conf, "yandex_datasphere_project", cloudId)
}

func sweepProjectOnce(conf *provider_config.Config, id string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Minute)

	defer cancel()

	op, err := conf.SDK.Datasphere().Project().Delete(ctx, &datasphere.DeleteProjectRequest{
		ProjectId: id,
	})
	return test.HandleSweepOperation(ctx, conf, op, err)
}

func TestAccDatasphereProjectResource_basic(t *testing.T) {
	var (
		communityName = test.ResourceName(63)

		projectName = test.ResourceName(63)
		projectDesc = acctest.RandStringFromCharSet(256, acctest.CharSetAlpha)
		labelKey    = acctest.RandStringFromCharSet(63, acctest.CharSetAlpha)
		labelValue  = acctest.RandStringFromCharSet(63, acctest.CharSetAlphaNum)
	)

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { test.AccPreCheck(t) },
		ProtoV6ProviderFactories: test.AccProviderFactories,
		CheckDestroy:             dataspheretest.AccCheckProjectDestroy,
		Steps: []resource.TestStep{
			basicDatasphereProjectTestStep(communityName, projectName, projectDesc, labelKey, labelValue),
			datasphereProjectImportTestStep(),
		},
	})
}

func TestAccDatasphereProjectResource_fullData(t *testing.T) {
	var (
		communityName = test.ResourceName(63)

		projectName = test.ResourceName(63)
		projectDesc = acctest.RandStringFromCharSet(256, acctest.CharSetAlpha)
		saName      = acctest.RandStringFromCharSet(63, acctest.CharSetAlpha)
	)

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { test.AccPreCheck(t) },
		ProtoV6ProviderFactories: test.AccProviderFactories,
		CheckDestroy:             dataspheretest.AccCheckProjectDestroy,
		Steps: []resource.TestStep{
			{
				Config: testDatasphereProjectFullConfig(communityName, projectName, projectDesc, saName),
				Check: resource.ComposeTestCheckFunc(
					dataspheretest.ProjectExists(dataspheretest.ProjectResourceName),
					resource.TestCheckResourceAttr(dataspheretest.ProjectResourceName, "name", projectName),
					resource.TestCheckResourceAttr(dataspheretest.ProjectResourceName, "description", projectDesc),
					resource.TestCheckResourceAttrSet(dataspheretest.ProjectResourceName, "created_at"),
					resource.TestCheckResourceAttrSet(dataspheretest.ProjectResourceName, "created_by"),
					resource.TestCheckResourceAttr(dataspheretest.ProjectResourceName, "limits.max_units_per_hour", "10"),
					resource.TestCheckResourceAttr(dataspheretest.ProjectResourceName, "limits.max_units_per_execution", "10"),
					resource.TestCheckResourceAttr(dataspheretest.ProjectResourceName, "limits.balance", "10"),
					resource.TestCheckResourceAttr(dataspheretest.ProjectResourceName, "settings.commit_mode", "AUTO"),
					resource.TestCheckResourceAttr(dataspheretest.ProjectResourceName, "settings.ide", "JUPYTER_LAB"),
					resource.TestCheckResourceAttr(dataspheretest.ProjectResourceName, "settings.stale_exec_timeout_mode", "ONE_HOUR"),
					test.AccCheckCreatedAtAttr(dataspheretest.ProjectResourceName),
				),
			},
			datasphereProjectImportTestStep(),
		},
	})
}

func TestAccDatasphereProjectResource_update(t *testing.T) {
	var (
		communityName = test.ResourceName(63)

		projectName = test.ResourceName(63)
		projectDesc = acctest.RandStringFromCharSet(256, acctest.CharSetAlpha)
		labelKey    = acctest.RandStringFromCharSet(63, acctest.CharSetAlpha)
		labelValue  = acctest.RandStringFromCharSet(63, acctest.CharSetAlphaNum)

		projectNameUpdated = test.ResourceName(63)
		projectDescUpdated = acctest.RandStringFromCharSet(256, acctest.CharSetAlpha)
		labelKeyUpdated    = acctest.RandStringFromCharSet(63, acctest.CharSetAlpha)
		labelValueUpdated  = acctest.RandStringFromCharSet(63, acctest.CharSetAlphaNum)
	)

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { test.AccPreCheck(t) },
		ProtoV6ProviderFactories: test.AccProviderFactories,
		CheckDestroy:             dataspheretest.AccCheckProjectDestroy,
		Steps: []resource.TestStep{
			basicDatasphereProjectTestStep(communityName, projectName, projectDesc, labelKey, labelValue),
			datasphereProjectImportTestStep(),
			basicDatasphereProjectTestStep(communityName, projectNameUpdated, projectDescUpdated, labelKeyUpdated, labelValueUpdated),
			datasphereProjectImportTestStep(),
		},
	})
}

func datasphereProjectImportTestStep() resource.TestStep {
	return resource.TestStep{
		ResourceName:      dataspheretest.ProjectResourceName,
		ImportState:       true,
		ImportStateVerify: true,
	}
}

func basicDatasphereProjectTestStep(communityName, projectName, projectDesc, labelKey, labelValue string) resource.TestStep {
	return resource.TestStep{
		Config: testDatasphereProjectBasicConfig(communityName, projectName, projectDesc, labelKey, labelValue),
		Check: resource.ComposeTestCheckFunc(
			dataspheretest.ProjectExists(dataspheretest.ProjectResourceName),
			resource.TestCheckResourceAttr(dataspheretest.ProjectResourceName, "name", projectName),
			resource.TestCheckResourceAttr(dataspheretest.ProjectResourceName, "description", projectDesc),
			resource.TestCheckResourceAttr(dataspheretest.ProjectResourceName, fmt.Sprintf("labels.%s", labelKey), labelValue),
			resource.TestCheckResourceAttrSet(dataspheretest.ProjectResourceName, "created_at"),
			resource.TestCheckResourceAttrSet(dataspheretest.ProjectResourceName, "created_by"),
			test.AccCheckCreatedAtAttr(dataspheretest.ProjectResourceName),
		),
	}
}

func testDatasphereProjectBasicConfig(communityName, projectName, desc, labelKey, labelValue string) string {
	return fmt.Sprintf(`
resource "yandex_datasphere_community" "test-community" {
  name = "%s"
  organization_id = "%s"
  billing_account_id = "%s"
}

resource "yandex_datasphere_project" "test-project" {
  name = "%s"
  description = "%s"
  labels = {
    "%s": "%s"
  }
  community_id = yandex_datasphere_community.test-community.id
}
`, communityName, test.GetExampleOrganizationID(), test.GetBillingAccountId(), projectName, desc, labelKey, labelValue)
}

func testDatasphereProjectFullConfig(communityName, projectName, desc, saName string) string {
	return fmt.Sprintf(`
resource "yandex_datasphere_community" "test-community" {
  name = "%s"
  organization_id = "%s"
  billing_account_id = "%s"
}

resource "yandex_vpc_network" "test-network" {}

resource "yandex_vpc_subnet" "test-subnet" {
  zone           = "ru-central1-a"
  network_id     = yandex_vpc_network.test-network.id
  v4_cidr_blocks = ["192.168.0.0/24"]
}

resource "yandex_vpc_security_group" "test-security-group" {
  network_id = yandex_vpc_network.test-network.id

  ingress {
    protocol       = "TCP"
    description    = "healthchecks"
    port           = 30080
    v4_cidr_blocks = ["198.18.235.0/24", "198.18.248.0/24"]
  }
}

resource "yandex_iam_service_account" "test-account" {
  name        = "%s"
  description = "tf-test"
}

resource "yandex_datasphere_project" "test-project" {
  name = "%s"
  description = "%s"

  labels = {
    test-label: "test-label-value"
  }

  community_id = yandex_datasphere_community.test-community.id
  
  limits = {
	max_units_per_hour = 10
    max_units_per_execution = 10
	balance = 10
  }

  settings = {
	service_account_id = yandex_iam_service_account.test-account.id
 	subnet_id = yandex_vpc_subnet.test-subnet.id
	commit_mode = "AUTO"
	security_group_ids = [yandex_vpc_security_group.test-security-group.id]
	ide = "JUPYTER_LAB"
	default_folder_id = "%s"
	stale_exec_timeout_mode = "ONE_HOUR"
  }
}

resource "yandex_resourcemanager_folder_iam_member" "test_account" {
  folder_id   = "%s"
  member      = "serviceAccount:${yandex_iam_service_account.test-account.id}"
  role        = "editor"
}
`, communityName, test.GetExampleOrganizationID(), test.GetBillingAccountId(), saName, projectName, desc, test.GetExampleFolderID(), test.GetExampleFolderID())
}
