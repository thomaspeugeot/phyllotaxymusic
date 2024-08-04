import { Component, OnInit } from '@angular/core';

import * as Tone from "tone"

import * as gongtone from '../../../../gongtone/src/public-api'


import { MatIconModule } from '@angular/material/icon';
import { MatDividerModule } from '@angular/material/divider';
import { MatButtonModule } from '@angular/material/button';

@Component({
  selector: 'lib-gongtone',
  standalone: true,
  imports: [
    MatButtonModule, MatDividerModule, MatIconModule
  ],
  templateUrl: './gongtone.component.html',
  styleUrl: './gongtone.component.css'
})
export class GongtoneComponent implements OnInit {

  synth: Tone.PolySynth<Tone.Synth<Tone.SynthOptions>> | undefined

  StacksNames = gongtone.StacksNames

  public frontRepo?: gongtone.FrontRepo

  constructor(
    private frontRepoService: gongtone.FrontRepoService,
  ) {

  }

  ngOnInit(): void {
    this.synth = new Tone.PolySynth().toDestination()

    console.log("ngOnInit")

    this.frontRepoService.connectToWebSocket(this.StacksNames.Gongtone).subscribe(
      gongtablesFrontRepo => {
        this.frontRepo = gongtablesFrontRepo
      }
    )
  }

  stop() {

    // transport must be started before it starts invoking events
    Tone.getTransport().stop();
  }

  play() {

    if (this.synth) {

      const now = Tone.now();

      const sampler = new Tone.Sampler({
        urls: {
          C3: "C3.mp3",
          "D#3": "Ds3.mp3",
          "F#3": "Fs3.mp3",
          A3: "A3.mp3",
          C4: "C4.mp3",
          "D#4": "Ds4.mp3",
          "F#4": "Fs4.mp3",
          A4: "A4.mp3",
        },
        release: 1,
        baseUrl: "https://tonejs.github.io/audio/salamander/",
      }).toDestination();

      Tone.loaded().then(() => {
        for (let note of this.frontRepo!.getFrontArray<gongtone.Note>(gongtone.Note.GONGSTRUCT_NAME)) {

          var frequencies = new Array<string>()
          for (let freq of note.Frequencies) {
            frequencies.push(freq.Name)
          }

          // console.log(now, now + note.Start, now + note.Start + note.Duration)
          sampler.triggerAttackRelease(frequencies, note.Duration, now + note.Start)
          // sampler.triggerAttackRelease(["Eb4", "G4", "Bb4"], 4, 4);

        }
      });




      // this.synth.triggerAttack(["D4", "A5"], now, 0.5);
      // this.synth.triggerAttack("F4", now + 0.5);
      // this.synth.triggerAttack("A4", now + 1);
      // this.synth.triggerAttack("C5", now + 1.5);
      // this.synth.triggerAttack("E5", now + 2);
      // this.synth.triggerRelease(["D4", "F4"], now + 8);
      // this.synth.triggerRelease(["A4", "C5", "E5"], now + 4);

      // this.synth.triggerRelease(["D4", "A5"], now + 2);

      Tone.start()

    }
  }



}
