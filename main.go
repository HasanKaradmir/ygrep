package main

import (
	"bufio"
	"fmt"
	"github.com/fatih/color"
	"log"
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

	searchKey := "" // searching value

	//Read YAML File
	if input == "" { // Value getting inline command
		if len(os.Args) > 1 {
			input = os.Args[1]
			if len(os.Args) > 2 {
				searchKey = os.Args[2]
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
			searchKey = os.Args[1]
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

	found := printKeyContent(content, searchKey, 0)
	if !found {
		fmt.Printf("Key '%s' not found\n", searchKey)
	}
}

func printKeyContent(node *yaml.Node, key string, depth int) bool {
	found := false
	lowercaseKey := strings.ToLower(key)
	if node.Kind == yaml.MappingNode {
		for i := 0; i < len(node.Content); i += 2 {
			keyNode := node.Content[i]
			valueNode := node.Content[i+1]
			isContains := strings.Contains(strings.ToLower(keyNode.Value), lowercaseKey)
			if isContains {
				key = keyNode.Value
			}

			if keyNode.Kind == yaml.ScalarNode && isContains {
				fmt.Printf(color.RedString("Line %v:\n", keyNode.Line))
				// Key Found, Print Content
				fmt.Printf(color.YellowString("%s%s:\n", strings.Repeat("  ", depth), key))
				marshalAndPrint(valueNode, depth+1)
				found = true // Key Found
			} else {
				// Continue searching the key
				if printKeyContent(valueNode, key, depth+1) {
					found = true
				}
			}
		}
	} else if node.Kind == yaml.SequenceNode {
		for _, n := range node.Content {
			if printKeyContent(n, key, depth) {
				found = true
			}
		}
	}
	return found
}

func marshalAndPrint(node *yaml.Node, depth int) {
	out, err := yaml.Marshal(node)
	if err != nil {
		log.Fatalf("error: %v", err)
	}
	lines := strings.Split(string(out), "\n")
	for _, line := range lines {
		if line != "" {
			fmt.Printf(color.CyanString("%s%s\n", strings.Repeat("  ", depth), line))
		}
	}
}
