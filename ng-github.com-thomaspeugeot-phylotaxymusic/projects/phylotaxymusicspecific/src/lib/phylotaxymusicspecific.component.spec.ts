import { ComponentFixture, TestBed } from '@angular/core/testing';

import { PhylotaxymusicspecificComponent } from './phylotaxymusicspecific.component';

describe('PhylotaxymusicspecificComponent', () => {
  let component: PhylotaxymusicspecificComponent;
  let fixture: ComponentFixture<PhylotaxymusicspecificComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [PhylotaxymusicspecificComponent]
    })
    .compileComponents();
    
    fixture = TestBed.createComponent(PhylotaxymusicspecificComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
