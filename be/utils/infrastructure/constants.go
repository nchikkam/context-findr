package utils

// upload rules/constraints
const Store = "./assets/txt/"

// 2MB limit for Upload
const FileUploadSizeLimit = 2 * 1024 * 1024

// allowed types
var SupportedMIMEs = map[string]struct{}{
	"text/plain; charset=utf-8": {},
	"application/octet-stream":  {},
}

// server configuration related
const Port = ":8080"
const ForwardedByClientIP = true

// white list proxies
var Proxies = []string{"127.0.0.1"}
