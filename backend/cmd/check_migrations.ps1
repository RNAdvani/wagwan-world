docker exec -i eventguests-db psql -U postgres -d eventguests -t -A -c "SELECT version, dirty FROM schema_migrations;"
