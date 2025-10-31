#!/bin/bash
export AWS_ACCESS_KEY_ID="$(rage -d -i $HOME/age-id-yubikey $HOME/.local/share/gopass/stores/root/pop-os/S3_ACCOUNT_ID.age)"
export AWS_SECRET_ACCESS_KEY="$(rage -d -i $HOME/age-id-yubikey $HOME/.local/share/gopass/stores/root/pop-os/S3_ACCOUNT_KEY.age)"
export RESTIC_REPOSITORY="$(rage -d -i $HOME/age-id-yubikey $HOME/.local/share/gopass/stores/root/pop-os/RESTIC_REPOSITORY.age)"
export RESTIC_PASSWORD_COMMAND="rage -d -i $HOME/age-id-yubikey $HOME/.local/share/gopass/stores/root/pop-os/RESTIC_PASSWORD.age"
