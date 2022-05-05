// Copyright 2022 Paul Greenberg greenpau@outlook.com
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package main

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"

	"github.com/google/uuid"
	"github.com/greenpau/go-faker/pkg/persons"
	"github.com/urfave/cli/v2"
)

type User struct {
	ID       string   `json:"id,omitempty" xml:"id,omitempty" yaml:"id,omitempty"`
	Username string   `json:"username,omitempty" xml:"username,omitempty" yaml:"username,omitempty"`
	Password string   `json:"password,omitempty" xml:"password,omitempty" yaml:"password,omitempty"`
	Name     string   `json:"name,omitempty" xml:"name,omitempty" yaml:"name,omitempty"`
	Email    string   `json:"email,omitempty" xml:"email,omitempty" yaml:"email,omitempty"`
	Roles    []string `json:"roles,omitempty" xml:"roles,omitempty" yaml:"roles,omitempty"`
}

// NewUser generates a new user.
func NewUser(c *cli.Context) error {
	domainName := c.String("domain")
	password := c.String("password")
	count := c.Int("count")
	roles := c.StringSlice("roles")
	emailAddrPrefix := c.String("email-prefix")

	if count == 0 {
		count++
	}
	if domainName == "" {
		domainName = "localdomain.local"
	}

	for i := 1; i <= count; i++ {
		user := &User{
			ID: uuid.New().String(),
		}
		sex := randomSex()
		firstName := strings.Title(strings.ToLower(randomFirstName(sex)))
		lastName := strings.Title(strings.ToLower(randomLastName()))
		user.Name = fmt.Sprintf("%s %s", firstName, lastName)
		user.Username = generateUsername(lastName, firstName)
		user.Password = password
		user.Email = fmt.Sprintf("%s@%s", user.Username, domainName)
		if emailAddrPrefix != "" {
			user.Email = emailAddrPrefix + user.Email
		}
		user.Roles = roles

		b, err := json.Marshal(user)
		if err != nil {
			return err
		}
		fmt.Fprintf(os.Stdout, string(b)+"\n")
	}

	return nil
}

func randomSex() string {
	if rand.Intn(2) == 1 {
		return "M"
	}
	return "F"
}

func randomFirstName(sex string) string {
	if sex == "F" {
		return randomFemaleFirstName()
	}
	return randomMaleFirstName()
}

func randomFemaleFirstName() string {
	return randomString(persons.FemaleNames)
}

func randomMaleFirstName() string {
	return randomString(persons.MaleNames)
}

func randomLastName() string {
	return randomString(persons.LastNames)
}

func randomString(arr []string) string {
	rand.Seed(time.Now().UnixNano())
	i := rand.Intn(len(arr))
	return string(arr[i])
}

func generateUsername(names ...string) string {
	var output string
	flat := strings.Join(names, "")
	if len(flat) < 9 {
		return flat
	}

	r := 8
	for _, n := range names {
		if r < 1 {
			break
		}
		if len(n) > 5 {
			output += n[:5]
			r -= 5
			continue
		}
		output += n
		r -= len(n)
	}

	output = strings.ToLower(output)
	if len(output) > 8 {
		output = output[:8]
	}
	return output
}
