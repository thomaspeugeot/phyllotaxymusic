import { Component, OnInit } from '@angular/core';


import * as phyllotaxymusic from '../../projects/phyllotaxymusic/src/public-api'

import { PhyllotaxymusicspecificComponent } from '../../projects/phyllotaxymusicspecific/src/public-api'

@Component({
    selector: 'app-root',
    imports: [
        PhyllotaxymusicspecificComponent
    ],
    templateUrl: './app.component.html'
})
export class AppComponent {

}
