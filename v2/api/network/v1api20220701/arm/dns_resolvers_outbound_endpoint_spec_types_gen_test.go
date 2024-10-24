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

func Test_DnsResolversOutboundEndpoint_Spec_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of DnsResolversOutboundEndpoint_Spec via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForDnsResolversOutboundEndpoint_Spec, DnsResolversOutboundEndpoint_SpecGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForDnsResolversOutboundEndpoint_Spec runs a test to see if a specific instance of DnsResolversOutboundEndpoint_Spec round trips to JSON and back losslessly
func RunJSONSerializationTestForDnsResolversOutboundEndpoint_Spec(subject DnsResolversOutboundEndpoint_Spec) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual DnsResolversOutboundEndpoint_Spec
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

// Generator of DnsResolversOutboundEndpoint_Spec instances for property testing - lazily instantiated by
// DnsResolversOutboundEndpoint_SpecGenerator()
var dnsResolversOutboundEndpoint_SpecGenerator gopter.Gen

// DnsResolversOutboundEndpoint_SpecGenerator returns a generator of DnsResolversOutboundEndpoint_Spec instances for property testing.
// We first initialize dnsResolversOutboundEndpoint_SpecGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func DnsResolversOutboundEndpoint_SpecGenerator() gopter.Gen {
	if dnsResolversOutboundEndpoint_SpecGenerator != nil {
		return dnsResolversOutboundEndpoint_SpecGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForDnsResolversOutboundEndpoint_Spec(generators)
	dnsResolversOutboundEndpoint_SpecGenerator = gen.Struct(reflect.TypeOf(DnsResolversOutboundEndpoint_Spec{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForDnsResolversOutboundEndpoint_Spec(generators)
	AddRelatedPropertyGeneratorsForDnsResolversOutboundEndpoint_Spec(generators)
	dnsResolversOutboundEndpoint_SpecGenerator = gen.Struct(reflect.TypeOf(DnsResolversOutboundEndpoint_Spec{}), generators)

	return dnsResolversOutboundEndpoint_SpecGenerator
}

// AddIndependentPropertyGeneratorsForDnsResolversOutboundEndpoint_Spec is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForDnsResolversOutboundEndpoint_Spec(gens map[string]gopter.Gen) {
	gens["Location"] = gen.PtrOf(gen.AlphaString())
	gens["Name"] = gen.AlphaString()
	gens["Tags"] = gen.MapOf(
		gen.AlphaString(),
		gen.AlphaString())
}

// AddRelatedPropertyGeneratorsForDnsResolversOutboundEndpoint_Spec is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForDnsResolversOutboundEndpoint_Spec(gens map[string]gopter.Gen) {
	gens["Properties"] = gen.PtrOf(OutboundEndpointPropertiesGenerator())
}

func Test_OutboundEndpointProperties_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 100
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of OutboundEndpointProperties via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForOutboundEndpointProperties, OutboundEndpointPropertiesGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForOutboundEndpointProperties runs a test to see if a specific instance of OutboundEndpointProperties round trips to JSON and back losslessly
func RunJSONSerializationTestForOutboundEndpointProperties(subject OutboundEndpointProperties) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual OutboundEndpointProperties
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

// Generator of OutboundEndpointProperties instances for property testing - lazily instantiated by
// OutboundEndpointPropertiesGenerator()
var outboundEndpointPropertiesGenerator gopter.Gen

// OutboundEndpointPropertiesGenerator returns a generator of OutboundEndpointProperties instances for property testing.
func OutboundEndpointPropertiesGenerator() gopter.Gen {
	if outboundEndpointPropertiesGenerator != nil {
		return outboundEndpointPropertiesGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddRelatedPropertyGeneratorsForOutboundEndpointProperties(generators)
	outboundEndpointPropertiesGenerator = gen.Struct(reflect.TypeOf(OutboundEndpointProperties{}), generators)

	return outboundEndpointPropertiesGenerator
}

// AddRelatedPropertyGeneratorsForOutboundEndpointProperties is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForOutboundEndpointProperties(gens map[string]gopter.Gen) {
	gens["Subnet"] = gen.PtrOf(DnsresolverSubResourceGenerator())
}
