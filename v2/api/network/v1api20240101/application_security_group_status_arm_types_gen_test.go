// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1api20240101

import (
	"encoding/json"
	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
	"github.com/kr/pretty"
	"github.com/kylelemons/godebug/diff"
	"github.com/leanovate/gopter"
	"github.com/leanovate/gopter/gen"
	"github.com/leanovate/gopter/prop"
	"os"
	"reflect"
	"testing"
)

func Test_ApplicationSecurityGroupPropertiesFormat_STATUS_ARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of ApplicationSecurityGroupPropertiesFormat_STATUS_ARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForApplicationSecurityGroupPropertiesFormat_STATUS_ARM, ApplicationSecurityGroupPropertiesFormat_STATUS_ARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForApplicationSecurityGroupPropertiesFormat_STATUS_ARM runs a test to see if a specific instance of ApplicationSecurityGroupPropertiesFormat_STATUS_ARM round trips to JSON and back losslessly
func RunJSONSerializationTestForApplicationSecurityGroupPropertiesFormat_STATUS_ARM(subject ApplicationSecurityGroupPropertiesFormat_STATUS_ARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual ApplicationSecurityGroupPropertiesFormat_STATUS_ARM
	err = json.Unmarshal(bin, &actual)
	if err != nil {
		return err.Error()
	}

	// Check for outcome
	match := cmp.Equal(subject, actual, cmpopts.EquateEmpty())
	if !match {
		actualFmt := pretty.Sprint(actual)
		subjectFmt := pretty.Sprint(subject)
		result := diff.Diff(subjectFmt, actualFmt)
		return result
	}

	return ""
}

// Generator of ApplicationSecurityGroupPropertiesFormat_STATUS_ARM instances for property testing - lazily instantiated
// by ApplicationSecurityGroupPropertiesFormat_STATUS_ARMGenerator()
var applicationSecurityGroupPropertiesFormat_STATUS_ARMGenerator gopter.Gen

// ApplicationSecurityGroupPropertiesFormat_STATUS_ARMGenerator returns a generator of ApplicationSecurityGroupPropertiesFormat_STATUS_ARM instances for property testing.
func ApplicationSecurityGroupPropertiesFormat_STATUS_ARMGenerator() gopter.Gen {
	if applicationSecurityGroupPropertiesFormat_STATUS_ARMGenerator != nil {
		return applicationSecurityGroupPropertiesFormat_STATUS_ARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForApplicationSecurityGroupPropertiesFormat_STATUS_ARM(generators)
	applicationSecurityGroupPropertiesFormat_STATUS_ARMGenerator = gen.Struct(reflect.TypeOf(ApplicationSecurityGroupPropertiesFormat_STATUS_ARM{}), generators)

	return applicationSecurityGroupPropertiesFormat_STATUS_ARMGenerator
}

// AddIndependentPropertyGeneratorsForApplicationSecurityGroupPropertiesFormat_STATUS_ARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForApplicationSecurityGroupPropertiesFormat_STATUS_ARM(gens map[string]gopter.Gen) {
	gens["ProvisioningState"] = gen.PtrOf(gen.OneConstOf(
		ProvisioningState_STATUS_ARM_Deleting,
		ProvisioningState_STATUS_ARM_Failed,
		ProvisioningState_STATUS_ARM_Succeeded,
		ProvisioningState_STATUS_ARM_Updating))
	gens["ResourceGuid"] = gen.PtrOf(gen.AlphaString())
}

func Test_ApplicationSecurityGroup_STATUS_ARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of ApplicationSecurityGroup_STATUS_ARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForApplicationSecurityGroup_STATUS_ARM, ApplicationSecurityGroup_STATUS_ARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForApplicationSecurityGroup_STATUS_ARM runs a test to see if a specific instance of ApplicationSecurityGroup_STATUS_ARM round trips to JSON and back losslessly
func RunJSONSerializationTestForApplicationSecurityGroup_STATUS_ARM(subject ApplicationSecurityGroup_STATUS_ARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual ApplicationSecurityGroup_STATUS_ARM
	err = json.Unmarshal(bin, &actual)
	if err != nil {
		return err.Error()
	}

	// Check for outcome
	match := cmp.Equal(subject, actual, cmpopts.EquateEmpty())
	if !match {
		actualFmt := pretty.Sprint(actual)
		subjectFmt := pretty.Sprint(subject)
		result := diff.Diff(subjectFmt, actualFmt)
		return result
	}

	return ""
}

// Generator of ApplicationSecurityGroup_STATUS_ARM instances for property testing - lazily instantiated by
// ApplicationSecurityGroup_STATUS_ARMGenerator()
var applicationSecurityGroup_STATUS_ARMGenerator gopter.Gen

// ApplicationSecurityGroup_STATUS_ARMGenerator returns a generator of ApplicationSecurityGroup_STATUS_ARM instances for property testing.
// We first initialize applicationSecurityGroup_STATUS_ARMGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func ApplicationSecurityGroup_STATUS_ARMGenerator() gopter.Gen {
	if applicationSecurityGroup_STATUS_ARMGenerator != nil {
		return applicationSecurityGroup_STATUS_ARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForApplicationSecurityGroup_STATUS_ARM(generators)
	applicationSecurityGroup_STATUS_ARMGenerator = gen.Struct(reflect.TypeOf(ApplicationSecurityGroup_STATUS_ARM{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForApplicationSecurityGroup_STATUS_ARM(generators)
	AddRelatedPropertyGeneratorsForApplicationSecurityGroup_STATUS_ARM(generators)
	applicationSecurityGroup_STATUS_ARMGenerator = gen.Struct(reflect.TypeOf(ApplicationSecurityGroup_STATUS_ARM{}), generators)

	return applicationSecurityGroup_STATUS_ARMGenerator
}

// AddIndependentPropertyGeneratorsForApplicationSecurityGroup_STATUS_ARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForApplicationSecurityGroup_STATUS_ARM(gens map[string]gopter.Gen) {
	gens["Etag"] = gen.PtrOf(gen.AlphaString())
	gens["Id"] = gen.PtrOf(gen.AlphaString())
	gens["Location"] = gen.PtrOf(gen.AlphaString())
	gens["Name"] = gen.PtrOf(gen.AlphaString())
	gens["Tags"] = gen.MapOf(
		gen.AlphaString(),
		gen.AlphaString())
	gens["Type"] = gen.PtrOf(gen.AlphaString())
}

// AddRelatedPropertyGeneratorsForApplicationSecurityGroup_STATUS_ARM is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForApplicationSecurityGroup_STATUS_ARM(gens map[string]gopter.Gen) {
	gens["Properties"] = gen.PtrOf(ApplicationSecurityGroupPropertiesFormat_STATUS_ARMGenerator())
}
