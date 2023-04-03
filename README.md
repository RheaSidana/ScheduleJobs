<h1>Schedule Jobs</h1>

<h3> Steps to run the application </h3>
<h4>  1. Clone the repo </h4>
<h4>  2. run the command: go mod tidy </h4>
<h4>  3. DB operations </h4>
<h4> &nbsp;&nbsp;&nbsp; (MAC)</h4>
<h4> &nbsp;&nbsp;&nbsp;&nbsp; 1. install postgres: brew install postgresql</h4>
<h4> &nbsp;&nbsp;&nbsp;&nbsp; 2. start/stop postgres service: brew services start/stop postgresql</h4>
<h4> &nbsp;&nbsp;&nbsp;&nbsp; 3. psql postgres</h4>
<h4> &nbsp;&nbsp;&nbsp;&nbsp; 4. add postgres password: \password {password};</h4><br/>
<h4> &nbsp;&nbsp;&nbsp; (Windows)</h4>
<h4> &nbsp;&nbsp;&nbsp;&nbsp; 1. install postgres: [Link text Here](https://www.postgresql.org/download/windows/)</h4>
<h4> &nbsp;&nbsp;&nbsp;&nbsp; 2. port: 5432 (defaut), user: postgres (defaut)</h4>
<h4> &nbsp;&nbsp;&nbsp;&nbsp; 3. add postgres password</h4>
<h4> &nbsp;&nbsp;&nbsp;&nbsp; 4. open psql sql shell</h4>
<h4> &nbsp;&nbsp;&nbsp; 5. create orm db: CREATE database job_scheduling;</h4>
<h4> &nbsp;&nbsp;&nbsp; 6. connect to db: \c "db"</h4>
<h4> &nbsp;&nbsp;&nbsp; 7. Edit .env file with postgres details </h4><br/>
<h4>  4. Migrate Tables (run command: ### go run .\migrations\migrate.go) </h4>
<h4> &nbsp;&nbsp;&nbsp; View DB Table Schemas: ### \d "tablename" </h4>
<h4>  5. Seed Data to the Table (run command: ### go run .\dataSeeding\dataSeed.go)</h4>
<h4>  6. Run the application (run command: ### go run .)</h4>
<h4>  7. Call APIs </h4>
<h4> &nbsp;&nbsp;&nbsp; Add slots to db <br/>
    Postman: POST: http://localhost:3033/slot  <br/>
    JSON Body: { <br/>
            "user_id": 1, <br/>
            "job_id": 3, <br/>
            "start_time": "2023-03-18T08:00:00.000Z" <br/>
        } <br/>
</h4>
<h4> &nbsp;&nbsp;&nbsp; Get Current Week slots: <br/>
    Postman: GET: http://localhost:3033/slot  <br/>
    JSON Body: { <br/>
            "user_id": 1 <br/>
        } <br/>
</h4>
