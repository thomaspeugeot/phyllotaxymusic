import { Component } from '@angular/core';

import * as phylotaxymusic from '../../../phylotaxymusic/src/public-api'


import { GongsvgDiagrammingComponent } from '@vendored_components/github.com/fullstack-lang/gongsvg/ng-github.com-fullstack-lang-gongsvg/projects/gongsvgspecific/src/lib/gongsvg-diagramming/gongsvg-diagramming'

@Component({
  selector: 'lib-phylotaxymusicspecific',
  standalone: true,
  imports: [
    GongsvgDiagrammingComponent
  ],
  template: `
    <lib-gongsvg-diagramming [GONG__StackPath]="StacksNames.GongsvgStackName">
    </lib-gongsvg-diagramming>
  
  `,
  styles: ``
})
export class PhylotaxymusicspecificComponent {

  StacksNames = phylotaxymusic.StacksNames
}
