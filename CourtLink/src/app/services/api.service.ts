import { Injectable } from '@angular/core';
import { HttpClient, HttpParams } from '@angular/common/http';
import { Observable } from 'rxjs';

@Injectable({
  providedIn: 'root'
})
export class ApiService {
  private BASE_URL = 'http://localhost:8080';

  constructor(private http: HttpClient) {}

  getCourts(sport: string): Observable<any> {
    const params = new HttpParams().set('sport', sport);
    return this.http.get<{ courts: any[] }>(`${this.BASE_URL}/getCourts`, { params });
  }
}
