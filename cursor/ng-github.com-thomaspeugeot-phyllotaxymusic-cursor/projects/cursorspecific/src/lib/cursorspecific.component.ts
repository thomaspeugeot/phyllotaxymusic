import { Component, OnInit } from '@angular/core';
import * as cursor from '../../../cursor/src/public-api'

@Component({
  selector: 'lib-cursorspecific',
  standalone: true,
  imports: [],
  template: `
    <svg [attr.width]="xe" height="1000">
     <line
        [attr.x1]="x"
        [attr.y1]="0"
        [attr.x2]="x"
        [attr.y2]="1000"
        stroke="grey"
        stroke-width="6"
        stroke-opacity="0.5"
      />

       <line 
        [attr.x1]="xe"
        [attr.y1]="0"
        [attr.x2]="xe"
        [attr.y2]="1000"
        stroke="grey"
        stroke-width="6"
        stroke-opacity="0.5"
      />
      
    </svg>
  `,
  styles: ``
})
export class CursorspecificComponent implements OnInit {
  x = 0;
  xe = 500
  private animationFrameId: number | null = null;  // Store animation frame ID

  StacksNames = cursor.StacksNames;
  public frontRepo?: cursor.FrontRepo;
  public cursor: cursor.Cursor | undefined

  constructor(
    private frontRepoService: cursor.FrontRepoService,
  ) { }

  ngOnInit(): void {
    console.log("ngOnInit");

    this.frontRepoService.connectToWebSocket(this.StacksNames.Cursorstakcname).subscribe({
      next: (gongtablesFrontRepo) => {
        this.frontRepo = gongtablesFrontRepo;

        let cursors = this.frontRepo.getFrontArray<cursor.Cursor>(cursor.Cursor.GONGSTRUCT_NAME);

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

  public startEmittingPosition(): void {
    if (this.cursor == undefined) {
      return;
    }

    const duration = 1000 * this.cursor.DurationSeconds;
    const progressAbciss = this.cursor.EndX - this.cursor.StartX;
    let startTime: number | null = null;
    let previousTimestamp: number | null = null;

    const animate = (timestamp: number) => {
      if (!startTime) {
        startTime = timestamp;
        previousTimestamp = timestamp;
      }

      // Calculate elapsed time since last frame
      const deltaTime = timestamp - previousTimestamp!;
      previousTimestamp = timestamp;

      // Subtract 50ms from the elapsed time calculation
      const elapsed = ((timestamp - startTime - 300) % duration);
      const progress = elapsed / duration;
      this.x = this.cursor!.StartX + progress * progressAbciss;

      this.animationFrameId = requestAnimationFrame(animate);
    }

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