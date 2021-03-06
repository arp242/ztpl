ztpl contains functiosn to deal with Go's text/template and html/template.

Import as `zgo.at/ztpl`; API docs: https://godocs.io/zgo.at/ztpl

What you can do with this:

- The `parse` subpackage is a copy/fork of `text/template/parse`; the default
  makes it impossible to write tooling, as it will [error out on undefined
  functions][f]. This changes `Parse()` to accept a `Mode`, and you can add
  `ParseRelaxFunctions` to not error out.

  Also adds some other methods to facilitate the writing of tooling for
  templates.

  [f]: https://github.com/golang/go/issues/22538

- You can set up templates with `ztpl.Init()`, which can then be reloaded from
  the filesystem with `ztpl.Reload()`, e.g. with github.com/teamwork/reload:

  ```go
  ztpl.Init("tpl", pack.Templates)
  
  go func() {
      err := reload.Do(zlog.Module("main").Debugf, reload.Dir("./tpl", ztpl.Reload))
      if err != nil {
          panic(errors.Errorf("reload.Do: %v", err))
      }
  }()
  ```

  Simple replacing a `templates` variable introduces race conditions, this takes
  care of that.

  This also automatically runs either `text/template` or `html/template`
  depending on the file extension (`.gotxt` or `.gohtml`).

- Trace template execution with `Trace()`/`TestTemplateExecution()`, as a kind
  of poor-man code coverage.

- Additional template functions in `tplfunc/`.

TODO:

- Proper coverage support, now that we can forked the parse package.

- Compile templates to Go code. Especially things like tight loops are
  surprisingly slow.

- A template format tool, like the (unfinished) https://github.com/gotpl/gtfmt

Alhough, maybe it makes more sense to use
https://github.com/valyala/quicktemplate or https://github.com/a-h/templ
