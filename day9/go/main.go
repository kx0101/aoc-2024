package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, err := os.ReadFile("../input.txt")
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	lines := strings.TrimSpace(string(file))

	blocks := parseDisk(lines)

	moveBlocks(blocks)

	checksum := calculateChecksum(blocks)

	fmt.Printf("checksum: %d\n", checksum)
}

func parseDisk(lines string) []rune {
	var blocks []rune
	fileID := 0

	for i := 0; i < len(lines); i += 2 {
		fileLength, _ := strconv.Atoi(string(lines[i]))
		freeSpaceLength := 0

		if i+1 < len(lines) {
			freeSpaceLength, _ = strconv.Atoi(string(lines[i+1]))
		}

		for j := 0; j < fileLength; j++ {
			blocks = append(blocks, rune('0'+fileID))
		}

		for j := 0; j < freeSpaceLength; j++ {
			blocks = append(blocks, '.')
		}

		fileID++
	}

	return blocks
}

func moveBlocks(blocks []rune) {
	length := len(blocks)
	maxFileID := -1

	for _, block := range blocks {
		if block != '.' && int(block-'0') > maxFileID {
			maxFileID = int(block - '0')
		}
	}

	for fileID := maxFileID; fileID >= 0; fileID-- {
		fileStartIndex := -1
		fileLength := 0

		for i, block := range blocks {
			if block == rune('0'+fileID) {
				if fileStartIndex == -1 {
					fileStartIndex = i
				}

				fileLength++
			} else if fileStartIndex != -1 {
				break
			}
		}

		if fileStartIndex == -1 {
			continue
		}

		freeSpaceStartIndex := -1
		for i := 0; i <= length-fileLength; i++ {
			canFit := true

			for j := 0; j < fileLength; j++ {
				if blocks[i+j] != '.' {
					canFit = false
					break
				}
			}

			if canFit && i < fileStartIndex {
				freeSpaceStartIndex = i
				break
			}
		}

		if freeSpaceStartIndex != -1 {
			for i := 0; i < fileLength; i++ {
				blocks[freeSpaceStartIndex+i] = rune('0' + fileID)
				blocks[fileStartIndex+i] = '.'
			}
		}
	}
}

func calculateChecksum(blocks []rune) int64 {
	var checksum int64

	for i, block := range blocks {
		if block != '.' {
			fileID := int(block - '0')
			checksum += int64(i * fileID)
		}
	}

	return checksum
}
