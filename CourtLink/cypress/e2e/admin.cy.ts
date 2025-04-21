/// <reference types="cypress" />

describe('Admin Portal Flow', () => {
  it('logs in and navigates through admin dashboard, adds new sport, adds new court, resets the newly added court and then deletes it', () => {
    // Go to login page
    cy.log('Visiting the admin login page');
    cy.visit('http://localhost:4200/admin');

    // Login
    cy.log('Filling in login credentials');
    cy.get('input[placeholder="Username"]').type('Admin123');
    cy.get('input[placeholder="Password"]').type('Password123');
    cy.contains('Log In').click();

    // Verify successful login
    cy.log('Verifying admin dashboard is loaded');
    cy.url().should('include', '/admin');

    // Navigate to sports
    cy.log('Navigating to the Sports section');
    cy.contains('View Sports').click();
    cy.url().should('include', '/admin/sports');

    // Add a new sport
    cy.log('Adding a new sport: Pickleball');
    cy.contains('Add New Sport').click();
    cy.get('input').type('Pickleball');
    cy.contains('Confirm').click();

    // Navigate to courts
    cy.log('Navigating to the Courts section');
    cy.visit('http://localhost:4200/admin/courts');

    // Add a new court
    cy.log('Adding a new court: test court for basketball');
    cy.contains('Add New Court').click();
    cy.get('input[placeholder="Court Name"]').type('test court');
    cy.get('input[placeholder="Sport Name"]').type('basketball');
    cy.contains('Confirm').click();

    // Reset court
    cy.log('Resetting the newly added court');
    cy.get('table').contains('test court').parent().contains('Reset').click();

    // Delete court
    cy.log('Deleting the newly added court');
    cy.get('table').contains('test court').parent().contains('Delete').click();
  });
});
