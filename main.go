// Copyright 2023 Andrey Martyanov. All Rights Reserved.
// Copyright 2017 Google Inc. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package main

import (
	"flag"
	"io/ioutil"
	"log"
	"net/http"
)

var (
	address     = flag.String("address", "localhost:8080", "address to listen on")
	configPath  = flag.String("config", "vans.yaml", "path to config file")
	tlsCertPath = flag.String("tls-cert", "", "path to TLS certificate file")
	tlsKeyPath  = flag.String("tls-key", "", "path to TLS key file")
)

func main() {
	flag.Parse()

	vanity, err := ioutil.ReadFile(*configPath)
	if err != nil {
		log.Fatal(err)
	}

	h, err := newHandler(vanity)
	if err != nil {
		log.Fatal(err)
	}
	http.Handle("/", h)

	server := &http.Server{Addr: *address}
	if *tlsCertPath != "" && *tlsKeyPath != "" {
		err = server.ListenAndServeTLS(*tlsCertPath, *tlsKeyPath)
	} else {
		err = server.ListenAndServe()
	}
	if err != nil {
		log.Fatal(err)
	}
}
