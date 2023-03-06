package main

import (
        "flag"
        "fmt"
        "os"
        "os/exec"
        "path/filepath"
        "sync"
)

func main() {
        dirPtr := flag.String("dir", ".", "Directory to search for .sh files")
        flag.Parse()


        shFiles, err := filepath.Glob(filepath.Join(*dirPtr, "*.sh"))
        if err != nil {
                fmt.Fprintln(os.Stderr, "Error finding .sh files:", err)
                os.Exit(1)
        }


        var wg sync.WaitGroup


        for _, shFile := range shFiles {

                wg.Add(1)
                go func(shFile string) {
                        cmd := exec.Command("sh", shFile)


                        cmd.Stdout = os.Stdout
                        cmd.Stderr = os.Stderr


                        if err := cmd.Run(); err != nil {
                                fmt.Fprintln(os.Stderr, "Error running command:", err)
                                os.Exit(1)
                        }

                        wg.Done()
                }(shFile)
        }


        wg.Wait()
}
