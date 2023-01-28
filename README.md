# goapidb
GO, Check API Endpoint 

This script sends a GET request to the specified API endpoint URL, check the response status code, if it's 200 OK it reads the response body, then it stores the data in the database, if there's an error in any step it logs the error and waits for 5 minutes before trying again. The logError function is used to write the error message to a log file named "api_failures.log" and the error message will be appended to the existing log file if it exist otherwise it will create a new one.
