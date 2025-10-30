ALTER TABLE events ADD COLUMN show BOOLEAN NOT NULL DEFAULT TRUE;

-- Set the 'show' column to FALSE for the event made during initial migration
UPDATE events
SET show = FALSE
WHERE id  = (SELECT MIN(id) FROM events);

INSERT INTO events (title, description, event_date, location, created_at, show)
VALUES
('Tech Innovators Summit 2025',
 'A high-energy conference bringing together startups, investors, and engineers shaping the future of AI and cloud computing.',
 '2025-11-20 10:00:00',
 'Bangalore International Exhibition Centre, Bengaluru',
 NOW(),
 true),

('Product Design Workshop',
 'Hands-on design sprint focusing on UX/UI principles and rapid prototyping with Figma and TailwindCSS.',
 '2025-12-05 09:30:00',
 '91Springboard, Koramangala, Bengaluru',
 NOW(),
 true),

('Remote Team Offsite 2026',
 'A 3-day retreat for distributed teams to collaborate, strategize, and unwind together.',
 '2026-01-18 08:00:00',
 'The Leela, Goa',
 NOW(),
 true);