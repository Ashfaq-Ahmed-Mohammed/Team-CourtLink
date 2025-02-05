import { Component } from '@angular/core';
import { CommonModule } from '@angular/common'; // Import CommonModule for *ngFor

@Component({
  selector: 'app-sports',
  standalone: true,
  imports: [CommonModule], // Include CommonModule for *ngFor support
  templateUrl: './sports.component.html',
  styleUrls: ['./sports.component.css']
})
export class SportsComponent {
  sports = [
    { name: 'Basketball', icon: 'https://img.icons8.com/?size=100&id=196&format=png&color=000000' },
    { name: 'Soccer', icon: 'https://img.icons8.com/?size=100&id=9820&format=png&color=000000' },
    { name: 'Tennis', icon: 'https://img.icons8.com/?size=100&id=48991&format=png&color=000000' },
    { name: 'Badminton', icon: 'https://img.icons8.com/?size=100&id=24308&format=png&color=000000' },
    { name: 'Cricket', icon: 'https://img.icons8.com/?size=100&id=4252&format=png&color=000000' },
  ];

  selectSport(sport: string): void {
    console.log('Selected Sport:', sport);
  }
}
