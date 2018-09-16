package conf

import (
	"time"

	"github.com/spf13/pflag"
	"github.com/spf13/viper"
)

type Flag func(flags *pflag.FlagSet)

func stringFlag(p *string, name, short, value, usage string) Flag {
	return func(flags *pflag.FlagSet) {
		flags.StringVarP(p, name, short, value, usage)
		viper.BindPFlag(name, flags.Lookup(name))
	}
}

func stringSliceFlag(p *[]string, name, usage string) Flag {
	return func(flags *pflag.FlagSet) {
		flags.StringArrayVar(p, name, nil, usage)
		viper.BindPFlag(name, flags.Lookup(name))
	}
}

func boolFlag(b *bool, name string, usage string) Flag {
	return func(flags *pflag.FlagSet) {
		flags.BoolVar(b, name, false, usage)
		viper.BindPFlag(name, flags.Lookup(name))
	}
}

func intFlag(p *int, name string, value int, usage string) Flag {
	return func(flags *pflag.FlagSet) {
		flags.IntVar(p, name, value, usage)
		viper.BindPFlag(name, flags.Lookup(name))
	}
}

func durationFlag(p *time.Duration, name string, value time.Duration, usage string) Flag {
	return func(flags *pflag.FlagSet) {
		flags.DurationVar(p, name, value, usage)
		viper.BindPFlag(name, flags.Lookup(name))
	}
}
