package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"
	"sync"
	"time"

	_ "github.com/lib/pq"
	"github.com/radovskyb/watcher"
	"gopkg.in/yaml.v2"
)

type Config struct {
	Path     string   `yaml:"path"`
	Commands []string `yaml:"commands"`
}

func parseConfiguration(filename string) ([]Config, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	var configArr []Config
	if err := yaml.NewDecoder(file).Decode(&configArr); err != nil {
		return nil, err
	}
	return configArr, nil
}

type Database struct {
	*sql.DB
}

func (db *Database) Init() error {
	var err error
	db.DB, err = sql.Open("postgres", "user=postgres password=12345 dbname=postgres sslmode=disable")
	if err != nil {
		return err
	}

	_, err = db.Exec("CREATE TABLE IF NOT EXISTS file_changes(" +
		"id serial CONSTRAINT file_changes_pk PRIMARY KEY," +
		"path TEXT NOT NULL, " +
		"event TEXT NOT NULL);")
	if err != nil {
		return err
	}
	return nil
}

func (db *Database) Close() error {
	return db.DB.Close()
}

func (db *Database) InsertFileChange(event watcher.Event) error {
	_, err := db.Exec("INSERT INTO file_changes (path, event) VALUES ($1, $2)", event.Path, event.String())
	if err != nil {
		return err
	}
	return nil
}

func RunCommands(config Config) error {
	for _, cmd := range config.Commands {
		args := strings.Fields(cmd)
		name := args[0]
		args = args[1:]

		cmd := exec.Command(name, args...)
		if err := cmd.Run(); err != nil {
			return err
		}
		fmt.Println("Command done:", cmd)
	}
	return nil
}

func watch(config Config, db *Database, wg *sync.WaitGroup) {
	defer wg.Done()

	w := watcher.New()
	w.SetMaxEvents(1)
	w.FilterOps(watcher.Create, watcher.Move, watcher.Remove, watcher.Rename, watcher.Write, watcher.Chmod)
	if err := w.Add(config.Path); err != nil {
		log.Fatal(err)
	}

	go func() {
		for {
			select {
			case event := <-w.Event:
				if err := RunCommands(config); err != nil {
					log.Println(err)
				}
				if err := db.InsertFileChange(event); err != nil {
					log.Fatal(err)
				}
			case err := <-w.Error:
				log.Println(err)
			case <-w.Closed:
				return
			}
		}
	}()

	if err := w.Start(time.Second * 1); err != nil {
		log.Fatal(err)
	}
}

func main() {

	var wg sync.WaitGroup

	db := &Database{}
	if err := db.Init(); err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	configArr, err := parseConfiguration("config.yaml")
	if err != nil {
		log.Fatal(err)
	}

	for _, config := range configArr {
		wg.Add(1)
		go watch(config, db, &wg)
	}

	wg.Wait()
}
