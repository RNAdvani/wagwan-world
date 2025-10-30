# Event Guest Manager 

Event Guest Manager is a web application designed to help event organizers manage their guest lists efficiently. With features like RSVP tracking, guest information management, and event reminders, this tool simplifies the process of organizing events.

For setup instructions, please refer to the [SETUP.md](SETUP.md) file.

##  Key features: 
  - Database migration to add events table, event_id in the guests table, and notes, plus_ones, dietary_restrictions fields.
  - Made sure to handle the case where the existing users are associated with a default "Legacy Event" which is a placeholder for existing data.
  - ```cmd``` folder in the ```backend``` contains migration scripts to implement the database changes using the `golang-migrate` tool.
  - It also contains scripts to check whether the database has become dirty and if so it tries to return to a clean state by rolling back to the last successful migration.
  - ```migrations``` folder contains the actual migration files, it contains both ```.up.sql``` and ```.down.sql``` files for each migration step.
  - Frontend Svelte page to display a list of events fetched from the backend.
  - User can select an event and fill out the RSVP form with their details, including name, email, phone number, RSVP status, plus ones, dietary restrictions, and notes.


Detailed information about the implementation and bugs fixed can be found in the [SUBMISSION.md](SUBMISSION.md) file.