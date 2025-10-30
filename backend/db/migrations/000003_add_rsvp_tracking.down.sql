-- Remove RSVP tracking columns from guests table
ALTER TABLE guests DROP COLUMN IF EXISTS dietary_restrictions;
ALTER TABLE guests DROP COLUMN IF EXISTS plus_ones;
ALTER TABLE guests DROP COLUMN IF EXISTS rsvp_date;