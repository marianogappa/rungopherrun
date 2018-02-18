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
	Speed1 int    `json:"speed1"`
	Speed2 int    `json:"speed2"`
	Result string `json:"result"`
}

type plug struct {
	name                        string
	calculateDistance           func(i, j int, graphEdges [][]t.Edge) int
	sortBasedOnPriceAndDistance func(ads []t.Ad, councilApproved func(t.Ad) bool)
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
	for i := range ads {
		ads[i].Distance = p.calculateDistance(0, ads[i].Location, graphEdges)
	}
	var elapsed1 = time.Now().Sub(start)
	p.sortBasedOnPriceAndDistance(ads, councilApproved)
	var res, elapsed2 = validate(ads), time.Now().Sub(start) - elapsed1
	return result{Name: p.name, Speed1: int(elapsed1.Nanoseconds() / 1e6), Speed2: int(elapsed2.Nanoseconds() / 1e6), Result: res}
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
			calculateDistance, ok := f2.(func(int, int, [][]t.Edge) int)
			if !ok {
				log.Println("Couldn't cast to string; Ignoring plugin")
			}
			sortBasedOnPriceAndDistance, ok := f3.(func([]t.Ad, func(t.Ad) bool))
			if !ok {
				log.Println("Couldn't cast to string; Ignoring plugin")
			}
			plugs = append(plugs, plug{*name, calculateDistance, sortBasedOnPriceAndDistance})
		}
	}

	return plugs
}

func validate(ls []t.Ad) string {
	var result = "OK!"
	if !reflect.DeepEqual(ls, sortedAds) {
		result = "The algorithm didn't produce the correct sorting order :("
		fmt.Println("It produced the following incorrect results")
		for _, ad := range ls {
			// if ad.Price != sortedAds[i].Price || ad.Distance != sortedAds[i].Distance {
			// 	fmt.Println(i)
			// }
			fmt.Println(ad.Price, ad.Distance)
		}
	}
	return result
}
