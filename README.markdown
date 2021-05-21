ztpl contains functiosn to deal with Go's text/template and html/template

You have to call `ztpl.Init()` first with the path to load templates from, and
an optional map for templates compiled in the binary.

You can use `ztpl.Reload()` to reload the templates from disk on changes, which
is useful for development. e.g. with github.com/teamwork/reload:

```go
ztpl.Init("tpl", pack.Templates)

go func() {
    err := reload.Do(zlog.Module("main").Debugf, reload.Dir("./tpl", ztpl.Reload))
    if err != nil {
        panic(errors.Errorf("reload.Do: %v", err))
    }
}()
```
