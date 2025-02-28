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

import { TreeComponent } from '@vendored_components/github.com/fullstack-lang/gongtree/ng-github.com-fullstack-lang-gongtree/projects/gongtreespecific/src/public-api'
import { MaterialTableComponent } from '@vendored_components/github.com/fullstack-lang/gongtable/ng-github.com-fullstack-lang-gongtable/projects/gongtablespecific/src/lib/material-table/material-table.component';
import { MaterialFormComponent } from '@vendored_components/github.com/fullstack-lang/gongtable/ng-github.com-fullstack-lang-gongtable/projects/gongtablespecific/src/lib/material-form/material-form.component';
import * as gongtable from '@vendored_components/github.com/fullstack-lang/gongtable/ng-github.com-fullstack-lang-gongtable/projects/gongtable/src/public-api';
import { PanelComponent } from '@vendored_components/github.com/fullstack-lang/gongdoc/ng-github.com-fullstack-lang-gongdoc/projects/gongdocspecific/src/public-api'
import { GongsvgDiagrammingComponent } from '@vendored_components/github.com/fullstack-lang/gongsvg/ng-github.com-fullstack-lang-gongsvg/projects/gongsvgspecific/src/lib/gongsvg-diagramming/gongsvg-diagramming'

import * as gongsvg from '@vendored_components/github.com/fullstack-lang/gongsvg/ng-github.com-fullstack-lang-gongsvg/projects/gongsvg/src/public-api';
import * as gongtree from '@vendored_components/github.com/fullstack-lang/gongtree/ng-github.com-fullstack-lang-gongtree/projects/gongtree/src/public-api';
import * as gongtone from '@vendored_components/github.com/fullstack-lang/gongtone/ng-github.com-fullstack-lang-gongtone/projects/gongtone/src/public-api';

@Component({
    selector: 'app-root',
    imports: [
        CommonModule,
        FormsModule,
        MatRadioModule,
        MatButtonModule,
        MatIconModule,
        AngularSplitModule,
        TreeComponent,
        MaterialTableComponent,
        MaterialFormComponent,
        PanelComponent,
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

  TableExtraPathEnum = gongtable.TableExtraPathEnum

  StacksNames = phyllotaxymusic.StacksNames
  SVGStackType = gongsvg.StackType
  TreeStackType = gongtree.StackType
  ToneStackType = gongtone.StackType


  constructor(
  ) {

  }

  ngOnInit(): void {
  }
}
