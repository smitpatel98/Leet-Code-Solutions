package main

import (
    "bufio"
    "fmt"
    "io"
    "os"
    "strconv"
    "strings"
)

func lonelyinteger(a []int32) int32 {
    number_of_occurance := make(map[int32]int)
    var single_occurance int32

    for _, i := range a {
        if _, ok := number_of_occurance[i]; !ok {
            number_of_occurance[i] = 1
            single_occurance = i
        } else {
            delete(number_of_occurance, i)
        }
    }
    return single_occurance
}

func main() {
    reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)

    stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
    checkError(err)

    defer stdout.Close()

    writer := bufio.NewWriterSize(stdout, 16*1024*1024)

    // Read the integer array size from input
    _, err = strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
    checkError(err)

    // Read the integer array elements from input
    aTemp := strings.Split(strings.TrimSpace(readLine(reader)), " ")
    var a []int32
    for i := 0; i < len(aTemp); i++ {
        aItemTemp, err := strconv.ParseInt(aTemp[i], 10, 64)
        checkError(err)
        aItem := int32(aItemTemp)
        a = append(a, aItem)
    }

    // Call the function and get the result
    result := lonelyinteger(a)

    // Write the result to output
    fmt.Fprintf(writer, "%d\n", result)

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
