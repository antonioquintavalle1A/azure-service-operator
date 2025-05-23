/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package astmodel

import (
	"github.com/dave/dst"
)

type FakeFunction struct {
	name       string
	Imported   *PackageReferenceSet
	Referenced TypeNameSet
}

func NewFakeFunction(name string) *FakeFunction {
	return &FakeFunction{
		name:     name,
		Imported: NewPackageReferenceSet(),
	}
}

var _ Function = &FakeFunction{}

func (fake *FakeFunction) Name() string {
	return fake.name
}

func (fake *FakeFunction) RequiredPackageReferences() *PackageReferenceSet {
	result := NewPackageReferenceSet()
	result.Merge(fake.Imported)
	return result
}

func (fake *FakeFunction) References() TypeNameSet {
	return fake.Referenced
}

func (fake *FakeFunction) AsFunc(_ *CodeGenerationContext, receiver InternalTypeName) (*dst.FuncDecl, error) {
	panic("implement me")
}

func (fake *FakeFunction) Equals(f Function, _ EqualityOverrides) bool {
	if fake == nil && f == nil {
		return true
	}

	if fake == nil || f == nil {
		return false
	}

	fn, ok := f.(*FakeFunction)
	if !ok {
		return false
	}

	// Check to see if they have the same references
	if !fake.Referenced.Equals(fn.Referenced) {
		return false
	}

	// Check to see if they have the same imports
	if fake.Imported.Length() != fn.Imported.Length() {
		return false
	}

	for imp := range fake.Imported.All() {
		if !fn.Imported.Contains(imp) {
			return false
		}
	}

	return true
}
