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


  onButtonClick() {

    if (this.synth) {

      const now = Tone.now();

      for (let note of this.frontRepo!.getFrontArray<gongtone.Note>(gongtone.Note.GONGSTRUCT_NAME)) {

        var frequencies = new Array<string>()
        for (let freq of note.Frequencies) {
          frequencies.push(freq.Name)
        }

        // console.log(now, now + note.Start, now + note.Start + note.Duration)
        this.synth.triggerAttack(frequencies, now + note.Start, note.Velocity)
        this.synth.triggerRelease(frequencies, now + note.Start + note.Duration)

      }


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
