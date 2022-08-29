#!/usr/bin/env sh

PATH=$PATH:/usr/local/bin
set -e

env

auth_usr=$(head -1 $1)
auth_passwd=$(tail -1 $1)

if [ $common_name = $username ]; then
  openvpn-user auth --db.path /etc/openvpn/easyrsa/pki/openvpn.db --user ${auth_usr} --password ${auth_passwd}
else
  echo "Authorization failed"
  exit 1
fi
