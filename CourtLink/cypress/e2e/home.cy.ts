// cypress/e2e/full_skip_auth.cy.js
/// <reference types="cypress" />

const SPORTS = ['Basketball', 'Soccer', 'Tennis', 'Badminton', 'Cricket'];

describe('CourtLink Slideshow & Booking Flow (skipping Auth0)', () => {
  beforeEach(() => {
    // 1) Explicitly intercept the absolute URL
    cy.intercept('POST', 'http://localhost:8080/getCourts', {
      statusCode: 200,
      body: [
        { id: 1, name: 'Court A' },
        { id: 2, name: 'Court B' }
      ]
    }).as('getCourts');
  });

  it('1) shows rotating slideshow on the login screen', () => {
    cy.visit('/');
    cy.contains('Login to Continue', { timeout: 10000 }).should('be.visible');

    cy.get('img[alt="Slideshow Image"]', { timeout: 10000 })
      .should('be.visible')
      .invoke('attr', 'src')
      .then(src1 => {
        cy.wait(3500); // slide interval = 3000ms
        cy.get('img[alt="Slideshow Image"]')
          .invoke('attr', 'src')
          .should(src2 => {
            expect(src2).not.eq(src1);
          });
      });
  });

  it('2) skips login, shows sports grid, selects a sport and navigates', () => {
    // 2a) Skip Auth0 via query‑param
    cy.visit('/?skipAuth=true');

    // 2b) Login gate gone
    cy.contains('Login to Continue').should('not.exist');

    // 2c) Wait for the grid to render
    cy.get('.cursor-pointer', { timeout: 10000 }).should('have.length', SPORTS.length);

    // 2d) Each sport visible by name
    cy.wrap(SPORTS).each(name => {
      cy.contains(name).should('be.visible');
    });

    // 2e) Click “Tennis”
    cy.contains('Tennis').click();

    // 2f) Wait for navigation (instead of XHR)
    cy.url({ timeout: 10000 }).should('include', '/courts/tennis');

    // 2g) Verify your stubbed courts render
    cy.contains('Court 1').should('be.visible');
    cy.contains('Court 2').should('be.visible');
  });
});
