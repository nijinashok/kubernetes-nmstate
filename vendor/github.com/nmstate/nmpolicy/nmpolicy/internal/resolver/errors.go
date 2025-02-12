/*
 * Copyright 2021 NMPolicy Authors.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at:
 *
 *	  http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package resolver

import "fmt"

func pathError(format string, a ...interface{}) error {
	return fmt.Errorf("invalid path: %v", fmt.Errorf(format, a...))
}

func wrapWithResolveError(err error) error {
	return fmt.Errorf("resolve error: %v", err)
}

func wrapWithEqFilterError(err error) error {
	return fmt.Errorf("eqfilter error: %v", err)
}

func replaceError(format string, a ...interface{}) error {
	return wrapWithReplaceError(fmt.Errorf(format, a...))
}

func wrapWithReplaceError(err error) error {
	return fmt.Errorf("replace error: %v", err)
}
