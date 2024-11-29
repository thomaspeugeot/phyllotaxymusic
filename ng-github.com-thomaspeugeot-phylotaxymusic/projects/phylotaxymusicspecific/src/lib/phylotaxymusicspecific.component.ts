import { Component, model, OnInit } from '@angular/core';

import * as phylotaxymusic from '../../../phylotaxymusic/src/public-api'

import { MatSliderModule } from '@angular/material/slider'
import { FormsModule } from '@angular/forms';  // Import FormsModule
import { MatGridListModule } from '@angular/material/grid-list';

import { MatRadioChange, MatRadioModule } from '@angular/material/radio';
import { MatFormFieldModule } from '@angular/material/form-field';
import { MatCardModule } from '@angular/material/card';
import { MatCheckboxModule } from '@angular/material/checkbox';

import { AngularSplitModule } from 'angular-split';

import { GongsvgDiagrammingComponent } from '@vendored_components/github.com/fullstack-lang/gongsvg/ng-github.com-fullstack-lang-gongsvg/projects/gongsvgspecific/src/lib/gongsvg-diagramming/gongsvg-diagramming'
import { TreeComponent } from '@vendored_components/github.com/fullstack-lang/gongtree/ng-github.com-fullstack-lang-gongtree/projects/gongtreespecific/src/public-api'
import { GongtoneComponent } from '@vendored_components/github.com/fullstack-lang/gongtone/ng-github.com-fullstack-lang-gongtone/projects/gongtonespecific/src/lib/gongtone/gongtone.component'


import { CommonModule } from '@angular/common';

@Component({
  selector: 'lib-phylotaxymusicspecific',
  standalone: true,
  imports: [
    CommonModule,
    MatSliderModule,
    MatRadioModule,

    MatCardModule,
    MatCheckboxModule,

    FormsModule,
    MatFormFieldModule,

    AngularSplitModule,
    MatGridListModule,

    GongsvgDiagrammingComponent,
    TreeComponent,
    GongtoneComponent,

  ],
  templateUrl: './phylotaxymusicspecific.component.html',
  styleUrls: ['./phylotaxymusicspecific.component.css'],
})
export class PhylotaxymusicspecificComponent implements OnInit {
  onChange(i: number) {

    let parameter = this.frontRepo!.array_Parameters[0]
    let noteInfo = parameter.NoteInfos[i]

    this.noteInfoService.updateFront(noteInfo, this.StacksNames.Phylotaxy).subscribe(
      () => {

        // in order to provoke backend rework
        let event2: Event = new Event('input');
        this.input(event2)
      }
    )
  }

  readonly checked = model(false);
  readonly indeterminate = model(false);

  inputMatRadio($event: MatRadioChange) {
    let event2: Event = new Event('input');
    this.input(event2)
  }

  input($event: Event) {
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
    private noteInfoService: phylotaxymusic.NoteInfoService,
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

  getNoteInfoAtOffset(i: number): phylotaxymusic.NoteInfo | undefined {
    const noteInfos = this.frontRepo!.array_Parameters[0].NoteInfos;
    const bruteOffsetIndex = i - this.frontRepo!.array_Parameters[0].ActualBeatsTemporalShift + noteInfos.length
    const offsetIndex = bruteOffsetIndex %
      noteInfos.length;
    // console.log(i, this.frontRepo!.array_Parameters[0].ActualBeatsTemporalShift, bruteOffsetIndex, offsetIndex)
    return noteInfos[offsetIndex];
  }
}
