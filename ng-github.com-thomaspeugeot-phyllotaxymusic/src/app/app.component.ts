import { Component, OnInit } from '@angular/core';

import { Observable, combineLatest, timer } from 'rxjs'

// for angular & angular material
import { FormsModule } from '@angular/forms';
import { CommonModule } from '@angular/common';
import { MatRadioModule } from '@angular/material/radio';
import { MatButtonModule } from '@angular/material/button';
import { MatIconModule } from '@angular/material/icon';

import { AngularSplitModule } from 'angular-split';

import * as phyllotaxymusic from '../../projects/phyllotaxymusic/src/public-api'

import { PhyllotaxymusicspecificComponent } from '../../projects/phyllotaxymusicspecific/src/public-api'

import * as gongtone from '@vendored_components/github.com/fullstack-lang/gongtone/ng-github.com-fullstack-lang-gongtone/projects/gongtone/src/public-api';

import { TreeSpecificComponent } from '@vendored_components/github.com/fullstack-lang/gong/lib/tree/ng-github.com-fullstack-lang-gong-lib-tree/projects/treespecific/src/lib/tree-specific/tree-specific.component'
import { TableSpecificComponent } from '@vendored_components/github.com/fullstack-lang/gong/lib/table/ng-github.com-fullstack-lang-gong-lib-table/projects/tablespecific/src/lib/table-specific/table-specific.component'
import { FormSpecificComponent } from '@vendored_components/github.com/fullstack-lang/gong/lib/table/ng-github.com-fullstack-lang-gong-lib-table/projects/tablespecific/src/lib/form-specific/form-specific.component'
import { SvgSpecificComponent } from '@vendored_components/github.com/fullstack-lang/gong/lib/svg/ng-github.com-fullstack-lang-gong-lib-svg/projects/svgspecific/src/lib/svg-specific/svg-specific.component'
import { DocSpecificComponent } from '@vendored_components/github.com/fullstack-lang/gong/lib/doc/ng-github.com-fullstack-lang-gong-lib-doc/projects/docspecific/src/lib/doc-specific/doc-specific.component'

import * as svg from '@vendored_components/github.com/fullstack-lang/gong/lib/svg/ng-github.com-fullstack-lang-gong-lib-svg/projects/svg/src/public-api'
import * as tree from '@vendored_components/github.com/fullstack-lang/gong/lib/tree/ng-github.com-fullstack-lang-gong-lib-tree/projects/tree/src/public-api'
import * as table from '@vendored_components/github.com/fullstack-lang/gong/lib/table/ng-github.com-fullstack-lang-gong-lib-table/projects/table/src/public-api'


@Component({
    selector: 'app-root',
    imports: [
        CommonModule,
        FormsModule,
        MatRadioModule,
        MatButtonModule,
        MatIconModule,
        AngularSplitModule,

        TreeSpecificComponent,
        TableSpecificComponent,
        FormSpecificComponent,
        SvgSpecificComponent,
        DocSpecificComponent,

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

  TableExtraPathEnum = table.TableExtraPathEnum

  StacksNames = phyllotaxymusic.StacksNames
  SVGStackType = svg.StackType
  TreeStackType = tree.StackType
  ToneStackType = gongtone.StackType


  constructor(
  ) {

  }

  ngOnInit(): void {
  }
}
