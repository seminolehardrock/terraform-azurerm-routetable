package test

import (
	"testing"

	"github.com/gruntwork-io/terratest/modules/terraform"
	test_structure "github.com/gruntwork-io/terratest/modules/test-structure"
)

func TestTerraformRouteTableNewResourceGroup(t *testing.T) {
	t.Parallel()

	fixtureFolder := "./fixture_new_rg"

	// Deploy the example
	test_structure.RunTestStage(t, "setup", func() {
		terraformOptions := &terraform.Options{
			TerraformDir: fixtureFolder,
			Vars:         map[string]interface{}{},
		}

		// Save the options so later test stages can use them
		test_structure.SaveTerraformOptions(t, fixtureFolder, terraformOptions)

		// This will init and apply the resources and fail the test if there are any errors
		terraform.InitAndApply(t, terraformOptions)
	})

	// Check whether the length of output meets the requirement
	test_structure.RunTestStage(t, "validate", func() {
		terraformOptions := test_structure.LoadTerraformOptions(t, fixtureFolder)

		routeTableID := terraform.Output(t, terraformOptions, "test_routetable_id")
		if len(routeTableID) <= 0 {
			t.Fatal("Wrong output")
		}
	})

	// At the end of the test, clean up any resources that were created
	test_structure.RunTestStage(t, "teardown", func() {
		terraformOptions := test_structure.LoadTerraformOptions(t, fixtureFolder)
		terraform.Destroy(t, terraformOptions)
	})

}

func TestTerraformRouteTableExistingResourceGroup(t *testing.T) {
	t.Parallel()

	fixtureFolder := "./fixture_existing_rg"

	// Deploy the example
	test_structure.RunTestStage(t, "setup", func() {
		terraformOptions := &terraform.Options{
			TerraformDir: fixtureFolder,
			Vars:         map[string]interface{}{},
		}

		// Save the options so later test stages can use them
		test_structure.SaveTerraformOptions(t, fixtureFolder, terraformOptions)

		// This will init and apply the resources and fail the test if there are any errors
		terraform.InitAndApply(t, terraformOptions)
	})

	// Check whether the length of output meets the requirement
	test_structure.RunTestStage(t, "validate", func() {
		terraformOptions := test_structure.LoadTerraformOptions(t, fixtureFolder)

		routeTableID := terraform.Output(t, terraformOptions, "test_routetable_id")
		if len(routeTableID) <= 0 {
			t.Fatal("Wrong output")
		}
	})

	// At the end of the test, clean up any resources that were created
	test_structure.RunTestStage(t, "teardown", func() {
		terraformOptions := test_structure.LoadTerraformOptions(t, fixtureFolder)
		terraform.Destroy(t, terraformOptions)
	})

}

func TestTerraformRouteTableNextHopIpAddress(t *testing.T) {
	t.Parallel()

	fixtureFolder := "./fixture_nexthop_ip_address"

	// Deploy the example
	test_structure.RunTestStage(t, "setup", func() {
		terraformOptions := &terraform.Options{
			TerraformDir: fixtureFolder,
			Vars:         map[string]interface{}{},
		}

		// Save the options so later test stages can use them
		test_structure.SaveTerraformOptions(t, fixtureFolder, terraformOptions)

		// This will init and apply the resources and fail the test if there are any errors
		terraform.InitAndApply(t, terraformOptions)
	})

	// Check whether the length of output meets the requirement
	test_structure.RunTestStage(t, "validate", func() {
		terraformOptions := test_structure.LoadTerraformOptions(t, fixtureFolder)

		routeTableID := terraform.Output(t, terraformOptions, "test_routetable_id")
		if len(routeTableID) <= 0 {
			t.Fatal("Wrong output")
		}
	})

	// At the end of the test, clean up any resources that were created
	test_structure.RunTestStage(t, "teardown", func() {
		terraformOptions := test_structure.LoadTerraformOptions(t, fixtureFolder)
		terraform.Destroy(t, terraformOptions)
	})

}
