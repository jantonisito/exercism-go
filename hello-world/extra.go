//some file
// note that the package name is not the same as the directory name
// this is a common practice in Go
// the package name is the name that will be used in the import statement
// the directory name is the name that will be used in the import path
// for example, if you have a directory named "greeting" that contains a file named "hello.go"
// the package name in "hello.go" should be "greeting"
//Note tat we cannot have file with different package
//e.g. "hello.go" with package "greeting" and "extra.go" with package "greeting2"
package greeting

func HelloWorld2() string {
	return "Hello, World!"
}
