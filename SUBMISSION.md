# Submission

**Candidate Name**: Rohan Advani  
**Date**: 30th June 2025
**Email**: advanirohan03@gmail.com

---

## Time Spent
Approximately 18 hours

---

## What I Built

### Main Feature: Public RSVP Page

Describe what you implemented:
- Location: `/rsvp` route
- Key features: 
  - Database migration to add events table, event_id in the guests table, and notes, plus_ones, dietary_restrictions fields.
  - Made sure to handle the case where the existing users are associated with a default "Legacy Event" which is a placeholder for existing data.
  - ```cmd``` folder in the ```backend``` contains migration scripts to implement the database changes using the `golang-migrate` tool.
  - It also contains scripts to check whether the database has become dirty and if so it tries to return to a clean state by rolling back to the last successful migration.
  - ```migrations``` folder contains the actual migration files, it contains both ```.up.sql``` and ```.down.sql``` files for each migration step.
  - Frontend Svelte page to display a list of events fetched from the backend.
  - User can select an event and fill out the RSVP form with their details, including name, email, phone number, RSVP status, plus ones, dietary restrictions, and notes.


- Design choices: 
    - All the events are fetched from the DB and displayed. Users can select the event they want to RSVP for. They provide their details such as name, email, phone number, and RSVP status, plus ones, dietary restrictions, and notes. The form validates inputs and provides feedback on submission.
    - As I had to make a single page for RSVP, I had to fetch all the events rather than just one event using params. This allows users to RSVP for any event listed.
    - While submitting the form, I ensured that all fields are validated properly. If any field is invalid, an error message is shown to the user.
---

## Bugs Found & Fixed

List any bugs you discovered in the existing code and how you fixed them:

### Bug 1: GetAllGuests query sorts by created_at instead of name
- **Problem**: The query used to fetch all guests was sorting the results by the `created_at` timestamp instead of the guest's name.
- **Location**: `db/queries.go` (line 23)
- **Solution**: Modified the SQL query to include an `ORDER BY name` clause.
- **Why**: This change ensures that the guests are listed in alphabetical order by name, which is more user-friendly.

### Bug 2: Missing email validation
- **Problem**: The RSVP form did not validate the email format, allowing invalid emails to be submitted.
- **Location**: `frontend/src/routes/rsvp/+page.svelte` (line 34)
- **Solution**: Added a regex email validation check before form submission (line 46).
- **Why**: Validating email format to ensure that users provide a correctly formatted email address.

### Bug 3: Expects 'status' but frontend sends 'filter'
- **Problem**: The RSVP form was sending a 'filter' parameter instead of the expected 'status' parameter.
- **Location**: `frontend/src/lib/api.ts` (line 35) and `backend/handlers/guests.go` (line 40)
- **Solution**: Updated the frontend to send 'status' instead of 'filter' and selected the correct parameter in the backend handler.
- **Why**: This makes sure that we fetch the the guests based on their RSVP status correctly.

### Bug 4: Docker and PostgreSQL port conflict
- **Problem**: Both Docker and PostgreSQL were configured to use the same port (5432), causing a conflict.
- **Location**: `docker-compose.yml` (line 12)
- **Solution**: Changed the host port mapping in `docker-compose.yml` to 55432 from 5432 and updated the instructions to set `DB_PORT` accordingly when running the backend locally.
- **Why**: This resolves the port conflict, allowing both Docker and PostgreSQL to run simultaneously without issues.

---

## Challenges & Solutions

Describe the biggest challenges you faced and how you overcame them:

### Challenge 1
**Problem**: Implementing the database migration system to handle schema changes without data loss.
**Solution**: I used the `golang-migrate` tool to create migration scripts that can be applied to the database. This allows for version controlled schema changes and the ability to roll back changes if needed.

### Challenge 2
**Problem**: Ensuring data integrity while migrating existing users to the new event system.
**Solution**: I created a default "Legacy Event" for existing users who did not have an associated event. This way, their data is preserved and can be migrated to the new structure without loss.

---

## How to Test My Work

Step-by-step instructions for testing your implementation:

### Setup
1. # Install dependencies
   ```bash
   cd backend
   go mod download
   cd ../frontend
   npm install
   ```
2. # Start Docker containers
   ```bash
   docker-compose up -d
   ```
3. # Run database migrations
   ```bash
   cd backend
   go run cmd/main.go up
   ```
4. # Start backend server
   ```bash
   go run main.go
   ```
5. # Start frontend server
   ```bash
   cd ../frontend
   npm run dev
   ```

### Testing the RSVP Page
1. The RSVP Page
   - Navigate to `http://localhost:5173/rsvp`
   - Or click on the top right "RSVP" link from the home page.
2. Event selection and form filling
   - (Desktop) Select an event from the list shown on the left side".
   - Fill out the RSVP form with valid details (name, email, phone number, status, plus ones, dietary restrictions, notes) and click "Submit RSVP".
3. Expected Behavior
   - A success message should appear confirming the RSVP submission.
   - The backend should store the RSVP details in the database associated with the selected event.
   - You can check the homepage for the guest to appear in the guest list.

### Testing Bug Fixes
1. Guest list status filtering
   - Navigate to the homepage.
   - Use the status filter dropdown to select different RSVP statuses (e.g., "Attending", "Declined" or "Pending").
   - Verify that the guest list updates correctly based on the selected status.

2. Guest list sorting
   - Verify that the guest list is sorted alphabetically by name.

3. Email validation
   - On the RSVP page, try submitting the form with an invalid email format.
   - Verify that an error message is displayed and the form is not submitted.

4. RSVP guest creation
   - Submit a valid RSVP form and check the database to ensure the guest is created with the correct details, including notes, plus_ones, and dietary_restrictions.
   - Guest is shown on the homepage ensuring backward compatibility.
---

## What I'd Improve With More Time

If you had additional time, what would you add or improve?

1. Add an event page to show all the events and then create a dynamic RSVP page with params to fill out the RSVP for a specific event.
2. Implement duplicate RSVP prevention by checking if a guest with the same email has already RSVP'd for the selected event.
3. Improve the UI/UX of the RSVP form by adding client-side validation and better error handling.
4. Send email confirmations to guests upon successful RSVP submission.
5. Add RSVP update functionality to allow guests to modify their RSVP details before a certain duration of the event.
6. Implement attendee count to show how many people are attending each event.

---

## Additional Notes

Any other comments or observations about the project:

- The existing codebase is well structured, making it easier to navigate and implement new features.
- I loved the choices of the tech stack, it made me think about both frontend and backend aspects of web development.
- Talking about further challenges:
    - To implement zero-downtime migrations, it is better to have additive changes first then modify existing columns. This ensures that the application can continue to function while migrations are being applied.
    - Migration should be gradual as much as possible to avoid long downtime periods (if any).
    - As this is a local setup, I won't be able to perfrom blue-green deployments but I had a thought of simulation this what we can do is make a separate environment of the new migrations on a different host, and then switch the traffic to the new environment once everything is verified to be working fine.
    - For schema validation, we can have some integration tests that check the schema against expected structure.
    - For larger tables, we can perform migrations during off-peak hours and gradually do the migrations in batches to avoid locking the entire table for a long time.
---

## Self-Assessment

How confident are you in:
- Code quality: [ ] Low [ ] Medium [x] High
- Functionality: [ ] Low [ ] Medium [x] High
- Design: [ ] Low [ ] Medium [x] High
- Bug fixes: [ ] Low [x] Medium [ ] High

What are you most proud of in this submission?

I am most proud of successfully implementing the database migration as per requirements while ensuring data integrity for existing users. The migration scripts are well-structured and can be easily applied or rolled back as needed. Additionally, the RSVP page is user friendly and provides a seamless experience for guests to register for events.

---

Thank you for reviewing my work!
