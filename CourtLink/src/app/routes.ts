import { Routes } from '@angular/router';
import { SportsComponent } from './components/sports/sports.component';
import { CourtsComponent } from './components/courts/courts.component';

export const routes: Routes = [
  { path: '', component: SportsComponent }, // Landing page with sports tiles
  { path: 'courts/:sport', component: CourtsComponent }, // Courts page based on selected sport
];
