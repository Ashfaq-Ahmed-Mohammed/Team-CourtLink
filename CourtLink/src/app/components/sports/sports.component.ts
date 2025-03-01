import { ApiService } from './../../services/api.service';
import { Component } from '@angular/core';
import { CommonModule } from '@angular/common';
import { Router, RouterModule } from '@angular/router';
import { HttpClient } from '@angular/common/http';
import { firstValueFrom } from 'rxjs';

@Component({
  selector: 'app-sports',
  standalone: true,
  imports: [CommonModule, RouterModule], // HttpClient does not need to be imported here
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

  private BASE_URL = 'http://localhost:8080'; // Backend URL

  constructor(private router: Router, private http: HttpClient, private apiService: ApiService) {} // Inject HttpClient and ApiService

  /* selectSport(sport: string): void {
    const requestData = { sport };

    // Send JSON request to backend before navigation
    this.http.post(`${this.BASE_URL}/getCourts`, requestData).subscribe(
      () => {
        // Navigate to the courts page after successful API call
        this.router.navigate(['/courts', sport.toLowerCase()]);
      },
      (error) => {
        console.error('Error fetching courts:', error);
      }
    );
  } */

    async selectSport(sport: string): Promise<void> {
  
      try {
        // Send JSON to the backend at localhost:8080/getCourts
        await firstValueFrom(this.apiService.getCourts(sport));
        // On success, navigate to the CourtsComponent with the sport parameter in the URL
        this.router.navigate(['/courts', sport.toLowerCase()]);
      } catch (error) {
        console.error('Error fetching courts:', error);
      }
    }
  
}
