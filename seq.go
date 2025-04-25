package cclib

import (
	"fmt"
	"io/fs"
	"regexp"
	"strconv"
	"strings"
)

// SeqInfo holds information about an image sequence.
type SeqInfo struct {
	BaseName   string
	Delimiter  string
	Padding    int
	StartFrame int
	EndFrame   int
	Extension  string
}

// findImageSequences scans the specified directory for image sequences.
func FindImageSequences(dir fs.FS) ([]SeqInfo, error) {
	files, err := fs.ReadDir(dir, ".")
	if err != nil {
		return nil, err
	}

	seqMap := make(map[string]SeqInfo)

	for _, file := range files {
		if file.IsDir() {
			continue
		}
		filename := file.Name()
		re := regexp.MustCompile(`^(.*?)([_.-])?(\d+)(\.[^.]+)$`)
		matches := re.FindStringSubmatch(filename)

		if matches == nil {
			continue
		}

		baseName := matches[1]
		delimiter := matches[2]
		numberStr := matches[3]
		extension := matches[4]

		padding := len(numberStr)
		frameNumber, err := strconv.Atoi(numberStr)
		if err != nil {
			continue
		}

		key := fmt.Sprintf("%s%s%s", baseName, delimiter, extension)
		info, exists := seqMap[key]

		if !exists {
			info = SeqInfo{
				BaseName:   baseName,
				Delimiter:  delimiter,
				Padding:    padding,
				StartFrame: frameNumber,
				EndFrame:   frameNumber,
				Extension:  strings.TrimPrefix(extension, "."),
			}
		} else {
			if frameNumber < info.StartFrame {
				info.StartFrame = frameNumber
			}
			if frameNumber > info.EndFrame {
				info.EndFrame = frameNumber
			}
		}

		info.Padding = max(padding, info.Padding)
		seqMap[key] = info
	}

	var sequences []SeqInfo
	for _, seqInfo := range seqMap {
		sequences = append(sequences, seqInfo)
	}

	return sequences, nil
}
