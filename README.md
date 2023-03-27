# PocketHealth Intern Challenge

## Setup
For this assignment, the backend is written in Golang. Please go [here](https://go.dev/), and follow the appropriate instructions to set up the environment on your machine.

The frontend is written in Angular. You can find instructions to set up your environment [here](https://angular.io/guide/setup-local).

Please unzip the starter code attached. You will find the source code for the server in "/backend", and the web app in "/frontend".

Please go ahead and build and run the server and web app, there are instructions below. The server is configured to run on your localhost, at port 8080, and the web app will run on port 4200. Once both the server and app are running, using any browser, you may wish to visit http://localhost:4200

## Run Project
To run the full project, you'll need to run the backend and frontend.

### Backend
Install dependencies using `go get ./...` and then build and run the project.

### Frontend
Install dependencies using `npm install` and then serve the Angular project.

## Changes
1.  We're considering allowing anyone to register for access to PocketHealth. This task will give you a flavour of the approach.

2. We'd like to modify the server to run on port 80. Please make the appropriate modification.

3. There is a bug where the POST "/register" isn’t returning the User ID. Instead it returns an empty string. Find and fix this bug.

4. The registration page attempts to redirect the user to "/home" but it fails since this page isn’t implemented yet. Please add a page to the web app on the "/home" path that displays the message: "Welcome to PocketHealth {Name}. Your User ID is: {User ID}"

You can use the "/register" page as a template, please keep the PocketHealth logo on the page.

5. After some time, we decide that we also want to capture the user's favourite colour. Please extend the form to collect the user's favourite colour and update the backend to store this with the user's other details.

6. Modify the home page to display the user's favourite colour.

7. You may wish to perform some basic input sanitization and form validation in the front end. If you wish, go ahead and do so.

## Submission
Please provide a link to your solution (Google Drive, Dropbox, GitHub, etc), and email it back  within the next 24 hours. 
