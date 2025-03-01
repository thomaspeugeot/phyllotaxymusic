import { Component, Input, model, OnInit } from '@angular/core';

import * as phyllotaxymusic from '../../../phyllotaxymusic/src/public-api'

import { MatSliderModule } from '@angular/material/slider'
import { FormsModule } from '@angular/forms';  // Import FormsModule
import { MatGridListModule } from '@angular/material/grid-list';

import { MatRadioChange, MatRadioModule } from '@angular/material/radio';
import { MatFormFieldModule } from '@angular/material/form-field';
import { MatCardModule } from '@angular/material/card';
import { MatCheckboxModule } from '@angular/material/checkbox';
import { MatIconModule } from '@angular/material/icon'

import { AngularSplitModule } from 'angular-split';

import { GongsvgDiagrammingComponent } from '@vendored_components/github.com/fullstack-lang/gongsvg/ng-github.com-fullstack-lang-gongsvg/projects/gongsvgspecific/src/lib/gongsvg-diagramming/gongsvg-diagramming'
import { TreeComponent } from '@vendored_components/github.com/fullstack-lang/gongtree/ng-github.com-fullstack-lang-gongtree/projects/gongtreespecific/src/public-api'
import { GongtoneComponent } from '@vendored_components/github.com/fullstack-lang/gongtone/ng-github.com-fullstack-lang-gongtone/projects/gongtonespecific/src/lib/gongtone/gongtone.component'

import { ButtonSpecificComponent } from '@vendored_components/github.com/fullstack-lang/gong/lib/button/ng-github.com-fullstack-lang-gong-lib-button/projects/buttonspecific/src/lib/button-specific/button-specific.component' 
import { SliderSpecificComponent } from '@vendored_components/github.com/fullstack-lang/gong/lib/slider/ng-github.com-fullstack-lang-gong-lib-slider/projects/sliderspecific/src/lib/slider-specific/slider-specific.component' 


import { CursorspecificComponent } from '../../../../../cursor/ng-github.com-thomaspeugeot-phyllotaxymusic-cursor/projects/cursorspecific/src/public-api'

import { CommonModule } from '@angular/common';

@Component({
    selector: 'lib-phyllotaxymusicspecific',
    imports: [
        CommonModule,
        MatSliderModule,
        MatRadioModule,
        MatCardModule,
        MatCheckboxModule,
        FormsModule,
        MatFormFieldModule,
        MatIconModule,
        AngularSplitModule,
        MatGridListModule,
        GongsvgDiagrammingComponent,
        TreeComponent,
        GongtoneComponent,
        CursorspecificComponent,
        SliderSpecificComponent,
        ButtonSpecificComponent
    ],
    templateUrl: './phyllotaxymusicspecific.component.html',
    styleUrls: ['./phyllotaxymusicspecific.component.css']
})
export class PhyllotaxymusicspecificComponent implements OnInit {


  private socket: WebSocket | undefined

  readonly checked = model(false);
  readonly indeterminate = model(false);


  StacksNames = phyllotaxymusic.StacksNames
  StackName = phyllotaxymusic.StacksNames.Phylotaxy
  TreeNames = phyllotaxymusic.TreeNames


  toto: number = 40.0
  rowHeight: string = "50px"

  public frontRepo?: phyllotaxymusic.FrontRepo

  constructor(
    private frontRepoService: phyllotaxymusic.FrontRepoService,
  ) {

  }

  ngOnInit(): void {

  }
}
