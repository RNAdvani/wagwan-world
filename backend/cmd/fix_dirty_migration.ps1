# Check migration status in Docker Postgres and fix if dirty

$Result = docker exec -i eventguests-db psql -U postgres -d eventguests -t -A -c "SELECT version, dirty FROM schema_migrations;"
$Parts = $Result -split '\|'
$Version = $Parts[0].Trim()
$Dirty = $Parts[1].Trim()

Write-Host "Current Version: $Version"
Write-Host "Dirty: $Dirty"

if ($Dirty -eq "t") {
    $LastGood = [int]$Version - 1
    Write-Host "Database is dirty. Forcing migration to version $LastGood..."
    migrate -path db/migrations -database "postgres://postgres:postgres@localhost:55432/eventguests?sslmode=disable" force $LastGood
    Write-Host "Re-running all migrations..."
    migrate -path db/migrations -database "postgres://postgres:postgres@localhost:55432/eventguests?sslmode=disable" up
} else {
    Write-Host "Database is clean. No action required."
}
