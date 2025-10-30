-- Remove foreign key constraint
ALTER TABLE guests DROP CONSTRAINT IF EXISTS fk_guests_event_id;

-- Remove event_id column from guests
ALTER TABLE guests DROP COLUMN IF EXISTS event_id;

-- Drop events table
DROP TABLE IF EXISTS events;