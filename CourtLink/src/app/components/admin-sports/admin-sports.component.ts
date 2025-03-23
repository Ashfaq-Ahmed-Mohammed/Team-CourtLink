import { Component, OnInit } from '@angular/core';
import { HttpClient } from '@angular/common/http';
import { CommonModule } from '@angular/common';
import { MatTableModule } from '@angular/material/table';
import { MatButtonModule } from '@angular/material/button';
import { MatDialog, MatDialogModule } from '@angular/material/dialog';
import { MatCardModule } from '@angular/material/card';

// ✅ Direct import of non-standalone dialog component
import { AddSportDialogComponent } from '../add-sport-dialog/add-sport-dialog.component';

@Component({
  selector: 'app-admin-sports',
  standalone: true,
  imports: [
    CommonModule,
    MatTableModule,
    MatButtonModule,
    MatDialogModule,
    MatCardModule,
    AddSportDialogComponent // ✅ since it's a regular component
  ],
  templateUrl: './admin-sports.component.html',
})
export class AdminSportsComponent implements OnInit {
  sports: { name: string }[] = [];
  displayedColumns: string[] = ['sno', 'name', 'action'];

  constructor(private http: HttpClient, private dialog: MatDialog) {}

  ngOnInit(): void {
    this.http.get<string[]>('http://localhost:8080/ListSports').subscribe({
      next: (data) => {
        this.sports = data.map((sportName) => ({ name: sportName }));
      },
      error: (err) => {
        console.error('Failed to fetch sports:', err);
      },
    });
  }

  openAddSportDialog(): void {
    const dialogRef = this.dialog.open(AddSportDialogComponent, {
      width: '400px',
    });

    dialogRef.afterClosed().subscribe((result: string | null) => {
      if (result) {
        const newSport = { Sport_name: result };

        this.http.post('http://localhost:8080/CreateSport', newSport).subscribe({
          next: () => {
            this.sports.push({ name: result });
            this.sports = [...this.sports]; // trigger update
          },
          error: (err) => {
            alert('Error creating sport: ' + err.error);
          },
        });
      }
    });
  }
}
