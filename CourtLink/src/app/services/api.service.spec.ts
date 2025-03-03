/* import { TestBed } from '@angular/core/testing';
import { ApiService } from './api.service';
import { provideHttpClientTesting } from '@angular/common/http/testing';

describe('ApiService', () => {
  let service: ApiService;

  beforeEach(() => {
    TestBed.configureTestingModule({
      providers: [
        ApiService,
        ...provideHttpClientTesting()  // Spread the providers array
      ]
    });
    service = TestBed.inject(ApiService);
  });

  it('should be created', () => {
    if (service === null || service === undefined) {
      fail('ApiService was not created.');
    }
  });
});
 */