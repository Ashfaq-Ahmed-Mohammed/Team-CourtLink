import { Component, effect, signal } from '@angular/core';
import { CommonModule } from '@angular/common';
import { ActivatedRoute, RouterModule } from '@angular/router';
import { ApiService } from '../../services/api.service';

@Component({
  selector: 'app-courts',
  standalone: true,
  imports: [CommonModule, RouterModule],
  templateUrl: './courts.component.html',
  styleUrls: ['./courts.component.css']
})
export class CourtsComponent {
  // Signals for state management
  selectedSport = signal<string>('');
  courts = signal<any[]>([]);

  constructor(private route: ActivatedRoute, private apiService: ApiService) {
    // Read sport name from URL after constructor initializes route
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
      next: (data) => this.courts.set(data.courts || []),
      error: (error) => console.error('Error fetching courts:', error)
    });
  }
}
