import { Component, OnInit } from '@angular/core';
import * as substackcursor from '../../../substackcursor/src/public-api'

@Component({
  selector: 'lib-substackcursorspecific',
  standalone: true,
  imports: [],
  template: `
    <svg width="1000" height="1000">
     <line
        [attr.x1]="x"
        [attr.y1]="0"
        [attr.x2]="x"
        [attr.y2]="1000"
        stroke="black"
        stroke-width="6"
      />

       <line 
        [attr.x1]="xe"
        [attr.y1]="0"
        [attr.x2]="xe"
        [attr.y2]="1000"
        stroke="black"
        stroke-width="6"
      />
      
    </svg>
  `,
  styles: ``
})
export class SubstackcursorspecificComponent implements OnInit {
  x = 0;
  xe = 500
  private animationFrameId: number | null = null;  // Store animation frame ID

  StacksNames = substackcursor.StacksNames;
  public frontRepo?: substackcursor.FrontRepo;
  public cursor: substackcursor.Cursor | undefined

  constructor(
    private frontRepoService: substackcursor.FrontRepoService,
  ) { }

  ngOnInit(): void {
    console.log("ngOnInit");

    this.frontRepoService.connectToWebSocket(this.StacksNames.Substackcursor).subscribe({
      next: (gongtablesFrontRepo) => {
        this.frontRepo = gongtablesFrontRepo;

        let cursors = this.frontRepo.getFrontArray<substackcursor.Cursor>(substackcursor.Cursor.GONGSTRUCT_NAME);

        console.assert(cursors.length == 1);
        this.cursor = cursors[0];
        this.x = this.cursor.StartX
        this.xe = this.cursor.EndX

        if (this.cursor.IsPlaying == true) {
          this.startEmittingPosition();
        } else {
          this.stopEmittingPosition();  // Stop the animation if IsPlaying is false
        }
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

    if (this.cursor == undefined) {
      return
    }

    const duration = 1000 * this.cursor.DurationSeconds   // 5 seconds
    const endPosition = this.cursor?.EndX;
    let startTime: number | null = null;

    const animate = (timestamp: number) => {
      if (!startTime) {
        startTime = timestamp;
      }
      const elapsed = timestamp - startTime;
      // progress goes from 0.0 to 1.0 as time goes by
      const progress = Math.min(elapsed / duration, 1);
      // Position is progress * 1000
      this.x = this.cursor!.StartX + progress * endPosition;

      // Keep going until we reach progress=1.0
      if (progress < 1) {
        this.animationFrameId = requestAnimationFrame(animate);
      }
    };

    // Kick off our "animation"
    this.animationFrameId = requestAnimationFrame(animate);
  }

  /**
   * Stop emitting position by canceling the animation frame.
   */
  public stopEmittingPosition(): void {
    if (this.animationFrameId !== null) {
      cancelAnimationFrame(this.animationFrameId);
      this.animationFrameId = null;
    }
  }
}
