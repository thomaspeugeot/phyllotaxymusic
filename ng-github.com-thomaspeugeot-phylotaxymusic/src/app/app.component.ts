import { Component, OnInit } from '@angular/core';

import { Observable, combineLatest, timer } from 'rxjs'

// for angular & angular material
import { FormsModule } from '@angular/forms';
import { CommonModule } from '@angular/common';
import { MatRadioModule } from '@angular/material/radio';
import { MatButtonModule } from '@angular/material/button';
import { MatIconModule } from '@angular/material/icon';

import { AngularSplitModule } from 'angular-split';

import * as phylotaxymusic from '../../projects/phylotaxymusic/src/public-api'

import { PhylotaxymusicspecificComponent } from '../../projects/phylotaxymusicspecific/src/public-api'

import { TreeComponent } from '@vendored_components/github.com/fullstack-lang/gongtree/ng-github.com-fullstack-lang-gongtree/projects/gongtreespecific/src/public-api'
import { MaterialTableComponent } from '@vendored_components/github.com/fullstack-lang/gongtable/ng-github.com-fullstack-lang-gongtable/projects/gongtablespecific/src/lib/material-table/material-table.component';
import { MaterialFormComponent } from '@vendored_components/github.com/fullstack-lang/gongtable/ng-github.com-fullstack-lang-gongtable/projects/gongtablespecific/src/lib/material-form/material-form.component';
import * as gongtable from '@vendored_components/github.com/fullstack-lang/gongtable/ng-github.com-fullstack-lang-gongtable/projects/gongtable/src/public-api';
import { PanelComponent } from '@vendored_components/github.com/fullstack-lang/gongdoc/ng-github.com-fullstack-lang-gongdoc/projects/gongdocspecific/src/public-api'
import { GongsvgDiagrammingComponent } from '@vendored_components/github.com/fullstack-lang/gongsvg/ng-github.com-fullstack-lang-gongsvg/projects/gongsvgspecific/src/lib/gongsvg-diagramming/gongsvg-diagramming'

import * as gongsvg from '@vendored_components/github.com/fullstack-lang/gongsvg/ng-github.com-fullstack-lang-gongsvg/projects/gongsvg/src/public-api';

@Component({
  selector: 'app-root',
  standalone: true,
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

    PhylotaxymusicspecificComponent

  ],

  templateUrl: './app.component.html',
})
export class AppComponent implements OnInit {

  phylotaxymusic = 'Phylotaxymusic'
  probe = 'Phylotaxymusic Data/Model'
  svg_probe = 'svg probe'
  view = this.phylotaxymusic

  views: string[] = [this.phylotaxymusic, this.probe, this.svg_probe];

  scrollStyle = {
    'overflow- x': 'auto',
    'width': '100%',  // Ensure the div takes the full width of its parent container
  }

  StackName = "phylotaxymusic"
  StackType = phylotaxymusic.StackType

  TableExtraPathEnum = gongtable.TableExtraPathEnum

  StacksNames = phylotaxymusic.StacksNames
  SVGStackType = gongsvg.StackType


  constructor(
  ) {

  }

  ngOnInit(): void {
  }
}
