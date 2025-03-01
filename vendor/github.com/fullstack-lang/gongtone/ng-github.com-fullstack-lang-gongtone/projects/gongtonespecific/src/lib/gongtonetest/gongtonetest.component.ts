import { Component } from '@angular/core';

import * as Tone from "tone"


import { MatIconModule } from '@angular/material/icon';
import { MatDividerModule } from '@angular/material/divider';
import { MatButtonModule } from '@angular/material/button';

@Component({
    selector: 'lib-gongtonetest',
    imports: [
        MatButtonModule, MatDividerModule, MatIconModule
    ],
    templateUrl: './gongtonetest.component.html',
    styleUrl: './gongtonetest.component.css'
})
export class GongtonetestComponent {


  playHelloTone() {

    //create a synth and connect it to the main output (your speakers)
    const synth = new Tone.Synth().toDestination();

    //play a middle 'C' for the duration of an 8th note
    synth.triggerAttackRelease("C4", "8n");
  }

  ToneSynth() {

    //create a synth and connect it to the main output (your speakers)
    const synth = new Tone.Synth().toDestination();

    const now = Tone.now();
    // trigger the attack immediately
    synth.triggerAttack("C4", now);
    // wait one second before triggering the release
    synth.triggerRelease(now + 1);
  }

  TriggerAttackRelease() {

    const synth = new Tone.Synth().toDestination();
    const now = Tone.now();
    synth.triggerAttackRelease("C4", "8n", now);
    synth.triggerAttackRelease("E4", "8n", now + 0.5);
    synth.triggerAttackRelease("G4", "8n", now + 1);

    setInterval(() => console.log(Tone.now()), 100);
  }

  Transport() {
    // create two monophonic synths
    const synthA = new Tone.FMSynth().toDestination();
    const synthB = new Tone.AMSynth().toDestination();
    //play a note every quarter-note
    const loopA = new Tone.Loop((time) => {
      synthA.triggerAttackRelease("C2", "8n", time);
    }, "4n").start(0);
    //play another note every off quarter-note, by starting it "8n"
    const loopB = new Tone.Loop((time) => {
      synthB.triggerAttackRelease("C4", "8n", time);
    }, "4n").start("8n");
    // all loops start when the Transport is started
    Tone.getTransport().start();
    // ramp up to 800 bpm over 10 seconds
    Tone.getTransport().bpm.rampTo(800, 10);
  }

  Instruments() {
    const synth = new Tone.PolySynth(Tone.Synth).toDestination();
    const now = Tone.now();
    synth.triggerAttack("D4", now);
    synth.triggerAttack("F4", now + 0.5);
    synth.triggerAttack("A4", now + 1);
    synth.triggerAttack("C5", now + 1.5);
    synth.triggerAttack("E5", now + 2);
    synth.triggerRelease(["D4", "F4", "A4", "C5", "E5"], now + 4);
  }

  Samples() {
    const player = new Tone.Player(
      "https://tonejs.github.io/audio/berklee/gong_1.mp3"
    ).toDestination();
    Tone.loaded().then(() => {
      player.start();
    });
  }

  Sampler() {
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
      baseUrl: "assets/audio/salamander/",
    }).toDestination();

    Tone.loaded().then(() => {
      // sampler.triggerAttackRelease(["Eb4", "G4", "Bb4"], 4);
      sampler.triggerAttackRelease(["Eb3", "G3", "Bb3"], 2);
    });
  }


  Effects() {
    const player = new Tone.Player({
      url: "https://tonejs.github.io/audio/berklee/gurgling_theremin_1.mp3",
      loop: true,
      autostart: true,
    });
    //create a distortion effect
    const distortion = new Tone.Distortion(0.4).toDestination();
    //connect a player to the distortion
    player.connect(distortion);

  }


  Connection() {
    const player = new Tone.Player({
      url: "https://tonejs.github.io/audio/drum-samples/loops/ominous.mp3",
      autostart: true,
    });
    const filter = new Tone.Filter(400, "lowpass").toDestination();
    const feedbackDelay = new Tone.FeedbackDelay(0.125, 0.5).toDestination();

    // connect the player to the feedback delay and filter in parallel
    player.connect(filter);
    player.connect(feedbackDelay);
  }


  Signals() {
    const osc = new Tone.Oscillator().toDestination();
    // start at "C4"
    osc.frequency.value = "C4";
    // ramp to "C2" over 2 seconds
    osc.frequency.rampTo("C2", 2);
    // start the oscillator for 2 seconds
    osc.start().stop("+3");
  }


  ToBeCompleted() {

  }
}
