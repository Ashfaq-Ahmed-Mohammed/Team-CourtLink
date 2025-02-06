// import { MatIconModule } from '@angular/material/icon';
// import { MatMenuModule } from '@angular/material/menu';
// import { Component } from '@angular/core';
// import { AuthService } from '@auth0/auth0-angular';  // Import AuthService
// import { AsyncPipe, NgIf } from '@angular/common';
// @Component({
//   selector: 'app-navbar',
//   imports: [MatMenuModule, MatIconModule, NgIf, AsyncPipe], // Keep the imports as they were
//   templateUrl: './navbar.component.html',
//   styleUrls: ['./navbar.component.css']
// })

// export class NavbarComponent {
//   constructor(public auth: AuthService) {}

//   login() {
//     this.auth.loginWithRedirect();
//   }
  
//   logout() {
//     this.auth.logout({ logoutParams: { federated: true } });
//   }
// }

import { MatIconModule } from '@angular/material/icon';
import { MatMenuModule } from '@angular/material/menu';
import { Component } from '@angular/core';
import { AuthService } from '@auth0/auth0-angular';  // Import AuthService
import { AsyncPipe, NgIf } from '@angular/common';
import { DomSanitizer, SafeUrl } from '@angular/platform-browser';  // Import DomSanitizer

@Component({
  selector: 'app-navbar',
  imports: [MatMenuModule, MatIconModule, NgIf, AsyncPipe], // Keep the imports as they were
  templateUrl: './navbar.component.html',
  styleUrls: ['./navbar.component.css']
})

export class NavbarComponent {
  constructor(
    public auth: AuthService,
    private sanitizer: DomSanitizer  // Inject DomSanitizer
  ) {}

  login() {
    this.auth.loginWithRedirect();
  }
  
  logout() {
    this.auth.logout({ logoutParams: { federated: true } });
  }

  // Method to sanitize the image URL
  sanitizeUrl(url: string): SafeUrl {
    return this.sanitizer.bypassSecurityTrustUrl(url);
  }
}



