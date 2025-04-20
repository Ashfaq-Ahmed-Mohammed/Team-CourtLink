/// <reference types="cypress" />// cypress/support/commands.js
Cypress.Commands.add('loginByAuth0', () => {
    const auth0Domain   = Cypress.env('auth0_domain');
    const clientId      = Cypress.env('auth0_client_id');
    const clientSecret  = Cypress.env('auth0_client_secret');
    const audience      = Cypress.env('auth0_audience');
    const username      = Cypress.env('auth0_username');
    const password      = Cypress.env('auth0_password');
  
    cy.request({
      method: 'POST',
      url: `https://${auth0Domain}/oauth/token`,
      body: {
        grant_type:    'password',
        username,
        password,
        audience,
        scope:         'openid profile email',
        client_id:     clientId,
        client_secret: clientSecret
      }
    }).then(({ body }) => {
      const { access_token, id_token, expires_in } = body;
      const cacheKey = `@@auth0spajs@@::${auth0Domain}::${clientId}::openid profile email`;
      const item = {
        body: {
          client_id:    clientId,
          access_token,
          id_token,
          scope:        'openid profile email',
          expires_in
        },
        expiresAt: Date.now() + expires_in * 1000
      };
      // seed the localStorage exactly as the SDK does
      window.localStorage.setItem(cacheKey, JSON.stringify(item));
    });
  });
  
// ***********************************************
// This example commands.ts shows you how to
// create various custom commands and overwrite
// existing commands.
//
// For more comprehensive examples of custom
// commands please read more here:
// https://on.cypress.io/custom-commands
// ***********************************************
//
//
// -- This is a parent command --
// Cypress.Commands.add('login', (email, password) => { ... })
//
//
// -- This is a child command --
// Cypress.Commands.add('drag', { prevSubject: 'element'}, (subject, options) => { ... })
//
//
// -- This is a dual command --
// Cypress.Commands.add('dismiss', { prevSubject: 'optional'}, (subject, options) => { ... })
//
//
// -- This will overwrite an existing command --
// Cypress.Commands.overwrite('visit', (originalFn, url, options) => { ... })
//
// declare global {
//   namespace Cypress {
//     interface Chainable {
//       login(email: string, password: string): Chainable<void>
//       drag(subject: string, options?: Partial<TypeOptions>): Chainable<Element>
//       dismiss(subject: string, options?: Partial<TypeOptions>): Chainable<Element>
//       visit(originalFn: CommandOriginalFn, url: string, options: Partial<VisitOptions>): Chainable<Element>
//     }
//   }
// }