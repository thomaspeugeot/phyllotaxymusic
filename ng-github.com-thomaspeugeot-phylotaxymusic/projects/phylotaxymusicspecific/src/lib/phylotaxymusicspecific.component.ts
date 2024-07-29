import { Component, OnInit } from '@angular/core';

import * as phylotaxymusic from '../../../phylotaxymusic/src/public-api'

import { MatSliderModule } from '@angular/material/slider'
import { FormsModule } from '@angular/forms';  // Import FormsModule
import { MatGridListModule } from '@angular/material/grid-list';

import { AngularSplitModule } from 'angular-split';

import { GongsvgDiagrammingComponent } from '@vendored_components/github.com/fullstack-lang/gongsvg/ng-github.com-fullstack-lang-gongsvg/projects/gongsvgspecific/src/lib/gongsvg-diagramming/gongsvg-diagramming'
import { TreeComponent } from '@vendored_components/github.com/fullstack-lang/gongtree/ng-github.com-fullstack-lang-gongtree/projects/gongtreespecific/src/public-api'

import { CommonModule } from '@angular/common';

@Component({
  selector: 'lib-phylotaxymusicspecific',
  standalone: true,
  imports: [
    CommonModule,
    MatSliderModule,
    FormsModule,
    AngularSplitModule,
    MatGridListModule,

    GongsvgDiagrammingComponent,
    TreeComponent,
  ],
  templateUrl: './phylotaxymusicspecific.component.html',
  styleUrls: ['./phylotaxymusicspecific.component.css'],
})
export class PhylotaxymusicspecificComponent implements OnInit {


  input($event: Event) {
    console.log(this.frontRepo!.array_Parameters[0].InsideAngle)

    let parameter = this.frontRepo!.array_Parameters[0]

    this.parameterService.updateFront(parameter, this.StacksNames.Phylotaxy).subscribe(
      () => {

      }
    )
  }

  StacksNames = phylotaxymusic.StacksNames
  StackName = phylotaxymusic.StacksNames.Phylotaxy
  TreeNames = phylotaxymusic.TreeNames


  toto: number = 40.0
  rowHeight: string = "50px"

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
