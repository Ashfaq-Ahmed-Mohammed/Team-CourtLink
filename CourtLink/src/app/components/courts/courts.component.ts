import { Component } from '@angular/core';
import { CommonModule } from '@angular/common';
import { ActivatedRoute } from '@angular/router'; // Import ActivatedRoute

@Component({
  selector: 'app-courts',
  standalone: true,
  imports: [CommonModule],
  templateUrl: './courts.component.html',
  styleUrls: ['./courts.component.css']
})
export class CourtsComponent {
  selectedSport: string | null = null; // Store the sport name dynamically
  courts = [
    { name: 'Center Court', location: 'National Sports Arena', image: 'https://img.icons8.com/?size=100&id=HVTfsoSWrdwS&format=png&color=000000' },
    { name: 'Grand Slam Court', location: 'Downtown Tennis Club', image: 'https://img.icons8.com/?size=100&id=HVTfsoSWrdwS&format=png&color=000000' },
    { name: 'Clay Court', location: 'City Sports Complex', image: 'https://img.icons8.com/?size=100&id=HVTfsoSWrdwS&format=png&color=000000' }
  ];

  constructor(private route: ActivatedRoute) {}

  ngOnInit(): void {
    this.selectedSport = this.route.snapshot.paramMap.get('sport'); // Retrieve sport from the route
  }
}
