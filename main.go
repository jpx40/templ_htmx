package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/a-h/templ"
	_ "github.com/jackc/pgx/v5/stdlib"
)

type Package struct {
	name         string
	version      string
	dependencies []string
}

func main() {
	t0 := time.Now()
	s := db(db_connect())

	// foot := foot()
	// count := button()
	// stencil := component()
	// dropdown := dropdown()

	var str_list []string
	for _, t := range s {
		str_list = append(str_list, t.name)
	}
	http.Handle("/", templ.Handler(layout(head(), grid(s), foot())))

	http.ListenAndServe(":8080", nil)

	t1 := time.Now()
	fmt.Printf("The call took %v to run.\n", t1.Sub(t0))
}

// func render() {
// 	f, err := os.Create("hello.html")
// 	if err != nil {
// 		log.Fatalf("failed to create output file: %v", err)
// 	}

// 	err = layout().Render(context.Background(), f)
// 	if err != nil {
// 		log.Fatalf("failed to write output file: %v", err)
// 	}
// }

func db_connect() *sql.DB {
	// c := "postgres://postgres:artemis34@127.17.0.2:5432//postgres"
	c := "host=0.0.0.0 user=jonas password=artemis34 dbname=postgres port=32768 sslmode=disable"

	conn, err := sql.Open("pgx", c)
	if err != nil {
		panic("error opening postgres connection: " + err.Error())
	}

	return conn
}

func db(c *sql.DB) []Package {
	rows, err := c.Query("SELECT * FROM arch_packages")
	if err != nil {
		panic(err)
	}

	defer rows.Close()

	pkgs := []Package{}
	for rows.Next() {
		var pkg_name string
		var version string
		var dependencies_b []byte
		var dependencies []string
		if err := rows.Scan(&pkg_name, &version, &dependencies_b); err != nil {
			panic(err)
		}
		err := json.Unmarshal(dependencies_b, &dependencies)
		if err != nil {
			fmt.Println("error:", err)
		}
		pkgs = append(pkgs, Package{name: pkg_name, version: version, dependencies: dependencies})

	}

	return pkgs
}
