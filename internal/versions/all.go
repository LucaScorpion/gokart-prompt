package versions

import (
	"sync"
)

var sections = []section{
	// Languages
	Dotnet,
	Go,
	Java,
	Lua,
	Node,
	Php,
	Python,
	Ruby,
	Rust,
	// Tools
	Docker,
}

func All(wdFiles []string) string {
	var wg sync.WaitGroup
	parts := make([]string, len(sections))

	for i, s := range sections {
		wg.Add(1)

		i := i
		s := s
		go func() {
			parts[i] = s.version(wdFiles)
			wg.Done()
		}()
	}
	wg.Wait()

	result := ""
	for _, part := range parts {
		if part != "" {
			result += " " + part
		}
	}

	return result
}
