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
	app.Version = "1.0.0"

	app.Flags = []cli.Flag{
		&cli.BoolFlag{
			Name:  "view,v",
			Usage: "show arrays",
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
			Usage: "start goroutines count",
			Value: 0,
		},
		&cli.IntFlag{
			Name:  "maxGo,mgo",
			Usage: "end goroutines count",
			Value: 0,
		},
		&cli.IntFlag{
			Name:  "dataCount,c",
			Usage: "count of testing data",
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
					c.Int("maxGo"), c.Int("dataCount"), c.Int("elementRange"), c.Bool("view"))
				return nil
			},
		},
		{
			Name:    "input",
			Aliases: []string{"i"},
			Action: func(c *cli.Context) error {
				array, err := enterArray()
				if err != nil {
					return err
				}
				inputArray(array, c.Int("go"), c.Int("mgo"))
				return nil
			},
		},
		{
			Name:    "perm",
			Aliases: []string{"p"},
			Action: func(c *cli.Context) error {
				permArraySort(c.Int("size"), c.Int("maxSize"), c.Int("go"),
					c.Int("maxGo"), c.Int("dataCount"), c.Bool("view"))
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
func randArraySort(size, maxSize, goroutineStart, goroutineEnd, dataCount, elementRange int, show bool) {
	sorting := *tester.NewTester(size, maxSize, goroutineStart, goroutineEnd, dataCount)
	sorting.GenerateRandomData(elementRange)
	sorting.TestLinear()
	sorting.DisplayStats()
	if show {
		sorting.Display()
	}
}

func permArraySort(size, maxSize, goroutineStart, goroutineEnd, dataCount int, show bool) {
	sorting := *tester.NewTester(size, maxSize, goroutineStart, goroutineEnd, dataCount)
	sorting.GeneratePermData()
	sorting.TestLinear()
	sorting.DisplayStats()
	if show {
		sorting.Display()
	}
}

func inputArray(array []int, startGoroutines, endGoroutines int) {
	size := len(array)
	sorting := *tester.NewTester(size, size, startGoroutines, endGoroutines, 1)
	sorting.SetData(array)
	sorting.TestLinear()
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
