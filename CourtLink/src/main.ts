import { bootstrapApplication } from '@angular/platform-browser';
import { provideAuth0 } from '@auth0/auth0-angular';
import { AppComponent } from './app/app.component'; // Correct path if app.component.ts is inside the 'app' folder

bootstrapApplication(AppComponent, {
  providers: [
    provideAuth0({
      domain: 'dev-zvcqb0yxk33pst6d.us.auth0.com',
      clientId: 'RW4WlJ5sD2pWNV0a4AdrNjKf9LJxWuqg',
      authorizationParams: {
        redirect_uri: window.location.origin
      }
    }),
  ]
});