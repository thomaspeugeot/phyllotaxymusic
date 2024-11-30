import { Component, OnInit, OnDestroy, NgZone } from '@angular/core';
import { MatIconModule } from '@angular/material/icon';
import { MatDividerModule } from '@angular/material/divider';
import { MatButtonModule } from '@angular/material/button';
import { MatSnackBar } from '@angular/material/snack-bar';

import * as Tone from 'tone';
import * as gongtone from '../../../../gongtone/src/public-api';

import { Subject } from 'rxjs';
import { takeUntil, catchError } from 'rxjs/operators';

@Component({
  selector: 'lib-gongtone',
  standalone: true,
  imports: [
    MatButtonModule,
    MatDividerModule,
    MatIconModule
  ],
  templateUrl: './gongtone.component.html',
  styleUrl: './gongtone.component.css'
})
export class GongtoneComponent implements OnInit, OnDestroy {
  private synth: Tone.PolySynth | undefined;
  private sampler: Tone.Sampler | undefined;
  private destroy$ = new Subject<void>();

  readonly StacksNames = gongtone.StacksNames;

  frontRepo?: gongtone.FrontRepo;
  isLoading = false;
  isPlaying = false;

  constructor(
    private frontRepoService: gongtone.FrontRepoService,
    private ngZone: NgZone,
    private snackBar: MatSnackBar
  ) { }

  ngOnInit(): void {
    this.initializeSynth();
    this.connectToWebSocket();
  }

  ngOnDestroy(): void {
    this.destroy$.next();
    this.destroy$.complete();
    this.stopPlayback();
  }

  private initializeSynth(): void {
    try {
      this.synth = new Tone.PolySynth().toDestination();
    } catch (error) {
      this.handleAudioInitError(error);
    }
  }

  private connectToWebSocket(): void {
    this.frontRepoService.connectToWebSocket(this.StacksNames.Gongtone)
      .pipe(
        takeUntil(this.destroy$),
        catchError(error => {
          this.handleWebSocketError(error);
          throw error;
        })
      )
      .subscribe(gongtablesFrontRepo => {
        this.frontRepo = gongtablesFrontRepo;
      });
  }

  stopPlayback(): void {
    this.ngZone.runOutsideAngular(() => {
      try {
        Tone.getTransport().stop();
        this.sampler?.dispose();
        this.isPlaying = false;
      } catch (error) {
        console.error('Error stopping playback:', error);
      }
    });
  }

  play(): void {
    if (!this.frontRepo) {
      this.showError('No data available for playback');
      return;
    }

    this.isLoading = true;
    this.ngZone.runOutsideAngular(() => {
      try {
        const notes = this.frontRepo!.getFrontArray<gongtone.Note>(gongtone.Note.GONGSTRUCT_NAME);
        const duration = this.calculateTotalDuration(notes);
        this.initializeSampler(duration, notes);
      } catch (error) {
        this.handlePlaybackError(error);
      }
    });
  }

  private initializeSampler(duration: number, notes: gongtone.Note[]): void {
    const audioBaseUrl = `${window.location.origin}/assets/audio/salamander/`;

    // Prefer OGG files, fall back to MP3
    const urls: { [key: string]: string } = {
      C3: `${audioBaseUrl}C3.ogg`,
      'D#3': `${audioBaseUrl}Ds3.ogg`,
      'F#3': `${audioBaseUrl}Fs3.ogg`,
      A3: `${audioBaseUrl}A3.ogg`,
      C4: `${audioBaseUrl}C4.ogg`,
      'D#4': `${audioBaseUrl}Ds4.ogg`,
      'F#4': `${audioBaseUrl}Fs4.ogg`,
      A4: `${audioBaseUrl}A4.ogg`,
      C5: `${audioBaseUrl}C5.ogg`,
      'D#5': `${audioBaseUrl}Ds5.ogg`
    };

    this.sampler = new Tone.Sampler({
      urls: urls,
      release: 1,
      onload: () => {
        console.log('Sampler loaded successfully');
        this.startPlayback(duration, notes);
      },
      onerror: (error) => {
        console.error('Sampler load error:', error);

        // Detailed error logging
        if (error instanceof Event) {
          const target = error.target as HTMLAudioElement;
          console.error('Failed Audio Source:', {
            src: target.src,
            error: target.error,
            errorCode: target.error?.code
          });
        }

        this.showError('Failed to load audio samples');
      }
    }).toDestination();
  }

  private calculateTotalDuration(notes: gongtone.Note[]): number {
    return notes.reduce((maxDuration, note) =>
      Math.max(maxDuration, note.Start + note.Duration), 0);
  }

  private startPlayback(duration: number, notes: gongtone.Note[]): void {
    this.ngZone.runOutsideAngular(() => {
      try {
        this.isLoading = false;
        this.isPlaying = true;

        const loop = new Tone.Loop((time) => {
          notes.forEach(note => {
            const frequencies = note.Frequencies.map(freq => freq.Name);
            this.sampler?.triggerAttackRelease(frequencies, note.Duration, time + note.Start);
          });
        }, duration).start(0);

        Tone.getTransport().start();
      } catch (error) {
        this.handlePlaybackError(error);
      }
    });
  }

  private handleWebSocketError(error: any): void {
    console.error('WebSocket connection error:', error);
    this.showError('Failed to connect to WebSocket');
  }

  private handleAudioInitError(error: any): void {
    console.error('Audio initialization error:', error);
    this.showError('Failed to initialize audio');
  }

  private handleSamplerLoadError(error: any): void {
    console.error('Sampler load error:', error);
    this.showError('Failed to load audio samples');
    this.isLoading = false;
  }

  private handlePlaybackError(error: any): void {
    console.error('Playback error:', error);
    this.showError('Playback failed');
    this.isLoading = false;
    this.isPlaying = false;
  }

  private showError(message: string): void {
    this.snackBar.open(message, 'Close', {
      duration: 3000,
      horizontalPosition: 'center',
      verticalPosition: 'top'
    });
  }
}