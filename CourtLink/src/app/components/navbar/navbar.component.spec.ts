// navbar.component.spec.ts
import { ComponentFixture, TestBed } from '@angular/core/testing';
import { NavbarComponent } from './navbar.component';
import { AuthService } from '@auth0/auth0-angular';
import { of, Observable } from 'rxjs';
import { RouterTestingModule } from '@angular/router/testing';
import { MatMenuModule } from '@angular/material/menu';
import { MatIconModule } from '@angular/material/icon';
import { MatButtonModule } from '@angular/material/button';

class MockAuth {
  public user$: Observable<any> = of(null);
  loginWithRedirect(): void {}
  logout(_opts?: any): void {}
}

describe('NavbarComponent (new functionality)', () => {
  let fixture: ComponentFixture<NavbarComponent>;
  let component: NavbarComponent;
  let mockAuth: MockAuth;
  let el: HTMLElement;

  beforeEach(async () => {
    mockAuth = new MockAuth();

    await TestBed.configureTestingModule({
      imports: [
        NavbarComponent,    // standalone :contentReference[oaicite:0]{index=0}&#8203;:contentReference[oaicite:1]{index=1}
        RouterTestingModule,
        MatMenuModule,
        MatIconModule,
        MatButtonModule
      ],
      providers: [
        { provide: AuthService, useValue: mockAuth }
      ]
    }).compileComponents();

    fixture = TestBed.createComponent(NavbarComponent);
    component = fixture.componentInstance;
    el = fixture.nativeElement;
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });

  it('shows account_circle icon when NOT logged in', () => {
    fixture.detectChanges();
    const icons = Array.from(el.querySelectorAll('mat-icon'));
    const hasAccount = icons.some(i => i.textContent?.trim() === 'account_circle');
    expect(hasAccount).toBeTrue();
  });

  it('shows menu icon when logged in', () => {
    mockAuth.user$ = of({ name: 'Test User' });
    fixture.detectChanges();
    const icons = Array.from(el.querySelectorAll('mat-icon'));
    const hasMenu = icons.some(i => i.textContent?.trim() === 'menu');
    expect(hasMenu).toBeTrue();
  });

  // ─── EXTRA TESTS BELOW ─────────────────────────────────────────────────────

  it('displays brand link to home with text UFCourtLink', () => {
    fixture.detectChanges();
    const brandLink = el.querySelector('a[href="/"]');
    expect(brandLink).not.toBeNull();
    expect(brandLink!.textContent).toContain('UFCourtLink');
  });

  it('renders a search input with placeholder "Search..."', () => {
    fixture.detectChanges();
    const input = el.querySelector('input[type="text"]');
    expect(input).not.toBeNull();
    expect(input!.getAttribute('placeholder')).toBe('Search...');
  });

  it('calls loginWithRedirect() when login icon is clicked', () => {
    spyOn(mockAuth, 'loginWithRedirect');
    fixture.detectChanges();
    // find the login button (only one mat-icon-button when logged out)
    const buttons = Array.from(el.querySelectorAll('button[mat-icon-button]'));
    const loginBtn = buttons.find(b => b.textContent?.includes('account_circle'));
    expect(loginBtn).not.toBeUndefined();
    (loginBtn as HTMLElement).click();
    expect(mockAuth.loginWithRedirect).toHaveBeenCalled();
  });
});
