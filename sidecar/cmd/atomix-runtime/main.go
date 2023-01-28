// SPDX-FileCopyrightText: 2022-present Intel Corporation
//
// SPDX-License-Identifier: Apache-2.0

package main

import (
	"fmt"
	"github.com/atomix/atomix/runtime/pkg/logging"
	"github.com/atomix/atomix/runtime/pkg/runtime"
	runtimev1 "github.com/atomix/atomix/runtime/pkg/runtime/v1"
	"github.com/atomix/atomix/sidecar/pkg/sidecar"
	"github.com/spf13/cobra"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	logging.SetLevel(logging.DebugLevel)

	cmd := &cobra.Command{
		Use: "atomix-runtime",
		Run: func(cmd *cobra.Command, args []string) {
			host, err := cmd.Flags().GetString("host")
			if err != nil {
				fmt.Fprintln(cmd.OutOrStderr(), err.Error())
				os.Exit(1)
			}
			port, err := cmd.Flags().GetInt("port")
			if err != nil {
				fmt.Fprintln(cmd.OutOrStderr(), err.Error())
				os.Exit(1)
			}
			runtimeHost, err := cmd.Flags().GetString("runtime-host")
			if err != nil {
				fmt.Fprintln(cmd.OutOrStderr(), err.Error())
				os.Exit(1)
			}
			runtimePort, err := cmd.Flags().GetInt("runtime-port")
			if err != nil {
				fmt.Fprintln(cmd.OutOrStderr(), err.Error())
				os.Exit(1)
			}
			pluginsDir, err := cmd.Flags().GetString("plugins")
			if err != nil {
				fmt.Fprintln(cmd.OutOrStderr(), err.Error())
				os.Exit(1)
			}

			// Initialize the runtime
			rt := runtimev1.New(
				runtimev1.WithDriverProvider(sidecar.NewDriverProvider(pluginsDir)))

			// Start the runtime service
			rtSvc := runtime.NewService(rt,
				runtime.WithHost(runtimeHost),
				runtime.WithPort(runtimePort))
			if err := rtSvc.Start(); err != nil {
				fmt.Fprintln(cmd.OutOrStderr(), err.Error())
				os.Exit(1)
			}

			// Start the proxy service
			proxySvc := sidecar.NewService(rt,
				sidecar.WithHost(host),
				sidecar.WithPort(port))
			if err := proxySvc.Start(); err != nil {
				fmt.Fprintln(cmd.OutOrStderr(), err.Error())
				os.Exit(1)
			}

			// Wait for an interrupt signal
			ch := make(chan os.Signal, 2)
			signal.Notify(ch, os.Interrupt, syscall.SIGTERM)
			<-ch

			// Stop the proxy
			if err := proxySvc.Stop(); err != nil {
				fmt.Println(err)
				os.Exit(1)
			}

			// Stop the runtime
			if err := rtSvc.Stop(); err != nil {
				fmt.Println(err)
				os.Exit(1)
			}
		},
	}
	cmd.Flags().String("host", "", "the host to which to bind the proxy server")
	cmd.Flags().Int("port", 5678, "the port to which to bind the proxy server")
	cmd.Flags().String("runtime-host", "", "the host to which to bind the runtime server")
	cmd.Flags().Int("runtime-port", 5679, "the port to which to bind the runtime server")
	cmd.Flags().StringP("plugins", "p", "/var/atomix/plugins", "the path to the plugins directory")

	_ = cmd.MarkFlagDirname("drivers")

	if err := cmd.Execute(); err != nil {
		panic(err)
	}
}
