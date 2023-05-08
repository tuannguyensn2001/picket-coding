package cmd

import (
	"bufio"
	"fmt"
	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
	"os"
	"os/exec"
)

func execute() *cobra.Command {
	return &cobra.Command{
		Use:   "execute",
		Short: "Execute a command",
		Run: func(cmd *cobra.Command, args []string) {
			userId := 1
			ex := "1"
			content := "function sum(a, b) {\n  return a + b\n}\n\nconsole.log(sum(1, 3))"

			if err := os.Mkdir(fmt.Sprintf("src/storage/coding/ex%s/user_%d", ex, userId), 0777); err != nil {
				//log.Fatal().Err(err).Send()
			}

			f, err := os.Create(fmt.Sprintf("src/storage/coding/ex%s/user_%d/index.js", ex, userId))
			if err != nil {
				log.Fatal().Err(err).Send()
			}
			defer f.Close()
			w := bufio.NewWriter(f)
			_, err = w.WriteString(content)
			if err != nil {
				log.Fatal().Err(err).Send()
			}
			w.Flush()

			command := fmt.Sprintf("docker run --rm -v $(pwd)/src/storage/coding/ex%s/user_%d:/tmp picket-js node /tmp/index.js /tmp/output", ex, userId)
			err = exec.Command("sh", "-c", command).Run()
			if err != nil {
				log.Fatal().Err(err).Send()
				return
			}
		},
	}
}
