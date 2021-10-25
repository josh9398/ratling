package cmd

import (
	"fmt"
	"io"
	"log"
	"math/rand"
	"os"
	"path/filepath"
	"time"

	"github.com/spf13/cobra"
	"go.uber.org/zap"
)

var (
	// flags
	fileToBeChunked string
	minChunkBytes   int64
	maxChunkBytes   int64
)

func init() {
	rootCmd.AddCommand(sendCmd)
	sendCmd.Flags().StringVarP(&fileToBeChunked, "file", "f", "", "path of file to send")
	sendCmd.Flags().Int64VarP(&minChunkBytes, "min-chunk", "m", 1000000, "min chunk size in bytes")
	sendCmd.Flags().Int64VarP(&maxChunkBytes, "max-chunk", "M", 50000000, "max chunk size in bytes")
}

var sendCmd = &cobra.Command{
	Use:   "send",
	Short: `chunk & send file`,
	Long:  `chunk file into cache & send`,
	RunE: func(cmd *cobra.Command, args []string) error {
		if fileToBeChunked == "" {
			log.Fatal("no file supplied")
		}

		fi, err := os.Open(fileToBeChunked)
		if err != nil {
			log.Fatalf("cannot open file: %v", err)
		}
		defer fi.Close()
		logger.Debugf("opened file: %s", fileToBeChunked)

		fiInfo, err := fi.Stat()
		if err != nil {
			log.Fatalf("cannot stat file: %v", err)
		}

		var fiName = filepath.Base(fileToBeChunked)
		var fiSize int64 = fiInfo.Size()
		logger.Debugf("file size: %d", fiSize)

		// small files
		if fiSize <= minChunkBytes {
			logger.Debugf("file %s too small to chunk, under %d bytes", fiName, minChunkBytes)
			if err := send(fi, "TODO"); err != nil {
				logger.Fatalf("error sending file: %v", err)
			}
			return nil
		}

		var cache = fmt.Sprintf("%s/%s", cacheDir, fiName)
		if err = os.MkdirAll(cache, os.ModePerm); err != nil {
			logger.Fatalf("cannot create cache: %v", err)
		}

		var part int64 = 0
		for {
			rand.Seed(time.Now().UnixNano())
			randSize := rand.Int63n((maxChunkBytes - minChunkBytes) + minChunkBytes)

			// read a chunk
			buf := make([]byte, randSize)
			n, err := fi.Read(buf)
			if err != nil && err != io.EOF {
				logger.Fatalf("error reading file: %v", err)
			}
			if n == 0 {
				break
			}

			// create cache file
			foPath := fmt.Sprintf("%s/%s_%d", cache, fiName, part)
			fo, err := os.Create(foPath)
			if err != nil {
				log.Fatalf("cannot create %s in cache: %v", foPath, err)
			}
			defer fo.Close()

			// write a chunk
			if _, err := fo.Write(buf[:n]); err != nil {
				log.Fatalf("cannot write to %s: %v", foPath, err)
			}
			logger.Debugf("chunked file: %s", fiName,
				zap.Int64("part", part),
				zap.Int("size", int(randSize)),
				zap.String("path", foPath),
			)

			send(fo, "TODO")
			part++
		}

		logger.Debugf("finished chunking file %s", fiName,
			zap.Int64("parts", part),
		)	
		return nil
	},
}

func send(fo *os.File, dest string) error {
	logger.Debugf("sending file %s to %s", fo.Name(), dest)
	return nil
}
