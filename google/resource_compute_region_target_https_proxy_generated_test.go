// ----------------------------------------------------------------------------
//
//     ***     AUTO GENERATED CODE    ***    AUTO GENERATED CODE     ***
//
// ----------------------------------------------------------------------------
//
//     This file is automatically generated by Magic Modules and manual
//     changes will be clobbered when the file is regenerated.
//
//     Please read more about how to change this file in
//     .github/CONTRIBUTING.md.
//
// ----------------------------------------------------------------------------

package google

import (
	"fmt"
	"strings"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

func TestAccComputeRegionTargetHttpsProxy_regionTargetHttpsProxyBasicExample(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix": randString(t, 10),
	}

	vcrTest(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckComputeRegionTargetHttpsProxyDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccComputeRegionTargetHttpsProxy_regionTargetHttpsProxyBasicExample(context),
			},
			{
				ResourceName:      "google_compute_region_target_https_proxy.default",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func testAccComputeRegionTargetHttpsProxy_regionTargetHttpsProxyBasicExample(context map[string]interface{}) string {
	return Nprintf(`
resource "google_compute_region_target_https_proxy" "default" {
  region           = "us-central1"
  name             = "tf-test-test-proxy%{random_suffix}"
  url_map          = google_compute_region_url_map.default.id
  ssl_certificates = [google_compute_region_ssl_certificate.default.id]
}

resource "google_compute_region_ssl_certificate" "default" {
  region      = "us-central1"
  name        = "tf-test-my-certificate%{random_suffix}"
  private_key = file("test-fixtures/ssl_cert/test.key")
  certificate = file("test-fixtures/ssl_cert/test.crt")
}

resource "google_compute_region_url_map" "default" {
  region      = "us-central1"
  name        = "tf-test-url-map%{random_suffix}"
  description = "a description"

  default_service = google_compute_region_backend_service.default.id

  host_rule {
    hosts        = ["mysite.com"]
    path_matcher = "allpaths"
  }

  path_matcher {
    name            = "allpaths"
    default_service = google_compute_region_backend_service.default.id

    path_rule {
      paths   = ["/*"]
      service = google_compute_region_backend_service.default.id
    }
  }
}

resource "google_compute_region_backend_service" "default" {
  region      = "us-central1"
  name        = "tf-test-backend-service%{random_suffix}"
  protocol    = "HTTP"
  timeout_sec = 10

  health_checks = [google_compute_region_health_check.default.id]
}

resource "google_compute_region_health_check" "default" {
  region = "us-central1"
  name   = "tf-test-http-health-check%{random_suffix}"
  http_health_check {
    port = 80
  }
}
`, context)
}

func testAccCheckComputeRegionTargetHttpsProxyDestroyProducer(t *testing.T) func(s *terraform.State) error {
	return func(s *terraform.State) error {
		for name, rs := range s.RootModule().Resources {
			if rs.Type != "google_compute_region_target_https_proxy" {
				continue
			}
			if strings.HasPrefix(name, "data.") {
				continue
			}

			config := googleProviderConfig(t)

			url, err := replaceVarsForTest(config, rs, "{{ComputeBasePath}}projects/{{project}}/regions/{{region}}/targetHttpsProxies/{{name}}")
			if err != nil {
				return err
			}

			_, err = sendRequest(config, "GET", "", url, nil)
			if err == nil {
				return fmt.Errorf("ComputeRegionTargetHttpsProxy still exists at %s", url)
			}
		}

		return nil
	}
}
