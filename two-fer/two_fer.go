// Package twofer implement function makes phrase.
package twofer

import "fmt"

// ShareWith returns phrase for the name.
// If the given name is "Alice", the result should be "One for Alice, one for me.".
// If no name is given, the result should be "One for you, one for me.".
func ShareWith(name string) string {
    if name == "" {
        name = "you"
    }
	return fmt.Sprintf("One for %s, one for me.", name)
}
