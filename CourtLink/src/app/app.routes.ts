import { Routes } from '@angular/router';
import { SportsComponent } from './components/sports/sports.component';
import { CourtsComponent } from './components/courts/courts.component';
import { AdminComponent } from './components/admin/admin.component';
export const routes: Routes = [
  { path: '', component: SportsComponent },            // Default: SportsComponent loads on startup
  { path: 'courts/:sport', component: CourtsComponent }, 
  { path: 'admin', component: AdminComponent } // CourtsComponent loads when a sport is selected
];
