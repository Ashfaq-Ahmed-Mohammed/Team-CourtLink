import { ComponentFixture, TestBed } from '@angular/core/testing';
import { AdminComponent } from './admin.component';
import { RouterTestingModule } from '@angular/router/testing';
import { By } from '@angular/platform-browser';
import { CommonModule } from '@angular/common';

describe('AdminComponent', () => {
  let component: AdminComponent;
  let fixture: ComponentFixture<AdminComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [AdminComponent, RouterTestingModule, CommonModule]
    }).compileComponents();

    fixture = TestBed.createComponent(AdminComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create the AdminComponent', () => {
    expect(component).toBeTruthy();
  });

  it('should display "ADMIN" in the navbar', () => {
    const adminText = fixture.debugElement.query(By.css('nav div')).nativeElement.textContent;
    expect(adminText).toContain('ADMIN');
  });

  it('should display "WELCOME ADMIN, EDIT AWAY"', () => {
    const heading = fixture.debugElement.query(By.css('h2')).nativeElement.textContent;
    expect(heading).toContain('WELCOME ADMIN, EDIT AWAY');
  });

  it('should display View Sports tile with correct text and image', () => {
    const sportsTile = fixture.debugElement.query(By.css('a[routerLink="/admin/sports"]'));
    const label = sportsTile.nativeElement.querySelector('p').textContent;
    const img = sportsTile.nativeElement.querySelector('img');
    expect(label).toContain('View Sports');
    expect(img).toBeTruthy();
  });

  it('should display View Courts tile with correct text and image', () => {
    const courtsTile = fixture.debugElement.query(By.css('a[routerLink="/admin/courts"]'));
    const label = courtsTile.nativeElement.querySelector('p').textContent;
    const img = courtsTile.nativeElement.querySelector('img');
    expect(label).toContain('View Courts');
    expect(img).toBeTruthy();
  });
});
