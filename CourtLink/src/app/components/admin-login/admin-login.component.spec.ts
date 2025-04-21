import { ComponentFixture, TestBed } from '@angular/core/testing';
import { AdminLoginComponent } from './admin-login.component';
import { HttpClientTestingModule } from '@angular/common/http/testing';
import { RouterTestingModule } from '@angular/router/testing';
import { Router } from '@angular/router';
import { of, throwError } from 'rxjs';
import { provideAnimations } from '@angular/platform-browser/animations'; // ✅ import this

describe('AdminLoginComponent', () => {
  let component: AdminLoginComponent;
  let fixture: ComponentFixture<AdminLoginComponent>;
  let router: Router;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [
        AdminLoginComponent,
        HttpClientTestingModule,
        RouterTestingModule,
      ],
      providers: [
        provideAnimations(), // ✅ Add this line
      ]
    }).compileComponents();

    fixture = TestBed.createComponent(AdminLoginComponent);
    component = fixture.componentInstance;
    router = TestBed.inject(Router);
    fixture.detectChanges();
  });

  it('should create the component', () => {
    expect(component).toBeTruthy();
  });

  it('should navigate to /admin on successful login', () => {
    spyOn(component['http'], 'post').and.returnValue(of({ message: 'Login successful' }));
    const navigateSpy = spyOn(router, 'navigate');

    component.username = 'Admin123';
    component.password = 'Password123';
    component.login();

    expect(navigateSpy).toHaveBeenCalledWith(['/admin-portal']);
  });

  it('should display error on failed login', () => {
    spyOn(component['http'], 'post').and.returnValue(throwError(() => ({ error: 'Invalid login' })));

    component.username = 'wrong';
    component.password = 'wrong';
    component.login();

    expect(component.errorMessage).toBe('Invalid login');
  });
});
