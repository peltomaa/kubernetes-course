#!/usr/bin/env bash
set -e

if [ $URL ]
then
  pg_dump -v $URL > /usr/src/app/backup.sql
  gcloud storage cp /usr/src/app/backup.sql  "gs://dwk-backups/crud-backup-$(date +'%Y-%m-%d')"
fi
