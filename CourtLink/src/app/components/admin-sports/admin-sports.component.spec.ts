import { ComponentFixture, TestBed, fakeAsync, tick } from '@angular/core/testing';
import { AdminSportsComponent } from './admin-sports.component';
import { HttpClientTestingModule, HttpTestingController } from '@angular/common/http/testing';
import { MatDialog, MatDialogModule } from '@angular/material/dialog';
import { of } from 'rxjs';
import { AddSportDialogComponent } from '../add-sport-dialog/add-sport-dialog.component';
import { BrowserAnimationsModule } from '@angular/platform-browser/animations';

describe('AdminSportsComponent', () => {
  let component: AdminSportsComponent;
  let fixture: ComponentFixture<AdminSportsComponent>;
  let httpMock: HttpTestingController;
  let dialogSpy: jasmine.Spy;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [
        HttpClientTestingModule,
        MatDialogModule,
        BrowserAnimationsModule,
        AdminSportsComponent,
        AddSportDialogComponent
      ],
    }).compileComponents();

    fixture = TestBed.createComponent(AdminSportsComponent);
    component = fixture.componentInstance;
    httpMock = TestBed.inject(HttpTestingController);
    fixture.detectChanges();
  });

  it('should create the component', () => {
    expect(component).toBeTruthy();
  });

  it('should fetch and display sports', fakeAsync(() => {
    const mockSports = ['Basketball', 'Tennis'];
    const req = httpMock.expectOne('http://localhost:8080/ListSports');
    req.flush(mockSports);

    tick();
    fixture.detectChanges();

    expect(component.sports.length).toBe(2);
    expect(component.sports[0].name).toBe('Basketball');
  }));

  it('should not add sport if dialog is cancelled', () => {
    const dialog = TestBed.inject(MatDialog);
    component.sports = [];

    spyOn(dialog, 'open').and.returnValue({
      afterClosed: () => of(null)
    } as any);

    component.openAddSportDialog();

    expect(component.sports.length).toBe(0);
  });
});