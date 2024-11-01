// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package arm

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

func Test_RedisFirewallRuleProperties_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 100
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of RedisFirewallRuleProperties via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForRedisFirewallRuleProperties, RedisFirewallRulePropertiesGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForRedisFirewallRuleProperties runs a test to see if a specific instance of RedisFirewallRuleProperties round trips to JSON and back losslessly
func RunJSONSerializationTestForRedisFirewallRuleProperties(subject RedisFirewallRuleProperties) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual RedisFirewallRuleProperties
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

// Generator of RedisFirewallRuleProperties instances for property testing - lazily instantiated by
// RedisFirewallRulePropertiesGenerator()
var redisFirewallRulePropertiesGenerator gopter.Gen

// RedisFirewallRulePropertiesGenerator returns a generator of RedisFirewallRuleProperties instances for property testing.
func RedisFirewallRulePropertiesGenerator() gopter.Gen {
	if redisFirewallRulePropertiesGenerator != nil {
		return redisFirewallRulePropertiesGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForRedisFirewallRuleProperties(generators)
	redisFirewallRulePropertiesGenerator = gen.Struct(reflect.TypeOf(RedisFirewallRuleProperties{}), generators)

	return redisFirewallRulePropertiesGenerator
}

// AddIndependentPropertyGeneratorsForRedisFirewallRuleProperties is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForRedisFirewallRuleProperties(gens map[string]gopter.Gen) {
	gens["EndIP"] = gen.PtrOf(gen.AlphaString())
	gens["StartIP"] = gen.PtrOf(gen.AlphaString())
}

func Test_RedisFirewallRule_Spec_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of RedisFirewallRule_Spec via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForRedisFirewallRule_Spec, RedisFirewallRule_SpecGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForRedisFirewallRule_Spec runs a test to see if a specific instance of RedisFirewallRule_Spec round trips to JSON and back losslessly
func RunJSONSerializationTestForRedisFirewallRule_Spec(subject RedisFirewallRule_Spec) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual RedisFirewallRule_Spec
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

// Generator of RedisFirewallRule_Spec instances for property testing - lazily instantiated by
// RedisFirewallRule_SpecGenerator()
var redisFirewallRule_SpecGenerator gopter.Gen

// RedisFirewallRule_SpecGenerator returns a generator of RedisFirewallRule_Spec instances for property testing.
// We first initialize redisFirewallRule_SpecGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func RedisFirewallRule_SpecGenerator() gopter.Gen {
	if redisFirewallRule_SpecGenerator != nil {
		return redisFirewallRule_SpecGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForRedisFirewallRule_Spec(generators)
	redisFirewallRule_SpecGenerator = gen.Struct(reflect.TypeOf(RedisFirewallRule_Spec{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForRedisFirewallRule_Spec(generators)
	AddRelatedPropertyGeneratorsForRedisFirewallRule_Spec(generators)
	redisFirewallRule_SpecGenerator = gen.Struct(reflect.TypeOf(RedisFirewallRule_Spec{}), generators)

	return redisFirewallRule_SpecGenerator
}

// AddIndependentPropertyGeneratorsForRedisFirewallRule_Spec is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForRedisFirewallRule_Spec(gens map[string]gopter.Gen) {
	gens["Name"] = gen.AlphaString()
}

// AddRelatedPropertyGeneratorsForRedisFirewallRule_Spec is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForRedisFirewallRule_Spec(gens map[string]gopter.Gen) {
	gens["Properties"] = gen.PtrOf(RedisFirewallRulePropertiesGenerator())
}
