package otel_rules
import service39054 "otelbuild-demo/service"
import service95044 "otelbuild-demo/otel_rules/service95044"
import http55767 "net/http"
import http87548 "otelbuild-demo/otel_rules/http87548"
import otel_log "log"
import otel_runtime_debug "runtime/debug"
func init() {
	service39054.CreateUserEnterHookImpl = service95044.CreateUserEnterHook
	service39054.CreateUserExitHookImpl = service95044.CreateUserExitHook
	service39054.OtelGetStackImpl = otel_runtime_debug.Stack
	service39054.OtelPrintStackImpl = func(bt []byte){ otel_log.Printf(string(bt)) }
	http55767.HttpClientEnterHookImpl = http87548.HttpClientEnterHook
	http55767.HttpClientExitHookImpl = http87548.HttpClientExitHook
	http55767.OtelGetStackImpl = otel_runtime_debug.Stack
	http55767.OtelPrintStackImpl = func(bt []byte){ otel_log.Printf(string(bt)) }
}
