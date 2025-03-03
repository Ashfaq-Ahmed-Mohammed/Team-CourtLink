import { ComponentFixture, TestBed } from '@angular/core/testing';
import { SportsComponent } from './sports.component';
import { ApiService } from './../../services/api.service';  // Import ApiService
import { Router } from '@angular/router';
import { HttpClientModule } from '@angular/common/http';
import { of, throwError } from 'rxjs';
import { RouterTestingModule } from '@angular/router/testing';

// Mock ApiService for testing
class MockApiService {
  getCourts(sport: string) {
    if (sport === 'Basketball') {
      return of([]);  // Simulate success response for Basketball
    }
    return throwError('Error fetching courts');  // Simulate error for other sports
  }
}

describe('SportsComponent', () => {
  let component: SportsComponent;
  let fixture: ComponentFixture<SportsComponent>;
  let router: Router;
  let apiService: ApiService;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [
        HttpClientModule,    // Import HttpClientModule
        RouterTestingModule, // Import RouterTestingModule
        SportsComponent      // Import the standalone component
      ],
      providers: [
        { provide: ApiService, useClass: MockApiService },  // Use MockApiService
      ],
    }).compileComponents();
  });

  beforeEach(() => {
    fixture = TestBed.createComponent(SportsComponent);
    component = fixture.componentInstance;
    router = TestBed.inject(Router);  // Inject router
    apiService = TestBed.inject(ApiService);  // Inject ApiService
    fixture.detectChanges();
  });

  it('should create the component', () => {
    expect(component).toBeTruthy();
  });

  it('should display sports names correctly', () => {
    const compiled = fixture.nativeElement as HTMLElement;
    const sportNames = compiled.querySelectorAll('.sport-name');
    
    // Check that each sport name is displayed in the template
    expect(sportNames.length).toBe(component.sports.length);
    expect(sportNames[0].textContent).toContain('Basketball');
    expect(sportNames[1].textContent).toContain('Soccer');
    expect(sportNames[2].textContent).toContain('Tennis');
    expect(sportNames[3].textContent).toContain('Badminton');
    expect(sportNames[4].textContent).toContain('Cricket');
  });

  it('should call selectSport and navigate on success', async () => {
    const navigateSpy = spyOn(router, 'navigate');
    const sportName = 'Basketball';

    // Call the selectSport method
    await component.selectSport(sportName);

    // Expect the API service to be called and navigation to happen
    expect(navigateSpy).toHaveBeenCalledWith(['/courts', sportName.toLowerCase()]);
  });

  it('should handle error and not navigate on failure', async () => {
    const navigateSpy = spyOn(router, 'navigate');
    const sportName = 'Soccer';

    // Call the selectSport method
    await component.selectSport(sportName);

    // Expect that the navigation was not called due to the error
    expect(navigateSpy).not.toHaveBeenCalled();
  });
});
