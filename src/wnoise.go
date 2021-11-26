// White noise generator written in GO
// With pictorial representation
package main

// Used imports
import (
	"fmt"
	"math/rand"
	"os"
	"os/exec"
	"strconv"
	"strings"
	"time"
)

const (
	// Number of command-line arguments
	maxArgumentCount int = 5
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
	// Default script path
	scriptPath string = "./png-parse/bin/png-parse.pyc"
)

// Global path to output file (mutable)
var defaultOutPath = "out/out.pgm"

// Main function
func main() {
	// Initial status message
	cliStatus("White Noise Generator Initialised!")

	// Receive passed command-line arguments
	args := func () []string {
		return os.Args[1:]
	}()

	// Default value to convert to PNG -> false
	var toConvertToPng = false

	// Store number of command-line arguments
	argc := len(args)

	// Default case
	if argc < 2 || argc > maxArgumentCount {
		warnUsage("usage")
    } else if argc == 3 {
		// Width, height, single flag
		flag := args[len(args) - 1]
		// Help flag
		if flag == "-h" {
			warnUsage("help")
			// Png converter flag
		} else if flag == "-png" {
			toConvertToPng = true
		} else {
			warnUsage("flag")
		}
	} else if argc == 4 {
		checkRelocateFlags(args, EXTENSION)
	// Max number of command line arguments
	} else if argc == maxArgumentCount {
        checkRelocateFlags(args, "png")
		for i := range args {
			if args[i] == "-png" {
				// Convert to png
				toConvertToPng = true
				break
			}
		}
    }
	// Write corresponding headers and generate scene
	width, height := outputParameters(args)
	writeHeader(width, height)
	generateWhiteNoise(width, height)

	// Convert to PNG (if required)
	if toConvertToPng {
		pngOutPath := func () string {
			return strings.Replace(defaultOutPath, EXTENSION, "png", 1)
		}()
        convertToPng(defaultOutPath, pngOutPath)
    }

	// End status
	cliStatus("Generating white noise...done!")
	os.Exit(0)
}

// Function to check command-line flags
func checkRelocateFlags(args []string, ext string) {
    // Check flags
    for i := range args {
		if args[i] == "-d" {
            newPath := args[i + 1]
            if checkFileExtension(newPath, ext) {
                updatePath(newPath)
				return
            } else {
				warnUsage("extension")
			}
		}
    }
	warnUsage("flag")
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
		ERR = "Invalid file extension."
	} else if key == "help" {
		fmt.Printf("%sDefault usage: ./wnoise <width> <height>\n" +
			"Optional usage: ./wnoise <width> <height> | `-d` <output_path> | -png\n" +
			"- to change default output directory `out/out.pgm`%s\n", YELLOW, RESET)
	} else if key == "flag" {
		ERR = "Unsupported flag"
	} else if key == "script" {
		ERR = "Script error detected."
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
			// Generate randomGrayScale value
			randomGrayScale := rand.New(rand.NewSource(time.Now().UnixNano())).Intn(maxGrayScale)
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
func updatePath(generatedPath string) {
	cliStatus("Updating default output path")
	// Ensure that defaultOutPath has file extension .pgm
	// We need to generate a file with .pgm extension to feed it into a .png converter
	generatedPath = func () string {
		return strings.Replace(generatedPath, "png", EXTENSION, 1)
	}()
	defaultOutPath = generatedPath
}

// Check correct file extension
func checkFileExtension(newPath string, extension string) bool {
    // Split the path
    fullSplitPath := strings.Split(newPath, "/")
    // Get the last element
    lastElement := fullSplitPath[len(fullSplitPath)-1]
    // Split the last element
    splitLastElement := strings.Split(lastElement, ".")
    // Get the last element of the file
    lastElement = splitLastElement[len(splitLastElement)-1]

    // Check if the last element is .pgm
    if lastElement == extension {
        return true
    }
    return false
}

// Function to transform .pgm to .png
// Using a prebuild `Python3` script
// Python Pip dependency: pypng
func convertToPng(inputPath string, outputPath string) {
	cliStatus("Converting to png...")
	// Execute the script to convert .pgm to .png and handle error
	arguments := []string{inputPath, outputPath}
	cmd := exec.Command("python3", scriptPath, arguments[0], arguments[1])
	// RTun script and handle error
	stdout, err := cmd.Output()
	if err != nil {
		warnUsage("script")
	}
	// Handle script STDOUT
	fmt.Println(string(stdout))
}

// Custom CLI status formatted message
func cliStatus(msg string) {
	fmt.Printf("[  %sOK%s  ] %s%s%s\n", GREEN, RESET, PURPLE, msg, RESET)
}