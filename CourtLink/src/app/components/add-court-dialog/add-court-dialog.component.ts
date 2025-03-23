import { Component } from '@angular/core';
import { MatDialogRef } from '@angular/material/dialog';
import { CommonModule } from '@angular/common';
import { FormsModule } from '@angular/forms';
import { MatButtonModule } from '@angular/material/button';
import { MatInputModule } from '@angular/material/input';

@Component({
  selector: 'app-add-court-dialog',
  standalone: true,
  templateUrl: './add-court-dialog.component.html',
  styleUrls: ['./add-court-dialog.component.css'],
  imports: [CommonModule, FormsModule, MatButtonModule, MatInputModule],
})
export class AddCourtDialogComponent {
  courtName: string = '';
  sportName: string = '';

  constructor(public dialogRef: MatDialogRef<AddCourtDialogComponent>) {}

  cancel(): void {
    this.dialogRef.close(null);
  }

  confirm(): void {
    if (this.courtName.trim() && this.sportName.trim()) {
      this.dialogRef.close({
        Court_Name: this.courtName.trim(),
        Sport_name: this.sportName.trim(),
      });
    }
  }
}
