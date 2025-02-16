import { Component } from '@angular/core';

import { AngularSplitModule } from 'angular-split';

import * as sliders from '../../../sliders/src/public-api'

import { MatSliderModule } from '@angular/material/slider'
import { FormsModule } from '@angular/forms';  // Import FormsModule
import { MatGridListModule } from '@angular/material/grid-list';

import { MatRadioModule } from '@angular/material/radio';
import { MatFormFieldModule } from '@angular/material/form-field';
import { MatCardModule } from '@angular/material/card';
import { MatCheckboxModule } from '@angular/material/checkbox';
import { MatIconModule } from '@angular/material/icon'

@Component({
  selector: 'lib-slidersspecific',
  standalone: true,
  imports: [
    AngularSplitModule,

    MatSliderModule,
    MatRadioModule,
    MatCardModule,
    MatCheckboxModule,
    FormsModule,
    MatFormFieldModule,
    MatIconModule,
    MatGridListModule,

  ],
  templateUrl: './slidersspecific.component.html',
  styleUrls: [
    './slidersspecific.component.css'
  ]
})
export class SlidersspecificComponent {


  rowHeight: string = "30px"

  StacksNames = sliders.StacksNames;
  public frontRepo?: sliders.FrontRepo;
  splitAreaSize = 0

  layout : sliders.Layout | undefined

  constructor(
    private frontRepoService: sliders.FrontRepoService,
    private sliderService: sliders.SliderService,
  ) { }

  formatLabel(value: number): string {
    if (value >= 1000) {
      return Math.round(value / 1000) + 'k';
    }

    return `${value}`;
  }

  ngOnInit(): void {
    console.log("ngOnInit");

    this.frontRepoService.connectToWebSocket(this.StacksNames.SliderStackName).subscribe({
      next: (frontRepo) => {
        this.frontRepo = frontRepo;

        console.log("hello")

        for (let layout_ of this.frontRepo.array_Layouts) {
          this.layout = layout_
        }

        this.splitAreaSize = 100.0 / this.frontRepo.array_Groups.length
      }
    }
    )
  }

  input($event: Event, slider: sliders.Slider) {
    this.sliderService.updateFront(slider, this.StacksNames.SliderStackName).subscribe(
      () => {
        console.log("slider updated")
      }
    )
  }

}
