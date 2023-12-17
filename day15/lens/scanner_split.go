package lens

import "bufio"

func ScanStep(data []byte, atEOF bool) (advance int, token []byte, err error) {
	token = nil
	err = nil

	for i := 0; i < len(data); i++ {
		if data[i] == ',' {
			advance = i + 1
			break
		}
	}

	if advance > 0 {
		token = data[:advance-1]
	}

	if atEOF && advance == 0 {
		advance = len(data)
		token = data
		err = bufio.ErrFinalToken
	}

	return
}
