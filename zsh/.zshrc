# Lines configured by zsh-newuser-install
HISTFILE=~/.histfile
HISTSIZE=1000
SAVEHIST=1000
setopt autocd extendedglob nomatch notify
unsetopt beep
# End of lines configured by zsh-newuser-install
# The following lines were added by compinstall
zstyle :compinstall filename '/home/username/.zshrc'

autoload -Uz compinit
compinit
# End of lines added by compinstall

PROMPT='%F{yellow}%n%f@%F{green}%M%f:%F{red}%0~%f(%F{14}%?%f)%# '


