import { TestBed } from '@angular/core/testing';

import { SlidersspecificService } from './slidersspecific.service';

describe('SlidersspecificService', () => {
  let service: SlidersspecificService;

  beforeEach(() => {
    TestBed.configureTestingModule({});
    service = TestBed.inject(SlidersspecificService);
  });

  it('should be created', () => {
    expect(service).toBeTruthy();
  });
});
