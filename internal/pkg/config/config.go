package config

import (
	"flag"
	"log"
)

func FromFlags(flags map[string]string) map[string]string {
	for flagName, flagInfo := range flags {
		flag.String(flagName, "", flagInfo)
	}
	flag.Parse()

	config := make(map[string]string, len(flags))
	flag.VisitAll(
		func(f *flag.Flag) {
			if f.Value.String() == "" {
				log.Fatalf("Missing argument: '-%s' (%s)", f.Name, f.Usage)
			}
			config[f.Name] = f.Value.String()
		},
	)
	return config
}
