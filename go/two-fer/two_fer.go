package twofer

import "fmt"
import "strings"

func ShareWith(name string) string {
    entry := strings.TrimSpace(name)
    if len(entry) == 0 {
        entry = "you"
    }
    return fmt.Sprintf("One for %s, one for me.", entry)
}
