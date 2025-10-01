#!/bin/sh
set -e
chmod 600 /var/lib/postgresql/server.key
chown postgres:postgres /var/lib/postgresql/server.key
