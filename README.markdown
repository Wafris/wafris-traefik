# Wafris traefik plugin 

## Usage

### Define the plugin in Static Configuration

Wafris plugin must be first defined in your traefik [static configuration][static]

[static]: https://doc.traefik.io/traefik/getting-started/configuration-overview/#the-static-configuration


There are three different, mutually exclusive (i.e. you can use only one at the same time), ways to define static configuration options in Traefik:
	1	In a configuration file
	2	In the command-line arguments
	3	As environment variables

### Static Config: YAML or TOML example

YAML Static configuration example:

```YAML
# Define the module name for the wafris plugin
# we use wafrisPlugin in this example, but any valid module name works
experimental:
  plugins:
    wafrisPlugin:
      moduleName: github.com/Wafris/wafris-traefik
      version: v0.0.1
```

TOML Static configuration example:

```TOML
# Define the module name for the wafris plugin
# we use wafrisPlugin in this example, but any valid module name works

experimental:
  plugins:
    wafrisPlugin:
      moduleName: github.com/Wafris/wafris-traefik
      version: v0.0.1
```

### Static Config: CLI example

In this example, we use the name wafrisPlugin.  Any valid module name should work.

```
--experimental.plugins.wafrisPlugin.modulename=github.com/Wafris/wafris-traefik --experimental.plugins.wafrisPlugin.version=v0.0.1
```

### Static Config: environment variables example

TODO


### Add the plugin to a router

TODO

<img src='https://uptimer.expeditedsecurity.com/wafris-traefik' width='0' height='0'>
