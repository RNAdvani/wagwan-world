-- Create events table

CREATE TABLE events (
    id SERIAL PRIMARY KEY,
    title VARCHAR(255) NOT NULL,
    description TEXT,
    event_date TIMESTAMP NOT NULL,
    location VARCHAR(255),
    created_at TIMESTAMP NOT NULL DEFAULT NOW()
);


-- Insert default event for existing guests
INSERT INTO events (title, description, event_date, location, created_at)
VALUES ('Legacy Event', 'Default event for existing guests', NOW(), 'TBD', NOW());

-- Add event_id column to guests table
ALTER TABLE guests ADD COLUMN event_id INTEGER;

-- Set event_id for existing guests to the default event
UPDATE guests SET event_id = (SELECT id FROM events WHERE title = 'Legacy Event' LIMIT 1);

ALTER TABLE guests
ALTER COLUMN event_id SET NOT NULL;

ALTER TABLE guests
ADD CONSTRAINT fk_event
FOREIGN KEY (event_id)
REFERENCES events(id)
ON DELETE CASCADE;

-- Create index for faster queries 
CREATE INDEX IF NOT EXISTS idx_guests_event_id ON guests(event_id);
CREATE INDEX IF NOT EXISTS idx_events_event_date ON events(event_date);
CREATE INDEX IF NOT EXISTS idx_events_title ON events(title);