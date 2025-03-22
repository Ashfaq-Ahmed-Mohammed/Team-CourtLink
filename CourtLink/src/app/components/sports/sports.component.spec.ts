import { ComponentFixture, TestBed } from '@angular/core/testing';
import { SportsComponent } from './sports.component';
import { ApiService } from './../../services/api.service';  
import { Router } from '@angular/router';
import { HttpClientModule } from '@angular/common/http';
import { of, throwError } from 'rxjs';
import { RouterTestingModule } from '@angular/router/testing';


class MockApiService {
  getCourts(sport: string) {
    if (sport === 'Basketball') {
      return of([]);  
    }
    return throwError('Error fetching courts');  
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
        HttpClientModule,    
        RouterTestingModule, 
        SportsComponent      
      ],
      providers: [
        { provide: ApiService, useClass: MockApiService },  
      ],
    }).compileComponents();
  });

  beforeEach(() => {
    fixture = TestBed.createComponent(SportsComponent);
    component = fixture.componentInstance;
    router = TestBed.inject(Router);  
    apiService = TestBed.inject(ApiService);  
    fixture.detectChanges();
  });

  it('should create the component', () => {
    expect(component).toBeTruthy();
  });


  it('should call selectSport and navigate on success', async () => {
    const navigateSpy = spyOn(router, 'navigate');
    const sportName = 'Basketball';


    await component.selectSport(sportName);


    expect(navigateSpy).toHaveBeenCalledWith(['/courts', sportName.toLowerCase()]);
  });

  it('should handle error and not navigate on failure', async () => {
    const navigateSpy = spyOn(router, 'navigate');
    const sportName = 'Soccer';


    await component.selectSport(sportName);


    expect(navigateSpy).not.toHaveBeenCalled();
  });
}); 