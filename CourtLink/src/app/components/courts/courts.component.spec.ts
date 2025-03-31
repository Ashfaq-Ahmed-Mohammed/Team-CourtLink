import { ComponentFixture, TestBed, waitForAsync, fakeAsync, tick } from '@angular/core/testing';
import { CourtsComponent } from './courts.component';
import { MatIconModule } from '@angular/material/icon';
import { provideRouter } from '@angular/router';
import { HttpClientModule } from '@angular/common/http';
import { of } from 'rxjs';
import { ApiService } from '../../services/api.service';
import { AuthService } from '@auth0/auth0-angular';

describe('CourtsComponent (Angular 17+)', () => {
  let fixture: ComponentFixture<CourtsComponent>;
  let component: CourtsComponent;

  beforeEach(waitForAsync(() => {
    TestBed.configureTestingModule({
      imports: [MatIconModule, HttpClientModule],
      providers: [
        { provide: ApiService, useValue: { getCourts: jasmine.createSpy() } },
        { provide: AuthService, useValue: { user$: of({ email: 'test@example.com' }) } },
        provideRouter([]),
      ],
    }).compileComponents().then(() => {
      fixture = TestBed.createComponent(CourtsComponent);
      component = fixture.componentInstance;
      fixture.detectChanges();
    });
  }));

  // Test 1: Check if component is created
  it('should create the CourtsComponent', () => {
    expect(component).toBeTruthy(); // Ensures the component is created successfully
  });

  // Test 2: Ensure that the courts signal is initialized as an empty array
  it('should initialize courts signal as an empty array', () => {
    expect(component.courts()).toEqual([]); // Ensures courts are initialized correctly
  });

  // Test 3: Check if default time slots are populated
  it('should have time slots initialized', () => {
    expect(component.timeSlots.length).toBeGreaterThan(0); // Ensures that time slots are populated
  });

  // Test 4: Ensure the `bookingSuccess` signal is initialized correctly
  it('should initialize bookingSuccess signal as false', () => {
    expect(component.bookingSuccess()).toBe(false); // Ensures booking success signal is initialized
  });

  // Test 5: Check if the modal can be closed (i.e., selectedBooking is reset)
  it('should close the modal and reset selectedBooking', () => {
    // Create a complete court object with all required properties
    const court = {
      name: 'Court A',
      id: 1,
      status: 1, // Add required status
      slots: [1, 1, 0], // Add required slots
      image: 'test-image-url', // Add required image
      location: 'Location A', // Add required location
      floor: 'Floor 1', // Add required floor
      surface: 'Hard', // Add required surface
      capacity: '4', // Add required capacity
      type: 'Tennis' // Add required type
    };

    // Set the selectedBooking signal with the complete court object
    component.selectedBooking.set({
      court, // Use the complete court object
      time: '10 AM',
      slotIndex: 0
    });

    // Call the closeModal method
    component.closeModal();

    // Assert that the selectedBooking signal is reset to null
    expect(component.selectedBooking()).toBeNull();
  });

  // Test 6: Check if the `selectedSport` signal is set when `sport` is passed
  it('should set selectedSport signal when sport is passed', () => {
    const sport = 'Basketball';
    component.selectedSport.set(sport); // Simulate selecting a sport
    expect(component.selectedSport()).toBe(sport); // Ensures selectedSport is set correctly
  });
});
