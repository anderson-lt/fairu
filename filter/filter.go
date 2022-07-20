// Copyright (C) 2022, Anderson Lizarazo Tellez

// Package filter contains filters for the filesystem.
//
// All exported and first-class functions defined here are filters.
//
// Multiple Arguments
//
// Almost all filters admit multiple arguments, in this case, it would be like
// writing the same filter successively one after another with each argument.
// For example:
//  FilterVariadic:
//      - Arg 1
//      - Arg 2
//      - Arg 3
// It is equivalent to:
//  FilterVariadic: Arg 1 // Or
//  FilterVariadic: Arg 2 // Or
//  FilterVariadic: Arg 3
// That is, it is only necessary that it matches an argument.
//
// The filters with which this rule applies is:
//  - Glob
//  - Identifier
//  - Name
//  - Type
//
// The rest of filters accept only an argument or have its own syntax.
package filter

// Filter defines the specification of a filter, where path is the path to an
// existing file, and args are the filter arguments. And passed is a boolean
// that indicates if the file passed the filter.
type Filter func(path string, args []string) (passed bool)
