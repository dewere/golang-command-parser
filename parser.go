//
//   Copyright 2019 Aristarh Deryapa.
//
//   Licensed under the Apache License, Version 2.0 (the "License");
//   you may not use this file except in compliance with the License.
//   You may obtain a copy of the License at
//
//       http://www.apache.org/licenses/LICENSE-2.0
//
//   Unless required by applicable law or agreed to in writing, software
//   distributed under the License is distributed on an "AS IS" BASIS,
//   WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//   See the License for the specific language governing permissions and
//   limitations under the License.
//

package parser

import "strings"

// Clearing arguments.
// • Triming leading and trailing spaces.
// • Then removes zero-length elements.
func filterArgs(args []string) []string {
	filtered := args[:0]

	for i := 0; i < len(args); i++ {
		el := strings.TrimSpace(args[i])

		if len(el) > 0 {
			filtered = append(filtered, el)
		}
	}

	return filtered
}

// String parser like in discord.py
// From `"Hello, Mike!" How are you?` will produce [ "Hello, Mike!", "How", "are", "you?" ]
// And from "1       2   3" will produce [ "1", "2", "3" ]
func parseArgs(args []string) []string {
	result := make([]string, len(args))
	pos := 0

	var temp string
	inString := false

	for i := 0; i < len(args); i++ {
		el := args[i]
		endPos := strings.Index(el, "\"")

		if endPos == len(el)-1 && endPos > -1 && inString == true {
			if len(el) > 1 {
				result[pos] = temp + el[:len(el)-1]
			} else {
				result[pos] = temp
			}

			inString = false
			pos++

			continue
		}

		if inString == true {
			temp = temp + el + " "
			continue
		}

		if strings.Index(el, "\"") == 0 {
			inString = true
			temp = el[1:] + " "
			continue
		}

		result[pos] = el
		pos++
	}

	if inString {
		result[pos] = temp[:len(temp)-2]
	}

	return filterArgs(result)
}

// Simple command parser.
// Will yield result like [command, arguments]
func ParseCommand(data string) (string, []string) {
	if strings.Index(data, " ") == -1 {
		return data, make([]string, 0)
	}

	args := strings.Split(data, " ")
	cmd := args[0]
	args = args[1:]

	if len(args) == 1 {
		return cmd, args
	}

	return cmd, parseArgs(args)
}
