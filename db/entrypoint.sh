#!/bin/bash
set -e

# Выполняем стандартный entrypoint PostgreSQL
exec docker-entrypoint.sh postgres "$@"