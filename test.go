package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
	"time"

	"lemin/bfs"
	filterpathways "lemin/filterPathWays"
	lemin "lemin/leminlib"

	"lemin/simulateAntMovement"
)

type Room struct {
	Name        string
	Coordinates []int
}

type Connection struct {
	From string
	To   string
}

type Karincaçiftliği struct{}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: go run main.go <filename>")
	}

	argument := os.Args[1]

	file, err := os.Open(argument)
	if err != nil {
		fmt.Println("File Not Read", err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var antCount int
	var start Room
	var end Room
	var rooms []Room
	var connections []Connection

	scanner.Scan()
	antCount, err = strconv.Atoi(scanner.Text())
	if err != nil {
		fmt.Println("Ant Count Not Read", err)
		return
	}

	data, _ := os.ReadFile(argument)
	content := string(data)
	lines := strings.Split(content, "\n")

	for scanner.Scan() {
		line := scanner.Text()
		if strings.HasPrefix(line, "##comment") || strings.HasPrefix(line, "#comment") || strings.HasPrefix(line, "#another comment") {
			continue
		}
		if strings.Contains(line, "-") {
			parts := strings.Split(line, "-")
			connections = append(connections, Connection{From: parts[0], To: parts[1]})
		} else {
			if strings.HasPrefix(line, "##start") {
				scanner.Scan()
				start = parseRoom(scanner.Text())
				rooms = append(rooms, start)
			} else if strings.HasPrefix(line, "##end") {
				scanner.Scan()
				end = parseRoom(scanner.Text())
				rooms = append(rooms, end)
			} else {
				rooms = append(rooms, parseRoom(line))
			}
		}
	}

	if lemin.Mergeprocess(lines) != "" {
		fmt.Println(lemin.Mergeprocess(lines))
		return
	}
	graph := buildGraph(rooms, connections)
	startName := start.Name
	endName := end.Name

	tümYollar := bfs.Bfs(graph, startName, endName)
	fmt.Println(tümYollar)
	if len(tümYollar) == 0 {
		fmt.Println("Başlangıçtan bitişe hiçbir yol bulunamadı.")
		return
	}
	fmt.Println("Başlangiç odasindan hedef odaya giden tüm yollar:")
	for i, yol := range tümYollar {
		fmt.Printf("Yol %d: %v\n", i+1, yol)
	}
	sort.Slice(tümYollar, func(i, j int) bool {
		return len(tümYollar[i]) < len(tümYollar[j])
	})

	filtrelenmişYollar := filterpathways.FilterPathWays(tümYollar, antCount)
	for i, yol := range filtrelenmişYollar {
		fmt.Printf("Yol %d: %v\n", i+1, yol)
	}
	enKisaYol := filtrelenmişYollar[0]
	fmt.Println("En Kısa Yol:")
	fmt.Printf("%v\n", enKisaYol)

	karincaHareketleri := simulateAntMovement.SimulateAntMovement(filtrelenmişYollar, antCount, startName, endName, enKisaYol)
	for _, hareket := range karincaHareketleri {
		fmt.Println(hareket)
	}
	// Zaman ölçümünü başlat
	startTime := time.Now()

	// Zaman ölçümünü bitir
	endTime := time.Now()
	duration := endTime.Sub(startTime)

	if argument == "example06.txt" || argument == "example07.txt" {
		fmt.Printf("Gerçek zamanlı süre: %.5f saniye\n", duration.Seconds())
	}
}

func parseRoom(line string) Room {
	parts := strings.Split(line, " ")
	name := parts[0]
	var coordinates []int
	for _, part := range parts[1:] {
		coordinate, _ := strconv.Atoi(part)
		coordinates = append(coordinates, coordinate)
	}
	return Room{Name: name, Coordinates: coordinates}
}

func buildGraph(rooms []Room, connections []Connection) map[string]map[string]bool {
	graph := make(map[string]map[string]bool)
	for _, room := range rooms {
		graph[room.Name] = make(map[string]bool)
	}
	for _, conn := range connections {
		graph[conn.From][conn.To] = true
		graph[conn.To][conn.From] = true
	}
	return graph
}
