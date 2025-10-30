-- Remove notes column from guests table
ALTER TABLE guests DROP COLUMN IF EXISTS notes;