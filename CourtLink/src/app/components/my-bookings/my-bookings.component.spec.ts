import { ComponentFixture, TestBed, fakeAsync, tick, waitForAsync } from '@angular/core/testing';
import { MyBookingsComponent } from './my-bookings.component';
import { HttpClientTestingModule } from '@angular/common/http/testing';
import { AuthService } from '@auth0/auth0-angular';
import { of } from 'rxjs';

describe('MyBookingsComponent (Angular 17)', () => {
  let fixture: ComponentFixture<MyBookingsComponent>;
  let component: MyBookingsComponent;

  const mockAuthService = {
    user$: of({ email: 'test@example.com' }),  // Simulate user being logged in with email
  };

  beforeEach(waitForAsync(() => {
    TestBed.configureTestingModule({
      imports: [HttpClientTestingModule],  // Use HttpClientTestingModule for mocking HTTP requests
      providers: [
        { provide: AuthService, useValue: mockAuthService },
      ],
    }).compileComponents().then(() => {
      fixture = TestBed.createComponent(MyBookingsComponent);
      component = fixture.componentInstance;
      fixture.detectChanges();
    });
  }));

  // Test 1: Check if the component is created successfully
  it('should create the MyBookingsComponent', fakeAsync(() => {
    fixture.detectChanges();
    tick();
    expect(component).toBeTruthy(); // Ensures the component is created successfully
  }));

  // Test 2: Ensure that the "No bookings yet" message appears when there are no bookings
  it('should display "No bookings yet" message when there are no bookings', fakeAsync(() => {
    component.bookings.set([]); // Simulating no bookings
    fixture.detectChanges();
    tick();

    const message = fixture.nativeElement.querySelector('.text-center');
    expect(message).not.toBeNull();
    expect(message.textContent.trim()).toBe('🏀 No bookings yet — time to hit the court!');
  }));

  // Test 5: Ensure the bookings list has correct initial length
  it('should have an empty bookings array initially', fakeAsync(() => {
    fixture.detectChanges();
    tick(); // Simulate async behavior

    expect(component.bookings().length).toBe(0); // Ensure bookings array is empty
  }));
});
