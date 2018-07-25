# We don't need to source the rcfile here (like the other shells do)  because
# the mechanism we are using to spawn the subshell is already loading it.
#
# Also, it's ineffectual to attempt setting a prompt in this script since those
# `set` variables are not inherited when we spawn the sub shell via exec.  Only
# `setenv` values are inherited.

{{range $K, $V := .Env}}
setenv {{$K}} "{{$V}}"
{{end}}
cd "{{.WD}}"

{{.UserScripts}}
