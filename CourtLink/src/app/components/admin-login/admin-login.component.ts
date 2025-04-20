import { Component } from '@angular/core';
import { CommonModule } from '@angular/common';
import { FormsModule } from '@angular/forms';
import { HttpClient } from '@angular/common/http';
import { Router } from '@angular/router';
import { MatInputModule } from '@angular/material/input';
import { MatButtonModule } from '@angular/material/button';

@Component({
  selector: 'app-admin-login',
  standalone: true,
  imports: [CommonModule, FormsModule, MatInputModule, MatButtonModule],
  templateUrl: './admin-login.component.html',
})
export class AdminLoginComponent {
  username: string = '';
  password: string = '';
  errorMessage: string = '';

  constructor(private http: HttpClient, private router: Router) {}

  login(): void {
    this.http.post<{ message: string }>('http://localhost:8080/AdminLogin', {
      username: this.username,
      password: this.password,
    }).subscribe({
      next: () => {
        localStorage.setItem('admin-auth', 'true'); // Simple session check
        this.router.navigate(['/admin-portal']);
      },
      error: (err) => {
        this.errorMessage = err.error || 'Login failed. Check credentials.';
      }
    });
  }
}
