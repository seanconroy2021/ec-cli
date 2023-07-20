// Copyright The Enterprise Contract Contributors
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
//
// SPDX-License-Identifier: Apache-2.0

package something

import (
	"errors"
	"math/rand"
)

func addNewErrors() error {
	switch rand.Int() {
	case 0:
		return errors.New("0") // want `don't return raw error, return pkg/error.Error`
	case 1:
		return errors.New("1") // want `don't return raw error, return pkg/error.Error`
	case 2:
		return errors.New("2") // want `don't return raw error, return pkg/error.Error`
	case 3:
		return errors.New("3") // want `don't return raw error, return pkg/error.Error`
	default:
		return errors.New("4") // want `don't return raw error, return pkg/error.Error`
	}
}
