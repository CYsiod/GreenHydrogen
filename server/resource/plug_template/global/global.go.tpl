package global

{{- if .HasGlobal }}

import "GreenHydrogen/server/plugin/{{ .Snake}}/config"

var GlobalConfig = new(config.{{ .PlugName}})
{{ end -}}