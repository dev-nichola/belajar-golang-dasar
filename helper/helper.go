package helper

var version string = "1.0.0"
var Application string = "10.0.0"

func sayGoodBye(name string) string {
	return "Good Bye " + name
}
func SayHello(name string) string {
	return "Hello " + name
}

func Contoh(name string) string {
	return sayGoodBye(name)
}
