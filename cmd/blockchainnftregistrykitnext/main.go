// cmd/blockchainnftregistrykitnext/main.go
package main

import (
"flag"
"log"
"os"

"blockchainnftregistrykitnext/internal/blockchainnftregistrykitnext"
)

func main() {
verbose := flag.Bool("verbose", false, "Enable verbose logging")
flag.Parse()

app := blockchainnftregistrykitnext.NewApp(*verbose)
if err := app.Run(); err != nil {
log.Fatal(err)
}
}
