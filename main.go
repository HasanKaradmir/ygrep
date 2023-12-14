package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"gopkg.in/yaml.v3"
)

// MappingNode = Dictionary
// SequenceNode = List
// ScalarNode = Solid String

func main() {

	stdin := bufio.NewReader(os.Stdin)
	input, _ := stdin.ReadString('\n')
	input = strings.TrimSpace(input)

	keyToFind := "" // searching value

	//Read YAML File
	if input == "" { // Value getting inline command
		if len(os.Args) > 1 {
			input = os.Args[1]
			if len(os.Args) > 2 {
				keyToFind = os.Args[2]
			} else {
				fmt.Println("No search value inline.")
				os.Exit(1)
			}
		} else {
			fmt.Println("No input provided")
			os.Exit(1)
		}
	} else { // Value getting STDIN
		if len(os.Args) > 1 {
			keyToFind = os.Args[1]
		} else {
			fmt.Println("No search value.")
			os.Exit(1)
		}
	}
	yamlFile, err := os.ReadFile(input)
	if err != nil {
		panic(err)
	}

	// Read the key to find from the command line arguments

	// Decode YAML File
	var Node yaml.Node
	if err := yaml.Unmarshal(yamlFile, &Node); err != nil {
		panic(err)
	}

	if len(Node.Content) == 0 {
		panic("No YAML docs found")
	}

	content := Node.Content[0] // Get Node Content

	findKeys(content, keyToFind)

	//Encode YAML File
	enc := yaml.NewEncoder(os.Stdout) // Create new encoder to os.Stdout
	enc.SetIndent(2)
	if err := enc.Encode(content); err != nil {
		panic(err)
	}

}

func findKeys(node *yaml.Node, key string) {

	lowerCase := strings.ToLower(key)

	for i, child := range node.Content { // scanning all content
		if node.Kind == yaml.SequenceNode && child.Kind == yaml.ScalarNode { // if node type = list and child of node type = string pass
			continue
		}
		if i%2 == 0 && child.Value != "" { // Keys = Twin, Values = Odd if child index = 0 (so, child = Key) and child value NOT empty
			if strings.Contains(strings.ToLower(child.Value), lowerCase) {
				child.Value = "FOUNDED_" + child.Value
			}
		}
		findKeys(child, key)
	}

}
