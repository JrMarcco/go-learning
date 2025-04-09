package cmd

import (
	"log"
	"strconv"
	"strings"
	"time"

	"github.com/JrMarcco/go-learning/word-cli/internal/timer"
	"github.com/spf13/cobra"
)

var timeCmd = &cobra.Command{
	Use:   "time",
	Short: "deal time format",
	Long:  "deal time format",
	Run:   func(cmd *cobra.Command, args []string) {},
}

var nowCmd = &cobra.Command{
	Use:   "now",
	Short: "get current time",
	Long:  "get current time",
	Run: func(cmd *cobra.Command, args []string) {
		now := time.Now()
		log.Printf("Res: %s, %d\n", now.Format("2006-01-02 15:04:05"), now.Unix())
	},
}

var calcTime string
var duration string

var calcTimeCmd = &cobra.Command{
	Use:   "calc",
	Short: "get calc time",
	Long:  "get calc time",
	Run: func(cmd *cobra.Command, args []string) {
		var current time.Time
		var layout = "2006-01-02 15:04:05"
		if calcTime == "" {
			current = timer.GetNow()
		} else {
			var err error
			space := strings.Count(calcTime, " ")
			if space == 0 {
				layout = "2006-01-02"
			}
			if space == 1 {
				layout = "2006-01-02 15:04:05"
			}
			current, err = time.Parse(layout, calcTime)
			if err != nil {
				t, _ := strconv.Atoi(calcTime)
				current = time.Unix(int64(t), 0)
			}
		}

		t, err := timer.GetCalcTime(current, duration)
		if err != nil {
			log.Fatalf("failt to get calc time: %v\n", err)
		}

		log.Printf("Res: %s, %d\n", t.Format(layout), t.Unix())
	},
}

func init() {
	timeCmd.AddCommand(nowCmd)
	timeCmd.AddCommand(calcTimeCmd)

	calcTimeCmd.Flags().StringVarP(&calcTime, "calc", "c", "", "Please input formatted time")
	calcTimeCmd.Flags().StringVarP(&duration, "duration", "d", "", "Please input time duration")
}
