import { ComponentFixture, TestBed } from '@angular/core/testing';

import { SlidersspecificComponent } from './slidersspecific.component';

describe('SlidersspecificComponent', () => {
  let component: SlidersspecificComponent;
  let fixture: ComponentFixture<SlidersspecificComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [SlidersspecificComponent]
    })
    .compileComponents();
    
    fixture = TestBed.createComponent(SlidersspecificComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
