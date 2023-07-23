package cmd

import (
	"github.com/spf13/cobra"

	"github.com/zhaoyunxing92/flexi-code/server/internal/app"
	"github.com/zhaoyunxing92/flexi-code/server/pkg/log"
)

func init() {
	log.NewDefault()
	flags := main.Flags()
	flags.StringVarP(&config, "app", "c", "", "flexi code server app file path")
}

var (
	config string
	main   = &cobra.Command{
		Use:   "code",
		Short: "Run the app",
		PreRun: func(cmd *cobra.Command, args []string) {
			log.Info("flexi code server start...")
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			return start()
		},
	}
)

func Main() {
	if err := main.Execute(); err != nil {
		log.Errorf("rootCmd.Execute err:%v", err)
	}
}

func start() error {
	conf, err := app.Load(config)
	if err != nil {
		return err
	}
	engine, err := initApplication(conf.App, conf.Storage)
	if err != nil {
		return err
	}
	return engine.Run(conf.Server.Address)
}
