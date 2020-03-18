package main

import (
	"fmt"
	"net/http"
	"os"
	"strconv"

	"github.com/spf13/cobra"

	"ipsec-exporter/ipsec"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/prometheus/common/log"
)

const (
	flagIpsecConfigFile  = "config-path"
	flagWebListenAddress = "web.listen-address"
)

var IpSecConfigFile string
var WebListenAddress int
var ipSecConfiguration *ipsec.Configuration
var Version string

// Package cobra is a commander providing a simple interface to create powerful modern CLI interfaces
var RootCmd = &cobra.Command{
	Use:     "ipsec_exporter",
	Short:   "Prometheus exporter for ipsec status.",
	Long:    "",
	Run:     defaultCommand,
	Version: Version,
}

func Server() {

	var err error
	ipSecConfiguration, err = ipsec.NewConfiguration(IpSecConfigFile)
	if err != nil {
		log.Fatal(err)
		return
	}
	if !ipSecConfiguration.HasTunnels() {
		log.Warn("No configured connections in " + IpSecConfigFile)
	}

	collector := ipsec.NewCollector(ipSecConfiguration)
	prometheus.MustRegister(collector)

	http.Handle("/metrics", promhttp.Handler())

	log.Infoln("Listening on", WebListenAddress)
	err = http.ListenAndServe(":"+strconv.Itoa(WebListenAddress), nil)
	if err != nil {
		log.Fatal(err)
	}
}

func init() {
	RootCmd.PersistentFlags().StringVar(&IpSecConfigFile, flagIpsecConfigFile, "/etc/ipsec.conf", "Path to the ipsec config file.")
	RootCmd.PersistentFlags().IntVar(&WebListenAddress, flagWebListenAddress, 10437, "Address on which to expose metrics.")
}

func Execute() {
	if err := RootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func defaultCommand(_ *cobra.Command, _ []string) {
	Server()
}

func main() {
	Execute()
}
