/*
Tiny Web server for testing k8s cluster, by Bryan Lee, 2017-08-31
*/

// A tiny web server for viewing the environment kubernetes creates for your
// containers. It exposes the filesystem and environment variables via http
// server.
//
// Modified from explorer tiny web server of k8s team.

package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
)

const (
        // Name of the application
        Name = "Tiny Webserver"
        // Version of the application
        Version = "0.71"
	// Default port
	DefaultPort = 8888
)

func main() {

	var (
		portNum = flag.Int("port", DefaultPort, "Port number to serve at.")
	)

	flag.Parse()

	// Getting hostname of Docker node or Podname
	podname, err := os.Hostname()
	if err != nil {
		log.Fatalf("Error getting hostname: %v", err)
	}

	links := []struct {
		link, desc string
	}{
		{"/", "Default landing page"},
		{"/info", "Show version & usage"},
		{"/fs", "Complete file system as seen by this container."},
		{"/env", "Environment variables as seen by this container."},
		{"/podname", "Podname or Hostname as seen by this container."},
		{"/healthz", "Just respond 200 ok for health checks"},
		{"/quit", "Cause this container to exit."},
	}

	// Handlers
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "<b> Pod/Hostname: %s Port: %d</b><br/><br/>", podname, *portNum)
	})

	http.HandleFunc("/info", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "<b> Name: %s Version: %s Port: %d</b><br/><br/>", Name, Version, *portNum)
		fmt.Fprintf(w, "<b> Kubernetes environment explorer usage</b><br/><br/>")
		for _, v := range links {
			fmt.Fprintf(w, `<a href="%v">%v: %v</a><br/>`, v.link, v.link, v.desc)
		}
	})

	http.Handle("/fs", http.StripPrefix("/fs", http.FileServer(http.Dir("/"))))
	http.HandleFunc("/env", func(w http.ResponseWriter, r *http.Request) {
		for _, v := range os.Environ() {
			fmt.Fprintf(w, "%v\n", v)
		}
	})
	http.HandleFunc("/podname", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, podname)
	})
	http.HandleFunc("/healthz", func(w http.ResponseWriter, r *http.Request) {
                w.WriteHeader(200)
		fmt.Fprintf(w, "ok")
	})
	http.HandleFunc("/quit", func(w http.ResponseWriter, r *http.Request) {
		os.Exit(0)
	})

	// Start and listen
	go log.Fatal(http.ListenAndServe(fmt.Sprintf("0.0.0.0:%d", *portNum), nil))

	select {}
}
