// Copyright 2022 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package govulncheck provides an experimental govulncheck API.
package govulncheck

import "github.com/hyangah/legacyvuln/internal/govulncheck"

var (
	// Source reports vulnerabilities that affect the analyzed packages.
	Source = govulncheck.Source

	// DefaultCache constructs cache for a vulnerability database client.
	DefaultCache = govulncheck.DefaultCache
)

type (
	// Config is the configuration for Main.
	Config = govulncheck.Config

	// Result is the result of executing Source.
	Result = govulncheck.Result

	// Vuln represents a single OSV entry.
	Vuln = govulncheck.Vuln

	// Module represents a specific vulnerability relevant to a
	// single module or package.
	Module = govulncheck.Module

	// Package is a Go package with known vulnerable symbols.
	Package = govulncheck.Package

	// CallStacks contains a representative call stack for each
	// vulnerable symbol that is called.
	CallStack = govulncheck.CallStack

	// StackFrame represents a call stack entry.
	StackFrame = govulncheck.StackFrame
)
