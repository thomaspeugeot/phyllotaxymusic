import { Component, OnInit } from '@angular/core';

import * as substackcursor from '../../../substackcursor/src/public-api'


@Component({
  selector: 'lib-substackcursorspecific',
  standalone: true,
  imports: [],
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
  styles: ``
})
export class SubstackcursorspecificComponent implements OnInit {
  x = 0

  StacksNames = substackcursor.StacksNames
  public frontRepo?: substackcursor.FrontRepo



  constructor(
    private frontRepoService: substackcursor.FrontRepoService,
  ) {

  }

  ngOnInit(): void {

    console.log("ngOnInit")

    this.frontRepoService.connectToWebSocket(this.StacksNames.Substackcursor).subscribe({
      next: (gongtablesFrontRepo) => {
        this.frontRepo = gongtablesFrontRepo;

        this.startEmittingPosition()
      },
      error: (err) => {
        // this will capture any errors thrown by the Observable
        console.error("WebSocket error from Observable:", err);
      },
      complete: () => {
        console.log("WebSocket connection complete.");
      },
    });

  }

  /**
 * Smoothly move the line from x=0 to x=1000 over 5 seconds using requestAnimationFrame.
 */
  public startEmittingPosition(): void {
    const duration = 5000;   // 5 seconds
    const endPosition = 1000;
    let startTime: number | null = null;

    const animate = (timestamp: number) => {
      if (!startTime) {
        startTime = timestamp;
      }
      const elapsed = timestamp - startTime;
      // progress goes from 0.0 to 1.0 as time goes by
      const progress = Math.min(elapsed / duration, 1);
      // Position is progress * 1000
      this.x = progress * endPosition

      // Keep going until we reach progress=1.0
      if (progress < 1) {
        requestAnimationFrame(animate);
      }
    };

    // Kick off our "animation"
    requestAnimationFrame(animate);
  }

}
