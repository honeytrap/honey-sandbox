package web

//go:generate go-bindata -modtime=0 -pkg web -o bindata_gen.go -ignore \.map\$ build/...

var Prefix = "build"
