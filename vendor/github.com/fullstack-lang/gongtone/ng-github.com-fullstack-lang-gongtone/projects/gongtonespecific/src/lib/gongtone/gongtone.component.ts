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
  sampler: Tone.Sampler | undefined


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

    Tone.getTransport().stop()

    // const now = Tone.now()
    // if (this.sampler != undefined) {
    //   this.sampler.releaseAll(now)
    //   console.log("releassed", this.sampler)
    // }
  }

  play() {

    if (this.synth) {

      const now = Tone.now();

      this.sampler = new Tone.Sampler({
        urls: {
          C3: "C3.mp3",
          "D#3": "Ds3.mp3",
          "F#3": "Fs3.mp3",
          A3: "A3.mp3",
          C4: "C4.mp3",
          "D#4": "Ds4.mp3",
          "F#4": "Fs4.mp3",
          A4: "A4.mp3",
          C5: "C5.mp3",
          "D#5": "Ds5.mp3",
        },
        release: 1,
        baseUrl: "https://tonejs.github.io/audio/salamander/",
        onload: () => {

          Tone.getTransport().start()
        }
      }).toDestination();

      //
      // compute full duration of theme
      //
      var duration: number = 0
      for (let note of this.frontRepo!.getFrontArray<gongtone.Note>(gongtone.Note.GONGSTRUCT_NAME)) {
        if (note.Start + note.Duration > duration) {
          duration = note.Start + note.Duration
        }
      }
      console.log("duration", duration)

      const loop = new Tone.Loop((time) => {
        for (let note of this.frontRepo!.getFrontArray<gongtone.Note>(gongtone.Note.GONGSTRUCT_NAME)) {

          var frequencies = new Array<string>()
          for (let freq of note.Frequencies) {
            frequencies.push(freq.Name)
          }

          if (this.sampler != undefined) {
            // console.log(now, now + note.Start, now + note.Start + note.Duration)
            this.sampler.triggerAttackRelease(frequencies, note.Duration, time + note.Start)
            // sampler.triggerAttackRelease(["Eb4", "G4", "Bb4"], 4, 4);
          }
        }
      }, duration).start(0);


    }
  }



}
