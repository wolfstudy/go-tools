// Licensed to the Apache Software Foundation (ASF) under one
// or more contributor license agreements.  See the NOTICE file
// distributed with this work for additional information
// regarding copyright ownership.  The ASF licenses this file
// to you under the Apache License, Version 2.0 (the
// "License"); you may not use this file except in compliance
// with the License.  You may obtain a copy of the License at
//
//   http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing,
// software distributed under the License is distributed on an
// "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
// KIND, either express or implied.  See the License for the
// specific language governing permissions and limitations
// under the License.

package tools

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"github.com/wolfstudy/go-tools/context"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"testing"
)

func TestDeduplicate(t *testing.T) {
	dedupSlice := []string{"apple", "banana", "cherry", "cherry"}
	newSlice := Deduplicate(dedupSlice)
	fmt.Println(newSlice)

	exceptSlice := []string{"apple", "banana", "cherry"}
	assert.Equal(t, exceptSlice, newSlice)
}

type Config struct {
	Kind           string `yaml:"kind,omitempty"`
	APIVersion     string `yaml:"apiVersion,omitempty"`
	CurrentContext string `yaml:"current-context"`
}

// operation yaml file
func TestReadFile(t *testing.T) {
	conf := &Config{
		Kind:           "aaa",
		APIVersion:     "bbb",
		CurrentContext: "ccc",
	}
	home := context.HomeDir()
	err := WriteToFile(conf, home+"/github.com/wolfstudy/go-tools/tools/config")
	if err != nil {
		log.Fatal(err)
	}
}

func WriteToFile(obj interface{}, filename string) error {
	out, err := yaml.Marshal(obj.(*Config))
	if err != nil {
		log.Fatalf("err: %v", err)
	}
	if err != nil {
		return err
	}
	dir := filepath.Dir(filename)
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		if err = os.MkdirAll(dir, 0755); err != nil {
			return err
		}
	}

	if err := ioutil.WriteFile(filename, out, 0600); err != nil {
		return err
	}
	return nil
}
