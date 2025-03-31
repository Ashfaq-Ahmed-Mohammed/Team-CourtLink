import { ComponentFixture, TestBed, fakeAsync, tick } from '@angular/core/testing';
import { AdminCourtsComponent } from './admin-courts.component';
import { HttpClientTestingModule, HttpTestingController } from '@angular/common/http/testing';
import { MatTableModule } from '@angular/material/table';
import { MatButtonModule } from '@angular/material/button';
import { MatFormFieldModule } from '@angular/material/form-field';
import { MatInputModule } from '@angular/material/input';
import { BrowserAnimationsModule } from '@angular/platform-browser/animations';

describe('AdminCourtsComponent', () => {
  let component: AdminCourtsComponent;
  let fixture: ComponentFixture<AdminCourtsComponent>;
  let httpMock: HttpTestingController;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [
        HttpClientTestingModule,
        MatTableModule,
        MatButtonModule,
        MatFormFieldModule,
        MatInputModule,
        BrowserAnimationsModule,
        AdminCourtsComponent,
      ],
    }).compileComponents();

    fixture = TestBed.createComponent(AdminCourtsComponent);
    component = fixture.componentInstance;
    httpMock = TestBed.inject(HttpTestingController);
    fixture.detectChanges();
  });

  it('should create the component', () => {
    expect(component).toBeTruthy();
  });

  it('should fetch and display courts', fakeAsync(() => {
    const mockCourts = [
      { court_name: 'Court A', sport_name: 'Tennis' },
      { court_name: 'Court B', sport_name: 'Basketball' },
    ];
    const req = httpMock.expectOne('http://localhost:8080/ListCourts');
    req.flush(mockCourts);

    tick();
    fixture.detectChanges();

    expect(component.courts.data.length).toBe(2);
    expect(component.courts.data[0].court_name).toBe('Court A');
  }));

  it('should filter courts when search is applied', () => {
    component.courts.data = [
      { court_name: 'Court A', sport_name: 'Tennis' },
      { court_name: 'Court B', sport_name: 'Basketball' },
    ];

    const inputEvent = new Event('input');
    const inputElement = { target: { value: 'tennis' } } as unknown as Event;
    component.applyFilter(inputElement);
    expect(component.courts.filter).toBe('tennis');
  });

  it('should render the correct number of displayed columns', () => {
    expect(component.displayedColumns).toEqual(['sno', 'court', 'sport', 'action']);
  });

  it('should handle empty court list without crashing', fakeAsync(() => {
    const req = httpMock.expectOne('http://localhost:8080/ListCourts');
    req.flush([]);

    tick();
    fixture.detectChanges();

    expect(component.courts.data.length).toBe(0);
  }));
});
