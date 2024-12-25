import { Component, OnInit } from '@angular/core';
import { CommonModule } from '@angular/common';
import { FakeWebsocketService } from './fake-websocket.service';

@Component({
    selector: 'app-moving-line',
    standalone: true,
    imports: [CommonModule],
    template: `
    <svg width="1000" height="1000">
      <!-- 
        A single vertical line, whose x-position is bound
        to the 'x' field updated by the fake websocket.
      -->
      <line
        [attr.x1]="x"
        y1="0"
        [attr.x2]="x"
        y2="1000"
        stroke="black"
        stroke-width="2"
      />
    </svg>
  `,
})
export class MovingLineComponent implements OnInit {
    x = 0;

    constructor(private fakeWebsocketService: FakeWebsocketService) { }

    ngOnInit(): void {
        // Whenever the service updates the position, set our local 'x' variable
        this.fakeWebsocketService.getPositionUpdates().subscribe((position) => {
            this.x = position;
        });
    }
}
