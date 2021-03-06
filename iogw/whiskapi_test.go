package main

import (
	"encoding/json"
	"strings"
)

// Manual test
// To run this test you need to init wit
// iosdk stop
// iosdk init --whisk-apikey=23bc46b1-71f6-4ed5-8c54-816aa4f8c502:123zO3xZCLrMN6v2BKK1dXYFpXlPkccOFqm12CdAsMgRU4VrNZ9lyGVCGuMDGIwP
// iosdk start
// then remove the "_"
func _ExamplePackageUpdate() {
	run("rm -Rvf /tmp/iogw-test ; mkdir /tmp/iogw-test")
	ConfigLoad()
	DryRunPush("123456")
	Configure("/tmp/iogw-test/javascript")
	ConfigLoad()
	Config.WhiskAPIKey = DefaultWhiskAPIKey
	m := ConfigMap()
	kv := whiskConfigKeyValues(m)
	grep("Key:", kv)
	rm := WhiskUpdatePackageParameters("iosdk", m)
	r, _ := json.MarshalIndent(rm, "", " ")
	grep("key", strings.Split(string(r), "\n"))
	// Unordered output:
	// Wrote /Users/michelesciabarra/.iosdk
	// Key: (string) (len=12) "whisk-apikey",
	// Key: (string) (len=15) "whisk-namespace",
	// Key: (string) (len=9) "io-apikey",
	// Key: (string) (len=11) "io-messages",
	// Key: (string) (len=7) "app-dir",
	// Key: (string) (len=19) "whisk-apihost-local",
	// Key: (string) (len=20) "whisk-apihost-docker",
	// (string) (len=20) "   \"key\": \"app-dir\",",
	// (string) (len=28) "   \"key\": \"whisk-namespace\",",
	// (string) (len=24) "   \"key\": \"io-messages\",",
	// (string) (len=22) "   \"key\": \"io-apikey\",",
	// (string) (len=33) "   \"key\": \"whisk-apihost-docker\",",
	// (string) (len=32) "   \"key\": \"whisk-apihost-local\",",
	// (string) (len=25) "   \"key\": \"whisk-apikey\",",

}
