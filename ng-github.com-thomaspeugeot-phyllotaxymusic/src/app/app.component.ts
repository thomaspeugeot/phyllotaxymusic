import { Component, OnInit } from '@angular/core';

import { Observable, combineLatest, timer } from 'rxjs'

// for angular & angular material
import { FormsModule } from '@angular/forms';
import { CommonModule } from '@angular/common';
import { MatRadioModule } from '@angular/material/radio';

import { AngularSplitModule } from 'angular-split';

import * as phyllotaxymusic from '../../projects/phyllotaxymusic/src/public-api'

import { PhyllotaxymusicspecificComponent } from '../../projects/phyllotaxymusicspecific/src/public-api'

@Component({
    selector: 'app-root',
    imports: [
      CommonModule,
      FormsModule,
      MatRadioModule,
        AngularSplitModule,

        PhyllotaxymusicspecificComponent
    ],
    templateUrl: './app.component.html'
})
export class AppComponent implements OnInit {

  phyllotaxymusic = 'Phyllotaxymusic'
  probe = 'Phyllotaxymusic Data/Model'
  svg_probe = 'svg probe'
  tree_probe = "tree probe"
  tone_probe = "tone probe"
  view = this.phyllotaxymusic

  views: string[] = [this.phyllotaxymusic, this.probe, this.svg_probe, this.tree_probe, this.tone_probe];

  scrollStyle = {
    'overflow- x': 'auto',
    'width': '100%',  // Ensure the div takes the full width of its parent container
  }

  StackName = "phyllotaxymusic"
  StackType = phyllotaxymusic.StackType

  StacksNames = phyllotaxymusic.StacksNames



  constructor(
  ) {

  }

  ngOnInit(): void {
  }
}
