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

import { GongtoneComponent } from '@vendored_components/github.com/fullstack-lang/gongtone/ng-github.com-fullstack-lang-gongtone/projects/gongtonespecific/src/lib/gongtone/gongtone.component'

import { ButtonSpecificComponent } from '@vendored_components/github.com/fullstack-lang/gong/lib/button/ng-github.com-fullstack-lang-gong-lib-button/projects/buttonspecific/src/lib/button-specific/button-specific.component' 
import { SliderSpecificComponent } from '@vendored_components/github.com/fullstack-lang/gong/lib/slider/ng-github.com-fullstack-lang-gong-lib-slider/projects/sliderspecific/src/lib/slider-specific/slider-specific.component' 
import { TreeSpecificComponent } from '@vendored_components/github.com/fullstack-lang/gong/lib/tree/ng-github.com-fullstack-lang-gong-lib-tree/projects/treespecific/src/lib/tree-specific/tree-specific.component'
import { SvgSpecificComponent } from '@vendored_components/github.com/fullstack-lang/gong/lib/svg/ng-github.com-fullstack-lang-gong-lib-svg/projects/svgspecific/src/lib/svg-specific/svg-specific.component'

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

        
        
        TreeSpecificComponent,
        GongtoneComponent,
        CursorspecificComponent,

        SliderSpecificComponent,
        ButtonSpecificComponent,
        SvgSpecificComponent
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
