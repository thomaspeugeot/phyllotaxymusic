import { Component, Input, model, OnInit } from '@angular/core';

import * as phyllotaxymusic from '../../../phyllotaxymusic/src/public-api'

import { SplitSpecificComponent } from '@vendored_components/github.com/fullstack-lang/gong/lib/split/ng-github.com-fullstack-lang-gong-lib-split/projects/splitspecific/src/lib/split-specific/split-specific.component';

@Component({
    selector: 'lib-phyllotaxymusicspecific',
    imports: [
        SplitSpecificComponent
    ],
    templateUrl: './phyllotaxymusicspecific.component.html',
    styleUrls: ['./phyllotaxymusicspecific.component.css']
})
export class PhyllotaxymusicspecificComponent {

  StacksNames = phyllotaxymusic.StacksNames

}
