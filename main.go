package main

import (
	"os"

	"gopkg.in/yaml.v3"
)

func main() {

	//Read YAML File
	yamlArg := os.Args[1]
	yamlFile, err := os.ReadFile(yamlArg)
	if err != nil {
		panic(err)
	}

	// Decode YAML File
	var Node yaml.Node
	if err := yaml.Unmarshal(yamlFile, &Node); err != nil {
		panic(err)
	}

	if len(Node.Content) == 0 {
		panic("No YAML docs found")
	}

	content := Node.Content[0]

	findKeys(content)

	//Encode YAML File
	enc := yaml.NewEncoder(os.Stdout)
	enc.SetIndent(2)
	if err := enc.Encode(content); err != nil {
		panic(err)
	}

}

func findKeys(node *yaml.Node) {
	for i, child := range node.Content {
		if node.Kind == yaml.SequenceNode && child.Kind == yaml.ScalarNode {
			continue
		}
		if i%2 == 0 && child.Value != "" {
			child.Value = "KEY_" + child.Value
		}
		findKeys(child)
	}
}
