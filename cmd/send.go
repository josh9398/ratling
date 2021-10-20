package cmd

import (
	"log"
	"math/rand"
	"os"
	"time"

	"github.com/spf13/cobra"
)

var (
	minChunkBytes int64 = 1000000

	// flags
	fileToBeChunked string
	maxChunkBytes   int64
)

func init() {
	rootCmd.AddCommand(sendCmd)
	sendCmd.Flags().StringVarP(&fileToBeChunked, "file", "f", "", "path of file to send")
	sendCmd.Flags().Int64VarP(&maxChunkBytes, "max-chunk", "m", 50000000, "max chunk size in bytes")
}

var sendCmd = &cobra.Command{
	Use:   "send",
	Short: `chunk & send file`,
	Long:  `chunk file into cache & send`,
	RunE: func(cmd *cobra.Command, args []string) error {
		if fileToBeChunked == "" {
			log.Fatal("no file supplied")
		}

		file, err := os.Open(fileToBeChunked)
		if err != nil {
			log.Fatalf("cannot open file: %v", err)
		}
		defer file.Close()

		fileInfo, err := file.Stat()
		if err != nil {
			log.Fatalf("cannot stat file: %v", err)
		}

		var fileSize int64 = fileInfo.Size()
		logger.Infof("file size: %d", fileSize)

		if fileSize < minChunkBytes {
			minChunkBytes = fileSize
		}
		if maxChunkBytes <= minChunkBytes {
			log.Fatalf("max-chunk too small, must be over: %d", minChunkBytes)
		}

		rand.Seed(time.Now().UnixNano())

		logger.Info(rand.Int63n(maxChunkBytes-minChunkBytes) + minChunkBytes)

		return nil
	},
}
