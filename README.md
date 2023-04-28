# vincent

[![PkgGoDev](https://pkg.go.dev/badge/github.com/rsteube/vincents)](https://pkg.go.dev/github.com/rsteube/vincent)
[![GoReportCard](https://goreportcard.com/badge/github.com/rsteube/vincent)](https://goreportcard.com/report/github.com/rsteube/vincent)
[![Coverage Status](https://coveralls.io/repos/github/rsteube/vincent/badge.svg?branch=master)](https://coveralls.io/github/rsteube/vincent?branch=master)
[![Packaging status](https://repology.org/badge/tiny-repos/vincent.svg)](https://repology.org/project/vincent/versions)

Terminal color scheme chooser ([source](https://gogh-co.github.io/Gogh/)).

[![asciicast](https://asciinema.org/a/581022.svg)](https://asciinema.org/a/581022)


# Getting Started

```sh
# bash (~/.bashrc)
source <(vincent _carapace)

# elvish (~/.elvish/rc.elv)
eval (vincent _carapace|slurp)

# fish (~/.config/fish/config.fish)
vincent _carapace | source

# nushell (~/.config/nushell/config.nu)
vincent _carapace nushell # update config.nu manually according to output

# oil (~/.config/oil/oshrc)
source <(vincent _carapace)

# powershell (~/.config/powershell/Microsoft.PowerShell_profile.ps1)
Set-PSReadLineOption -Colors @{ "Selection" = "`e[7m" }
Set-PSReadlineKeyHandler -Key Tab -Function MenuComplete
vincent _carapace | Out-String | Invoke-Expression

# tcsh (~/.tcshrc)
set autolist
eval `vincent _carapace`

# xonsh (~/.config/xonsh/rc.xsh)
COMPLETIONS_CONFIRM=True
exec($(vincent _carapace))

# zsh (~/.zshrc)
source <(vincent _carapace)
```
