// Package greeting contains a function that prints out a customized greeting
package greeting

const testVersion = 3

// HelloWorld takes a name as a parameter and returns a string 'Hello [Name]' or 'Hello World' if no name is given
func HelloWorld(name string) string {
	// Write some code here to pass the test suite.
	if name == "" {
		name = "World"
	}
	return "Hello, " + name + "!"
}
