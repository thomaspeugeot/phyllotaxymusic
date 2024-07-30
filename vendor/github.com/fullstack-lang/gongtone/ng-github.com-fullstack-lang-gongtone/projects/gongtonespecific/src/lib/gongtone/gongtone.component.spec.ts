import { ComponentFixture, TestBed } from '@angular/core/testing';

import { GongtoneComponent } from './gongtone.component';

describe('GongtoneComponent', () => {
  let component: GongtoneComponent;
  let fixture: ComponentFixture<GongtoneComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [GongtoneComponent]
    })
    .compileComponents();
    
    fixture = TestBed.createComponent(GongtoneComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
