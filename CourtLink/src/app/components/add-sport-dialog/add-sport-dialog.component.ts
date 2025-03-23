import { Component } from '@angular/core';
import { MatDialogRef } from '@angular/material/dialog';
import { CommonModule } from '@angular/common';
import { FormsModule } from '@angular/forms';
import { MatButtonModule } from '@angular/material/button';
import { MatInputModule } from '@angular/material/input';

@Component({
  selector: 'app-add-sport-dialog',
  standalone: true,
  templateUrl: './add-sport-dialog.component.html',
  styleUrls: ['./add-sport-dialog.component.css'],
  imports: [CommonModule, FormsModule, MatButtonModule, MatInputModule],
})
export class AddSportDialogComponent {
  sportName: string = '';

  constructor(public dialogRef: MatDialogRef<AddSportDialogComponent>) {}

  cancel(): void {
    this.dialogRef.close(null);
  }

  confirm(): void {
    this.dialogRef.close(this.sportName.trim());
  }
}
