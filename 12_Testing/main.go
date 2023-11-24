/*
WHY WRITE TESTS?
1- Initial Build: Ensure correctness of app
2- Production: Identify errors before users are impacted
3- Enhancements: Prevent new development from breaking older features

WHAT SHOULD YOU TEST?
-> correctness: the program does what it should
-> performance: the program does not consume excessive resources

Test Hierarchie: Unit -> Component -> Integration -> End to End

-> Unit: Prove that individual units of logic are correct 
-> Component: Confirm that different app concerns (package) perform correctly
-> Integration: Validate that entire program works as expexted
-> End to End: Demonstrate that entire system works correctly together

HOW DOES GO SUPPORT TESTING?
elements of correctness -> test - fuzz - example (http://pkg.go.dev/testing)
performance -> benchmark (http://go.dev/doc/diagnostics)
resource usage -> profiling - tracing (http://go.dev/doc/diagnostics)

Check: "Testing Go Application" Course!

↓↓↓↓
package main
func main() {
}
func Add(l, r int) int {
	return l + r
}
The first thing that we need to do is to give a location for the test to live. The way we do that in Go is to create a source file with --> main_test.go
For unit tests create using the name of the source file that contains the unit, so add function is in the source file main.go.
This _ (suffix) is critical, because this is what Go's test runner is going to use to discover the tests that we write. When we create production builds, the Go compiler will explicitly leaves these files out. These files are only included when we're executing tests, and they're not included in production builds to prevent any possibility of our tests introducing any security vulnerabilities.

First: The source file is a package declaration (package main). And then I need to create the test itself on the file main_test.go. Like this:

					package main
					func TestXxx(t *testing.T) {
					}

The rule with naming a test is that the first letter after the word test does have to be capitalized. If it is not, it is not a valid test for Go.
Normally, what we'll do is we'll just say what we're testing, especially with unit testing.

					package main
					func TestAdd(t *testing.T) {
					}

So we are testing the add function, we'll call this func TestAdd. And then notice that I do receive one parameter, and that is a pointer to a testing.T object.

We use that T object to communicate back to the test runner, what's going to inside of our test. That'S why have have a pointer, we're sharing memory back with the test runner. 

In the arrange section, we set up our test conditions. So what is the environment, what are the variables that we need before we execute the test that we're going to be confirming.

						package main

						import "testing"

						func TestAdd(t *testing.T) {
							// arrange

							// act

							// assert
						}

I declare her two variables, L and R variable. I set up also my expectation, so what do I expect to come out of this test.

						package main

						import "testing"

						func TestAdd(t *testing.T) {
							// arrange
							l, r := 1, 2
							expect := 3
							// act

							// assert
						}

In the action section this test is going to be to call the add function, passing in L and R. And I am receving that in a variable called got. And this is only my convention, you can name the result variables whatever you want. I've gotten into the habit of using expect and got for expect setting up my expectation and got capturing whatever I received, especially in a function call like this.

							package main

							import "testing"

							func TestAdd(t *testing.T) {
								// arrange
								l, r := 1, 2
								expect := 3
								// act
								got := Add(l, r)
								// assert
							}

In the last section assertion: We do not have an assertion API in Go. In some languages or some testing frameworks that are out there, you might see that there are certain assertions that are pre-provided for you. While Go does have third-party libraries that provide that, the standard library does not include any kind of assertions.The way we write assertions when we're just using the standard library is simply using control flow, i.e. usinf if statements. How do I want to confirm if this test worked properly? I want to check if expect is not equal to got, that means I've got a failure, because I would expect the expect variable to equal the got variable. If they are not, then I've got a failure. I report that failure is I'm going to call the Errorf method on the T object.
Remember, that T object allows me to communicate the status of my test back to the test runner. This is a formatted error statement. It means, it's going to take a pattern: ("Failed to add %v and %v. Got %v, expected %v\").. And then I ned provide those parameters-> l, r, got, expected.

							package main

							import "testing"

							func TestAdd(t *testing.T) {
								// arrange
								l, r := 1, 2
								expect := 3
								// act
								got := Add(l, r)
								// assert
								if expect != got {
									t.Errorf("Failed to add %v and %v. Got %v, expected %v\", l, r, got, expected)
								}
							}

VSCode has already offered to help me run these tests. Close the window main.go.

There is several ways that I can run this test.
The first is I can click up "run package tests". If you click that, VSCode is going to execute your test for you, and you can see that the test passes. We can see what a failure looks like by setting our expectation to an erroneous value (expect := 4). And if I run that again, then Go is going to run the package test for me. We actually see the command that Go ran in order to do that, up above at the top of terminal window. 

Running tool: /usr/local/go/bin/go test -timeout 30s -coverprofile=/tmp/vscode-go4YNOaQ/go-code-cover

So you can see that we have this go.exe, that's actually the Go command that we installed. Test is the subcommand, and then we have some additional information that's being provided. Now you can see that it did still tell me that it passed, and that's because the source file did not change, you can tell by this dot up here.

If you want to run from the terminal which is sth commonly done, fo over the terminal, and run the 'go test' command yourself. Often, the way that we do this is if your current working directory is where your test is located, you can simply provided a dot after the go test command, and that's going to invoke the test runner --> go test .
