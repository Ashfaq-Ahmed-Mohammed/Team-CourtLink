// app/app.routes.ts
import { Routes } from '@angular/router';
import { SportsComponent }   from './components/sports/sports.component';
import { CourtsComponent }   from './components/courts/courts.component';
import { AdminComponent }    from './components/admin/admin.component';
import { MyBookingsComponent }   from './components/my-bookings/my-bookings.component';
import { AdminSportsComponent }  from './components/admin-sports/admin-sports.component';
import { AdminCourtsComponent }  from './components/admin-courts/admin-courts.component';

export const routes: Routes = [
  { path: '',            component: SportsComponent },
  { path: 'courts/:sport', component: CourtsComponent },
  { path: 'admin',       component: AdminComponent },
  { path: 'my-bookings', component: MyBookingsComponent },
  { path: 'admin/sports',  component: AdminSportsComponent },
  { path: 'admin/courts',  component: AdminCourtsComponent },
  { path: '**',          redirectTo: '' }
];
