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
  styles: ``
})
export class SlidersspecificComponent {
input($event: Event) {
throw new Error('Method not implemented.');
}

  rowHeight: string = "50px"

  StacksNames = sliders.StacksNames;
  public frontRepo?: sliders.FrontRepo;
  splitAreaSize = 0

  constructor(
    private frontRepoService: sliders.FrontRepoService,
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

        this.splitAreaSize = 100.0 / this.frontRepo.array_Groups.length
      }
    }
    )
  }

}
