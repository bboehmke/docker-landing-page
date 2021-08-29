package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
)

func main() {
	ctx := context.Background()
	cli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	if err != nil {
		panic(err)
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		log.Printf("GET from %s", r.RemoteAddr)
		containers, err := cli.ContainerList(ctx, types.ContainerListOptions{})
		if err != nil {
			panic(err)
		}

		var lines []string

		for _, container := range containers {
			if strings.ToLower(container.Labels["landing-page.enabled"]) != "true" {
				continue
			}

			var name string
			if v, ok := container.Labels["landing-page.name"]; ok {
				name = v
			} else {
				if len(container.Names) > 0 {
					name = strings.TrimPrefix(container.Names[0], "/")
				} else {
					name = container.Image
				}
			}

			port, err := strconv.Atoi(container.Labels["landing-page.port"])
			if err != nil {
				if len(container.Ports) > 0 {
					port = int(container.Ports[0].PublicPort)
				}
			}

			lines = append(lines,
				fmt.Sprintf("<b><a href=\"http://%s:%d\">%s</a></b>",
					strings.Split(r.Host, ":")[0],
					port,
					name))
		}

		w.WriteHeader(200)

		_, _ = fmt.Fprint(w, "<html><br/><ul>")
		for _, entry := range lines {
			_, _ = fmt.Fprintf(w, "<li>%s</li>", entry)
		}
		_, _ = fmt.Fprint(w, "</ul></html>")
	})

	log.Printf("Start server")
	err = http.ListenAndServe(":8080", nil)
	if err != nil {
		panic(err)
	}
}
