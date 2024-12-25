import { ComponentFixture, TestBed } from '@angular/core/testing';

import { SubstackcursorspecificComponent } from './substackcursorspecific.component';

describe('SubstackcursorspecificComponent', () => {
  let component: SubstackcursorspecificComponent;
  let fixture: ComponentFixture<SubstackcursorspecificComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [SubstackcursorspecificComponent]
    })
    .compileComponents();
    
    fixture = TestBed.createComponent(SubstackcursorspecificComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
