package Logs

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/fatih/color"
)

/*PrintRouteTrace - function for output request */
func PrintRouteTrace(r *http.Request, flagError bool) {
	defer r.Body.Close()
	body, _ := ioutil.ReadAll(r.Body)

	fmt.Printf("%d-%02d-%02dT%02d:%02d:%02d-00:00 ",
		time.Now().Year(), time.Now().Month(), time.Now().Day(),
		time.Now().Hour(), time.Now().Minute(), time.Now().Second())

	printMethodType(r.Method, flagError)
	printAPIRoute(r.RequestURI, flagError)
	printRequestContentType(r.Header.Get("Content-Type"), flagError)
	printBodyRequest(string(body), flagError)

	fmt.Println("")
}

func printMethodType(method string, flagError bool) {
	if flagError {
		outRed := color.New(color.FgRed).Add(color.Underline)
		outRed.Print("[" + method + "]")
		fmt.Print("  ")
		return
	}
	outGreen := color.New(color.FgGreen)
	outGreen.Print("[" + method + "]")
	fmt.Print("  ")
}

func printAPIRoute(addr string, flagError bool) {
	if flagError {
		outRed := color.New(color.FgRed).Add(color.Underline)
		outRed.Print("api request: " + addr)
		fmt.Print("  ")
		return
	}
	outYellow := color.New(color.FgYellow).Add(color.Underline)
	outYellow.Print("api request: " + addr)
	fmt.Print("  ")
}

func printRequestContentType(contType string, flagError bool) {
	if flagError {
		outFit := color.New(color.FgRed)
		outFit.Print("Content-Type: " + string(contType))
		fmt.Print("  ")
		return
	}
	outFit := color.New(color.FgCyan)
	outFit.Print("Content-Type: " + string(contType))
	fmt.Print("  ")
}

func printBodyRequest(body string, flagError bool) {
	if flagError {
		outK := color.New(color.FgRed)
		outK.Print("body: " + string(body))
		return
	}
	outK := color.New(color.FgBlue)
	outK.Print("body: " + string(body))
}
