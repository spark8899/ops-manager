package global

{{- if .HasGlobal }}

import "github.com/spark8899/ops-manager/server/plugin/{{ .Snake}}/config"

var GlobalConfig = new(config.{{ .PlugName}})
{{ end -}}