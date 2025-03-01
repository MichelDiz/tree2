package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/spf13/cobra"
)

var (
	dir         string
	fileExt     string
	outputFile  string
	depthLimit  int
	ignoreList  []string
	jsonOutput  bool
	showPerms   bool
	showModTime bool
	showSize    bool
	ignoredDirs = map[string]bool{
		".git":         true,
		"node_modules": true,
		"tmp":          true,
	}
)

var rootCmd = &cobra.Command{
	Use:   "tree-cli",
	Short: "Displays the directory and file tree",
	Long:  "tree-cli is a simple tool to display the structure of directories and files.",
	RunE: func(cmd *cobra.Command, args []string) error {
		var output *os.File
		var err error
		if outputFile != "" {
			output, err = os.Create(outputFile)
			if err != nil {
				return fmt.Errorf("error creating output file: %w", err)
			}
			defer output.Close()
		}

		absRoot, err := filepath.Abs(dir)
		if err != nil {
			return fmt.Errorf("error getting absolute path: %w", err)
		}

		return printTree(absRoot, "", fileExt, output, absRoot, 0)
	},
}

func init() {
	rootCmd.Flags().StringVarP(&dir, "directory", "d", ".", "Directory to display file tree")
	rootCmd.Flags().StringVarP(&fileExt, "extension", "e", "", "File extension to be displayed")
	rootCmd.Flags().StringVarP(&outputFile, "output", "o", "", "File to save output")
	rootCmd.Flags().IntVarP(&depthLimit, "depth", "l", -1, "Depth limit for display")
	rootCmd.Flags().StringSliceVarP(&ignoreList, "ignore", "i", []string{}, "List of files/directories to ignore")
	rootCmd.Flags().BoolVarP(&jsonOutput, "json", "j", false, "Export tree in JSON format")
	rootCmd.Flags().BoolVar(&showPerms, "permissions", false, "Show file permissions")
	rootCmd.Flags().BoolVar(&showModTime, "modtime", false, "Show file modification time")
	rootCmd.Flags().BoolVar(&showSize, "size", false, "Show file size")
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, "Error:", err)
		os.Exit(1)
	}
}

func printTree(root string, prefix string, fileExt string, output *os.File, basePath string, currentDepth int) error {
	if depthLimit != -1 && currentDepth > depthLimit {
		return nil
	}

	entries, err := os.ReadDir(root)
	if err != nil {
		return fmt.Errorf("error reading directory %s: %w", root, err)
	}

	for i, entry := range entries {
		if ignoredDirs[entry.Name()] || isIgnored(entry.Name()) {
			continue
		}

		info, _ := entry.Info()
		connString := ""
		if showPerms {
			connString += fmt.Sprintf(" [%s]", info.Mode().String())
		}
		if showModTime {
			connString += fmt.Sprintf(" [%s]", info.ModTime().Format("2006-01-02 15:04:05"))
		}
		if showSize {
			connString += fmt.Sprintf(" [%s]", formatSize(info.Size()))
		}

		connector := "├──"
		if i == len(entries)-1 {
			connector = "└──"
		}
		relPath, _ := filepath.Rel(basePath, filepath.Join(root, entry.Name()))
		line := fmt.Sprintf("%s%s%s %s", prefix, connector, connString, relPath)
		fmt.Println(line)
		if output != nil {
			_, _ = output.WriteString(line + "\n")
		}

		if entry.IsDir() {
			nextPrefix := prefix + "│   "
			if i == len(entries)-1 {
				nextPrefix = prefix + "    "
			}
			if err := printTree(filepath.Join(root, entry.Name()), nextPrefix, fileExt, output, basePath, currentDepth+1); err != nil {
				return err
			}
		} else if fileExt != "" && strings.HasSuffix(entry.Name(), fileExt) {
			printFileContent(filepath.Join(root, entry.Name()), output, basePath)
		}
	}
	return nil
}

func printFileContent(filename string, output *os.File, basePath string) {
	relPath, _ := filepath.Rel(basePath, filename)
	content, err := os.ReadFile(filename)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error reading file %s: %v\n", relPath, err)
		return
	}
	formattedContent := fmt.Sprintf("\n=== Content of %s ===\n%s", relPath, string(content))
	fmt.Println(formattedContent)
	if output != nil {
		_, _ = output.WriteString(formattedContent + "\n")
	}
}

func formatSize(size int64) string {
	if size < 1024 {
		return fmt.Sprintf("%dB", size)
	} else if size < 1024*1024 {
		return fmt.Sprintf("%.1fK", float64(size)/1024)
	} else if size < 1024*1024*1024 {
		return fmt.Sprintf("%.1fM", float64(size)/(1024*1024))
	} else {
		return fmt.Sprintf("%.1fG", float64(size)/(1024*1024*1024))
	}
}

func isIgnored(name string) bool {
	for _, pattern := range ignoreList {
		if strings.Contains(name, pattern) {
			return true
		}
	}
	return false
}
