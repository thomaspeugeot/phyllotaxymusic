import { Component } from '@angular/core';

import { AngularSplitModule } from 'angular-split';

import * as buttons from '../../../buttons/src/public-api'

import { MatSliderModule } from '@angular/material/slider'
import { FormsModule } from '@angular/forms';  // Import FormsModule
import { MatGridListModule } from '@angular/material/grid-list';

import { MatRadioChange, MatRadioModule } from '@angular/material/radio';
import { MatFormFieldModule } from '@angular/material/form-field';
import { MatCardModule } from '@angular/material/card';
import { MatButtonModule } from '@angular/material/button';
import { MatIconModule } from '@angular/material/icon'

@Component({
  selector: 'lib-buttonsspecific',
  standalone: true,
  imports: [
    AngularSplitModule,

    MatSliderModule,
    MatRadioModule,
    MatCardModule,
    MatButtonModule,
    FormsModule,
    MatFormFieldModule,
    MatIconModule,
    MatGridListModule,
  ],
  templateUrl: 'buttonsspecific.component.html',
  styleUrls: ['./buttonsspecific.component.css'],
})
export class ButtonsspecificComponent {


  StacksNames = buttons.StacksNames;
  public frontRepo?: buttons.FrontRepo;
  splitAreaSize = 0

  layout: buttons.Layout | undefined

  constructor(
    private frontRepoService: buttons.FrontRepoService,
    private buttonService: buttons.ButtonService,
  ) { }

  formatLabel(value: number): string {
    if (value >= 1000) {
      return Math.round(value / 1000) + 'k';
    }

    return `${value}`;
  }

  ngOnInit(): void {
    console.log("ngOnInit");

    this.frontRepoService.connectToWebSocket(this.StacksNames.ButtonsStackName).subscribe({
      next: (frontRepo) => {
        this.frontRepo = frontRepo;

        for (let layout_ of this.frontRepo.array_Layouts) {
          this.layout = layout_
        }

        this.splitAreaSize = 100.0 / this.frontRepo.array_Groups.length
      }
    }
    )
  }

  onClick(button: buttons.Button) {
    this.buttonService.updateFront(button, this.StacksNames.ButtonsStackName).subscribe(
      () => {
        console.log("checkbox updated")
      }
    )
  }
}
