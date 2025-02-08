import { routes } from './app/app.routes';
import { bootstrapApplication } from '@angular/platform-browser';
import { provideRouter, withComponentInputBinding } from '@angular/router';
import { provideAuth0 } from '@auth0/auth0-angular';
import { AppComponent } from './app/app.component';
import { RouterModule } from '@angular/router'; // Correct path if app.component.ts is inside the 'app' folder


bootstrapApplication(AppComponent, {
  providers: [
    provideAuth0({
      domain: 'dev-7gppji8v3bdbsj6k.us.auth0.com',
      clientId: 'K3yZGflpa3qLYWXQtrBUsNaO4xXrfwtv',
      authorizationParams: {
        redirect_uri: window.location.origin
      }
    }),
    provideRouter(routes, withComponentInputBinding()),
  ]

});