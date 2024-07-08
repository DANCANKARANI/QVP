#!/bin/bash

until mysql -h "db" -u "root" -p"$DB_PASSWORD" -e "show databases;" > /dev/null 2>&1; do
  echo "Waiting for database connection..."
  sleep 5
done

echo "Database is up - executing command"
exec "$@"
