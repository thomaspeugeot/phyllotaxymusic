import { Component, Input } from '@angular/core';

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
    @Input() Name: string = ""
}
