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
  selectedTime = signal<string>('');  // Track a globally selected time if needed

  // Time labels from 8 AM to 5 PM (or 6 PM)
  timeSlots: string[] = [];

  constructor(private route: ActivatedRoute, private apiService: ApiService) {
    // Generate times from 8 AM (8) to 6 PM (18) => 10 slots
    for (let hour = 8; hour < 18; hour++) {
      // Convert 24-hour format to a user-friendly label
      if (hour < 12) {
        this.timeSlots.push(`${hour} AM`);
      } else if (hour === 12) {
        this.timeSlots.push(`12 PM`);
      } else {
        this.timeSlots.push(`${hour - 12} PM`);
      }
    }

    // Reactive effect to fetch courts based on the route param
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
        // Handle the response whether it's an array or an object with 'courts'
        let courtsData: any[] = [];
        if (data && data.courts && Array.isArray(data.courts)) {
          courtsData = data.courts;
        } else if (Array.isArray(data)) {
          courtsData = data;
        }

        // Map to our extended Court interface
        const mapped = courtsData.map((c: any, i: number) => ({
          name: c.CourtName || `Basketball Court #${i + 1}`,
          id: c.CourtID || 0,
          status: c.CourtStatus ?? 1,     // 1 => available by default
          slots: c.Slots || [],
          image: c.Image || 'https://i.bleacherreport.net/images/team_logos/328x328/florida_gators_football.png?canvas=492,328',
        })) as Court[];

        this.courts.set(mapped);
      },
      error: (err) => console.error('Error fetching courts:', err)
    });
  }

  /**
   * Called when a time button is clicked for a specific court/time.
   * Only invoked if the slot is available (1).
   */
  selectTime(court: Court, time: string): void {
    // Example: Log the selection (extend this to handle reservations)
    console.log(`Selected ${time} for ${court.name}`);
    this.selectedTime.set(time);
  }

  /**
   * Utility to determine if a given time is available for a court.
   */
  isSlotAvailable(court: Court, time: string): boolean {
    const index = this.timeSlots.indexOf(time);
    return court.slots[index] === 1;
  }

  /**
   * Returns a text label (e.g. "Available" / "Not Available") based on court.status.
   */
  getAvailabilityLabel(court: Court): string {
    return court.status === 1 ? 'Available' : 'Not Available';
  }
}
