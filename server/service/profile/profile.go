package profile

import (
	"io/ioutil"
	"runtime/debug"
	"runtime/pprof"
	"runtime/trace"
	"os"
)

func HeapDump() (err error) {
	file, err := ioutil.TempFile(".", "dump")
	if err != nil {
		return
	}
	defer file.Close()
	debug.WriteHeapDump(file.Fd())
	return
}
var cpufile,memfile, tracefile *os.File
//go tool pprof [cpu.prof|mem.prof]
func PProf() (err error) {
	cpufile, err = ioutil.TempFile(".", "cpu")
	if err != nil {
		return
	}
	memfile, err = ioutil.TempFile(".", "mem")
	if err != nil {
		return
	}
	defer memfile.Close()
	err = pprof.StartCPUProfile(cpufile)
	if err != nil {
		return
	}
	err = pprof.WriteHeapProfile(memfile)
	if err != nil {
		return
	}
	return
}

func StopPProf() (err error){
	defer cpufile.Close()
	defer pprof.StopCPUProfile()
	return
}

func Trace() (err error){
	tracefile, err= ioutil.TempFile(".", "trace")
	if err != nil {
		return
	}
	err = trace.Start(tracefile)
	if err != nil {
		return
	}
	return
}

func StopTrace() (err error){
	defer tracefile.Close()
	defer trace.Stop()
	return
}

func CpuFile() string{
	return cpufile.Name()
}

func MemFile() string{
	return memfile.Name()
}

func TraceFile() string{
	return tracefile.Name()
}
