/*
Copyright 2014 Google Inc. All Rights Reserved.
Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at
    http://www.apache.org/licenses/LICENSE-2.0
Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package godata

import (
	"log"

	"github.com/google/btree"
)

// RowData maps column names to column values for a given row.
type RowData map[string]interface{}

// RowOf wraps the given arguments into a RowData. Arguments must be given as
// string keys followed by values. The function panics if the arguments are not
// consistent with this requirement.
func RowOf(args ...interface{}) RowData {
	if len(args)%2 == 1 {
		log.Fatalf("Cannot call RowOf(%v) for odd number of arguments", args)
	}
	data := make(map[string]interface{})
	for i := 0; i < len(args)-1; i += 2 {
		key, ok := args[i].(string)
		if !ok {
			log.Fatalf("RowOf(%v) needs string keys on even indices", args)
		}
		data[key] = args[i+1]
	}
	return data
}

// Row represents a single entry in a Frame.
type Row struct {
	index Index
	data  RowData
}

// Less returns true if the row is less than the given Row sharing the same
// index type, or if the row index is less than the given Index sharing the
// same index type.
func (r Row) Less(item btree.Item) bool {
	switch item := item.(type) {
	default:
		log.Fatal("btree.Item is not a Row or Index")
	case Row:
		return r.index.Less(item.index)
	case Index:
		return r.index.Less(item)
	}
	return false
}
