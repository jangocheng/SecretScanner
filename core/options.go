package core

import (
	"flag"
	"os"
)

const (
	TempDirSuffix = "SecretScanning"
	ExtractedImageFilesDir = "ExtractedFiles"
)

type Options struct {
	Threads          *int
	DebugLevel       *string
	MaximumFileSize  *uint
	TempDirectory    *string
	Local            *string
	ConfigPath       *string
	OutputPath       *string
	ImageName        *string
	MultipleMatch    *bool
	MaxMultiMatch    *uint
	MaxSecrets       *uint
}

func ParseOptions() (*Options, error) {
	options := &Options{
		Threads:             flag.Int("threads", 0, "Number of concurrent threads (default number of logical CPUs)"),
		DebugLevel:          flag.String("debug-level", "ERROR", "Debug levels are one of FATAL, ERROR, IMPORTANT, WARN, INFO, DEBUG. Only levels higher than the debug-level are displayed"),
		MaximumFileSize:     flag.Uint("maximum-file-size", 256, "Maximum file size to process in KB"),
		TempDirectory:       flag.String("temp-directory", os.TempDir(), "Directory to process and store repositories/matches"),
		Local:               flag.String("local", "", "Specify local directory (absolute path) which to scan. Scans only given directory recursively."),
		ConfigPath:          flag.String("config-path", "", "Searches for config.yaml from given directory. If not set, tries to find it from SecretScanner binary's and current directory"),
		OutputPath:          flag.String("output-path", "", "Outputs json file with secrets to this dir/file. If not set, it will output to a default filename in current directory"),
		ImageName:           flag.String("image-name", "", "Name of the image along with tag to scan for secrets"),
		MultipleMatch:       flag.Bool("multi-match", false, "Output multiple matches of same pattern in one file. By default, only one match of a pattern is output for a file for better performance"),
		MaxMultiMatch:       flag.Uint("max-multi-match", 3, "Maximum number of matches of same pattern in one file. This is used only when multi-match option is enabled."),
		MaxSecrets:          flag.Uint("max-secrets", 1000, "Maximum number of secrets to find in one container image or file system."),
	}
	flag.Parse()
	return options, nil
}
