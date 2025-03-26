package cclib

import (
	"bufio"
	"os"
	"strings"
)

/*
	Simplest possible key storage that store keys in text file (one key per line).
*/

type KeyDB struct {
	dbPath string
	keys   map[string]bool
}

func (ks *KeyDB) Init(dbPath string) error {
	ks.dbPath = dbPath
	ks.keys = make(map[string]bool)
	return ks.Load()
}

func (ks *KeyDB) Has(value string) bool {
	return ks.keys[value]
}

func (ks *KeyDB) Add(value string) error {
	if !ks.Has(value) {
		ks.keys[value] = true
		return ks.Save(value)
	}
	return nil
}

func (ks *KeyDB) Load() error {
	file, err := os.OpenFile(ks.dbPath, os.O_RDONLY|os.O_CREATE, 0644)
	if err != nil {
		return err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		ks.keys[strings.TrimSpace(scanner.Text())] = true
	}

	return scanner.Err()
}

func (ks *KeyDB) Save(value string) error {
	file, err := os.OpenFile(ks.dbPath, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0644)
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = file.WriteString(value + "\n")
	return err
}
