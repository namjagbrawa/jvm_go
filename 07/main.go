package main

import (
	"flag"
	"fmt"
	"github.com/namjagbrawa/jvm_go/instructions"
	"github.com/namjagbrawa/jvm_go/instructions/base"
	"github.com/namjagbrawa/jvm_go/rtda"
	"os"
)
import "strings"
import "github.com/namjagbrawa/jvm_go/classpath"
import "github.com/namjagbrawa/jvm_go/rtda/heap"

func main() {
	cmd := parseCmd()

	if cmd.versionFlag {
		fmt.Println("version 0.0.1")
	} else if cmd.helpFlag || cmd.class == "" {
		printUsage()
	} else {
		startJVM(cmd)
	}
}

func startJVM(cmd *Cmd) {
	cp := classpath.Parse(cmd.XjreOption, cmd.cpOption)
	classLoader := heap.NewClassLoader(cp, cmd.verboseClassFlag)

	className := strings.Replace(cmd.class, ".", "/", -1)
	mainClass := classLoader.LoadClass(className)
	mainMethod := mainClass.GetMainMethod()
	if mainMethod != nil {
		interpret(mainMethod, cmd.verboseInstFlag)
	} else {
		fmt.Printf("Main method not found in class %s\n", cmd.class)
	}
}


// java [-options] class [args...]
type Cmd struct {
	helpFlag         bool
	versionFlag      bool
	verboseClassFlag bool
	verboseInstFlag  bool
	cpOption         string
	XjreOption       string
	class            string
	args             []string
}

func parseCmd() *Cmd {
	cmd := &Cmd{}

	flag.Usage = printUsage
	flag.BoolVar(&cmd.helpFlag, "help", false, "print help message")
	flag.BoolVar(&cmd.helpFlag, "?", false, "print help message")
	flag.BoolVar(&cmd.versionFlag, "version", false, "print version and exit")
	flag.BoolVar(&cmd.verboseClassFlag, "verbose", false, "enable verbose output")
	flag.BoolVar(&cmd.verboseClassFlag, "verbose:class", false, "enable verbose output")
	flag.BoolVar(&cmd.verboseInstFlag, "verbose:inst", false, "enable verbose output")
	flag.StringVar(&cmd.cpOption, "classpath", "", "classpath")
	flag.StringVar(&cmd.cpOption, "cp", "", "classpath")
	flag.StringVar(&cmd.XjreOption, "Xjre", "", "path to jre")
	flag.Parse()

	args := flag.Args()
	if len(args) > 0 {
		cmd.class = args[0]
		cmd.args = args[1:]
	}

	return cmd
}

func printUsage() {
	fmt.Printf("Usage: %s [-options] class [args...]\n", os.Args[0])
	//flag.PrintDefaults()
}


func interpret(method *heap.Method, logInst bool) {
	thread := rtda.NewThread()
	frame := thread.NewFrame(method)
	thread.PushFrame(frame)

	defer catchErr(thread)
	loop(thread, logInst)
}

func catchErr(thread *rtda.Thread) {
	if r := recover(); r != nil {
		logFrames(thread)
		panic(r)
	}
}

func loop(thread *rtda.Thread, logInst bool) {
	reader := &base.BytecodeReader{}
	for {
		// 新的方法执行：将帧压入栈，并通过此命令取出当期帧
		frame := thread.CurrentFrame()
		pc := frame.NextPC()
		thread.SetPC(pc)

		// decode
		reader.Reset(frame.Method().Code(), pc)
		opcode := reader.ReadUint8()
		inst := instructions.NewInstruction(opcode)
		inst.FetchOperands(reader)
		frame.SetNextPC(reader.PC())  // Reader的下一个指令的位置

		if logInst {
			logInstruction(frame, inst)
		}

		// execute
		inst.Execute(frame)
		if thread.IsStackEmpty() {
			break
		}
	}
}

func logInstruction(frame *rtda.Frame, inst base.Instruction) {
	method := frame.Method()
	className := method.Class().Name()
	methodName := method.Name()
	pc := frame.Thread().PC()
	fmt.Printf("%v.%v() #%2d %T %v\n", className, methodName, pc, inst, inst)
}

func logFrames(thread *rtda.Thread) {
	for !thread.IsStackEmpty() {
		frame := thread.PopFrame()
		method := frame.Method()
		className := method.Class().Name()
		fmt.Printf(">> pc:%4d %v.%v%v \n",
			frame.NextPC(), className, method.Name(), method.Descriptor())
	}
}
