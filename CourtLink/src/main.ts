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
      clientId: 'TqW28zdd6xCyQ9If2HU7nCo86rmvraC9',
      authorizationParams: {
        redirect_uri: window.location.origin, // Audience set to your Auth0 API Identifier
        scope: 'openid profile email'  // Add the required scopes here
      }
    }),
    provideRouter(routes, withComponentInputBinding()),
    provideHttpClient()
  ]
}).catch(err => console.error(err));

