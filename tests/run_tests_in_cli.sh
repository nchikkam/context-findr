#!/bin/sh

npm install
npx cypress run --spec cypress/e2e/rest/home_page.cy.js cypress/e2e/rest/signup.cy.js cypress/e2e/rest/signin.cy.js cypress/e2e/rest/upload.cy.js cypress/e2e/rest/uploads.cy.js cypress/e2e/rest/search.cy.js

