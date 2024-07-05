import { Component, OnInit } from '@angular/core';

import * as phylotaxymusic from '../../../phylotaxymusic/src/public-api'

import { MatSliderModule } from '@angular/material/slider'
import { FormsModule } from '@angular/forms';  // Import FormsModule

import { GongsvgDiagrammingComponent } from '@vendored_components/github.com/fullstack-lang/gongsvg/ng-github.com-fullstack-lang-gongsvg/projects/gongsvgspecific/src/lib/gongsvg-diagramming/gongsvg-diagramming'
import { CommonModule } from '@angular/common';

@Component({
  selector: 'lib-phylotaxymusicspecific',
  standalone: true,
  imports: [
    CommonModule,
    MatSliderModule,
    FormsModule,
    GongsvgDiagrammingComponent
  ],
  templateUrl: './phylotaxymusicspecific.component.html',
  styles: `
  mat-slider {
  width: 800px;
  }`
})
export class PhylotaxymusicspecificComponent implements OnInit {


  input($event: Event) {
    console.log(this.frontRepo!.array_Parameters[0].Angle)

    let parameter = this.frontRepo!.array_Parameters[0]

    this.parameterService.updateFront(parameter, this.StacksNames.Phylotaxy).subscribe(
      () => {

      }
    )
  }

  StacksNames = phylotaxymusic.StacksNames
  StackName = phylotaxymusic.StacksNames.Phylotaxy

  toto: number = 40.0

  public frontRepo?: phylotaxymusic.FrontRepo

  constructor(
    private frontRepoService: phylotaxymusic.FrontRepoService,

    private parameterService: phylotaxymusic.ParameterService,
  ) {

  }

  ngOnInit(): void {

    console.log("ngOnInit")

    this.frontRepoService.connectToWebSocket(this.StacksNames.Phylotaxy).subscribe(
      gongtablesFrontRepo => {
        this.frontRepo = gongtablesFrontRepo
      }
    )
  }

  formatLabel(value: number): string {
    if (value >= 1000) {
      return Math.round(value / 1000) + 'k';
    }

    return `${value}`;
  }
}
