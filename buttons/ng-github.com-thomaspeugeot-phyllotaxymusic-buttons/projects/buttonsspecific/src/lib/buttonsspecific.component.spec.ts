import { ComponentFixture, TestBed } from '@angular/core/testing';

import { ButtonsspecificComponent } from './buttonsspecific.component';

describe('ButtonsspecificComponent', () => {
  let component: ButtonsspecificComponent;
  let fixture: ComponentFixture<ButtonsspecificComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [ButtonsspecificComponent]
    })
    .compileComponents();
    
    fixture = TestBed.createComponent(ButtonsspecificComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
