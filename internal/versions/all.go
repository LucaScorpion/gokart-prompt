package versions

import (
	"strings"
	"sync"
)

var sections = []section{
	// Languages
	Go,
	Java,
	Node,
	Php,
	Ruby,
	Rust,
	// Tools
	Docker,
}

func All(wdFiles []string) string {
	var wg sync.WaitGroup
	result := make([]string, len(sections))

	for i, s := range sections {
		wg.Add(1)

		i := i
		s := s
		go func() {
			result[i] = s.version()
			wg.Done()
		}()
	}

	wg.Wait()
	return strings.Join(result, " ")
}
