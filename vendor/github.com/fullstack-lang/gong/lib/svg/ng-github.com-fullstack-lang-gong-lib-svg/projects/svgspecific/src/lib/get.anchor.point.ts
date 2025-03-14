
import * as svg from '../../../svg/src/public-api'
import { createPoint } from './draw.segments';

type Segment = {
    StartPointWithoutRadius: svg.Point;
    EndPointWithoutRadius: svg.Point;
};

// Calculate the distance between two points
function getDistance(p1: svg.Point, p2: svg.Point): number {
    const dx = p2.X - p1.X;
    const dy = p2.Y - p1.Y;
    return Math.sqrt(dx * dx + dy * dy);
}

// Calculate the position between two points given a ratio
function interpolate(p1: svg.Point, p2: svg.Point, ratio: number): svg.Point {
    return createPoint(
        p1.X + (p2.X - p1.X) * ratio,
        p1.Y + (p2.Y - p1.Y) * ratio,
    )
}

// Function to get the anchor point
export function getAnchorPoint(polyline: Segment[], targetAnchorPosition: number): svg.Point | null {
    // Calculate total length of the polyline
    let totalLength = 0;
    for (let segment of polyline) {
        totalLength += getDistance(segment.StartPointWithoutRadius, segment.EndPointWithoutRadius);
    }

    // Find target distance along the line
    const targetDistance = totalLength * targetAnchorPosition;

    // Iterate through the segments again, tracking the accumulated distance
    let accumulatedDistance = 0;
    for (let segment of polyline) {
        const segmentLength = getDistance(segment.StartPointWithoutRadius, segment.EndPointWithoutRadius);

        if (accumulatedDistance + segmentLength >= targetDistance) {
            // If the target falls within this segment, interpolate the position
            const remainingDistance = targetDistance - accumulatedDistance
            const ratio = remainingDistance / segmentLength
            let anchor = interpolate(segment.StartPointWithoutRadius, segment.EndPointWithoutRadius, ratio)
            if (Number.isNaN(anchor.X)) {
                return null
            }

            return anchor
        } else {
            // Otherwise, add this segment's length to the accumulated distance and continue
            accumulatedDistance += segmentLength;
        }
    }




    // If we've gone through all segments and haven't found the position, return null
    return null;
}
