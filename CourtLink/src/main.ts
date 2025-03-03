import { routes } from './app/app.routes';
import { bootstrapApplication } from '@angular/platform-browser';
import { provideRouter, withComponentInputBinding } from '@angular/router';
import { provideAuth0 } from '@auth0/auth0-angular';
import { AppComponent } from './app/app.component';
import { RouterModule } from '@angular/router'; // Correct path if app.component.ts is inside the 'app' folder
import { provideHttpClient } from '@angular/common/http';

bootstrapApplication(AppComponent, {
  providers: [
    provideAuth0({
      domain: 'dev-7gppji8v3bdbsj6k.us.auth0.com',
      clientId: 'K3yZGflpa3qLYWXQtrBUsNaO4xXrfwtv',
      authorizationParams: {
        redirect_uri: window.location.origin,
        audience: 'https://dev-7gppji8v3bdbsj6k.us.auth0.com/api/v2/', // Audience set to your Auth0 API Identifier
        scope: 'read:users write:users'  // Add the required scopes here
      }
    }),
    provideRouter(routes, withComponentInputBinding()),
    provideHttpClient()
  ]
}).catch(err => console.error(err));
