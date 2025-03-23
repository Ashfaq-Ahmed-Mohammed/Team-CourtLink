import { signal } from '@angular/core';
import { Component } from '@angular/core';
import { CommonModule } from '@angular/common';
import { HttpClient } from '@angular/common/http';
import { AuthService } from '@auth0/auth0-angular';

@Component({
  selector: 'app-my-bookings',
  standalone: true,
  imports: [CommonModule],
  templateUrl: './my-bookings.component.html',
  styleUrls: ['./my-bookings.component.css']
})
export class MyBookingsComponent {
  bookings = signal<any[]>([]);
  userEmail = signal<string | null>(null);

  constructor(private http: HttpClient, private auth: AuthService) {
    this.auth.user$.subscribe(user => {
      if (user?.email) {
        this.userEmail.set(user.email);
        this.fetchBookings(user.email);
      }
    });
  }

  fetchBookings(email: string) {
    this.http.get<any[]>(`http://localhost:8080/GetUserBookings?email=${email}`)
      .subscribe(data => this.bookings.set(data));
  }
}
