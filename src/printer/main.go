package printer

import "fmt"

func PrintMissingRequired(files []string) {
	for _, item := range files {
		fmt.Println("[COMPLIANCE FAILED] >", item, "is missing.")
	}
}

func PrintSuccessfulRequired() {
	fmt.Println("[COMPLIANCE SUCCESSFUL] >", "All required files present.")
}

func PrintStructureFailed(files []string) {
	fmt.Println("[COMPLIANCE FAILED] for following items:")

	for _, item := range files {
		fmt.Println(">", item, "is not allowed.")
	}
}

func PrintStructureSuccess() {
	fmt.Println("[COMPLIANCE SUCCESSFUL] >", "Source code matches expected patterns.")
}
