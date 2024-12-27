import { ComponentFixture, TestBed } from '@angular/core/testing';

import { PhyllotaxymusicspecificComponent } from './phyllotaxymusicspecific.component';

describe('PhyllotaxymusicspecificComponent', () => {
  let component: PhyllotaxymusicspecificComponent;
  let fixture: ComponentFixture<PhyllotaxymusicspecificComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [PhyllotaxymusicspecificComponent]
    })
      .compileComponents();

    fixture = TestBed.createComponent(PhyllotaxymusicspecificComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
