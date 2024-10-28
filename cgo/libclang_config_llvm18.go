//go:build !byollvm && llvm18

package cgo

/*
#cgo linux        CFLAGS:  -I/usr/lib64/llvm18/include
#cgo darwin,amd64 CFLAGS:  -I/usr/local/opt/llvm@18/include
#cgo darwin,arm64 CFLAGS:  -I/opt/homebrew/opt/llvm@18/include
#cgo freebsd      CFLAGS:  -I/usr/local/llvm18/include
#cgo linux        LDFLAGS: -L/usr/lib64/llvm18/lib -lclang
#cgo darwin,amd64 LDFLAGS: -L/usr/local/opt/llvm@18/lib -lclang
#cgo darwin,arm64 LDFLAGS: -L/opt/homebrew/opt/llvm@18/lib -lclang
#cgo freebsd      LDFLAGS: -L/usr/local/llvm18/lib -lclang
*/
import "C"
