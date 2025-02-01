# Go PHP Parser - LOG6302

---
This project is a php parser written in Go. It uses tree-sitter to parse the php code and generate an AST. The AST is then used to generate an AST JSON file that contains the structure of the code.
Multiple analysis can be done on the AST JSON file to extract information about the code.

## Installation
### Docker (Recommended)
1. Clone the repository
2. Run `docker-compose up` in the root directory of the project
3. The parser is then available as a CLI tool. Run `go-php-parser --help` to see the available commands

### Manual
1. Clone the repository
2. Install the project dependencies by running `go mod download`
3. Ensure you have a gcc compiler installed on your machine
4. Run `go build -o go-php-parser main.go` with CGO_ENABLED=1

## Usage
The parser can be used as a CLI tool. The following commands are available:
- `parse` : Parse a php file and generate an AST JSON file. Command: `go-php-parser parse <path-to-php-file>`. Consult `go-php-parser parse --help` for more information.
- `operations` : Perform operations on the AST JSON file. Command: `go-php-parser operations <path-to-ast-json-file>`. Consult `go-php-parser operations --help` for more information.


## Examples
### Parse

```bash
# parse a single php file
go-php-parser parse ./examples/test.php
```
```bash
# parse a directory of php files recursively
go-php-parser parse --directory --recursive ./data
```
```bash
# Specify the output file/directory
go-php-parser parse --output ./output/file.ast.json ./examples/test.php
go-php-parser parse --output ./output/directory --directory --recursive ./data
```

### Operations
#### count-kind
```bash
# Count the number of nodes of a specific kind in the AST JSON file/directory
go-php-parser operations ./output/file.ast.json count-kind "kind"
go-php-parser operations --recursive ./output/directory count-kind "kind"
```
You can also specify multiple kinds to count with the count-kinds operation :
```bash
go-php-parser operations ./output/file.ast.json count-kinds "kind1" "kind2" "kind3"
go-php-parser operations --recursive ./output/directory count-kinds "kind1" "kind2" "kind3"
```

#### find-kind-tree
Kind trees are a way to represent the structure of the AST. The find-kind-tree operation allows you to find all the trees in the AST JSON file that match a specific kind tree.
Default extension of a kind tree file is `.kt.json`.
The structure of a kind tree is a JSON structure with the following format:
```json
{
    "kind": "<kind>",
    "attributes": {
      "<attribute>": "<value>",
      ...
    },
    "children": [
      {
        "kind": "<kind>",
        "attributes": {
          "<attribute>": "<value>",
          ...
        },
        "children": [
          ...
        ]
      },
      ...
    ]
  }
```
You can also provide a map of kind trees to find in the AST JSON file with the `find-kind-trees` operation. The operation will return all the trees that match any of the provided kind trees. The map has the following format:
```json
{
    "<name>": {
        "kind": "<kind>",
        "attributes": {
            "<attribute>": "<value>",
            ...
        },
        "children": [
            ...
        ]
    },
    ...
}
```

Example Kind trees are available in the `examples` directory.

To use the find-kind-tree operation, you can use the following commands:
```bash
# Find all the trees in the AST JSON file that match the provided kind tree
go-php-parser operations ./output/file.ast.json find-kind-tree "<kind-tree.kt.json>"
# Find all the trees in the AST JSON file that match any of the provided kind trees
go-php-parser operations ./output/file.ast.json find-kind-trees "<kind-trees.kt.json>"
```

## Contributors
- [Valentin Lemaire](https://github.com/28Pollux28)
- [Mattéo Ricard](https://github.com/RicardMatteo)

## License
This project is licensed under the CC BY-NC-SA 4.0 License - see the [LICENSE](LICENSE) file for details.
 