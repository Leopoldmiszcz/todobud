package main

import (
	"flag"

	jsonHand "github.com/Leopoldmiszcz/todobud/pkg/jsonHandler"
)

// import filecheck "github.com/Leopoldmiszcz/todobud/internal/fileCheck"

// type ServerConfig struct {
// 	port    int
// 	workers int
// 	env     string
// }

type Task struct {
	task string
	done bool
}

func addTask(task Task, input string) {
	task.task = input
	task.done = false

	jsonHand.FileCheck()
	jsonHand.WriteJSONFile(task.task)

}

func main() {

	var task Task
	var input string
	flag.StringVar(&input, "add", "none", "Adds task")

	flag.Parse()

	addTask(task, input)
	// filecheck.FileCheck()
	// var age int

	// var serverConfig ServerConfig

	// flag.IntVar(&serverConfig.port, "port", 8000, "Server port")
	// flag.IntVar(&serverConfig.workers, "workers", 4, "Number of worker processes")
	// flag.StringVar(&serverConfig.env, "env", "dev", "Environment")

	// // var age = flag.Int("age", 20, "An age int")
	// flag.Parse()

	// fmt.Println(serverConfig)
}
