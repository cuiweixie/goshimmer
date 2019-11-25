package cli

import (
	"fmt"
	"os"
	"path/filepath"
	"sort"
	"strings"

	flag "github.com/spf13/pflag"

	"github.com/iotaledger/goshimmer/packages/node"
)

var enabledPlugins []string
var disabledPlugins []string

func AddPluginStatus(name string, status int) {
	switch status {
	case node.Enabled:
		enabledPlugins = append(enabledPlugins, name)
	case node.Disabled:
		disabledPlugins = append(disabledPlugins, name)
	}
}

func getList(a []string) string {
	sort.Strings(a)
	return strings.Join(a, " ")
}

func printUsage() {
	fmt.Fprintf(
		os.Stderr,
		"\n"+
			"SHIMMER\n\n"+
			"  A lightweight modular IOTA node.\n\n"+
			"Usage:\n\n"+
			"  %s [OPTIONS]\n\n"+
			"Options:\n",
		filepath.Base(os.Args[0]),
	)
	flag.PrintDefaults()

	fmt.Fprintf(os.Stderr, "\nThe following plugins are enabled by default and can be disabled with -%s:\n  %s\n", node.CFG_DISABLE_PLUGINS, getList(enabledPlugins))
	fmt.Fprintf(os.Stderr, "The following plugins are disabled by default and can be enabled with -%s:\n  %s\n", node.CFG_ENABLE_PLGUINS, getList(disabledPlugins))
	fmt.Fprintf(os.Stderr, "The enabled/disabled plugins can be overriden by altering %s/%s inside config.json\n\n", node.CFG_ENABLE_PLGUINS, node.CFG_DISABLE_PLUGINS)
}
