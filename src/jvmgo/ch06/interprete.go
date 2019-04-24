// interprete.go
package main

import (
    "fmt"
    "jvmgo/ch06/classfile"
    "jvmgo/ch06/instructions"
    "jvmgo/ch06/instructions/base"
    "jvmgo/ch06/rtda"
)

func interpret(methodInfo *classfile.MemberInfo) {
    //  获取 MemberInfo 的 Code 属性
    codeAttr := methodInfo.CodeAttribute()

    maxLocals := codeAttr.MaxLocals()
    maxStack := codeAttr.MaxStack()
    bytecode := codeAttr.Code()

    // 创建一个线程,创建一个帧并把它推入Java虚拟机栈顶
    thread := rtda.NewThread()
    frame := thread.NewFrame(maxLocals, maxStack)
    thread.PushFrame(frame)

    defer catchErr(frame)
    loop(thread, bytecode)
}

func catchErr(frame *rtda.Frame) {
    if r := recover(); r != nil {
        fmt.Printf("LocalVars:%v\n", frame.LocalVars())
        fmt.Printf("OperandStack:%v\n", frame.OperandStack())
        panic(r)
    }
}

func loop(thread *rtda.Thread, bytecode []byte) {
    frame := thread.PopFrame()
    reader := &base.BytecodeReader{}

    for {
        pc := frame.NextPC()
        thread.SetPC(pc)

        // decode
        reader.Reset(bytecode, pc)
        opcode := reader.ReadUint8()
        inst := instructions.NewInstruction(opcode)
        inst.FetchOperands(reader)
        frame.SetNextPC(reader.PC())

        // execute
        fmt.Printf("pc:%2d inst:%T %v\n", pc, inst, inst)
        inst.Execute(frame)
    }
}