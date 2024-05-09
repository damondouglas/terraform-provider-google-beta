// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

// ----------------------------------------------------------------------------
//
//     ***     AUTO GENERATED CODE    ***    Type: MMv1     ***
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

package workstations_test

import (
	"fmt"
	"strings"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"

	"github.com/hashicorp/terraform-provider-google-beta/google-beta/acctest"
	"github.com/hashicorp/terraform-provider-google-beta/google-beta/tpgresource"
	transport_tpg "github.com/hashicorp/terraform-provider-google-beta/google-beta/transport"
)

func TestAccWorkstationsWorkstationCluster_workstationClusterBasicExample(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix": acctest.RandString(t, 10),
	}

	acctest.VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderBetaFactories(t),
		CheckDestroy:             testAccCheckWorkstationsWorkstationClusterDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccWorkstationsWorkstationCluster_workstationClusterBasicExample(context),
			},
			{
				ResourceName:            "google_workstations_workstation_cluster.default",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"annotations", "labels", "location", "terraform_labels", "workstation_cluster_id"},
			},
		},
	})
}

func testAccWorkstationsWorkstationCluster_workstationClusterBasicExample(context map[string]interface{}) string {
	return acctest.Nprintf(`
resource "google_workstations_workstation_cluster" "default" {
  provider               = google-beta
  workstation_cluster_id = "tf-test-workstation-cluster%{random_suffix}"
  network                = google_compute_network.default.id
  subnetwork             = google_compute_subnetwork.default.id
  location               = "us-central1"
  
  labels = {
    "label" = "key"
  }

  annotations = {
    label-one = "value-one"
  }
}

data "google_project" "project" {
  provider = google-beta
}

resource "google_compute_network" "default" {
  provider                = google-beta
  name                    = "tf-test-workstation-cluster%{random_suffix}"
  auto_create_subnetworks = false
}

resource "google_compute_subnetwork" "default" {
  provider      = google-beta
  name          = "tf-test-workstation-cluster%{random_suffix}"
  ip_cidr_range = "10.0.0.0/24"
  region        = "us-central1"
  network       = google_compute_network.default.name
}
`, context)
}

func TestAccWorkstationsWorkstationCluster_workstationClusterPrivateExample(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix": acctest.RandString(t, 10),
	}

	acctest.VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderBetaFactories(t),
		CheckDestroy:             testAccCheckWorkstationsWorkstationClusterDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccWorkstationsWorkstationCluster_workstationClusterPrivateExample(context),
			},
			{
				ResourceName:            "google_workstations_workstation_cluster.default",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"annotations", "labels", "location", "terraform_labels", "workstation_cluster_id"},
			},
		},
	})
}

func testAccWorkstationsWorkstationCluster_workstationClusterPrivateExample(context map[string]interface{}) string {
	return acctest.Nprintf(`
resource "google_workstations_workstation_cluster" "default" {
  provider               = google-beta
  workstation_cluster_id = "tf-test-workstation-cluster-private%{random_suffix}"
  network                = google_compute_network.default.id
  subnetwork             = google_compute_subnetwork.default.id
  location               = "us-central1"

  private_cluster_config {
    enable_private_endpoint = true
  }

  labels = {
    "label" = "key"
  }

  annotations = {
    label-one = "value-one"
  }
}

data "google_project" "project" {
  provider = google-beta
}

resource "google_compute_network" "default" {
  provider                = google-beta
  name                    = "tf-test-workstation-cluster-private%{random_suffix}"
  auto_create_subnetworks = false
}

resource "google_compute_subnetwork" "default" {
  provider = google-beta
  name          = "tf-test-workstation-cluster-private%{random_suffix}"
  ip_cidr_range = "10.0.0.0/24"
  region        = "us-central1"
  network       = google_compute_network.default.name
}
`, context)
}

func TestAccWorkstationsWorkstationCluster_workstationClusterCustomDomainExample(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix": acctest.RandString(t, 10),
	}

	acctest.VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderBetaFactories(t),
		CheckDestroy:             testAccCheckWorkstationsWorkstationClusterDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccWorkstationsWorkstationCluster_workstationClusterCustomDomainExample(context),
			},
			{
				ResourceName:            "google_workstations_workstation_cluster.default",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"annotations", "labels", "location", "terraform_labels", "workstation_cluster_id"},
			},
		},
	})
}

func testAccWorkstationsWorkstationCluster_workstationClusterCustomDomainExample(context map[string]interface{}) string {
	return acctest.Nprintf(`
resource "google_workstations_workstation_cluster" "default" {
  provider               = google-beta
  workstation_cluster_id = "tf-test-workstation-cluster-custom-domain%{random_suffix}"
  network                = google_compute_network.default.id
  subnetwork             = google_compute_subnetwork.default.id
  location               = "us-central1"

  private_cluster_config {
    enable_private_endpoint = true
  }

  domain_config {
    domain = "workstations.example.com"
  }

  labels = {
    "label" = "key"
  }

  annotations = {
    label-one = "value-one"
  }
}

data "google_project" "project" {
  provider = google-beta
}

resource "google_compute_network" "default" {
  provider                = google-beta
  name                    = "tf-test-workstation-cluster-custom-domain%{random_suffix}"
  auto_create_subnetworks = false
}

resource "google_compute_subnetwork" "default" {
  provider = google-beta
  name          = "tf-test-workstation-cluster-custom-domain%{random_suffix}"
  ip_cidr_range = "10.0.0.0/24"
  region        = "us-central1"
  network       = google_compute_network.default.name
}
`, context)
}

func testAccCheckWorkstationsWorkstationClusterDestroyProducer(t *testing.T) func(s *terraform.State) error {
	return func(s *terraform.State) error {
		for name, rs := range s.RootModule().Resources {
			if rs.Type != "google_workstations_workstation_cluster" {
				continue
			}
			if strings.HasPrefix(name, "data.") {
				continue
			}

			config := acctest.GoogleProviderConfig(t)

			url, err := tpgresource.ReplaceVarsForTest(config, rs, "{{WorkstationsBasePath}}projects/{{project}}/locations/{{location}}/workstationClusters/{{workstation_cluster_id}}")
			if err != nil {
				return err
			}

			billingProject := ""

			if config.BillingProject != "" {
				billingProject = config.BillingProject
			}

			_, err = transport_tpg.SendRequest(transport_tpg.SendRequestOptions{
				Config:    config,
				Method:    "GET",
				Project:   billingProject,
				RawURL:    url,
				UserAgent: config.UserAgent,
			})
			if err == nil {
				return fmt.Errorf("WorkstationsWorkstationCluster still exists at %s", url)
			}
		}

		return nil
	}
}
