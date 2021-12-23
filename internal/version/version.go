package version

//go:generate go run gen.go
func GetSDKVersion() string {
	return version
}
