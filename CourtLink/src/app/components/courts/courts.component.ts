import { Component, effect, signal } from '@angular/core';
import { CommonModule } from '@angular/common';
import { ActivatedRoute, RouterModule } from '@angular/router';
import { ApiService } from '../../services/api.service';

// Extend your interface to hold extra fields (floor, surface, capacity, type)
export interface Court {
  name: string;
  id: number;
  status: number;       // e.g. 1 => "Available", 0 => "Not Available"
  slots: number[];      // 1 => can reserve, 0 => disabled
  image: string;        // court image URL or placeholder
  location: string;     // e.g. "Main Building, Floor 1"
  floor: string;        // e.g. "Floor 1"
  surface: string;      // e.g. "Hardwood"
  capacity: string;     // e.g. "10 players"
  type: string;         // e.g. "Indoor Basketball" or "Outdoor Basketball"
}

@Component({
  selector: 'app-courts',
  standalone: true,
  imports: [CommonModule, RouterModule],
  templateUrl: './courts.component.html',
  styleUrls: ['./courts.component.css']
})
export class CourtsComponent {
  // Signals
  selectedSport = signal<string>('');
  courts = signal<Court[]>([]);
  selectedTime = signal<string>('');  // Track selected time slot
  selectedBooking = signal<{ court: Court, time: string } | null>(null);  // Booking details

  timeSlots: string[] = [];

  constructor(private route: ActivatedRoute, private apiService: ApiService) {
    // Generate time slots from 8 AM to 6 PM
    for (let hour = 8; hour < 18; hour++) {
      this.timeSlots.push(hour < 12 ? `${hour} AM` : hour === 12 ? `12 PM` : `${hour - 12} PM`);
    }

    // Fetch courts when the sport changes
    effect(() => {
      const sport = this.route.snapshot.paramMap.get('sport') || '';
      this.selectedSport.set(sport);
      if (sport) {
        this.fetchCourts(sport);
      }
    });
  }

  fetchCourts(sport: string): void {
    this.apiService.getCourts(sport).subscribe({
      next: (data) => {
        let courtsData: any[] = [];
        if (data && data.courts && Array.isArray(data.courts)) {
          courtsData = data.courts;
        } else if (Array.isArray(data)) {
          courtsData = data;
        }

        const mapped = courtsData.map((c: any, i: number) => ({
          name: c.CourtName || `Basketball Court #${i + 1}`,
          id: c.CourtID || 0,
          status: c.CourtStatus ?? 1,     // 1 => available by default
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

  /**
   * Selects a time slot and prepares the booking modal.
   */
  selectTime(court: Court, time: string): void {
    if (this.isSlotAvailable(court, time)) {
      this.selectedTime.set(time);
      this.selectedBooking.set({ court, time });
    }
  }

  /**
   * Checks if the time slot is available for a court.
   */
  isSlotAvailable(court: Court, time: string): boolean {
    const index = this.timeSlots.indexOf(time);
    return court.slots[index] === 1;
  }

  /**
   * Returns a label based on court availability.
   */
  getAvailabilityLabel(court: Court): string {
    return court.status === 1 ? 'Available' : 'Not Available';
  }

  /**
   * Cancels the booking process.
   */
  cancelBooking(): void {
    this.selectedBooking.set(null);
  }

  /**
   * Confirms booking and sends data to backend.
   */
  confirmBooking(): void {
    const booking = this.selectedBooking();
    if (!booking) return;

    const bookingData = {
      sport: this.selectedSport(),
      court: booking.court.name,
      time: booking.time,
      user: "test-user@example.com" // Replace with actual user data (Auth0 integration)
    };

    this.apiService.bookCourt(bookingData).subscribe({
      next: () => {
        alert('Booking successful!');
        this.selectedBooking.set(null);
      },
      error: () => {
        alert('Booking failed. Try again.');
      }
    });
  }
}
