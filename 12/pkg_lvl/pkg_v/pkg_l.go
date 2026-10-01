package pkg_l

var (
	AppName string
)

func init() {
	if AppName == "" {
		AppName = "MyApp v1.0"
	}
}
