import { Injectable } from '@angular/core';
import { BehaviorSubject, Observable } from 'rxjs';

import * as phylotaxymusic from '../../../phylotaxymusic/src/public-api'


@Injectable({ providedIn: 'root' })
export class FakeWebsocketService {

    StacksNames = phylotaxymusic.StacksNames

    private positionSubject = new BehaviorSubject<number>(0);

    /**
     * Expose an observable that components can subscribe to.
     */
    public getPositionUpdates(): Observable<number> {
        return this.positionSubject.asObservable();
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
            this.positionSubject.next(progress * endPosition);

            // Keep going until we reach progress=1.0
            if (progress < 1) {
                requestAnimationFrame(animate);
            }
        };

        // Kick off our "animation"
        requestAnimationFrame(animate);
    }
}
