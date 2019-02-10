// Copyright 2019 Google Inc.
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

// options.go contains configurable options for Button.

import (
	"fmt"

	"github.com/mum4k/termdash/cell"
	"github.com/mum4k/termdash/keyboard"
)

// Option is used to provide options.
type Option interface {
	// set sets the provided option.
	set(*options)
}

// option implements Option.
type option func(*options)

// set implements Option.set.
func (o option) set(opts *options) {
	o(opts)
}

// options holds the provided options.
type options struct {
	fillColor   cell.Color
	textColor   cell.Color
	shadowColor cell.Color
	height      int
	width       int
	key         keyboard.Key
}

// validate validates the provided options.
func (o *options) validate() error {
	if min := 1; o.height < min {
		return fmt.Errorf("invalid height %d, must be %d <= height", o.height, min)
	}
	if min := 1; o.width < min {
		return fmt.Errorf("invalid width %d, must be %d <= width", o.width, min)
	}
	return nil
}

// newOptions returns options with the default values set.
func newOptions(textWidth int) *options {
	return &options{
		fillColor:   DefaultFillColor,
		textColor:   DefaultTextColor,
		shadowColor: DefaultShadowColor,
		height:      DefaultHeight,
		width:       textWidth + 2, // One empty cell on each side of the text.
	}
}

// DefaultFillColor is the default for the FillColor option.
const DefaultFillColor = cell.ColorCyan

// FillColor sets the fill color of the button.
// Defaults to DefaultFillColor.
func FillColor(c cell.Color) Option {
	return option(func(opts *options) {
		opts.fillColor = c
	})
}

// DefaultTextColor is the default for the TextColor option.
const DefaultTextColor = cell.ColorBlack

// TextColor sets the color of the text label in the button.
// Defaults to DefaultTextColor.
func TextColor(c cell.Color) Option {
	return option(func(opts *options) {
		opts.textColor = c
	})
}

// DefaultShadowColor is the default of the ShadowColor option.
const DefaultShadowColor = cell.Color(250)

// ShadowColor sets the color of the shadow under the button.
// Defaults to DefaultShadowColor.
func ShadowColor(c cell.Color) Option {
	return option(func(opts *options) {
		opts.shadowColor = c
	})
}

// DefaultHeight is the default for the Height option.
const DefaultHeight = 2

// Height sets the height of the button in cells.
// Must be a positive non-zero integer.
// Defaults to DefaultHeight.
func Height(cells int) Option {
	return option(func(opts *options) {
		opts.height = cells
	})
}

// Width sets the width of the button in cells.
// Must be a positive non-zero integer.
// Defaults to the auto-width based on the length of the text label.
func Width(cells int) Option {
	return option(func(opts *options) {
		opts.width = cells
	})
}

// DefaultKey is the default value for the Key option.
const DefaultKey = keyboard.KeyEnter

// Key configures the keyboard key that presses the button.
// Defaults to DefaultKey.
func Key(k keyboard.Key) Option {
	return option(func(opts *options) {
		opts.key = k
	})
}
