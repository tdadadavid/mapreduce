package cmd

import (
	"github.com/spf13/cobra"
	"github.com/tdadadavid/mapreduce/pkg/implementation"
	"github.com/tdadadavid/mapreduce/pkg/mapreduce"
	"os"
)

var mapCount int
var reduceCount int

var rootCmd = cobra.Command{
	Use:   "mr",
	Long:  "A map reduce implementation in Go",
	Short: "A map reduce implementation in Go",
	Run: func(cmd *cobra.Command, args []string) {

		// start the mapreduce
		mr := mr.New()

		// implementations
		// word-counter
		counter := implementation.NewWordCounter()
		adder := implementation.NewWordReducer()

		// register mappers and reducers
		mr.RegisterMapper(&counter)
		mr.RegisterReducer(&adder)

		// specify configuration for the mapreduce
		mr.SetMapCount(mapCount).
			SetReduceCount(reduceCount)

		// optionally, you add register combiners (mini-reduce) for better performance
		//TODO: add optional combiner support

		// run mapreduce framework
		mr.Run()
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.PersistentFlags().IntVarP(&mapCount, "map", "m", 10, "Number of mappers to run")
	rootCmd.PersistentFlags().IntVarP(&reduceCount, "reduce", "r", 10, "Number of reducers to run")
}
