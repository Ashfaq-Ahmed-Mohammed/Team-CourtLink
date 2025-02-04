import { Component } from '@angular/core';
import { MatMenuModule } from '@angular/material/menu'; // Import MatMenuModule
import { MatIconModule } from '@angular/material/icon';


@Component({
  selector: 'app-navbar',
  imports: [MatMenuModule, MatIconModule],
  templateUrl: './navbar.component.html',
  styleUrl: './navbar.component.css'
})
export class NavbarComponent {
  menuOpen: boolean = false;
  dropdownOpen = false;
}