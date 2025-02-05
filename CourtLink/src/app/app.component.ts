import { Component } from '@angular/core';
import { RouterOutlet } from '@angular/router';
import { NavbarComponent } from './components/navbar/navbar.component';
import { SportsComponent } from './components/sports/sports.component';


@Component({
  selector: 'app-root',
  imports: [RouterOutlet, NavbarComponent, SportsComponent],
  templateUrl: './app.component.html',
  styleUrl: './app.component.css'
})
export class AppComponent {
  title = 'CourtLink';
}
