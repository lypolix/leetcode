package main


import (
    "bufio"
    "fmt"
    "io"
    "os"
    "strconv"
    "strings"
)



func biggerIsGreater(w string) string {
	bs := []byte(w)
	
	n := len(bs)

	i := n - 2
	for i >= 0 && bs[i] >= bs[i + 1] {
		i --
	}
	if i < 0 {
		return "no answer"
	}

	j := n - 1

	for j > i && bs[j] <= bs[i]{
		j --
	}

	bs[i], bs[j] = bs[j], bs[i]

	for l, r := i+1, n-1; l < r; l, r = l+1, r-1 {
        bs[l], bs[r] = bs[r], bs[l]
    }

	return string(bs)
}

func main() {
    reader := bufio.NewReaderSize(os.Stdin, 16 * 1024 * 1024)

    stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
    checkError(err)

    defer stdout.Close()

    writer := bufio.NewWriterSize(stdout, 16 * 1024 * 1024)

    TTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
    checkError(err)
    T := int32(TTemp)

    for TItr := 0; TItr < int(T); TItr++ {
        w := readLine(reader)

        result := biggerIsGreater(w)

        fmt.Fprintf(writer, "%s\n", result)
    }

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
