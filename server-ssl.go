package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"strings"
)

/*
{
    "untrustedCertResponseControl": "drop",
    "uncleanShutdown": "enabled",
    "strictResume": "disabled",
    "sslSignHash": "any",
    "sslForwardProxyBypass": "disabled",
    "sslForwardProxy": "disabled",
    "sniRequire": "false",
    "sniDefault": "false",
    "expireCertResponseControl": "drop",
    "defaultsFrom": "/Common/serverssl",
    "ciphers": "DEFAULT",
    "chain": "/Common/myob-chain.crt",
    "cert": "/Common/sit.store.myob.com.au.crt",
    "cacheTimeout": 3600,
    "cacheSize": 262144,
    "authenticateDepth": 9,
    "kind": "tm:ltm:profile:server-ssl:server-sslstate",
    "name": "sit.store.myob.com.au",
    "partition": "Common",
    "fullPath": "/Common/sit.store.myob.com.au",
    "generation": 690,
    "selfLink": "https://localhost/mgmt/tm/ltm/profile/server-ssl/~Common~sit.store.myob.com.au?ver=11.6.0",
    "alertTimeout": "10",
    "authenticate": "once",
    "genericAlert": "enabled",
    "handshakeTimeout": "10",
    "key": "/Common/sit.store.myob.com.au.key",
    "modSslMethods": "disabled",
    "mode": "enabled",
    "tmOptions": [
      "dont-insert-empty-fragments"
    ],
    "peerCertMode": "ignore",
    "proxySsl": "disabled",
    "proxySslPassthrough": "disabled",
    "renegotiatePeriod": "indefinite",
    "renegotiateSize": "indefinite",
    "renegotiation": "enabled",
    "retainCertificate": "true",
    "secureRenegotiation": "require-strict",
    "sessionMirroring": "disabled",
    "sessionTicket": "disabled"
  },
*/

type LBServerSsl struct {
	Name                         string   `json:"name"`
	Partition                    string   `json:"partition"`
	Fullpath                     string   `json:"fullPath"`
	Generation                   int      `json:"generation"`
	UntrustedCertResponseControl string   `json:"untrustedCertResponseControl"`
	UncleanShutdown              string   `json:"uncleanShutdown"`
	StrictResume                 string   `json:"strictResume"`
	SslSignHash                  string   `json:"sslSignHash"`
	SslForwardProxyBypass        string   `json:"sslForwardProxyBypass"`
	SslForwardProxy              string   `json:"sslForwardProxy"`
	SniRequire                   string   `json:"sniRequire"`
	SniDefault                   string   `json:"sniDefault"`
	ExpireCertResponseControl    string   `json:"expireCertResponseControl"`
	DefaultsFrom                 string   `json:"defaultsFrom"`
	Ciphers                      string   `json:"ciphers"`
	Chain                        string   `json:"chain"`
	Cert                         string   `json:"cert"`
	Key                          string   `json:"key"`
	CacheTimeout                 int      `json:"cacheTimeout"`
	CacheSize                    int      `json:"cacheSize"`
	AuthenticateDepth            int      `json:"authenticateDepth"`
	AlertTimeout                 string   `json:"alertTimeout"`
	SelfLink                     string   `json:"selfLink"`
	Authenticate                 string   `json:"authenticate"`
	GenericAlert                 string   `json:"genericAlert"`
	HandshakeTimeout             string   `json:"handshakeTimeout"`
	ModSslMethods                string   `json:"modSslMethods"`
	Mode                         string   `json:"mode"`
	TmOptions                    []string `json:"tmOptions"`
	PeerCertMode                 string   `json:"peerCertMode"`
	ProxySsl                     string   `json:"proxySsl"`
	ProxySslPassthrough          string   `json:"proxySslPassthrough"`
	RenegotiatePeriod            string   `json:"renegotiatePeriod"`
	RenegotiateSize              string   `json:"renegotiateSize"`
	Renegotiation                string   `json:"renegotiation"`
	RetainCertificate            string   `json:"retainCertificate"`
	SecureRenegotiation          string   `json:"secureRenegotiation"`
	SessionMirroring             string   `json:"sessionMirroring"`
	SessionTicket                string   `json:"sessionTicket"`
}

type LBServerSsls struct {
	Items []LBServerSsl `json:"items"`
}

func showServerSsls() {

	u := "https://" + f5Host + "/mgmt/tm/ltm/profile/server-ssl"
	res := LBServerSsls{}

	err, resp := SendRequest(u, GET, &sessn, nil, &res)
	if err != nil {
		log.Fatalf("%s : %s\n", resp.HttpResponse().Status, err)
	}

	for _, v := range res.Items {
		fmt.Printf("%s\n", v.Fullpath)
	}
}

func showServerSsl(sname string) {

	server := strings.Replace(sname, "/", "~", -1)
	u := "https://" + f5Host + "/mgmt/tm/ltm/profile/server-ssl/" + server
	res := LBServerSsl{}

	err, resp := SendRequest(u, GET, &sessn, nil, &res)
	if err != nil {
		log.Fatalf("%s : %s\n", resp.HttpResponse().Status, err)
	}
	printResponse(&res)

}

func addServerSsl() {

	u := "https://" + f5Host + "/mgmt/tm/ltm/profile/server-ssl"
	res := LBServerSsl{}
	// we use raw so we can modify the input file without using a struct
	// use of a struct will send all available fields, some of which can't be modified
	body := json.RawMessage{}

	// read in json file
	dat, err := ioutil.ReadFile(f5Input)
	if err != nil {
		log.Fatal(err)
	}

	// convert json to a node struct
	err = json.Unmarshal(dat, &body)
	if err != nil {
		log.Fatal(err)
	}

	// post the request
	err, resp := SendRequest(u, POST, &sessn, &body, &res)
	if err != nil {
		log.Fatalf("%s : %s\n", resp.HttpResponse().Status, err)
	}
	printResponse(&res)
}

func updateServerSsl(sname string) {

	server := strings.Replace(sname, "/", "~", -1)
	u := "https://" + f5Host + "/mgmt/tm/ltm/profile/server-ssl/" + server
	res := LBServerSsl{}
	body := json.RawMessage{}

	// read in json file
	dat, err := ioutil.ReadFile(f5Input)
	if err != nil {
		log.Fatal(err)
	}

	// convert json to a node struct
	err = json.Unmarshal(dat, &body)
	if err != nil {
		log.Fatal(err)
	}

	// put the request
	err, resp := SendRequest(u, PUT, &sessn, &body, &res)
	if err != nil {
		log.Fatalf("%s : %s\n", resp.HttpResponse().Status, err)
	}
	printResponse(&res)
}

func deleteServerSsl(sname string) {

	server := strings.Replace(sname, "/", "~", -1)
	u := "https://" + f5Host + "/mgmt/tm/ltm/profile/server-ssl/" + server
	res := json.RawMessage{}

	err, resp := SendRequest(u, DELETE, &sessn, nil, &res)
	if err != nil {
		log.Fatalf("%s : %s\n", resp.HttpResponse().Status, err)
	} else {
		log.Printf("%s : %s deleted\n", resp.HttpResponse().Status, sname)
	}

}
