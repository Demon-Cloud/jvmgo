/*
	命令行参数定义
*/
package main

import (
	"flag"
	"fmt"
	"os"
)

type Cmd struct {
	helpFlag 		bool
	versionFlag 	bool
	// classpath option
	cpOption 		string
	class			string
	args			[]string
}

// 定义解析命令行参数
func parseCmd() *Cmd {
	cmd := &Cmd{}

	// 如果调用 flag.Parse() 失败,则会调用该方法
	flag.Usage = printUsage
	flag.BoolVar(&cmd.helpFlag, "help", false, "print help message")
	flag.BoolVar(&cmd.helpFlag, "?", false, "print help message")
	flag.BoolVar(&cmd.versionFlag, "version", false, "print help message")
	flag.StringVar(&cmd.cpOption, "classpath", "", "classpath")
	flag.StringVar(&cmd.cpOption, "cp", "", "classpath")
	flag.Parse()

	args := flag.Args()
	if len(args) > 0 {
		cmd.class = args[0]
		cmd.args = args[1:]
	}

	return cmd
}

/**
 * 解析失败时调用
 */
func printUsage() {
	fmt.Printf("Usage: %s [-options] class [args...]\n", os.Args[0])
}