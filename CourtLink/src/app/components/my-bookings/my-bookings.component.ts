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
  showCancelled = signal(false);

  // Toast & animation signals
  toastMessage = signal('');
  showToast = signal(false);
  cancellingId = signal<number | null>(null);

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
      next: (data) => {
        const sorted = data.sort((a, b) => {
          if (a.booking_status === 'booked' && b.booking_status !== 'booked') return -1;
          if (a.booking_status !== 'booked' && b.booking_status === 'booked') return 1;
          return 0;
        });
        this.bookings.set(sorted);
      },
      error: (err) => console.error("❌ Error fetching bookings:", err)
    });
  }

  cancelBooking(bookingId: number) {
    this.cancellingId.set(bookingId); // 🔥 Start animation

    const booking = this.bookings().find(b => b.booking_id === bookingId);
    if (!booking) {
      console.error("❌ Booking not found for cancellation:", bookingId);
      return;
    }

    const url = `http://localhost:8080/CancelBookingandUpdateSlot`;
    const body = {
      booking_id: booking.booking_id,
      court_id: booking.court_id,
      slot_index: booking.slot_index,
      sport_id: booking.sport_id
    };

    this.http.put(url, body, { responseType: 'text' as 'json' }).subscribe({
      next: () => {
        this.toastMessage.set("✅ Booking cancelled successfully!!");
        this.showToast.set(true);

        // 🕒 Delay updating list until animation finishes
        setTimeout(() => {
          const updated = this.bookings().map(b =>
            b.booking_id === bookingId ? { ...b, booking_status: 'cancelled' } : b
          );
          this.bookings.set(updated);
          this.cancellingId.set(null);
        }, 400);

        // Hide toast after 3s
        setTimeout(() => {
          this.showToast.set(false);
        }, 3000);
      },
      error: (err) => {
        console.error("❌ Error from backend:", err);
        alert("❌ Failed to cancel booking. Please try again.");
        this.cancellingId.set(null);
      }
    });
  }

  onToggleShowCancelled(event: Event) {
    const input = event.target as HTMLInputElement;
    this.showCancelled.set(input?.checked ?? false);
  }

  get filteredBookings() {
    const all = this.bookings();
    return this.showCancelled()
      ? all.sort((a, b) => {
          if (a.booking_status === 'booked' && b.booking_status !== 'booked') return -1;
          if (a.booking_status !== 'booked' && b.booking_status === 'booked') return 1;
          return 0;
        })
      : all.filter(b => b.booking_status === 'booked');
  }
}
