// Copyright Amazon.com, Inc. or its affiliates. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License"). You may
// not use this file except in compliance with the License. A copy of the
// License is located at
//
//	http://aws.amazon.com/apache2.0/
//
// or in the "license" file accompanying this file. This file is distributed
// on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either
// express or implied. See the License for the specific language governing
// permissions and limitations under the License.

package firecracker

import "strings"

// kernelArgs serializes+deserializes kernel boot parameters.
// Kernel docs: https://www.kernel.org/doc/Documentation/admin-guide/kernel-parameters.txt
type kernelArgs struct {
	parsed   map[string]*string
	original string
}

// serialize the kernelArgs back to a string that can be provided
// to the kernel
func (kargs kernelArgs) String() string {
	return kargs.original
}

func (kargs *kernelArgs) GetValue(key string) *string {
	return kargs.parsed[key]
}

// deserialize the provided string to a kernelArgs map
func parseKernelArgs(rawString string) *kernelArgs {
	argMap := make(map[string]*string)
	for _, kv := range strings.Fields(rawString) {
		// only split into up to 2 fields (before and after the first "=")
		kvSplit := strings.SplitN(kv, "=", 2)

		key := kvSplit[0]

		var value *string
		if len(kvSplit) == 2 {
			value = &kvSplit[1]
		}

		argMap[key] = value
	}

	return &kernelArgs{
		argMap,
		rawString,
	}
}
