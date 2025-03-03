import { ComponentFixture, TestBed } from '@angular/core/testing';
import { CourtsComponent } from './courts.component';
import { ActivatedRoute } from '@angular/router';
import { NO_ERRORS_SCHEMA } from '@angular/core';
import { ApiService } from '../../services/api.service';

// Dummy ApiService that does nothing.
class DummyApiService {}

describe('CourtsComponent (Standalone)', () => {
  let fixture: ComponentFixture<CourtsComponent>;
  let component: CourtsComponent;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [CourtsComponent],
      providers: [
        // Provide a dummy ActivatedRoute with a proper paramMap.get() method.
        { 
          provide: ActivatedRoute, 
          useValue: { snapshot: { paramMap: { get: (key: string) => null } } } 
        },
        // Replace ApiService with a dummy version so that HttpClient is not required.
        { provide: ApiService, useClass: DummyApiService }
      ],
      schemas: [NO_ERRORS_SCHEMA]
    }).compileComponents();
  });

  beforeEach(() => {
    fixture = TestBed.createComponent(CourtsComponent);
    component = fixture.componentInstance;
    // Update the WritableSignals: simulate no available courts and a dummy sport.
    component.courts.set([]);
    component.selectedSport.set('basketball');
    fixture.detectChanges();
  });

  // Test 1: Ensure the Courts component renders.
  it('should render the courts component', () => {
    if (!component) {
      fail('CourtsComponent was not created.');
    }
    const container: HTMLElement = fixture.nativeElement.querySelector('div');
    if (!container) {
      fail('Courts component container not rendered.');
    }
  });

  // Test 2: Verify that "No courts available or loading..." is displayed when there are no courts.
  it('should display "No courts available or loading..." when there are no courts', () => {
    const compiled = fixture.nativeElement as HTMLElement;
    const expectedText = 'No courts available or loading...';
    if (!compiled.textContent || compiled.textContent.indexOf(expectedText) === -1) {
      fail('Expected text "' + expectedText + '" not found.');
    }
  });
});
