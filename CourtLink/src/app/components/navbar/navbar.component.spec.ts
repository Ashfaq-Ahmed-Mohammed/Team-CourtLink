import { ComponentFixture, TestBed } from '@angular/core/testing';
import { NavbarComponent } from './navbar.component';
import { NO_ERRORS_SCHEMA } from '@angular/core';
import { AuthService } from '@auth0/auth0-angular';

class DummyAuthService {
  // Provide a minimal observable for user$
  user$ = { subscribe: () => {} };
  login() {}
  logout() {}
}

describe('NavbarComponent (Standalone)', () => {
  let fixture: ComponentFixture<NavbarComponent>;
  let component: NavbarComponent;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [NavbarComponent],
      providers: [
        // Provide a dummy for the auth0.client injection token
        { provide: 'auth0.client', useValue: {} },
        // Override AuthService with our dummy implementation
        { provide: AuthService, useClass: DummyAuthService }
      ],
      // This schema ignores unknown elements and attributes in the template
      schemas: [NO_ERRORS_SCHEMA]
    }).compileComponents();
  });

  beforeEach(() => {
    fixture = TestBed.createComponent(NavbarComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  // Test 1: Ensure the Navbar component renders
  it('should render the navbar', () => {
    if (!component) {
      fail("NavbarComponent was not created.");
    }
    const navElement: HTMLElement = fixture.nativeElement.querySelector('nav');
    if (!navElement) {
      fail("Navbar element not rendered.");
    }
  });

  // Test 2: Verify that the logo and title are displayed correctly
  it('should display the logo and title', () => {
    const compiled = fixture.nativeElement as HTMLElement;
    const logoImg = compiled.querySelector('img[alt="UF Logo"]');
    if (!logoImg) {
      fail("Logo image not found.");
    }
    const titleSpan = compiled.querySelector('span');
    if (!titleSpan) {
      fail("Title element not found.");
    }
    if (!titleSpan || !titleSpan.textContent || titleSpan.textContent.trim() !== 'UFCourtLink') {
      fail("Expected title 'UFCourtLink', but got: " + (titleSpan && titleSpan.textContent ? titleSpan.textContent.trim() : ''));
    }
  });
});
