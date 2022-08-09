// -*- Mode: Go; indent-tabs-mode: t -*-

/*
 * Copyright (C) 2022 Canonical Ltd
 *
 * This program is free software: you can redistribute it and/or modify
 * it under the terms of the GNU General Public License version 3 as
 * published by the Free Software Foundation.
 *
 * This program is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 * GNU General Public License for more details.
 *
 * You should have received a copy of the GNU General Public License
 * along with this program.  If not, see <http://www.gnu.org/licenses/>.
 *
 */

// Package i18n provides an implementation agnostic API to be used in
// both library packages and application packages importing this library.
// The application has the options to:
//
// 1. Initialise this package by implementing the i18n interface
// 2. Leave the package uninitialised, disabling translation
//
// Warning:
//
// None of the i18n functionality may be used during early
// initialisation code such as when defining a package 'const', 'var'
// or 'init()'. This will prevent the application initialising
// the i18n interface before it gets used in packages, and will
// result in translation being disabled until the initialisation
// call is made.

package i18n

// API defines the interface that provides internationalisation
// for packages. An implementation specific instance of interface
// must be provided by the application wanting to enable
// runtime translation.
type MarkerAPI interface {
	G(msgid string) string
	NG(msgid, msgidPlural string, n int) string
}

var instance MarkerAPI

// Initialise must be called before any of the API calls will
// have translation enabled. It is a valid use case to not
// call Initialise, and still use the API.
func Initialise(markers MarkerAPI) {
	instance = markers
}

// G is the shorthand for Gettext behaviour
func G(msgid string) string {
	if instance == nil {
		return msgid
	} else {
		return instance.G(msgid)
	}
}

// NG is the shorthand for NGettext behaviour
func NG(msgid string, msgidPlural string, n int) string {
	if instance == nil {
		if n == 1 {
			// Singular
			return msgid
		} else {
			// Plural
			return msgidPlural
		}
	} else {
		return instance.NG(msgid, msgidPlural, n)
	}
}
