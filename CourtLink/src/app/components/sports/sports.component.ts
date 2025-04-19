import { Component, OnInit } from '@angular/core';
import { CommonModule } from '@angular/common';
import { RouterModule, Router } from '@angular/router';
import { HttpClientModule } from '@angular/common/http';
import { AuthService } from '@auth0/auth0-angular';
import { ApiService } from './../../services/api.service';
import { Observable, firstValueFrom } from 'rxjs';

@Component({
  selector: 'app-sports',
  standalone: true,
  imports: [
    CommonModule,
    RouterModule,
    HttpClientModule
  ],
  templateUrl: './sports.component.html',
  styleUrls: ['./sports.component.css']
})
export class SportsComponent implements OnInit {
  isAuthenticated$!: Observable<boolean>;
  isLoading$!: Observable<boolean>;

  currentSlideIndex = 0;
  currentSlide = '';
  slides = [
    'https://snworksceo.imgix.net/ufa/0d1ef788-7d24-4f52-b28f-bc25f8a28e14.sized-1000x1000.jpg?w=1000',
    'https://gatorswire.usatoday.com/wp-content/uploads/sites/7/2023/11/UF-Arky-23-14.jpeg?w=1000&h=600&crop=1',
    'https://www.onlygators.com/wp-content/uploads/2024/11/albert-swamp-fans-650x300.png',
    'https://d2b5htfb6s9xp9.cloudfront.net/images/2024/8/15/230916_FBTNEditor_9984_EmmaBissell_kI24p.jpg?width=1024&height=576&mode=crop',
    'https://snworksceo.imgix.net/ufa/b41a8d6a-0d1e-4827-90c3-0c182f355037.sized-1000x1000.jpeg?w=1000',
    'https://assets.goal.com/images/v3/blt0d938608573ececc/Florida%20Gators%20football%20.jpg?auto=webp&format=pjpg&width=3840&quality=60',
    'https://cdn.learfield.com/wp-content/uploads/2022/04/GatorsLearfield.com-1.jpg',
    'https://www.gatorcountry.com/wp-content/uploads/2024/11/Florida-Gators-running-back-Jadan-Baugh-13_Florida-Gators-Football-vs-Georgia-Bulldogs_0258-1021x580.jpg'
  ];

  sports = [
    { name: 'Basketball', icon: 'https://img.icons8.com/?size=100&id=196&format=png&color=000000' },
    { name: 'Soccer',     icon: 'https://img.icons8.com/?size=100&id=9820&format=png&color=000000' },
    { name: 'Tennis',     icon: 'https://img.icons8.com/?size=100&id=48991&format=png&color=000000' },
    { name: 'Badminton',  icon: 'https://img.icons8.com/?size=100&id=24308&format=png&color=000000' },
    { name: 'Cricket',    icon: 'https://img.icons8.com/?size=100&id=4252&format=png&color=000000' },
  ];

  constructor(
    private router: Router,
    private apiService: ApiService,
    public auth: AuthService
  ) {}

  ngOnInit(): void {
    this.isAuthenticated$ = this.auth.isAuthenticated$;
    this.isLoading$       = this.auth.isLoading$;

    this.currentSlide = this.slides[0];
    setInterval(() => {
      this.currentSlideIndex = (this.currentSlideIndex + 1) % this.slides.length;
      this.currentSlide = this.slides[this.currentSlideIndex];
    }, 3000);
  }

  login(): void {
    this.auth.loginWithRedirect();
  }

  async selectSport(sport: string): Promise<void> {
    const key = sport.toLowerCase();
    console.log('Sport clicked (key):', key);
    try {
      await firstValueFrom(this.apiService.getCourts(key));
      console.log('API success, navigating…');
      this.router.navigate(['/courts', key]);
    } catch (error) {
      console.error('Error fetching courts:', error);
    }
  }
}
