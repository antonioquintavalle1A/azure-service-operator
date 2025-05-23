// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1api20240101

import (
	"encoding/json"
	storage "github.com/Azure/azure-service-operator/v2/api/servicebus/v1api20240101/storage"
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

func Test_NamespacesTopic_WhenConvertedToHub_RoundTripsWithoutLoss(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	parameters.MinSuccessfulTests = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip from NamespacesTopic to hub returns original",
		prop.ForAll(RunResourceConversionTestForNamespacesTopic, NamespacesTopicGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(false, 240, os.Stdout))
}

// RunResourceConversionTestForNamespacesTopic tests if a specific instance of NamespacesTopic round trips to the hub storage version and back losslessly
func RunResourceConversionTestForNamespacesTopic(subject NamespacesTopic) string {
	// Copy subject to make sure conversion doesn't modify it
	copied := subject.DeepCopy()

	// Convert to our hub version
	var hub storage.NamespacesTopic
	err := copied.ConvertTo(&hub)
	if err != nil {
		return err.Error()
	}

	// Convert from our hub version
	var actual NamespacesTopic
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

func Test_NamespacesTopic_WhenPropertiesConverted_RoundTripsWithoutLoss(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip from NamespacesTopic to NamespacesTopic via AssignProperties_To_NamespacesTopic & AssignProperties_From_NamespacesTopic returns original",
		prop.ForAll(RunPropertyAssignmentTestForNamespacesTopic, NamespacesTopicGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(false, 240, os.Stdout))
}

// RunPropertyAssignmentTestForNamespacesTopic tests if a specific instance of NamespacesTopic can be assigned to storage and back losslessly
func RunPropertyAssignmentTestForNamespacesTopic(subject NamespacesTopic) string {
	// Copy subject to make sure assignment doesn't modify it
	copied := subject.DeepCopy()

	// Use AssignPropertiesTo() for the first stage of conversion
	var other storage.NamespacesTopic
	err := copied.AssignProperties_To_NamespacesTopic(&other)
	if err != nil {
		return err.Error()
	}

	// Use AssignPropertiesFrom() to convert back to our original type
	var actual NamespacesTopic
	err = actual.AssignProperties_From_NamespacesTopic(&other)
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

func Test_NamespacesTopic_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 20
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of NamespacesTopic via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForNamespacesTopic, NamespacesTopicGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForNamespacesTopic runs a test to see if a specific instance of NamespacesTopic round trips to JSON and back losslessly
func RunJSONSerializationTestForNamespacesTopic(subject NamespacesTopic) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual NamespacesTopic
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

// Generator of NamespacesTopic instances for property testing - lazily instantiated by NamespacesTopicGenerator()
var namespacesTopicGenerator gopter.Gen

// NamespacesTopicGenerator returns a generator of NamespacesTopic instances for property testing.
func NamespacesTopicGenerator() gopter.Gen {
	if namespacesTopicGenerator != nil {
		return namespacesTopicGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddRelatedPropertyGeneratorsForNamespacesTopic(generators)
	namespacesTopicGenerator = gen.Struct(reflect.TypeOf(NamespacesTopic{}), generators)

	return namespacesTopicGenerator
}

// AddRelatedPropertyGeneratorsForNamespacesTopic is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForNamespacesTopic(gens map[string]gopter.Gen) {
	gens["Spec"] = NamespacesTopic_SpecGenerator()
	gens["Status"] = NamespacesTopic_STATUSGenerator()
}

func Test_NamespacesTopicOperatorSpec_WhenPropertiesConverted_RoundTripsWithoutLoss(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip from NamespacesTopicOperatorSpec to NamespacesTopicOperatorSpec via AssignProperties_To_NamespacesTopicOperatorSpec & AssignProperties_From_NamespacesTopicOperatorSpec returns original",
		prop.ForAll(RunPropertyAssignmentTestForNamespacesTopicOperatorSpec, NamespacesTopicOperatorSpecGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(false, 240, os.Stdout))
}

// RunPropertyAssignmentTestForNamespacesTopicOperatorSpec tests if a specific instance of NamespacesTopicOperatorSpec can be assigned to storage and back losslessly
func RunPropertyAssignmentTestForNamespacesTopicOperatorSpec(subject NamespacesTopicOperatorSpec) string {
	// Copy subject to make sure assignment doesn't modify it
	copied := subject.DeepCopy()

	// Use AssignPropertiesTo() for the first stage of conversion
	var other storage.NamespacesTopicOperatorSpec
	err := copied.AssignProperties_To_NamespacesTopicOperatorSpec(&other)
	if err != nil {
		return err.Error()
	}

	// Use AssignPropertiesFrom() to convert back to our original type
	var actual NamespacesTopicOperatorSpec
	err = actual.AssignProperties_From_NamespacesTopicOperatorSpec(&other)
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

func Test_NamespacesTopicOperatorSpec_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 100
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of NamespacesTopicOperatorSpec via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForNamespacesTopicOperatorSpec, NamespacesTopicOperatorSpecGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForNamespacesTopicOperatorSpec runs a test to see if a specific instance of NamespacesTopicOperatorSpec round trips to JSON and back losslessly
func RunJSONSerializationTestForNamespacesTopicOperatorSpec(subject NamespacesTopicOperatorSpec) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual NamespacesTopicOperatorSpec
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

// Generator of NamespacesTopicOperatorSpec instances for property testing - lazily instantiated by
// NamespacesTopicOperatorSpecGenerator()
var namespacesTopicOperatorSpecGenerator gopter.Gen

// NamespacesTopicOperatorSpecGenerator returns a generator of NamespacesTopicOperatorSpec instances for property testing.
func NamespacesTopicOperatorSpecGenerator() gopter.Gen {
	if namespacesTopicOperatorSpecGenerator != nil {
		return namespacesTopicOperatorSpecGenerator
	}

	generators := make(map[string]gopter.Gen)
	namespacesTopicOperatorSpecGenerator = gen.Struct(reflect.TypeOf(NamespacesTopicOperatorSpec{}), generators)

	return namespacesTopicOperatorSpecGenerator
}

func Test_NamespacesTopic_STATUS_WhenPropertiesConverted_RoundTripsWithoutLoss(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip from NamespacesTopic_STATUS to NamespacesTopic_STATUS via AssignProperties_To_NamespacesTopic_STATUS & AssignProperties_From_NamespacesTopic_STATUS returns original",
		prop.ForAll(RunPropertyAssignmentTestForNamespacesTopic_STATUS, NamespacesTopic_STATUSGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(false, 240, os.Stdout))
}

// RunPropertyAssignmentTestForNamespacesTopic_STATUS tests if a specific instance of NamespacesTopic_STATUS can be assigned to storage and back losslessly
func RunPropertyAssignmentTestForNamespacesTopic_STATUS(subject NamespacesTopic_STATUS) string {
	// Copy subject to make sure assignment doesn't modify it
	copied := subject.DeepCopy()

	// Use AssignPropertiesTo() for the first stage of conversion
	var other storage.NamespacesTopic_STATUS
	err := copied.AssignProperties_To_NamespacesTopic_STATUS(&other)
	if err != nil {
		return err.Error()
	}

	// Use AssignPropertiesFrom() to convert back to our original type
	var actual NamespacesTopic_STATUS
	err = actual.AssignProperties_From_NamespacesTopic_STATUS(&other)
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

func Test_NamespacesTopic_STATUS_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of NamespacesTopic_STATUS via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForNamespacesTopic_STATUS, NamespacesTopic_STATUSGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForNamespacesTopic_STATUS runs a test to see if a specific instance of NamespacesTopic_STATUS round trips to JSON and back losslessly
func RunJSONSerializationTestForNamespacesTopic_STATUS(subject NamespacesTopic_STATUS) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual NamespacesTopic_STATUS
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

// Generator of NamespacesTopic_STATUS instances for property testing - lazily instantiated by
// NamespacesTopic_STATUSGenerator()
var namespacesTopic_STATUSGenerator gopter.Gen

// NamespacesTopic_STATUSGenerator returns a generator of NamespacesTopic_STATUS instances for property testing.
// We first initialize namespacesTopic_STATUSGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func NamespacesTopic_STATUSGenerator() gopter.Gen {
	if namespacesTopic_STATUSGenerator != nil {
		return namespacesTopic_STATUSGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForNamespacesTopic_STATUS(generators)
	namespacesTopic_STATUSGenerator = gen.Struct(reflect.TypeOf(NamespacesTopic_STATUS{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForNamespacesTopic_STATUS(generators)
	AddRelatedPropertyGeneratorsForNamespacesTopic_STATUS(generators)
	namespacesTopic_STATUSGenerator = gen.Struct(reflect.TypeOf(NamespacesTopic_STATUS{}), generators)

	return namespacesTopic_STATUSGenerator
}

// AddIndependentPropertyGeneratorsForNamespacesTopic_STATUS is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForNamespacesTopic_STATUS(gens map[string]gopter.Gen) {
	gens["AccessedAt"] = gen.PtrOf(gen.AlphaString())
	gens["AutoDeleteOnIdle"] = gen.PtrOf(gen.AlphaString())
	gens["CreatedAt"] = gen.PtrOf(gen.AlphaString())
	gens["DefaultMessageTimeToLive"] = gen.PtrOf(gen.AlphaString())
	gens["DuplicateDetectionHistoryTimeWindow"] = gen.PtrOf(gen.AlphaString())
	gens["EnableBatchedOperations"] = gen.PtrOf(gen.Bool())
	gens["EnableExpress"] = gen.PtrOf(gen.Bool())
	gens["EnablePartitioning"] = gen.PtrOf(gen.Bool())
	gens["Id"] = gen.PtrOf(gen.AlphaString())
	gens["Location"] = gen.PtrOf(gen.AlphaString())
	gens["MaxMessageSizeInKilobytes"] = gen.PtrOf(gen.Int())
	gens["MaxSizeInMegabytes"] = gen.PtrOf(gen.Int())
	gens["Name"] = gen.PtrOf(gen.AlphaString())
	gens["RequiresDuplicateDetection"] = gen.PtrOf(gen.Bool())
	gens["SizeInBytes"] = gen.PtrOf(gen.Int())
	gens["Status"] = gen.PtrOf(gen.OneConstOf(
		EntityStatus_STATUS_Active,
		EntityStatus_STATUS_Creating,
		EntityStatus_STATUS_Deleting,
		EntityStatus_STATUS_Disabled,
		EntityStatus_STATUS_ReceiveDisabled,
		EntityStatus_STATUS_Renaming,
		EntityStatus_STATUS_Restoring,
		EntityStatus_STATUS_SendDisabled,
		EntityStatus_STATUS_Unknown))
	gens["SubscriptionCount"] = gen.PtrOf(gen.Int())
	gens["SupportOrdering"] = gen.PtrOf(gen.Bool())
	gens["Type"] = gen.PtrOf(gen.AlphaString())
	gens["UpdatedAt"] = gen.PtrOf(gen.AlphaString())
}

// AddRelatedPropertyGeneratorsForNamespacesTopic_STATUS is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForNamespacesTopic_STATUS(gens map[string]gopter.Gen) {
	gens["CountDetails"] = gen.PtrOf(MessageCountDetails_STATUSGenerator())
	gens["SystemData"] = gen.PtrOf(SystemData_STATUSGenerator())
}

func Test_NamespacesTopic_Spec_WhenPropertiesConverted_RoundTripsWithoutLoss(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip from NamespacesTopic_Spec to NamespacesTopic_Spec via AssignProperties_To_NamespacesTopic_Spec & AssignProperties_From_NamespacesTopic_Spec returns original",
		prop.ForAll(RunPropertyAssignmentTestForNamespacesTopic_Spec, NamespacesTopic_SpecGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(false, 240, os.Stdout))
}

// RunPropertyAssignmentTestForNamespacesTopic_Spec tests if a specific instance of NamespacesTopic_Spec can be assigned to storage and back losslessly
func RunPropertyAssignmentTestForNamespacesTopic_Spec(subject NamespacesTopic_Spec) string {
	// Copy subject to make sure assignment doesn't modify it
	copied := subject.DeepCopy()

	// Use AssignPropertiesTo() for the first stage of conversion
	var other storage.NamespacesTopic_Spec
	err := copied.AssignProperties_To_NamespacesTopic_Spec(&other)
	if err != nil {
		return err.Error()
	}

	// Use AssignPropertiesFrom() to convert back to our original type
	var actual NamespacesTopic_Spec
	err = actual.AssignProperties_From_NamespacesTopic_Spec(&other)
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

func Test_NamespacesTopic_Spec_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of NamespacesTopic_Spec via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForNamespacesTopic_Spec, NamespacesTopic_SpecGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForNamespacesTopic_Spec runs a test to see if a specific instance of NamespacesTopic_Spec round trips to JSON and back losslessly
func RunJSONSerializationTestForNamespacesTopic_Spec(subject NamespacesTopic_Spec) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual NamespacesTopic_Spec
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

// Generator of NamespacesTopic_Spec instances for property testing - lazily instantiated by
// NamespacesTopic_SpecGenerator()
var namespacesTopic_SpecGenerator gopter.Gen

// NamespacesTopic_SpecGenerator returns a generator of NamespacesTopic_Spec instances for property testing.
// We first initialize namespacesTopic_SpecGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func NamespacesTopic_SpecGenerator() gopter.Gen {
	if namespacesTopic_SpecGenerator != nil {
		return namespacesTopic_SpecGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForNamespacesTopic_Spec(generators)
	namespacesTopic_SpecGenerator = gen.Struct(reflect.TypeOf(NamespacesTopic_Spec{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForNamespacesTopic_Spec(generators)
	AddRelatedPropertyGeneratorsForNamespacesTopic_Spec(generators)
	namespacesTopic_SpecGenerator = gen.Struct(reflect.TypeOf(NamespacesTopic_Spec{}), generators)

	return namespacesTopic_SpecGenerator
}

// AddIndependentPropertyGeneratorsForNamespacesTopic_Spec is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForNamespacesTopic_Spec(gens map[string]gopter.Gen) {
	gens["AutoDeleteOnIdle"] = gen.PtrOf(gen.AlphaString())
	gens["AzureName"] = gen.AlphaString()
	gens["DefaultMessageTimeToLive"] = gen.PtrOf(gen.AlphaString())
	gens["DuplicateDetectionHistoryTimeWindow"] = gen.PtrOf(gen.AlphaString())
	gens["EnableBatchedOperations"] = gen.PtrOf(gen.Bool())
	gens["EnableExpress"] = gen.PtrOf(gen.Bool())
	gens["EnablePartitioning"] = gen.PtrOf(gen.Bool())
	gens["MaxMessageSizeInKilobytes"] = gen.PtrOf(gen.Int())
	gens["MaxSizeInMegabytes"] = gen.PtrOf(gen.Int())
	gens["RequiresDuplicateDetection"] = gen.PtrOf(gen.Bool())
	gens["SupportOrdering"] = gen.PtrOf(gen.Bool())
}

// AddRelatedPropertyGeneratorsForNamespacesTopic_Spec is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForNamespacesTopic_Spec(gens map[string]gopter.Gen) {
	gens["OperatorSpec"] = gen.PtrOf(NamespacesTopicOperatorSpecGenerator())
}
