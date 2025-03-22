import { Component } from '@angular/core';
import { CommonModule } from '@angular/common';
@Component({
  selector: 'app-admin',
  standalone: true, // ✅ This is required
  imports: [CommonModule],      // Add other imports like Angular Material modules if needed
  templateUrl: './admin.component.html',
  styleUrls: ['./admin.component.css'] // <- typo fix: should be 'styleUrls' not 'styleUrl'
})
export class AdminComponent {}
