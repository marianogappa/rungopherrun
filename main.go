package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"plugin"
	"reflect"
	"strings"
	"time"

	t "github.com/marianogappa/rungopherrun/types"
)

func main() {
	newEndpoint(9000, runPlugins).start()
}

type result struct {
	Name   string `json:"name"`
	Speed  int    `json:"speed"`
	Result string `json:"result"`
}

type plug struct {
	name                        string
	filterMandatoryRequirements func(ads []t.Ad, requirements []string) []t.Ad
	calculateDistance           func(i, j int, graphEdges [][]t.Edge) int
	sortBasedOnPriceAndDistance func(ads []t.Ad)
}

func runPlugins() []result {
	var results = make([]result, 0)
	for _, p := range loadPlugins() {
		results = append(results, runPlugin(p))
	}
	return results
}

func runPlugin(p plug) result {
	fmt.Printf("Running %v\n", p.name)
	var start = time.Now()
	filteredList := p.filterMandatoryRequirements(ads, requirements)
	fmt.Printf("Finished filterMandatoryRequirements for %v\n", p.name)
	for i := range filteredList {
		filteredList[i].Distance = p.calculateDistance(0, filteredList[i].Location, graphEdges)
	}
	fmt.Printf("Finished calculateDistance for %v\n", p.name)
	p.sortBasedOnPriceAndDistance(filteredList)
	fmt.Printf("Finished sortBasedOnPriceAndDistance for %v\n", p.name)
	var res, elapsed = validate(filteredList, time.Now().Sub(start))
	return result{Name: p.name, Speed: elapsed, Result: res}
}

func loadPlugins() []plug {
	var (
		plugs      = make([]plug, 0)
		files, err = ioutil.ReadDir("/mnt/plugins/")
	)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Looking for plugins...")
	for _, file := range files {
		if file.IsDir() {
			log.Println("Found plugin", file.Name())
			//find .so file
			fs, err := ioutil.ReadDir("/mnt/plugins/" + file.Name())
			if err != nil {
				log.Printf("%v\nIgnoring plugin\n", err)
				continue
			}
			var so = ""
			for _, f := range fs {
				if strings.HasSuffix(f.Name(), ".so") {
					so = f.Name()
				}
			}
			if so == "" {
				log.Print("Haven't found a plugin here.\nIgnoring plugin\n", err)
				continue
			}
			p, err := plugin.Open("/mnt/plugins/" + file.Name() + "/" + so)
			if err != nil {
				log.Printf("%v\nIgnoring plugin\n", err)
			}
			n, err := p.Lookup("Name")
			if err != nil {
				log.Printf("%v\nIgnoring plugin\n", err)
			}
			f1, err := p.Lookup("FilterMandatoryRequirements")
			if err != nil {
				log.Printf("%v\nIgnoring plugin\n", err)
			}
			f2, err := p.Lookup("CalculateDistance")
			if err != nil {
				log.Printf("%v\nIgnoring plugin\n", err)
			}
			f3, err := p.Lookup("SortBasedOnPriceAndDistance")
			if err != nil {
				log.Printf("%v\nIgnoring plugin\n", err)
			}
			name, ok := n.(*string)
			if !ok {
				log.Println("Couldn't cast to string; Ignoring plugin")
			}
			filterMandatoryRequirements, ok := f1.(func([]t.Ad, []string) []t.Ad)
			if !ok {
				log.Println("Couldn't cast to string; Ignoring plugin")
			}
			calculateDistance, ok := f2.(func(int, int, [][]t.Edge) int)
			if !ok {
				log.Println("Couldn't cast to string; Ignoring plugin")
			}
			sortBasedOnPriceAndDistance, ok := f3.(func([]t.Ad))
			if !ok {
				log.Println("Couldn't cast to string; Ignoring plugin")
			}
			plugs = append(plugs, plug{*name, filterMandatoryRequirements, calculateDistance, sortBasedOnPriceAndDistance})
		}
	}

	return plugs
}

func validate(ls []t.Ad, speed time.Duration) (result string, elapsed int) {
	elapsed = int(speed.Nanoseconds() / 1e6)
	result = "OK!"
	if !reflect.DeepEqual(ls, sortedAds) {
		// elapsed = 1<<31 - 1
		result = "The algorithm didn't produce the correct sorting order :("
	}
	return
}
