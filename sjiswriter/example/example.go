package main

import (
	"fmt"
	"log"
	"os"
	"strings"

	"golang.org/x/text/encoding/japanese"
	"golang.org/x/text/transform"

	"github.com/ktakenaka/charsetsupport/sjiswriter"
)

func main() {
	out := os.Stdout
	defer out.Close()

	w := sjiswriter.NewSJISWriter(
		transform.NewWriter(out, japanese.ShiftJIS.NewEncoder()),
	)

	str := strings.Join([]string{"〜", "−", "¢", "£", "¬", "–", "—", "‖", "‾", "ø", "›", "び"}, " ")
	if _, err := w.Write([]byte(str)); err != nil {
		log.Println(fmt.Errorf("failed to write: %w", err))
		return
	}
}
