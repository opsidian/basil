// Copyright (c) 2017 Opsidian Ltd.
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at http://mozilla.org/MPL/2.0/.

package multierror

type Error struct {
	errors    []error
	formatter Formatter
}

func (e *Error) WithFormatter(formatter Formatter) *Error {
	if e == nil {
		return nil
	}

	e.formatter = formatter
	return e
}

func (e *Error) Err() error {
	if e == nil {
		return nil
	}

	if len(e.errors) == 0 {
		return nil
	}

	return e
}

func (e *Error) Error() string {
	formatter := e.formatter
	if formatter == nil {
		formatter = FormatterFunc(DefaultFormatter)
	}

	return formatter.FormatMultiError(e.errors)
}

func (e *Error) Errors() []error {
	if e == nil {
		return nil
	}

	return e.errors
}

func (e *Error) Append(err error) *Error {
	if err == nil {
		return e
	}

	if e == nil {
		return &Error{errors: []error{err}}
	}

	if me, ok := err.(*Error); ok {
		e.errors = append(e.errors, me.Errors()...)
	} else {
		e.errors = append(e.errors, err)
	}

	return e
}
