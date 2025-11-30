package main

import (
    "bufio"
    "fmt"
    "io"
    "os"
    "strings"
)



func superReducedString(s string) string {
    ns := []byte(s)
	stck := []byte{}
	for _, i := range ns {
		if len(stck) != 0 && stck[len(stck) - 1] == i {
			stck = stck[:len(stck) - 1]
		} else {
			stck = append(stck, i)
		}
	}

	if len(stck) == 0 {
		return "Empty String"
	} 
	return string(stck)
}

func main() {
    reader := bufio.NewReaderSize(os.Stdin, 16 * 1024 * 1024)

    stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
    checkError(err)

    defer stdout.Close()

    writer := bufio.NewWriterSize(stdout, 16 * 1024 * 1024)

    s := readLine(reader)

    result := superReducedString(s)

    fmt.Fprintf(writer, "%s\n", result)

    writer.Flush()
}

func readLine(reader *bufio.Reader) string {
    str, _, err := reader.ReadLine()
    if err == io.EOF {
        return ""
    }

    return strings.TrimRight(string(str), "\r\n")
}

func checkError(err error) {
    if err != nil {
        panic(err)
    }
}
