import { Component, effect, signal } from '@angular/core';
import { CommonModule } from '@angular/common';
import { ActivatedRoute, Router, RouterModule } from '@angular/router';
import { ApiService } from '../../services/api.service';
import { HttpClientModule } from '@angular/common/http'; // <-- Add HttpClientModule here
import { HttpClient } from '@angular/common/http';
import { AuthService } from '@auth0/auth0-angular';

export interface Court {
  name: string;
  id: number;
  status: number;
  slots: number[];
  image: string;
  location: string;
  floor: string;
  surface: string;
  capacity: string;
  type: string;
}

interface Booking {
  court: Court;
  time: string;
  slotIndex: number;
}

@Component({
  selector: 'app-courts',
  standalone: true,  // This marks the component as standalone
  imports: [CommonModule, RouterModule, HttpClientModule],  // <-- Add HttpClientModule to imports
  templateUrl: './courts.component.html',
  styleUrls: ['./courts.component.css']
})
export class CourtsComponent {
  selectedSport = signal<string>('');
  courts = signal<Court[]>([]);
  selectedTime = signal<string>('');
  selectedBooking = signal<Booking | null>(null);
  userEmail = signal<string | null>(null);
  bookingSuccess = signal<boolean>(false); // ✅ Toast flag

  timeSlots: string[] = [];

  constructor(
    private route: ActivatedRoute,
    private apiService: ApiService,
    private http: HttpClient,
    private auth: AuthService,
    private router: Router
  ) {
    for (let hour = 8; hour < 18; hour++) {
      this.timeSlots.push(hour < 12 ? `${hour} AM` : hour === 12 ? `12 PM` : `${hour - 12} PM`);
    }

    effect(() => {
      const sport = this.route.snapshot.paramMap.get('sport') || '';
      this.selectedSport.set(sport);
      if (sport) {
        this.fetchCourts(sport);
      }
    });

    this.auth.user$.subscribe((user) => {
      if (user?.email) {
        this.userEmail.set(user.email);
      }
    });
  }

  fetchCourts(sport: string): void {
    this.apiService.getCourts(sport).subscribe({
      next: (data) => {
        let courtsData: any[] = [];
        if (data?.courts && Array.isArray(data.courts)) {
          courtsData = data.courts;
        } else if (Array.isArray(data)) {
          courtsData = data;
        }

        const mapped = courtsData.map((c: any, i: number) => ({
          name: c.CourtName || `Basketball Court #${i + 1}`,
          id: c.CourtID || 0,
          status: c.CourtStatus ?? 1,
          slots: c.Slots || [],
          image: c.Image || 'https://i.bleacherreport.net/images/team_logos/328x328/florida_gators_football.png?canvas=492,328',
          location: c.Location || 'Unknown',
          floor: c.Floor || 'N/A',
          surface: c.Surface || 'Unknown',
          capacity: c.Capacity || 'Unknown',
          type: c.Type || 'Unknown'
        })) as Court[];

        this.courts.set(mapped);
      },
      error: (err) => console.error('Error fetching courts:', err)
    });
  }

  selectTime(court: Court, time: string): void {
    const slotIndex = this.timeSlots.indexOf(time);
    if (this.isSlotAvailable(court, time)) {
      this.selectedTime.set(time);
      this.selectedBooking.set({ court, time, slotIndex });
    }
  }

  isSlotAvailable(court: Court, time: string): boolean {
    const index = this.timeSlots.indexOf(time);
    return court.slots[index] === 1;
  }

  getAvailabilityLabel(court: Court): string {
    return court.status === 1 ? 'Available' : 'Not Available';
  }

  closeModal(): void {
    this.selectedBooking.set(null);
  }

  confirmBooking(): void {
    const booking = this.selectedBooking();
    if (!booking) {
      alert("No booking selected.");
      return;
    }

    const sportName = this.selectedSport();
    const bookingPayload = {
      Court_Name: booking.court.name,
      Court_ID: booking.court.id,
      Slot_Index: booking.slotIndex,
      Sport_name: sportName,
      Sport_ID: sportName.toUpperCase(),
      Customer_email: this.userEmail()
    };

    this.closeModal();

    this.http.put('http://localhost:8080/UpdateCourtSlotandBooking', bookingPayload, {
      responseType: 'text'
    }).subscribe({
      next: () => {
        this.bookingSuccess.set(true); // ✅ Show toast
        setTimeout(() => {
          window.location.reload();
        }, 1000); // ✅ Reload after 1.5s
      },
      error: (error) => {
        console.error('⛔ Booking failed:', error);
        alert('❌ Booking failed. Please try again.');
      }
    });
  }
}
