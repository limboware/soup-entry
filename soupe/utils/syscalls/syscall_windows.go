//go:build windows

package sysutils

import "syscall"

func SetNonblock(fd uintptr, nonb bool) error {
	return syscall.SetNonblock(syscall.Handle(fd), nonb)
}

func Read(fd uintptr, buff []byte) (int, error) {
	return syscall.Read(syscall.Handle(fd), buff)
}

func Write(fd uintptr, buff []byte) (int, error) {
	return syscall.Write(syscall.Handle(fd), buff)
}