# tree-sitter-coconut

A [tree-sitter][] grammar for [Coconut][].

[tree-sitter]: https://github.com/tree-sitter/tree-sitter
[Coconut]: https://coconut-lang.org/

This is very much a work in progress.


## Coconut Features

<https://coconut.readthedocs.io/en/latest/DOCS.html>

- Operators
  - [x] lambda syntax
- Keywords
- Expressions
- Function Definition
- Statements
- Built-Ins



## Development

- <https://tree-sitter.github.io/tree-sitter/creating-parsers>

`cargo install tree-sitter-cli`

### helix

You can see how the grammar works for syntax highlighting by trying it out in the [helix editor](https://helix-editor.com/).

Add the following to your [language.toml file](https://docs.helix-editor.com/languages.html):

```toml
[[language]]
name = "coconut"
scope = "source.coconut"
injection-regex = "coconut"
file-types = ["coco"]
shebangs = ["coconut"]
roots = ["pyproject.toml", "setup.py", "poetry.lock", "pyrightconfig.json"]
comment-token = "#"
# language-servers = [ "pylsp" ]
# TODO: pyls needs utf-8 offsets
indent = { tab-width = 4, unit = "    " }

[[grammar]]
name = "coconut"
# source = { git = "https://github.com/tree-sitter/tree-sitter-python", rev = "4bfdd9033a2225cc95032ce77066b7aeca9e2efc" }
source = { path = "/PATH/TO/THIS/REPO"}
```

Helix uses syntax highlighting and indentation queries from it's own runtime directory, not from the tree-sitter package (because Helix uses some of its own names in the queries). So, for now we can do something like this:

```bash
$ ln -s $PWD/queries/helix ~/.config/helix/runtime/queries/coconut
```

Later, the `queries/helix` folder should probably become a pull request into helix's source.

```bash
$ hx --grammar fetch
$ hx --grammar build
$ hx --health coconut
```

Now try opening a coconut file in helix, and you should have syntax highlighting!



<!--
*Ideally we'll have all these things, but for now these are just links for tree-sitter-python*

[![CI][ci]](https://github.com/tree-sitter/tree-sitter-python/actions/workflows/ci.yml)
[![discord][discord]](https://discord.gg/w7nTvsVJhm)
[![matrix][matrix]](https://matrix.to/#/#tree-sitter-chat:matrix.org)
[![crates][crates]](https://crates.io/crates/tree-sitter-python)
[![npm][npm]](https://www.npmjs.com/package/tree-sitter-python)
[![pypi][pypi]](https://pypi.org/project/tree-sitter-python/)


## References

- [Python 2 Grammar](https://docs.python.org/2/reference/grammar.html)
- [Python 3 Grammar](https://docs.python.org/3/reference/grammar.html)

[ci]: https://img.shields.io/github/actions/workflow/status/tree-sitter/tree-sitter-python/ci.yml?logo=github&label=CI
[discord]: https://img.shields.io/discord/1063097320771698699?logo=discord&label=discord
[matrix]: https://img.shields.io/matrix/tree-sitter-chat%3Amatrix.org?logo=matrix&label=matrix
[npm]: https://img.shields.io/npm/v/tree-sitter-python?logo=npm
[crates]: https://img.shields.io/crates/v/tree-sitter-python?logo=rust
[pypi]: https://img.shields.io/pypi/v/tree-sitter-python?logo=pypi&logoColor=ffd242

-->
