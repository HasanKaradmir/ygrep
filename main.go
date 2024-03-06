package main

import (
	"flag"
	"fmt"
	"github.com/fatih/color"
	"io"
	"os"
	"strings"

	"gopkg.in/yaml.v3"
)

// MappingNode = Dictionary
// SequenceNode = List
// ScalarNode = Solid String

func main() {
	value := flag.Bool("v", false, "value-based search")
	help := flag.Bool("h", false, "show help message")

	flag.Usage = func() {
		printHelp()
	}

	flag.Parse()

	if *help {
		printHelp()
		return
	}
	nFlags := flag.NFlag()
	searchKey, yamlContent, err := processArguments(os.Args, nFlags)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error processing arguments: %v\n", err)
		printHelp()
		os.Exit(1)
	}

	node, err := unmarshalYAML(yamlContent)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error unmarshaling YAML: %v\n", err)
		printHelp()
		os.Exit(1)
	}

	if len(node.Content) == 0 {
		fmt.Fprintln(os.Stderr, "No YAML docs found")
		os.Exit(1)
	}

	content := node.Content[0] // Get Node Content
	found, err := printKeyContent(content, searchKey, 0, *value)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Occurred Error: %v\n", err)
		printHelp()
		os.Exit(1)
	}
	if !found {
		fmt.Printf("Key '%s' not found\n", searchKey)
	}
}

// printHelp shows help messages.
func printHelp() {
	help := `USAGE: ygep [-v] <file-path>

search the contents of <file-path>.

Options:
  ygrep       : default setted to key-based value
  ygrep -v    : value based search
  
  ygrep -h    : show this message

Examples:
  key-based YAML search
  $ ygrep key ./search_file.yaml
  
  value-based YAML search
  $ ygrep -v value ./search_file.yaml`
	fmt.Fprintf(os.Stderr, "%s\n", help)
}

// processArguments processes and validates command line arguments.
func processArguments(args []string, flagsCount int) (searchKey string, yamlContent []byte, err error) {
	switch argCount := len(args); {
	case argCount > flagsCount+2: // Get Values from file
		yamlContent, err = os.ReadFile(args[flagsCount+2])
		if err != nil {
			return "", nil, err
		}
		searchKey = args[flagsCount+1]

	case argCount == flagsCount+2: // Get Values from STDIN
		if isStdinEmpty() {
			fmt.Println("No data provided in standard input.")
			os.Exit(1)
		}
		yamlContent, err = io.ReadAll(os.Stdin)
		if err != nil {
			return "", nil, err
		}
		searchKey = args[flag.NFlag()+1]

	default:
		err = fmt.Errorf("Invalid number of arguments")
	}
	return
}

// isStdinEmpty checks that Stdin is full or empty.
func isStdinEmpty() bool {
	stat, _ := os.Stdin.Stat()
	return (stat.Mode() & os.ModeCharDevice) != 0
}

// unmarshalYAML unmarshals the YAML content into a YAML node.
func unmarshalYAML(yamlContent []byte) (*yaml.Node, error) {
	var node yaml.Node
	err := yaml.Unmarshal(yamlContent, &node)
	return &node, err
}

// printKeyContent searches for a given key in a YAML node and prints its content.
func printKeyContent(node *yaml.Node, key string, depth int, kind bool) (found bool, err error) {
	if node.Kind != yaml.MappingNode && node.Kind != yaml.SequenceNode {
		return false, nil
	}

	lowercaseKey := strings.ToLower(key)
	return searchNode(node, lowercaseKey, depth, kind)
}

// searchNode handles the iteration and searching logic within a node.
func searchNode(node *yaml.Node, searchKey string, depth int, kind bool) (found bool, err error) {
	if node.Kind == yaml.MappingNode {
		return searchMappingNode(node, searchKey, depth, kind)
	}
	return searchSequenceNode(node, searchKey, depth, kind)
}

// searchMappingNode handles searching within a mapping node.
func searchMappingNode(node *yaml.Node, searchKey string, depth int, kind bool) (bool, error) {
	foundAny := false
	for i := 0; i < len(node.Content); i += 2 {
		keyNode, valueNode := node.Content[i], node.Content[i+1]
		if !kind {
			if containsKey(keyNode, searchKey) {
				foundAny = true // Mark that a match is found
				if valueNode.Kind == yaml.ScalarNode {
					printKeyValue(keyNode, valueNode, depth)
				} else {
					printKey(keyNode, depth)
					if err := marshalAndPrint(valueNode, depth); err != nil {
						return foundAny, err
					}
				}
			}

			if found, err := printKeyContent(valueNode, searchKey, depth, kind); err != nil || found {
				foundAny = foundAny || found
			}
		} else {
			if containsKey(valueNode, searchKey) {
				foundAny = true // Mark that a match is found
				if valueNode.Kind == yaml.ScalarNode {
					printKeyValue(keyNode, valueNode, depth)
				} else {
					printKey(keyNode, depth)
					if err := marshalAndPrint(valueNode, depth); err != nil {
						return foundAny, err
					}
				}
			}

			if found, err := printKeyContent(valueNode, searchKey, depth, kind); err != nil || found {
				foundAny = foundAny || found
			}
		}

	}
	return foundAny, nil
}

// searchSequenceNode handles searching within a sequence node.
func searchSequenceNode(node *yaml.Node, searchKey string, depth int, kind bool) (bool, error) {
	foundAny := false
	for _, n := range node.Content {
		if found, err := printKeyContent(n, searchKey, depth, kind); err != nil || found {
			foundAny = foundAny || found
		}
	}
	return foundAny, nil
}

// containsKey checks if the keyNode contains the search key.
func containsKey(keyNode *yaml.Node, searchKey string) bool {
	return keyNode.Kind == yaml.ScalarNode && strings.Contains(strings.ToLower(keyNode.Value), searchKey)
}

// printKeyValue prints the key-value pair in the desired format.
func printKeyValue(keyNode, valueNode *yaml.Node, depth int) {

	coloredKey := color.YellowString("%s%s:", strings.Repeat("  ", depth), keyNode.Value)
	coloredValue := color.CyanString(" %s", valueNode.Value)

	fmt.Printf(color.RedString("Line %v:\n", keyNode.Line))
	fmt.Printf("%s%s\n", coloredKey, coloredValue)
}

// printKey prints the key in a formatted manner.
func printKey(keyNode *yaml.Node, depth int) {
	fmt.Printf(color.RedString("Line %v:\n", keyNode.Line))
	fmt.Printf(color.YellowString("%s%s:\n", strings.Repeat("  ", depth), keyNode.Value))
}

// marshalAndPrint marshals a YAML node and prints its content.
func marshalAndPrint(node *yaml.Node, depth int) error {
	out, err := yaml.Marshal(node)
	if err != nil {
		return fmt.Errorf("Error Marshaling YAML: %v", err)
	}
	//if node.Kind == yaml.MappingNode {
	//	printIndented(out, 1)
	//} else {
	//	printIndented(out, depth)
	//}
	printIndented(out, depth)
	return nil
}

// printIndented prints the string with indentation.
func printIndented(text []byte, depth int) {
	lines := strings.Split(string(text), "\n")
	for _, line := range lines {
		if line != "" {
			fmt.Printf(color.CyanString("%s%s\n", strings.Repeat("  ", depth), line))
		}
	}
}
