//go:build !ignore_autogenerated

/*
Copyright (c) Microsoft Corporation.
Licensed under the MIT license.
*/

// Code generated by controller-gen. DO NOT EDIT.

package v1api20220615

import (
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime/conditions"
	"k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Actions) DeepCopyInto(out *Actions) {
	*out = *in
	if in.ActionGroupsReferences != nil {
		in, out := &in.ActionGroupsReferences, &out.ActionGroupsReferences
		*out = make([]genruntime.ResourceReference, len(*in))
		copy(*out, *in)
	}
	if in.CustomProperties != nil {
		in, out := &in.CustomProperties, &out.CustomProperties
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Actions.
func (in *Actions) DeepCopy() *Actions {
	if in == nil {
		return nil
	}
	out := new(Actions)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Actions_STATUS) DeepCopyInto(out *Actions_STATUS) {
	*out = *in
	if in.ActionGroups != nil {
		in, out := &in.ActionGroups, &out.ActionGroups
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.CustomProperties != nil {
		in, out := &in.CustomProperties, &out.CustomProperties
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Actions_STATUS.
func (in *Actions_STATUS) DeepCopy() *Actions_STATUS {
	if in == nil {
		return nil
	}
	out := new(Actions_STATUS)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Condition) DeepCopyInto(out *Condition) {
	*out = *in
	if in.Dimensions != nil {
		in, out := &in.Dimensions, &out.Dimensions
		*out = make([]Dimension, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.FailingPeriods != nil {
		in, out := &in.FailingPeriods, &out.FailingPeriods
		*out = new(Condition_FailingPeriods)
		(*in).DeepCopyInto(*out)
	}
	if in.MetricMeasureColumn != nil {
		in, out := &in.MetricMeasureColumn, &out.MetricMeasureColumn
		*out = new(string)
		**out = **in
	}
	if in.MetricName != nil {
		in, out := &in.MetricName, &out.MetricName
		*out = new(string)
		**out = **in
	}
	if in.Operator != nil {
		in, out := &in.Operator, &out.Operator
		*out = new(Condition_Operator)
		**out = **in
	}
	if in.Query != nil {
		in, out := &in.Query, &out.Query
		*out = new(string)
		**out = **in
	}
	if in.ResourceIdColumnReference != nil {
		in, out := &in.ResourceIdColumnReference, &out.ResourceIdColumnReference
		*out = new(genruntime.ResourceReference)
		**out = **in
	}
	if in.Threshold != nil {
		in, out := &in.Threshold, &out.Threshold
		*out = new(float64)
		**out = **in
	}
	if in.TimeAggregation != nil {
		in, out := &in.TimeAggregation, &out.TimeAggregation
		*out = new(Condition_TimeAggregation)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Condition.
func (in *Condition) DeepCopy() *Condition {
	if in == nil {
		return nil
	}
	out := new(Condition)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Condition_FailingPeriods) DeepCopyInto(out *Condition_FailingPeriods) {
	*out = *in
	if in.MinFailingPeriodsToAlert != nil {
		in, out := &in.MinFailingPeriodsToAlert, &out.MinFailingPeriodsToAlert
		*out = new(int)
		**out = **in
	}
	if in.NumberOfEvaluationPeriods != nil {
		in, out := &in.NumberOfEvaluationPeriods, &out.NumberOfEvaluationPeriods
		*out = new(int)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Condition_FailingPeriods.
func (in *Condition_FailingPeriods) DeepCopy() *Condition_FailingPeriods {
	if in == nil {
		return nil
	}
	out := new(Condition_FailingPeriods)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Condition_FailingPeriods_STATUS) DeepCopyInto(out *Condition_FailingPeriods_STATUS) {
	*out = *in
	if in.MinFailingPeriodsToAlert != nil {
		in, out := &in.MinFailingPeriodsToAlert, &out.MinFailingPeriodsToAlert
		*out = new(int)
		**out = **in
	}
	if in.NumberOfEvaluationPeriods != nil {
		in, out := &in.NumberOfEvaluationPeriods, &out.NumberOfEvaluationPeriods
		*out = new(int)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Condition_FailingPeriods_STATUS.
func (in *Condition_FailingPeriods_STATUS) DeepCopy() *Condition_FailingPeriods_STATUS {
	if in == nil {
		return nil
	}
	out := new(Condition_FailingPeriods_STATUS)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Condition_STATUS) DeepCopyInto(out *Condition_STATUS) {
	*out = *in
	if in.Dimensions != nil {
		in, out := &in.Dimensions, &out.Dimensions
		*out = make([]Dimension_STATUS, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.FailingPeriods != nil {
		in, out := &in.FailingPeriods, &out.FailingPeriods
		*out = new(Condition_FailingPeriods_STATUS)
		(*in).DeepCopyInto(*out)
	}
	if in.MetricMeasureColumn != nil {
		in, out := &in.MetricMeasureColumn, &out.MetricMeasureColumn
		*out = new(string)
		**out = **in
	}
	if in.MetricName != nil {
		in, out := &in.MetricName, &out.MetricName
		*out = new(string)
		**out = **in
	}
	if in.Operator != nil {
		in, out := &in.Operator, &out.Operator
		*out = new(Condition_Operator_STATUS)
		**out = **in
	}
	if in.Query != nil {
		in, out := &in.Query, &out.Query
		*out = new(string)
		**out = **in
	}
	if in.ResourceIdColumn != nil {
		in, out := &in.ResourceIdColumn, &out.ResourceIdColumn
		*out = new(string)
		**out = **in
	}
	if in.Threshold != nil {
		in, out := &in.Threshold, &out.Threshold
		*out = new(float64)
		**out = **in
	}
	if in.TimeAggregation != nil {
		in, out := &in.TimeAggregation, &out.TimeAggregation
		*out = new(Condition_TimeAggregation_STATUS)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Condition_STATUS.
func (in *Condition_STATUS) DeepCopy() *Condition_STATUS {
	if in == nil {
		return nil
	}
	out := new(Condition_STATUS)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Dimension) DeepCopyInto(out *Dimension) {
	*out = *in
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
	if in.Operator != nil {
		in, out := &in.Operator, &out.Operator
		*out = new(Dimension_Operator)
		**out = **in
	}
	if in.Values != nil {
		in, out := &in.Values, &out.Values
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Dimension.
func (in *Dimension) DeepCopy() *Dimension {
	if in == nil {
		return nil
	}
	out := new(Dimension)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Dimension_STATUS) DeepCopyInto(out *Dimension_STATUS) {
	*out = *in
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
	if in.Operator != nil {
		in, out := &in.Operator, &out.Operator
		*out = new(Dimension_Operator_STATUS)
		**out = **in
	}
	if in.Values != nil {
		in, out := &in.Values, &out.Values
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Dimension_STATUS.
func (in *Dimension_STATUS) DeepCopy() *Dimension_STATUS {
	if in == nil {
		return nil
	}
	out := new(Dimension_STATUS)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *HeaderField) DeepCopyInto(out *HeaderField) {
	*out = *in
	if in.Key != nil {
		in, out := &in.Key, &out.Key
		*out = new(string)
		**out = **in
	}
	if in.Value != nil {
		in, out := &in.Value, &out.Value
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HeaderField.
func (in *HeaderField) DeepCopy() *HeaderField {
	if in == nil {
		return nil
	}
	out := new(HeaderField)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *HeaderField_STATUS) DeepCopyInto(out *HeaderField_STATUS) {
	*out = *in
	if in.Key != nil {
		in, out := &in.Key, &out.Key
		*out = new(string)
		**out = **in
	}
	if in.Value != nil {
		in, out := &in.Value, &out.Value
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HeaderField_STATUS.
func (in *HeaderField_STATUS) DeepCopy() *HeaderField_STATUS {
	if in == nil {
		return nil
	}
	out := new(HeaderField_STATUS)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ScheduledQueryRule) DeepCopyInto(out *ScheduledQueryRule) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ScheduledQueryRule.
func (in *ScheduledQueryRule) DeepCopy() *ScheduledQueryRule {
	if in == nil {
		return nil
	}
	out := new(ScheduledQueryRule)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ScheduledQueryRule) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ScheduledQueryRuleCriteria) DeepCopyInto(out *ScheduledQueryRuleCriteria) {
	*out = *in
	if in.AllOf != nil {
		in, out := &in.AllOf, &out.AllOf
		*out = make([]Condition, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ScheduledQueryRuleCriteria.
func (in *ScheduledQueryRuleCriteria) DeepCopy() *ScheduledQueryRuleCriteria {
	if in == nil {
		return nil
	}
	out := new(ScheduledQueryRuleCriteria)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ScheduledQueryRuleCriteria_STATUS) DeepCopyInto(out *ScheduledQueryRuleCriteria_STATUS) {
	*out = *in
	if in.AllOf != nil {
		in, out := &in.AllOf, &out.AllOf
		*out = make([]Condition_STATUS, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ScheduledQueryRuleCriteria_STATUS.
func (in *ScheduledQueryRuleCriteria_STATUS) DeepCopy() *ScheduledQueryRuleCriteria_STATUS {
	if in == nil {
		return nil
	}
	out := new(ScheduledQueryRuleCriteria_STATUS)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ScheduledQueryRuleList) DeepCopyInto(out *ScheduledQueryRuleList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]ScheduledQueryRule, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ScheduledQueryRuleList.
func (in *ScheduledQueryRuleList) DeepCopy() *ScheduledQueryRuleList {
	if in == nil {
		return nil
	}
	out := new(ScheduledQueryRuleList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ScheduledQueryRuleList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ScheduledQueryRule_STATUS) DeepCopyInto(out *ScheduledQueryRule_STATUS) {
	*out = *in
	if in.Actions != nil {
		in, out := &in.Actions, &out.Actions
		*out = new(Actions_STATUS)
		(*in).DeepCopyInto(*out)
	}
	if in.AutoMitigate != nil {
		in, out := &in.AutoMitigate, &out.AutoMitigate
		*out = new(bool)
		**out = **in
	}
	if in.CheckWorkspaceAlertsStorageConfigured != nil {
		in, out := &in.CheckWorkspaceAlertsStorageConfigured, &out.CheckWorkspaceAlertsStorageConfigured
		*out = new(bool)
		**out = **in
	}
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]conditions.Condition, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.CreatedWithApiVersion != nil {
		in, out := &in.CreatedWithApiVersion, &out.CreatedWithApiVersion
		*out = new(string)
		**out = **in
	}
	if in.Criteria != nil {
		in, out := &in.Criteria, &out.Criteria
		*out = new(ScheduledQueryRuleCriteria_STATUS)
		(*in).DeepCopyInto(*out)
	}
	if in.Description != nil {
		in, out := &in.Description, &out.Description
		*out = new(string)
		**out = **in
	}
	if in.DisplayName != nil {
		in, out := &in.DisplayName, &out.DisplayName
		*out = new(string)
		**out = **in
	}
	if in.Enabled != nil {
		in, out := &in.Enabled, &out.Enabled
		*out = new(bool)
		**out = **in
	}
	if in.Etag != nil {
		in, out := &in.Etag, &out.Etag
		*out = new(string)
		**out = **in
	}
	if in.EvaluationFrequency != nil {
		in, out := &in.EvaluationFrequency, &out.EvaluationFrequency
		*out = new(string)
		**out = **in
	}
	if in.Id != nil {
		in, out := &in.Id, &out.Id
		*out = new(string)
		**out = **in
	}
	if in.IsLegacyLogAnalyticsRule != nil {
		in, out := &in.IsLegacyLogAnalyticsRule, &out.IsLegacyLogAnalyticsRule
		*out = new(bool)
		**out = **in
	}
	if in.IsWorkspaceAlertsStorageConfigured != nil {
		in, out := &in.IsWorkspaceAlertsStorageConfigured, &out.IsWorkspaceAlertsStorageConfigured
		*out = new(bool)
		**out = **in
	}
	if in.Kind != nil {
		in, out := &in.Kind, &out.Kind
		*out = new(ScheduledQueryRule_Kind_STATUS)
		**out = **in
	}
	if in.Location != nil {
		in, out := &in.Location, &out.Location
		*out = new(string)
		**out = **in
	}
	if in.MuteActionsDuration != nil {
		in, out := &in.MuteActionsDuration, &out.MuteActionsDuration
		*out = new(string)
		**out = **in
	}
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
	if in.OverrideQueryTimeRange != nil {
		in, out := &in.OverrideQueryTimeRange, &out.OverrideQueryTimeRange
		*out = new(string)
		**out = **in
	}
	if in.Scopes != nil {
		in, out := &in.Scopes, &out.Scopes
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.Severity != nil {
		in, out := &in.Severity, &out.Severity
		*out = new(ScheduledQueryRuleProperties_Severity_STATUS)
		**out = **in
	}
	if in.SkipQueryValidation != nil {
		in, out := &in.SkipQueryValidation, &out.SkipQueryValidation
		*out = new(bool)
		**out = **in
	}
	if in.SystemData != nil {
		in, out := &in.SystemData, &out.SystemData
		*out = new(SystemData_STATUS)
		(*in).DeepCopyInto(*out)
	}
	if in.Tags != nil {
		in, out := &in.Tags, &out.Tags
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.TargetResourceTypes != nil {
		in, out := &in.TargetResourceTypes, &out.TargetResourceTypes
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.Type != nil {
		in, out := &in.Type, &out.Type
		*out = new(string)
		**out = **in
	}
	if in.WindowSize != nil {
		in, out := &in.WindowSize, &out.WindowSize
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ScheduledQueryRule_STATUS.
func (in *ScheduledQueryRule_STATUS) DeepCopy() *ScheduledQueryRule_STATUS {
	if in == nil {
		return nil
	}
	out := new(ScheduledQueryRule_STATUS)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ScheduledQueryRule_Spec) DeepCopyInto(out *ScheduledQueryRule_Spec) {
	*out = *in
	if in.Actions != nil {
		in, out := &in.Actions, &out.Actions
		*out = new(Actions)
		(*in).DeepCopyInto(*out)
	}
	if in.AutoMitigate != nil {
		in, out := &in.AutoMitigate, &out.AutoMitigate
		*out = new(bool)
		**out = **in
	}
	if in.CheckWorkspaceAlertsStorageConfigured != nil {
		in, out := &in.CheckWorkspaceAlertsStorageConfigured, &out.CheckWorkspaceAlertsStorageConfigured
		*out = new(bool)
		**out = **in
	}
	if in.Criteria != nil {
		in, out := &in.Criteria, &out.Criteria
		*out = new(ScheduledQueryRuleCriteria)
		(*in).DeepCopyInto(*out)
	}
	if in.Description != nil {
		in, out := &in.Description, &out.Description
		*out = new(string)
		**out = **in
	}
	if in.DisplayName != nil {
		in, out := &in.DisplayName, &out.DisplayName
		*out = new(string)
		**out = **in
	}
	if in.Enabled != nil {
		in, out := &in.Enabled, &out.Enabled
		*out = new(bool)
		**out = **in
	}
	if in.EvaluationFrequency != nil {
		in, out := &in.EvaluationFrequency, &out.EvaluationFrequency
		*out = new(string)
		**out = **in
	}
	if in.Kind != nil {
		in, out := &in.Kind, &out.Kind
		*out = new(ScheduledQueryRule_Kind_Spec)
		**out = **in
	}
	if in.Location != nil {
		in, out := &in.Location, &out.Location
		*out = new(string)
		**out = **in
	}
	if in.MuteActionsDuration != nil {
		in, out := &in.MuteActionsDuration, &out.MuteActionsDuration
		*out = new(string)
		**out = **in
	}
	if in.OverrideQueryTimeRange != nil {
		in, out := &in.OverrideQueryTimeRange, &out.OverrideQueryTimeRange
		*out = new(string)
		**out = **in
	}
	if in.Owner != nil {
		in, out := &in.Owner, &out.Owner
		*out = new(genruntime.KnownResourceReference)
		**out = **in
	}
	if in.ScopesReferences != nil {
		in, out := &in.ScopesReferences, &out.ScopesReferences
		*out = make([]genruntime.ResourceReference, len(*in))
		copy(*out, *in)
	}
	if in.Severity != nil {
		in, out := &in.Severity, &out.Severity
		*out = new(ScheduledQueryRuleProperties_Severity)
		**out = **in
	}
	if in.SkipQueryValidation != nil {
		in, out := &in.SkipQueryValidation, &out.SkipQueryValidation
		*out = new(bool)
		**out = **in
	}
	if in.Tags != nil {
		in, out := &in.Tags, &out.Tags
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.TargetResourceTypes != nil {
		in, out := &in.TargetResourceTypes, &out.TargetResourceTypes
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.WindowSize != nil {
		in, out := &in.WindowSize, &out.WindowSize
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ScheduledQueryRule_Spec.
func (in *ScheduledQueryRule_Spec) DeepCopy() *ScheduledQueryRule_Spec {
	if in == nil {
		return nil
	}
	out := new(ScheduledQueryRule_Spec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SystemData_STATUS) DeepCopyInto(out *SystemData_STATUS) {
	*out = *in
	if in.CreatedAt != nil {
		in, out := &in.CreatedAt, &out.CreatedAt
		*out = new(string)
		**out = **in
	}
	if in.CreatedBy != nil {
		in, out := &in.CreatedBy, &out.CreatedBy
		*out = new(string)
		**out = **in
	}
	if in.CreatedByType != nil {
		in, out := &in.CreatedByType, &out.CreatedByType
		*out = new(SystemData_CreatedByType_STATUS)
		**out = **in
	}
	if in.LastModifiedAt != nil {
		in, out := &in.LastModifiedAt, &out.LastModifiedAt
		*out = new(string)
		**out = **in
	}
	if in.LastModifiedBy != nil {
		in, out := &in.LastModifiedBy, &out.LastModifiedBy
		*out = new(string)
		**out = **in
	}
	if in.LastModifiedByType != nil {
		in, out := &in.LastModifiedByType, &out.LastModifiedByType
		*out = new(SystemData_LastModifiedByType_STATUS)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SystemData_STATUS.
func (in *SystemData_STATUS) DeepCopy() *SystemData_STATUS {
	if in == nil {
		return nil
	}
	out := new(SystemData_STATUS)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *WebTestGeolocation) DeepCopyInto(out *WebTestGeolocation) {
	*out = *in
	if in.Id != nil {
		in, out := &in.Id, &out.Id
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new WebTestGeolocation.
func (in *WebTestGeolocation) DeepCopy() *WebTestGeolocation {
	if in == nil {
		return nil
	}
	out := new(WebTestGeolocation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *WebTestGeolocation_STATUS) DeepCopyInto(out *WebTestGeolocation_STATUS) {
	*out = *in
	if in.Id != nil {
		in, out := &in.Id, &out.Id
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new WebTestGeolocation_STATUS.
func (in *WebTestGeolocation_STATUS) DeepCopy() *WebTestGeolocation_STATUS {
	if in == nil {
		return nil
	}
	out := new(WebTestGeolocation_STATUS)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *WebTestProperties_Configuration) DeepCopyInto(out *WebTestProperties_Configuration) {
	*out = *in
	if in.WebTest != nil {
		in, out := &in.WebTest, &out.WebTest
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new WebTestProperties_Configuration.
func (in *WebTestProperties_Configuration) DeepCopy() *WebTestProperties_Configuration {
	if in == nil {
		return nil
	}
	out := new(WebTestProperties_Configuration)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *WebTestProperties_Configuration_STATUS) DeepCopyInto(out *WebTestProperties_Configuration_STATUS) {
	*out = *in
	if in.WebTest != nil {
		in, out := &in.WebTest, &out.WebTest
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new WebTestProperties_Configuration_STATUS.
func (in *WebTestProperties_Configuration_STATUS) DeepCopy() *WebTestProperties_Configuration_STATUS {
	if in == nil {
		return nil
	}
	out := new(WebTestProperties_Configuration_STATUS)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *WebTestProperties_Request) DeepCopyInto(out *WebTestProperties_Request) {
	*out = *in
	if in.FollowRedirects != nil {
		in, out := &in.FollowRedirects, &out.FollowRedirects
		*out = new(bool)
		**out = **in
	}
	if in.Headers != nil {
		in, out := &in.Headers, &out.Headers
		*out = make([]HeaderField, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.HttpVerb != nil {
		in, out := &in.HttpVerb, &out.HttpVerb
		*out = new(string)
		**out = **in
	}
	if in.ParseDependentRequests != nil {
		in, out := &in.ParseDependentRequests, &out.ParseDependentRequests
		*out = new(bool)
		**out = **in
	}
	if in.RequestBody != nil {
		in, out := &in.RequestBody, &out.RequestBody
		*out = new(string)
		**out = **in
	}
	if in.RequestUrl != nil {
		in, out := &in.RequestUrl, &out.RequestUrl
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new WebTestProperties_Request.
func (in *WebTestProperties_Request) DeepCopy() *WebTestProperties_Request {
	if in == nil {
		return nil
	}
	out := new(WebTestProperties_Request)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *WebTestProperties_Request_STATUS) DeepCopyInto(out *WebTestProperties_Request_STATUS) {
	*out = *in
	if in.FollowRedirects != nil {
		in, out := &in.FollowRedirects, &out.FollowRedirects
		*out = new(bool)
		**out = **in
	}
	if in.Headers != nil {
		in, out := &in.Headers, &out.Headers
		*out = make([]HeaderField_STATUS, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.HttpVerb != nil {
		in, out := &in.HttpVerb, &out.HttpVerb
		*out = new(string)
		**out = **in
	}
	if in.ParseDependentRequests != nil {
		in, out := &in.ParseDependentRequests, &out.ParseDependentRequests
		*out = new(bool)
		**out = **in
	}
	if in.RequestBody != nil {
		in, out := &in.RequestBody, &out.RequestBody
		*out = new(string)
		**out = **in
	}
	if in.RequestUrl != nil {
		in, out := &in.RequestUrl, &out.RequestUrl
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new WebTestProperties_Request_STATUS.
func (in *WebTestProperties_Request_STATUS) DeepCopy() *WebTestProperties_Request_STATUS {
	if in == nil {
		return nil
	}
	out := new(WebTestProperties_Request_STATUS)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *WebTestProperties_ValidationRules) DeepCopyInto(out *WebTestProperties_ValidationRules) {
	*out = *in
	if in.ContentValidation != nil {
		in, out := &in.ContentValidation, &out.ContentValidation
		*out = new(WebTestProperties_ValidationRules_ContentValidation)
		(*in).DeepCopyInto(*out)
	}
	if in.ExpectedHttpStatusCode != nil {
		in, out := &in.ExpectedHttpStatusCode, &out.ExpectedHttpStatusCode
		*out = new(int)
		**out = **in
	}
	if in.IgnoreHttpStatusCode != nil {
		in, out := &in.IgnoreHttpStatusCode, &out.IgnoreHttpStatusCode
		*out = new(bool)
		**out = **in
	}
	if in.SSLCertRemainingLifetimeCheck != nil {
		in, out := &in.SSLCertRemainingLifetimeCheck, &out.SSLCertRemainingLifetimeCheck
		*out = new(int)
		**out = **in
	}
	if in.SSLCheck != nil {
		in, out := &in.SSLCheck, &out.SSLCheck
		*out = new(bool)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new WebTestProperties_ValidationRules.
func (in *WebTestProperties_ValidationRules) DeepCopy() *WebTestProperties_ValidationRules {
	if in == nil {
		return nil
	}
	out := new(WebTestProperties_ValidationRules)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *WebTestProperties_ValidationRules_ContentValidation) DeepCopyInto(out *WebTestProperties_ValidationRules_ContentValidation) {
	*out = *in
	if in.ContentMatch != nil {
		in, out := &in.ContentMatch, &out.ContentMatch
		*out = new(string)
		**out = **in
	}
	if in.IgnoreCase != nil {
		in, out := &in.IgnoreCase, &out.IgnoreCase
		*out = new(bool)
		**out = **in
	}
	if in.PassIfTextFound != nil {
		in, out := &in.PassIfTextFound, &out.PassIfTextFound
		*out = new(bool)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new WebTestProperties_ValidationRules_ContentValidation.
func (in *WebTestProperties_ValidationRules_ContentValidation) DeepCopy() *WebTestProperties_ValidationRules_ContentValidation {
	if in == nil {
		return nil
	}
	out := new(WebTestProperties_ValidationRules_ContentValidation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *WebTestProperties_ValidationRules_ContentValidation_STATUS) DeepCopyInto(out *WebTestProperties_ValidationRules_ContentValidation_STATUS) {
	*out = *in
	if in.ContentMatch != nil {
		in, out := &in.ContentMatch, &out.ContentMatch
		*out = new(string)
		**out = **in
	}
	if in.IgnoreCase != nil {
		in, out := &in.IgnoreCase, &out.IgnoreCase
		*out = new(bool)
		**out = **in
	}
	if in.PassIfTextFound != nil {
		in, out := &in.PassIfTextFound, &out.PassIfTextFound
		*out = new(bool)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new WebTestProperties_ValidationRules_ContentValidation_STATUS.
func (in *WebTestProperties_ValidationRules_ContentValidation_STATUS) DeepCopy() *WebTestProperties_ValidationRules_ContentValidation_STATUS {
	if in == nil {
		return nil
	}
	out := new(WebTestProperties_ValidationRules_ContentValidation_STATUS)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *WebTestProperties_ValidationRules_STATUS) DeepCopyInto(out *WebTestProperties_ValidationRules_STATUS) {
	*out = *in
	if in.ContentValidation != nil {
		in, out := &in.ContentValidation, &out.ContentValidation
		*out = new(WebTestProperties_ValidationRules_ContentValidation_STATUS)
		(*in).DeepCopyInto(*out)
	}
	if in.ExpectedHttpStatusCode != nil {
		in, out := &in.ExpectedHttpStatusCode, &out.ExpectedHttpStatusCode
		*out = new(int)
		**out = **in
	}
	if in.IgnoreHttpStatusCode != nil {
		in, out := &in.IgnoreHttpStatusCode, &out.IgnoreHttpStatusCode
		*out = new(bool)
		**out = **in
	}
	if in.SSLCertRemainingLifetimeCheck != nil {
		in, out := &in.SSLCertRemainingLifetimeCheck, &out.SSLCertRemainingLifetimeCheck
		*out = new(int)
		**out = **in
	}
	if in.SSLCheck != nil {
		in, out := &in.SSLCheck, &out.SSLCheck
		*out = new(bool)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new WebTestProperties_ValidationRules_STATUS.
func (in *WebTestProperties_ValidationRules_STATUS) DeepCopy() *WebTestProperties_ValidationRules_STATUS {
	if in == nil {
		return nil
	}
	out := new(WebTestProperties_ValidationRules_STATUS)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Webtest) DeepCopyInto(out *Webtest) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Webtest.
func (in *Webtest) DeepCopy() *Webtest {
	if in == nil {
		return nil
	}
	out := new(Webtest)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Webtest) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *WebtestList) DeepCopyInto(out *WebtestList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Webtest, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new WebtestList.
func (in *WebtestList) DeepCopy() *WebtestList {
	if in == nil {
		return nil
	}
	out := new(WebtestList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *WebtestList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Webtest_STATUS) DeepCopyInto(out *Webtest_STATUS) {
	*out = *in
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]conditions.Condition, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Configuration != nil {
		in, out := &in.Configuration, &out.Configuration
		*out = new(WebTestProperties_Configuration_STATUS)
		(*in).DeepCopyInto(*out)
	}
	if in.Description != nil {
		in, out := &in.Description, &out.Description
		*out = new(string)
		**out = **in
	}
	if in.Enabled != nil {
		in, out := &in.Enabled, &out.Enabled
		*out = new(bool)
		**out = **in
	}
	if in.Frequency != nil {
		in, out := &in.Frequency, &out.Frequency
		*out = new(int)
		**out = **in
	}
	if in.Id != nil {
		in, out := &in.Id, &out.Id
		*out = new(string)
		**out = **in
	}
	if in.Kind != nil {
		in, out := &in.Kind, &out.Kind
		*out = new(WebTestProperties_Kind_STATUS)
		**out = **in
	}
	if in.Location != nil {
		in, out := &in.Location, &out.Location
		*out = new(string)
		**out = **in
	}
	if in.Locations != nil {
		in, out := &in.Locations, &out.Locations
		*out = make([]WebTestGeolocation_STATUS, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
	if in.PropertiesName != nil {
		in, out := &in.PropertiesName, &out.PropertiesName
		*out = new(string)
		**out = **in
	}
	if in.ProvisioningState != nil {
		in, out := &in.ProvisioningState, &out.ProvisioningState
		*out = new(string)
		**out = **in
	}
	if in.Request != nil {
		in, out := &in.Request, &out.Request
		*out = new(WebTestProperties_Request_STATUS)
		(*in).DeepCopyInto(*out)
	}
	if in.RetryEnabled != nil {
		in, out := &in.RetryEnabled, &out.RetryEnabled
		*out = new(bool)
		**out = **in
	}
	if in.SyntheticMonitorId != nil {
		in, out := &in.SyntheticMonitorId, &out.SyntheticMonitorId
		*out = new(string)
		**out = **in
	}
	if in.Tags != nil {
		in, out := &in.Tags, &out.Tags
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.Timeout != nil {
		in, out := &in.Timeout, &out.Timeout
		*out = new(int)
		**out = **in
	}
	if in.Type != nil {
		in, out := &in.Type, &out.Type
		*out = new(string)
		**out = **in
	}
	if in.ValidationRules != nil {
		in, out := &in.ValidationRules, &out.ValidationRules
		*out = new(WebTestProperties_ValidationRules_STATUS)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Webtest_STATUS.
func (in *Webtest_STATUS) DeepCopy() *Webtest_STATUS {
	if in == nil {
		return nil
	}
	out := new(Webtest_STATUS)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Webtest_Spec) DeepCopyInto(out *Webtest_Spec) {
	*out = *in
	if in.Configuration != nil {
		in, out := &in.Configuration, &out.Configuration
		*out = new(WebTestProperties_Configuration)
		(*in).DeepCopyInto(*out)
	}
	if in.Description != nil {
		in, out := &in.Description, &out.Description
		*out = new(string)
		**out = **in
	}
	if in.Enabled != nil {
		in, out := &in.Enabled, &out.Enabled
		*out = new(bool)
		**out = **in
	}
	if in.Frequency != nil {
		in, out := &in.Frequency, &out.Frequency
		*out = new(int)
		**out = **in
	}
	if in.Kind != nil {
		in, out := &in.Kind, &out.Kind
		*out = new(WebTestProperties_Kind)
		**out = **in
	}
	if in.Location != nil {
		in, out := &in.Location, &out.Location
		*out = new(string)
		**out = **in
	}
	if in.Locations != nil {
		in, out := &in.Locations, &out.Locations
		*out = make([]WebTestGeolocation, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
	if in.Owner != nil {
		in, out := &in.Owner, &out.Owner
		*out = new(genruntime.KnownResourceReference)
		**out = **in
	}
	if in.Request != nil {
		in, out := &in.Request, &out.Request
		*out = new(WebTestProperties_Request)
		(*in).DeepCopyInto(*out)
	}
	if in.RetryEnabled != nil {
		in, out := &in.RetryEnabled, &out.RetryEnabled
		*out = new(bool)
		**out = **in
	}
	if in.SyntheticMonitorId != nil {
		in, out := &in.SyntheticMonitorId, &out.SyntheticMonitorId
		*out = new(string)
		**out = **in
	}
	if in.Tags != nil {
		in, out := &in.Tags, &out.Tags
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.Timeout != nil {
		in, out := &in.Timeout, &out.Timeout
		*out = new(int)
		**out = **in
	}
	if in.ValidationRules != nil {
		in, out := &in.ValidationRules, &out.ValidationRules
		*out = new(WebTestProperties_ValidationRules)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Webtest_Spec.
func (in *Webtest_Spec) DeepCopy() *Webtest_Spec {
	if in == nil {
		return nil
	}
	out := new(Webtest_Spec)
	in.DeepCopyInto(out)
	return out
}
