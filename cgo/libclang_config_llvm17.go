//go:build !byollvm && llvm17

package cgo

/*
#cgo linux        CFLAGS:  -I/usr/lib64/llvm17/include
#cgo darwin,amd64 CFLAGS:  -I/usr/local/opt/llvm@17/include
#cgo darwin,arm64 CFLAGS:  -I/opt/homebrew/opt/llvm@17/include
#cgo freebsd      CFLAGS:  -I/usr/local/llvm17/include
#cgo linux        LDFLAGS: -L/usr/lib64/llvm17/lib -lclang
#cgo darwin,amd64 LDFLAGS: -L/usr/local/opt/llvm@17/lib -lclang
#cgo darwin,arm64 LDFLAGS: -L/opt/homebrew/opt/llvm@17/lib -lclang
#cgo freebsd      LDFLAGS: -L/usr/local/llvm17/lib -lclang
*/
import "C"
