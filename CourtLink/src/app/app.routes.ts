import { Routes } from '@angular/router';
import { SportsComponent } from './components/sports/sports.component';
import { CourtsComponent } from './components/courts/courts.component';

export const routes: Routes = [
  { path: '', component: SportsComponent }, // Default route shows SportsComponent first
  { path: 'courts/:sport', component: CourtsComponent }, // Navigates to CourtsComponent when a sport is selected
];
