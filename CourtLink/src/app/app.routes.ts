import { Routes } from '@angular/router';
import { SportsComponent } from './components/sports/sports.component';
import { CourtsComponent } from './components/courts/courts.component';
import { AdminComponent } from './components/admin/admin.component';
import { MyBookingsComponent } from './components/my-bookings/my-bookings.component'; // ✅ Import MyBookingsComponent
import { AdminSportsComponent } from './components/admin-sports/admin-sports.component';
import { AdminCourtsComponent } from './components/admin-courts/admin-courts.component';

export const routes: Routes = [
  { path: '', component: SportsComponent },                      // Home page - sports selector
  { path: 'courts/:sport', component: CourtsComponent },         // Courts for selected sport
  { path: 'admin', component: AdminComponent },                  // Admin panel
  { path: 'my-bookings', component: MyBookingsComponent },
  { path: 'admin/sports', component: AdminSportsComponent },
  { path: 'admin/courts', component: AdminCourtsComponent }
];      // ✅ My Bookings page



