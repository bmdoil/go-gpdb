package main

var (
	programName    = "gpdb"
	programVersion = "3.0.1-dl"
)

func main() {
	// Execute the cobra CLI & run the program
	rootCmd.Execute()
}
