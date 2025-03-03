import { MatIconModule } from '@angular/material/icon';
import { CommonModule } from '@angular/common';
import { Component, OnInit } from '@angular/core';
import { AuthService } from '@auth0/auth0-angular';
import { HttpClient, HttpHeaders } from '@angular/common/http';  // Import HttpClient and HttpHeaders

@Component({
  selector: 'app-navbar',
  imports: [MatIconModule, CommonModule],
  templateUrl: './navbar.component.html',
  styleUrls: ['./navbar.component.css']
})
export class NavbarComponent{

  constructor(
    public auth: AuthService,
    private http: HttpClient // Inject HttpClient
  ) {}

  // The login method will trigger fetching users and posting data.
  login() {
    this.auth.loginWithRedirect(); // Trigger Auth0 login and once successful, send user data
  }

  // Logout method
  logout() {
    this.auth.logout();
  }
}
