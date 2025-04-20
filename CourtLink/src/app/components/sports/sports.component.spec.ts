// sports.component.spec.ts
import { ComponentFixture, TestBed } from '@angular/core/testing';
import { SportsComponent } from './sports.component';
import { RouterTestingModule } from '@angular/router/testing';
import { HttpClientTestingModule } from '@angular/common/http/testing';
import { AuthService } from '@auth0/auth0-angular';
import { ApiService } from './../../services/api.service';
import { of } from 'rxjs';

class MockAuthService {
  isAuthenticated$ = of(false);
  isLoading$       = of(false);
  loginWithRedirect(): void {}
}

class MockApiService {
  getCourts(_sport: string) {
    return of({});
  }
}

describe('SportsComponent', () => {
  let fixture: ComponentFixture<SportsComponent>;
  let component: SportsComponent;
  let mockAuth: MockAuthService;

  beforeEach(async () => {
    mockAuth = new MockAuthService();

    await TestBed.configureTestingModule({
      imports: [
        SportsComponent,           // standalone component
        RouterTestingModule,
        HttpClientTestingModule
      ],
      providers: [
        { provide: AuthService, useValue: mockAuth },
        { provide: ApiService,  useValue: new MockApiService() }
      ]
    }).compileComponents();

    fixture = TestBed.createComponent(SportsComponent);
    component = fixture.componentInstance;
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });

  it('should show login button when not authenticated', () => {
    fixture.detectChanges();
    const btn = fixture.nativeElement.querySelector('button');
    expect(btn).not.toBeNull();
    expect(btn.textContent).toContain('Login');
  });

  it('should render slideshow image before login', () => {
    fixture.detectChanges();
    const img = fixture.nativeElement.querySelector('img[alt="Slideshow Image"]');
    expect(img).not.toBeNull();
  });

  it('should display the "CHOMP GATOR STYLE" tagline before login', () => {
    fixture.detectChanges();
    const taglineEl = fixture.nativeElement.querySelector('.tracking-wide');
    expect(taglineEl).not.toBeNull();
    expect(taglineEl.textContent).toContain('CHOMP GATOR STYLE');
  });

  it('should show 5 sport cards when authenticated', () => {
    mockAuth.isAuthenticated$ = of(true);
    fixture.detectChanges();
    const cards = fixture.nativeElement.querySelectorAll('.cursor-pointer');
    expect(cards.length).toBe(5);
  });
});
