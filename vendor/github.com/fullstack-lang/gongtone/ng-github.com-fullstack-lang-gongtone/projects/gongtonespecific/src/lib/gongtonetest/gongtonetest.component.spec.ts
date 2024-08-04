import { ComponentFixture, TestBed } from '@angular/core/testing';

import { GongtonetestComponent } from './gongtonetest.component';

describe('GongtonetestComponent', () => {
  let component: GongtonetestComponent;
  let fixture: ComponentFixture<GongtonetestComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [GongtonetestComponent]
    })
    .compileComponents();
    
    fixture = TestBed.createComponent(GongtonetestComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
