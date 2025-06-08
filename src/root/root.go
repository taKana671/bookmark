package root

import "github.com/spf13/cobra"

func NewRootCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "bookmark",
		Short: "CLI to bookmark a website.",
		Long:  `Bookmark websites, delete bookmarks, search for bookmarks,
		        and access bookmarked sites.
				Information on bookmarked websites (datetime of bookmarking, category,
				site title, and site URL) is saved in a CSV file that is output
				on the same level as this CLI.`,

		RunE: func(cmd *cobra.Command, args []string) error {
			cmd.Help()
			return nil
		},
	}
	return cmd
}
