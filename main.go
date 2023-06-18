package main

import (
	"github.com/mu-box/butter/api"
	"github.com/mu-box/butter/auth"
	"github.com/mu-box/butter/config"
	"github.com/mu-box/butter/deploy"
	"github.com/mu-box/butter/repo"
	"github.com/mu-box/butter/server"

	"github.com/spf13/cobra"
)

func main() {
	server := true
	configFile := ""
	command := cobra.Command{
		Use:   "butter",
		Short: "butter makes the breads silky smooth",
		Long:  `Butter is a solid dairy product made by churning fresh or fermented cream or milk, to separate the butterfat from the buttermilk. It is generally used as a spread on plain or toasted bread products and a condiment on cooked vegetables, as well as in cooking, such as baking, sauce making, and pan frying. Butter consists of butterfat, milk proteins and water.`,
		Run: func(ccmd *cobra.Command, args []string) {
			if !server {
				ccmd.HelpFunc()(ccmd, args)
				return
			}
			if configFile != "" {
				config.Parse(configFile)
			}
			serverStart()
		},
	}
	config.AddFlags(&command)

	command.Flags().BoolVarP(&server, "server", "s", false, "Run as server")
	command.Flags().StringVarP(&configFile, "configFile", "", "", "[server] config file location")

	// when we create a cli i will add it here
	// cli.AddCli(command)

	command.Execute()

}

func serverStart() {
	err := deploy.Setup()
	if err != nil {
		panic(err)
	}
	err = repo.Setup()
	if err != nil {
		panic(err)
	}
	err = auth.Setup()
	if err != nil {
		panic(err)
	}
	sshServer, err := server.StartServer()
	if err != nil {
		panic(err)
	}
	defer sshServer.Close()

	err = api.Start()
	if err != nil {
		panic(err)
	}
}
