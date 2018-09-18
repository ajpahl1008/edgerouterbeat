package beater

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/elastic/beats/libbeat/common"
	"log"
	"os"
	"os/exec"
	"strconv"
	"strings"
)

type InterfacesCondensed struct {
	Interfaces    []InterfaceCondensed `json:"interfaces"`
}

func (i *InterfacesCondensed) AddStats(iface InterfaceCondensed) {
	i.Interfaces = append(i.Interfaces, iface)
}

type InterfaceCondensed struct {
	InterfaceName string `json:"interfaceName"`
	ReceivePackets int64 `json:"receivePackets"`
	ReceiveBytes int64`json:"receiveBytes"`
	TransmitPackets int64 `json:"transmitPackets"`
	TransmitBytes int64 `json:"transmitBytes"`
}

func CollectEdgeStats() InterfacesCondensed {

	//input, err := exec.Command("/opt/vyatta/bin/vyatta-op-cmd-wrapper","show","interfaces","counters").CombinedOutput()
	input, err := exec.Command("cat","/Users/aj/Development/Go/src/examples/sample.header").CombinedOutput()
	if err != nil {
		log.Fatal(err)
	}

	linescanner := bufio.NewScanner(bytes.NewReader(input))

	onNewline := func(data []byte, atEOF bool) (advance int, token []byte, err error) {
		for i := 0; i < len(data); i++ {
			if data[i] == '\n' {
				return i + 1, data[:i], nil
			}
		}
		return 0, data, bufio.ErrFinalToken
	}

	edgeMetrics := InterfacesCondensed{}
	mapStringReturn := common.MapStr{}
	linescanner.Split(onNewline)
	for linescanner.Scan() {

		wordscanner := bufio.NewScanner(strings.NewReader(linescanner.Text()))
		wordscanner.Split(bufio.ScanWords)
		wordcount :=0

		edgeMetric := InterfaceCondensed{}
		for wordscanner.Scan() {
			if wordcount == 0 {edgeMetric.InterfaceName = wordscanner.Text()}
			if wordcount == 1 {edgeMetric.ReceivePackets, _ = strconv.ParseInt(wordscanner.Text(), 10, 64)}
			if wordcount == 2 {edgeMetric.ReceiveBytes, _ = strconv.ParseInt(wordscanner.Text(), 10, 64)}
			if wordcount == 3 {edgeMetric.TransmitPackets, _ = strconv.ParseInt(wordscanner.Text(), 10, 64)}
			if wordcount == 4 {edgeMetric.TransmitBytes, _ = strconv.ParseInt(wordscanner.Text(), 10, 64)}
			wordcount++
		}

		if edgeMetric.InterfaceName != "Interface" && edgeMetric.InterfaceName != ""{
			edgeMetric = InterfaceCondensed {
				InterfaceName: edgeMetric.InterfaceName,
				ReceivePackets: edgeMetric.ReceivePackets,
				ReceiveBytes: edgeMetric.ReceiveBytes,
				TransmitPackets: edgeMetric.TransmitPackets,
				TransmitBytes: edgeMetric.TransmitBytes,
			}
			edgeMetrics.AddStats(edgeMetric)
			tmpString,_ := json.Marshal(edgeMetric)
			mapStringReturn.CopyFieldsTo(mapStringReturn,string(tmpString))
		}
	}

	jsonReturn,errj := json.Marshal(edgeMetrics)
	if errj != nil {
		log.Fatal(errj)
	}
	fmt.Println(string(jsonReturn))

	if err := linescanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading input:", err)
	}

	//return string(jsonReturn)
	return edgeMetrics
}
