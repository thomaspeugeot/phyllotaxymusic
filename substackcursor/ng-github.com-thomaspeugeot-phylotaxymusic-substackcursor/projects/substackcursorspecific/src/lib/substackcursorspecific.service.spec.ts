import { TestBed } from '@angular/core/testing';

import { SubstackcursorspecificService } from './substackcursorspecific.service';

describe('SubstackcursorspecificService', () => {
  let service: SubstackcursorspecificService;

  beforeEach(() => {
    TestBed.configureTestingModule({});
    service = TestBed.inject(SubstackcursorspecificService);
  });

  it('should be created', () => {
    expect(service).toBeTruthy();
  });
});
