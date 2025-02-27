import { Routes } from '@angular/router';
import { SportsComponent } from './components/sports/sports.component';
import { CourtsComponent } from './components/courts/courts.component';

export const routes: Routes = [
  { path: '', component: SportsComponent },            // Default: SportsComponent loads on startup
  { path: 'courts/:sport', component: CourtsComponent }  // CourtsComponent loads when a sport is selected
];
