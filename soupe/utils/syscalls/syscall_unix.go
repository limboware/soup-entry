//go:build unix

package sysutils

import "syscall"

func SetNonblock(fd uintptr, nonb bool) error {
	return syscall.SetNonblock(int(fd), nonb)
}

func Read(fd uintptr, buff []byte) (int, error) {
	return syscall.Read(int(fd), buff)
}
