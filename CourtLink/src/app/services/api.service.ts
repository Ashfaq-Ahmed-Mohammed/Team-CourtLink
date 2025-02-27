import { Injectable } from '@angular/core';
import { HttpClient } from '@angular/common/http';
import { Observable } from 'rxjs';

@Injectable({
  providedIn: 'root'
})
export class ApiService {
  private BASE_URL = 'http://localhost:8080';

  constructor(private http: HttpClient) {}

  getCourts(sport: string): Observable<any> {
    return this.http.post<{ courts: any[] }>(`${this.BASE_URL}/getCourts`, { sport });
  }
}
