// Copyright 2020 Google Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package button

// write_options.go contains options used when writing content to the widget.

import "github.com/mum4k/termdash/cell"

// WriteOption is used to provide options to NewFromChunks().
type WriteOption interface {
	// set sets the provided option.
	set(*writeOptions)
}

// writeOptions stores the provided options.
type writeOptions struct {
	cellOpts []cell.Option
}

// setDefaultFgColor configures a default color for text if one isn't specified
// in the write options.
func (wo *writeOptions) setDefaultFgColor(c cell.Color) {
	wo.cellOpts = append(
		[]cell.Option{cell.FgColor(c)},
		wo.cellOpts...,
	)
}

// newWriteOptions returns new writeOptions instance.
func newWriteOptions(wOpts ...WriteOption) *writeOptions {
	wo := &writeOptions{}
	for _, o := range wOpts {
		o.set(wo)
	}
	return wo
}

// writeOption implements WriteOption.
type writeOption func(*writeOptions)

// set implements WriteOption.set.
func (wo writeOption) set(wOpts *writeOptions) {
	wo(wOpts)
}

// WriteCellOpts sets options on the cells that contain the text.
func WriteCellOpts(opts ...cell.Option) WriteOption {
	return writeOption(func(wOpts *writeOptions) {
		wOpts.cellOpts = opts
	})
}
