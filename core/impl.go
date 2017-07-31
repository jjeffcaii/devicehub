package core

import (
	"encoding/csv"
	"bufio"
	"os"
	"io"
	"strings"
	"log"
)

type MapIndexer struct {
	cache map[string]string
}

func (impl *MapIndexer) Search(device *string) string {
	return impl.cache[strings.ToLower(*device)]
}

func parseAndroid(path string) (map[string]string, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	ret := make(map[string]string)
	reader := csv.NewReader(bufio.NewReader(f))
	var i int
	for i = 0; ; i++ {
		record, err := reader.Read()
		if err == io.EOF {
			break
		}
		if i > 0 {
			k, v := record[3], record[1]
			ret[k] = v
		}
	}
	log.Printf("load %d android devices.\n", i)
	return ret, nil
}

func parseIOS(path string) (map[string]string, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	ret := make(map[string]string)

	reader := bufio.NewReader(f)
	var i int
	for ; ; i++ {
		l, _, err := reader.ReadLine()
		if err == io.EOF {
			break
		}
		line := string(l)
		sp := strings.SplitN(line, "=", 2)
		k, v := sp[0], sp[1]
		ret[k] = v
	}
	log.Printf("load %d ios devices.\n", i)
	return ret, nil
}

func New(android, ios string) (Indexer, error) {
	m1, err1 := parseAndroid(android)
	if err1 != nil {
		return nil, err1
	}
	m2, err2 := parseIOS(ios)
	if err2 != nil {
		return nil, err2
	}
	m := make(map[string]string)
	for k, v := range m1 {
		m[strings.ToLower(strings.TrimSpace(k))] = strings.TrimSpace(v)
	}
	for k, v := range m2 {
		m[strings.ToLower(strings.TrimSpace(k))] = strings.TrimSpace(v)
	}
	return &MapIndexer{m}, nil
}
