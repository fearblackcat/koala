// +build koala_go

package internal

import (
	"runtime"
	"github.com/v2pro/koala/countlog"
	"syscall"
	"net"
)

func SetCurrentGoRoutineIsKoala() {
	countlog.Trace("event!internal.set_is_koala", "threadID", GetCurrentGoRoutineId())
	runtime.SetCurrentGoRoutineIsKoala()
}

func GetCurrentGoRoutineIsKoala() bool {
	return runtime.GetCurrentGoRoutineIsKoala()
}

func GetCurrentGoRoutineId() int64 {
	return runtime.GetCurrentGoRoutineId()
}

func RegisterOnConnect(callback func(fd int, sa syscall.Sockaddr)) {
	syscall.OnConnect = callback
}

func RegisterOnAccept(callback func(serverSocketFD int, clientSocketFD int, sa syscall.Sockaddr)) {
	syscall.OnAccept = callback
}

func RegisterOnBind(callback func(fd int, sa syscall.Sockaddr)) {
	syscall.OnBind = callback
}

func RegisterOnRecv(callback func(fd int, span []byte)) {
	net.OnRead = callback
}

func RegisterOnSend(callback func(fd int, span []byte)) {
	net.OnWrite = callback
}
