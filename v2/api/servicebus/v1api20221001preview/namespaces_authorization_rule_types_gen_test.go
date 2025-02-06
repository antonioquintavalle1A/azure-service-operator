// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1api20221001preview

import (
	"encoding/json"
	v20221001ps "github.com/Azure/azure-service-operator/v2/api/servicebus/v1api20221001preview/storage"
	v20240101s "github.com/Azure/azure-service-operator/v2/api/servicebus/v1api20240101/storage"
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

func Test_NamespacesAuthorizationRule_WhenConvertedToHub_RoundTripsWithoutLoss(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	parameters.MinSuccessfulTests = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip from NamespacesAuthorizationRule to hub returns original",
		prop.ForAll(RunResourceConversionTestForNamespacesAuthorizationRule, NamespacesAuthorizationRuleGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(false, 240, os.Stdout))
}

// RunResourceConversionTestForNamespacesAuthorizationRule tests if a specific instance of NamespacesAuthorizationRule round trips to the hub storage version and back losslessly
func RunResourceConversionTestForNamespacesAuthorizationRule(subject NamespacesAuthorizationRule) string {
	// Copy subject to make sure conversion doesn't modify it
	copied := subject.DeepCopy()

	// Convert to our hub version
	var hub v20240101s.NamespacesAuthorizationRule
	err := copied.ConvertTo(&hub)
	if err != nil {
		return err.Error()
	}

	// Convert from our hub version
	var actual NamespacesAuthorizationRule
	err = actual.ConvertFrom(&hub)
	if err != nil {
		return err.Error()
	}

	// Compare actual with what we started with
	match := cmp.Equal(subject, actual, cmpopts.EquateEmpty())
	if !match {
		actualFmt := pretty.Sprint(actual)
		subjectFmt := pretty.Sprint(subject)
		result := diff.Diff(subjectFmt, actualFmt)
		return result
	}

	return ""
}

func Test_NamespacesAuthorizationRule_WhenPropertiesConverted_RoundTripsWithoutLoss(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip from NamespacesAuthorizationRule to NamespacesAuthorizationRule via AssignProperties_To_NamespacesAuthorizationRule & AssignProperties_From_NamespacesAuthorizationRule returns original",
		prop.ForAll(RunPropertyAssignmentTestForNamespacesAuthorizationRule, NamespacesAuthorizationRuleGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(false, 240, os.Stdout))
}

// RunPropertyAssignmentTestForNamespacesAuthorizationRule tests if a specific instance of NamespacesAuthorizationRule can be assigned to storage and back losslessly
func RunPropertyAssignmentTestForNamespacesAuthorizationRule(subject NamespacesAuthorizationRule) string {
	// Copy subject to make sure assignment doesn't modify it
	copied := subject.DeepCopy()

	// Use AssignPropertiesTo() for the first stage of conversion
	var other v20221001ps.NamespacesAuthorizationRule
	err := copied.AssignProperties_To_NamespacesAuthorizationRule(&other)
	if err != nil {
		return err.Error()
	}

	// Use AssignPropertiesFrom() to convert back to our original type
	var actual NamespacesAuthorizationRule
	err = actual.AssignProperties_From_NamespacesAuthorizationRule(&other)
	if err != nil {
		return err.Error()
	}

	// Check for a match
	match := cmp.Equal(subject, actual, cmpopts.EquateEmpty())
	if !match {
		actualFmt := pretty.Sprint(actual)
		subjectFmt := pretty.Sprint(subject)
		result := diff.Diff(subjectFmt, actualFmt)
		return result
	}

	return ""
}

func Test_NamespacesAuthorizationRule_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 20
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of NamespacesAuthorizationRule via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForNamespacesAuthorizationRule, NamespacesAuthorizationRuleGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForNamespacesAuthorizationRule runs a test to see if a specific instance of NamespacesAuthorizationRule round trips to JSON and back losslessly
func RunJSONSerializationTestForNamespacesAuthorizationRule(subject NamespacesAuthorizationRule) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual NamespacesAuthorizationRule
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

// Generator of NamespacesAuthorizationRule instances for property testing - lazily instantiated by
// NamespacesAuthorizationRuleGenerator()
var namespacesAuthorizationRuleGenerator gopter.Gen

// NamespacesAuthorizationRuleGenerator returns a generator of NamespacesAuthorizationRule instances for property testing.
func NamespacesAuthorizationRuleGenerator() gopter.Gen {
	if namespacesAuthorizationRuleGenerator != nil {
		return namespacesAuthorizationRuleGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddRelatedPropertyGeneratorsForNamespacesAuthorizationRule(generators)
	namespacesAuthorizationRuleGenerator = gen.Struct(reflect.TypeOf(NamespacesAuthorizationRule{}), generators)

	return namespacesAuthorizationRuleGenerator
}

// AddRelatedPropertyGeneratorsForNamespacesAuthorizationRule is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForNamespacesAuthorizationRule(gens map[string]gopter.Gen) {
	gens["Spec"] = NamespacesAuthorizationRule_SpecGenerator()
	gens["Status"] = NamespacesAuthorizationRule_STATUSGenerator()
}

func Test_NamespacesAuthorizationRuleOperatorSecrets_WhenPropertiesConverted_RoundTripsWithoutLoss(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip from NamespacesAuthorizationRuleOperatorSecrets to NamespacesAuthorizationRuleOperatorSecrets via AssignProperties_To_NamespacesAuthorizationRuleOperatorSecrets & AssignProperties_From_NamespacesAuthorizationRuleOperatorSecrets returns original",
		prop.ForAll(RunPropertyAssignmentTestForNamespacesAuthorizationRuleOperatorSecrets, NamespacesAuthorizationRuleOperatorSecretsGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(false, 240, os.Stdout))
}

// RunPropertyAssignmentTestForNamespacesAuthorizationRuleOperatorSecrets tests if a specific instance of NamespacesAuthorizationRuleOperatorSecrets can be assigned to storage and back losslessly
func RunPropertyAssignmentTestForNamespacesAuthorizationRuleOperatorSecrets(subject NamespacesAuthorizationRuleOperatorSecrets) string {
	// Copy subject to make sure assignment doesn't modify it
	copied := subject.DeepCopy()

	// Use AssignPropertiesTo() for the first stage of conversion
	var other v20221001ps.NamespacesAuthorizationRuleOperatorSecrets
	err := copied.AssignProperties_To_NamespacesAuthorizationRuleOperatorSecrets(&other)
	if err != nil {
		return err.Error()
	}

	// Use AssignPropertiesFrom() to convert back to our original type
	var actual NamespacesAuthorizationRuleOperatorSecrets
	err = actual.AssignProperties_From_NamespacesAuthorizationRuleOperatorSecrets(&other)
	if err != nil {
		return err.Error()
	}

	// Check for a match
	match := cmp.Equal(subject, actual, cmpopts.EquateEmpty())
	if !match {
		actualFmt := pretty.Sprint(actual)
		subjectFmt := pretty.Sprint(subject)
		result := diff.Diff(subjectFmt, actualFmt)
		return result
	}

	return ""
}

func Test_NamespacesAuthorizationRuleOperatorSecrets_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 100
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of NamespacesAuthorizationRuleOperatorSecrets via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForNamespacesAuthorizationRuleOperatorSecrets, NamespacesAuthorizationRuleOperatorSecretsGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForNamespacesAuthorizationRuleOperatorSecrets runs a test to see if a specific instance of NamespacesAuthorizationRuleOperatorSecrets round trips to JSON and back losslessly
func RunJSONSerializationTestForNamespacesAuthorizationRuleOperatorSecrets(subject NamespacesAuthorizationRuleOperatorSecrets) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual NamespacesAuthorizationRuleOperatorSecrets
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

// Generator of NamespacesAuthorizationRuleOperatorSecrets instances for property testing - lazily instantiated by
// NamespacesAuthorizationRuleOperatorSecretsGenerator()
var namespacesAuthorizationRuleOperatorSecretsGenerator gopter.Gen

// NamespacesAuthorizationRuleOperatorSecretsGenerator returns a generator of NamespacesAuthorizationRuleOperatorSecrets instances for property testing.
func NamespacesAuthorizationRuleOperatorSecretsGenerator() gopter.Gen {
	if namespacesAuthorizationRuleOperatorSecretsGenerator != nil {
		return namespacesAuthorizationRuleOperatorSecretsGenerator
	}

	generators := make(map[string]gopter.Gen)
	namespacesAuthorizationRuleOperatorSecretsGenerator = gen.Struct(reflect.TypeOf(NamespacesAuthorizationRuleOperatorSecrets{}), generators)

	return namespacesAuthorizationRuleOperatorSecretsGenerator
}

func Test_NamespacesAuthorizationRuleOperatorSpec_WhenPropertiesConverted_RoundTripsWithoutLoss(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip from NamespacesAuthorizationRuleOperatorSpec to NamespacesAuthorizationRuleOperatorSpec via AssignProperties_To_NamespacesAuthorizationRuleOperatorSpec & AssignProperties_From_NamespacesAuthorizationRuleOperatorSpec returns original",
		prop.ForAll(RunPropertyAssignmentTestForNamespacesAuthorizationRuleOperatorSpec, NamespacesAuthorizationRuleOperatorSpecGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(false, 240, os.Stdout))
}

// RunPropertyAssignmentTestForNamespacesAuthorizationRuleOperatorSpec tests if a specific instance of NamespacesAuthorizationRuleOperatorSpec can be assigned to storage and back losslessly
func RunPropertyAssignmentTestForNamespacesAuthorizationRuleOperatorSpec(subject NamespacesAuthorizationRuleOperatorSpec) string {
	// Copy subject to make sure assignment doesn't modify it
	copied := subject.DeepCopy()

	// Use AssignPropertiesTo() for the first stage of conversion
	var other v20221001ps.NamespacesAuthorizationRuleOperatorSpec
	err := copied.AssignProperties_To_NamespacesAuthorizationRuleOperatorSpec(&other)
	if err != nil {
		return err.Error()
	}

	// Use AssignPropertiesFrom() to convert back to our original type
	var actual NamespacesAuthorizationRuleOperatorSpec
	err = actual.AssignProperties_From_NamespacesAuthorizationRuleOperatorSpec(&other)
	if err != nil {
		return err.Error()
	}

	// Check for a match
	match := cmp.Equal(subject, actual, cmpopts.EquateEmpty())
	if !match {
		actualFmt := pretty.Sprint(actual)
		subjectFmt := pretty.Sprint(subject)
		result := diff.Diff(subjectFmt, actualFmt)
		return result
	}

	return ""
}

func Test_NamespacesAuthorizationRuleOperatorSpec_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 100
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of NamespacesAuthorizationRuleOperatorSpec via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForNamespacesAuthorizationRuleOperatorSpec, NamespacesAuthorizationRuleOperatorSpecGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForNamespacesAuthorizationRuleOperatorSpec runs a test to see if a specific instance of NamespacesAuthorizationRuleOperatorSpec round trips to JSON and back losslessly
func RunJSONSerializationTestForNamespacesAuthorizationRuleOperatorSpec(subject NamespacesAuthorizationRuleOperatorSpec) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual NamespacesAuthorizationRuleOperatorSpec
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

// Generator of NamespacesAuthorizationRuleOperatorSpec instances for property testing - lazily instantiated by
// NamespacesAuthorizationRuleOperatorSpecGenerator()
var namespacesAuthorizationRuleOperatorSpecGenerator gopter.Gen

// NamespacesAuthorizationRuleOperatorSpecGenerator returns a generator of NamespacesAuthorizationRuleOperatorSpec instances for property testing.
func NamespacesAuthorizationRuleOperatorSpecGenerator() gopter.Gen {
	if namespacesAuthorizationRuleOperatorSpecGenerator != nil {
		return namespacesAuthorizationRuleOperatorSpecGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddRelatedPropertyGeneratorsForNamespacesAuthorizationRuleOperatorSpec(generators)
	namespacesAuthorizationRuleOperatorSpecGenerator = gen.Struct(reflect.TypeOf(NamespacesAuthorizationRuleOperatorSpec{}), generators)

	return namespacesAuthorizationRuleOperatorSpecGenerator
}

// AddRelatedPropertyGeneratorsForNamespacesAuthorizationRuleOperatorSpec is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForNamespacesAuthorizationRuleOperatorSpec(gens map[string]gopter.Gen) {
	gens["Secrets"] = gen.PtrOf(NamespacesAuthorizationRuleOperatorSecretsGenerator())
}

func Test_NamespacesAuthorizationRule_STATUS_WhenPropertiesConverted_RoundTripsWithoutLoss(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip from NamespacesAuthorizationRule_STATUS to NamespacesAuthorizationRule_STATUS via AssignProperties_To_NamespacesAuthorizationRule_STATUS & AssignProperties_From_NamespacesAuthorizationRule_STATUS returns original",
		prop.ForAll(RunPropertyAssignmentTestForNamespacesAuthorizationRule_STATUS, NamespacesAuthorizationRule_STATUSGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(false, 240, os.Stdout))
}

// RunPropertyAssignmentTestForNamespacesAuthorizationRule_STATUS tests if a specific instance of NamespacesAuthorizationRule_STATUS can be assigned to storage and back losslessly
func RunPropertyAssignmentTestForNamespacesAuthorizationRule_STATUS(subject NamespacesAuthorizationRule_STATUS) string {
	// Copy subject to make sure assignment doesn't modify it
	copied := subject.DeepCopy()

	// Use AssignPropertiesTo() for the first stage of conversion
	var other v20221001ps.NamespacesAuthorizationRule_STATUS
	err := copied.AssignProperties_To_NamespacesAuthorizationRule_STATUS(&other)
	if err != nil {
		return err.Error()
	}

	// Use AssignPropertiesFrom() to convert back to our original type
	var actual NamespacesAuthorizationRule_STATUS
	err = actual.AssignProperties_From_NamespacesAuthorizationRule_STATUS(&other)
	if err != nil {
		return err.Error()
	}

	// Check for a match
	match := cmp.Equal(subject, actual, cmpopts.EquateEmpty())
	if !match {
		actualFmt := pretty.Sprint(actual)
		subjectFmt := pretty.Sprint(subject)
		result := diff.Diff(subjectFmt, actualFmt)
		return result
	}

	return ""
}

func Test_NamespacesAuthorizationRule_STATUS_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of NamespacesAuthorizationRule_STATUS via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForNamespacesAuthorizationRule_STATUS, NamespacesAuthorizationRule_STATUSGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForNamespacesAuthorizationRule_STATUS runs a test to see if a specific instance of NamespacesAuthorizationRule_STATUS round trips to JSON and back losslessly
func RunJSONSerializationTestForNamespacesAuthorizationRule_STATUS(subject NamespacesAuthorizationRule_STATUS) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual NamespacesAuthorizationRule_STATUS
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

// Generator of NamespacesAuthorizationRule_STATUS instances for property testing - lazily instantiated by
// NamespacesAuthorizationRule_STATUSGenerator()
var namespacesAuthorizationRule_STATUSGenerator gopter.Gen

// NamespacesAuthorizationRule_STATUSGenerator returns a generator of NamespacesAuthorizationRule_STATUS instances for property testing.
// We first initialize namespacesAuthorizationRule_STATUSGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func NamespacesAuthorizationRule_STATUSGenerator() gopter.Gen {
	if namespacesAuthorizationRule_STATUSGenerator != nil {
		return namespacesAuthorizationRule_STATUSGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForNamespacesAuthorizationRule_STATUS(generators)
	namespacesAuthorizationRule_STATUSGenerator = gen.Struct(reflect.TypeOf(NamespacesAuthorizationRule_STATUS{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForNamespacesAuthorizationRule_STATUS(generators)
	AddRelatedPropertyGeneratorsForNamespacesAuthorizationRule_STATUS(generators)
	namespacesAuthorizationRule_STATUSGenerator = gen.Struct(reflect.TypeOf(NamespacesAuthorizationRule_STATUS{}), generators)

	return namespacesAuthorizationRule_STATUSGenerator
}

// AddIndependentPropertyGeneratorsForNamespacesAuthorizationRule_STATUS is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForNamespacesAuthorizationRule_STATUS(gens map[string]gopter.Gen) {
	gens["Id"] = gen.PtrOf(gen.AlphaString())
	gens["Location"] = gen.PtrOf(gen.AlphaString())
	gens["Name"] = gen.PtrOf(gen.AlphaString())
	gens["Rights"] = gen.SliceOf(gen.OneConstOf(Namespaces_AuthorizationRule_Properties_Rights_STATUS_Listen, Namespaces_AuthorizationRule_Properties_Rights_STATUS_Manage, Namespaces_AuthorizationRule_Properties_Rights_STATUS_Send))
	gens["Type"] = gen.PtrOf(gen.AlphaString())
}

// AddRelatedPropertyGeneratorsForNamespacesAuthorizationRule_STATUS is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForNamespacesAuthorizationRule_STATUS(gens map[string]gopter.Gen) {
	gens["SystemData"] = gen.PtrOf(SystemData_STATUSGenerator())
}

func Test_NamespacesAuthorizationRule_Spec_WhenPropertiesConverted_RoundTripsWithoutLoss(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip from NamespacesAuthorizationRule_Spec to NamespacesAuthorizationRule_Spec via AssignProperties_To_NamespacesAuthorizationRule_Spec & AssignProperties_From_NamespacesAuthorizationRule_Spec returns original",
		prop.ForAll(RunPropertyAssignmentTestForNamespacesAuthorizationRule_Spec, NamespacesAuthorizationRule_SpecGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(false, 240, os.Stdout))
}

// RunPropertyAssignmentTestForNamespacesAuthorizationRule_Spec tests if a specific instance of NamespacesAuthorizationRule_Spec can be assigned to storage and back losslessly
func RunPropertyAssignmentTestForNamespacesAuthorizationRule_Spec(subject NamespacesAuthorizationRule_Spec) string {
	// Copy subject to make sure assignment doesn't modify it
	copied := subject.DeepCopy()

	// Use AssignPropertiesTo() for the first stage of conversion
	var other v20221001ps.NamespacesAuthorizationRule_Spec
	err := copied.AssignProperties_To_NamespacesAuthorizationRule_Spec(&other)
	if err != nil {
		return err.Error()
	}

	// Use AssignPropertiesFrom() to convert back to our original type
	var actual NamespacesAuthorizationRule_Spec
	err = actual.AssignProperties_From_NamespacesAuthorizationRule_Spec(&other)
	if err != nil {
		return err.Error()
	}

	// Check for a match
	match := cmp.Equal(subject, actual, cmpopts.EquateEmpty())
	if !match {
		actualFmt := pretty.Sprint(actual)
		subjectFmt := pretty.Sprint(subject)
		result := diff.Diff(subjectFmt, actualFmt)
		return result
	}

	return ""
}

func Test_NamespacesAuthorizationRule_Spec_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of NamespacesAuthorizationRule_Spec via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForNamespacesAuthorizationRule_Spec, NamespacesAuthorizationRule_SpecGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForNamespacesAuthorizationRule_Spec runs a test to see if a specific instance of NamespacesAuthorizationRule_Spec round trips to JSON and back losslessly
func RunJSONSerializationTestForNamespacesAuthorizationRule_Spec(subject NamespacesAuthorizationRule_Spec) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual NamespacesAuthorizationRule_Spec
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

// Generator of NamespacesAuthorizationRule_Spec instances for property testing - lazily instantiated by
// NamespacesAuthorizationRule_SpecGenerator()
var namespacesAuthorizationRule_SpecGenerator gopter.Gen

// NamespacesAuthorizationRule_SpecGenerator returns a generator of NamespacesAuthorizationRule_Spec instances for property testing.
// We first initialize namespacesAuthorizationRule_SpecGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func NamespacesAuthorizationRule_SpecGenerator() gopter.Gen {
	if namespacesAuthorizationRule_SpecGenerator != nil {
		return namespacesAuthorizationRule_SpecGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForNamespacesAuthorizationRule_Spec(generators)
	namespacesAuthorizationRule_SpecGenerator = gen.Struct(reflect.TypeOf(NamespacesAuthorizationRule_Spec{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForNamespacesAuthorizationRule_Spec(generators)
	AddRelatedPropertyGeneratorsForNamespacesAuthorizationRule_Spec(generators)
	namespacesAuthorizationRule_SpecGenerator = gen.Struct(reflect.TypeOf(NamespacesAuthorizationRule_Spec{}), generators)

	return namespacesAuthorizationRule_SpecGenerator
}

// AddIndependentPropertyGeneratorsForNamespacesAuthorizationRule_Spec is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForNamespacesAuthorizationRule_Spec(gens map[string]gopter.Gen) {
	gens["AzureName"] = gen.AlphaString()
	gens["Rights"] = gen.SliceOf(gen.OneConstOf(Namespaces_AuthorizationRule_Properties_Rights_Spec_Listen, Namespaces_AuthorizationRule_Properties_Rights_Spec_Manage, Namespaces_AuthorizationRule_Properties_Rights_Spec_Send))
}

// AddRelatedPropertyGeneratorsForNamespacesAuthorizationRule_Spec is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForNamespacesAuthorizationRule_Spec(gens map[string]gopter.Gen) {
	gens["OperatorSpec"] = gen.PtrOf(NamespacesAuthorizationRuleOperatorSpecGenerator())
}
