import { TestBed } from '@angular/core/testing';

import { PhylotaxymusicspecificService } from './phylotaxymusicspecific.service';

describe('PhylotaxymusicspecificService', () => {
  let service: PhylotaxymusicspecificService;

  beforeEach(() => {
    TestBed.configureTestingModule({});
    service = TestBed.inject(PhylotaxymusicspecificService);
  });

  it('should be created', () => {
    expect(service).toBeTruthy();
  });
});
