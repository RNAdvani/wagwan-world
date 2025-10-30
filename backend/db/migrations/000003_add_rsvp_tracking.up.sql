-- Add RSVP tracking columns to guests table
ALTER TABLE guests ADD COLUMN rsvp_date TIMESTAMP;
ALTER TABLE guests ADD COLUMN plus_ones INTEGER NOT NULL DEFAULT 0;
ALTER TABLE guests ADD COLUMN dietary_restrictions TEXT;