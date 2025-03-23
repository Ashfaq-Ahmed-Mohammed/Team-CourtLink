import { Component, OnInit, inject } from '@angular/core';
import { HttpClient } from '@angular/common/http';
import { CommonModule } from '@angular/common';
import { MatTableDataSource, MatTableModule } from '@angular/material/table';
import { MatButtonModule } from '@angular/material/button';
import { MatFormFieldModule } from '@angular/material/form-field';
import { MatInputModule } from '@angular/material/input';
import { MatDialog, MatDialogModule } from '@angular/material/dialog';

import { AddCourtDialogComponent } from '../add-court-dialog/add-court-dialog.component'; // ✅ Update path if needed

@Component({
  selector: 'app-admin-courts',
  standalone: true,
  imports: [
    CommonModule,
    MatTableModule,
    MatButtonModule,
    MatFormFieldModule,
    MatInputModule,
    MatDialogModule,
    AddCourtDialogComponent // ✅ Required to open the modal
  ],
  templateUrl: './admin-courts.component.html',
})
export class AdminCourtsComponent implements OnInit {
  http = inject(HttpClient);
  dialog = inject(MatDialog);

  courts = new MatTableDataSource<{ court_name: string; sport_name: string }>([]);
  displayedColumns: string[] = ['sno', 'court', 'sport', 'action'];

  ngOnInit(): void {
    this.fetchCourts();
  }

  fetchCourts(): void {
    this.http.get<any[]>('http://localhost:8080/ListCourts').subscribe({
      next: (data) => {
        this.courts.data = data;
      },
      error: (err) => {
        console.error('Failed to fetch courts:', err);
      },
    });
  }

  applyFilter(event: Event) {
    const filterValue = (event.target as HTMLInputElement).value;
    this.courts.filter = filterValue.trim().toLowerCase();
  }

  openAddCourtDialog(): void {
    const dialogRef = this.dialog.open(AddCourtDialogComponent, {
      width: '400px',
    });

    dialogRef.afterClosed().subscribe((result) => {
      if (result) {
        this.http.post('http://localhost:8080/CreateCourt', result).subscribe({
          next: () => {
            this.fetchCourts(); // 🔄 Refresh after adding court
          },
          error: (err) => {
            alert('Error creating court: ' + err.error);
          },
        });
      }
    });
  }
}
