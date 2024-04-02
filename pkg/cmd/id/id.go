package id

import (
	"github.com/MakeNowJust/heredoc"
	"github.com/mskelton/go-template/internal/utils"
	"github.com/spf13/cobra"
)

var IdCmd = &cobra.Command{
	Use:   "id",
	Short: "Generate a new id",
	Long:  heredoc.Doc(``),
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Println(utils.GenerateId())
	},
}
