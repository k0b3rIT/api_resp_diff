package main

import (
	"bufio"
	"bytes"
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/kylelemons/godebug/diff"
	"gopkg.in/yaml.v2"
)

var Reset = "\033[0m"
var Red = "\033[31m"
var Green = "\033[32m"
var Yellow = "\033[33m"
var Blue = "\033[34m"
var Purple = "\033[35m"
var Cyan = "\033[36m"
var Gray = "\033[37m"
var White = "\033[97m"

type ApiDiffConfig struct {
	Hosts []string
	Tests []ApiTest
}

type ApiTest struct {
	Api    string
	Params []map[string]string
}

func readConfig(configFilePath string) ApiDiffConfig {
	configStr, err := ioutil.ReadFile(configFilePath)
	if err != nil {
		fmt.Print(err)
	}

	fmt.Printf("Reading config: %v\n", configFilePath)

	var apiDiffConfig ApiDiffConfig

	err = yaml.Unmarshal([]byte(configStr), &apiDiffConfig)

	if err != nil {
		panic(fmt.Sprint("Unable to parse config file, %v", err))
	}

	return apiDiffConfig
}

func parseArgs() map[string]string {
	configPtr := flag.String("config", "", "config file path")
	flag.Parse()
	configFilePath := *configPtr

	if configFilePath == "" {
		panic("you have to define the config file")
	}
	configs := make(map[string]string)
	configs["configFilePath"] = configFilePath
	return configs
}

func executeApiCompare(apiDiffConfig ApiDiffConfig) {
	for _, apiTest := range apiDiffConfig.Tests {
		for _, param := range apiTest.Params {
			endpoint := substituteParams(apiTest.Api, param, "{", "}")
			fmt.Printf("######## Execute api test: %s\n", endpoint)
			responses := make(map[string]string)
			for _, host := range apiDiffConfig.Hosts {
				resp, err := executeApiCall(host, endpoint)
				if err != nil {
					panic(err)
				}
				responses[host] = resp
			}
			compareResponses(responses)
		}
	}
}

func prettyString(str string) (string, error) {
	var prettyJSON bytes.Buffer
	if err := json.Indent(&prettyJSON, []byte(str), "", "    "); err != nil {
		return "", err
	}
	return prettyJSON.String(), nil
}

func compareResponses(responses map[string]string) {
	keys := make([]string, 0, len(responses))
	values := make([]string, 0, len(responses))

	for k, v := range responses {
		keys = append(keys, k)
		values = append(values, v)
	}

	a, _ := prettyString(values[0])
	b, _ := prettyString(values[1])

	diff := diff.Diff(b, a)

	fmt.Println(colorize(diff))
}

func colorize(diff string) string {
	var b bytes.Buffer

	scanner := bufio.NewScanner(strings.NewReader(diff))
	for scanner.Scan() {
		s := scanner.Text()
		if s[0:1] == "+" {
			s = inGreen(s)
		} else if s[0:1] == "-" {
			s = inRed(s)
		}
		b.WriteString(s)
		b.WriteString("\n")
	}
	return b.String()
}

func inRed(str string) string {
	return Red + str + Reset
}

func inGreen(str string) string {
	return Green + str + Reset
}

func executeApiCall(host string, endpoint string) (string, error) {
	requestUrl := fmt.Sprintf("%s%s", host, endpoint)
	fmt.Printf("%s", requestUrl)
	resp, err := http.Get(requestUrl)
	if err != nil {
		panic(fmt.Sprint("failed to send request to [%s]", requestUrl))
	}

	defer resp.Body.Close()

	fmt.Printf(" [HTTP-%v]\n", resp.StatusCode)

	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	bodyString := string(bodyBytes)

	if resp.StatusCode == http.StatusOK {
		return bodyString, nil
	}

	return "", fmt.Errorf("API error: %v", bodyString)
}

func substituteParams(api string, params map[string]string, paramPrefix string, paramPostfix string) string {
	afterReplace := api
	for k, v := range params {
		afterReplace = strings.Replace(afterReplace, paramPrefix+k+paramPostfix, v, -1)
	}
	return afterReplace
}

func main() {

	args := parseArgs()
	apiDiffConfig := readConfig(args["configFilePath"])
	executeApiCompare(apiDiffConfig)

}
