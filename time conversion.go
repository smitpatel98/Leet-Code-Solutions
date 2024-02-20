package main

import (
    "bufio"
    "fmt"
    "io"
    "os"
    "strings"
    "strconv" // Added strconv package
)

/*
 * Complete the 'timeConversion' function below.
 *
 * The function is expected to return a STRING.
 * The function accepts STRING s as parameter.
 */

func timeConversion(s string) string {
    newFormat := ""
    if s[len(s)-2] == 'P' {
        if s[:2] == "12" {
            newFormat = "12"
        } else {
            tempHour, err := strconv.Atoi(s[:2]) // Converted string to int using strconv.Atoi
            checkError(err)
            tempHour += 12
            newFormat = strconv.Itoa(tempHour) // Converted int to string using strconv.Itoa
        }
    } else {
        if s[:2] == "12" {
            newFormat = "00"
        } else {
            newFormat = s[:2]
        }
        return newFormat + s[2:len(s)-2]
    }
    return newFormat + s[2:len(s)-2]
}

func main() {
    reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)

    stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
    checkError(err)

    defer stdout.Close()

    writer := bufio.NewWriterSize(stdout, 16*1024*1024)

    s := readLine(reader)

    result := timeConversion(s)

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
