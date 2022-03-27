package console

import (
	"QuickSort/internal/tester"
	"fmt"
	"github.com/urfave/cli/v2"
	"log"
	"os"
)

func Run() {
	app := initializeCLI()

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}

func initializeCLI() *cli.App {
	app := cli.NewApp()

	app.Name = "QuickSort"
	app.Usage = "To sort arrays"
	app.Version = "2.0.0"

	app.Flags = []cli.Flag{
		&cli.BoolFlag{
			Name:  "view,v",
			Usage: "show arrays inputs and results",
			Value: false,
		},
		&cli.IntFlag{
			Name:  "size,s",
			Usage: "set array size",
			Value: 10000,
		},
		&cli.IntFlag{
			Name:  "maxSize, ms",
			Usage: "set max array size",
			Value: 100000,
		},
		&cli.IntFlag{
			Name:  "elementRange,er",
			Usage: "set random element range",
			Value: 1000,
		},
		&cli.IntFlag{
			Name:  "go",
			Usage: "set test start goroutines count",
			Value: 1,
		},
		&cli.IntFlag{
			Name:  "maxGo,mgo",
			Usage: "set test end goroutines count",
			Value: 1,
		},
		&cli.IntFlag{
			Name:  "goStep,gos",
			Usage: "set test end goroutines count",
			Value: 1,
		},
		&cli.IntFlag{
			Name:  "dataCount,c",
			Usage: "set count of testing data to each size",
			Value: 1,
		},
	}

	app.Commands = []*cli.Command{
		{
			Name:    "rand",
			Aliases: []string{"r"},
			Usage:   "Use random arrays",
			Action: func(c *cli.Context) error {
				randArraySort(c.Int("size"), c.Int("maxSize"), c.Int("go"),
					c.Int("maxGo"), c.Int("dataCount"), c.Int("elementRange"), c.Int("goStep"),
					c.Bool("view"))
				return nil
			},
		},
		{
			Name:    "input",
			Aliases: []string{"i"},
			Usage:   "Input array manually",
			Action: func(c *cli.Context) error {
				array, err := enterArray()
				if err != nil {
					return err
				}
				inputArray(array, c.Int("go"), c.Int("mgo"), c.Int("goStep"))
				return nil
			},
		},
		{
			Name:    "perm",
			Aliases: []string{"p"},
			Usage:   "Use unique elements in arrays",
			Action: func(c *cli.Context) error {
				permArraySort(c.Int("size"), c.Int("maxSize"), c.Int("go"),
					c.Int("maxGo"), c.Int("dataCount"), c.Int("goStep"), c.Bool("view"))
				return nil
			},
		},
	}

	app.Action = func(c *cli.Context) error {
		fmt.Println("Test")
		return nil
	}

	return app
}
func randArraySort(size, maxSize, goroutineStart, goroutineEnd, dataCount, elementRange, goStep int, show bool) {
	sorting := *tester.NewTester(size, maxSize, goroutineStart, goroutineEnd, dataCount, goStep)
	sorting.GenerateRandomData(elementRange)
	sorting.TestSequential()
	sorting.TestConcurrent()
	sorting.TestParallel()
	sorting.DisplayStats()
	if show {
		sorting.Display()
	}
}

func permArraySort(size, maxSize, goroutineStart, goroutineEnd, dataCount, goStep int, show bool) {
	sorting := *tester.NewTester(size, maxSize, goroutineStart, goroutineEnd, dataCount, goStep)
	sorting.GeneratePermData()
	sorting.TestSequential()
	sorting.TestConcurrent()
	sorting.TestParallel()
	sorting.DisplayStats()
	if show {
		sorting.Display()
	}
}

func inputArray(array []int, startGoroutines, endGoroutines, goStep int) {
	size := len(array)
	sorting := *tester.NewTester(size, size, startGoroutines, endGoroutines, 1, goStep)
	sorting.SetData(array)
	sorting.TestSequential()
	sorting.TestConcurrent()
	sorting.TestParallel()
	sorting.Display()
	sorting.DisplayStats()
}

func enterArray() ([]int, error) {
	fmt.Printf("Enter size of your array: ")
	var size int
	_, err := fmt.Scanln(&size)
	if err != nil {
		return nil, err
	}

	var array = make([]int, size)
	for i := 0; i < size; i++ {
		fmt.Printf("Enter %dth element: ", i)
		_, err := fmt.Scan(&array[i])
		if err != nil {
			return nil, err
		}
	}
	fmt.Println()

	return array, nil
}
