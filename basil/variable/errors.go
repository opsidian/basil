// Copyright (c) 2017 Opsidian Ltd.
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at http://mozilla.org/MPL/2.0/.

package variable

import "errors"

// ErrNotDefined is an error when a variable is not defined
var ErrNotDefined = errors.New("variable not defined")

// Errors when expecting a certain variable type
var (
	ErrExpectingAny          = errors.New("was expecting any valid type")
	ErrExpectingArray        = errors.New("was expecting array")
	ErrExpectingBasic        = errors.New("was expecting any basic type")
	ErrExpectingBool         = errors.New("was expecting boolean")
	ErrExpectingFloat        = errors.New("was expecting float")
	ErrExpectingIdentifier   = errors.New("was expecting identifier")
	ErrExpectingInteger      = errors.New("was expecting integer")
	ErrExpectingMap          = errors.New("was expecting map")
	ErrExpectingNumber       = errors.New("was expecting number")
	ErrExpectingStream       = errors.New("was expecting stream")
	ErrExpectingString       = errors.New("was expecting string")
	ErrExpectingStringArray  = errors.New("was expecting string array")
	ErrExpectingTime         = errors.New("was expecting RFC3339 time")
	ErrExpectingTimeDuration = errors.New("was expecting time duration")
	ErrExpectingWithLength   = errors.New("was expecting string, array or map")
)

// TypeErrors contains the type errors for all variable types
var TypeErrors = map[string]error{
	TypeAny:          ErrExpectingAny,
	TypeArray:        ErrExpectingArray,
	TypeBasic:        ErrExpectingBasic,
	TypeBool:         ErrExpectingBool,
	TypeFloat:        ErrExpectingFloat,
	TypeIdentifier:   ErrExpectingIdentifier,
	TypeInteger:      ErrExpectingInteger,
	TypeMap:          ErrExpectingMap,
	TypeNumber:       ErrExpectingNumber,
	TypeStream:       ErrExpectingStream,
	TypeString:       ErrExpectingString,
	TypeStringArray:  ErrExpectingStringArray,
	TypeTime:         ErrExpectingTime,
	TypeTimeDuration: ErrExpectingTimeDuration,
	TypeWithLength:   ErrExpectingWithLength,
}
