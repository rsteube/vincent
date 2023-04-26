# vincent

Terminal color theme chooser.

[![asciicast](https://asciinema.org/a/580784.svg)](https://asciinema.org/a/580784)


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
