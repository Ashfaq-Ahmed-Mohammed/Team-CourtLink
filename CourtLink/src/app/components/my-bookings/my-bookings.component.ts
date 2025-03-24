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
      } else {
        console.warn("❌ Could not fetch user email.");
      }
    });
  }

  fetchBookings(email: string) {
    const url = `http://localhost:8080/listBookings?email=${encodeURIComponent(email)}`;
    this.http.get<any[]>(url).subscribe({
      next: (data) => this.bookings.set(data),
      error: (err) => console.error("❌ Error fetching bookings:", err)
    });
  }

  cancelBooking(bookingId: number) {
    const url = `http://localhost:8080/cancelBooking?id=${bookingId}`;
    this.http.delete(url).subscribe({
      next: () => {
        console.log("✅ Booking cancelled:", bookingId);
        this.bookings.set(this.bookings().filter(b => b.booking_id !== bookingId));
      },
      error: (err) => {
        console.error("❌ Failed to cancel booking:", err);
        alert("Failed to cancel booking. Please try again.");
      }
    });
  }
}
