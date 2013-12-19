// +build OMIT

package main

import ()

// STARTREADFILE OMIT
func readFile(filename string) (string, error) {
	buf := bytes.NewBuffer(nil)

	f, err := os.Open(filename)
	if nil != err {
		return ``, err
	}
	defer f.Close() // HL

	_, err = io.Copy(buf, f) // HL
	if nil != err {
		return ``, err
	}

	return string(buf.Bytes())
}

// STOPREADFILE OMIT
