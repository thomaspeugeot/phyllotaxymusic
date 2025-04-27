import { Component, OnInit, OnDestroy, NgZone, Input } from '@angular/core';
import { MatIconModule } from '@angular/material/icon';
import { MatDividerModule } from '@angular/material/divider';
import { MatButtonModule } from '@angular/material/button';
import { MatSnackBar } from '@angular/material/snack-bar';
import { MatProgressSpinnerModule } from '@angular/material/progress-spinner'; // Import progress spinner
import { CommonModule } from '@angular/common'; // Import CommonModule for *ngIf

// Import Recorder
import * as Tone from 'tone';
// No need for specific Recorder import if using Tone.Recorder
// import { Recorder } from 'tone';

import * as tonelocal from '../../../../tone/src/public-api';

import { Subject } from 'rxjs';
import { takeUntil, catchError } from 'rxjs/operators';

@Component({
  selector: 'lib-tone-specific',
  standalone: true, // Make it standalone
  imports: [
    CommonModule, // Add CommonModule for *ngIf
    MatButtonModule,
    MatDividerModule,
    MatIconModule,
    MatProgressSpinnerModule // Add progress spinner module
  ],
  templateUrl: './tone-specific.component.html',
  styleUrl: './tone-specific.component.css'
})
export class ToneSpecificComponent implements OnInit, OnDestroy {

  @Input() Name: string = ""

  // Tone.js specific properties
  private synth: Tone.PolySynth | undefined;
  private sampler: Tone.Sampler | undefined;
  private currentLoop: Tone.Loop | undefined;
  private recorder: Tone.Recorder | undefined; // Recorder instance

  // RxJS Subject for unsubscribing
  private destroy$ = new Subject<void>();

  // Component state properties
  frontRepo?: tonelocal.FrontRepo;
  isLoading = false; // For loading indicators
  isPlaying = false; // True if audio is currently playing
  isRecording = false; // True if recorder is active

  constructor(
    // Inject necessary services
    private frontRepoService: tonelocal.FrontRepoService,
    private playerService: tonelocal.PlayerService,
    private ngZone: NgZone, // Angular zone for UI updates
    private snackBar: MatSnackBar // For showing messages
  ) { }

  /**
   * Initializes audio components and connects to WebSocket on component init.
   */
  ngOnInit(): void {
    // Initialize audio context and recorder. Requires user gesture in some browsers.
    this.initializeAudio();
    // Establish WebSocket connection for data updates.
    this.connectToWebSocket();
  }

  /**
   * Cleans up resources like subscriptions, loops, and Tone.js objects.
   */
  ngOnDestroy(): void {
    // Signal subscriptions to complete.
    this.destroy$.next();
    this.destroy$.complete();
    // Stop any ongoing playback and recording.
    this.stopPlayback(); // Ensure cleanup happens
    // Dispose of the recorder instance.
    this.recorder?.dispose();
    // Dispose sampler if it exists
    this.sampler?.dispose();
    // Dispose synth if it exists
    this.synth?.dispose();
    // Dispose loop if it exists
    this.currentLoop?.dispose();
    console.log("ToneSpecificComponent destroyed and cleaned up.");
  }

  /**
   * Initializes the base audio components (Synth, Recorder).
   * Must be called after a user interaction (e.g., button click) to ensure
   * AudioContext starts correctly in all browsers.
   */
  private async initializeAudio(): Promise<void> {
    try {
      // Start Tone.js AudioContext. Best practice is to call this within a user gesture handler.
      await Tone.start();
      console.log('AudioContext started successfully.');

      // Initialize a PolySynth (though sampler is primarily used here).
      this.synth = new Tone.PolySynth().toDestination();
      // Initialize the recorder.
      this.recorder = new Tone.Recorder();
      console.log('Synth and Recorder initialized.');

    } catch (error) {
      // Handle errors during audio initialization.
      this.handleAudioInitError(error);
    }
  }

  /**
   * Connects to the WebSocket service to receive FrontRepo updates.
   */
  private connectToWebSocket(): void {
    this.frontRepoService.connectToWebSocket(this.Name)
      .pipe(
        takeUntil(this.destroy$), // Automatically unsubscribe on component destroy.
        catchError(error => {
          // Handle WebSocket connection errors.
          this.handleWebSocketError(error);
          throw error; // Rethrow or handle as needed.
        })
      )
      .subscribe(gongtablesFrontRepo => {
        // Store the received FrontRepo data.
        this.frontRepo = gongtablesFrontRepo;
        console.log("WebSocket connected, FrontRepo received.");
      });
  }

  /**
   * Stops the currently playing audio sources (loop, transport, sampler notes).
   * Does not stop the recorder.
   */
  private stopAudioSources(): void {
    // Run Tone.js operations outside Angular zone for performance.
    this.ngZone.runOutsideAngular(() => {
       try {
        // Stop and dispose the current loop if it exists.
        if (this.currentLoop) {
          this.currentLoop.stop(0); // Stop immediately at the current position.
          // Dispose is handled later or in ngOnDestroy to avoid race conditions if stopPlayback is called rapidly
          // this.currentLoop.dispose();
          this.currentLoop = undefined; // Clear reference
          console.log("Loop stopped.");
        }

        // Stop the Tone.js transport and cancel any scheduled events.
        Tone.Transport.stop();
        Tone.Transport.cancel();
        console.log("Transport stopped and events cancelled.");

        // Release any notes currently held by the sampler.
        this.sampler?.releaseAll();
        console.log("Sampler notes released.");

        // Update playing state (outside Angular zone initially).
        this.isPlaying = false;
        console.log("Audio sources stopped.");

      } catch (error) {
        console.error('Error stopping audio sources:', error);
        // Potentially show an error message to the user here if needed.
      }
    });
    // Ensure Angular UI reflects the state change.
    this.ngZone.run(() => {
        this.isPlaying = false;
    });
  }

  /**
   * Stops all playback and recording, updates backend status, and returns the recorded audio Blob.
   * @returns A Promise resolving to the recorded audio Blob, or null if no recording was active or an error occurred.
   */
  async stopPlayback(): Promise<Blob | null> {
    console.log("stopPlayback called");
    // Stop the audio loops, transport, etc.
    this.stopAudioSources();

    let recordingBlob: Blob | null = null;
    // Check if recording was active and recorder exists.
    if (this.isRecording && this.recorder) {
      try {
        console.log("Stopping recorder...");
        // Make sure recorder state is actually recording before stopping
        if (this.recorder.state === "started") {
             // Stop the recorder and wait for the Blob data.
             // Run recorder stop outside zone, handle result inside.
             await this.ngZone.runOutsideAngular(async () => {
                 // Use ! here as we checked recorder existence and state
                 recordingBlob = await this.recorder!.stop();
             });
             console.log("Recorder stopped, blob received.");
        } else {
            console.warn("Recorder was marked as recording, but state was not 'started'. No blob generated.");
        }
        // Update recording state regardless of whether stop was called
        this.ngZone.run(() => {
            this.isRecording = false; // Update recording state in zone
        });
      } catch (error) {
        console.error('Error stopping recorder:', error);
        this.showError('Failed to stop recording');
        this.ngZone.run(() => { this.isRecording = false; }); // Ensure state reset on error
        recordingBlob = null; // Ensure null is returned on error
      }
    } else {
         // If not recording, ensure the flag is false.
         this.ngZone.run(() => { this.isRecording = false; });
    }

    // Dispose of the sampler instance after stopping everything.
    // Use optional chaining for safety.
    this.sampler?.dispose();
    this.sampler = undefined;
    console.log("Sampler disposed.");

    // Dispose of the loop instance
    this.currentLoop?.dispose();
    this.currentLoop = undefined;
    console.log("Loop disposed.");


    // Update the backend player status to PAUSED.
    // Use optional chaining for frontRepo access
    const players = this.frontRepo?.getFrontArray<tonelocal.Player>(tonelocal.Player.GONGSTRUCT_NAME);
    if (players && players.length === 1) {
      const player = players[0];
      // Update status only if it's not already PAUSED.
      if (player.Status !== tonelocal.Status.PAUSED) {
          player.Status = tonelocal.Status.PAUSED;
          // Update the backend via the service.
          this.playerService.updateFront(player, this.Name).subscribe({
              next: () => console.log("gongtone: status set to PAUSED"),
              error: (err) => console.error("Failed to update player status:", err)
          });
       }
    } else if (this.frontRepo) { // Only log if frontRepo exists but players are wrong
        console.log("Player not found or multiple players exist, status not updated.");
    }

    // Return the recorded blob (or null).
    return recordingBlob;
  }

  /**
   * Triggers the download of the provided audio Blob.
   * @param blob The recorded audio Blob to download.
   */
  private downloadRecording(blob: Blob | null): void {
    if (!blob) {
      this.showError('No recording available to download.');
      return;
    }
    // Run DOM manipulation within Angular zone if it affects UI/bindings.
    this.ngZone.run(() => {
        try {
            // Create a URL for the Blob object.
            const url = URL.createObjectURL(blob);
            // Create a temporary anchor element.
            const anchor = document.createElement('a');
            // Set the download filename. Uses component Name or default.
            // Tone.Recorder usually produces webm with opus codec
            anchor.download = `gongtone_recording_${this.Name || 'audio'}.webm`;
            // Set the anchor's href to the Blob URL.
            anchor.href = url;
            // Append to body to ensure click works on all browsers (especially Firefox)
            document.body.appendChild(anchor);
            // Programmatically click the anchor to trigger the download.
            anchor.click();
            // Clean up: remove anchor and revoke Object URL.
            document.body.removeChild(anchor);
            URL.revokeObjectURL(url);
            console.log("Download triggered for recording.");
            // Show a confirmation message.
            this.snackBar.open('Recording download started.', 'Close', { duration: 3000 });
        } catch(error) {
            console.error("Error creating download link:", error);
            this.showError('Failed to initiate download.');
        }
    });
  }

  /**
   * Stops playback/recording and initiates the download of the recorded audio.
   * Called by the download button.
   */
  async stopAndDownload(): Promise<void> {
    // Prevent action if nothing is playing or recording.
    if (!this.isPlaying && !this.isRecording) {
       this.showError("Nothing is playing or recording.");
       return;
    }
    console.log("Stop and Download initiated...");
    // Stop playback/recording and get the resulting Blob.
    const blob = await this.stopPlayback();
    // Trigger the download process for the Blob.
    this.downloadRecording(blob);
  }

  /**
   * Starts the audio playback and recording process.
   * Called by the play button.
   */
  play(): void {
    // --- Check 1: frontRepo must exist ---
    if (!this.frontRepo) {
      this.showError('No data available for playback (frontRepo is undefined).');
      return;
    }
    // --- Check 2: Not already playing or loading ---
    if (this.isPlaying || this.isLoading) {
        console.log("Already playing or loading.");
        this.showError("Playback or loading is already in progress.");
        return;
    }
    // --- Check 3: Recorder must be initialized ---
    if (!this.recorder) {
        this.showError('Recorder not initialized. Please try again.');
        return;
    }

    // --- Hold a reference to frontRepo after the check ---
    const currentFrontRepo = this.frontRepo;

    // Ensure Tone.js AudioContext is running before proceeding.
    Tone.start().then(() => {
        console.log("AudioContext confirmed running for play.");

        // Stop any previous playback/recording cleanly first.
        this.stopPlayback().then(() => {
            console.log("Previous playback stopped, starting new playback...");
            this.ngZone.run(() => { // Set loading state within zone
                this.isLoading = true;
                this.isRecording = false; // Ensure recording flag is reset
            });


            // Perform Tone.js setup outside Angular zone.
            this.ngZone.runOutsideAngular(() => {
                try {
                    // --- Access notes using the checked reference ---
                    const notes = currentFrontRepo.getFrontArray<tonelocal.Note>(tonelocal.Note.GONGSTRUCT_NAME);
                    if (!notes || notes.length === 0) {
                        throw new Error("No notes data found to play.");
                    }
                    const duration = this.calculateTotalDuration(notes);
                    if (duration <= 0) {
                         throw new Error("Invalid playback duration calculated (0 or less).");
                    }
                    this.initializeSampler(duration, notes); // Sampler init leads to startPlayback

                } catch (error) {
                    this.handlePlaybackError(error); // Handles state reset and shows error
                }
            });

            // Update backend player status to PLAYING.
            // --- Access players using the checked reference ---
            const players = currentFrontRepo.getFrontArray<tonelocal.Player>(tonelocal.Player.GONGSTRUCT_NAME);
            if (players && players.length === 1) {
              const player = players[0];
              if (player.Status !== tonelocal.Status.PLAYING) {
                  player.Status = tonelocal.Status.PLAYING;
                  this.playerService.updateFront(player, this.Name).subscribe({
                      next: () => console.log("gongtone: status set to PLAYING"),
                      error: (err) => console.error("Failed to update player status:", err)
                  });
              }
            } else {
                console.log("Player not found or multiple players exist, status not updated.");
            }

        }); // End .then() for stopPlayback

    }).catch(err => {
        // Handle errors starting the AudioContext.
        console.error("Failed to start AudioContext for play:", err);
        this.showError("Could not start audio playback. Please interact with the page first.");
        this.ngZone.run(() => { this.isLoading = false; }); // Reset loading state
    });
  }

  /**
   * Initializes the Tone.Sampler with audio files and connects it for playback and recording.
   * @param duration The total duration of the piece to be played.
   * @param notes The array of Note objects to schedule.
   */
  private initializeSampler(duration: number, notes: tonelocal.Note[]): void {
    // Ensure recorder is ready before creating sampler.
    if (!this.recorder) {
        console.error("Recorder not available for sampler initialization.");
        this.handlePlaybackError(new Error("Recorder not initialized"));
        return; // Exit early
    }

    // Define the base URL for audio assets.
    const audioBaseUrl = `${window.location.origin}/assets/audio/salamander/`;
    // Define the mapping of note names to audio file URLs.
    const urls: { [key: string]: string } = {
        C3: `C3.mp3`, 'D#3': `Ds3.mp3`, 'F#3': `Fs3.mp3`,
        A3: `A3.mp3`, C4: `C4.mp3`, 'D#4': `Ds4.mp3`,
        'F#4': `Fs4.mp3`, A4: `A4.mp3`, C5: `C5.mp3`,
        'D#5': `Ds5.mp3`
        // Add other necessary notes
    };

    // Dispose previous sampler if it exists to prevent memory leaks
    this.sampler?.dispose();
    this.sampler = undefined; // Clear reference
    console.log("Disposed previous sampler instance if any.");


    // Create a new Tone.Sampler instance.
    this.sampler = new Tone.Sampler({
      urls: urls, // The map of note names to URLs.
      release: 1, // Envelope release time.
      baseUrl: audioBaseUrl, // Base URL for loading samples.
      onload: () => {
        // This runs when all samples are loaded successfully.
        this.ngZone.run(() => { // Run UI updates in zone if needed.
            console.log('Sampler loaded successfully');
             // Ensure sampler and recorder still exist before connecting
            if (!this.sampler || !this.recorder) {
                 console.error("Sampler or recorder became undefined before connection in onload.");
                 this.handlePlaybackError(new Error("Audio components missing during setup"));
                 return;
            }
            // Connect sampler output to speakers (destination) AND recorder.
            this.sampler.connect(Tone.getDestination());
            this.sampler.connect(this.recorder);
            console.log("Sampler connected to destination and recorder.");

            // Now that sampler is ready and connected, start the actual playback loop.
            this.startPlayback(duration, notes);
        });
      },
      onerror: (error) => {
        // This runs if there's an error loading samples.
        this.ngZone.run(() => { // Ensure error handling runs in NgZone.
            console.error('Sampler load error:', error);
             if (error instanceof Event && error.target instanceof HTMLAudioElement) {
                 console.error('Failed Audio Source:', {
                    src: error.target.src,
                    error: error.target.error, // Log the MediaError object
                    errorCode: error.target.error?.code // Log specific error code
                 });
             }
            this.handleSamplerLoadError(error); // Handle error (shows snackbar, sets loading=false).
        });
      }
    });
    // Note: Do not connect to destination here; it happens in onload.
  }

  /**
   * Calculates the total duration needed for the playback loop based on note start times and durations.
   * @param notes Array of Note objects.
   * @returns The maximum end time among all notes, in seconds. Returns 0 if no notes.
   */
  private calculateTotalDuration(notes: tonelocal.Note[]): number {
     if (!notes || notes.length === 0) return 0;
     // Find the maximum time point reached by any note's end.
     return notes.reduce((maxDuration, note) =>
      // Ensure Start and Duration are numbers
      Math.max(maxDuration, (note.Start || 0) + (note.Duration || 0)),
     0);
  }

  /**
   * Creates the Tone.Loop, starts the recorder, and starts the Tone.Transport.
   * @param duration The duration for the loop.
   * @param notes The notes to schedule within the loop.
   */
  private startPlayback(duration: number, notes: tonelocal.Note[]): void {
    // Double-check required components are ready.
    if (!this.sampler || !this.recorder) {
        console.error("Sampler or Recorder not ready for playback start.");
        this.handlePlaybackError(new Error("Audio components not ready"));
        return; // Exit early
    }
    // Validate duration.
     if (duration <= 0) {
         console.warn("Calculated duration is zero or negative, cannot start loop.");
         this.handlePlaybackError(new Error("Invalid playback duration"));
         return; // Exit early
     }

    // Run Tone.js scheduling outside Angular zone.
    this.ngZone.runOutsideAngular(() => {
        try {
            // Ensure transport is stopped and cleared before starting new loop.
            Tone.Transport.stop();
            Tone.Transport.cancel();

            // Dispose previous loop if any to prevent duplicates
            this.currentLoop?.dispose();
            this.currentLoop = undefined; // Clear reference
            console.log("Disposed previous loop if any.");


            // Create the main playback loop.
            this.currentLoop = new Tone.Loop((time) => {
                // --- ADDED CHECK ---
                // Check if sampler still exists at the start of each loop iteration
                if (!this.sampler) {
                    console.warn("Sampler became undefined during loop execution. Stopping loop.");
                    // Stop the loop itself from within the callback
                    if (this.currentLoop) { // Check if currentLoop itself is defined
                       this.currentLoop.stop(0);
                       this.currentLoop.dispose();
                       this.currentLoop = undefined;
                    }
                    // Trigger the full stop procedure outside the loop callback (schedule it)
                    this.ngZone.run(() => this.stopPlayback());
                    return; // Exit this iteration of the loop callback
                }
                // --- END ADDED CHECK ---

                // Schedule each note within the loop.
                notes.forEach(note => {
                    // Map frequency names (e.g., "C4") from the note data.
                    const frequencies = note.Frequencies.map(freq => freq.Name);
                    // Trigger the sampler note - Sampler existence is checked above
                    // Use non-null assertion '!' as check guarantees existence
                    this.sampler!.triggerAttackRelease(frequencies, note.Duration, time + note.Start);
                });
            }, duration).start(0); // Start the loop immediately (at time 0).

            // --- Recorder Start ---
            // Check recorder *again* immediately before starting it.
            if (!this.recorder) {
                 console.error("Recorder became undefined before starting.");
                 throw new Error("Recorder not available"); // Throw to be caught by outer try/catch
            }
            // Start the recorder asynchronously.
            this.recorder.start().then(() => {
                // Recorder successfully started.
                this.ngZone.run(() => { this.isRecording = true; }); // Update state in zone.
                console.log("Recorder started.");

                // Start the transport only AFTER recorder confirms start.
                Tone.Transport.start();
                console.log("Transport started.");

                // Update component state within Angular zone.
                 this.ngZone.run(() => {
                    this.isLoading = false; // Loading finished.
                    this.isPlaying = true;  // Playback started.
                 });

            }).catch(err => {
                 // Handle errors specifically from recorder.start().
                 this.ngZone.run(() => {
                    console.error("Failed to start recorder:", err);
                    // Pass specific error message, converting non-Error types to string
                    this.handlePlaybackError(new Error(`Recorder start failed: ${err instanceof Error ? err.message : String(err)}`));
                });
            });
            // --- End Recorder Start ---

        } catch (error) {
            // Catch synchronous errors in loop creation or recorder check.
            this.ngZone.run(() => { // Ensure error handling runs in NgZone.
                this.handlePlaybackError(error);
            });
        }
    });
}


  // --- Error Handling and UI Feedback Methods ---

  /** Handles WebSocket connection errors. */
  private handleWebSocketError(error: any): void {
    console.error('WebSocket connection error:', error);
    this.showError('Failed to connect to WebSocket');
  }

  /** Handles errors during Tone.js/AudioContext initialization. */
  private handleAudioInitError(error: any): void {
    console.error('Audio initialization error:', error);
    // Provide more specific feedback if possible (e.g., about user gesture).
    let message = 'Failed to initialize audio.';
    if (error instanceof Error && error.message.toLowerCase().includes('user gesture')) {
        message += ' Please click the page or a button first.';
    } else if (error instanceof Error) {
         message += ` ${error.message}`;
    }
    this.showError(message);
  }

  /** Handles errors during sampler audio file loading. */
  private handleSamplerLoadError(error: any): void {
    console.error('Sampler load error:', error);
    this.showError('Failed to load audio samples. Check console for details.');
    // Reset loading state within Angular zone.
    this.ngZone.run(() => { this.isLoading = false; });
  }

  /** Handles generic playback errors (scheduling, unexpected issues). */
  private handlePlaybackError(error: any): void {
    console.error('Playback error:', error);
    // Display a user-friendly message.
    this.showError(`Playback failed: ${error instanceof Error ? error.message : 'Unknown error'}`);
     // Ensure flags are reset correctly within the Angular zone.
     this.ngZone.run(() => {
        this.isLoading = false;
        this.isPlaying = false;
        this.isRecording = false; // Also reset recording flag on playback error.
     });
     // Optionally stop audio sources again just in case, outside zone
     this.stopAudioSources();
  }

  /**
   * Displays an error message to the user using MatSnackBar.
   * @param message The message to display.
   */
  private showError(message: string): void {
    // Ensure snackbar display runs within the Angular zone for proper UI updates.
    this.ngZone.run(() => {
      // Check if snackBar is available (it should be injected, but as a safeguard)
      if (this.snackBar) {
          this.snackBar.open(message, 'Close', {
            duration: 4000, // Slightly longer duration for errors.
            horizontalPosition: 'center',
            verticalPosition: 'bottom' // Bottom is often less intrusive.
          });
      } else {
          console.error("Snackbar service not available to show error:", message);
      }
    });
  }
}
