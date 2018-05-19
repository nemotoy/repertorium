package service

import (
	"bufio"
	"fmt"
	"os"

	"go.uber.org/zap"
)

// Filter ...
func Filter(name, language, listupOutputPath, filterOutputPath string) error {
	logger, _ := zap.NewProduction()
	defer logger.Sync()

	// r := regexp.MustCompilePOSIX(name)

	fp, err := os.Open(listupOutputPath)
	if err != nil {
		logger.Error("@os.Open", zap.String("listupOutputPath", listupOutputPath), zap.String("error", err.Error()))
		return err
	}
	defer func() {
		if fp != nil {
			fp.Close()
		}
	}()
	scanner := bufio.NewScanner(fp)

	fmt.Println("######################################################################")
	fmt.Println("######################################################################")
	fmt.Println("######################################################################")
	for scanner.Scan() {
		fmt.Println(scanner.Text())

		// FIXME json parse
		// FIXME regexp find
		// FIXME write output file
	}
	if err := scanner.Err(); err != nil {
		logger.Error("@scanner.Scan", zap.String("error", err.Error()))
		return err
	}
	fmt.Println("######################################################################")
	fmt.Println("######################################################################")
	fmt.Println("######################################################################")
	return nil
}
