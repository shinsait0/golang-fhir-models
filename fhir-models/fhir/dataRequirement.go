// Copyright 2019 - 2021 The Samply Community
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package fhir

// THIS FILE IS GENERATED BY https://github.com/samply/golang-fhir-models
// PLEASE DO NOT EDIT BY HAND

// DataRequirement is documented here http://hl7.org/fhir/StructureDefinition/DataRequirement
type DataRequirement struct {
	Id                     *string                     `bson:"id,omitempty" json:"id,omitempty"`
	Extension              []Extension                 `bson:"extension,omitempty" json:"extension,omitempty"`
	Type                   string                      `bson:"type" json:"type"`
	Profile                []string                    `bson:"profile,omitempty" json:"profile,omitempty"`
	SubjectCodeableConcept *CodeableConcept            `bson:"subjectCodeableConcept,omitempty" json:"subjectCodeableConcept,omitempty"`
	SubjectReference       *Reference                  `bson:"subjectReference,omitempty" json:"subjectReference,omitempty"`
	MustSupport            []string                    `bson:"mustSupport,omitempty" json:"mustSupport,omitempty"`
	CodeFilter             []DataRequirementCodeFilter `bson:"codeFilter,omitempty" json:"codeFilter,omitempty"`
	DateFilter             []DataRequirementDateFilter `bson:"dateFilter,omitempty" json:"dateFilter,omitempty"`
	Limit                  *int                        `bson:"limit,omitempty" json:"limit,omitempty"`
	Sort                   []DataRequirementSort       `bson:"sort,omitempty" json:"sort,omitempty"`
}
type DataRequirementCodeFilter struct {
	Id          *string     `bson:"id,omitempty" json:"id,omitempty"`
	Extension   []Extension `bson:"extension,omitempty" json:"extension,omitempty"`
	Path        *string     `bson:"path,omitempty" json:"path,omitempty"`
	SearchParam *string     `bson:"searchParam,omitempty" json:"searchParam,omitempty"`
	ValueSet    *string     `bson:"valueSet,omitempty" json:"valueSet,omitempty"`
	Code        []Coding    `bson:"code,omitempty" json:"code,omitempty"`
}
type DataRequirementDateFilter struct {
	Id            *string     `bson:"id,omitempty" json:"id,omitempty"`
	Extension     []Extension `bson:"extension,omitempty" json:"extension,omitempty"`
	Path          *string     `bson:"path,omitempty" json:"path,omitempty"`
	SearchParam   *string     `bson:"searchParam,omitempty" json:"searchParam,omitempty"`
	ValueDateTime *string     `bson:"valueDateTime,omitempty" json:"valueDateTime,omitempty"`
	ValuePeriod   *Period     `bson:"valuePeriod,omitempty" json:"valuePeriod,omitempty"`
	ValueDuration *Duration   `bson:"valueDuration,omitempty" json:"valueDuration,omitempty"`
}
type DataRequirementSort struct {
	Id        *string       `bson:"id,omitempty" json:"id,omitempty"`
	Extension []Extension   `bson:"extension,omitempty" json:"extension,omitempty"`
	Path      string        `bson:"path" json:"path"`
	Direction SortDirection `bson:"direction" json:"direction"`
}
