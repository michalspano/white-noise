// White noise generator written in GO
// With pictorial representation
package main

// Used imports
import (
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"strings"
)

const (
	// Number of command-line arguments
	maxArgumentCount int = 4
	// Gray scale limit
	maxGrayScale int = 15
	// ANSI escape colors
	RED string = "\033[0;31m"
	RESET string = "\033[0m"
	PURPLE string = "\033[0;35m"
	GREEN string = "\033[0;32m"
	YELLOW string = "\033[0;33m"
	// HEADER - Default primary header
	HEADER string = "P2\n"
	// EXTENSION - default file extension
	EXTENSION string = "pgm"
)

// Global path to output file (mutable)
var defaultOutPath = "out/out.pgm"

// Main function
func main() {
	cliStatus("White Noise Generator Initialised!")
	args := getArgs()

	// Check if the number of arguments is valid
	// Size: <2;4>
	if len(args) > maxArgumentCount || len(args) < maxArgumentCount - 2 {
		warnUsage("usage")
    }
	// Detect size 4 -> `-d` flag, to change default output path
	if len(args) == maxArgumentCount {
		// Detect if any of the arguments is '-d'
		for i := range args {
			if args[i] == "-d" {
				newPath := args[i + 1]
				if checkFileExtension(newPath) {
					updatePath(args[i + 1])
				} else {
					warnUsage("extension")
				}
				break
			}
		}
	}
	// Detect 3 command-line arguments
	if len(args) == maxArgumentCount - 1 {
		// Detect `-h` flag
		if args[len(args) - 1] == "-h" {
			warnUsage("help")
		} else {
			warnUsage("flag")
		}
	}

	// Write corresponding headers and generate scene
	width, height := outputParameters(args)
	writeHeader(width, height)
	generateWhiteNoise(width, height)

	// End status
	cliStatus("Generating white noise...done!")
	os.Exit(0)
}

// Get command line arguments
func getArgs() []string {
    return os.Args[1:]
}

// Ensure correct usage
func warnUsage(key string) {
	var ERR string
	if key == "usage" {
		ERR = "Usage: ./wnoise <width> <height> | `-d` <path>"
    } else if key == "type" {
		ERR = "Type must be an integer."
	} else if key == "file" {
		ERR = "File error detected."
	} else if key == "size" {
		ERR = "Image size must be greater than 0."
	} else if key == "extension" {
		ERR = "File extension must be `.pgm`."
	} else if key == "help" {
		fmt.Printf("%sDefault usage: ./wnoise <width> <height>\n" +
			"Optional usage: ./wnoise <width> <height> | `-d` <output_path>\n" +
			"- to change default output directory `out/out.pgm`%s\n", YELLOW, RESET)
	} else if key == "flag" {
		ERR = "Unsupported flag"
	}
	fmt.Printf("%s%s%s\n", RED, ERR, RESET)
	os.Exit(1)
}

// Function to get output parameters with error handling -> (int, int)
func outputParameters(args []string) (int, int) {
    width, err := strconv.Atoi(args[0])
    if err != nil {
        warnUsage("type")
    }
    height, err := strconv.Atoi(args[1])
    if err != nil {
        warnUsage("type")
    }

	// Check if the image size is valid
	if width <= 0 || height <= 0 {
		warnUsage("size")
	}
    return width, height
}

/* Write headers
Source: https://en.wikipedia.org/wiki/Netpbm
*/
func writeHeader(w int, h int) {
	cliStatus("Writing headers...")

	// Split by '/'
	fullSplitPath := strings.Split(defaultOutPath, "/")

	// A subdirectory is detected
	if len(fullSplitPath) > 1 {
		dir := strings.Join(fullSplitPath[:len(fullSplitPath)-1], "/")
		_, pathErr := os.Stat(dir)

		// Make directory if it doesn't exist and handle errors
		if os.IsNotExist(pathErr) {
			// Handle errors
			dirErr := os.Mkdir(dir, 0755)
			if dirErr != nil {
				warnUsage("file")
			}
		}
	}

	// Open output file
	file, err := os.Create(defaultOutPath)
	if err != nil {
        warnUsage("file")
    }

	// Write first level header
	_, firstLevelHeader := file.WriteString(HEADER)
	if firstLevelHeader != nil {
        warnUsage("file")
    }

	// Write second level header
	_, secondLevelHeader := file.WriteString(strconv.Itoa(w) + " " + strconv.Itoa(h) + "\n")
	if secondLevelHeader != nil {
		warnUsage("file")
	}

	// Write third level header, i.e. color limit
	_, thirdLevelHeader := file.WriteString(strconv.Itoa(maxGrayScale) + "\n")
	if thirdLevelHeader != nil {
		warnUsage("file")
	}
	// Close file
	_ = file.Close()
}

// Function to generate white noise with limit of maxGrayScale
// With given width and height
func generateWhiteNoise(width int, height int) {
	// Open an existing file at path with permission to write (a.k.a. append)
	file, err := os.OpenFile(defaultOutPath, os.O_APPEND|os.O_WRONLY, 0600)
	if err != nil {
		warnUsage("file")
	}

	// White noise height x width
	for i := 0; i < height; i++ {
		for j := 0; j < width; j++ {
			randomGrayScale := rand.Intn(maxGrayScale)
			// Handle file writing error
			_, Err := file.WriteString(strconv.Itoa(randomGrayScale) + " ")
			if Err != nil {
				warnUsage("file")
			}
		}
		// Add new line each row
		_, _ = file.WriteString("\n")
	}
}

// Create a function to update defaultOutPath
func updatePath(newPath string) {
	cliStatus("Updating default output path")
	defaultOutPath = newPath
}

// Check correct file extension
func checkFileExtension(newPath string) bool {
    // Split the path
    fullSplitPath := strings.Split(newPath, "/")
    // Get the last element
    lastElement := fullSplitPath[len(fullSplitPath)-1]
    // Split the last element
    splitLastElement := strings.Split(lastElement, ".")
    // Get the last element of the file
    lastElement = splitLastElement[len(splitLastElement)-1]

    // Check if the last element is .pgm
    if lastElement == EXTENSION {
        return true
    }
    return false
}

// Custom CLI status formatted message
func cliStatus(msg string) {
	fmt.Printf("[  %sOK%s  ] %s%s%s\n", GREEN, RESET, PURPLE, msg, RESET)
}