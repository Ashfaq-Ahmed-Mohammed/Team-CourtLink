import { ComponentFixture, TestBed, waitForAsync } from '@angular/core/testing';
import { NavbarComponent } from './navbar.component';
import { AuthService } from '@auth0/auth0-angular';
import { MatIconModule } from '@angular/material/icon';
import { of, Observable } from 'rxjs';
import { provideRouter } from '@angular/router';
import { HttpClientTestingModule, HttpTestingController } from '@angular/common/http/testing';

describe('NavbarComponent (HttpClientTestingModule)', () => {
  let fixture: ComponentFixture<NavbarComponent>;
  let component: NavbarComponent;
  let httpTestingController: HttpTestingController;

  const authServiceMock: {
    user$: Observable<{ name: string; picture: string } | null>;
    loginWithRedirect: jasmine.Spy;
    logout: jasmine.Spy;
  } = {
    user$: of({ name: 'Test User', picture: 'https://example.com/image.png' }),
    loginWithRedirect: jasmine.createSpy('loginWithRedirect'),
    logout: jasmine.createSpy('logout'),
  };

  beforeEach(waitForAsync(() => {
    TestBed.configureTestingModule({
      imports: [
        MatIconModule,
        HttpClientTestingModule, // ✅ Add HttpClientTestingModule
      ],
      providers: [
        { provide: AuthService, useValue: authServiceMock },
        provideRouter([]),
      ],
    }).compileComponents().then(() => {
      fixture = TestBed.createComponent(NavbarComponent);
      component = fixture.componentInstance;
      httpTestingController = TestBed.inject(HttpTestingController); // ✅ Inject HttpTestingController
      fixture.detectChanges();
    });
  }));

  it('should create the component', () => {
    if (!(component instanceof NavbarComponent)) {
      throw new Error('NavbarComponent was not created');
    }
  });

  it('should call logout on profile icon click', () => {
    authServiceMock.logout.calls.reset();

    const logoutBtn = fixture.nativeElement.querySelector('button[mat-icon-button] img');
    if (!logoutBtn) {
      throw new Error('Logout button not found');
    }

    logoutBtn.click();

    if (authServiceMock.logout.calls.count() !== 1) {
      throw new Error('Logout function was not called exactly once');
    }
  });

  it('should show "My Bookings" button with correct text', () => {
    const bookingsBtn = fixture.nativeElement.querySelector('a[routerlink="/my-bookings"]');
    if (!bookingsBtn) {
      throw new Error('"My Bookings" button not found');
    }

    const btnText = bookingsBtn.textContent?.trim();
    if (btnText !== 'My Bookings') {
      throw new Error(`Expected "My Bookings", but got "${btnText}"`);
    }
  });

  it('should display "My Bookings" link in the navbar', () => {
    // Trigger change detection to ensure the DOM is updated
    fixture.detectChanges();
  
    // Query for the "My Bookings" link
    const bookingsLink = fixture.nativeElement.querySelector('a[routerlink="/my-bookings"]');
    
    // If the link is not found, throw an error
    if (!bookingsLink) {
      throw new Error('"My Bookings" link not found');
    }
  
    // Check if the link text is correct
    const linkText = bookingsLink.textContent.trim();
    if (linkText !== 'My Bookings') {
      throw new Error(`Expected "My Bookings" but got "${linkText}"`);
    }
  });
  

  afterEach(() => {
    httpTestingController.verify();  // ✅ Verify that there are no outstanding HTTP requests
  });
});
