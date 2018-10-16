/*
  Copyright (c) 2013 José Carlos Nieto, http://xiam.menteslibres.org/
  Copyright (c) 2018 David L. Parsley, parsley@linuxjedi.org

  Permission is hereby granted, free of charge, to any person obtaining
  a copy of this software and associated documentation files (the
  "Software"), to deal in the Software without restriction, including
  without limitation the rights to use, copy, modify, merge, publish,
  distribute, sublicense, and/or sell copies of the Software, and to
  permit persons to whom the Software is furnished to do so, subject to
  the following conditions:

  The above copyright notice and this permission notice shall be
  included in all copies or substantial portions of the Software.

  THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND,
  EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF
  MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND
  NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT HOLDERS BE
  LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER IN AN ACTION
  OF CONTRACT, TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN CONNECTION
  WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.
*/

package dig

import "fmt"

// InterfaceMap is a type for arbitrary data structures that can be used with
// dig. The top-level of the value is always a map[string]interface, but
// values of any type can be stored and retrieved.
type InterfaceMap map[string]interface{}

// New returns an empty InterfaceMap
func New() InterfaceMap {
	return map[string]interface{}{}
}

// Set sets an arbitrary value in a struct given a route follwed by the
// value, or returns an error.
func (im InterfaceMap) Set(params ...interface{}) error {

	l := len(params)

	if l < 2 {
		return fmt.Errorf("missing value")
	}

	route := params[0 : l-1]
	value := params[l-1]

	if err := Dig(&im, route...); err != nil {
		return err
	}
	return Set(&im, value, route...)
}

// Get returns an arbitrary interface value, or interface of type error.
func (im InterfaceMap) Get(route ...interface{}) interface{} {
	var i interface{}

	err := Get(&im, &i, route...)
	if err != nil {
		return err
	}
	return i
}